// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithBytes:objCType:"), value, type_)
	return rv
}

func (nc _NumberClass) Alloc() Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("alloc"))
	return rv
}

func (nc _NumberClass) New() Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewNumber() Number {
	return NumberClass.New()
}

func (n_ Number) Init() Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("init"))
	return rv
}

func (nc _NumberClass) NumberWithBool(value bool) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithBool:"), value)
	return rv
}

func (nc _NumberClass) NumberWithChar(value byte) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithChar:"), value)
	return rv
}

func (nc _NumberClass) NumberWithDouble(value float64) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithDouble:"), value)
	return rv
}

func (nc _NumberClass) NumberWithFloat(value float32) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithFloat:"), value)
	return rv
}

func (nc _NumberClass) NumberWithInt(value int32) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithInt:"), value)
	return rv
}

func (nc _NumberClass) NumberWithInteger(value int) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithInteger:"), value)
	return rv
}

func (nc _NumberClass) NumberWithLong(value int64) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithLong:"), value)
	return rv
}

func (nc _NumberClass) NumberWithLongLong(value int64) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithLongLong:"), value)
	return rv
}

func (nc _NumberClass) NumberWithShort(value int16) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithShort:"), value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedChar(value byte) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithUnsignedChar:"), value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedInt(value uint32) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithUnsignedInt:"), value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedInteger(value uint) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithUnsignedInteger:"), value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedLong(value uint64) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithUnsignedLong:"), value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedLongLong(value uint64) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithUnsignedLongLong:"), value)
	return rv
}

func (nc _NumberClass) NumberWithUnsignedShort(value uint16) Number {
	rv := objc.CallMethod[Number](nc, objc.SEL("numberWithUnsignedShort:"), value)
	return rv
}

func (n_ Number) InitWithBool(value bool) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithBool:"), value)
	return rv
}

func (n_ Number) InitWithChar(value byte) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithChar:"), value)
	return rv
}

func (n_ Number) InitWithDouble(value float64) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithDouble:"), value)
	return rv
}

func (n_ Number) InitWithFloat(value float32) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithFloat:"), value)
	return rv
}

func (n_ Number) InitWithInt(value int32) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithInt:"), value)
	return rv
}

func (n_ Number) InitWithInteger(value int) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithInteger:"), value)
	return rv
}

func (n_ Number) InitWithLong(value int64) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithLong:"), value)
	return rv
}

func (n_ Number) InitWithLongLong(value int64) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithLongLong:"), value)
	return rv
}

func (n_ Number) InitWithShort(value int16) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithShort:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedChar(value byte) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithUnsignedChar:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedInt(value uint32) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithUnsignedInt:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedInteger(value uint) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithUnsignedInteger:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedLong(value uint64) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithUnsignedLong:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedLongLong(value uint64) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithUnsignedLongLong:"), value)
	return rv
}

func (n_ Number) InitWithUnsignedShort(value uint16) Number {
	rv := objc.CallMethod[Number](n_, objc.SEL("initWithUnsignedShort:"), value)
	return rv
}

func (n_ Number) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.CallMethod[string](n_, objc.SEL("descriptionWithLocale:"), objc.ExtractPtr(locale))
	return rv
}

func (n_ Number) Compare(otherNumber INumber) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](n_, objc.SEL("compare:"), objc.ExtractPtr(otherNumber))
	return rv
}

func (n_ Number) IsEqualToNumber(number INumber) bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("isEqualToNumber:"), objc.ExtractPtr(number))
	return rv
}

func (n_ Number) BoolValue() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("boolValue"))
	return rv
}

func (n_ Number) CharValue() byte {
	rv := objc.CallMethod[byte](n_, objc.SEL("charValue"))
	return rv
}

func (n_ Number) DecimalValue() Decimal {
	rv := objc.CallMethod[Decimal](n_, objc.SEL("decimalValue"))
	return rv
}

func (n_ Number) DoubleValue() float64 {
	rv := objc.CallMethod[float64](n_, objc.SEL("doubleValue"))
	return rv
}

func (n_ Number) FloatValue() float32 {
	rv := objc.CallMethod[float32](n_, objc.SEL("floatValue"))
	return rv
}

func (n_ Number) IntValue() int32 {
	rv := objc.CallMethod[int32](n_, objc.SEL("intValue"))
	return rv
}

func (n_ Number) IntegerValue() int {
	rv := objc.CallMethod[int](n_, objc.SEL("integerValue"))
	return rv
}

func (n_ Number) LongLongValue() int64 {
	rv := objc.CallMethod[int64](n_, objc.SEL("longLongValue"))
	return rv
}

func (n_ Number) LongValue() int64 {
	rv := objc.CallMethod[int64](n_, objc.SEL("longValue"))
	return rv
}

func (n_ Number) ShortValue() int16 {
	rv := objc.CallMethod[int16](n_, objc.SEL("shortValue"))
	return rv
}

func (n_ Number) UnsignedCharValue() byte {
	rv := objc.CallMethod[byte](n_, objc.SEL("unsignedCharValue"))
	return rv
}

func (n_ Number) UnsignedIntegerValue() uint {
	rv := objc.CallMethod[uint](n_, objc.SEL("unsignedIntegerValue"))
	return rv
}

func (n_ Number) UnsignedIntValue() uint32 {
	rv := objc.CallMethod[uint32](n_, objc.SEL("unsignedIntValue"))
	return rv
}

func (n_ Number) UnsignedLongLongValue() uint64 {
	rv := objc.CallMethod[uint64](n_, objc.SEL("unsignedLongLongValue"))
	return rv
}

func (n_ Number) UnsignedLongValue() uint64 {
	rv := objc.CallMethod[uint64](n_, objc.SEL("unsignedLongValue"))
	return rv
}

func (n_ Number) UnsignedShortValue() uint16 {
	rv := objc.CallMethod[uint16](n_, objc.SEL("unsignedShortValue"))
	return rv
}

func (n_ Number) StringValue() string {
	rv := objc.CallMethod[string](n_, objc.SEL("stringValue"))
	return rv
}
