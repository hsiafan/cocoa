// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type ViewToolTipOwner interface {
	// required
	View_StringForToolTip_Point_UserData(view View, tag ToolTipTag, point foundation.Point, data unsafe.Pointer) string
}

type ViewToolTipOwnerWrapper struct {
	objc.Object
}

func (v_ ViewToolTipOwnerWrapper) View_StringForToolTip_Point_UserData(view IView, tag ToolTipTag, point foundation.Point, data unsafe.Pointer) string {
	rv := ffi.CallMethod[string](v_, "view:stringForToolTip:point:userData:", view, tag, point, data)
	return rv
}
