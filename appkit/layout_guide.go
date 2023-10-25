// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var LayoutGuideClass = _LayoutGuideClass{objc.GetClass("NSLayoutGuide")}

type _LayoutGuideClass struct {
	objc.Class
}

type ILayoutGuide interface {
	objc.IObject
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
	Frame() foundation.Rect
	OwningView() View
	SetOwningView(value IView)
	BottomAnchor() LayoutYAxisAnchor
	CenterXAnchor() LayoutXAxisAnchor
	CenterYAnchor() LayoutYAxisAnchor
	HeightAnchor() LayoutDimension
	LeadingAnchor() LayoutXAxisAnchor
	LeftAnchor() LayoutXAxisAnchor
	RightAnchor() LayoutXAxisAnchor
	TopAnchor() LayoutYAxisAnchor
	TrailingAnchor() LayoutXAxisAnchor
	WidthAnchor() LayoutDimension
	HasAmbiguousLayout() bool
}

type LayoutGuide struct {
	objc.Object
}

func MakeLayoutGuide(ptr unsafe.Pointer) LayoutGuide {
	return LayoutGuide{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LayoutGuideClass) Alloc() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](lc, objc.SEL("alloc"))
	return rv
}

func (lc _LayoutGuideClass) New() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](lc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutGuide() LayoutGuide {
	return LayoutGuideClass.New()
}

func (l_ LayoutGuide) Init() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](l_, objc.SEL("init"))
	return rv
}

func (l_ LayoutGuide) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](l_, objc.SEL("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}

func (l_ LayoutGuide) Identifier() UserInterfaceItemIdentifier {
	rv := objc.CallMethod[UserInterfaceItemIdentifier](l_, objc.SEL("identifier"))
	return rv
}

func (l_ LayoutGuide) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](l_, objc.SEL("setIdentifier:"), value)
}

func (l_ LayoutGuide) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.SEL("frame"))
	return rv
}

// weak property
func (l_ LayoutGuide) OwningView() View {
	rv := objc.CallMethod[View](l_, objc.SEL("owningView"))
	return rv
}

// weak property
func (l_ LayoutGuide) SetOwningView(value IView) {
	objc.CallMethod[objc.Void](l_, objc.SEL("setOwningView:"), objc.ExtractPtr(value))
}

func (l_ LayoutGuide) BottomAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, objc.SEL("bottomAnchor"))
	return rv
}

func (l_ LayoutGuide) CenterXAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.SEL("centerXAnchor"))
	return rv
}

func (l_ LayoutGuide) CenterYAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, objc.SEL("centerYAnchor"))
	return rv
}

func (l_ LayoutGuide) HeightAnchor() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.SEL("heightAnchor"))
	return rv
}

func (l_ LayoutGuide) LeadingAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.SEL("leadingAnchor"))
	return rv
}

func (l_ LayoutGuide) LeftAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.SEL("leftAnchor"))
	return rv
}

func (l_ LayoutGuide) RightAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.SEL("rightAnchor"))
	return rv
}

func (l_ LayoutGuide) TopAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, objc.SEL("topAnchor"))
	return rv
}

func (l_ LayoutGuide) TrailingAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.SEL("trailingAnchor"))
	return rv
}

func (l_ LayoutGuide) WidthAnchor() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.SEL("widthAnchor"))
	return rv
}

func (l_ LayoutGuide) HasAmbiguousLayout() bool {
	rv := objc.CallMethod[bool](l_, objc.SEL("hasAmbiguousLayout"))
	return rv
}
