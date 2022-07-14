// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var WindowFeaturesClass = _WindowFeaturesClass{objc.GetClass("WKWindowFeatures")}

type _WindowFeaturesClass struct {
	objc.Class
}

type IWindowFeatures interface {
	objc.IObject
	AllowsResizing() foundation.Number
	Height() foundation.Number
	Width() foundation.Number
	X() foundation.Number
	Y() foundation.Number
	MenuBarVisibility() foundation.Number
	StatusBarVisibility() foundation.Number
	ToolbarsVisibility() foundation.Number
}

type WindowFeatures struct {
	objc.Object
}

func MakeWindowFeatures(ptr unsafe.Pointer) WindowFeatures {
	return WindowFeatures{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WindowFeaturesClass) Alloc() WindowFeatures {
	rv := ffi.CallMethod[WindowFeatures](wc, "alloc")
	return rv
}

func (w_ WindowFeatures) Init() WindowFeatures {
	rv := ffi.CallMethod[WindowFeatures](w_, "init")
	rv.Autorelease()
	return rv
}

func (wc _WindowFeaturesClass) New() WindowFeatures {
	rv := ffi.CallMethod[WindowFeatures](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWindowFeatures() WindowFeatures {
	return WindowFeaturesClass.New()
}

func (w_ WindowFeatures) AllowsResizing() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](w_, "allowsResizing")
	return rv
}

func (w_ WindowFeatures) Height() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](w_, "height")
	return rv
}

func (w_ WindowFeatures) Width() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](w_, "width")
	return rv
}

func (w_ WindowFeatures) X() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](w_, "x")
	return rv
}

func (w_ WindowFeatures) Y() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](w_, "y")
	return rv
}

func (w_ WindowFeatures) MenuBarVisibility() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](w_, "menuBarVisibility")
	return rv
}

func (w_ WindowFeatures) StatusBarVisibility() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](w_, "statusBarVisibility")
	return rv
}

func (w_ WindowFeatures) ToolbarsVisibility() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](w_, "toolbarsVisibility")
	return rv
}
