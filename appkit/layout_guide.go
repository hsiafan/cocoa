// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[LayoutGuide](lc, "alloc")
	return rv
}

func (lc _LayoutGuideClass) New() LayoutGuide {
	rv := ffi.CallMethod[LayoutGuide](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLayoutGuide() LayoutGuide {
	return LayoutGuideClass.New()
}

func (l_ LayoutGuide) Init() LayoutGuide {
	rv := ffi.CallMethod[LayoutGuide](l_, "init")
	return rv
}

func (l_ LayoutGuide) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := ffi.CallMethod[[]LayoutConstraint](l_, "constraintsAffectingLayoutForOrientation:", orientation)
	return rv
}

func (l_ LayoutGuide) Identifier() UserInterfaceItemIdentifier {
	rv := ffi.CallMethod[UserInterfaceItemIdentifier](l_, "identifier")
	return rv
}

func (l_ LayoutGuide) SetIdentifier(value UserInterfaceItemIdentifier) {
	ffi.CallMethod[ffi.Void](l_, "setIdentifier:", value)
}

func (l_ LayoutGuide) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](l_, "frame")
	return rv
}

func (l_ LayoutGuide) OwningView() View {
	rv := ffi.CallMethod[View](l_, "owningView")
	return rv
}

func (l_ LayoutGuide) SetOwningView(value IView) {
	ffi.CallMethod[ffi.Void](l_, "setOwningView:", value)
}

func (l_ LayoutGuide) BottomAnchor() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](l_, "bottomAnchor")
	return rv
}

func (l_ LayoutGuide) CenterXAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](l_, "centerXAnchor")
	return rv
}

func (l_ LayoutGuide) CenterYAnchor() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](l_, "centerYAnchor")
	return rv
}

func (l_ LayoutGuide) HeightAnchor() LayoutDimension {
	rv := ffi.CallMethod[LayoutDimension](l_, "heightAnchor")
	return rv
}

func (l_ LayoutGuide) LeadingAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](l_, "leadingAnchor")
	return rv
}

func (l_ LayoutGuide) LeftAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](l_, "leftAnchor")
	return rv
}

func (l_ LayoutGuide) RightAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](l_, "rightAnchor")
	return rv
}

func (l_ LayoutGuide) TopAnchor() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](l_, "topAnchor")
	return rv
}

func (l_ LayoutGuide) TrailingAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](l_, "trailingAnchor")
	return rv
}

func (l_ LayoutGuide) WidthAnchor() LayoutDimension {
	rv := ffi.CallMethod[LayoutDimension](l_, "widthAnchor")
	return rv
}

func (l_ LayoutGuide) HasAmbiguousLayout() bool {
	rv := ffi.CallMethod[bool](l_, "hasAmbiguousLayout")
	return rv
}
