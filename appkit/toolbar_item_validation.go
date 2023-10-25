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
