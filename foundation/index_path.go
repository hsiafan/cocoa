// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var IndexPathClass = _IndexPathClass{objc.GetClass("NSIndexPath")}

type _IndexPathClass struct {
	objc.Class
}

type IIndexPath interface {
	objc.IObject
	IndexPathByAddingIndex(index uint) IndexPath
	IndexPathByRemovingLastIndex() IndexPath
	Compare(otherObject IIndexPath) ComparisonResult
	IndexAtPosition(position uint) uint
	GetIndexes_Range(indexes *uint, positionRange Range)
	// deprecated
	GetIndexes(indexes *uint)
	Length() uint
}

type IndexPath struct {
	objc.Object
}

func MakeIndexPath(ptr unsafe.Pointer) IndexPath {
	return IndexPath{
		Object: objc.MakeObject(ptr),
	}
}

func (ic _IndexPathClass) IndexPathWithIndex(index uint) IndexPath {
	rv := objc.CallMethod[IndexPath](ic, "indexPathWithIndex:", index)
	return rv
}

func (ic _IndexPathClass) IndexPathWithIndexes_Length(indexes *uint, length uint) IndexPath {
	rv := objc.CallMethod[IndexPath](ic, "indexPathWithIndexes:length:", indexes, length)
	return rv
}

func (i_ IndexPath) InitWithIndex(index uint) IndexPath {
	rv := objc.CallMethod[IndexPath](i_, "initWithIndex:", index)
	return rv
}

func (i_ IndexPath) InitWithIndexes_Length(indexes *uint, length uint) IndexPath {
	rv := objc.CallMethod[IndexPath](i_, "initWithIndexes:length:", indexes, length)
	return rv
}

func (ic _IndexPathClass) Alloc() IndexPath {
	rv := objc.CallMethod[IndexPath](ic, "alloc")
	return rv
}

func (ic _IndexPathClass) New() IndexPath {
	rv := objc.CallMethod[IndexPath](ic, "new")
	rv.Autorelease()
	return rv
}

func NewIndexPath() IndexPath {
	return IndexPathClass.New()
}

func (i_ IndexPath) Init() IndexPath {
	rv := objc.CallMethod[IndexPath](i_, "init")
	return rv
}

func (i_ IndexPath) IndexPathByAddingIndex(index uint) IndexPath {
	rv := objc.CallMethod[IndexPath](i_, "indexPathByAddingIndex:", index)
	return rv
}

func (i_ IndexPath) IndexPathByRemovingLastIndex() IndexPath {
	rv := objc.CallMethod[IndexPath](i_, "indexPathByRemovingLastIndex")
	return rv
}

func (i_ IndexPath) Compare(otherObject IIndexPath) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](i_, "compare:", otherObject)
	return rv
}

func (i_ IndexPath) IndexAtPosition(position uint) uint {
	rv := objc.CallMethod[uint](i_, "indexAtPosition:", position)
	return rv
}

func (i_ IndexPath) GetIndexes_Range(indexes *uint, positionRange Range) {
	objc.CallMethod[objc.Void](i_, "getIndexes:range:", indexes, positionRange)
}

// deprecated
func (i_ IndexPath) GetIndexes(indexes *uint) {
	objc.CallMethod[objc.Void](i_, "getIndexes:", indexes)
}

func (i_ IndexPath) Length() uint {
	rv := objc.CallMethod[uint](i_, "length")
	return rv
}
