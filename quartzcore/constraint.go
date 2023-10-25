// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ConstraintClass = _ConstraintClass{objc.GetClass("CAConstraint")}

type _ConstraintClass struct {
	objc.Class
}

type IConstraint interface {
	objc.IObject
	Attribute() ConstraintAttribute
	Offset() float64
	Scale() float64
	SourceAttribute() ConstraintAttribute
	SourceName() string
}

type Constraint struct {
	objc.Object
}

func MakeConstraint(ptr unsafe.Pointer) Constraint {
	return Constraint{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ConstraintClass) ConstraintWithAttribute_RelativeTo_Attribute_Scale_Offset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, m float64, c float64) Constraint {
	rv := objc.CallMethod[Constraint](cc, objc.SEL("constraintWithAttribute:relativeTo:attribute:scale:offset:"), attr, srcId, srcAttr, m, c)
	return rv
}

func (cc _ConstraintClass) ConstraintWithAttribute_RelativeTo_Attribute_Offset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, c float64) Constraint {
	rv := objc.CallMethod[Constraint](cc, objc.SEL("constraintWithAttribute:relativeTo:attribute:offset:"), attr, srcId, srcAttr, c)
	return rv
}

func (cc _ConstraintClass) ConstraintWithAttribute_RelativeTo_Attribute(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute) Constraint {
	rv := objc.CallMethod[Constraint](cc, objc.SEL("constraintWithAttribute:relativeTo:attribute:"), attr, srcId, srcAttr)
	return rv
}

func (c_ Constraint) InitWithAttribute_RelativeTo_Attribute_Scale_Offset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, m float64, c float64) Constraint {
	rv := objc.CallMethod[Constraint](c_, objc.SEL("initWithAttribute:relativeTo:attribute:scale:offset:"), attr, srcId, srcAttr, m, c)
	return rv
}

func (cc _ConstraintClass) Alloc() Constraint {
	rv := objc.CallMethod[Constraint](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ConstraintClass) New() Constraint {
	rv := objc.CallMethod[Constraint](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewConstraint() Constraint {
	return ConstraintClass.New()
}

func (c_ Constraint) Init() Constraint {
	rv := objc.CallMethod[Constraint](c_, objc.SEL("init"))
	return rv
}

func (c_ Constraint) Attribute() ConstraintAttribute {
	rv := objc.CallMethod[ConstraintAttribute](c_, objc.SEL("attribute"))
	return rv
}

func (c_ Constraint) Offset() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("offset"))
	return rv
}

func (c_ Constraint) Scale() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("scale"))
	return rv
}

func (c_ Constraint) SourceAttribute() ConstraintAttribute {
	rv := objc.CallMethod[ConstraintAttribute](c_, objc.SEL("sourceAttribute"))
	return rv
}

func (c_ Constraint) SourceName() string {
	rv := objc.CallMethod[string](c_, objc.SEL("sourceName"))
	return rv
}
