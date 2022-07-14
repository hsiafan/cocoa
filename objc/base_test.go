package objc

import (
	"testing"
	"unsafe"
)

func TestGetClass(t *testing.T) {
	cls := GetClass("NSObject")

	version := cls.GetVersion()
	if version != 0 {
		t.Failed()
	}

	if cls.GetName() != "NSObject" {
		t.Failed()
	}
}

func TestGetMethod(t *testing.T) {
	cls := GetClass("NSObject")

	m := cls.GetClassMethod(SelectorRegisterName("alloc"))
	name := m.GetName().GetName()
	if name != "alloc" {
		t.Failed()
	}
	te := m.GetTypeEncoding()
	if te != "@16@0:8" {
		t.Failed()
	}
}

func TestWithAutoreleasePool(t *testing.T) {

	var o = NewObject()
	WithAutoreleasePool(func() {
		o.Autorelease()
		if o.RetainCount() != 1 {
			t.Failed()
		}
	})
}

func TestClass_CopyMethodList(t *testing.T) {
	cls := GetClass("NSObject")
	ms := cls.CopyMethodList()
	for _, m := range ms {
		if m.ptr == nil {
			t.Fail()
		}
	}
}

func TestSetAssociatedObject(t *testing.T) {
	o1 := NewObject()
	o2 := NewObject()
	var i = 0
	key := unsafe.Pointer(&i)
	SetAssociatedObject(o1, key, o2, ASSOCIATION_RETAIN)
	if o2.RetainCount() != 2 {
		t.Fail()
	}
	o1.Release()
	if o2.RetainCount() != 1 {
		t.Fail()
	}
	o2.Release()

	o1 = NewObject()
	o2 = NewObject()
	o3 := NewObject()
	SetAssociatedObject(o1, key, o2, ASSOCIATION_RETAIN)
	if o2.RetainCount() != 2 {
		t.Fail()
	}
	SetAssociatedObject(o1, key, o3, ASSOCIATION_RETAIN)
	if o2.RetainCount() != 1 {
		t.Fail()
	}
	o1.Release()
	if o2.RetainCount() != 1 || o3.RetainCount() != 1 {
		t.Fail()
	}
	o2.Release()
	o3.Release()
}

func TestClass_CopyPropertyList(t *testing.T) {
	oc := GetClass("NSObject")
	ps := oc.CopyPropertyList()
	for _, p := range ps {
		_ = p
		// fmt.Println(p.GetAttributes())
		// fmt.Println(p.CopyAttributeList())
	}
}
