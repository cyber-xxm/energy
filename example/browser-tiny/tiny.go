package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS

func main() {
	inits.Init(nil, &resources)
	fmt.Println("main", common.Args.ProcessType())

	config := cef.NewApplicationConfig()
	config.SetMultiThreadedMessageLoop(false)
	config.SetExternalMessagePump(false)
	//config.SetChromeRuntime(true)
	application := cef.NewCEFApplication(config)
	application.SetOnContextCreated(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, context *cef.ICefV8Context) bool {
		fmt.Println("OnContextCreated")
		return false
	})
	application.SetOnContextInitialized(func() {
		fmt.Println("OnContextInitialized()")
		component := lcl.NewComponent(nil)
		chromiumConfig := cef.NewChromiumConfig()
		chromium := cef.NewChromium(component, chromiumConfig)
		browserViewComponent := cef.NewBrowserViewComponent(component)
		windowComponent := cef.NewWindowComponent(component)
		chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
			fmt.Println("OnBeforeClose")
		})
		chromium.SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
			fmt.Println("OnTitleChange", title)
			windowComponent.SetTitle(title)
		})
		chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, client *cef.ICefClient, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup")
			return false
		})

		windowComponent.SetOnWindowCreated(func(sender lcl.IObject, window *cef.ICefWindow) {
			fmt.Println("OnWindowCreated")
			b := chromium.CreateBrowserByBrowserViewComponent("https://www.baidu.com", browserViewComponent)
			fmt.Println("\tCreateBrowserByBrowserViewComponent", b)
			windowComponent.AddChildView(browserViewComponent)
			display := windowComponent.Display()
			fmt.Println("\tdisplay", display, "ClientAreaBoundsInScreen", windowComponent.ClientAreaBoundsInScreen(), display.ID(), display.DeviceScaleFactor())
			fmt.Println("\t", display.Bounds(), display.WorkArea())
			windowComponent.CenterWindow(cef.NewCefSize(1024, 768))
			browserViewComponent.RequestFocus()
			windowComponent.Show()
			windowComponent.SetWindowAppIcon(1, "resources/gitcode.png")
			appIcon := windowComponent.WindowAppIcon()
			fmt.Println("WindowIcon", appIcon.GetHeight())
		})
		windowComponent.SetOnCanClose(func(sender lcl.IObject, window *cef.ICefWindow, aResult *bool) {
			fmt.Println("OnCanClose")
			application.QuitMessageLoop()
			*aResult = true
		})
		windowComponent.SetOnGetInitialBounds(func(sender lcl.IObject, window *cef.ICefWindow, aResult *cef.TCefRect) {
			fmt.Println("OnGetInitialBounds")
		})
		windowComponent.SetOnGetInitialShowState(func(sender lcl.IObject, window *cef.ICefWindow, aResult *consts.TCefShowState) {
			fmt.Println("OnGetInitialShowState", *aResult)
		})
		windowComponent.CreateTopLevelWindow()
	})
	application.SetOnGetDefaultClient(func(client *cef.ICefClient) {
		fmt.Println("OnGetDefaultClient")
	})
	process := application.StartMainProcess()
	fmt.Println("application.StartMainProcess()", process)
	if process {
		fmt.Println("application.RunMessageLoop()")
		application.RunMessageLoop()
	}
}
