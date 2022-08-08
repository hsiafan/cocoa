// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ExpressionClass = _ExpressionClass{objc.GetClass("NSExpression")}

type _ExpressionClass struct {
	objc.Class
}

type IExpression interface {
	objc.IObject
	ExpressionValueWithObject_Context(object objc.IObject, context IMutableDictionary) objc.Object
	AllowEvaluation()
	Arguments() []Expression
	Collection() objc.Object
	ConstantValue() objc.Object
	ExpressionType() ExpressionType
	Function() string
	KeyPath() string
	Operand() Expression
	Predicate() Predicate
	LeftExpression() Expression
	RightExpression() Expression
	Variable() string
	FalseExpression() Expression
	TrueExpression() Expression
	ExpressionBlock() func(param1 objc.IObject, param2 []IExpression, param3 IMutableDictionary) objc.Object
}

type Expression struct {
	objc.Object
}

func MakeExpression(ptr unsafe.Pointer) Expression {
	return Expression{
		Object: objc.MakeObject(ptr),
	}
}

func (e_ Expression) InitWithExpressionType(_type ExpressionType) Expression {
	rv := ffi.CallMethod[Expression](e_, "initWithExpressionType:", _type)
	return rv
}

func (ec _ExpressionClass) Alloc() Expression {
	rv := ffi.CallMethod[Expression](ec, "alloc")
	return rv
}

func (e_ Expression) Init() Expression {
	rv := ffi.CallMethod[Expression](e_, "init")
	return rv
}

func (ec _ExpressionClass) New() Expression {
	rv := ffi.CallMethod[Expression](ec, "new")
	rv.Autorelease()
	return rv
}

func NewExpression() Expression {
	return ExpressionClass.New()
}

func (ec _ExpressionClass) ExpressionWithFormat_ArgumentArray(expressionFormat string, arguments []objc.IObject) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionWithFormat:argumentArray:", expressionFormat, arguments)
	return rv
}

func (ec _ExpressionClass) ExpressionForConstantValue(obj objc.IObject) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForConstantValue:", obj)
	return rv
}

func (ec _ExpressionClass) ExpressionForEvaluatedObject() Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForEvaluatedObject")
	return rv
}

func (ec _ExpressionClass) ExpressionForKeyPath(keyPath string) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForKeyPath:", keyPath)
	return rv
}

func (ec _ExpressionClass) ExpressionForVariable(_string string) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForVariable:", _string)
	return rv
}

func (ec _ExpressionClass) ExpressionForAnyKey() Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForAnyKey")
	return rv
}

func (ec _ExpressionClass) ExpressionForAggregate(subexpressions []IExpression) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForAggregate:", subexpressions)
	return rv
}

func (ec _ExpressionClass) ExpressionForUnionSet_With(left IExpression, right IExpression) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForUnionSet:with:", left, right)
	return rv
}

func (ec _ExpressionClass) ExpressionForIntersectSet_With(left IExpression, right IExpression) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForIntersectSet:with:", left, right)
	return rv
}

func (ec _ExpressionClass) ExpressionForMinusSet_With(left IExpression, right IExpression) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForMinusSet:with:", left, right)
	return rv
}

func (ec _ExpressionClass) ExpressionForSubquery_UsingIteratorVariable_Predicate(expression IExpression, variable string, predicate IPredicate) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForSubquery:usingIteratorVariable:predicate:", expression, variable, predicate)
	return rv
}

func (ec _ExpressionClass) ExpressionForConditional_TrueExpression_FalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForConditional:trueExpression:falseExpression:", predicate, trueExpression, falseExpression)
	return rv
}

func (ec _ExpressionClass) ExpressionForBlock_Arguments(block func(evaluatedObject objc.Object, expressions []Expression, context MutableDictionary) objc.IObject, arguments []IExpression) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForBlock:arguments:", block, arguments)
	return rv
}

func (ec _ExpressionClass) ExpressionForFunction_Arguments(name string, parameters []objc.IObject) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForFunction:arguments:", name, parameters)
	return rv
}

func (ec _ExpressionClass) ExpressionForFunction_SelectorName_Arguments(target IExpression, name string, parameters []objc.IObject) Expression {
	rv := ffi.CallMethod[Expression](ec, "expressionForFunction:selectorName:arguments:", target, name, parameters)
	return rv
}

func (e_ Expression) ExpressionValueWithObject_Context(object objc.IObject, context IMutableDictionary) objc.Object {
	rv := ffi.CallMethod[objc.Object](e_, "expressionValueWithObject:context:", object, context)
	return rv
}

func (e_ Expression) AllowEvaluation() {
	ffi.CallMethod[ffi.Void](e_, "allowEvaluation")
}

func (e_ Expression) Arguments() []Expression {
	rv := ffi.CallMethod[[]Expression](e_, "arguments")
	return rv
}

func (e_ Expression) Collection() objc.Object {
	rv := ffi.CallMethod[objc.Object](e_, "collection")
	return rv
}

func (e_ Expression) ConstantValue() objc.Object {
	rv := ffi.CallMethod[objc.Object](e_, "constantValue")
	return rv
}

func (e_ Expression) ExpressionType() ExpressionType {
	rv := ffi.CallMethod[ExpressionType](e_, "expressionType")
	return rv
}

func (e_ Expression) Function() string {
	rv := ffi.CallMethod[string](e_, "function")
	return rv
}

func (e_ Expression) KeyPath() string {
	rv := ffi.CallMethod[string](e_, "keyPath")
	return rv
}

func (e_ Expression) Operand() Expression {
	rv := ffi.CallMethod[Expression](e_, "operand")
	return rv
}

func (e_ Expression) Predicate() Predicate {
	rv := ffi.CallMethod[Predicate](e_, "predicate")
	return rv
}

func (e_ Expression) LeftExpression() Expression {
	rv := ffi.CallMethod[Expression](e_, "leftExpression")
	return rv
}

func (e_ Expression) RightExpression() Expression {
	rv := ffi.CallMethod[Expression](e_, "rightExpression")
	return rv
}

func (e_ Expression) Variable() string {
	rv := ffi.CallMethod[string](e_, "variable")
	return rv
}

func (e_ Expression) FalseExpression() Expression {
	rv := ffi.CallMethod[Expression](e_, "falseExpression")
	return rv
}

func (e_ Expression) TrueExpression() Expression {
	rv := ffi.CallMethod[Expression](e_, "trueExpression")
	return rv
}

func (e_ Expression) ExpressionBlock() func(param1 objc.IObject, param2 []IExpression, param3 IMutableDictionary) objc.Object {
	rv := ffi.CallMethod[func(param1 objc.IObject, param2 []IExpression, param3 IMutableDictionary) objc.Object](e_, "expressionBlock")
	return rv
}

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

func (p_ Predicate) Init() Predicate {
	rv := ffi.CallMethod[Predicate](p_, "init")
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

func (pc _PredicateClass) PredicateWithFormat_ArgumentArray(predicateFormat string, arguments []objc.IObject) Predicate {
	rv := ffi.CallMethod[Predicate](pc, "predicateWithFormat:argumentArray:", predicateFormat, arguments)
	return rv
}

func (pc _PredicateClass) PredicateWithValue(value bool) Predicate {
	rv := ffi.CallMethod[Predicate](pc, "predicateWithValue:", value)
	return rv
}

func (pc _PredicateClass) PredicateWithBlock(block func(evaluatedObject objc.Object, bindings map[string]objc.Object) bool) Predicate {
	rv := ffi.CallMethod[Predicate](pc, "predicateWithBlock:", block)
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
