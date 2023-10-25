// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[LayoutConstraint](lc, objc.SEL("constraintWithItem:attribute:relatedBy:toItem:attribute:multiplier:constant:"), objc.ExtractPtr(view1), attr1, relation, objc.ExtractPtr(view2), attr2, multiplier, c)
	return rv
}

func (lc _LayoutConstraintClass) Alloc() LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](lc, objc.SEL("alloc"))
	return rv
}

func (lc _LayoutConstraintClass) New() LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](lc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutConstraint() LayoutConstraint {
	return LayoutConstraintClass.New()
}

func (l_ LayoutConstraint) Init() LayoutConstraint {
	rv := objc.CallMethod[LayoutConstraint](l_, objc.SEL("init"))
	return rv
}

func (lc _LayoutConstraintClass) ConstraintsWithVisualFormat_Options_Metrics_Views(format string, opts LayoutFormatOptions, metrics map[string]objc.IObject, views map[string]objc.IObject) []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](lc, objc.SEL("constraintsWithVisualFormat:options:metrics:views:"), format, opts, metrics, views)
	return rv
}

func (lc _LayoutConstraintClass) ActivateConstraints(constraints []ILayoutConstraint) {
	objc.CallMethod[objc.Void](lc, objc.SEL("activateConstraints:"), constraints)
}

func (lc _LayoutConstraintClass) DeactivateConstraints(constraints []ILayoutConstraint) {
	objc.CallMethod[objc.Void](lc, objc.SEL("deactivateConstraints:"), constraints)
}

func (l_ LayoutConstraint) IsActive() bool {
	rv := objc.CallMethod[bool](l_, objc.SEL("isActive"))
	return rv
}

func (l_ LayoutConstraint) SetActive(value bool) {
	objc.CallMethod[objc.Void](l_, objc.SEL("setActive:"), value)
}

// weak property
func (l_ LayoutConstraint) FirstItem() objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.SEL("firstItem"))
	return rv
}

func (l_ LayoutConstraint) FirstAttribute() LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](l_, objc.SEL("firstAttribute"))
	return rv
}

func (l_ LayoutConstraint) Relation() LayoutRelation {
	rv := objc.CallMethod[LayoutRelation](l_, objc.SEL("relation"))
	return rv
}

// weak property
func (l_ LayoutConstraint) SecondItem() objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.SEL("secondItem"))
	return rv
}

func (l_ LayoutConstraint) SecondAttribute() LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](l_, objc.SEL("secondAttribute"))
	return rv
}

func (l_ LayoutConstraint) Multiplier() float64 {
	rv := objc.CallMethod[float64](l_, objc.SEL("multiplier"))
	return rv
}

func (l_ LayoutConstraint) Constant() float64 {
	rv := objc.CallMethod[float64](l_, objc.SEL("constant"))
	return rv
}

func (l_ LayoutConstraint) SetConstant(value float64) {
	objc.CallMethod[objc.Void](l_, objc.SEL("setConstant:"), value)
}

func (l_ LayoutConstraint) FirstAnchor() LayoutAnchor {
	rv := objc.CallMethod[LayoutAnchor](l_, objc.SEL("firstAnchor"))
	return rv
}

func (l_ LayoutConstraint) SecondAnchor() LayoutAnchor {
	rv := objc.CallMethod[LayoutAnchor](l_, objc.SEL("secondAnchor"))
	return rv
}

func (l_ LayoutConstraint) Priority() LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](l_, objc.SEL("priority"))
	return rv
}

func (l_ LayoutConstraint) SetPriority(value LayoutPriority) {
	objc.CallMethod[objc.Void](l_, objc.SEL("setPriority:"), value)
}

func (l_ LayoutConstraint) Identifier() string {
	rv := objc.CallMethod[string](l_, objc.SEL("identifier"))
	return rv
}

func (l_ LayoutConstraint) SetIdentifier(value string) {
	objc.CallMethod[objc.Void](l_, objc.SEL("setIdentifier:"), value)
}

func (l_ LayoutConstraint) ShouldBeArchived() bool {
	rv := objc.CallMethod[bool](l_, objc.SEL("shouldBeArchived"))
	return rv
}

func (l_ LayoutConstraint) SetShouldBeArchived(value bool) {
	objc.CallMethod[objc.Void](l_, objc.SEL("setShouldBeArchived:"), value)
}
