package ffi

import (
	"testing"

	"github.com/hsiafan/cocoa/objc"
)

func TestCallMethod(t *testing.T) {
	var o = objc.NewObject()
	var count = CallMethod[uint](o, "retainCount")
	if count != 1 {
		t.Fail()
	}
}
