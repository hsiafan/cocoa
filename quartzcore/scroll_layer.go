// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
)

var ScrollLayerClass = _ScrollLayerClass{objc.GetClass("CAScrollLayer")}

type _ScrollLayerClass struct {
	objc.Class
}

type IScrollLayer interface {
	ILayer
	ScrollToPoint(p coregraphics.Point)
	ScrollToRect(r coregraphics.Rect)
	ScrollMode() ScrollLayerScrollMode
	SetScrollMode(value ScrollLayerScrollMode)
}

type ScrollLayer struct {
	Layer
}

func MakeScrollLayer(ptr unsafe.Pointer) ScrollLayer {
	return ScrollLayer{
		Layer: MakeLayer(ptr),
	}
}

func (sc _ScrollLayerClass) Layer() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](sc, objc.GetSelector("layer"))
	return rv
}

func (s_ ScrollLayer) Init() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](s_, objc.GetSelector("init"))
	return rv
}

func (s_ ScrollLayer) InitWithLayer(layer objc.IObject) ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](s_, objc.GetSelector("initWithLayer:"), objc.ExtractPtr(layer))
	return rv
}

func (s_ ScrollLayer) PresentationLayer() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](s_, objc.GetSelector("presentationLayer"))
	return rv
}

func (s_ ScrollLayer) ModelLayer() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](s_, objc.GetSelector("modelLayer"))
	return rv
}

func (sc _ScrollLayerClass) Alloc() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](sc, objc.GetSelector("alloc"))
	return rv
}

func (sc _ScrollLayerClass) New() ScrollLayer {
	rv := objc.CallMethod[ScrollLayer](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewScrollLayer() ScrollLayer {
	return ScrollLayerClass.New()
}

func (s_ ScrollLayer) ScrollToPoint(p coregraphics.Point) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("scrollToPoint:"), p)
}

func (s_ ScrollLayer) ScrollToRect(r coregraphics.Rect) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("scrollToRect:"), r)
}

func (s_ ScrollLayer) ScrollMode() ScrollLayerScrollMode {
	rv := objc.CallMethod[ScrollLayerScrollMode](s_, objc.GetSelector("scrollMode"))
	return rv
}

func (s_ ScrollLayer) SetScrollMode(value ScrollLayerScrollMode) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setScrollMode:"), value)
}
