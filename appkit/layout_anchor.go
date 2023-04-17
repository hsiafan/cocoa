// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var LayoutAnchorClass = _LayoutAnchorClass{objc.GetClass("NSLayoutAnchor")}

type _LayoutAnchorClass struct {
	objc.Class
}

type ILayoutAnchor interface {
	objc.IObject
	ConstraintEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint
	ConstraintEqualToAnchor_Constant(anchor ILayoutAnchor, c float64) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor_Constant(anchor ILayoutAnchor, c float64) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint
	ConstraintLessThanOrEqualToAnchor_Constant(anchor ILayoutAnchor, c float64) LayoutConstraint
	ConstraintsAffectingLayout() []LayoutConstraint
	HasAmbiguousLayout() bool
	Name() string
	Item() objc.Object
}

type LayoutAnchor struct {
	objc.Object
}

func MakeLayoutAnchor(ptr unsafe.Pointer) LayoutAnchor {
	return LayoutAnchor{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LayoutAnchorClass) Alloc() LayoutAnchor {
	rv := ffi.CallMethod[LayoutAnchor](lc, "alloc")
	return rv
}

func (lc _LayoutAnchorClass) New() LayoutAnchor {
	rv := ffi.CallMethod[LayoutAnchor](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLayoutAnchor() LayoutAnchor {
	return LayoutAnchorClass.New()
}

func (l_ LayoutAnchor) Init() LayoutAnchor {
	rv := ffi.CallMethod[LayoutAnchor](l_, "init")
	return rv
}

func (l_ LayoutAnchor) ConstraintEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintEqualToAnchor:", anchor)
	return rv
}

func (l_ LayoutAnchor) ConstraintEqualToAnchor_Constant(anchor ILayoutAnchor, c float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintEqualToAnchor:constant:", anchor, c)
	return rv
}

func (l_ LayoutAnchor) ConstraintGreaterThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintGreaterThanOrEqualToAnchor:", anchor)
	return rv
}

func (l_ LayoutAnchor) ConstraintGreaterThanOrEqualToAnchor_Constant(anchor ILayoutAnchor, c float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintGreaterThanOrEqualToAnchor:constant:", anchor, c)
	return rv
}

func (l_ LayoutAnchor) ConstraintLessThanOrEqualToAnchor(anchor ILayoutAnchor) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintLessThanOrEqualToAnchor:", anchor)
	return rv
}

func (l_ LayoutAnchor) ConstraintLessThanOrEqualToAnchor_Constant(anchor ILayoutAnchor, c float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "constraintLessThanOrEqualToAnchor:constant:", anchor, c)
	return rv
}

func (l_ LayoutAnchor) ConstraintsAffectingLayout() []LayoutConstraint {
	rv := ffi.CallMethod[[]LayoutConstraint](l_, "constraintsAffectingLayout")
	return rv
}

func (l_ LayoutAnchor) HasAmbiguousLayout() bool {
	rv := ffi.CallMethod[bool](l_, "hasAmbiguousLayout")
	return rv
}

func (l_ LayoutAnchor) Name() string {
	rv := ffi.CallMethod[string](l_, "name")
	return rv
}

func (l_ LayoutAnchor) Item() objc.Object {
	rv := ffi.CallMethod[objc.Object](l_, "item")
	return rv
}
