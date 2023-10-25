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

func WrapViewToolTipOwner(v ViewToolTipOwner) objc.Object {
	return objc.WrapAsProtocol("NSViewToolTipOwner", v)
}

type ViewToolTipOwnerBase struct {
}
