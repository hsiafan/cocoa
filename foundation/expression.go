// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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

func (e_ Expression) InitWithExpressionType(type_ ExpressionType) Expression {
	rv := objc.CallMethod[Expression](e_, objc.SEL("initWithExpressionType:"), type_)
	return rv
}

func (ec _ExpressionClass) Alloc() Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("alloc"))
	return rv
}

func (ec _ExpressionClass) New() Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewExpression() Expression {
	return ExpressionClass.New()
}

func (e_ Expression) Init() Expression {
	rv := objc.CallMethod[Expression](e_, objc.SEL("init"))
	return rv
}

func (ec _ExpressionClass) ExpressionWithFormat_ArgumentArray(expressionFormat string, arguments []objc.IObject) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionWithFormat:argumentArray:"), expressionFormat, arguments)
	return rv
}

func (ec _ExpressionClass) ExpressionForConstantValue(obj objc.IObject) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForConstantValue:"), objc.ExtractPtr(obj))
	return rv
}

func (ec _ExpressionClass) ExpressionForEvaluatedObject() Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForEvaluatedObject"))
	return rv
}

func (ec _ExpressionClass) ExpressionForKeyPath(keyPath string) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForKeyPath:"), keyPath)
	return rv
}

func (ec _ExpressionClass) ExpressionForVariable(string_ string) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForVariable:"), string_)
	return rv
}

func (ec _ExpressionClass) ExpressionForAnyKey() Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForAnyKey"))
	return rv
}

func (ec _ExpressionClass) ExpressionForAggregate(subexpressions []IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForAggregate:"), subexpressions)
	return rv
}

func (ec _ExpressionClass) ExpressionForUnionSet_With(left IExpression, right IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForUnionSet:with:"), objc.ExtractPtr(left), objc.ExtractPtr(right))
	return rv
}

func (ec _ExpressionClass) ExpressionForIntersectSet_With(left IExpression, right IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForIntersectSet:with:"), objc.ExtractPtr(left), objc.ExtractPtr(right))
	return rv
}

func (ec _ExpressionClass) ExpressionForMinusSet_With(left IExpression, right IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForMinusSet:with:"), objc.ExtractPtr(left), objc.ExtractPtr(right))
	return rv
}

func (ec _ExpressionClass) ExpressionForSubquery_UsingIteratorVariable_Predicate(expression IExpression, variable string, predicate IPredicate) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForSubquery:usingIteratorVariable:predicate:"), objc.ExtractPtr(expression), variable, objc.ExtractPtr(predicate))
	return rv
}

func (ec _ExpressionClass) ExpressionForConditional_TrueExpression_FalseExpression(predicate IPredicate, trueExpression IExpression, falseExpression IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForConditional:trueExpression:falseExpression:"), objc.ExtractPtr(predicate), objc.ExtractPtr(trueExpression), objc.ExtractPtr(falseExpression))
	return rv
}

func (ec _ExpressionClass) ExpressionForBlock_Arguments(block func(evaluatedObject objc.Object, expressions []Expression, context MutableDictionary) objc.IObject, arguments []IExpression) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForBlock:arguments:"), block, arguments)
	return rv
}

func (ec _ExpressionClass) ExpressionForFunction_Arguments(name string, parameters []objc.IObject) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForFunction:arguments:"), name, parameters)
	return rv
}

func (ec _ExpressionClass) ExpressionForFunction_SelectorName_Arguments(target IExpression, name string, parameters []objc.IObject) Expression {
	rv := objc.CallMethod[Expression](ec, objc.SEL("expressionForFunction:selectorName:arguments:"), objc.ExtractPtr(target), name, parameters)
	return rv
}

func (e_ Expression) ExpressionValueWithObject_Context(object objc.IObject, context IMutableDictionary) objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.SEL("expressionValueWithObject:context:"), objc.ExtractPtr(object), objc.ExtractPtr(context))
	return rv
}

func (e_ Expression) AllowEvaluation() {
	objc.CallMethod[objc.Void](e_, objc.SEL("allowEvaluation"))
}

func (e_ Expression) Arguments() []Expression {
	rv := objc.CallMethod[[]Expression](e_, objc.SEL("arguments"))
	return rv
}

func (e_ Expression) Collection() objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.SEL("collection"))
	return rv
}

func (e_ Expression) ConstantValue() objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.SEL("constantValue"))
	return rv
}

func (e_ Expression) ExpressionType() ExpressionType {
	rv := objc.CallMethod[ExpressionType](e_, objc.SEL("expressionType"))
	return rv
}

func (e_ Expression) Function() string {
	rv := objc.CallMethod[string](e_, objc.SEL("function"))
	return rv
}

func (e_ Expression) KeyPath() string {
	rv := objc.CallMethod[string](e_, objc.SEL("keyPath"))
	return rv
}

func (e_ Expression) Operand() Expression {
	rv := objc.CallMethod[Expression](e_, objc.SEL("operand"))
	return rv
}

func (e_ Expression) Predicate() Predicate {
	rv := objc.CallMethod[Predicate](e_, objc.SEL("predicate"))
	return rv
}

func (e_ Expression) LeftExpression() Expression {
	rv := objc.CallMethod[Expression](e_, objc.SEL("leftExpression"))
	return rv
}

func (e_ Expression) RightExpression() Expression {
	rv := objc.CallMethod[Expression](e_, objc.SEL("rightExpression"))
	return rv
}

func (e_ Expression) Variable() string {
	rv := objc.CallMethod[string](e_, objc.SEL("variable"))
	return rv
}

func (e_ Expression) FalseExpression() Expression {
	rv := objc.CallMethod[Expression](e_, objc.SEL("falseExpression"))
	return rv
}

func (e_ Expression) TrueExpression() Expression {
	rv := objc.CallMethod[Expression](e_, objc.SEL("trueExpression"))
	return rv
}

func (e_ Expression) ExpressionBlock() func(param1 objc.IObject, param2 []IExpression, param3 IMutableDictionary) objc.Object {
	rv := objc.CallMethod[func(param1 objc.IObject, param2 []IExpression, param3 IMutableDictionary) objc.Object](e_, objc.SEL("expressionBlock"))
	return rv
}
