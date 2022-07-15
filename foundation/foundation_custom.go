package foundation

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"

	"github.com/hsiafan/cocoa/objc"
)

type IString interface {
	objc.IObject
	String() string
}

type String struct {
	objc.Object
}

func MakeString(ptr unsafe.Pointer) String {
	return String{
		Object: objc.MakeObject(ptr),
	}
}

func NewString(str string) String {
	return MakeString(ffi.ToNSString(str))
}

func (N String) String() string {
	return ffi.ToGoString(N.Ptr())
}

func (N String) ToString() string {
	return ffi.ToGoString(N.Ptr())
}

type IMutableString interface {
	IString
	AppendString(str string)
}

type MutableString struct {
	String
}

func MakeMutableString(ptr unsafe.Pointer) MutableString {
	return MutableString{
		String: MakeString(ptr),
	}
}

var _MutableStringClass = objc.GetClass("NSMutableString")

func NewMutableString() MutableString {
	return ffi.CallMethod[MutableString](_MutableStringClass, "string")
}

func NewMutableStringWithString(str string) MutableString {
	return ffi.CallMethod[MutableString](_MutableStringClass, "stringWithString:", str)
}

func NewMutableStringWithCapacity(capacity uint) MutableString {
	return ffi.CallMethod[MutableString](_MutableStringClass, "stringWithCapacity:", capacity)
}

func (s MutableString) AppendString(str string) {
	ffi.CallMethod[ffi.Void](s, "appendString:", str)
}

type IData interface {
	objc.IObject
	ToBytes() []byte
}

type Data struct {
	objc.Object
}

func MakeData(ptr unsafe.Pointer) Data {
	return Data{
		Object: objc.MakeObject(ptr),
	}
}

func NewData(bytes []byte) Data {
	return MakeData(ffi.ToNSData(bytes))
}

func (d Data) ToBytes() []byte {
	return ffi.ToGoBytes(d.Ptr())
}

type IMutableData interface {
	IData
	AppendData(data []byte)
}

type MutableData struct {
	Data
}

func MakeMutableData(ptr unsafe.Pointer) MutableData {
	return MutableData{
		Data: MakeData(ptr),
	}
}

var _MutableDataClass = objc.GetClass("NSMutableData")

func NewMutableData() MutableData {
	return ffi.CallMethod[MutableData](_MutableDataClass, "data")
}

func NewMutableDataWithData(data []byte) MutableData {
	return ffi.CallMethod[MutableData](_MutableDataClass, "dataWithData:", data)
}

func NewMutableDataWithCapacity(capacity uint) MutableData {
	return ffi.CallMethod[MutableData](_MutableDataClass, "dataWithCapacity:", capacity)
}

func NewMutableDataWithLength(length uint) MutableData {
	return ffi.CallMethod[MutableData](_MutableDataClass, "dataWithLength:", length)
}

func (d MutableData) AppendData(data []byte) {
	ffi.CallMethod[ffi.Void](d, "appendData:", data)
}

type ISet interface {
	objc.IObject
	Count() uint
	AllObjects() Array
}

type Set struct {
	objc.Object
}

func MakeSet(ptr unsafe.Pointer) Set {
	return Set{
		Object: objc.MakeObject(ptr),
	}
}

var _SetClass = objc.GetClass("NSSet")

func NewSet() Set {
	return ffi.CallMethod[Set](_SetClass, "set")
}

func NewSetWithArray(array IArray) Set {
	return ffi.CallMethod[Set](_SetClass, "setWithArray:", array)
}

func (s Set) Count() uint {
	return ffi.CallMethod[uint](s, "count")
}

func (s Set) AllObjects() Array {
	return ffi.CallMethod[Array](s, "allObjects")
}

type IMutableSet interface {
	ISet
	AddObject(o objc.IObject)
}

type MutableSet struct {
	Set
}

func MakeMutableSet(ptr unsafe.Pointer) MutableSet {
	return MutableSet{
		Set: MakeSet(ptr),
	}
}

var _MutableSetClass = objc.GetClass("NSMutableSet")

func NewMutableSet() MutableSet {
	return ffi.CallMethod[MutableSet](_MutableSetClass, "set")
}

func NewMutableSetWithCapacity(size uint) MutableSet {
	return ffi.CallMethod[MutableSet](_MutableSetClass, "setWithCapacity:", size)
}

func NewMutableSetWithArray(array IArray) Set {
	return ffi.CallMethod[Set](_MutableSetClass, "setWithArray:", array)
}

func (s MutableSet) AddObject(o objc.IObject) {
	ffi.CallMethod[ffi.Void](s, "addObject:", o)
}

type IArray interface {
	objc.IObject
	Count() uint
}

type Array struct {
	objc.Object
}

func MakeArray(ptr unsafe.Pointer) Array {
	return Array{
		Object: objc.MakeObject(ptr),
	}
}

func (a Array) Count() uint {
	return ffi.CallMethod[uint](a, "count")
}

func ArrayOf[T any](slice []T) Array {
	return MakeArray(ffi.ToNSArray(reflect.ValueOf(slice)))
}

func ArrayToSlice[T any](a Array) []T {
	var zero []T
	return ffi.ToGoSlice(a.Ptr(), reflect.TypeOf(zero)).Interface().([]T)
}

type IMutableArray interface {
	IArray
	AddObject(o objc.IObject)
}

type MutableArray struct {
	Array
}

func MakeMutableArray(ptr unsafe.Pointer) MutableArray {
	return MutableArray{
		Array: MakeArray(ptr),
	}
}

var _MutableArrayClass = objc.GetClass("NSMutableArray")

func NewMutableArray() MutableArray {
	return ffi.CallMethod[MutableArray](_MutableArrayClass, "array")
}

func NewMutableArrayWithCapacity(size uint) MutableArray {
	return ffi.CallMethod[MutableArray](_MutableArrayClass, "arrayWithCapacity:", size)
}

func NewMutableSetWithSet(set ISet) MutableSet {
	return ffi.CallMethod[MutableSet](_MutableSetClass, "setWithSet:", set)
}

func (a MutableArray) AddObject(o objc.IObject) {
	ffi.CallMethod[ffi.Void](a, "addObject:", o)
}

type IDictionary interface {
	objc.IObject
	Count() uint
}

type Dictionary struct {
	objc.Object
}

func MakeDictionary(ptr unsafe.Pointer) Dictionary {
	return Dictionary{
		Object: objc.MakeObject(ptr),
	}
}

func (d Dictionary) Count() uint {
	return ffi.CallMethod[uint](d, "count")
}

func DictOf[K comparable, V any](m map[K]V) Dictionary {
	return MakeDictionary(ffi.ToNSDict(reflect.ValueOf(m)))
}

func DictToMap[K comparable, V any](d Dictionary) map[K]V {
	var zero map[K]V
	return ffi.ToGoMap(d.Ptr(), reflect.TypeOf(zero)).Interface().(map[K]V)
}

type IMutableDictionary interface {
	IDictionary
}

type MutableDictionary struct {
	Dictionary
}

func MakeMutableDictionary(ptr unsafe.Pointer) MutableDictionary {
	return MutableDictionary{
		Dictionary: MakeDictionary(ptr),
	}
}

var _MutableDictionaryClass = objc.GetClass("NSMutableDictionary")

func NewMutableDictionary() MutableDictionary {
	return ffi.CallMethod[MutableDictionary](_MutableDictionaryClass, "dictionary")
}

func NewMutableDictionaryWithDictionary(dict IDictionary) MutableDictionary {
	return ffi.CallMethod[MutableDictionary](_MutableDictionaryClass, "dictionaryWithDictionary:", dict)
}

func (d MutableDictionary) SetObject(k objc.IObject, v objc.IObject) {
	ffi.CallMethod[ffi.Void](d, "setObject:forKey:", v, k)
}

func ToNSError(code int, err error) Error {
	if err == nil {
		return Error{}
	}
	return ErrorClass.ErrorWithDomain_Code_UserInfo("cocoa.hsiafan.github.com", code, map[ErrorUserInfoKey]objc.IObject{
		"Error reason": NewString(err.Error()),
	})
}

// Error convert

var _ error = ocErr{}

type ocErr struct {
	code    int
	message string
}

// Error implements error
func (e ocErr) Error() string {
	return fmt.Sprintf("object error, code: %v, message: %v", e.code, e.message)
}

func ToGoError(err Error) error {
	if err.IsNil() {
		return nil
	}
	return ocErr{
		code:    err.Code(),
		message: err.Description(),
	}
}

// JSONSerialization begin
type JSONReadingOptions uint

const (
	JSONReadingMutableContainers         JSONReadingOptions = 1 << 0
	JSONReadingMutableLeaves             JSONReadingOptions = 1 << 1
	JSONReadingFragmentsAllowed          JSONReadingOptions = 1 << 2
	JSONReadingJSON5Allowed              JSONReadingOptions = 1 << 3
	JSONReadingTopLevelDictionaryAssumed JSONReadingOptions = 1 << 4
)

type JSONWritingOptions uint

const (
	JSONWritingPrettyPrinted          JSONWritingOptions = 1 << 0
	JSONWritingSortedKeys             JSONWritingOptions = 1 << 1
	JSONWritingFragmentsAllowed       JSONWritingOptions = 1 << 2
	JSONWritingWithoutEscapingSlashes JSONWritingOptions = 1 << 3
)

var _NSJSONSerializationClass = objc.GetClass("NSJSONSerialization")

func JSONObjectWithData(data []byte, options JSONReadingOptions) (objc.Object, error) {
	var err Error
	r := ffi.CallMethod[objc.Object](_NSJSONSerializationClass, "JSONObjectWithData:options:error:", data, options, &err)
	return r, ToGoError(err)
}

func DataWithJSONObject(o objc.IObject, options JSONWritingOptions) ([]byte, error) {
	var err Error
	r := ffi.CallMethod[[]byte](_NSJSONSerializationClass, "dataWithJSONObject:options:error:", o, options, &err)
	return r, ToGoError(err)
}

func IsValidJSONObject(o objc.IObject) bool {
	return ffi.CallMethod[bool](_NSJSONSerializationClass, "isValidJSONObject:", o)
}
