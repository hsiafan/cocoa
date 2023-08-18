package objc

// #import <stdint.h>
// #import <stdlib.h>
// #import <stdbool.h>
// void* to_ns_string(void* str, unsigned long len);
// const char* to_c_string(void* ptr, bool* shouldFree);
//
// void* to_ns_data(void* data, unsigned long len);
// void* to_c_bytes(void* ptr, unsigned long *len);
//
// void* to_ns_array(void** items, unsigned long len);
// unsigned long ns_array_len(void* ptr);
// void ns_array_get(void* ptr, void** item, unsigned long len);
//
// void* to_ns_dict(void** keys, void** values, unsigned long size);
// unsigned long ns_dict_size(void* ptr);
// void ns_dict_get(void* ptr, void** keys, void** values);
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
)

var bytesType = reflect.TypeOf([]byte{})
var pointerHolderType = reflect.TypeOf((*Holder)(nil)).Elem()
var selectorType = reflect.TypeOf(Selector{})
var classType = reflect.TypeOf(Class{})

type reValue struct {
	// typ holds the type of the value represented by a Value.
	typ unsafe.Pointer

	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.
	ptr unsafe.Pointer

	flag uintptr
}

// flagIndir: val holds a pointer to the data
var flagIndir uintptr = 1 << 7

func autoRelease(p unsafe.Pointer) unsafe.Pointer {
	MakeObject(p).Autorelease()
	return p
}

func ToNSString(s string) unsafe.Pointer {
	sp := unsafe.StringData(s)
	p := C.to_ns_string(unsafe.Pointer(sp), C.ulong(len(s)))
	return autoRelease(p)
}

func ToGoString(p unsafe.Pointer) string {
	if p == nil {
		return ""
	}
	var shouldFree C.bool
	cstr := C.to_c_string(p, &shouldFree)
	str := C.GoString(cstr)
	if shouldFree {
		C.free(unsafe.Pointer(cstr))
	}
	return str
}

func ToNSData(bytes []byte) unsafe.Pointer {
	if bytes == nil {
		return nil
	}
	size := len(bytes)
	var p unsafe.Pointer
	if size == 0 {
		p = C.to_ns_data(nil, C.ulong(0))
	} else {
		p = C.to_ns_data(unsafe.Pointer(&bytes[0]), C.ulong(size))
	}
	return autoRelease(p)
}

func ToGoBytes(p unsafe.Pointer) []byte {
	if p == nil {
		return nil
	}
	var len C.ulong
	data := C.to_c_bytes(p, &len)
	size := int(len)
	if size <= 0 {
		return nil
	}
	bytes := unsafe.Slice((*byte)(data), size)
	newBytes := make([]byte, size)
	copy(newBytes, bytes)
	return newBytes
}

// slice/dict elements convert
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
		if t.AssignableTo(pointerHolderType) {
			return v.Interface().(Holder).Ptr()
		}
	case reflect.Struct:
		if t.Implements(pointerHolderType) {
			return v.Interface().(Holder).Ptr()
		}
	default:

	}
	panic("not supported types: " + t.Name())
}

// slice/dict elements convert
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
		if t.Implements(pointerHolderType) {
			// objc object pointer holder struct should have same layout as a pointer
			return reflect.NewAt(t, unsafe.Pointer(&ptr)).Elem()
		}
	default:

	}
	panic("not supported types: " + t.Name())
}

func ToNSArray(slice reflect.Value) unsafe.Pointer {
	if slice.IsNil() {
		return nil
	}
	if slice.Len() == 0 {
		return autoRelease(C.to_ns_array(nil, C.ulong(0)))
	}
	cArrayData := make([]unsafe.Pointer, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		cArrayData[i] = toNSElement(slice.Index(i))
	}
	return autoRelease(C.to_ns_array((*unsafe.Pointer)(&cArrayData[0]), C.ulong(len(cArrayData))))
}

func ToGoSlice(ptr unsafe.Pointer, sliceType reflect.Type) reflect.Value {
	if ptr == nil {
		return reflect.New(sliceType).Elem()
	}
	length := int(C.ns_array_len(ptr))
	if length == 0 {
		return reflect.MakeSlice(sliceType, 0, 0)
	}
	ptrSlice := make([]unsafe.Pointer, length)
	C.ns_array_get(ptr, (*unsafe.Pointer)(&ptrSlice[0]), C.ulong(length))
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
	if m.Len() == 0 {
		return C.to_ns_dict(nil, nil, 0)
	}
	keys := make([]unsafe.Pointer, m.Len())
	values := make([]unsafe.Pointer, m.Len())
	for idx, k := range m.MapKeys() {
		v := m.MapIndex(k)
		keys[idx] = toNSElement(k)
		values[idx] = toNSElement(v)
	}
	rp := C.to_ns_dict((*unsafe.Pointer)(&keys[0]), (*unsafe.Pointer)(&values[0]), C.ulong(m.Len()))
	return autoRelease(rp)
}

func ToGoMap(ptr unsafe.Pointer, mapType reflect.Type) reflect.Value {
	if ptr == nil {
		return reflect.New(mapType).Elem()
	}
	size := int(C.ns_dict_size(ptr))
	if size == 0 {
		return reflect.MakeMap(mapType)
	}
	keys := make([]unsafe.Pointer, size)
	values := make([]unsafe.Pointer, size)
	C.ns_dict_get(ptr, (*unsafe.Pointer)(&keys[0]), (*unsafe.Pointer)(&values[0]))
	var m = reflect.MakeMap(mapType)
	for idx, k := range keys {
		v := values[idx]
		m.SetMapIndex(toGoElement(k, mapType.Key()), toGoElement(v, mapType.Elem()))
	}
	return m
}

// convertToGoValue convert objc value to go value.
// param p: the pointer to objc value
// param t: the go value type
func convertToGoValue(p unsafe.Pointer, t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(*(*uint8)(p) == 1).Convert(t)
	case reflect.Int8:
		return reflect.ValueOf(*(*int8)(p)).Convert(t)
	case reflect.Int16:
		return reflect.ValueOf(*(*int16)(p)).Convert(t)
	case reflect.Int32:
		return reflect.ValueOf(*(*int32)(p)).Convert(t)
	case reflect.Int, reflect.Int64:
		return reflect.ValueOf(*(*int64)(p)).Convert(t)
	case reflect.Uint8:
		return reflect.ValueOf(*(*uint8)(p)).Convert(t)
	case reflect.Uint16:
		return reflect.ValueOf(*(*uint16)(p)).Convert(t)
	case reflect.Uint32:
		return reflect.ValueOf(*(*uint32)(p)).Convert(t)
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		return reflect.ValueOf(*(*uint64)(p)).Convert(t)
	case reflect.Float32:
		return reflect.ValueOf(*(*float32)(p)).Convert(t)
	case reflect.Float64:
		return reflect.ValueOf(*(*float64)(p)).Convert(t)
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
			// objc object pointer holder struct should have same layout as a pointer
			return reflect.NewAt(t, p).Elem()
		} else {
			return reflect.NewAt(t, p).Elem()
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

	rt := v.Type()
	switch rt.Kind() {
	case reflect.Bool:
		var cv uint8 = 0
		if v.Bool() {
			cv = 1
		}
		return unsafe.Pointer(&cv)
	case reflect.Int8:
		cv := int8(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Int16:
		cv := int16(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Int32:
		cv := int32(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Int, reflect.Int64:
		cv := int64(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Uint8:
		cv := uint8(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Uint16:
		cv := uint16(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Uint32:
		cv := uint32(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		cv := uint64(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Float32:
		cv := float32(v.Float())
		return unsafe.Pointer(&cv)
	case reflect.Float64:
		cv := float64(v.Float())
		return unsafe.Pointer(&cv)
	case reflect.UnsafePointer, reflect.Pointer:
		cv := v.UnsafePointer()
		return unsafe.Pointer(&cv)
	case reflect.Interface:
		if v.Type().AssignableTo(pointerHolderType) {
			cv := v.Interface().(Holder).Ptr()
			return unsafe.Pointer(&cv)
		}
		panic(fmt.Sprintf("not support type: %T", v))
	case reflect.Struct:
		if v.Type().Implements(pointerHolderType) {
			cv := v.Interface().(Holder).Ptr()
			return unsafe.Pointer(&cv)
		}
		return getStructValuePointer(v)
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
		if rv.Type().AssignableTo(pointerHolderType) {
			*(*unsafe.Pointer)(p) = rv.Interface().(Holder).Ptr()
		} else {
			panic(fmt.Sprintf("not support type: %v", rv.Type()))
		}
	case reflect.Struct:
		if rv.Type().Implements(pointerHolderType) {
			*(*unsafe.Pointer)(p) = rv.Interface().(Holder).Ptr()
		} else {
			sp := getStructValuePointer(rv)
			size := rv.Type().Size()
			copy(unsafe.Slice((*byte)(p), size), unsafe.Slice((*byte)(sp), size))
		}
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

func getStructValuePointer(value reflect.Value) unsafe.Pointer {
	vsp := (*reValue)(unsafe.Pointer(&value))

	if vsp.flag&flagIndir == 0 {
		return unsafe.Pointer(&vsp.ptr)
	} else {
		return vsp.ptr
	}
}

// toFFIType return the ffi type
func toFFIType(rt reflect.Type) *ffi.Type {
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
		if rt.AssignableTo(pointerHolderType) {
			return ffi.TypePointer
		}
		panic(fmt.Sprintf("not support type: %v", rt))
	case reflect.Struct:
		if rt.Size() == 0 {
			return ffi.TypeVoid
		}
		if rt.Implements(pointerHolderType) {
			return ffi.TypePointer
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
	return ffi.MakeStructType(fts)
}
