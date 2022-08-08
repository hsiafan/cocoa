// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var EnumeratorClass = _EnumeratorClass{objc.GetClass("NSEnumerator")}

type _EnumeratorClass struct {
	objc.Class
}

type IEnumerator interface {
	objc.IObject
	NextObject() objc.Object
	AllObjects() []objc.Object
}

type Enumerator struct {
	objc.Object
}

func MakeEnumerator(ptr unsafe.Pointer) Enumerator {
	return Enumerator{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _EnumeratorClass) Alloc() Enumerator {
	rv := ffi.CallMethod[Enumerator](ec, "alloc")
	return rv
}

func (e_ Enumerator) Init() Enumerator {
	rv := ffi.CallMethod[Enumerator](e_, "init")
	return rv
}

func (ec _EnumeratorClass) New() Enumerator {
	rv := ffi.CallMethod[Enumerator](ec, "new")
	rv.Autorelease()
	return rv
}

func NewEnumerator() Enumerator {
	return EnumeratorClass.New()
}

func (e_ Enumerator) NextObject() objc.Object {
	rv := ffi.CallMethod[objc.Object](e_, "nextObject")
	return rv
}

func (e_ Enumerator) AllObjects() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](e_, "allObjects")
	return rv
}
