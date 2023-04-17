// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionLayoutSectionClass = _CollectionLayoutSectionClass{objc.GetClass("NSCollectionLayoutSection")}

type _CollectionLayoutSectionClass struct {
	objc.Class
}

type ICollectionLayoutSection interface {
	objc.IObject
	OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior
	SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior)
	InterGroupSpacing() float64
	SetInterGroupSpacing(value float64)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem)
	DecorationItems() []CollectionLayoutDecorationItem
	SetDecorationItems(value []ICollectionLayoutDecorationItem)
	SupplementariesFollowContentInsets() bool
	SetSupplementariesFollowContentInsets(value bool)
}

type CollectionLayoutSection struct {
	objc.Object
}

func MakeCollectionLayoutSection(ptr unsafe.Pointer) CollectionLayoutSection {
	return CollectionLayoutSection{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _CollectionLayoutSectionClass) SectionWithGroup(group ICollectionLayoutGroup) CollectionLayoutSection {
	rv := ffi.CallMethod[CollectionLayoutSection](cc, "sectionWithGroup:", group)
	return rv
}

func (cc _CollectionLayoutSectionClass) Alloc() CollectionLayoutSection {
	rv := ffi.CallMethod[CollectionLayoutSection](cc, "alloc")
	return rv
}

func (cc _CollectionLayoutSectionClass) New() CollectionLayoutSection {
	rv := ffi.CallMethod[CollectionLayoutSection](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSection() CollectionLayoutSection {
	return CollectionLayoutSectionClass.New()
}

func (c_ CollectionLayoutSection) Init() CollectionLayoutSection {
	rv := ffi.CallMethod[CollectionLayoutSection](c_, "init")
	return rv
}

func (c_ CollectionLayoutSection) OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior {
	rv := ffi.CallMethod[CollectionLayoutSectionOrthogonalScrollingBehavior](c_, "orthogonalScrollingBehavior")
	return rv
}

func (c_ CollectionLayoutSection) SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior) {
	ffi.CallMethod[ffi.Void](c_, "setOrthogonalScrollingBehavior:", value)
}

func (c_ CollectionLayoutSection) InterGroupSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "interGroupSpacing")
	return rv
}

func (c_ CollectionLayoutSection) SetInterGroupSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setInterGroupSpacing:", value)
}

func (c_ CollectionLayoutSection) ContentInsets() DirectionalEdgeInsets {
	rv := ffi.CallMethod[DirectionalEdgeInsets](c_, "contentInsets")
	return rv
}

func (c_ CollectionLayoutSection) SetContentInsets(value DirectionalEdgeInsets) {
	ffi.CallMethod[ffi.Void](c_, "setContentInsets:", value)
}

func (c_ CollectionLayoutSection) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := ffi.CallMethod[[]CollectionLayoutBoundarySupplementaryItem](c_, "boundarySupplementaryItems")
	return rv
}

func (c_ CollectionLayoutSection) SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem) {
	ffi.CallMethod[ffi.Void](c_, "setBoundarySupplementaryItems:", value)
}

func (c_ CollectionLayoutSection) DecorationItems() []CollectionLayoutDecorationItem {
	rv := ffi.CallMethod[[]CollectionLayoutDecorationItem](c_, "decorationItems")
	return rv
}

func (c_ CollectionLayoutSection) SetDecorationItems(value []ICollectionLayoutDecorationItem) {
	ffi.CallMethod[ffi.Void](c_, "setDecorationItems:", value)
}

func (c_ CollectionLayoutSection) SupplementariesFollowContentInsets() bool {
	rv := ffi.CallMethod[bool](c_, "supplementariesFollowContentInsets")
	return rv
}

func (c_ CollectionLayoutSection) SetSupplementariesFollowContentInsets(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSupplementariesFollowContentInsets:", value)
}
