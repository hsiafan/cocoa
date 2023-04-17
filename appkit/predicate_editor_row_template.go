// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PredicateEditorRowTemplateClass = _PredicateEditorRowTemplateClass{objc.GetClass("NSPredicateEditorRowTemplate")}

type _PredicateEditorRowTemplateClass struct {
	objc.Class
}

type IPredicateEditorRowTemplate interface {
	objc.IObject
	MatchForPredicate(predicate foundation.IPredicate) float64
	SetPredicate(predicate foundation.IPredicate)
	DisplayableSubpredicatesOfPredicate(predicate foundation.IPredicate) []foundation.Predicate
	PredicateWithSubpredicates(subpredicates []foundation.IPredicate) foundation.Predicate
	TemplateViews() []View
	LeftExpressions() []foundation.Expression
	RightExpressions() []foundation.Expression
	CompoundTypes() []foundation.Number
	Modifier() foundation.ComparisonPredicateModifier
	Operators() []foundation.Number
	Options() uint
}

type PredicateEditorRowTemplate struct {
	objc.Object
}

func MakePredicateEditorRowTemplate(ptr unsafe.Pointer) PredicateEditorRowTemplate {
	return PredicateEditorRowTemplate{
		Object: objc.MakeObject(ptr),
	}
}

func (p_ PredicateEditorRowTemplate) InitWithLeftExpressions_RightExpressions_Modifier_Operators_Options(leftExpressions []foundation.IExpression, rightExpressions []foundation.IExpression, modifier foundation.ComparisonPredicateModifier, operators []foundation.INumber, options uint) PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](p_, "initWithLeftExpressions:rightExpressions:modifier:operators:options:", leftExpressions, rightExpressions, modifier, operators, options)
	return rv
}

func (p_ PredicateEditorRowTemplate) InitWithCompoundTypes(compoundTypes []foundation.INumber) PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](p_, "initWithCompoundTypes:", compoundTypes)
	return rv
}

func (pc _PredicateEditorRowTemplateClass) Alloc() PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](pc, "alloc")
	return rv
}

func (pc _PredicateEditorRowTemplateClass) New() PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPredicateEditorRowTemplate() PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.New()
}

func (p_ PredicateEditorRowTemplate) Init() PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](p_, "init")
	return rv
}

func (p_ PredicateEditorRowTemplate) MatchForPredicate(predicate foundation.IPredicate) float64 {
	rv := ffi.CallMethod[float64](p_, "matchForPredicate:", predicate)
	return rv
}

func (p_ PredicateEditorRowTemplate) SetPredicate(predicate foundation.IPredicate) {
	ffi.CallMethod[ffi.Void](p_, "setPredicate:", predicate)
}

func (p_ PredicateEditorRowTemplate) DisplayableSubpredicatesOfPredicate(predicate foundation.IPredicate) []foundation.Predicate {
	rv := ffi.CallMethod[[]foundation.Predicate](p_, "displayableSubpredicatesOfPredicate:", predicate)
	return rv
}

func (p_ PredicateEditorRowTemplate) PredicateWithSubpredicates(subpredicates []foundation.IPredicate) foundation.Predicate {
	rv := ffi.CallMethod[foundation.Predicate](p_, "predicateWithSubpredicates:", subpredicates)
	return rv
}

func (p_ PredicateEditorRowTemplate) TemplateViews() []View {
	rv := ffi.CallMethod[[]View](p_, "templateViews")
	return rv
}

func (p_ PredicateEditorRowTemplate) LeftExpressions() []foundation.Expression {
	rv := ffi.CallMethod[[]foundation.Expression](p_, "leftExpressions")
	return rv
}

func (p_ PredicateEditorRowTemplate) RightExpressions() []foundation.Expression {
	rv := ffi.CallMethod[[]foundation.Expression](p_, "rightExpressions")
	return rv
}

func (p_ PredicateEditorRowTemplate) CompoundTypes() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](p_, "compoundTypes")
	return rv
}

func (p_ PredicateEditorRowTemplate) Modifier() foundation.ComparisonPredicateModifier {
	rv := ffi.CallMethod[foundation.ComparisonPredicateModifier](p_, "modifier")
	return rv
}

func (p_ PredicateEditorRowTemplate) Operators() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](p_, "operators")
	return rv
}

func (p_ PredicateEditorRowTemplate) Options() uint {
	rv := ffi.CallMethod[uint](p_, "options")
	return rv
}
