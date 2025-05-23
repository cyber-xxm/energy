//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	. "github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/energye/golcl/energy/emfs"
)

const (
	MagicNumber = 0xCEF0CEF0
	Key         = "xxm" // AES-256-GCM
)

// IXHRProxy
//
//	本地资源加载 XHR 请求代理接口
type IXHRProxy interface {
	Send(request *ICefRequest) (*XHRProxyResponse, error) // 被动调用，发送请求，在浏览器进程同步执行
}

// XHRProxy
//
//	数据请求代理
type XHRProxy struct {
	Scheme     LocalProxyScheme // http/https/tcp default: http
	IP         string           // default: localhost
	Port       int              // default: 80
	SSL        XHRProxySSL      // https 安全证书配置
	HttpClient *HttpClient      // http/https 客户端, 可自定义配置
	TcpClient  *TcpClient       // tcp 客户端, 可自定义配置
}

// XHRProxySSL
//
//	https证书配置，如果其中某一配置为空，则跳过ssl检查, 如果证书配置错误则请求失败
type XHRProxySSL struct {
	FS      emfs.IEmbedFS // 证书到内置执行文件时需要设置
	RootDir string        // 根目录 如果使用 FS 时目录名 root/path, 否则本地目录/to/root/path
	Cert    string        // RootDir/to/path/cert.crt
	Key     string        // RootDir/to/path/key.key
	CARoots []string      // RootDir/to/path/ca.crt
}

// HttpClient
//
//	http/https 客户端
type HttpClient struct {
	Transport *http.Transport
	Client    *http.Client
	Jar       *cookiejar.Jar
	Timeout   time.Duration
}

// TcpClient
//
//	tcp 客户端
type TcpClient struct {
	rw              net.Conn
	Dialer          *tls.Dialer
	Timeout         time.Duration
	KeepAlive       time.Duration
	TLSClientConfig *tls.Config
}

// XHRProxyResponse
//
//	代理响应数据
type XHRProxyResponse struct {
	Data       []byte              // 响应数据
	DataSize   int                 // 响应数据大小
	StatusCode int32               // 响应状态码
	Status     string              //
	Header     map[string][]string // 响应头
}

func (m *XHRProxySSL) skipVerify() bool {
	return m.RootDir == "" || m.Cert == "" || m.Key == "" || len(m.CARoots) == 0
}

// Send
//
//	被动调用，发送请求，在浏览器进程同步执行
func (m *XHRProxy) Send(request *ICefRequest) (*XHRProxyResponse, error) {
	if m.Scheme == LpsHttp {
		return m.sendHttp(request)
	} else if m.Scheme == LpsHttps {
		return m.sendHttps(request)
	} else if m.Scheme == LpsTcp {
		return m.sendTcp(request)
	} else if m.Scheme == LpsTls || m.Scheme == LpsTlcp {
		return m.sendTcp(request)
	}
	return nil, errors.New("incorrect scheme")
}

// XHR代理配置
// 如果配置代理，并且是 XHRProxy 时调用
// 否则你可以自己实现代理， 实现 IXHRProxy 接口，自定义代理请求
func (m *XHRProxy) init() {
	if m.Scheme == LpsHttp || m.Scheme == LpsHttps {
		if m.IP == "" {
			m.IP = "localhost"
		}
		if m.HttpClient == nil {
			m.HttpClient = new(HttpClient)
		}
		if m.Scheme == LpsHttps {
			if m.SSL.skipVerify() {
				if m.HttpClient.Transport == nil {
					m.HttpClient.Transport = &http.Transport{
						TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
					}
				}
			} else {
				if m.HttpClient.Transport == nil {
					tlsCfg := loadCert(m)
					m.HttpClient.Transport = &http.Transport{
						TLSClientConfig: tlsCfg,
					}
				}
			}
		}
		if m.HttpClient.Jar == nil {
			if jar, err := cookiejar.New(nil); err == nil {
				m.HttpClient.Jar = jar
			} else {
				println("[Error] XHRProxy SSL New cookiejar:", err.Error())
			}
		}
		if m.HttpClient.Client == nil {
			if m.HttpClient.Timeout <= 0 {
				m.HttpClient.Timeout = time.Second * 30
			}
			m.HttpClient.Client = &http.Client{
				Jar:     m.HttpClient.Jar,
				Timeout: m.HttpClient.Timeout,
			}
		}
		if m.HttpClient.Client.Transport == nil && m.HttpClient.Transport != nil {
			m.HttpClient.Client.Transport = m.HttpClient.Transport
		}
	} else if m.Scheme == LpsTcp || m.Scheme == LpsTls || m.Scheme == LpsTlcp {
		if m.IP == "" {
			m.IP = "localhost"
		}
		if m.TcpClient == nil {
			m.TcpClient = new(TcpClient)
		}
		if m.TcpClient.Dialer == nil {
			if m.TcpClient.Timeout <= 0 {
				m.TcpClient.Timeout = time.Second * 10
			}
			if m.TcpClient.KeepAlive <= 0 {
				m.TcpClient.KeepAlive = time.Second * 30
			}
			// 声明一个 Dialer 实例，并设置一些可选参数
			m.TcpClient.Dialer = &tls.Dialer{
				NetDialer: &net.Dialer{
					Timeout:   m.TcpClient.Timeout,   // 设置拨号超时时间
					KeepAlive: m.TcpClient.KeepAlive, // 设置保持连接的时间
				},
			}
		}
		if m.Scheme == LpsTls || m.Scheme == LpsTlcp {
			if m.SSL.skipVerify() {
				m.TcpClient.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
			} else {
				m.TcpClient.TLSClientConfig = loadCert(m)
			}
			m.TcpClient.Dialer.Config = m.TcpClient.TLSClientConfig
		}
	}
}

func (m *XHRProxy) sendHttp(request *ICefRequest) (*XHRProxyResponse, error) {
	return m.send("http://", request)
}

func (m *XHRProxy) sendHttps(request *ICefRequest) (*XHRProxyResponse, error) {
	return m.send("https://", request)
}

func (m *XHRProxy) send(scheme string, request *ICefRequest) (*XHRProxyResponse, error) {
	reqUrl, err := url.Parse(request.URL())
	if err != nil {
		return nil, err
	}
	// 构造目标地址
	targetUrl := new(bytes.Buffer)
	targetUrl.WriteString(scheme)
	targetUrl.WriteString(m.IP)
	if m.Port > 0 { // ip:port
		targetUrl.WriteString(":")
		targetUrl.WriteString(strconv.Itoa(m.Port))
	}
	targetUrl.WriteString(reqUrl.Path)
	if reqUrl.RawQuery != "" {
		targetUrl.WriteString("?")
		targetUrl.WriteString(reqUrl.RawQuery)
	}
	// 读取请求数据
	requestData := new(bytes.Buffer)
	data := request.GetPostData()
	if data.IsValid() {
		dataCount := int(data.GetElementCount())
		elements := data.GetElements()
		for i := 0; i < dataCount; i++ {
			element := elements.Get(uint32(i))
			switch element.GetType() {
			case PDE_TYPE_EMPTY:
			case PDE_TYPE_BYTES:
				if byt, c := element.GetBytes(); c > 0 {
					requestData.Write(byt)
				}
			case PDE_TYPE_FILE:
				if f := element.GetFile(); f != "" {
					if byt, err := os.ReadFile(f); err == nil {
						requestData.Write(byt)
					}
				}
			}
			element.Free()
		}
		data.Free()
	}
	tarUrl := targetUrl.String()
	if logger.Enable() {
		logger.Debug("XHRProxy URL:", tarUrl, "method:", request.Method(), "data-size:", requestData.Len())
	}
	httpRequest, err := http.NewRequest(request.Method(), tarUrl, requestData)
	if err != nil {
		return nil, err
	}
	// 设置请求头
	header := request.GetHeaderMap()
	if header.IsValid() {
		size := header.GetSize()
		for i := 0; i < int(size); i++ {
			key := header.GetKey(uint32(i))
			c := header.FindCount(key)
			for j := 0; j < int(c); j++ {
				value := header.GetEnumerate(key, uint32(j))
				httpRequest.Header.Add(key, value)
			}
		}
		header.Free()
	}
	//httpRequest.Header.Add("Host", "www.example.com")
	//httpRequest.Header.Add("Origin", "https://www.example.com")
	//httpRequest.Header.Add("Referer", "https://www.example.com/")
	if m.HttpClient.Client == nil {
		return nil, errors.New("http client is nil")
	}
	httpResponse, err := m.HttpClient.Client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	// 读取响应头
	responseHeader := make(map[string][]string)
	for key, value := range httpResponse.Header {
		for _, vs := range value {
			if header, ok := responseHeader[key]; ok {
				responseHeader[key] = append(header, vs)
			} else {
				responseHeader[key] = []string{vs}
			}
		}
	}
	// 读取响应数据
	buf := new(bytes.Buffer)
	c, err := buf.ReadFrom(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	status := "OK"
	if httpResponse.StatusCode != 200 {
		rs := strings.Split(httpResponse.Status, " ")
		if len(rs) > 1 {
			status = rs[1]
		} else {
			status = httpResponse.Status
		}
	}
	result := &XHRProxyResponse{
		Data:       buf.Bytes(),
		DataSize:   int(c),
		StatusCode: int32(httpResponse.StatusCode),
		Status:     status,
		Header:     responseHeader,
	}
	return result, nil
}

func (m *XHRProxy) sendTcp(request *ICefRequest) (*XHRProxyResponse, error) {
	scheme := "http://"
	reqUrl, err := url.Parse(request.URL())
	if err != nil {
		return nil, err
	}
	// 构造目标地址
	targetUrl := new(bytes.Buffer)
	targetUrl.WriteString(scheme)
	targetUrl.WriteString(m.IP)
	if m.Port > 0 { // ip:port
		targetUrl.WriteString(":")
		targetUrl.WriteString(strconv.Itoa(m.Port))
	} else {
		targetUrl.WriteString(":")
		targetUrl.WriteString(strconv.Itoa(8081))
	}
	targetUrl.WriteString(reqUrl.Path)
	if reqUrl.RawQuery != "" {
		targetUrl.WriteString("?")
		targetUrl.WriteString(reqUrl.RawQuery)
	}
	// 读取请求数据
	requestData := new(bytes.Buffer)
	data := request.GetPostData()
	if data.IsValid() {
		dataCount := int(data.GetElementCount())
		elements := data.GetElements()
		for i := 0; i < dataCount; i++ {
			element := elements.Get(uint32(i))
			switch element.GetType() {
			case PDE_TYPE_EMPTY:
			case PDE_TYPE_BYTES:
				if byt, c := element.GetBytes(); c > 0 {
					requestData.Write(byt)
				}
			case PDE_TYPE_FILE:
				if f := element.GetFile(); f != "" {
					if byt, err := os.ReadFile(f); err == nil {
						requestData.Write(byt)
					}
				}
			}
			element.Free()
		}
		data.Free()
	}
	tarUrl := targetUrl.String()
	if logger.Enable() {
		logger.Debug("XHRProxy URL:", tarUrl, "method:", request.Method(), "data-size:", requestData.Len())
	}
	httpRequest, err := http.NewRequest(request.Method(), tarUrl, requestData)
	if err != nil {
		return nil, err
	}
	// 设置请求头
	header := request.GetHeaderMap()
	if header.IsValid() {
		size := header.GetSize()
		for i := 0; i < int(size); i++ {
			key := header.GetKey(uint32(i))
			c := header.FindCount(key)
			for j := 0; j < int(c); j++ {
				value := header.GetEnumerate(key, uint32(j))
				httpRequest.Header.Add(key, value)
			}
		}
		header.Free()
	}
	if m.TcpClient.Dialer != nil && m.TcpClient.rw == nil {
		fmt.Println("tcp client is nil, reconnecting...")
		for i := 0; i < 10; i++ {
			fmt.Println(fmt.Sprintf("tcp client is nil, reconnecting count %d...", i+1))
			conn, err := m.TcpClient.Dialer.Dial("tcp", fmt.Sprintf("%s:%d", m.IP, m.Port))
			if err == nil {
				m.TcpClient.rw = conn
				fmt.Println("tcp client connect success")
				break // 连接成功，返回连接对象
			}
			fmt.Println(fmt.Sprintf("[Error] XHRProxy TCP Dial: %s, trying reconnect", err))
			// time.Sleep(10 * time.Second) // 等待一段时间后重试
		}
		if m.TcpClient.rw == nil {
			return nil, errors.New("tcp client is nil")
		}
	} else if m.TcpClient.Dialer == nil {
		panic("tcp dialer config no configured")
	}
	fmt.Println("tcp client start send message")
	httpResponse, err := m.sendHttpRequestOverTcp(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()
	// 读取响应头
	responseHeader := make(map[string][]string)
	for key, value := range httpResponse.Header {
		for _, vs := range value {
			if header, ok := responseHeader[key]; ok {
				responseHeader[key] = append(header, vs)
			} else {
				responseHeader[key] = []string{vs}
			}
		}
	}
	// 读取响应数据
	buf := new(bytes.Buffer)
	c, err := buf.ReadFrom(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	status := "OK"
	if httpResponse.StatusCode != 200 {
		rs := strings.Split(httpResponse.Status, " ")
		if len(rs) > 1 {
			status = rs[1]
		} else {
			status = httpResponse.Status
		}
	}
	result := &XHRProxyResponse{
		Data:       buf.Bytes(),
		DataSize:   int(c),
		StatusCode: int32(httpResponse.StatusCode),
		Status:     status,
		Header:     responseHeader,
	}
	return result, nil
}

func (m *XHRProxy) sendHttpRequestOverTcp(request *http.Request) (*http.Response, error) {
	requestData, err := requestToHTTP(request)
	if err != nil {
		println("[Error] XHRProxy TCP Read Request: ", err.Error())
		return nil, errors.New("http格式错误")
	}
	// 加密数据
	// encryptedData, iv, err := encryptData(requestData)
	// if err != nil {
	// 	println("加密失败: ", err)
	// }

	// if m.TcpClient.Client == nil {
	// 	return nil, errors.New("tcp client is nil")
	// }

	writer := bufio.NewWriter(m.TcpClient.rw)
	// 构建报文头
	// header := make([]byte, 8)
	// binary.LittleEndian.PutUint32(header[0:4], MagicNumber)
	// binary.LittleEndian.PutUint32(header[4:8], uint32(len(encryptedData)+len(iv)))

	// 组合完整报文
	// fullPacket := append(header, iv...)
	// fullPacket = append(fullPacket, encryptedData...)
	fmt.Println("tcp client start write data")
	wl, err := writer.Write(requestData)
	fmt.Println(fmt.Sprintf("tcp client start write data done: %d", wl))
	if err != nil {
		println("[Error] XHRProxy TCP Write: ", err.Error())
		return nil, err
	}
	if wl <= 0 {
		println("[Error] XHRProxy TCP Write Data Zero: ", err.Error())
		return nil, errors.New("写入数据失败，长度为0")
	}
	writer.Flush()
	fmt.Println("tcp client start read response")
	// 使用bufio.Reader包装conn，以便逐行读取数据
	reader := bufio.NewReader(m.TcpClient.rw)

	// 读取报文头（8字节）
	// header = make([]byte, 8)
	// if _, err = io.ReadFull(reader, header); err != nil {
	// 	println("读取头失败:", err)
	// 	return nil, err
	// }

	// 验证魔数
	// magic := binary.LittleEndian.Uint32(header[0:4])
	// if magic != MagicNumber {
	// 	println("无效协议")
	// 	return nil, errors.New("无效协议")
	// }

	// 获取数据长度
	//dataLength := binary.LittleEndian.Uint32(header[4:8])

	// 读取IV+密文
	// data := make([]byte, dataLength)
	// if _, err = io.ReadFull(reader, data); err != nil {
	// 	println("读取数据失败:", err)
	// 	return nil, err
	// }

	// 解密数据
	// plaintext, err := decryptData(data)
	// if err != nil {
	// 	println("解密失败:", err)
	// 	return nil, err
	// }

	// 将收集到的数据转换为http.Response
	resp, err := http.ReadResponse(reader, request)
	fmt.Println("tcp client read response done")
	if err != nil {
		fmt.Println("Error parsing response:", err.Error())
		return nil, err
	}
	return resp, nil
}

// 将 http.Request 转换为符合 HTTP 协议格式的字符串
// requestToHTTP 将 *http.Request 转换为符合 TCP 传输格式的 HTTP 请求字节数组，并添加 Content-Length 头。
func requestToHTTP(req *http.Request) ([]byte, error) {
	var buf bytes.Buffer

	// 如果 Request 的 URL 是相对路径，需要构造完整的 URL 路径
	uri := req.URL.RequestURI()

	// 写请求行: METHOD PATH HTTP/1.1\r\n
	_, err := fmt.Fprintf(&buf, "%s %s HTTP/1.1\r\n", req.Method, uri)
	if err != nil {
		return nil, err
	}

	// 确保 Host 头存在
	if req.Host != "" {
		req.Header.Set("Host", req.Host)
	} else if req.URL.Host != "" {
		req.Header.Set("Host", req.URL.Host)
	}

	// 计算请求体长度（如果有 body）
	var bodyBytes []byte
	var contentLength int
	if req.Body != nil {
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		// Reset body 以便复用 request
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		contentLength = len(bodyBytes)
	}

	// 如果请求体不为空，设置 Content-Length
	if contentLength > 0 {
		req.Header.Set("Content-Length", fmt.Sprintf("%d", contentLength))
	}

	// 写 headers
	err = req.Header.Write(&buf)
	if err != nil {
		return nil, err
	}

	// 写空行分隔 headers 和 body
	_, err = buf.Write([]byte("\r\n"))
	if err != nil {
		return nil, err
	}

	// 写 body（如果有）
	if contentLength > 0 {
		_, err = buf.Write(bodyBytes)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func loadCert(m *XHRProxy) *tls.Config {
	var (
		certPEMBlock, keyPEMBlock []byte
		err                       error
	)
	var readFile = func(path string) (data []byte, err error) {
		if m.SSL.FS != nil {
			path = strings.Replace(filepath.Join(m.SSL.RootDir, path), "\\", "/", -1)
			data, err = m.SSL.FS.ReadFile(path)
		} else {
			path = filepath.Join(m.SSL.RootDir, path)
			data, err = os.ReadFile(path)
		}
		return
	}
	certPEMBlock, err = readFile(m.SSL.Cert)
	if err != nil {
		panic(err)
		return nil
	}
	keyPEMBlock, err = readFile(m.SSL.Key)
	if err != nil {
		panic(err)
		return nil
	}
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		panic(err)
		return nil
	}
	pool := x509.NewCertPool()
	for _, path := range m.SSL.CARoots {
		if ca, err := readFile(path); err == nil {
			pool.AppendCertsFromPEM(ca)
		} else {
			panic(err)
		}
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
		MinVersion:   tls.VersionTLS12,
	}
}
