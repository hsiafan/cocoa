// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[LayoutYAxisAnchor](lc, objc.SEL("alloc"))
	return rv
}

func (lc _LayoutYAxisAnchorClass) New() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](lc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutYAxisAnchor() LayoutYAxisAnchor {
	return LayoutYAxisAnchorClass.New()
}

func (l_ LayoutYAxisAnchor) Init() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](l_, objc.SEL("init"))
	return rv
}

func (l_ LayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.SEL("constraintEqualToSystemSpacingBelowAnchor:multiplier:"), objc.ExtractPtr(anchor), multiplier)
	return rv
}

func (l_ LayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.SEL("constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), objc.ExtractPtr(anchor), multiplier)
	return rv
}

func (l_ LayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(anchor ILayoutYAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.SEL("constraintLessThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), objc.ExtractPtr(anchor), multiplier)
	return rv
}

func (l_ LayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutYAxisAnchor) LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.SEL("anchorWithOffsetToAnchor:"), objc.ExtractPtr(otherAnchor))
	return rv
}
