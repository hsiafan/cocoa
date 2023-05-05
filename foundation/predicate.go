// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var PredicateClass = _PredicateClass{objc.GetClass("NSPredicate")}

type _PredicateClass struct {
	objc.Class
}

type IPredicate interface {
	objc.IObject
	EvaluateWithObject(object objc.IObject) bool
	EvaluateWithObject_SubstitutionVariables(object objc.IObject, bindings map[string]objc.IObject) bool
	AllowEvaluation()
	PredicateFormat() string
}

type Predicate struct {
	objc.Object
}

func MakePredicate(ptr unsafe.Pointer) Predicate {
	return Predicate{
		Object: objc.MakeObject(ptr),
	}
}

func (p_ Predicate) PredicateWithSubstitutionVariables(variables map[string]objc.IObject) Predicate {
	rv := objc.CallMethod[Predicate](p_, objc.GetSelector("predicateWithSubstitutionVariables:"), variables)
	return rv
}

func (pc _PredicateClass) Alloc() Predicate {
	rv := objc.CallMethod[Predicate](pc, objc.GetSelector("alloc"))
	return rv
}

func (pc _PredicateClass) New() Predicate {
	rv := objc.CallMethod[Predicate](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPredicate() Predicate {
	return PredicateClass.New()
}

func (p_ Predicate) Init() Predicate {
	rv := objc.CallMethod[Predicate](p_, objc.GetSelector("init"))
	return rv
}

func (pc _PredicateClass) PredicateWithFormat_ArgumentArray(predicateFormat string, arguments []objc.IObject) Predicate {
	rv := objc.CallMethod[Predicate](pc, objc.GetSelector("predicateWithFormat:argumentArray:"), predicateFormat, arguments)
	return rv
}

func (pc _PredicateClass) PredicateWithValue(value bool) Predicate {
	rv := objc.CallMethod[Predicate](pc, objc.GetSelector("predicateWithValue:"), value)
	return rv
}

func (pc _PredicateClass) PredicateFromMetadataQueryString(queryString string) Predicate {
	rv := objc.CallMethod[Predicate](pc, objc.GetSelector("predicateFromMetadataQueryString:"), queryString)
	return rv
}

func (p_ Predicate) EvaluateWithObject(object objc.IObject) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("evaluateWithObject:"), objc.ExtractPtr(object))
	return rv
}

func (p_ Predicate) EvaluateWithObject_SubstitutionVariables(object objc.IObject, bindings map[string]objc.IObject) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("evaluateWithObject:substitutionVariables:"), objc.ExtractPtr(object), bindings)
	return rv
}

func (p_ Predicate) AllowEvaluation() {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("allowEvaluation"))
}

func (p_ Predicate) PredicateFormat() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("predicateFormat"))
	return rv
}
