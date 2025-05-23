package main

import (
	"embed"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	var app = cef.NewApplication()
	if common.IsDarwin() {
		app.SetUseMockKeyChain(true)
	}
	cef.BrowserWindow.Config.Title = "Energy - Local load"
	// 本地加载资源方式, 直接读取本地或内置执行文件资源
	// 该模块不使用 http server
	// 默认访问地址fs://energy/index.html, 仅能在应用内访问
	//   fs: 默认的自定义协议名, 你可以任意设置
	//   energy: 默认的自定义域, 你可以任意设置
	//   index.html: 默认打开的页面名，你可以任意设置
	// 页面ajax xhr数据获取
	// xhr数据获取通过Proxy配置, 支持http, https证书配置
	cef.BrowserWindow.Config.Url = "fs://energy" // 设置默认
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		Scheme:     "fs",             // 自定义协议名
		Domain:     "energy",         // 自定义域名
		ResRootDir: "resources/dist", // 资源存放目录, FS不为空时是内置资源目录名, 空时当前文件执行目录, @/to/path @开头表示当前目录下开始
		FS:         resources,        //静态资源所在的 embed.FS
		Proxy: &cef.XHRProxy{ // 页面Ajax XHR请求接口代理转发配置
			Scheme: consts.LpsTls, // http's 支持ssl配置
			IP:     "127.0.0.1",   //http服务ip或domain
			Port:   8040,
			SSL: cef.XHRProxySSL{
				FS:      resources,
				RootDir: "resources/certs",
				Cert:    "client.crt",
				Key:     "client.key",
				CARoots: []string{"ca.crt"},
			},
		},
	}.Build())
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {

	})
	//运行应用
	cef.Run(app)
}
