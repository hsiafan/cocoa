// auto-generated code, do not modify
package appkit

import (
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

func WrapCollectionViewElement(v CollectionViewElement) objc.Object {
	return objc.WrapAsProtocol("NSCollectionViewElement", v)
}

type CollectionViewElementBase struct {
	UserInterfaceItemIdentificationBase
}

func (p *CollectionViewElementBase) ImplementsPrepareForReuse() bool {
	return false
}

func (p *CollectionViewElementBase) PrepareForReuse() {
	panic("unimpemented protocol method")
}

func (p *CollectionViewElementBase) ImplementsPreferredLayoutAttributesFittingAttributes() bool {
	return false
}

func (p *CollectionViewElementBase) PreferredLayoutAttributesFittingAttributes(layoutAttributes CollectionViewLayoutAttributes) ICollectionViewLayoutAttributes {
	panic("unimpemented protocol method")
}

func (p *CollectionViewElementBase) ImplementsApplyLayoutAttributes() bool {
	return false
}

func (p *CollectionViewElementBase) ApplyLayoutAttributes(layoutAttributes CollectionViewLayoutAttributes) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewElementBase) ImplementsWillTransitionFromLayout_ToLayout() bool {
	return false
}

func (p *CollectionViewElementBase) WillTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewElementBase) ImplementsDidTransitionFromLayout_ToLayout() bool {
	return false
}

func (p *CollectionViewElementBase) DidTransitionFromLayout_ToLayout(oldLayout CollectionViewLayout, newLayout CollectionViewLayout) {
	panic("unimpemented protocol method")
}

type CollectionViewElementWrapper struct {
	objc.Object
}

func (c_ CollectionViewElementWrapper) PrepareForReuse() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("prepareForReuse"))
}

func (c_ CollectionViewElementWrapper) PreferredLayoutAttributesFittingAttributes(layoutAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("preferredLayoutAttributesFittingAttributes:"), objc.ExtractPtr(layoutAttributes))
	return rv
}

func (c_ CollectionViewElementWrapper) ApplyLayoutAttributes(layoutAttributes ICollectionViewLayoutAttributes) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("applyLayoutAttributes:"), objc.ExtractPtr(layoutAttributes))
}

func (c_ CollectionViewElementWrapper) WillTransitionFromLayout_ToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("willTransitionFromLayout:toLayout:"), objc.ExtractPtr(oldLayout), objc.ExtractPtr(newLayout))
}

func (c_ CollectionViewElementWrapper) DidTransitionFromLayout_ToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("didTransitionFromLayout:toLayout:"), objc.ExtractPtr(oldLayout), objc.ExtractPtr(newLayout))
}

func (c_ CollectionViewElementWrapper) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setIdentifier:"), value)
}

func (c_ CollectionViewElementWrapper) Identifier() UserInterfaceItemIdentifier {
	rv := objc.CallMethod[UserInterfaceItemIdentifier](c_, objc.GetSelector("identifier"))
	return rv
}
