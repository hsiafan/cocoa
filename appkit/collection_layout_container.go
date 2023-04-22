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

type CollectionLayoutContainerWrapper struct {
	objc.Object
}

func (c_ *CollectionLayoutContainerWrapper) ImplementsContentSize() bool {
	return c_.RespondsToSelector(objc.GetSelector("contentSize"))
}

func (c_ CollectionLayoutContainerWrapper) ContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("contentSize"))
	return rv
}

func (c_ *CollectionLayoutContainerWrapper) ImplementsEffectiveContentSize() bool {
	return c_.RespondsToSelector(objc.GetSelector("effectiveContentSize"))
}

func (c_ CollectionLayoutContainerWrapper) EffectiveContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("effectiveContentSize"))
	return rv
}

func (c_ *CollectionLayoutContainerWrapper) ImplementsContentInsets() bool {
	return c_.RespondsToSelector(objc.GetSelector("contentInsets"))
}

func (c_ CollectionLayoutContainerWrapper) ContentInsets() DirectionalEdgeInsets {
	rv := objc.CallMethod[DirectionalEdgeInsets](c_, objc.GetSelector("contentInsets"))
	return rv
}

func (c_ *CollectionLayoutContainerWrapper) ImplementsEffectiveContentInsets() bool {
	return c_.RespondsToSelector(objc.GetSelector("effectiveContentInsets"))
}

func (c_ CollectionLayoutContainerWrapper) EffectiveContentInsets() DirectionalEdgeInsets {
	rv := objc.CallMethod[DirectionalEdgeInsets](c_, objc.GetSelector("effectiveContentInsets"))
	return rv
}
