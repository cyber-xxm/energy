//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 应用程序
package cef

import (
	"fmt"
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"strings"
	"unsafe"
)

var application *TCEFApplication

// CEF应用对象
type TCEFApplication struct {
	instance unsafe.Pointer
	cfg      *tCefApplicationConfig
}

// 创建CEF应用程序
func NewCEFApplication(cfg *tCefApplicationConfig) *TCEFApplication {
	if application != nil {
		return application
	}
	if cfg == nil {
		cfg = NewApplicationConfig()
	}
	cfg.framework()
	application = new(TCEFApplication)
	r1, _, _ := imports.Proc(internale_CEFApplication_Create).Call(uintptr(unsafe.Pointer(cfg)))
	application.instance = unsafe.Pointer(r1)
	application.cfg = cfg

	return application
}

// 创建应用程序
//
// 带有默认的应用事件
func NewApplication(cfg *tCefApplicationConfig) *TCEFApplication {
	cefApp := NewCEFApplication(cfg)
	cefApp.registerDefaultEvent()
	return cefApp
}

func (m *TCEFApplication) registerDefaultEvent() {
	//register default function
	m.defaultSetOnContextCreated()
	m.defaultSetOnProcessMessageReceived()
	m.defaultSetOnBeforeChildProcessLaunch()
}

func (m *TCEFApplication) Instance() uintptr {
	return uintptr(m.instance)
}

// 启动主进程
func (m *TCEFApplication) StartMainProcess() bool {
	if m.instance != nullptr {
		logger.Debug("application single exe,", common.Args.ProcessType(), "process start")
		r1, _, _ := imports.Proc(internale_CEFStartMainProcess).Call(m.Instance())
		return api.GoBool(r1)
	}
	return false
}

// 启动子进程, 如果指定了子进程执行程序, 将执行指定的子进程程序
func (m *TCEFApplication) StartSubProcess() (result bool) {
	if m.instance != nullptr {
		logger.Debug("application multiple exe,", common.Args.ProcessType(), "process start")
		r1, _, _ := imports.Proc(internale_CEFStartSubProcess).Call(m.Instance())
		result = api.GoBool(r1)
	}
	return false
}

func (m *TCEFApplication) RunMessageLoop() {
	defer func() {
		logger.Debug("application run message loop end")
		api.EnergyLibRelease()
	}()
	logger.Debug("application run message loop start")
	imports.Proc(internale_CEFApplication_RunMessageLoop).Call()
}

func (m *TCEFApplication) QuitMessageLoop() {
	logger.Debug("application quit message loop")
	imports.Proc(internale_CEFApplication_QuitMessageLoop).Call()
}

func (m *TCEFApplication) StopScheduler() {
	imports.Proc(internale_CEFApplication_StopScheduler).Call()
}

func (m *TCEFApplication) Destroy() {
	imports.Proc(internale_CEFApplication_Destroy).Call()
}

func (m *TCEFApplication) Free() {
	if m.instance != nullptr {
		imports.Proc(internale_CEFApplication_Free).Call()
		m.instance = nullptr
	}
}

func (m *TCEFApplication) ExecuteJS(browserId int32, code string) {
	imports.Proc(internale_CEFApplication_ExecuteJS).Call()
}

// 上下文件创建回调
//
// 返回值 false 将会创建 render进程的IPC和GO绑定变量
//
// 对于一些不想GO绑定变量的URL地址，实现该函数，通过 frame.Url
func (m *TCEFApplication) SetOnContextCreated(fn GlobalCEFAppEventOnContextCreated) {
	imports.Proc(internale_CEFGlobalApp_SetOnContextCreated).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) defaultSetOnContextCreated() {
	m.SetOnContextCreated(func(browse *ICefBrowser, frame *ICefFrame, context *ICefV8Context) bool {
		return false
	})
}

func (m *TCEFApplication) SetOnContextInitialized(fn GlobalCEFAppEventOnContextInitialized) {
	imports.Proc(internale_CEFGlobalApp_SetOnContextInitialized).Call(api.MakeEventDataPtr(fn))
}

// 初始化设置全局回调
func (m *TCEFApplication) SetOnWebKitInitialized(fn GlobalCEFAppEventOnWebKitInitialized) {
	imports.Proc(internale_CEFGlobalApp_SetOnWebKitInitialized).Call(api.MakeEventDataPtr(fn))
}

// 进程间通信处理消息接收
func (m *TCEFApplication) SetOnProcessMessageReceived(fn RenderProcessMessageReceived) {
	imports.Proc(internale_CEFGlobalApp_SetOnProcessMessageReceived).Call(api.MakeEventDataPtr(fn))
}
func (m *TCEFApplication) defaultSetOnProcessMessageReceived() {
	m.SetOnProcessMessageReceived(func(browse *ICefBrowser, frame *ICefFrame, sourceProcess CefProcessId, processMessage *ipc.ICefProcessMessage) bool {
		return false
	})
}

func (m *TCEFApplication) AddCustomCommandLine(commandLine, value string) {
	imports.Proc(internale_AddCustomCommandLine).Call(api.PascalStr(commandLine), api.PascalStr(value))
}

// 启动子进程之前自定义命令行参数设置
func (m *TCEFApplication) SetOnBeforeChildProcessLaunch(fn GlobalCEFAppEventOnBeforeChildProcessLaunch) {
	imports.Proc(internale_CEFGlobalApp_SetOnBeforeChildProcessLaunch).Call(api.MakeEventDataPtr(fn))
}
func (m *TCEFApplication) defaultSetOnBeforeChildProcessLaunch() {
	m.SetOnBeforeChildProcessLaunch(func(commandLine *TCefCommandLine) {})
}

func (m *TCEFApplication) SetOnBrowserDestroyed(fn GlobalCEFAppEventOnBrowserDestroyed) {
	imports.Proc(internale_CEFGlobalApp_SetOnBrowserDestroyed).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadStart(fn GlobalCEFAppEventOnRenderLoadStart) {
	imports.Proc(internale_CEFGlobalApp_SetOnLoadStart).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadEnd(fn GlobalCEFAppEventOnRenderLoadEnd) {
	imports.Proc(internale_CEFGlobalApp_SetOnLoadEnd).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadError(fn GlobalCEFAppEventOnRenderLoadError) {
	imports.Proc(internale_CEFGlobalApp_SetOnLoadError).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnLoadingStateChange(fn GlobalCEFAppEventOnRenderLoadingStateChange) {
	imports.Proc(internale_CEFGlobalApp_SetOnLoadingStateChange).Call(api.MakeEventDataPtr(fn))
}

func (m *TCEFApplication) SetOnGetDefaultClient(fn GlobalCEFAppEventOnGetDefaultClient) {
	imports.Proc(internale_CEFGlobalApp_SetOnGetDefaultClient).Call(api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case GlobalCEFAppEventOnBrowserDestroyed:
			fn.(GlobalCEFAppEventOnBrowserDestroyed)(&ICefBrowser{instance: getPtr(0)})
		case GlobalCEFAppEventOnRenderLoadStart:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadStart)(browser, frame, TCefTransitionType(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadEnd:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadEnd)(browser, frame, int32(getVal(2)))
		case GlobalCEFAppEventOnRenderLoadError:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadError)(browser, frame, TCefErrorCode(getVal(2)), api.GoStr(getVal(3)), api.GoStr(getVal(4)))
		case GlobalCEFAppEventOnRenderLoadingStateChange:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{getPtr(1)}
			fn.(GlobalCEFAppEventOnRenderLoadingStateChange)(browser, frame, api.GoBool(getVal(2)), api.GoBool(getVal(3)), api.GoBool(getVal(4)))
		case RenderProcessMessageReceived:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{getPtr(1)}
			cefProcMsg := (*ipc.CefProcessMessagePtr)(getPtr(3))
			args := ipc.NewArgumentList()
			args.UnPackageBytePtr(cefProcMsg.Data, int32(cefProcMsg.DataLen))
			processMessage := &ipc.ICefProcessMessage{
				Name:         api.GoStr(cefProcMsg.Name),
				ArgumentList: args,
			}
			var result = (*bool)(getPtr(4))
			*result = fn.(RenderProcessMessageReceived)(browser, frame, CefProcessId(getVal(2)), processMessage)
			args.Clear()
			cefProcMsg.Data = 0
			cefProcMsg.DataLen = 0
			cefProcMsg.Name = 0
			cefProcMsg = nil
			args = nil
		case GlobalCEFAppEventOnContextCreated:
			browser := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			if strings.Index(frame.Url(), "devtools://") == 0 {
				processName = common.PT_DEVTOOLS
				return true
			} else {
				processName = common.Args.ProcessType()
			}
			ctx := &ICefV8Context{instance: getPtr(2)}
			browser.Identifier()
			var status = (*bool)(getPtr(3))
			//用户定义返回 false 创建 render IPC 及 变量绑定
			var result = fn.(GlobalCEFAppEventOnContextCreated)(browser, frame, ctx)
			if !result {
				cefAppContextCreated(browser, frame)
				*status = true
			} else {
				*status = false
			}
		case GlobalCEFAppEventOnWebKitInitialized:
			fn.(GlobalCEFAppEventOnWebKitInitialized)()
		case GlobalCEFAppEventOnContextInitialized:
			fn.(GlobalCEFAppEventOnContextInitialized)()
		case GlobalCEFAppEventOnBeforeChildProcessLaunch:
			commands := (*uintptr)(getPtr(0))
			commandLine := &TCefCommandLine{commandLines: make(map[string]string)}
			ipc.IPC.SetPort()
			commandLine.AppendSwitch(MAINARGS_NETIPCPORT, fmt.Sprintf("%d", ipc.IPC.Port()))
			fn.(GlobalCEFAppEventOnBeforeChildProcessLaunch)(commandLine)
			*commands = api.PascalStr(commandLine.toString())
		case GlobalCEFAppEventOnGetDefaultClient:
			client := (*uintptr)(getPtr(0))
			getClient := &ICefClient{instance: unsafe.Pointer(client)}
			fn.(GlobalCEFAppEventOnGetDefaultClient)(getClient)
			*client = uintptr(getClient.instance)
		default:
			return false
		}
		return true
	})
}
