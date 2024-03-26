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
	"unsafe"
)

// IFindDialog Parent: ICommonDialog
type IFindDialog interface {
	ICommonDialog
	Left() int32                      // property
	SetLeft(AValue int32)             // property
	Position() (resultPoint TPoint)   // property
	SetPosition(AValue *TPoint)       // property
	Top() int32                       // property
	SetTop(AValue int32)              // property
	FindText() string                 // property
	SetFindText(AValue string)        // property
	Options() TFindOptions            // property
	SetOptions(AValue TFindOptions)   // property
	CloseDialog()                     // procedure
	SetOnFind(fn TNotifyEvent)        // property event
	SetOnHelpClicked(fn TNotifyEvent) // property event
}

// TFindDialog Parent: TCommonDialog
type TFindDialog struct {
	TCommonDialog
	findPtr        uintptr
	helpClickedPtr uintptr
}

func NewFindDialog(AOwner IComponent) IFindDialog {
	r1 := LCL().SysCallN(2767, GetObjectUintptr(AOwner))
	return AsFindDialog(r1)
}

func (m *TFindDialog) Left() int32 {
	r1 := LCL().SysCallN(2769, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFindDialog) SetLeft(AValue int32) {
	LCL().SysCallN(2769, 1, m.Instance(), uintptr(AValue))
}

func (m *TFindDialog) Position() (resultPoint TPoint) {
	LCL().SysCallN(2771, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TFindDialog) SetPosition(AValue *TPoint) {
	LCL().SysCallN(2771, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TFindDialog) Top() int32 {
	r1 := LCL().SysCallN(2774, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFindDialog) SetTop(AValue int32) {
	LCL().SysCallN(2774, 1, m.Instance(), uintptr(AValue))
}

func (m *TFindDialog) FindText() string {
	r1 := LCL().SysCallN(2768, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFindDialog) SetFindText(AValue string) {
	LCL().SysCallN(2768, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFindDialog) Options() TFindOptions {
	r1 := LCL().SysCallN(2770, 0, m.Instance(), 0)
	return TFindOptions(r1)
}

func (m *TFindDialog) SetOptions(AValue TFindOptions) {
	LCL().SysCallN(2770, 1, m.Instance(), uintptr(AValue))
}

func FindDialogClass() TClass {
	ret := LCL().SysCallN(2765)
	return TClass(ret)
}

func (m *TFindDialog) CloseDialog() {
	LCL().SysCallN(2766, m.Instance())
}

func (m *TFindDialog) SetOnFind(fn TNotifyEvent) {
	if m.findPtr != 0 {
		RemoveEventElement(m.findPtr)
	}
	m.findPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2772, m.Instance(), m.findPtr)
}

func (m *TFindDialog) SetOnHelpClicked(fn TNotifyEvent) {
	if m.helpClickedPtr != 0 {
		RemoveEventElement(m.helpClickedPtr)
	}
	m.helpClickedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2773, m.Instance(), m.helpClickedPtr)
}
