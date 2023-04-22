// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var DecimalNumberHandlerClass = _DecimalNumberHandlerClass{objc.GetClass("NSDecimalNumberHandler")}

type _DecimalNumberHandlerClass struct {
	objc.Class
}

type IDecimalNumberHandler interface {
	objc.IObject
}

type DecimalNumberHandler struct {
	objc.Object
}

func MakeDecimalNumberHandler(ptr unsafe.Pointer) DecimalNumberHandler {
	return DecimalNumberHandler{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DecimalNumberHandlerClass) DecimalNumberHandlerWithRoundingMode_Scale_RaiseOnExactness_RaiseOnOverflow_RaiseOnUnderflow_RaiseOnDivideByZero(roundingMode RoundingMode, scale int16, exact bool, overflow bool, underflow bool, divideByZero bool) DecimalNumberHandler {
	rv := objc.CallMethod[DecimalNumberHandler](dc, objc.GetSelector("decimalNumberHandlerWithRoundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:"), roundingMode, scale, exact, overflow, underflow, divideByZero)
	return rv
}

func (d_ DecimalNumberHandler) InitWithRoundingMode_Scale_RaiseOnExactness_RaiseOnOverflow_RaiseOnUnderflow_RaiseOnDivideByZero(roundingMode RoundingMode, scale int16, exact bool, overflow bool, underflow bool, divideByZero bool) DecimalNumberHandler {
	rv := objc.CallMethod[DecimalNumberHandler](d_, objc.GetSelector("initWithRoundingMode:scale:raiseOnExactness:raiseOnOverflow:raiseOnUnderflow:raiseOnDivideByZero:"), roundingMode, scale, exact, overflow, underflow, divideByZero)
	return rv
}

func (dc _DecimalNumberHandlerClass) Alloc() DecimalNumberHandler {
	rv := objc.CallMethod[DecimalNumberHandler](dc, objc.GetSelector("alloc"))
	return rv
}

func (dc _DecimalNumberHandlerClass) New() DecimalNumberHandler {
	rv := objc.CallMethod[DecimalNumberHandler](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDecimalNumberHandler() DecimalNumberHandler {
	return DecimalNumberHandlerClass.New()
}

func (d_ DecimalNumberHandler) Init() DecimalNumberHandler {
	rv := objc.CallMethod[DecimalNumberHandler](d_, objc.GetSelector("init"))
	return rv
}

func (dc _DecimalNumberHandlerClass) DefaultDecimalNumberHandler() DecimalNumberHandler {
	rv := objc.CallMethod[DecimalNumberHandler](dc, objc.GetSelector("defaultDecimalNumberHandler"))
	return rv
}
