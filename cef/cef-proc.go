//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common/imports"
)

//--------TCEFWindowParent proc end--------

// 针对 MacOSX 设置命令行参数
//
// 没找到什么好的方式，只能这样设置
func setMacOSXCommandLine(commandLine uintptr) {
	imports.Proc(internale_SetMacOSXCommandLine).Call(commandLine)
}

func AddGoForm(windowId int32, instance uintptr) {
	imports.Proc(internale_CEF_AddGoForm).Call(uintptr(windowId), instance)
}

func RemoveGoForm(windowId int32) {
	imports.Proc(internale_CEF_RemoveGoForm).Call(uintptr(windowId))
}
