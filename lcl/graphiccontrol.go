//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IGraphicControl Parent: IControl
type IGraphicControl interface {
	IControl
	Canvas() ICanvas // property
}

// TGraphicControl Parent: TControl
type TGraphicControl struct {
	TControl
}

func NewGraphicControl(AOwner IComponent) IGraphicControl {
	r1 := LCL().SysCallN(2931, GetObjectUintptr(AOwner))
	return AsGraphicControl(r1)
}

func (m *TGraphicControl) Canvas() ICanvas {
	r1 := LCL().SysCallN(2929, m.Instance())
	return AsCanvas(r1)
}

func GraphicControlClass() TClass {
	ret := LCL().SysCallN(2930)
	return TClass(ret)
}
