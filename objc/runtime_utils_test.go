package objc

import (
	"testing"
)

func TestSetDeallocListener(t *testing.T) {
	var notified = false
	o := NewObject()
	SetDeallocListener(o, func() {
		notified = true
	})
	o.Release()
	if !notified {
		t.Fail()
	}
}
