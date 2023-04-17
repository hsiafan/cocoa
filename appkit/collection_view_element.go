// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type CollectionViewElement interface {
	UserInterfaceItemIdentification
	ImplementsPrepareForReuse() bool
	// optional
	PrepareForReuse()
	ImplementsPreferredLayoutAttributesFittingAttributes() bool
	// optional
	PreferredLayoutAttributesFittingAttributes(layoutAttributes CollectionViewLayoutAttributes) ICollectionViewLayoutAttributes
	ImplementsApplyLayoutAttributes() bool
	// optional
	ApplyLayoutAttributes(layoutAttributes CollectionViewLayoutAttributes)
	ImplementsWillTransitionFromLayout_ToLayout() bool
	// optional
	WillTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
	ImplementsDidTransitionFromLayout_ToLayout() bool
	// optional
	DidTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout)
}

type CollectionViewElementWrapper struct {
	UserInterfaceItemIdentificationWrapper
}

func (c_ *CollectionViewElementWrapper) ImplementsPrepareForReuse() bool {
	return c_.RespondsToSelector(objc.GetSelector("prepareForReuse"))
}

func (c_ CollectionViewElementWrapper) PrepareForReuse() {
	ffi.CallMethod[ffi.Void](c_, "prepareForReuse")
}

func (c_ *CollectionViewElementWrapper) ImplementsPreferredLayoutAttributesFittingAttributes() bool {
	return c_.RespondsToSelector(objc.GetSelector("preferredLayoutAttributesFittingAttributes:"))
}

func (c_ CollectionViewElementWrapper) PreferredLayoutAttributesFittingAttributes(layoutAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutAttributes {
	rv := ffi.CallMethod[CollectionViewLayoutAttributes](c_, "preferredLayoutAttributesFittingAttributes:", layoutAttributes)
	return rv
}

func (c_ *CollectionViewElementWrapper) ImplementsApplyLayoutAttributes() bool {
	return c_.RespondsToSelector(objc.GetSelector("applyLayoutAttributes:"))
}

func (c_ CollectionViewElementWrapper) ApplyLayoutAttributes(layoutAttributes ICollectionViewLayoutAttributes) {
	ffi.CallMethod[ffi.Void](c_, "applyLayoutAttributes:", layoutAttributes)
}

func (c_ *CollectionViewElementWrapper) ImplementsWillTransitionFromLayout_ToLayout() bool {
	return c_.RespondsToSelector(objc.GetSelector("willTransitionFromLayout:toLayout:"))
}

func (c_ CollectionViewElementWrapper) WillTransitionFromLayout_ToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	ffi.CallMethod[ffi.Void](c_, "willTransitionFromLayout:toLayout:", oldLayout, newLayout)
}

func (c_ *CollectionViewElementWrapper) ImplementsDidTransitionFromLayout_ToLayout() bool {
	return c_.RespondsToSelector(objc.GetSelector("didTransitionFromLayout:toLayout:"))
}

func (c_ CollectionViewElementWrapper) DidTransitionFromLayout_ToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	ffi.CallMethod[ffi.Void](c_, "didTransitionFromLayout:toLayout:", oldLayout, newLayout)
}
