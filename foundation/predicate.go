// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[Predicate](p_, "predicateWithSubstitutionVariables:", variables)
	return rv
}

func (pc _PredicateClass) Alloc() Predicate {
	rv := ffi.CallMethod[Predicate](pc, "alloc")
	return rv
}

func (pc _PredicateClass) New() Predicate {
	rv := ffi.CallMethod[Predicate](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPredicate() Predicate {
	return PredicateClass.New()
}

func (p_ Predicate) Init() Predicate {
	rv := ffi.CallMethod[Predicate](p_, "init")
	return rv
}

func (pc _PredicateClass) PredicateWithFormat_ArgumentArray(predicateFormat string, arguments []objc.IObject) Predicate {
	rv := ffi.CallMethod[Predicate](pc, "predicateWithFormat:argumentArray:", predicateFormat, arguments)
	return rv
}

func (pc _PredicateClass) PredicateWithValue(value bool) Predicate {
	rv := ffi.CallMethod[Predicate](pc, "predicateWithValue:", value)
	return rv
}

func (pc _PredicateClass) PredicateFromMetadataQueryString(queryString string) Predicate {
	rv := ffi.CallMethod[Predicate](pc, "predicateFromMetadataQueryString:", queryString)
	return rv
}

func (p_ Predicate) EvaluateWithObject(object objc.IObject) bool {
	rv := ffi.CallMethod[bool](p_, "evaluateWithObject:", object)
	return rv
}

func (p_ Predicate) EvaluateWithObject_SubstitutionVariables(object objc.IObject, bindings map[string]objc.IObject) bool {
	rv := ffi.CallMethod[bool](p_, "evaluateWithObject:substitutionVariables:", object, bindings)
	return rv
}

func (p_ Predicate) AllowEvaluation() {
	ffi.CallMethod[ffi.Void](p_, "allowEvaluation")
}

func (p_ Predicate) PredicateFormat() string {
	rv := ffi.CallMethod[string](p_, "predicateFormat")
	return rv
}
