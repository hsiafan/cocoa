// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[DecimalNumber](d_, "initWithDecimal:", dcm)
	return rv
}

func (d_ DecimalNumber) InitWithMantissa_Exponent_IsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "initWithMantissa:exponent:isNegative:", mantissa, exponent, flag)
	return rv
}

func (d_ DecimalNumber) InitWithString(numberValue string) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "initWithString:", numberValue)
	return rv
}

func (d_ DecimalNumber) InitWithString_Locale(numberValue string, locale objc.IObject) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "initWithString:locale:", numberValue, locale)
	return rv
}

func (d_ DecimalNumber) InitWithBytes_ObjCType(value unsafe.Pointer, type_ *byte) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "initWithBytes:objCType:", value, type_)
	return rv
}

func (dc _DecimalNumberClass) Alloc() DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "alloc")
	return rv
}

func (dc _DecimalNumberClass) New() DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDecimalNumber() DecimalNumber {
	return DecimalNumberClass.New()
}

func (d_ DecimalNumber) Init() DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "init")
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithDecimal(dcm Decimal) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "decimalNumberWithDecimal:", dcm)
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithMantissa_Exponent_IsNegative(mantissa uint64, exponent int16, flag bool) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "decimalNumberWithMantissa:exponent:isNegative:", mantissa, exponent, flag)
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithString(numberValue string) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "decimalNumberWithString:", numberValue)
	return rv
}

func (dc _DecimalNumberClass) DecimalNumberWithString_Locale(numberValue string, locale objc.IObject) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "decimalNumberWithString:locale:", numberValue, locale)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByAdding(decimalNumber IDecimalNumber) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "decimalNumberByAdding:", decimalNumber)
	return rv
}

func (d_ DecimalNumber) DecimalNumberBySubtracting(decimalNumber IDecimalNumber) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "decimalNumberBySubtracting:", decimalNumber)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByMultiplyingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "decimalNumberByMultiplyingBy:", decimalNumber)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByDividingBy(decimalNumber IDecimalNumber) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "decimalNumberByDividingBy:", decimalNumber)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByRaisingToPower(power uint) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "decimalNumberByRaisingToPower:", power)
	return rv
}

func (d_ DecimalNumber) DecimalNumberByMultiplyingByPowerOf10(power int16) DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](d_, "decimalNumberByMultiplyingByPowerOf10:", power)
	return rv
}

func (dc _DecimalNumberClass) One() DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "one")
	return rv
}

func (dc _DecimalNumberClass) Zero() DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "zero")
	return rv
}

func (dc _DecimalNumberClass) NotANumber() DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "notANumber")
	return rv
}

func (dc _DecimalNumberClass) MaximumDecimalNumber() DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "maximumDecimalNumber")
	return rv
}

func (dc _DecimalNumberClass) MinimumDecimalNumber() DecimalNumber {
	rv := ffi.CallMethod[DecimalNumber](dc, "minimumDecimalNumber")
	return rv
}
