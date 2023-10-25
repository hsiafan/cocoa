// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionViewFlowLayout](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionViewFlowLayoutClass) New() CollectionViewFlowLayout {
	rv := objc.CallMethod[CollectionViewFlowLayout](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewFlowLayout() CollectionViewFlowLayout {
	return CollectionViewFlowLayoutClass.New()
}

func (c_ CollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	rv := objc.CallMethod[CollectionViewFlowLayout](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionViewFlowLayout) CollapseSectionAtIndex(sectionIndex uint) {
	objc.CallMethod[objc.Void](c_, objc.SEL("collapseSectionAtIndex:"), sectionIndex)
}

func (c_ CollectionViewFlowLayout) ExpandSectionAtIndex(sectionIndex uint) {
	objc.CallMethod[objc.Void](c_, objc.SEL("expandSectionAtIndex:"), sectionIndex)
}

func (c_ CollectionViewFlowLayout) SectionAtIndexIsCollapsed(sectionIndex uint) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("sectionAtIndexIsCollapsed:"), sectionIndex)
	return rv
}

func (c_ CollectionViewFlowLayout) ScrollDirection() CollectionViewScrollDirection {
	rv := objc.CallMethod[CollectionViewScrollDirection](c_, objc.SEL("scrollDirection"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetScrollDirection(value CollectionViewScrollDirection) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setScrollDirection:"), value)
}

func (c_ CollectionViewFlowLayout) MinimumLineSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("minimumLineSpacing"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetMinimumLineSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setMinimumLineSpacing:"), value)
}

func (c_ CollectionViewFlowLayout) MinimumInteritemSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("minimumInteritemSpacing"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetMinimumInteritemSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setMinimumInteritemSpacing:"), value)
}

func (c_ CollectionViewFlowLayout) EstimatedItemSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("estimatedItemSize"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetEstimatedItemSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setEstimatedItemSize:"), value)
}

func (c_ CollectionViewFlowLayout) ItemSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("itemSize"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetItemSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setItemSize:"), value)
}

func (c_ CollectionViewFlowLayout) SectionInset() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](c_, objc.SEL("sectionInset"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionInset(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setSectionInset:"), value)
}

func (c_ CollectionViewFlowLayout) HeaderReferenceSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("headerReferenceSize"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetHeaderReferenceSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setHeaderReferenceSize:"), value)
}

func (c_ CollectionViewFlowLayout) FooterReferenceSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.SEL("footerReferenceSize"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetFooterReferenceSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setFooterReferenceSize:"), value)
}

func (c_ CollectionViewFlowLayout) SectionFootersPinToVisibleBounds() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("sectionFootersPinToVisibleBounds"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionFootersPinToVisibleBounds(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setSectionFootersPinToVisibleBounds:"), value)
}

func (c_ CollectionViewFlowLayout) SectionHeadersPinToVisibleBounds() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("sectionHeadersPinToVisibleBounds"))
	return rv
}

func (c_ CollectionViewFlowLayout) SetSectionHeadersPinToVisibleBounds(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setSectionHeadersPinToVisibleBounds:"), value)
}
