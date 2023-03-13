//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Response
package cef

import (
	"github.com/energye/energy/common/imports"
	. "github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ResponseRef -> ICefResponse
var ResponseRef response

// request
type response uintptr

func (*response) UnWrap(data *ICefResponse) *ICefResponse {
	var result uintptr
	imports.Proc(internale_CefResponseRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	data.instance = unsafe.Pointer(result)
	return data
}

// Instance 实例
func (m *ICefResponse) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefResponse) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

// IsReadOnly 是否只读
func (m *ICefResponse) IsReadOnly() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_CefResponse_IsReadOnly).Call(m.Instance())
	return api.GoBool(r1)
}

// SetURL 设置URL
func (m *ICefResponse) SetURL(url string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponse_SetURL).Call(m.Instance(), api.PascalStr(url))
}

// SetError 设置错误码
func (m *ICefResponse) SetError(error TCefErrorCode) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponse_SetError).Call(m.Instance(), error.ToPtr())
}

// SetStatus 设置状态码
func (m *ICefResponse) SetStatus(status int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponse_SetStatus).Call(m.Instance(), uintptr(status))
}

// SetStatusText 设置状态文本
func (m *ICefResponse) SetStatusText(statusText string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponse_SetStatusText).Call(m.Instance(), api.PascalStr(statusText))
}

// SetMimeType mime类型
func (m *ICefResponse) SetMimeType(mimetype string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponse_SetMimeType).Call(m.Instance(), api.PascalStr(mimetype))
}

// SetCharset 设置编码
func (m *ICefResponse) SetCharset(charset string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponse_SetCharset).Call(m.Instance(), api.PascalStr(charset))
}

// GetHeaderByName
func (m *ICefResponse) GetHeaderByName(name string) string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(internale_CefResponse_GetHeaderByName).Call(m.Instance(), api.PascalStr(name))
	return api.GoStr(r1)
}

// SetHeaderByName
func (m *ICefResponse) SetHeaderByName(name, value string, overwrite bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponse_SetHeaderByName).Call(m.Instance(), api.PascalStr(name), api.PascalStr(value), api.PascalBool(overwrite))
}

// GetHeaderMap
func (m *ICefResponse) GetHeaderMap() *ICefStringMultiMap {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(internale_CefResponse_GetHeaderMap).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefStringMultiMap{instance: unsafe.Pointer(result)}
}

// GetHeaderMap
func (m *ICefResponse) SetHeaderMap(headerMap *ICefStringMultiMap) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefResponse_SetHeaderMap).Call(m.Instance(), headerMap.Instance())
}

func (m *ICefResponse) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}
