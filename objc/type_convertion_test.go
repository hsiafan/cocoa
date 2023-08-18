package objc

import (
	"reflect"
	"testing"

	"unsafe"

	"github.com/stretchr/testify/assert"
)

// for more array/dict tests, see foundation/foundation_custom_test.go

type TestStruct struct {
	A int64
	B int64
}

type TestDirectStruct struct {
	Ptr unsafe.Pointer
}

func Test_convertStruct(t *testing.T) {
	s := TestStruct{
		A: 10,
		B: 20,
	}
	p := convertToObjcValue(reflect.ValueOf(s))
	ns := *(*TestStruct)(p)
	assert.Equal(t, s, ns)

	ns2 := convertToGoValue(p, reflect.TypeOf(TestStruct{})).Interface().(TestStruct)
	assert.Equal(t, s, ns2)

	ds := TestDirectStruct{
		Ptr: p,
	}
	dp := convertToObjcValue(reflect.ValueOf(ds))
	nds := *(*TestDirectStruct)(dp)
	assert.Equal(t, ds, nds)

	nds2 := convertToGoValue(dp, reflect.TypeOf(TestDirectStruct{})).Interface().(TestDirectStruct)
	assert.Equal(t, ds, nds2)
}

func Test_setGoValueToObjcPointer(t *testing.T) {
	s := TestStruct{
		A: 10,
		B: 20,
	}
	rv := reflect.ValueOf(s)
	var ns TestStruct
	setGoValueToObjcPointer(rv, unsafe.Pointer(&ns))
	assert.Equal(t, s, ns)

	ds := TestDirectStruct{
		Ptr: unsafe.Pointer(&ns),
	}
	drv := reflect.ValueOf(ds)
	var nds TestDirectStruct
	setGoValueToObjcPointer(drv, unsafe.Pointer(&nds))
	assert.Equal(t, ds, nds)
}

func Test_convertString(t *testing.T) {
	s := "this is a test string, and never duplicated with others。这是一个测试🐶"
	var o Object

	WithAutoreleasePool(func() {
		p := ToNSString(s)
		o = MakeObject(p)
		o.Retain()
	})
	assert.Equal(t, uint(1), o.RetainCount())
	gs := ToGoString(o.Ptr())
	assert.Equal(t, s, gs)
	o.Release()
}

func Test_convertData(t *testing.T) {
	var o Object
	WithAutoreleasePool(func() {
		p := ToNSData([]byte("this is a byte array data"))
		o = MakeObject(p)
		o.Retain()
	})

	assert.Equal(t, uint(1), o.RetainCount())
	var data = ToGoBytes(o.Ptr())
	assert.Equal(t, []byte("this is a byte array data"), data)

	o.Release()
}

func Test_convertArray(t *testing.T) {
	slice := []string{"test1", "test2"}
	var o Object
	WithAutoreleasePool(func() {
		p := ToNSArray(reflect.ValueOf(slice))
		o = MakeObject(p)
		o.Retain()
	})

	assert.Equal(t, uint(1), o.RetainCount())
	var newSlice = ToGoSlice(o.Ptr(), reflect.TypeOf(slice)).Interface().([]string)
	assert.Equal(t, slice, newSlice)

	o.Release()
}

func Test_convertDict(t *testing.T) {
	m := map[string]string{"test1": "1", "test2": "2"}
	var o Object
	WithAutoreleasePool(func() {
		p := ToNSDict(reflect.ValueOf(m))
		o = MakeObject(p)
		o.Retain()
	})

	assert.Equal(t, uint(1), o.RetainCount())
	var nm = ToGoMap(o.Ptr(), reflect.TypeOf(m)).Interface().(map[string]string)
	assert.Equal(t, m, nm)

	o.Release()
}
