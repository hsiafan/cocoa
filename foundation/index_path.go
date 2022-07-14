// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[IndexPath](ic, "indexPathWithIndex:", index)
	return rv
}

func (ic _IndexPathClass) IndexPathWithIndexes_Length(indexes *uint, length uint) IndexPath {
	rv := ffi.CallMethod[IndexPath](ic, "indexPathWithIndexes:length:", indexes, length)
	return rv
}

func (i_ IndexPath) InitWithIndex(index uint) IndexPath {
	rv := ffi.CallMethod[IndexPath](i_, "initWithIndex:", index)
	rv.Autorelease()
	return rv
}

func (i_ IndexPath) InitWithIndexes_Length(indexes *uint, length uint) IndexPath {
	rv := ffi.CallMethod[IndexPath](i_, "initWithIndexes:length:", indexes, length)
	rv.Autorelease()
	return rv
}

func (ic _IndexPathClass) Alloc() IndexPath {
	rv := ffi.CallMethod[IndexPath](ic, "alloc")
	return rv
}

func (i_ IndexPath) Init() IndexPath {
	rv := ffi.CallMethod[IndexPath](i_, "init")
	rv.Autorelease()
	return rv
}

func (ic _IndexPathClass) New() IndexPath {
	rv := ffi.CallMethod[IndexPath](ic, "new")
	rv.Autorelease()
	return rv
}

func NewIndexPath() IndexPath {
	return IndexPathClass.New()
}

func (i_ IndexPath) IndexPathByAddingIndex(index uint) IndexPath {
	rv := ffi.CallMethod[IndexPath](i_, "indexPathByAddingIndex:", index)
	return rv
}

func (i_ IndexPath) IndexPathByRemovingLastIndex() IndexPath {
	rv := ffi.CallMethod[IndexPath](i_, "indexPathByRemovingLastIndex")
	return rv
}

func (i_ IndexPath) Compare(otherObject IIndexPath) ComparisonResult {
	rv := ffi.CallMethod[ComparisonResult](i_, "compare:", otherObject)
	return rv
}

func (i_ IndexPath) IndexAtPosition(position uint) uint {
	rv := ffi.CallMethod[uint](i_, "indexAtPosition:", position)
	return rv
}

func (i_ IndexPath) GetIndexes_Range(indexes *uint, positionRange Range) {
	ffi.CallMethod[ffi.Void](i_, "getIndexes:range:", indexes, positionRange)
}

// deprecated
func (i_ IndexPath) GetIndexes(indexes *uint) {
	ffi.CallMethod[ffi.Void](i_, "getIndexes:", indexes)
}

func (i_ IndexPath) Length() uint {
	rv := ffi.CallMethod[uint](i_, "length")
	return rv
}
