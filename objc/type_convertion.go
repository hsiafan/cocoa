package objc

// #import <stdint.h>
// #include "type_convertion.h"
// void* to_ns_string(const char* str);
// const char* to_c_string(void* ptr);
//
// void* to_ns_data(data);
// data to_c_bytes(void* ptr);
//
// void* to_ns_array(array array);
// array to_c_array(void* ptr);
//
// dict to_c_items(void* ptr);
// void* to_ns_dict(dict cDict);
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/internal"
)

var intType = reflect.TypeOf(0)
var uintType = reflect.TypeOf(uint(0))
var int32Type = reflect.TypeOf(int32(0))
var uint32Type = reflect.TypeOf(uint32(0))
var int64Type = reflect.TypeOf(int64(0))
var uint64Type = reflect.TypeOf(uint64(0))
var float32Type = reflect.TypeOf(float32(0))
var float64Type = reflect.TypeOf(float64(0))
var bytesType = reflect.TypeOf([]byte{})
var unsafePointerType = reflect.TypeOf(unsafe.Pointer(nil))
var pointerHolderType = reflect.TypeOf((*Holder)(nil)).Elem()
var selectorType = reflect.TypeOf(Selector{})
var classType = reflect.TypeOf(Class{})

type holder struct {
	ptr unsafe.Pointer
}

type Value struct {
	// typ holds the type of the value represented by a Value.
	typ unsafe.Pointer

	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.
	ptr unsafe.Pointer

	flag uintptr
}

var flagIndir uintptr = 1 << 7

var magicStringPtr uintptr = uintptr(unsafe.Pointer(&flagIndir))

// MagicNilString is a magic string value corresponding to objc nil NSString
var MagicNilString = "$$__magic_nil_string__$$"

func ToNSString(s string) unsafe.Pointer {
	if s == MagicNilString {
		return nil
	}
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return C.to_ns_string(cs)
}

func ToGoString(p unsafe.Pointer) string {
	if p == nil {
		return ""
	}
	return C.GoString(C.to_c_string(p))
}

func ToNSData(bytes []byte) unsafe.Pointer {
	if bytes == nil {
		return nil
	}
	size := len(bytes)
	var p unsafe.Pointer
	var d C.data
	if size == 0 {
		p = C.to_ns_data(d)
	} else {
		p = C.to_ns_data(C.data{
			len:  C.ulong(size),
			data: unsafe.Pointer(&bytes[0]),
		})
	}
	return p
}

func ToGoBytes(p unsafe.Pointer) []byte {
	if p == nil {
		return nil
	}
	d := C.to_c_bytes(p)
	size := int(d.len)
	if size <= 0 {
		return nil
	}
	bytes := unsafe.Slice((*byte)(d.data), size)
	newBytes := make([]byte, size)
	copy(newBytes, bytes)
	return newBytes
}

func objectToGoHolder(ptr unsafe.Pointer, paramType reflect.Type) reflect.Value {
	v := reflect.Zero(paramType)
	vp := (*Value)(unsafe.Pointer(&v))
	h := holder{
		ptr: ptr,
	}
	if vp.flag&flagIndir != 0 {
		*(*holder)(vp.ptr) = h
	} else {
		vp.ptr = internal.ForceCast[holder, unsafe.Pointer](h)
	}
	return v
}

func toNSElement(v reflect.Value) unsafe.Pointer {
	var t = v.Type()
	switch t.Kind() {
	case reflect.String:
		return ToNSString(v.String())
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return ToNSData(v.Convert(bytesType).Interface().([]byte))
		} else {
			return ToNSArray(v)
		}
	case reflect.Interface:
		if t.ConvertibleTo(pointerHolderType) {
			return v.Interface().(Holder).Ptr()
		}
	default:

	}
	panic("not supported types: " + t.Name())
}

func toGoElement(ptr unsafe.Pointer, t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.String:
		return reflect.ValueOf(ToGoString(ptr)).Convert(t)
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return reflect.ValueOf(ToGoBytes(ptr)).Convert(t)
		} else {
			return ToGoSlice(ptr, t)
		}
	case reflect.Struct:
		if t.ConvertibleTo(pointerHolderType) {
			return objectToGoHolder(ptr, t)
		}
	default:

	}
	panic("not supported types: " + t.Name())
}

func ToNSArray(slice reflect.Value) unsafe.Pointer {
	if slice.IsNil() {
		return nil
	}
	var cArray C.array
	if slice.Len() > 0 {
		cArrayData := make([]unsafe.Pointer, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			cArrayData[i] = toNSElement(slice.Index(i))
		}
		cArray.data = unsafe.Pointer(&cArrayData[0])
		cArray.len = C.ulong(len(cArrayData))
	}
	return C.to_ns_array(cArray)
}

func ToGoSlice(ptr unsafe.Pointer, sliceType reflect.Type) reflect.Value {
	if ptr == nil {
		return reflect.New(sliceType).Elem()
	}
	ca := C.to_c_array(ptr)
	if ca.len > 0 {
		defer C.free(ca.data)
	}
	ptrSlice := unsafe.Slice((*unsafe.Pointer)(ca.data), int(ca.len))
	var slice = reflect.MakeSlice(sliceType, len(ptrSlice), len(ptrSlice))
	for idx, ptr := range ptrSlice {
		slice.Index(idx).Set(toGoElement(ptr, sliceType.Elem()))
	}
	return slice
}

func ToNSDict(m reflect.Value) unsafe.Pointer {
	if m.IsNil() {
		return nil
	}
	var cDict C.dict
	if m.Len() > 0 {
		keyData := make([]unsafe.Pointer, m.Len())
		valueData := make([]unsafe.Pointer, m.Len())
		for idx, k := range m.MapKeys() {
			v := m.MapIndex(k)
			keyData[idx] = toNSElement(k)
			valueData[idx] = toNSElement(v)
		}
		cDict.key_data = unsafe.Pointer(&keyData[0])
		cDict.value_data = unsafe.Pointer(&valueData[0])
		cDict.len = C.ulong(m.Len())
	}
	return C.to_ns_dict(cDict)
}

func ToGoMap(ptr unsafe.Pointer, mapType reflect.Type) reflect.Value {
	if ptr == nil {
		return reflect.New(mapType).Elem()
	}
	cDict := C.to_c_items(ptr)
	if cDict.len > 0 {
		defer C.free(cDict.key_data)
		defer C.free(cDict.value_data)
	}
	keys := unsafe.Slice((*unsafe.Pointer)(cDict.key_data), int(cDict.len))
	values := unsafe.Slice((*unsafe.Pointer)(cDict.value_data), int(cDict.len))
	var m = reflect.MakeMap(mapType)
	for idx, k := range keys {
		v := values[idx]
		m.SetMapIndex(toGoElement(k, mapType.Key()), toGoElement(v, mapType.Elem()))
	}
	return m
}

// parseGoType return a pointer to receive objc value, for given go type
func parseGoType(t reflect.Type) unsafe.Pointer {
	switch t.Kind() {
	case reflect.Bool, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		var cv C.int32_t
		return unsafe.Pointer(&cv)
	case reflect.Int, reflect.Uint, reflect.Int64, reflect.Uint64:
		var cv C.int64_t
		return unsafe.Pointer(&cv)
	case reflect.Float32:
		var cv C.float
		return unsafe.Pointer(&cv)
	case reflect.Float64:
		var cv C.double
		return unsafe.Pointer(&cv)
	case reflect.UnsafePointer, reflect.Pointer:
		var cv unsafe.Pointer
		return unsafe.Pointer(&cv)
	case reflect.String:
		var cv unsafe.Pointer
		return unsafe.Pointer(&cv)
	case reflect.Slice, reflect.Map:
		var cv unsafe.Pointer
		return unsafe.Pointer(&cv)
	case reflect.Struct:
		if t.Implements(pointerHolderType) {
			var cv unsafe.Pointer
			return unsafe.Pointer(&cv)
		} else {
			rv := reflect.New(t).Elem()
			return getStructPointer(rv)
		}
	case reflect.Func:
		// a pointer to Block
		var cv unsafe.Pointer
		return unsafe.Pointer(&cv)
	default:
		panic(fmt.Sprintf("not support type: %v", t.Name()))
	}
}

// convertToGoValue convert objc value to go value.
// param p: the pointer to objc value
// param t: the go value type
func convertToGoValue(p unsafe.Pointer, t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(int(*((*C.uint8_t)(p))) == 1).Convert(t)
	case reflect.Int8:
		return reflect.ValueOf(int8(*((*C.int8_t)(p)))).Convert(t)
	case reflect.Int16:
		return reflect.ValueOf(int16(*((*C.int16_t)(p)))).Convert(t)
	case reflect.Int32:
		return reflect.ValueOf(int32(*((*C.int32_t)(p)))).Convert(t)
	case reflect.Int, reflect.Int64:
		return reflect.ValueOf(int64(*((*C.int64_t)(p)))).Convert(t)
	case reflect.Uint8:
		return reflect.ValueOf(uint8(*((*C.uint8_t)(p)))).Convert(t)
	case reflect.Uint16:
		return reflect.ValueOf(uint16(*((*C.uint16_t)(p)))).Convert(t)
	case reflect.Uint32:
		return reflect.ValueOf(uint32(*((*C.uint32_t)(p)))).Convert(t)
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		return reflect.ValueOf(uint64(*((*C.uint64_t)(p)))).Convert(t)
	case reflect.Float32:
		return reflect.ValueOf(float32(*((*C.float)(p)))).Convert(t)
	case reflect.Float64:
		return reflect.ValueOf(float64(*((*C.double)(p)))).Convert(t)
	case reflect.Pointer, reflect.UnsafePointer:
		return reflect.ValueOf(*(*unsafe.Pointer)(p)).Convert(t)
	case reflect.String:
		return reflect.ValueOf(ToGoString(*(*unsafe.Pointer)(p))).Convert(t)
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return reflect.ValueOf(ToGoBytes(*(*unsafe.Pointer)(p)))
		} else {
			return reflect.ValueOf(ToGoSlice(*(*unsafe.Pointer)(p), t))
		}
	case reflect.Map:
		return reflect.ValueOf(ToGoMap(*(*unsafe.Pointer)(p), t))
	case reflect.Struct:
		if t.Implements(pointerHolderType) {
			return objectToGoHolder(*(*unsafe.Pointer)(p), t)
		} else {
			rv := reflect.New(t).Elem()
			storeStructValue(&rv, p)
			return rv
		}
	case reflect.Func:
		rv := wrapBlockInGoFunc(*(*unsafe.Pointer)(p), t)
		return rv
	default:
		panic(fmt.Sprintf("not support type: %v", t))
	}
}

// convertToObjcValue convert go value to objc value, return the pointer to objc value, and the ffi-type
func convertToObjcValue(v reflect.Value) unsafe.Pointer {
	if !v.IsValid() {
		var p unsafe.Pointer = nil
		return unsafe.Pointer(&p)
	}

	if v.Type().AssignableTo(pointerHolderType) {
		cv := v.Interface().(Holder).Ptr()
		return unsafe.Pointer(&cv)
	}

	rt := v.Type()
	switch rt.Kind() {
	case reflect.Bool:
		var cv C.uint8_t = 0
		if v.Bool() {
			cv = 1
		}
		return unsafe.Pointer(&cv)
	case reflect.Int8:
		cv := C.int8_t(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Int16:
		cv := C.int16_t(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Int32:
		cv := C.int32_t(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Int, reflect.Int64:
		cv := C.int64_t(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Uint8:
		cv := C.uint8_t(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Uint16:
		cv := C.uint16_t(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Uint32:
		cv := C.uint32_t(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		cv := C.uint64_t(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Float32:
		cv := C.float(v.Float())
		return unsafe.Pointer(&cv)
	case reflect.Float64:
		cv := C.double(v.Float())
		return unsafe.Pointer(&cv)
	case reflect.UnsafePointer, reflect.Pointer:
		cv := v.UnsafePointer()
		return unsafe.Pointer(&cv)
	case reflect.Interface:
		panic(fmt.Sprintf("not support type: %T", v))
	case reflect.Struct:
		return getStructPointer(v)
	case reflect.String:
		sp := ToNSString(v.String())
		return unsafe.Pointer(&sp)
	case reflect.Slice:
		if rt.Elem().Kind() == reflect.Uint8 {
			dp := ToNSData(v.Bytes())
			return unsafe.Pointer(&dp)
		} else {
			sp := ToNSArray(v)
			return unsafe.Pointer(&sp)
		}
	case reflect.Map:
		dp := ToNSDict(v)
		return unsafe.Pointer(&dp)
	case reflect.Func:
		fp := CreateMallocBlock(v.Interface())
		return unsafe.Pointer(&fp.ptr)
	default:
		panic(fmt.Sprintf("not support type: %T, kind: %v", v, v.Kind()))
	}
}

func setGoValueToObjcPointer(rv reflect.Value, p unsafe.Pointer) {
	if rv.Type().AssignableTo(pointerHolderType) {
		*(*unsafe.Pointer)(p) = rv.Interface().(Holder).Ptr()
	}

	rt := rv.Type()
	switch rt.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			*(*uint8)(p) = 1
		} else {
			*(*uint8)(p) = 0
		}
	case reflect.Int8:
		*(*int8)(p) = int8(rv.Int())
	case reflect.Int16:
		*(*int16)(p) = int16(rv.Int())
	case reflect.Int32:
		*(*int32)(p) = int32(rv.Int())
	case reflect.Int, reflect.Int64:
		*(*int64)(p) = int64(rv.Int())
	case reflect.Uint8:
		*(*uint8)(p) = uint8(rv.Uint())
	case reflect.Uint16:
		*(*uint16)(p) = uint16(rv.Uint())
	case reflect.Uint32:
		*(*uint32)(p) = uint32(rv.Uint())
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		*(*uint64)(p) = uint64(rv.Uint())
	case reflect.Float32:
		*(*float32)(p) = float32(rv.Float())
	case reflect.Float64:
		*(*float64)(p) = float64(rv.Float())
	case reflect.UnsafePointer, reflect.Pointer:
		*(*unsafe.Pointer)(p) = rv.UnsafePointer()
	case reflect.Interface:
		panic(fmt.Sprintf("not support type: %v", rv.Type()))
	case reflect.Struct:
		sp := getStructPointer(rv)
		size := rv.Type().Size()
		copy(unsafe.Slice((*byte)(p), size), unsafe.Slice((*byte)(sp), size))
	case reflect.String:
		sp := ToNSString(rv.String())
		*(*unsafe.Pointer)(p) = sp
	case reflect.Slice:
		if rt.Elem().Kind() == reflect.Uint8 {
			dp := ToNSData(rv.Bytes())
			*(*unsafe.Pointer)(p) = dp
		} else {
			sp := ToNSArray(rv)
			*(*unsafe.Pointer)(p) = sp
		}
	case reflect.Map:
		dp := ToNSDict(rv)
		*(*unsafe.Pointer)(p) = dp
	case reflect.Func:
		fp := CreateMallocBlock(rv.Interface())
		*(*unsafe.Pointer)(p) = fp.ptr
	default:
		panic(fmt.Sprintf("not support type: %v, kind: %v", rv.Type(), rv.Kind()))
	}
}

func getStructPointer(value reflect.Value) unsafe.Pointer {
	vsp := (*Value)(unsafe.Pointer(&value))

	if vsp.flag&flagIndir == 0 {
		return unsafe.Pointer(&vsp.ptr)
	} else {
		return vsp.ptr
	}
}

func storeStructValue(value *reflect.Value, ptr unsafe.Pointer) {
	vsp := (*Value)(unsafe.Pointer(value))

	if vsp.flag&flagIndir == 0 {
		vsp.ptr = *(*unsafe.Pointer)(ptr)
	} else {
		vsp.ptr = ptr
	}
}

// getFFIType return the ffi type
func getFFIType(v any) *ffi.Type {
	if v == nil {
		return ffi.TypePointer
	}

	if _, ok := v.(Holder); ok {
		return ffi.TypePointer
	}

	rv := reflect.ValueOf(v)
	rt := rv.Type()
	return toFFIType(rt)
}

func toFFIType(rt reflect.Type) *ffi.Type {
	if rt.AssignableTo(pointerHolderType) {
		return ffi.TypePointer
	}
	switch rt.Kind() {
	case reflect.Bool:
		return ffi.TypeUint8
	case reflect.Int8:
		return ffi.TypeSint8
	case reflect.Int16:
		return ffi.TypeSint16
	case reflect.Int32:
		return ffi.TypeSint32
	case reflect.Int, reflect.Int64:
		return ffi.TypeSint64
	case reflect.Uint8:
		return ffi.TypeUint8
	case reflect.Uint16:
		return ffi.TypeUint16
	case reflect.Uint32:
		return ffi.TypeUint32
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		return ffi.TypeUint64
	case reflect.Float32:
		return ffi.TypeFloat
	case reflect.Float64:
		return ffi.TypeDouble
	case reflect.UnsafePointer, reflect.Pointer:
		return ffi.TypePointer
	case reflect.Interface:
		panic(fmt.Sprintf("not support type: %v", rt))
	case reflect.Struct:
		if rt.Size() == 0 {
			return ffi.TypeVoid
		}
		return getStructFFIType(rt)
	case reflect.String:
		return ffi.TypePointer
	case reflect.Slice:
		return ffi.TypePointer
	case reflect.Map:
		return ffi.TypePointer
	case reflect.Func:
		return ffi.TypePointer
	default:
		panic(fmt.Sprintf("not support type: %v, kind: %v", rt, rt.Kind()))
	}
}

func getStructFFIType(t reflect.Type) *ffi.Type {
	fn := t.NumField()
	var fts []*ffi.Type
	for i := 0; i < fn; i++ {
		ft := t.Field(i).Type
		switch ft.Kind() {
		case reflect.Array:
			for i := 0; i < ft.Len(); i++ {
				fts = append(fts, toFFIType(ft.Elem()))
			}
		default:
			fts = append(fts, toFFIType(ft))
		}
	}
	return ffi.StructType(fts)
}
