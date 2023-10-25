// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionLayoutSection](cc, objc.SEL("sectionWithGroup:"), objc.ExtractPtr(group))
	return rv
}

func (cc _CollectionLayoutSectionClass) Alloc() CollectionLayoutSection {
	rv := objc.CallMethod[CollectionLayoutSection](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionLayoutSectionClass) New() CollectionLayoutSection {
	rv := objc.CallMethod[CollectionLayoutSection](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionLayoutSection() CollectionLayoutSection {
	return CollectionLayoutSectionClass.New()
}

func (c_ CollectionLayoutSection) Init() CollectionLayoutSection {
	rv := objc.CallMethod[CollectionLayoutSection](c_, objc.SEL("init"))
	return rv
}

func (c_ CollectionLayoutSection) OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior {
	rv := objc.CallMethod[CollectionLayoutSectionOrthogonalScrollingBehavior](c_, objc.SEL("orthogonalScrollingBehavior"))
	return rv
}

func (c_ CollectionLayoutSection) SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setOrthogonalScrollingBehavior:"), value)
}

func (c_ CollectionLayoutSection) InterGroupSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.SEL("interGroupSpacing"))
	return rv
}

func (c_ CollectionLayoutSection) SetInterGroupSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setInterGroupSpacing:"), value)
}

func (c_ CollectionLayoutSection) ContentInsets() DirectionalEdgeInsets {
	rv := objc.CallMethod[DirectionalEdgeInsets](c_, objc.SEL("contentInsets"))
	return rv
}

func (c_ CollectionLayoutSection) SetContentInsets(value DirectionalEdgeInsets) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setContentInsets:"), value)
}

func (c_ CollectionLayoutSection) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	rv := objc.CallMethod[[]CollectionLayoutBoundarySupplementaryItem](c_, objc.SEL("boundarySupplementaryItems"))
	return rv
}

func (c_ CollectionLayoutSection) SetBoundarySupplementaryItems(value []ICollectionLayoutBoundarySupplementaryItem) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setBoundarySupplementaryItems:"), value)
}

func (c_ CollectionLayoutSection) DecorationItems() []CollectionLayoutDecorationItem {
	rv := objc.CallMethod[[]CollectionLayoutDecorationItem](c_, objc.SEL("decorationItems"))
	return rv
}

func (c_ CollectionLayoutSection) SetDecorationItems(value []ICollectionLayoutDecorationItem) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setDecorationItems:"), value)
}

func (c_ CollectionLayoutSection) SupplementariesFollowContentInsets() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("supplementariesFollowContentInsets"))
	return rv
}

func (c_ CollectionLayoutSection) SetSupplementariesFollowContentInsets(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setSupplementariesFollowContentInsets:"), value)
}
