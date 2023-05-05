// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var LayoutDimensionClass = _LayoutDimensionClass{objc.GetClass("NSLayoutDimension")}

type _LayoutDimensionClass struct {
	objc.Class
}

type ILayoutDimension interface {
	ILayoutAnchor
	ConstraintEqualToAnchor_Multiplier(anchor ILayoutDimension, m float64) LayoutConstraint
	ConstraintEqualToAnchor_Multiplier_Constant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint
	ConstraintEqualToConstant(c float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor_Multiplier(anchor ILayoutDimension, m float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToConstant(c float64) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor_Multiplier(anchor ILayoutDimension, m float64) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint
	ConstraintLessThanOrEqualToConstant(c float64) LayoutConstraint
}

type LayoutDimension struct {
	LayoutAnchor
}

func MakeLayoutDimension(ptr unsafe.Pointer) LayoutDimension {
	return LayoutDimension{
		LayoutAnchor: MakeLayoutAnchor(ptr),
	}
}

func (lc _LayoutDimensionClass) Alloc() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](lc, objc.GetSelector("alloc"))
	return rv
}

func (lc _LayoutDimensionClass) New() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutDimension() LayoutDimension {
	return LayoutDimensionClass.New()
}

func (l_ LayoutDimension) Init() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](l_, objc.GetSelector("init"))
	return rv
}

func (l_ LayoutDimension) ConstraintEqualToAnchor_Multiplier(anchor ILayoutDimension, m float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToAnchor:multiplier:"), objc.ExtractPtr(anchor), m)
	return rv
}

func (l_ LayoutDimension) ConstraintEqualToAnchor_Multiplier_Constant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToAnchor:multiplier:constant:"), objc.ExtractPtr(anchor), m, c)
	return rv
}

func (l_ LayoutDimension) ConstraintEqualToConstant(c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintEqualToConstant:"), c)
	return rv
}

func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToAnchor_Multiplier(anchor ILayoutDimension, m float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToAnchor:multiplier:"), objc.ExtractPtr(anchor), m)
	return rv
}

func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToAnchor:multiplier:constant:"), objc.ExtractPtr(anchor), m, c)
	return rv
}

func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToConstant(c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintGreaterThanOrEqualToConstant:"), c)
	return rv
}

func (l_ LayoutDimension) ConstraintLessThanOrEqualToAnchor_Multiplier(anchor ILayoutDimension, m float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToAnchor:multiplier:"), objc.ExtractPtr(anchor), m)
	return rv
}

func (l_ LayoutDimension) ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToAnchor:multiplier:constant:"), objc.ExtractPtr(anchor), m, c)
	return rv
}

func (l_ LayoutDimension) ConstraintLessThanOrEqualToConstant(c float64) LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.GetSelector("constraintLessThanOrEqualToConstant:"), c)
	return rv
}
