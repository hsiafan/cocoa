// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var SortDescriptorClass = _SortDescriptorClass{objc.GetClass("NSSortDescriptor")}

type _SortDescriptorClass struct {
	objc.Class
}

type ISortDescriptor interface {
	objc.IObject
	CompareObject_ToObject(object1 objc.IObject, object2 objc.IObject) ComparisonResult
	AllowEvaluation()
	Ascending() bool
	Key() string
	Selector() objc.Selector
	Comparator() func(obj1 objc.IObject, obj2 objc.IObject) ComparisonResult
	ReversedSortDescriptor() objc.Object
}

type SortDescriptor struct {
	objc.Object
}

func MakeSortDescriptor(ptr unsafe.Pointer) SortDescriptor {
	return SortDescriptor{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _SortDescriptorClass) SortDescriptorWithKey_Ascending(key string, ascending bool) SortDescriptor {
	rv := ffi.CallMethod[SortDescriptor](sc, "sortDescriptorWithKey:ascending:", key, ascending)
	return rv
}

func (s_ SortDescriptor) InitWithKey_Ascending(key string, ascending bool) SortDescriptor {
	rv := ffi.CallMethod[SortDescriptor](s_, "initWithKey:ascending:", key, ascending)
	return rv
}

func (sc _SortDescriptorClass) SortDescriptorWithKey_Ascending_Selector(key string, ascending bool, selector objc.Selector) SortDescriptor {
	rv := ffi.CallMethod[SortDescriptor](sc, "sortDescriptorWithKey:ascending:selector:", key, ascending, selector)
	return rv
}

func (s_ SortDescriptor) InitWithKey_Ascending_Selector(key string, ascending bool, selector objc.Selector) SortDescriptor {
	rv := ffi.CallMethod[SortDescriptor](s_, "initWithKey:ascending:selector:", key, ascending, selector)
	return rv
}

func (sc _SortDescriptorClass) SortDescriptorWithKey_Ascending_Comparator(key string, ascending bool, cmptr func(obj1 objc.Object, obj2 objc.Object) ComparisonResult) SortDescriptor {
	rv := ffi.CallMethod[SortDescriptor](sc, "sortDescriptorWithKey:ascending:comparator:", key, ascending, cmptr)
	return rv
}

func (s_ SortDescriptor) InitWithKey_Ascending_Comparator(key string, ascending bool, cmptr func(obj1 objc.Object, obj2 objc.Object) ComparisonResult) SortDescriptor {
	rv := ffi.CallMethod[SortDescriptor](s_, "initWithKey:ascending:comparator:", key, ascending, cmptr)
	return rv
}

func (sc _SortDescriptorClass) Alloc() SortDescriptor {
	rv := ffi.CallMethod[SortDescriptor](sc, "alloc")
	return rv
}

func (sc _SortDescriptorClass) New() SortDescriptor {
	rv := ffi.CallMethod[SortDescriptor](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSortDescriptor() SortDescriptor {
	return SortDescriptorClass.New()
}

func (s_ SortDescriptor) Init() SortDescriptor {
	rv := ffi.CallMethod[SortDescriptor](s_, "init")
	return rv
}

func (s_ SortDescriptor) CompareObject_ToObject(object1 objc.IObject, object2 objc.IObject) ComparisonResult {
	rv := ffi.CallMethod[ComparisonResult](s_, "compareObject:toObject:", object1, object2)
	return rv
}

func (s_ SortDescriptor) AllowEvaluation() {
	ffi.CallMethod[ffi.Void](s_, "allowEvaluation")
}

func (s_ SortDescriptor) Ascending() bool {
	rv := ffi.CallMethod[bool](s_, "ascending")
	return rv
}

func (s_ SortDescriptor) Key() string {
	rv := ffi.CallMethod[string](s_, "key")
	return rv
}

func (s_ SortDescriptor) Selector() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](s_, "selector")
	return rv
}

func (s_ SortDescriptor) Comparator() func(obj1 objc.IObject, obj2 objc.IObject) ComparisonResult {
	rv := ffi.CallMethod[func(obj1 objc.IObject, obj2 objc.IObject) ComparisonResult](s_, "comparator")
	return rv
}

func (s_ SortDescriptor) ReversedSortDescriptor() objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "reversedSortDescriptor")
	return rv
}
