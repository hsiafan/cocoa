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
