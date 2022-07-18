// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type Editor interface {
	// required
	CommitEditing() bool
	// required
	CommitEditingWithDelegate_DidCommitSelector_ContextInfo(delegate objc.Object, didCommitSelector objc.Selector, contextInfo unsafe.Pointer)
	// required
	CommitEditingAndReturnError(error *foundation.Error) bool
	// required
	DiscardEditing()
}

type EditorWrapper struct {
	objc.Object
}

func (e_ EditorWrapper) CommitEditing() bool {
	rv := ffi.CallMethod[bool](e_, "commitEditing")
	return rv
}

func (e_ EditorWrapper) CommitEditingWithDelegate_DidCommitSelector_ContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](e_, "commitEditingWithDelegate:didCommitSelector:contextInfo:", delegate, didCommitSelector, contextInfo)
}

func (e_ EditorWrapper) CommitEditingAndReturnError(error *foundation.Error) bool {
	rv := ffi.CallMethod[bool](e_, "commitEditingAndReturnError:", unsafe.Pointer(error))
	return rv
}

func (e_ EditorWrapper) DiscardEditing() {
	ffi.CallMethod[ffi.Void](e_, "discardEditing")
}
