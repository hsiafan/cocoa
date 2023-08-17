// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type CollectionViewSectionHeaderView interface {
	CollectionViewElement
	ImplementsSetSectionCollapseButton() bool
	// optional
	SetSectionCollapseButton(value Button)
	ImplementsSectionCollapseButton() bool
	// optional
	SectionCollapseButton() IButton
}

func WrapCollectionViewSectionHeaderView(v CollectionViewSectionHeaderView) objc.Object {
	return objc.WrapAsProtocol("NSCollectionViewSectionHeaderView", v)
}

type CollectionViewSectionHeaderViewBase struct {
	CollectionViewElementBase
}

func (p *CollectionViewSectionHeaderViewBase) ImplementsSetSectionCollapseButton() bool {
	return false
}

func (p *CollectionViewSectionHeaderViewBase) SetSectionCollapseButton(value Button) {
	panic("unimpemented protocol method")
}

func (p *CollectionViewSectionHeaderViewBase) ImplementsSectionCollapseButton() bool {
	return false
}

func (p *CollectionViewSectionHeaderViewBase) SectionCollapseButton() IButton {
	panic("unimpemented protocol method")
}

type CollectionViewSectionHeaderViewWrapper struct {
	objc.Object
}

// weak property
func (c_ CollectionViewSectionHeaderViewWrapper) SetSectionCollapseButton(value IButton) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setSectionCollapseButton:"), objc.ExtractPtr(value))
}

// weak property
func (c_ CollectionViewSectionHeaderViewWrapper) SectionCollapseButton() Button {
	rv := objc.CallMethod[Button](c_, objc.GetSelector("sectionCollapseButton"))
	return rv
}

func (c_ CollectionViewSectionHeaderViewWrapper) PrepareForReuse() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("prepareForReuse"))
}

func (c_ CollectionViewSectionHeaderViewWrapper) PreferredLayoutAttributesFittingAttributes(layoutAttributes ICollectionViewLayoutAttributes) CollectionViewLayoutAttributes {
	rv := objc.CallMethod[CollectionViewLayoutAttributes](c_, objc.GetSelector("preferredLayoutAttributesFittingAttributes:"), objc.ExtractPtr(layoutAttributes))
	return rv
}

func (c_ CollectionViewSectionHeaderViewWrapper) ApplyLayoutAttributes(layoutAttributes ICollectionViewLayoutAttributes) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("applyLayoutAttributes:"), objc.ExtractPtr(layoutAttributes))
}

func (c_ CollectionViewSectionHeaderViewWrapper) WillTransitionFromLayout_ToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("willTransitionFromLayout:toLayout:"), objc.ExtractPtr(oldLayout), objc.ExtractPtr(newLayout))
}

func (c_ CollectionViewSectionHeaderViewWrapper) DidTransitionFromLayout_ToLayout(oldLayout ICollectionViewLayout, newLayout ICollectionViewLayout) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("didTransitionFromLayout:toLayout:"), objc.ExtractPtr(oldLayout), objc.ExtractPtr(newLayout))
}

func (c_ CollectionViewSectionHeaderViewWrapper) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setIdentifier:"), value)
}

func (c_ CollectionViewSectionHeaderViewWrapper) Identifier() UserInterfaceItemIdentifier {
	rv := objc.CallMethod[UserInterfaceItemIdentifier](c_, objc.GetSelector("identifier"))
	return rv
}
