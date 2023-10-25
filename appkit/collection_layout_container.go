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
