// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[LayoutXAxisAnchor](lc, objc.GetSelector("alloc"))
	return rv
}

func (lc _LayoutXAxisAnchorClass) New() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutXAxisAnchor() LayoutXAxisAnchor {
	return LayoutXAxisAnchorClass.New()
}

func (l_ LayoutXAxisAnchor) Init() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](l_, objc.GetSelector("init"))
	return rv
}

func (l_ LayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchor_Multiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return rv
}

func (l_ LayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return rv
}

func (l_ LayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchor_Multiplier(anchor ILayoutXAxisAnchor, multiplier float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return rv
}

func (l_ LayoutXAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutXAxisAnchor) LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.GetSelector("anchorWithOffsetToAnchor:"), otherAnchor)
	return rv
}
