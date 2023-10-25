// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var AppleEventDescriptorClass = _AppleEventDescriptorClass{objc.GetClass("NSAppleEventDescriptor")}

type _AppleEventDescriptorClass struct {
	objc.Class
}

type IAppleEventDescriptor interface {
	objc.IObject
	DescriptorAtIndex(index int) AppleEventDescriptor
	InsertDescriptor_AtIndex(descriptor IAppleEventDescriptor, index int)
	RemoveDescriptorAtIndex(index int)
	SendEventWithOptions_Timeout_Error(sendOptions AppleEventSendOptions, timeoutInSeconds TimeInterval, error *Error) AppleEventDescriptor
	BooleanValue() bool
	Data() []byte
	Int32Value() int32
	NumberOfItems() int
	StringValue() string
	DateValue() Date
	DoubleValue() float64
	FileURLValue() URL
	IsRecordDescriptor() bool
}

type AppleEventDescriptor struct {
	objc.Object
}

func MakeAppleEventDescriptor(ptr unsafe.Pointer) AppleEventDescriptor {
	return AppleEventDescriptor{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ AppleEventDescriptor) InitListDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.SEL("initListDescriptor"))
	return rv
}

func (a_ AppleEventDescriptor) InitRecordDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.SEL("initRecordDescriptor"))
	return rv
}

func (ac _AppleEventDescriptorClass) Alloc() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("alloc"))
	return rv
}

func (ac _AppleEventDescriptorClass) New() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewAppleEventDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.New()
}

func (a_ AppleEventDescriptor) Init() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.SEL("init"))
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithBoolean(boolean bool) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("descriptorWithBoolean:"), boolean)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithInt32(signedInt int32) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("descriptorWithInt32:"), signedInt)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithString(string_ string) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("descriptorWithString:"), string_)
	return rv
}

func (ac _AppleEventDescriptorClass) ListDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("listDescriptor"))
	return rv
}

func (ac _AppleEventDescriptorClass) NullDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("nullDescriptor"))
	return rv
}

func (ac _AppleEventDescriptorClass) RecordDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("recordDescriptor"))
	return rv
}

func (a_ AppleEventDescriptor) DescriptorAtIndex(index int) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.SEL("descriptorAtIndex:"), index)
	return rv
}

func (a_ AppleEventDescriptor) InsertDescriptor_AtIndex(descriptor IAppleEventDescriptor, index int) {
	objc.CallMethod[objc.Void](a_, objc.SEL("insertDescriptor:atIndex:"), objc.ExtractPtr(descriptor), index)
}

func (a_ AppleEventDescriptor) RemoveDescriptorAtIndex(index int) {
	objc.CallMethod[objc.Void](a_, objc.SEL("removeDescriptorAtIndex:"), index)
}

func (ac _AppleEventDescriptorClass) DescriptorWithApplicationURL(applicationURL IURL) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("descriptorWithApplicationURL:"), objc.ExtractPtr(applicationURL))
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("descriptorWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithDate(date IDate) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("descriptorWithDate:"), objc.ExtractPtr(date))
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("descriptorWithDouble:"), doubleValue)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithFileURL(fileURL IURL) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("descriptorWithFileURL:"), objc.ExtractPtr(fileURL))
	return rv
}

func (a_ AppleEventDescriptor) SendEventWithOptions_Timeout_Error(sendOptions AppleEventSendOptions, timeoutInSeconds TimeInterval, error *Error) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.SEL("sendEventWithOptions:timeout:error:"), sendOptions, timeoutInSeconds, unsafe.Pointer(error))
	return rv
}

func (ac _AppleEventDescriptorClass) CurrentProcessDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.SEL("currentProcessDescriptor"))
	return rv
}

func (a_ AppleEventDescriptor) BooleanValue() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("booleanValue"))
	return rv
}

func (a_ AppleEventDescriptor) Data() []byte {
	rv := objc.CallMethod[[]byte](a_, objc.SEL("data"))
	return rv
}

func (a_ AppleEventDescriptor) Int32Value() int32 {
	rv := objc.CallMethod[int32](a_, objc.SEL("int32Value"))
	return rv
}

func (a_ AppleEventDescriptor) NumberOfItems() int {
	rv := objc.CallMethod[int](a_, objc.SEL("numberOfItems"))
	return rv
}

func (a_ AppleEventDescriptor) StringValue() string {
	rv := objc.CallMethod[string](a_, objc.SEL("stringValue"))
	return rv
}

func (a_ AppleEventDescriptor) DateValue() Date {
	rv := objc.CallMethod[Date](a_, objc.SEL("dateValue"))
	return rv
}

func (a_ AppleEventDescriptor) DoubleValue() float64 {
	rv := objc.CallMethod[float64](a_, objc.SEL("doubleValue"))
	return rv
}

func (a_ AppleEventDescriptor) FileURLValue() URL {
	rv := objc.CallMethod[URL](a_, objc.SEL("fileURLValue"))
	return rv
}

func (a_ AppleEventDescriptor) IsRecordDescriptor() bool {
	rv := objc.CallMethod[bool](a_, objc.SEL("isRecordDescriptor"))
	return rv
}
