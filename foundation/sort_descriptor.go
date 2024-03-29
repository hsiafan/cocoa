// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[SortDescriptor](sc, objc.SEL("sortDescriptorWithKey:ascending:"), key, ascending)
	return rv
}

func (s_ SortDescriptor) InitWithKey_Ascending(key string, ascending bool) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](s_, objc.SEL("initWithKey:ascending:"), key, ascending)
	return rv
}

func (sc _SortDescriptorClass) SortDescriptorWithKey_Ascending_Selector(key string, ascending bool, selector objc.Selector) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](sc, objc.SEL("sortDescriptorWithKey:ascending:selector:"), key, ascending, selector)
	return rv
}

func (s_ SortDescriptor) InitWithKey_Ascending_Selector(key string, ascending bool, selector objc.Selector) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](s_, objc.SEL("initWithKey:ascending:selector:"), key, ascending, selector)
	return rv
}

func (sc _SortDescriptorClass) SortDescriptorWithKey_Ascending_Comparator(key string, ascending bool, cmptr func(obj1 objc.Object, obj2 objc.Object) ComparisonResult) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](sc, objc.SEL("sortDescriptorWithKey:ascending:comparator:"), key, ascending, cmptr)
	return rv
}

func (s_ SortDescriptor) InitWithKey_Ascending_Comparator(key string, ascending bool, cmptr func(obj1 objc.Object, obj2 objc.Object) ComparisonResult) SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](s_, objc.SEL("initWithKey:ascending:comparator:"), key, ascending, cmptr)
	return rv
}

func (sc _SortDescriptorClass) Alloc() SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SortDescriptorClass) New() SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSortDescriptor() SortDescriptor {
	return SortDescriptorClass.New()
}

func (s_ SortDescriptor) Init() SortDescriptor {
	rv := objc.CallMethod[SortDescriptor](s_, objc.SEL("init"))
	return rv
}

func (s_ SortDescriptor) CompareObject_ToObject(object1 objc.IObject, object2 objc.IObject) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](s_, objc.SEL("compareObject:toObject:"), objc.ExtractPtr(object1), objc.ExtractPtr(object2))
	return rv
}

func (s_ SortDescriptor) AllowEvaluation() {
	objc.CallMethod[objc.Void](s_, objc.SEL("allowEvaluation"))
}

func (s_ SortDescriptor) Ascending() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("ascending"))
	return rv
}

func (s_ SortDescriptor) Key() string {
	rv := objc.CallMethod[string](s_, objc.SEL("key"))
	return rv
}

func (s_ SortDescriptor) Selector() objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, objc.SEL("selector"))
	return rv
}

func (s_ SortDescriptor) Comparator() func(obj1 objc.IObject, obj2 objc.IObject) ComparisonResult {
	rv := objc.CallMethod[func(obj1 objc.IObject, obj2 objc.IObject) ComparisonResult](s_, objc.SEL("comparator"))
	return rv
}

func (s_ SortDescriptor) ReversedSortDescriptor() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("reversedSortDescriptor"))
	return rv
}
