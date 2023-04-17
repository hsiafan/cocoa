// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var LayoutConstraintClass = _LayoutConstraintClass{objc.GetClass("NSLayoutConstraint")}

type _LayoutConstraintClass struct {
	objc.Class
}

type ILayoutConstraint interface {
	objc.IObject
	IsActive() bool
	SetActive(value bool)
	FirstItem() objc.Object
	FirstAttribute() LayoutAttribute
	Relation() LayoutRelation
	SecondItem() objc.Object
	SecondAttribute() LayoutAttribute
	Multiplier() float64
	Constant() float64
	SetConstant(value float64)
	FirstAnchor() LayoutAnchor
	SecondAnchor() LayoutAnchor
	Priority() LayoutPriority
	SetPriority(value LayoutPriority)
	Identifier() string
	SetIdentifier(value string)
	ShouldBeArchived() bool
	SetShouldBeArchived(value bool)
}

type LayoutConstraint struct {
	objc.Object
}

func MakeLayoutConstraint(ptr unsafe.Pointer) LayoutConstraint {
	return LayoutConstraint{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LayoutConstraintClass) ConstraintWithItem_Attribute_RelatedBy_ToItem_Attribute_Multiplier_Constant(view1 objc.IObject, attr1 LayoutAttribute, relation LayoutRelation, view2 objc.IObject, attr2 LayoutAttribute, multiplier float64, c float64) LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](lc, "constraintWithItem:attribute:relatedBy:toItem:attribute:multiplier:constant:", view1, attr1, relation, view2, attr2, multiplier, c)
	return rv
}

func (lc _LayoutConstraintClass) Alloc() LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](lc, "alloc")
	return rv
}

func (lc _LayoutConstraintClass) New() LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLayoutConstraint() LayoutConstraint {
	return LayoutConstraintClass.New()
}

func (l_ LayoutConstraint) Init() LayoutConstraint {
	rv := ffi.CallMethod[LayoutConstraint](l_, "init")
	return rv
}

func (lc _LayoutConstraintClass) ConstraintsWithVisualFormat_Options_Metrics_Views(format string, opts LayoutFormatOptions, metrics map[string]objc.IObject, views map[string]objc.IObject) []LayoutConstraint {
	rv := ffi.CallMethod[[]LayoutConstraint](lc, "constraintsWithVisualFormat:options:metrics:views:", format, opts, metrics, views)
	return rv
}

func (lc _LayoutConstraintClass) ActivateConstraints(constraints []ILayoutConstraint) {
	ffi.CallMethod[ffi.Void](lc, "activateConstraints:", constraints)
}

func (lc _LayoutConstraintClass) DeactivateConstraints(constraints []ILayoutConstraint) {
	ffi.CallMethod[ffi.Void](lc, "deactivateConstraints:", constraints)
}

func (l_ LayoutConstraint) IsActive() bool {
	rv := ffi.CallMethod[bool](l_, "isActive")
	return rv
}

func (l_ LayoutConstraint) SetActive(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setActive:", value)
}

func (l_ LayoutConstraint) FirstItem() objc.Object {
	rv := ffi.CallMethod[objc.Object](l_, "firstItem")
	return rv
}

func (l_ LayoutConstraint) FirstAttribute() LayoutAttribute {
	rv := ffi.CallMethod[LayoutAttribute](l_, "firstAttribute")
	return rv
}

func (l_ LayoutConstraint) Relation() LayoutRelation {
	rv := ffi.CallMethod[LayoutRelation](l_, "relation")
	return rv
}

func (l_ LayoutConstraint) SecondItem() objc.Object {
	rv := ffi.CallMethod[objc.Object](l_, "secondItem")
	return rv
}

func (l_ LayoutConstraint) SecondAttribute() LayoutAttribute {
	rv := ffi.CallMethod[LayoutAttribute](l_, "secondAttribute")
	return rv
}

func (l_ LayoutConstraint) Multiplier() float64 {
	rv := ffi.CallMethod[float64](l_, "multiplier")
	return rv
}

func (l_ LayoutConstraint) Constant() float64 {
	rv := ffi.CallMethod[float64](l_, "constant")
	return rv
}

func (l_ LayoutConstraint) SetConstant(value float64) {
	ffi.CallMethod[ffi.Void](l_, "setConstant:", value)
}

func (l_ LayoutConstraint) FirstAnchor() LayoutAnchor {
	rv := ffi.CallMethod[LayoutAnchor](l_, "firstAnchor")
	return rv
}

func (l_ LayoutConstraint) SecondAnchor() LayoutAnchor {
	rv := ffi.CallMethod[LayoutAnchor](l_, "secondAnchor")
	return rv
}

func (l_ LayoutConstraint) Priority() LayoutPriority {
	rv := ffi.CallMethod[LayoutPriority](l_, "priority")
	return rv
}

func (l_ LayoutConstraint) SetPriority(value LayoutPriority) {
	ffi.CallMethod[ffi.Void](l_, "setPriority:", value)
}

func (l_ LayoutConstraint) Identifier() string {
	rv := ffi.CallMethod[string](l_, "identifier")
	return rv
}

func (l_ LayoutConstraint) SetIdentifier(value string) {
	ffi.CallMethod[ffi.Void](l_, "setIdentifier:", value)
}

func (l_ LayoutConstraint) ShouldBeArchived() bool {
	rv := ffi.CallMethod[bool](l_, "shouldBeArchived")
	return rv
}

func (l_ LayoutConstraint) SetShouldBeArchived(value bool) {
	ffi.CallMethod[ffi.Void](l_, "setShouldBeArchived:", value)
}