// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var NumberClass = _NumberClass{objc.GetClass("NSNumber")}

type _NumberClass struct {
	objc.Class
}

type INumber interface {
	IValue
	InitWithBool(value bool) Number
	InitWithChar(value byte) Number
	InitWithDouble(value float64) Number
	InitWithFloat(value float32) Number
	InitWithInt(value int32) Number
	InitWithInteger(value int) Number
	InitWithLong(value int64) Number
	InitWithLongLong(value int64) Number
	InitWithShort(value int16) Number
	InitWithUnsignedChar(value byte) Number
	InitWithUnsignedInt(value uint32) Number
	InitWithUnsignedInteger(value uint) Number
	InitWithUnsignedLong(value uint64) Number
	InitWithUnsignedLongLong(value uint64) Number
	InitWithUnsignedShort(value uint16) Number
	DescriptionWithLocale(locale objc.IObject) string
	Compare(otherNumber INumber) ComparisonResult
	IsEqualToNumber(number INumber) bool
	BoolValue() bool
	CharValue() byte
	DecimalValue() Decimal
	DoubleValue() float64
	FloatValue() float32
	IntValue() int32
	IntegerValue() int
	LongLongValue() int64
	LongValue() int64
	ShortValue() int16
	UnsignedCharValue() byte
	UnsignedIntegerValue() uint
	UnsignedIntValue() uint32
	UnsignedLongLongValue() uint64
	UnsignedLongValue() uint64
	UnsignedShortValue() uint16
	StringValue() string
}

type Number struct {
	Value
}

func MakeNumber(ptr unsafe.Pointer) Number {
	return Number{
		Value: MakeValue(ptr),
	}
}

func (n_ Number) InitWithBytes_ObjCType(value unsafe.Pointer, type_ *byte) Number {
	rv := ffi.CallMethod[Number](n_, "initWithBytes:objCType:", value, type_)
	return rv
}

func (nc _NumberClass) Alloc() Number {
	rv := ffi.CallMethod[Number](nc, "alloc")
	return rv
}

func (nc _NumberClass) New() Number {
	rv := ffi.CallMethod[Number](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNumber() Number {
	return NumberClass.New()
}

func (n_ Number) Init() Number {
	rv := ffi.CallMethod[Number](n_, "init")
	return rv
}

func (nc _NumberClass) NumberWithBool(value bool) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithBool:", value)
	return rv
}

func (nc _NumberClass) NumberWithChar(value byte) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithChar:", value)
	return rv
}

func (nc _NumberClass) NumberWithDouble(value float64) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithDouble:", value)
	return rv
}

func (nc _NumberClass) NumberWithFloat(value float32) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithFloat:", value)
	return rv
}

func (nc _NumberClass) NumberWithInt(value int32) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithInt:", value)
	return rv
}

func (nc _NumberClass) NumberWithInteger(value int) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithInteger:", value)
	return rv
}

func (nc _NumberClass) NumberWithLong(value int64) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithLong:", value)
	return rv
}

func (nc _NumberClass) NumberWithLongLong(value int64) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithLongLong:", value)
	return rv
}

func (nc _NumberClass) NumberWithShort(value int16) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithShort:", value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedChar(value byte) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithUnsignedChar:", value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedInt(value uint32) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithUnsignedInt:", value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedInteger(value uint) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithUnsignedInteger:", value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedLong(value uint64) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithUnsignedLong:", value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedLongLong(value uint64) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithUnsignedLongLong:", value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedShort(value uint16) Number {
	rv := ffi.CallMethod[Number](nc, "numberWithUnsignedShort:", value)
	return rv
}

func (n_ Number) InitWithBool(value bool) Number {
	rv := ffi.CallMethod[Number](n_, "initWithBool:", value)
	return rv
}

func (n_ Number) InitWithChar(value byte) Number {
	rv := ffi.CallMethod[Number](n_, "initWithChar:", value)
	return rv
}

func (n_ Number) InitWithDouble(value float64) Number {
	rv := ffi.CallMethod[Number](n_, "initWithDouble:", value)
	return rv
}

func (n_ Number) InitWithFloat(value float32) Number {
	rv := ffi.CallMethod[Number](n_, "initWithFloat:", value)
	return rv
}

func (n_ Number) InitWithInt(value int32) Number {
	rv := ffi.CallMethod[Number](n_, "initWithInt:", value)
	return rv
}

func (n_ Number) InitWithInteger(value int) Number {
	rv := ffi.CallMethod[Number](n_, "initWithInteger:", value)
	return rv
}

func (n_ Number) InitWithLong(value int64) Number {
	rv := ffi.CallMethod[Number](n_, "initWithLong:", value)
	return rv
}

func (n_ Number) InitWithLongLong(value int64) Number {
	rv := ffi.CallMethod[Number](n_, "initWithLongLong:", value)
	return rv
}

func (n_ Number) InitWithShort(value int16) Number {
	rv := ffi.CallMethod[Number](n_, "initWithShort:", value)
	return rv
}

func (n_ Number) InitWithUnsignedChar(value byte) Number {
	rv := ffi.CallMethod[Number](n_, "initWithUnsignedChar:", value)
	return rv
}

func (n_ Number) InitWithUnsignedInt(value uint32) Number {
	rv := ffi.CallMethod[Number](n_, "initWithUnsignedInt:", value)
	return rv
}

func (n_ Number) InitWithUnsignedInteger(value uint) Number {
	rv := ffi.CallMethod[Number](n_, "initWithUnsignedInteger:", value)
	return rv
}

func (n_ Number) InitWithUnsignedLong(value uint64) Number {
	rv := ffi.CallMethod[Number](n_, "initWithUnsignedLong:", value)
	return rv
}

func (n_ Number) InitWithUnsignedLongLong(value uint64) Number {
	rv := ffi.CallMethod[Number](n_, "initWithUnsignedLongLong:", value)
	return rv
}

func (n_ Number) InitWithUnsignedShort(value uint16) Number {
	rv := ffi.CallMethod[Number](n_, "initWithUnsignedShort:", value)
	return rv
}

func (n_ Number) DescriptionWithLocale(locale objc.IObject) string {
	rv := ffi.CallMethod[string](n_, "descriptionWithLocale:", locale)
	return rv
}

func (n_ Number) Compare(otherNumber INumber) ComparisonResult {
	rv := ffi.CallMethod[ComparisonResult](n_, "compare:", otherNumber)
	return rv
}

func (n_ Number) IsEqualToNumber(number INumber) bool {
	rv := ffi.CallMethod[bool](n_, "isEqualToNumber:", number)
	return rv
}

func (n_ Number) BoolValue() bool {
	rv := ffi.CallMethod[bool](n_, "boolValue")
	return rv
}

func (n_ Number) CharValue() byte {
	rv := ffi.CallMethod[byte](n_, "charValue")
	return rv
}

func (n_ Number) DecimalValue() Decimal {
	rv := ffi.CallMethod[Decimal](n_, "decimalValue")
	return rv
}

func (n_ Number) DoubleValue() float64 {
	rv := ffi.CallMethod[float64](n_, "doubleValue")
	return rv
}

func (n_ Number) FloatValue() float32 {
	rv := ffi.CallMethod[float32](n_, "floatValue")
	return rv
}

func (n_ Number) IntValue() int32 {
	rv := ffi.CallMethod[int32](n_, "intValue")
	return rv
}

func (n_ Number) IntegerValue() int {
	rv := ffi.CallMethod[int](n_, "integerValue")
	return rv
}

func (n_ Number) LongLongValue() int64 {
	rv := ffi.CallMethod[int64](n_, "longLongValue")
	return rv
}

func (n_ Number) LongValue() int64 {
	rv := ffi.CallMethod[int64](n_, "longValue")
	return rv
}

func (n_ Number) ShortValue() int16 {
	rv := ffi.CallMethod[int16](n_, "shortValue")
	return rv
}

func (n_ Number) UnsignedCharValue() byte {
	rv := ffi.CallMethod[byte](n_, "unsignedCharValue")
	return rv
}

func (n_ Number) UnsignedIntegerValue() uint {
	rv := ffi.CallMethod[uint](n_, "unsignedIntegerValue")
	return rv
}

func (n_ Number) UnsignedIntValue() uint32 {
	rv := ffi.CallMethod[uint32](n_, "unsignedIntValue")
	return rv
}

func (n_ Number) UnsignedLongLongValue() uint64 {
	rv := ffi.CallMethod[uint64](n_, "unsignedLongLongValue")
	return rv
}

func (n_ Number) UnsignedLongValue() uint64 {
	rv := ffi.CallMethod[uint64](n_, "unsignedLongValue")
	return rv
}

func (n_ Number) UnsignedShortValue() uint16 {
	rv := ffi.CallMethod[uint16](n_, "unsignedShortValue")
	return rv
}

func (n_ Number) StringValue() string {
	rv := ffi.CallMethod[string](n_, "stringValue")
	return rv
}
