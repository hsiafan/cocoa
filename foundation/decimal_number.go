// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var DecimalNumberClass = _DecimalNumberClass{objc.GetClass("NSDecimalNumber")}

type _DecimalNumberClass struct {
	objc.Class
}

type IDecimalNumber interface {
	INumber
	DecimalNumberByAdding(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByDividingBy(decimalNumber IDecimalNumber) DecimalNumber
	DecimalNumberByRaisingToPower(power uint) DecimalNumber
	DecimalNumberByMultiplyingByPowerOf10(power int16) DecimalNumber
}

type DecimalNumber struct {
	Number
}

func MakeDecimalNumber(ptr unsafe.Pointer) DecimalNumber {
	return DecimalNumber{
		Number: MakeNumber(ptr),
	}
}

func (d_ DecimalNumber) InitWithDecimal(dcm Decimal) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithDecimal:"), dcm)
	return rv
}

func (d_ DecimalNumber) InitWithMantissa_Exponent_IsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

func (d_ DecimalNumber) InitWithString(numberValue string) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithString:"), numberValue)
	return rv
}

func (d_ DecimalNumber) InitWithString_Locale(numberValue string, locale objc.IObject) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithString:locale:"), numberValue, locale)
	return rv
}

func (d_ DecimalNumber) InitWithBytes_ObjCType(value unsafe.Pointer, type_ *byte) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("initWithBytes:objCType:"), value, type_)
	return rv
}

func (dc _DecimalNumberClass) Alloc() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("alloc"))
	return rv
}

func (dc _DecimalNumberClass) New() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDecimalNumber() DecimalNumber {
	return DecimalNumberClass.New()
}

func (d_ DecimalNumber) Init() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("init"))
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithDecimal(dcm Decimal) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("decimalNumberWithDecimal:"), dcm)
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithMantissa_Exponent_IsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("decimalNumberWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithString(numberValue string) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("decimalNumberWithString:"), numberValue)
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithString_Locale(numberValue string, locale objc.IObject) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("decimalNumberWithString:locale:"), numberValue, locale)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByAdding(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByAdding:"), decimalNumber)
	return rv
}

func (d_ DecimalNumber) DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberBySubtracting:"), decimalNumber)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByMultiplyingBy:"), decimalNumber)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByDividingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByDividingBy:"), decimalNumber)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByRaisingToPower(power uint) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByRaisingToPower:"), power)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10(power int16) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.GetSelector("decimalNumberByMultiplyingByPowerOf10:"), power)
	return rv
}

func (dc _DecimalNumberClass) One() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("one"))
	return rv
}

func (dc _DecimalNumberClass) Zero() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("zero"))
	return rv
}

func (dc _DecimalNumberClass) NotANumber() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("notANumber"))
	return rv
}

func (dc _DecimalNumberClass) MaximumDecimalNumber() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("maximumDecimalNumber"))
	return rv
}

func (dc _DecimalNumberClass) MinimumDecimalNumber() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.GetSelector("minimumDecimalNumber"))
	return rv
}
