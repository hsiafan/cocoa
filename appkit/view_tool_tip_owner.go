// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[string](v_, objc.GetSelector("view:stringForToolTip:point:userData:"), objc.ExtractPtr(view), tag, point, data)
	return rv
}
