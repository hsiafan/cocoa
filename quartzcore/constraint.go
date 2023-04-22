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
	rv := objc.CallMethod[Constraint](cc, objc.GetSelector("constraintWithAttribute:relativeTo:attribute:scale:offset:"), attr, srcId, srcAttr, m, c)
	return rv
}

func (cc _ConstraintClass) ConstraintWithAttribute_RelativeTo_Attribute_Offset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, c float64) Constraint {
	rv := objc.CallMethod[Constraint](cc, objc.GetSelector("constraintWithAttribute:relativeTo:attribute:offset:"), attr, srcId, srcAttr, c)
	return rv
}

func (cc _ConstraintClass) ConstraintWithAttribute_RelativeTo_Attribute(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute) Constraint {
	rv := objc.CallMethod[Constraint](cc, objc.GetSelector("constraintWithAttribute:relativeTo:attribute:"), attr, srcId, srcAttr)
	return rv
}

func (c_ Constraint) InitWithAttribute_RelativeTo_Attribute_Scale_Offset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, m float64, c float64) Constraint {
	rv := objc.CallMethod[Constraint](c_, objc.GetSelector("initWithAttribute:relativeTo:attribute:scale:offset:"), attr, srcId, srcAttr, m, c)
	return rv
}

func (cc _ConstraintClass) Alloc() Constraint {
	rv := objc.CallMethod[Constraint](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _ConstraintClass) New() Constraint {
	rv := objc.CallMethod[Constraint](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewConstraint() Constraint {
	return ConstraintClass.New()
}

func (c_ Constraint) Init() Constraint {
	rv := objc.CallMethod[Constraint](c_, objc.GetSelector("init"))
	return rv
}

func (c_ Constraint) Attribute() ConstraintAttribute {
	rv := objc.CallMethod[ConstraintAttribute](c_, objc.GetSelector("attribute"))
	return rv
}

func (c_ Constraint) Offset() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("offset"))
	return rv
}

func (c_ Constraint) Scale() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("scale"))
	return rv
}

func (c_ Constraint) SourceAttribute() ConstraintAttribute {
	rv := objc.CallMethod[ConstraintAttribute](c_, objc.GetSelector("sourceAttribute"))
	return rv
}

func (c_ Constraint) SourceName() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("sourceName"))
	return rv
}
