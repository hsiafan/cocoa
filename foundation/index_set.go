// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var IndexSetClass = _IndexSetClass{objc.GetClass("NSIndexSet")}

type _IndexSetClass struct {
	objc.Class
}

type IIndexSet interface {
	objc.IObject
	ContainsIndex(value uint) bool
	ContainsIndexes(indexSet IIndexSet) bool
	ContainsIndexesInRange(_range Range) bool
	IntersectsIndexesInRange(_range Range) bool
	CountOfIndexesInRange(_range Range) uint
	IndexPassingTest(predicate func(idx uint, stop *bool) bool) uint
	IndexesPassingTest(predicate func(idx uint, stop *bool) bool) IndexSet
	IndexWithOptions_PassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint
	IndexesWithOptions_PassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet
	IndexInRange_Options_PassingTest(_range Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint
	IndexesInRange_Options_PassingTest(_range Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet
	EnumerateRangesInRange_Options_UsingBlock(_range Range, opts EnumerationOptions, block func(_range Range, stop *bool))
	EnumerateRangesUsingBlock(block func(_range Range, stop *bool))
	EnumerateRangesWithOptions_UsingBlock(opts EnumerationOptions, block func(_range Range, stop *bool))
	IsEqualToIndexSet(indexSet IIndexSet) bool
	IndexLessThanIndex(value uint) uint
	IndexLessThanOrEqualToIndex(value uint) uint
	IndexGreaterThanOrEqualToIndex(value uint) uint
	IndexGreaterThanIndex(value uint) uint
	GetIndexes_MaxCount_InIndexRange(indexBuffer *uint, bufferSize uint, _range *Range) uint
	EnumerateIndexesUsingBlock(block func(idx uint, stop *bool))
	EnumerateIndexesWithOptions_UsingBlock(opts EnumerationOptions, block func(idx uint, stop *bool))
	EnumerateIndexesInRange_Options_UsingBlock(_range Range, opts EnumerationOptions, block func(idx uint, stop *bool))
	Count() uint
	FirstIndex() uint
	LastIndex() uint
}

type IndexSet struct {
	objc.Object
}

func MakeIndexSet(ptr unsafe.Pointer) IndexSet {
	return IndexSet{
		Object: objc.MakeObject(ptr),
	}
}

func (ic _IndexSetClass) IndexSet() IndexSet {
	rv := ffi.CallMethod[IndexSet](ic, "indexSet")
	return rv
}

func (ic _IndexSetClass) IndexSetWithIndex(value uint) IndexSet {
	rv := ffi.CallMethod[IndexSet](ic, "indexSetWithIndex:", value)
	return rv
}

func (ic _IndexSetClass) IndexSetWithIndexesInRange(_range Range) IndexSet {
	rv := ffi.CallMethod[IndexSet](ic, "indexSetWithIndexesInRange:", _range)
	return rv
}

func (i_ IndexSet) InitWithIndex(value uint) IndexSet {
	rv := ffi.CallMethod[IndexSet](i_, "initWithIndex:", value)
	return rv
}

func (i_ IndexSet) InitWithIndexesInRange(_range Range) IndexSet {
	rv := ffi.CallMethod[IndexSet](i_, "initWithIndexesInRange:", _range)
	return rv
}

func (i_ IndexSet) InitWithIndexSet(indexSet IIndexSet) IndexSet {
	rv := ffi.CallMethod[IndexSet](i_, "initWithIndexSet:", indexSet)
	return rv
}

func (ic _IndexSetClass) Alloc() IndexSet {
	rv := ffi.CallMethod[IndexSet](ic, "alloc")
	return rv
}

func (i_ IndexSet) Init() IndexSet {
	rv := ffi.CallMethod[IndexSet](i_, "init")
	return rv
}

func (ic _IndexSetClass) New() IndexSet {
	rv := ffi.CallMethod[IndexSet](ic, "new")
	rv.Autorelease()
	return rv
}

func NewIndexSet() IndexSet {
	return IndexSetClass.New()
}

func (i_ IndexSet) ContainsIndex(value uint) bool {
	rv := ffi.CallMethod[bool](i_, "containsIndex:", value)
	return rv
}

func (i_ IndexSet) ContainsIndexes(indexSet IIndexSet) bool {
	rv := ffi.CallMethod[bool](i_, "containsIndexes:", indexSet)
	return rv
}

func (i_ IndexSet) ContainsIndexesInRange(_range Range) bool {
	rv := ffi.CallMethod[bool](i_, "containsIndexesInRange:", _range)
	return rv
}

func (i_ IndexSet) IntersectsIndexesInRange(_range Range) bool {
	rv := ffi.CallMethod[bool](i_, "intersectsIndexesInRange:", _range)
	return rv
}

func (i_ IndexSet) CountOfIndexesInRange(_range Range) uint {
	rv := ffi.CallMethod[uint](i_, "countOfIndexesInRange:", _range)
	return rv
}

func (i_ IndexSet) IndexPassingTest(predicate func(idx uint, stop *bool) bool) uint {
	rv := ffi.CallMethod[uint](i_, "indexPassingTest:", predicate)
	return rv
}

func (i_ IndexSet) IndexesPassingTest(predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := ffi.CallMethod[IndexSet](i_, "indexesPassingTest:", predicate)
	return rv
}

func (i_ IndexSet) IndexWithOptions_PassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint {
	rv := ffi.CallMethod[uint](i_, "indexWithOptions:passingTest:", opts, predicate)
	return rv
}

func (i_ IndexSet) IndexesWithOptions_PassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := ffi.CallMethod[IndexSet](i_, "indexesWithOptions:passingTest:", opts, predicate)
	return rv
}

func (i_ IndexSet) IndexInRange_Options_PassingTest(_range Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint {
	rv := ffi.CallMethod[uint](i_, "indexInRange:options:passingTest:", _range, opts, predicate)
	return rv
}

func (i_ IndexSet) IndexesInRange_Options_PassingTest(_range Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := ffi.CallMethod[IndexSet](i_, "indexesInRange:options:passingTest:", _range, opts, predicate)
	return rv
}

func (i_ IndexSet) EnumerateRangesInRange_Options_UsingBlock(_range Range, opts EnumerationOptions, block func(_range Range, stop *bool)) {
	ffi.CallMethod[ffi.Void](i_, "enumerateRangesInRange:options:usingBlock:", _range, opts, block)
}

func (i_ IndexSet) EnumerateRangesUsingBlock(block func(_range Range, stop *bool)) {
	ffi.CallMethod[ffi.Void](i_, "enumerateRangesUsingBlock:", block)
}

func (i_ IndexSet) EnumerateRangesWithOptions_UsingBlock(opts EnumerationOptions, block func(_range Range, stop *bool)) {
	ffi.CallMethod[ffi.Void](i_, "enumerateRangesWithOptions:usingBlock:", opts, block)
}

func (i_ IndexSet) IsEqualToIndexSet(indexSet IIndexSet) bool {
	rv := ffi.CallMethod[bool](i_, "isEqualToIndexSet:", indexSet)
	return rv
}

func (i_ IndexSet) IndexLessThanIndex(value uint) uint {
	rv := ffi.CallMethod[uint](i_, "indexLessThanIndex:", value)
	return rv
}

func (i_ IndexSet) IndexLessThanOrEqualToIndex(value uint) uint {
	rv := ffi.CallMethod[uint](i_, "indexLessThanOrEqualToIndex:", value)
	return rv
}

func (i_ IndexSet) IndexGreaterThanOrEqualToIndex(value uint) uint {
	rv := ffi.CallMethod[uint](i_, "indexGreaterThanOrEqualToIndex:", value)
	return rv
}

func (i_ IndexSet) IndexGreaterThanIndex(value uint) uint {
	rv := ffi.CallMethod[uint](i_, "indexGreaterThanIndex:", value)
	return rv
}

func (i_ IndexSet) GetIndexes_MaxCount_InIndexRange(indexBuffer *uint, bufferSize uint, _range *Range) uint {
	rv := ffi.CallMethod[uint](i_, "getIndexes:maxCount:inIndexRange:", indexBuffer, bufferSize, _range)
	return rv
}

func (i_ IndexSet) EnumerateIndexesUsingBlock(block func(idx uint, stop *bool)) {
	ffi.CallMethod[ffi.Void](i_, "enumerateIndexesUsingBlock:", block)
}

func (i_ IndexSet) EnumerateIndexesWithOptions_UsingBlock(opts EnumerationOptions, block func(idx uint, stop *bool)) {
	ffi.CallMethod[ffi.Void](i_, "enumerateIndexesWithOptions:usingBlock:", opts, block)
}

func (i_ IndexSet) EnumerateIndexesInRange_Options_UsingBlock(_range Range, opts EnumerationOptions, block func(idx uint, stop *bool)) {
	ffi.CallMethod[ffi.Void](i_, "enumerateIndexesInRange:options:usingBlock:", _range, opts, block)
}

func (i_ IndexSet) Count() uint {
	rv := ffi.CallMethod[uint](i_, "count")
	return rv
}

func (i_ IndexSet) FirstIndex() uint {
	rv := ffi.CallMethod[uint](i_, "firstIndex")
	return rv
}

func (i_ IndexSet) LastIndex() uint {
	rv := ffi.CallMethod[uint](i_, "lastIndex")
	return rv
}
