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

// IObject Root Interface
type IObject interface {
	Instance() uintptr
	UnsafeAddr() unsafe.Pointer
	IsValid() bool
	Is() TIs
	SetInstance(instance unsafe.Pointer)
	Equals(Obj IObject) bool           // function
	GetHashCode() uint32               // function
	ToString() string                  // function
	InheritsFrom(AClass TClass) bool   // function
	ClassType() TClass                 // function
	ClassName() string                 // function
	ClassParent() TClass               // function
	InstanceSize() (resultInt64 int64) // function
	Free()                             // procedure
}

// TObject Root Object
type TObject struct {
	instance unsafe.Pointer
}

func NewObject() IObject {
	r1 := LCL().SysCallN(3720)
	return AsObject(r1)
}

func (m *TObject) Equals(Obj IObject) bool {
	r1 := LCL().SysCallN(3721, m.Instance(), GetObjectUintptr(Obj))
	return GoBool(r1)
}

func (m *TObject) GetHashCode() uint32 {
	r1 := LCL().SysCallN(3723, m.Instance())
	return uint32(r1)
}

func ObjectClass() TClass {
	ret := LCL().SysCallN(3716)
	return TClass(ret)
}

func (m *TObject) ToString() string {
	r1 := LCL().SysCallN(3726, m.Instance())
	return GoStr(r1)
}

func (m *TObject) InheritsFrom(AClass TClass) bool {
	r1 := LCL().SysCallN(3724, m.Instance(), uintptr(AClass))
	return GoBool(r1)
}

func (m *TObject) ClassType() TClass {
	r1 := LCL().SysCallN(3719, m.Instance())
	return TClass(r1)
}

func (m *TObject) ClassName() string {
	r1 := LCL().SysCallN(3717, m.Instance())
	return GoStr(r1)
}

func (m *TObject) ClassParent() TClass {
	r1 := LCL().SysCallN(3718, m.Instance())
	return TClass(r1)
}

func (m *TObject) InstanceSize() (resultInt64 int64) {
	LCL().SysCallN(3725, m.Instance(), uintptr(unsafe.Pointer(&resultInt64)))
	return
}

func (m *TObject) Free() {
	m.free(3722)
}
