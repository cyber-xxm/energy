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

// IComponent Parent: IPersistent
type IComponent interface {
	IPersistent
	Components(Index int32) IComponent            // property
	ComponentCount() int32                        // property
	ComponentIndex() int32                        // property
	SetComponentIndex(AValue int32)               // property
	ComponentState() TComponentStates             // property
	ComponentStyle() TComponentStyles             // property
	DesignInfo() int32                            // property
	SetDesignInfo(AValue int32)                   // property
	Owner() IComponent                            // property
	VCLComObject() uintptr                        // property
	SetVCLComObject(AValue uintptr)               // property
	Name() string                                 // property
	SetName(AValue string)                        // property
	Tag() uint32                                  // property
	SetTag(AValue uint32)                         // property
	ExecuteAction(Action IBasicAction) bool       // function
	FindComponent(AName string) IComponent        // function
	GetEnumerator() IComponentEnumerator          // function
	GetParentComponent() IComponent               // function
	HasParent() bool                              // function
	UpdateAction(Action IBasicAction) bool        // function
	DestroyComponents()                           // procedure
	Destroying()                                  // procedure
	FreeNotification(AComponent IComponent)       // procedure
	RemoveFreeNotification(AComponent IComponent) // procedure
	FreeOnRelease()                               // procedure
	InsertComponent(AComponent IComponent)        // procedure
	RemoveComponent(AComponent IComponent)        // procedure
	SetSubComponent(ASubComponent bool)           // procedure
}

// TComponent Parent: TPersistent
type TComponent struct {
	TPersistent
}

func NewComponent(AOwner IComponent) IComponent {
	r1 := LCL().SysCallN(695, GetObjectUintptr(AOwner))
	return AsComponent(r1)
}

func (m *TComponent) Components(Index int32) IComponent {
	r1 := LCL().SysCallN(694, m.Instance(), uintptr(Index))
	return AsComponent(r1)
}

func (m *TComponent) ComponentCount() int32 {
	r1 := LCL().SysCallN(690, m.Instance())
	return int32(r1)
}

func (m *TComponent) ComponentIndex() int32 {
	r1 := LCL().SysCallN(691, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComponent) SetComponentIndex(AValue int32) {
	LCL().SysCallN(691, 1, m.Instance(), uintptr(AValue))
}

func (m *TComponent) ComponentState() TComponentStates {
	r1 := LCL().SysCallN(692, m.Instance())
	return TComponentStates(r1)
}

func (m *TComponent) ComponentStyle() TComponentStyles {
	r1 := LCL().SysCallN(693, m.Instance())
	return TComponentStyles(r1)
}

func (m *TComponent) DesignInfo() int32 {
	r1 := LCL().SysCallN(696, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComponent) SetDesignInfo(AValue int32) {
	LCL().SysCallN(696, 1, m.Instance(), uintptr(AValue))
}

func (m *TComponent) Owner() IComponent {
	r1 := LCL().SysCallN(708, m.Instance())
	return AsComponent(r1)
}

func (m *TComponent) VCLComObject() uintptr {
	r1 := LCL().SysCallN(714, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TComponent) SetVCLComObject(AValue uintptr) {
	LCL().SysCallN(714, 1, m.Instance(), uintptr(AValue))
}

func (m *TComponent) Name() string {
	r1 := LCL().SysCallN(707, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TComponent) SetName(AValue string) {
	LCL().SysCallN(707, 1, m.Instance(), PascalStr(AValue))
}

func (m *TComponent) Tag() uint32 {
	r1 := LCL().SysCallN(712, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TComponent) SetTag(AValue uint32) {
	LCL().SysCallN(712, 1, m.Instance(), uintptr(AValue))
}

func (m *TComponent) ExecuteAction(Action IBasicAction) bool {
	r1 := LCL().SysCallN(699, m.Instance(), GetObjectUintptr(Action))
	return GoBool(r1)
}

func (m *TComponent) FindComponent(AName string) IComponent {
	r1 := LCL().SysCallN(700, m.Instance(), PascalStr(AName))
	return AsComponent(r1)
}

func (m *TComponent) GetEnumerator() IComponentEnumerator {
	r1 := LCL().SysCallN(703, m.Instance())
	return AsComponentEnumerator(r1)
}

func (m *TComponent) GetParentComponent() IComponent {
	r1 := LCL().SysCallN(704, m.Instance())
	return AsComponent(r1)
}

func (m *TComponent) HasParent() bool {
	r1 := LCL().SysCallN(705, m.Instance())
	return GoBool(r1)
}

func (m *TComponent) UpdateAction(Action IBasicAction) bool {
	r1 := LCL().SysCallN(713, m.Instance(), GetObjectUintptr(Action))
	return GoBool(r1)
}

func ComponentClass() TClass {
	ret := LCL().SysCallN(689)
	return TClass(ret)
}

func (m *TComponent) DestroyComponents() {
	LCL().SysCallN(697, m.Instance())
}

func (m *TComponent) Destroying() {
	LCL().SysCallN(698, m.Instance())
}

func (m *TComponent) FreeNotification(AComponent IComponent) {
	LCL().SysCallN(701, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TComponent) RemoveFreeNotification(AComponent IComponent) {
	LCL().SysCallN(710, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TComponent) FreeOnRelease() {
	LCL().SysCallN(702, m.Instance())
}

func (m *TComponent) InsertComponent(AComponent IComponent) {
	LCL().SysCallN(706, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TComponent) RemoveComponent(AComponent IComponent) {
	LCL().SysCallN(709, m.Instance(), GetObjectUintptr(AComponent))
}

func (m *TComponent) SetSubComponent(ASubComponent bool) {
	LCL().SysCallN(711, m.Instance(), PascalBool(ASubComponent))
}
