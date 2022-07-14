package objc

import (
	"testing"
)

func TestNSObject_GetClass(t *testing.T) {
	o := NewObject()
	if o.GetClass().GetName() != "NSObject" {
		t.Fail()
	}
}

func TestNSObject_RetainCount(t *testing.T) {
	o := NewObject()
	if o.RetainCount() != 1 {
		t.Fail()
	}

	o.Retain()
	if o.RetainCount() != 2 {
		t.Fail()
	}

	o.Release()
	o.Release()
}
