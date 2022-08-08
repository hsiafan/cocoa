// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	DescriptorForKeyword(keyword uint32) AppleEventDescriptor
	KeywordForDescriptorAtIndex(index int) uint32
	RemoveDescriptorWithKeyword(keyword uint32)
	SetDescriptor_ForKeyword(descriptor IAppleEventDescriptor, keyword uint32)
	AttributeDescriptorForKeyword(keyword uint32) AppleEventDescriptor
	ParamDescriptorForKeyword(keyword uint32) AppleEventDescriptor
	RemoveParamDescriptorWithKeyword(keyword uint32)
	SetAttributeDescriptor_ForKeyword(descriptor IAppleEventDescriptor, keyword uint32)
	SetParamDescriptor_ForKeyword(descriptor IAppleEventDescriptor, keyword uint32)
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
	rv := ffi.CallMethod[AppleEventDescriptor](a_, "initListDescriptor")
	return rv
}

func (a_ AppleEventDescriptor) InitRecordDescriptor() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](a_, "initRecordDescriptor")
	return rv
}

func (ac _AppleEventDescriptorClass) Alloc() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "alloc")
	return rv
}

func (a_ AppleEventDescriptor) Init() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](a_, "init")
	return rv
}

func (ac _AppleEventDescriptorClass) New() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "new")
	rv.Autorelease()
	return rv
}

func NewAppleEventDescriptor() AppleEventDescriptor {
	return AppleEventDescriptorClass.New()
}

func (ac _AppleEventDescriptorClass) DescriptorWithBoolean(boolean bool) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "descriptorWithBoolean:", boolean)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithInt32(signedInt int32) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "descriptorWithInt32:", signedInt)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithString(_string string) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "descriptorWithString:", _string)
	return rv
}

func (ac _AppleEventDescriptorClass) ListDescriptor() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "listDescriptor")
	return rv
}

func (ac _AppleEventDescriptorClass) NullDescriptor() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "nullDescriptor")
	return rv
}

func (ac _AppleEventDescriptorClass) RecordDescriptor() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "recordDescriptor")
	return rv
}

func (a_ AppleEventDescriptor) DescriptorAtIndex(index int) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](a_, "descriptorAtIndex:", index)
	return rv
}

func (a_ AppleEventDescriptor) InsertDescriptor_AtIndex(descriptor IAppleEventDescriptor, index int) {
	ffi.CallMethod[ffi.Void](a_, "insertDescriptor:atIndex:", descriptor, index)
}

func (a_ AppleEventDescriptor) RemoveDescriptorAtIndex(index int) {
	ffi.CallMethod[ffi.Void](a_, "removeDescriptorAtIndex:", index)
}

func (a_ AppleEventDescriptor) DescriptorForKeyword(keyword uint32) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](a_, "descriptorForKeyword:", keyword)
	return rv
}

func (a_ AppleEventDescriptor) KeywordForDescriptorAtIndex(index int) uint32 {
	rv := ffi.CallMethod[uint32](a_, "keywordForDescriptorAtIndex:", index)
	return rv
}

func (a_ AppleEventDescriptor) RemoveDescriptorWithKeyword(keyword uint32) {
	ffi.CallMethod[ffi.Void](a_, "removeDescriptorWithKeyword:", keyword)
}

func (a_ AppleEventDescriptor) SetDescriptor_ForKeyword(descriptor IAppleEventDescriptor, keyword uint32) {
	ffi.CallMethod[ffi.Void](a_, "setDescriptor:forKeyword:", descriptor, keyword)
}

func (a_ AppleEventDescriptor) AttributeDescriptorForKeyword(keyword uint32) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](a_, "attributeDescriptorForKeyword:", keyword)
	return rv
}

func (a_ AppleEventDescriptor) ParamDescriptorForKeyword(keyword uint32) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](a_, "paramDescriptorForKeyword:", keyword)
	return rv
}

func (a_ AppleEventDescriptor) RemoveParamDescriptorWithKeyword(keyword uint32) {
	ffi.CallMethod[ffi.Void](a_, "removeParamDescriptorWithKeyword:", keyword)
}

func (a_ AppleEventDescriptor) SetAttributeDescriptor_ForKeyword(descriptor IAppleEventDescriptor, keyword uint32) {
	ffi.CallMethod[ffi.Void](a_, "setAttributeDescriptor:forKeyword:", descriptor, keyword)
}

func (a_ AppleEventDescriptor) SetParamDescriptor_ForKeyword(descriptor IAppleEventDescriptor, keyword uint32) {
	ffi.CallMethod[ffi.Void](a_, "setParamDescriptor:forKeyword:", descriptor, keyword)
}

func (ac _AppleEventDescriptorClass) DescriptorWithApplicationURL(applicationURL IURL) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "descriptorWithApplicationURL:", applicationURL)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "descriptorWithBundleIdentifier:", bundleIdentifier)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithDate(date IDate) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "descriptorWithDate:", date)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "descriptorWithDouble:", doubleValue)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithFileURL(fileURL IURL) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "descriptorWithFileURL:", fileURL)
	return rv
}

func (ac _AppleEventDescriptorClass) DescriptorWithProcessIdentifier(processIdentifier int32) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "descriptorWithProcessIdentifier:", processIdentifier)
	return rv
}

func (a_ AppleEventDescriptor) SendEventWithOptions_Timeout_Error(sendOptions AppleEventSendOptions, timeoutInSeconds TimeInterval, error *Error) AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](a_, "sendEventWithOptions:timeout:error:", sendOptions, timeoutInSeconds, unsafe.Pointer(error))
	return rv
}

func (ac _AppleEventDescriptorClass) CurrentProcessDescriptor() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](ac, "currentProcessDescriptor")
	return rv
}

func (a_ AppleEventDescriptor) BooleanValue() bool {
	rv := ffi.CallMethod[bool](a_, "booleanValue")
	return rv
}

func (a_ AppleEventDescriptor) Data() []byte {
	rv := ffi.CallMethod[[]byte](a_, "data")
	return rv
}

func (a_ AppleEventDescriptor) Int32Value() int32 {
	rv := ffi.CallMethod[int32](a_, "int32Value")
	return rv
}

func (a_ AppleEventDescriptor) NumberOfItems() int {
	rv := ffi.CallMethod[int](a_, "numberOfItems")
	return rv
}

func (a_ AppleEventDescriptor) StringValue() string {
	rv := ffi.CallMethod[string](a_, "stringValue")
	return rv
}

func (a_ AppleEventDescriptor) DateValue() Date {
	rv := ffi.CallMethod[Date](a_, "dateValue")
	return rv
}

func (a_ AppleEventDescriptor) DoubleValue() float64 {
	rv := ffi.CallMethod[float64](a_, "doubleValue")
	return rv
}

func (a_ AppleEventDescriptor) FileURLValue() URL {
	rv := ffi.CallMethod[URL](a_, "fileURLValue")
	return rv
}

func (a_ AppleEventDescriptor) IsRecordDescriptor() bool {
	rv := ffi.CallMethod[bool](a_, "isRecordDescriptor")
	return rv
}
