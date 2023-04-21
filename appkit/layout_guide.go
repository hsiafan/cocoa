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
	rv := objc.CallMethod[LayoutGuide](lc, "alloc")
	return rv
}

func (lc _LayoutGuideClass) New() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLayoutGuide() LayoutGuide {
	return LayoutGuideClass.New()
}

func (l_ LayoutGuide) Init() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](l_, "init")
	return rv
}

func (l_ LayoutGuide) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](l_, "constraintsAffectingLayoutForOrientation:", orientation)
	return rv
}

func (l_ LayoutGuide) Identifier() UserInterfaceItemIdentifier {
	rv := objc.CallMethod[UserInterfaceItemIdentifier](l_, "identifier")
	return rv
}

func (l_ LayoutGuide) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](l_, "setIdentifier:", value)
}

func (l_ LayoutGuide) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, "frame")
	return rv
}

func (l_ LayoutGuide) OwningView() View {
	rv := objc.CallMethod[View](l_, "owningView")
	return rv
}

func (l_ LayoutGuide) SetOwningView(value IView) {
	objc.CallMethod[objc.Void](l_, "setOwningView:", value)
}

func (l_ LayoutGuide) BottomAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, "bottomAnchor")
	return rv
}

func (l_ LayoutGuide) CenterXAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, "centerXAnchor")
	return rv
}

func (l_ LayoutGuide) CenterYAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, "centerYAnchor")
	return rv
}

func (l_ LayoutGuide) HeightAnchor() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, "heightAnchor")
	return rv
}

func (l_ LayoutGuide) LeadingAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, "leadingAnchor")
	return rv
}

func (l_ LayoutGuide) LeftAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, "leftAnchor")
	return rv
}

func (l_ LayoutGuide) RightAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, "rightAnchor")
	return rv
}

func (l_ LayoutGuide) TopAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, "topAnchor")
	return rv
}

func (l_ LayoutGuide) TrailingAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, "trailingAnchor")
	return rv
}

func (l_ LayoutGuide) WidthAnchor() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, "widthAnchor")
	return rv
}

func (l_ LayoutGuide) HasAmbiguousLayout() bool {
	rv := objc.CallMethod[bool](l_, "hasAmbiguousLayout")
	return rv
}
