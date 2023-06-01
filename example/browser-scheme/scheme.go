package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unsafe"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS
var cefApp *cef.TCEFApplication

func main() {
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp = cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com/"
	cef.BrowserWindow.Config.Title = "Energy - Scheme"
	if common.IsLinux() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	// 注册自定义scheme
	cefApp.SetOnRegCustomSchemes(func(registrar *cef.TCefSchemeRegistrarRef) {
		registrar.AddCustomScheme("hello", consts.CEF_SCHEME_OPTION_STANDARD|consts.CEF_SCHEME_OPTION_LOCAL)
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		var idScheme consts.MenuId
		var idClear consts.MenuId
		var idURL consts.MenuId
		// 在右键菜单实现这个示例
		event.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, model *cef.ICefMenuModel) {
			model.AddSeparator()
			idScheme = model.CefMis.NextCommandId()
			model.AddItem(idScheme, "RegScheme")
			idClear = model.CefMis.NextCommandId()
			model.AddItem(idClear, "ClearScheme")
			idURL = model.CefMis.NextCommandId()
			model.AddItem(idURL, "URL")
		})
		// 右键菜单命令
		event.SetOnContextMenuCommand(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, result *bool) {
			if commandId == idScheme {
				// 创建 SchemeHandlerFactory
				factory := cef.SchemeHandlerFactoryRef.New()
				// 创建SchemeHandlerFactory的回调函数
				factory.SetNew(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, schemeName string, request *cef.ICefRequest) *cef.ICefResourceHandler {
					// 创建资源处理器
					handler := cef.ResourceHandlerRef.New(browser, frame, schemeName, request)
					var (
						// 预设一些变量
						fileBytes             []byte
						fileBytesReadPosition = 0
						err                   error
						status                int32 = 0
						statusText                  = ""
						mimeType                    = ""
					)
					// 当加载指定的 scheme 后, 回调以下函数, 按step
					// step 1, ProcessRequest , 处理请求, 在这里预先加载需要的资源和 step 2 需要的响应状态
					handler.ProcessRequest(func(request *cef.ICefRequest, callback *cef.ICefCallback) bool {
						// 默认404
						status = 404
						statusText = "ERROR"
						mimeType = ""
						// 请求地址是我们预告定义好的地址
						if strings.Index(request.URL(), "hello-scheme") != 0 {
							wd, _ := os.Getwd()
							filePath := filepath.Join(wd, "example", "browser-scheme", "resources", "hello-scheme.html")
							fmt.Println(filePath)
							fileBytes, err = ioutil.ReadFile(filePath)
							if err != nil {
								fileBytes = nil
								return false
							}
							fileBytesReadPosition = 0 //每次都将读取位置归0
							// 加载资源成功后设置成功响应状态
							status = 200
							statusText = "OK"
							mimeType = cef.GetMimeType("html") // get html MimeType
							callback.Cont()                    // 继续
							return true
						} else {
							// 返回 false 后不会执行 ReadResponse 回调函数
							fileBytes = nil
							return false
						}
					})
					// step 2, 响应处理器, 将 step 1 的处理结果返回
					handler.GetResponseHeaders(func(response *cef.ICefResponse) (responseLength int64, redirectUrl string) {
						if fileBytes != nil {
							response.SetStatus(status)
							response.SetStatusText(statusText)
							response.SetMimeType(mimeType)
							responseLength = int64(len(fileBytes))
						}
						return
					})
					// step3, 读取响应内容
					handler.ReadResponse(func(dataOut uintptr, bytesToRead int32, callback *cef.ICefCallback) (bytesRead int32, result bool) {
						// 这个函数可能会被多次调用, 如果响应流大于 bytesToRead 时
						// 我们是按块读取 每个块最大bytesToRead且小于实际的要响应的流字节数
						// 从最新的读取位置fileBytesReadPosition把流返回到 dataOut
						if fileBytes != nil && len(fileBytes) > 0 {
							var i int32 = 0 // 默认 0
							// 循环读取字节流内容
							for i < bytesToRead && fileBytesReadPosition < len(fileBytes) {
								// 这里是通过指针地址将赋值, []byte 缓存数组
								// 描述: dataOut byte[i] = fileBytes byte[fileBytesReadPosition]
								*(*byte)(unsafe.Pointer(dataOut + uintptr(i))) = fileBytes[fileBytesReadPosition]
								fileBytesReadPosition++ // 缓存数据的下一个位置, 如果 len(fileBytes) 大于 bytesToRead 时
								i++                     // 计数, 当前最后的输出大小
							}
							// i当前读取的字节数
							return i, i > 0
						}
						return
					})
					return handler
				})
				requestContext := browser.GetRequestContext()
				requestContext.RegisterSchemeHandlerFactory("hello", "", factory)
			} else if commandId == idURL {
				window.Chromium().LoadUrl("hello://hello-scheme")
			} else if commandId == idClear {
				browser.GetRequestContext().ClearSchemeHandlerFactories()
			}
		})
	})
	//运行应用
	cef.Run(cefApp)
}