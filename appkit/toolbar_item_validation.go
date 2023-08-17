// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type ToolbarItemValidation interface {
	// required
	ValidateToolbarItem(item ToolbarItem) bool
}

func WrapToolbarItemValidation(v ToolbarItemValidation) objc.Object {
	return objc.WrapAsProtocol("NSToolbarItemValidation", v)
}

type ToolbarItemValidationBase struct {
}

type ToolbarItemValidationWrapper struct {
	objc.Object
}

func (t_ ToolbarItemValidationWrapper) ValidateToolbarItem(item IToolbarItem) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("validateToolbarItem:"), objc.ExtractPtr(item))
	return rv
}
