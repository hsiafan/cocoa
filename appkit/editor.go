// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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

func WrapEditor(v Editor) objc.Object {
	return objc.WrapAsProtocol("NSEditor", v)
}

type EditorBase struct {
}
