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
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("initListDescriptor"))
	return rv
}

func (a_ AppleEventDescriptor) InitRecordDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("initRecordDescriptor"))
	return rv
}

func (ac _AppleEventDescriptorClass) Alloc() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("alloc"))
	return rv
}

func (ac _AppleEventDescriptorClass) New() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewAppleEventDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.New()
}

func (a_ AppleEventDescriptor) Init() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("init"))
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithBoolean(boolean bool) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithBoolean:"), boolean)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithInt32(signedInt int32) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithInt32:"), signedInt)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithString(string_ string) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithString:"), string_)
	return rv
}

func (ac _AppleEventDescriptorClass) ListDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("listDescriptor"))
	return rv
}

func (ac _AppleEventDescriptorClass) NullDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("nullDescriptor"))
	return rv
}

func (ac _AppleEventDescriptorClass) RecordDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("recordDescriptor"))
	return rv
}

func (a_ AppleEventDescriptor) DescriptorAtIndex(index int) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("descriptorAtIndex:"), index)
	return rv
}

func (a_ AppleEventDescriptor) InsertDescriptor_AtIndex(descriptor IAppleEventDescriptor, index int) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("insertDescriptor:atIndex:"), descriptor, index)
}

func (a_ AppleEventDescriptor) RemoveDescriptorAtIndex(index int) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("removeDescriptorAtIndex:"), index)
}

func (ac _AppleEventDescriptorClass) DescriptorWithApplicationURL(applicationURL IURL) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithApplicationURL:"), applicationURL)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithDate(date IDate) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithDate:"), date)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithDouble:"), doubleValue)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithFileURL(fileURL IURL) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("descriptorWithFileURL:"), fileURL)
	return rv
}

func (a_ AppleEventDescriptor) SendEventWithOptions_Timeout_Error(sendOptions AppleEventSendOptions, timeoutInSeconds TimeInterval, error *Error) AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](a_, objc.GetSelector("sendEventWithOptions:timeout:error:"), sendOptions, timeoutInSeconds, unsafe.Pointer(error))
	return rv
}

func (ac _AppleEventDescriptorClass) CurrentProcessDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](ac, objc.GetSelector("currentProcessDescriptor"))
	return rv
}

func (a_ AppleEventDescriptor) BooleanValue() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("booleanValue"))
	return rv
}

func (a_ AppleEventDescriptor) Data() []byte {
	rv := objc.CallMethod[[]byte](a_, objc.GetSelector("data"))
	return rv
}

func (a_ AppleEventDescriptor) Int32Value() int32 {
	rv := objc.CallMethod[int32](a_, objc.GetSelector("int32Value"))
	return rv
}

func (a_ AppleEventDescriptor) NumberOfItems() int {
	rv := objc.CallMethod[int](a_, objc.GetSelector("numberOfItems"))
	return rv
}

func (a_ AppleEventDescriptor) StringValue() string {
	rv := objc.CallMethod[string](a_, objc.GetSelector("stringValue"))
	return rv
}

func (a_ AppleEventDescriptor) DateValue() Date {
	rv := objc.CallMethod[Date](a_, objc.GetSelector("dateValue"))
	return rv
}

func (a_ AppleEventDescriptor) DoubleValue() float64 {
	rv := objc.CallMethod[float64](a_, objc.GetSelector("doubleValue"))
	return rv
}

func (a_ AppleEventDescriptor) FileURLValue() URL {
	rv := objc.CallMethod[URL](a_, objc.GetSelector("fileURLValue"))
	return rv
}

func (a_ AppleEventDescriptor) IsRecordDescriptor() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isRecordDescriptor"))
	return rv
}
