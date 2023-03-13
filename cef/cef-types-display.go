// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

// CefDisplay
package cef

import (
	"github.com/energye/energy/common/imports"
	"unsafe"
)

func (m *ICefDisplay) ID() (result int64) {
	imports.Proc(internale_CEFDisplay_ID).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) DeviceScaleFactor() float32 {
	var result uintptr
	imports.Proc(internale_CEFDisplay_DeviceScaleFactor).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return *(*float32)(unsafe.Pointer(result))
}

func (m *ICefDisplay) Rotation() int32 {
	r1, _, _ := imports.Proc(internale_CEFDisplay_Rotation).Call(m.Instance())
	return int32(r1)
}

func (m *ICefDisplay) Bounds() (result TCefRect) {
	imports.Proc(internale_CEFDisplay_Bounds).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) WorkArea() (result TCefRect) {
	imports.Proc(internale_CEFDisplay_WorkArea).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefDisplay) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
