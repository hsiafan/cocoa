// auto-generated code, do not modify
package webkit

import (
	"unsafe"

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
	rv := objc.CallMethod[WindowFeatures](wc, objc.SEL("alloc"))
	return rv
}

func (wc _WindowFeaturesClass) New() WindowFeatures {
	rv := objc.CallMethod[WindowFeatures](wc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewWindowFeatures() WindowFeatures {
	return WindowFeaturesClass.New()
}

func (w_ WindowFeatures) Init() WindowFeatures {
	rv := objc.CallMethod[WindowFeatures](w_, objc.SEL("init"))
	return rv
}

func (w_ WindowFeatures) AllowsResizing() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.SEL("allowsResizing"))
	return rv
}

func (w_ WindowFeatures) Height() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.SEL("height"))
	return rv
}

func (w_ WindowFeatures) Width() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.SEL("width"))
	return rv
}

func (w_ WindowFeatures) X() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.SEL("x"))
	return rv
}

func (w_ WindowFeatures) Y() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.SEL("y"))
	return rv
}

func (w_ WindowFeatures) MenuBarVisibility() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.SEL("menuBarVisibility"))
	return rv
}

func (w_ WindowFeatures) StatusBarVisibility() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.SEL("statusBarVisibility"))
	return rv
}

func (w_ WindowFeatures) ToolbarsVisibility() foundation.Number {
	rv := objc.CallMethod[foundation.Number](w_, objc.SEL("toolbarsVisibility"))
	return rv
}
