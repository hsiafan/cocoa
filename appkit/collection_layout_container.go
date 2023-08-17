// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type CollectionLayoutContainer interface {
	ImplementsContentSize() bool
	// optional
	ContentSize() foundation.Size
	ImplementsEffectiveContentSize() bool
	// optional
	EffectiveContentSize() foundation.Size
	ImplementsContentInsets() bool
	// optional
	ContentInsets() DirectionalEdgeInsets
	ImplementsEffectiveContentInsets() bool
	// optional
	EffectiveContentInsets() DirectionalEdgeInsets
}

func WrapCollectionLayoutContainer(v CollectionLayoutContainer) objc.Object {
	return objc.WrapAsProtocol("NSCollectionLayoutContainer", v)
}

type CollectionLayoutContainerBase struct {
}

func (p *CollectionLayoutContainerBase) ImplementsContentSize() bool {
	return false
}

func (p *CollectionLayoutContainerBase) ContentSize() foundation.Size {
	panic("unimpemented protocol method")
}

func (p *CollectionLayoutContainerBase) ImplementsEffectiveContentSize() bool {
	return false
}

func (p *CollectionLayoutContainerBase) EffectiveContentSize() foundation.Size {
	panic("unimpemented protocol method")
}

func (p *CollectionLayoutContainerBase) ImplementsContentInsets() bool {
	return false
}

func (p *CollectionLayoutContainerBase) ContentInsets() DirectionalEdgeInsets {
	panic("unimpemented protocol method")
}

func (p *CollectionLayoutContainerBase) ImplementsEffectiveContentInsets() bool {
	return false
}

func (p *CollectionLayoutContainerBase) EffectiveContentInsets() DirectionalEdgeInsets {
	panic("unimpemented protocol method")
}

type CollectionLayoutContainerWrapper struct {
	objc.Object
}

func (c_ CollectionLayoutContainerWrapper) ContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("contentSize"))
	return rv
}

func (c_ CollectionLayoutContainerWrapper) EffectiveContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("effectiveContentSize"))
	return rv
}

func (c_ CollectionLayoutContainerWrapper) ContentInsets() DirectionalEdgeInsets {
	rv := objc.CallMethod[DirectionalEdgeInsets](c_, objc.GetSelector("contentInsets"))
	return rv
}

func (c_ CollectionLayoutContainerWrapper) EffectiveContentInsets() DirectionalEdgeInsets {
	rv := objc.CallMethod[DirectionalEdgeInsets](c_, objc.GetSelector("effectiveContentInsets"))
	return rv
}
