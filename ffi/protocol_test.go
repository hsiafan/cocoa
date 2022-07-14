package ffi

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hsiafan/cocoa/objc"
)

func Test_stringConvert(t *testing.T) {
	var s = "test"
	ptr := ToNSString(s)
	ns := ToGoString(ptr)
	if ns != s {
		t.Fail()
	}
}

func Test_bytesConvert(t *testing.T) {
	var bs = []byte("test")
	ptr := ToNSData(bs)
	nbs := ToGoBytes(ptr)
	if bytes.Compare(bs, nbs) != 0 {
		t.Fail()
	}
}

func Test_sliceConvert(t *testing.T) {
	s := []string{"1", "2"}
	ptr := ToNSArray(reflect.ValueOf(s))
	ns := ToGoSlice(ptr, reflect.TypeOf(s)).Interface().([]string)
	if len(s) != len(ns) || s[0] != ns[0] {
		t.Fail()
	}

	bs := [][]byte{[]byte("test")}
	ptr = ToNSArray(reflect.ValueOf(bs))
	nbs := ToGoSlice(ptr, reflect.TypeOf(bs)).Interface().([][]byte)
	if len(bs) != len(nbs) || bytes.Compare(bs[0], nbs[0]) != 0 {
		t.Fail()
	}
}

func Test_mapConvert(t *testing.T) {
	o1 := objc.NewObject()
	o2 := objc.NewObject()
	m := map[string]objc.IObject{
		"1": o1,
		"2": o2,
	}
	ptr := ToNSDict(reflect.ValueOf(m))
	nm := ToGoMap(ptr, reflect.TypeOf(map[string]objc.Object{})).Interface().(map[string]objc.Object)
	if len(m) != len(nm) || m["1"].(objc.Object) != nm["1"] {
		t.Fail()
	}
}

func Test_Wrap(t *testing.T) {

}
