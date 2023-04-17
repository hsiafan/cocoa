// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var LayoutXAxisAnchorClass = _LayoutXAxisAnchorClass{objc.GetClass("NSLayoutXAxisAnchor")}

type _LayoutXAxisAnchorClass struct {
	objc.Class
}

type ILayoutXAxisAnchor interface {
	ILayoutAnchor
	ConstraintEqualToSystemSpacingAfterAnchor_Multiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor ILayoutXAxisAnchor) LayoutDimension
}

type LayoutXAxisAnchor struct {
	LayoutAnchor
}

func MakeLayoutXAxisAnchor(ptr unsafe.Pointer) LayoutXAxisAnchor {
	return LayoutXAxisAnchor{
		LayoutAnchor: MakeLayoutAnchor(ptr),
	}
}

func (lc _LayoutXAxisAnchorClass) Alloc() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](lc, "alloc")
	return rv
}

func (lc _LayoutXAxisAnchorClass) New() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLayoutXAxisAnchor() LayoutXAxisAnchor {
	return LayoutXAxisAnchorClass.New()
}

func (l_ LayoutXAxisAnchor) Init() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](l_, "init")
	return rv
}

func (l_ LayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchor_Multiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintEqualToSystemSpacingAfterAnchor:multiplier:", anchor, multiplier)
	return rv
}

func (l_ LayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintGreaterThanOrEqualToSystemSpacingAfterAnchor:multiplier:", anchor, multiplier)
	return rv
}

func (l_ LayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintLessThanOrEqualToSystemSpacingAfterAnchor:multiplier:", anchor, multiplier)
	return rv
}

func (l_ LayoutXAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutXAxisAnchor) LayoutDimension {
	rv := ffi.CallMethod[LayoutDimension](l_, "anchorWithOffsetToAnchor:", otherAnchor)
	return rv
}
