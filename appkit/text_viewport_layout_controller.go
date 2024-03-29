// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
)

var TextViewportLayoutControllerClass = _TextViewportLayoutControllerClass{objc.GetClass("NSTextViewportLayoutController")}

type _TextViewportLayoutControllerClass struct {
	objc.Class
}

type ITextViewportLayoutController interface {
	objc.IObject
	AdjustViewportByVerticalOffset(verticalOffset float64)
	LayoutViewport()
	TextLayoutManager() TextLayoutManager
	ViewportBounds() coregraphics.Rect
	ViewportRange() TextRange
}

type TextViewportLayoutController struct {
	objc.Object
}

func MakeTextViewportLayoutController(ptr unsafe.Pointer) TextViewportLayoutController {
	return TextViewportLayoutController{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextViewportLayoutController) InitWithTextLayoutManager(textLayoutManager ITextLayoutManager) TextViewportLayoutController {
	rv := objc.CallMethod[TextViewportLayoutController](t_, objc.SEL("initWithTextLayoutManager:"), objc.ExtractPtr(textLayoutManager))
	return rv
}

func (tc _TextViewportLayoutControllerClass) Alloc() TextViewportLayoutController {
	rv := objc.CallMethod[TextViewportLayoutController](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextViewportLayoutControllerClass) New() TextViewportLayoutController {
	rv := objc.CallMethod[TextViewportLayoutController](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextViewportLayoutController() TextViewportLayoutController {
	return TextViewportLayoutControllerClass.New()
}

func (t_ TextViewportLayoutController) Init() TextViewportLayoutController {
	rv := objc.CallMethod[TextViewportLayoutController](t_, objc.SEL("init"))
	return rv
}

func (t_ TextViewportLayoutController) AdjustViewportByVerticalOffset(verticalOffset float64) {
	objc.CallMethod[objc.Void](t_, objc.SEL("adjustViewportByVerticalOffset:"), verticalOffset)
}

func (t_ TextViewportLayoutController) LayoutViewport() {
	objc.CallMethod[objc.Void](t_, objc.SEL("layoutViewport"))
}

// weak property
func (t_ TextViewportLayoutController) TextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.SEL("textLayoutManager"))
	return rv
}

func (t_ TextViewportLayoutController) ViewportBounds() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.SEL("viewportBounds"))
	return rv
}

func (t_ TextViewportLayoutController) ViewportRange() TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.SEL("viewportRange"))
	return rv
}
