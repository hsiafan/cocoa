// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var LayoutYAxisAnchorClass = _LayoutYAxisAnchorClass{objc.GetClass("NSLayoutYAxisAnchor")}

type _LayoutYAxisAnchorClass struct {
	objc.Class
}

type ILayoutYAxisAnchor interface {
	ILayoutAnchor
	ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint
	AnchorWithOffsetToAnchor(otherAnchor ILayoutYAxisAnchor) LayoutDimension
}

type LayoutYAxisAnchor struct {
	LayoutAnchor
}

func MakeLayoutYAxisAnchor(ptr unsafe.Pointer) LayoutYAxisAnchor {
	return LayoutYAxisAnchor{
		LayoutAnchor: MakeLayoutAnchor(ptr),
	}
}

func (lc _LayoutYAxisAnchorClass) Alloc() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](lc, "alloc")
	return rv
}

func (lc _LayoutYAxisAnchorClass) New() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLayoutYAxisAnchor() LayoutYAxisAnchor {
	return LayoutYAxisAnchorClass.New()
}

func (l_ LayoutYAxisAnchor) Init() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](l_, "init")
	return rv
}

func (l_ LayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintEqualToSystemSpacingBelowAnchor:multiplier:", anchor, multiplier)
	return rv
}

func (l_ LayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:multiplier:", anchor, multiplier)
	return rv
}

func (l_ LayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintLessThanOrEqualToSystemSpacingBelowAnchor:multiplier:", anchor, multiplier)
	return rv
}

func (l_ LayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutYAxisAnchor) LayoutDimension {
	rv := ffi.CallMethod[LayoutDimension](l_, "anchorWithOffsetToAnchor:", otherAnchor)
	return rv
}
