// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewFlowLayoutClass = _CollectionViewFlowLayoutClass{objc.GetClass("NSCollectionViewFlowLayout")}

type _CollectionViewFlowLayoutClass struct {
	objc.Class
}

type ICollectionViewFlowLayout interface {
	ICollectionViewLayout
	CollapseSectionAtIndex(sectionIndex uint)
	ExpandSectionAtIndex(sectionIndex uint)
	SectionAtIndexIsCollapsed(sectionIndex uint) bool
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	EstimatedItemSize() foundation.Size
	SetEstimatedItemSize(value foundation.Size)
	ItemSize() foundation.Size
	SetItemSize(value foundation.Size)
	SectionInset() foundation.EdgeInsets
	SetSectionInset(value foundation.EdgeInsets)
	HeaderReferenceSize() foundation.Size
	SetHeaderReferenceSize(value foundation.Size)
	FooterReferenceSize() foundation.Size
	SetFooterReferenceSize(value foundation.Size)
	SectionFootersPinToVisibleBounds() bool
	SetSectionFootersPinToVisibleBounds(value bool)
	SectionHeadersPinToVisibleBounds() bool
	SetSectionHeadersPinToVisibleBounds(value bool)
}

type CollectionViewFlowLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewFlowLayout(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (cc _CollectionViewFlowLayoutClass) Alloc() CollectionViewFlowLayout {
	rv := ffi.CallMethod[CollectionViewFlowLayout](cc, "alloc")
	return rv
}

func (cc _CollectionViewFlowLayoutClass) New() CollectionViewFlowLayout {
	rv := ffi.CallMethod[CollectionViewFlowLayout](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewFlowLayout() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.New()
}

func (c_ CollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	rv := ffi.CallMethod[CollectionViewFlowLayout](c_, "init")
	return rv
}

func (c_ CollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	ffi.CallMethod[ffi.Void](c_, "collapseSectionAtIndex:", sectionIndex)
}

func (c_ CollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	ffi.CallMethod[ffi.Void](c_, "expandSectionAtIndex:", sectionIndex)
}

func (c_ CollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	rv := ffi.CallMethod[bool](c_, "sectionAtIndexIsCollapsed:", sectionIndex)
	return rv
}

func (c_ CollectionViewFlowLayout) ScrollDirection() CollectionViewScrollDirection {
	rv := ffi.CallMethod[CollectionViewScrollDirection](c_, "scrollDirection")
	return rv
}

func (c_ CollectionViewFlowLayout) SetScrollDirection(value CollectionViewScrollDirection) {
	ffi.CallMethod[ffi.Void](c_, "setScrollDirection:", value)
}

func (c_ CollectionViewFlowLayout) MinimumLineSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "minimumLineSpacing")
	return rv
}

func (c_ CollectionViewFlowLayout) SetMinimumLineSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumLineSpacing:", value)
}

func (c_ CollectionViewFlowLayout) MinimumInteritemSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "minimumInteritemSpacing")
	return rv
}

func (c_ CollectionViewFlowLayout) SetMinimumInteritemSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumInteritemSpacing:", value)
}

func (c_ CollectionViewFlowLayout) EstimatedItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "estimatedItemSize")
	return rv
}

func (c_ CollectionViewFlowLayout) SetEstimatedItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setEstimatedItemSize:", value)
}

func (c_ CollectionViewFlowLayout) ItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "itemSize")
	return rv
}

func (c_ CollectionViewFlowLayout) SetItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setItemSize:", value)
}

func (c_ CollectionViewFlowLayout) SectionInset() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](c_, "sectionInset")
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionInset(value foundation.EdgeInsets) {
	ffi.CallMethod[ffi.Void](c_, "setSectionInset:", value)
}

func (c_ CollectionViewFlowLayout) HeaderReferenceSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "headerReferenceSize")
	return rv
}

func (c_ CollectionViewFlowLayout) SetHeaderReferenceSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setHeaderReferenceSize:", value)
}

func (c_ CollectionViewFlowLayout) FooterReferenceSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "footerReferenceSize")
	return rv
}

func (c_ CollectionViewFlowLayout) SetFooterReferenceSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setFooterReferenceSize:", value)
}

func (c_ CollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	rv := ffi.CallMethod[bool](c_, "sectionFootersPinToVisibleBounds")
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSectionFootersPinToVisibleBounds:", value)
}

func (c_ CollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	rv := ffi.CallMethod[bool](c_, "sectionHeadersPinToVisibleBounds")
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSectionHeadersPinToVisibleBounds:", value)
}
