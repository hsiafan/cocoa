package ffi

import (
	"reflect"
	"testing"

	"github.com/hsiafan/cocoa/objc"
)

func Test_getBlockTypeEncoding(t *testing.T) {
	var f = func(index uint, value objc.Object) {

	}
	encoding := getBlockTypeEncoding(reflect.TypeOf(f))
	if encoding != "v@?Q@" {
		t.Fail()
	}
}
