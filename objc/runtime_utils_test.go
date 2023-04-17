package objc

import (
	"runtime"
	"testing"
	"time"
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

func TestPossessObject(t *testing.T) {
	o := NewObject()
	op := PossessObject(o)
	op.Description()

	runtime.GC()
	time.Sleep(time.Microsecond * 10)
	if o.RetainCount() != 1 {
		t.Fail()
	}
}
