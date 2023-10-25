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
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("initWithDecimal:"), dcm)
	return rv
}

func (d_ DecimalNumber) InitWithMantissa_Exponent_IsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("initWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

func (d_ DecimalNumber) InitWithString(numberValue string) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("initWithString:"), numberValue)
	return rv
}

func (d_ DecimalNumber) InitWithString_Locale(numberValue string, locale objc.IObject) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("initWithString:locale:"), numberValue, objc.ExtractPtr(locale))
	return rv
}

func (d_ DecimalNumber) InitWithBytes_ObjCType(value unsafe.Pointer, type_ *byte) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("initWithBytes:objCType:"), value, type_)
	return rv
}

func (dc _DecimalNumberClass) Alloc() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DecimalNumberClass) New() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDecimalNumber() DecimalNumber {
	return DecimalNumberClass.New()
}

func (d_ DecimalNumber) Init() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("init"))
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithDecimal(dcm Decimal) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("decimalNumberWithDecimal:"), dcm)
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithMantissa_Exponent_IsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("decimalNumberWithMantissa:exponent:isNegative:"), mantissa, exponent, flag)
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithString(numberValue string) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("decimalNumberWithString:"), numberValue)
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithString_Locale(numberValue string, locale objc.IObject) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("decimalNumberWithString:locale:"), numberValue, objc.ExtractPtr(locale))
	return rv
}

func (d_ DecimalNumber) DecimalNumberByAdding(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("decimalNumberByAdding:"), objc.ExtractPtr(decimalNumber))
	return rv
}

func (d_ DecimalNumber) DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("decimalNumberBySubtracting:"), objc.ExtractPtr(decimalNumber))
	return rv
}

func (d_ DecimalNumber) DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("decimalNumberByMultiplyingBy:"), objc.ExtractPtr(decimalNumber))
	return rv
}

func (d_ DecimalNumber) DecimalNumberByDividingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("decimalNumberByDividingBy:"), objc.ExtractPtr(decimalNumber))
	return rv
}

func (d_ DecimalNumber) DecimalNumberByRaisingToPower(power uint) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("decimalNumberByRaisingToPower:"), power)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10(power int16) DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](d_, objc.SEL("decimalNumberByMultiplyingByPowerOf10:"), power)
	return rv
}

func (dc _DecimalNumberClass) One() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("one"))
	return rv
}

func (dc _DecimalNumberClass) Zero() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("zero"))
	return rv
}

func (dc _DecimalNumberClass) NotANumber() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("notANumber"))
	return rv
}

func (dc _DecimalNumberClass) MaximumDecimalNumber() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("maximumDecimalNumber"))
	return rv
}

func (dc _DecimalNumberClass) MinimumDecimalNumber() DecimalNumber {
	rv := objc.CallMethod[DecimalNumber](dc, objc.SEL("minimumDecimalNumber"))
	return rv
}
