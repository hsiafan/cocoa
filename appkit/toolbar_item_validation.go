// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type ToolbarItemValidation interface {
	// required
	ValidateToolbarItem(item ToolbarItem) bool
}

type ToolbarItemValidationWrapper struct {
	objc.Object
}

func (t_ ToolbarItemValidationWrapper) ValidateToolbarItem(item IToolbarItem) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("validateToolbarItem:"), item)
	return rv
}
