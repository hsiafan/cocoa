package objc

import (
	"testing"
)

func TestCallMethod(t *testing.T) {
	var o = NewObject()
	var count = CallMethod[uint](o, "retainCount")
	if count != 1 {
		t.Fail()
	}
}
