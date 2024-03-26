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

// IToggleBox Parent: ICustomCheckBox
type IToggleBox interface {
	ICustomCheckBox
	Checked() bool                                 // property
	SetChecked(AValue bool)                        // property
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TToggleBox Parent: TCustomCheckBox
type TToggleBox struct {
	TCustomCheckBox
	contextPopupPtr   uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
}

func NewToggleBox(TheOwner IComponent) IToggleBox {
	r1 := LCL().SysCallN(4777, GetObjectUintptr(TheOwner))
	return AsToggleBox(r1)
}

func (m *TToggleBox) Checked() bool {
	r1 := LCL().SysCallN(4775, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToggleBox) SetChecked(AValue bool) {
	LCL().SysCallN(4775, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToggleBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(4778, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TToggleBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4778, 1, m.Instance(), uintptr(AValue))
}

func (m *TToggleBox) DragKind() TDragKind {
	r1 := LCL().SysCallN(4779, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TToggleBox) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4779, 1, m.Instance(), uintptr(AValue))
}

func (m *TToggleBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(4780, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TToggleBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4780, 1, m.Instance(), uintptr(AValue))
}

func (m *TToggleBox) ParentFont() bool {
	r1 := LCL().SysCallN(4781, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToggleBox) SetParentFont(AValue bool) {
	LCL().SysCallN(4781, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToggleBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(4782, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToggleBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4782, 1, m.Instance(), PascalBool(AValue))
}

func ToggleBoxClass() TClass {
	ret := LCL().SysCallN(4776)
	return TClass(ret)
}

func (m *TToggleBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4783, m.Instance(), m.contextPopupPtr)
}

func (m *TToggleBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4784, m.Instance(), m.dragDropPtr)
}

func (m *TToggleBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4785, m.Instance(), m.dragOverPtr)
}

func (m *TToggleBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4786, m.Instance(), m.endDragPtr)
}

func (m *TToggleBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4787, m.Instance(), m.mouseDownPtr)
}

func (m *TToggleBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4788, m.Instance(), m.mouseEnterPtr)
}

func (m *TToggleBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4789, m.Instance(), m.mouseLeavePtr)
}

func (m *TToggleBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4790, m.Instance(), m.mouseMovePtr)
}

func (m *TToggleBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4791, m.Instance(), m.mouseUpPtr)
}

func (m *TToggleBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4792, m.Instance(), m.mouseWheelPtr)
}

func (m *TToggleBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4793, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TToggleBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4794, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TToggleBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4795, m.Instance(), m.startDragPtr)
}
