// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[foundation.Size](c_, "contentSize")
	return rv
}

func (c_ *CollectionLayoutContainerWrapper) ImplementsEffectiveContentSize() bool {
	return c_.RespondsToSelector(objc.GetSelector("effectiveContentSize"))
}

func (c_ CollectionLayoutContainerWrapper) EffectiveContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "effectiveContentSize")
	return rv
}

func (c_ *CollectionLayoutContainerWrapper) ImplementsContentInsets() bool {
	return c_.RespondsToSelector(objc.GetSelector("contentInsets"))
}

func (c_ CollectionLayoutContainerWrapper) ContentInsets() DirectionalEdgeInsets {
	rv := ffi.CallMethod[DirectionalEdgeInsets](c_, "contentInsets")
	return rv
}

func (c_ *CollectionLayoutContainerWrapper) ImplementsEffectiveContentInsets() bool {
	return c_.RespondsToSelector(objc.GetSelector("effectiveContentInsets"))
}

func (c_ CollectionLayoutContainerWrapper) EffectiveContentInsets() DirectionalEdgeInsets {
	rv := ffi.CallMethod[DirectionalEdgeInsets](c_, "effectiveContentInsets")
	return rv
}
