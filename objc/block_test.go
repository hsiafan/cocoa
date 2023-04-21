package objc

import (
	"reflect"
	"testing"
)

func Test_getBlockTypeEncoding(t *testing.T) {
	var f = func(index uint, value Object) {

	}
	encoding := getBlockTypeEncoding(reflect.TypeOf(f))
	if encoding != "v@?Q@" {
		t.Fail()
	}
}
