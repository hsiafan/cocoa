// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	ContainsIndexesInRange(range_ Range) bool
	IntersectsIndexesInRange(range_ Range) bool
	CountOfIndexesInRange(range_ Range) uint
	IndexPassingTest(predicate func(idx uint, stop *bool) bool) uint
	IndexesPassingTest(predicate func(idx uint, stop *bool) bool) IndexSet
	IndexWithOptions_PassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint
	IndexesWithOptions_PassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet
	IndexInRange_Options_PassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint
	IndexesInRange_Options_PassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet
	EnumerateRangesInRange_Options_UsingBlock(range_ Range, opts EnumerationOptions, block func(range_ Range, stop *bool))
	EnumerateRangesUsingBlock(block func(range_ Range, stop *bool))
	EnumerateRangesWithOptions_UsingBlock(opts EnumerationOptions, block func(range_ Range, stop *bool))
	IsEqualToIndexSet(indexSet IIndexSet) bool
	IndexLessThanIndex(value uint) uint
	IndexLessThanOrEqualToIndex(value uint) uint
	IndexGreaterThanOrEqualToIndex(value uint) uint
	IndexGreaterThanIndex(value uint) uint
	GetIndexes_MaxCount_InIndexRange(indexBuffer *uint, bufferSize uint, range_ *Range) uint
	EnumerateIndexesUsingBlock(block func(idx uint, stop *bool))
	EnumerateIndexesWithOptions_UsingBlock(opts EnumerationOptions, block func(idx uint, stop *bool))
	EnumerateIndexesInRange_Options_UsingBlock(range_ Range, opts EnumerationOptions, block func(idx uint, stop *bool))
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
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("indexSet"))
	return rv
}

func (ic _IndexSetClass) IndexSetWithIndex(value uint) IndexSet {
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("indexSetWithIndex:"), value)
	return rv
}

func (ic _IndexSetClass) IndexSetWithIndexesInRange(range_ Range) IndexSet {
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("indexSetWithIndexesInRange:"), range_)
	return rv
}

func (i_ IndexSet) InitWithIndex(value uint) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("initWithIndex:"), value)
	return rv
}

func (i_ IndexSet) InitWithIndexesInRange(range_ Range) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("initWithIndexesInRange:"), range_)
	return rv
}

func (i_ IndexSet) InitWithIndexSet(indexSet IIndexSet) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("initWithIndexSet:"), objc.ExtractPtr(indexSet))
	return rv
}

func (ic _IndexSetClass) Alloc() IndexSet {
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("alloc"))
	return rv
}

func (ic _IndexSetClass) New() IndexSet {
	rv := objc.CallMethod[IndexSet](ic, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewIndexSet() IndexSet {
	return IndexSetClass.New()
}

func (i_ IndexSet) Init() IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("init"))
	return rv
}

func (i_ IndexSet) ContainsIndex(value uint) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("containsIndex:"), value)
	return rv
}

func (i_ IndexSet) ContainsIndexes(indexSet IIndexSet) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("containsIndexes:"), objc.ExtractPtr(indexSet))
	return rv
}

func (i_ IndexSet) ContainsIndexesInRange(range_ Range) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("containsIndexesInRange:"), range_)
	return rv
}

func (i_ IndexSet) IntersectsIndexesInRange(range_ Range) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("intersectsIndexesInRange:"), range_)
	return rv
}

func (i_ IndexSet) CountOfIndexesInRange(range_ Range) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("countOfIndexesInRange:"), range_)
	return rv
}

func (i_ IndexSet) IndexPassingTest(predicate func(idx uint, stop *bool) bool) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexPassingTest:"), predicate)
	return rv
}

func (i_ IndexSet) IndexesPassingTest(predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("indexesPassingTest:"), predicate)
	return rv
}

func (i_ IndexSet) IndexWithOptions_PassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexWithOptions:passingTest:"), opts, predicate)
	return rv
}

func (i_ IndexSet) IndexesWithOptions_PassingTest(opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("indexesWithOptions:passingTest:"), opts, predicate)
	return rv
}

func (i_ IndexSet) IndexInRange_Options_PassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexInRange:options:passingTest:"), range_, opts, predicate)
	return rv
}

func (i_ IndexSet) IndexesInRange_Options_PassingTest(range_ Range, opts EnumerationOptions, predicate func(idx uint, stop *bool) bool) IndexSet {
	rv := objc.CallMethod[IndexSet](i_, objc.GetSelector("indexesInRange:options:passingTest:"), range_, opts, predicate)
	return rv
}

func (i_ IndexSet) EnumerateRangesInRange_Options_UsingBlock(range_ Range, opts EnumerationOptions, block func(range_ Range, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateRangesInRange:options:usingBlock:"), range_, opts, block)
}

func (i_ IndexSet) EnumerateRangesUsingBlock(block func(range_ Range, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateRangesUsingBlock:"), block)
}

func (i_ IndexSet) EnumerateRangesWithOptions_UsingBlock(opts EnumerationOptions, block func(range_ Range, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateRangesWithOptions:usingBlock:"), opts, block)
}

func (i_ IndexSet) IsEqualToIndexSet(indexSet IIndexSet) bool {
	rv := objc.CallMethod[bool](i_, objc.GetSelector("isEqualToIndexSet:"), objc.ExtractPtr(indexSet))
	return rv
}

func (i_ IndexSet) IndexLessThanIndex(value uint) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexLessThanIndex:"), value)
	return rv
}

func (i_ IndexSet) IndexLessThanOrEqualToIndex(value uint) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexLessThanOrEqualToIndex:"), value)
	return rv
}

func (i_ IndexSet) IndexGreaterThanOrEqualToIndex(value uint) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexGreaterThanOrEqualToIndex:"), value)
	return rv
}

func (i_ IndexSet) IndexGreaterThanIndex(value uint) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("indexGreaterThanIndex:"), value)
	return rv
}

func (i_ IndexSet) GetIndexes_MaxCount_InIndexRange(indexBuffer *uint, bufferSize uint, range_ *Range) uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("getIndexes:maxCount:inIndexRange:"), indexBuffer, bufferSize, range_)
	return rv
}

func (i_ IndexSet) EnumerateIndexesUsingBlock(block func(idx uint, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateIndexesUsingBlock:"), block)
}

func (i_ IndexSet) EnumerateIndexesWithOptions_UsingBlock(opts EnumerationOptions, block func(idx uint, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateIndexesWithOptions:usingBlock:"), opts, block)
}

func (i_ IndexSet) EnumerateIndexesInRange_Options_UsingBlock(range_ Range, opts EnumerationOptions, block func(idx uint, stop *bool)) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("enumerateIndexesInRange:options:usingBlock:"), range_, opts, block)
}

func (i_ IndexSet) Count() uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("count"))
	return rv
}

func (i_ IndexSet) FirstIndex() uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("firstIndex"))
	return rv
}

func (i_ IndexSet) LastIndex() uint {
	rv := objc.CallMethod[uint](i_, objc.GetSelector("lastIndex"))
	return rv
}
