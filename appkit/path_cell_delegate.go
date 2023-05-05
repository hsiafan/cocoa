// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type PathCellDelegate interface {
	ImplementsPathCell_WillDisplayOpenPanel() bool
	// optional
	PathCell_WillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel)
	ImplementsPathCell_WillPopUpMenu() bool
	// optional
	PathCell_WillPopUpMenu(pathCell PathCell, menu Menu)
}

type PathCellDelegateImpl struct {
	_PathCell_WillDisplayOpenPanel func(pathCell PathCell, openPanel OpenPanel)
	_PathCell_WillPopUpMenu        func(pathCell PathCell, menu Menu)
}

func (di *PathCellDelegateImpl) ImplementsPathCell_WillDisplayOpenPanel() bool {
	return di._PathCell_WillDisplayOpenPanel != nil
}

func (di *PathCellDelegateImpl) SetPathCell_WillDisplayOpenPanel(f func(pathCell PathCell, openPanel OpenPanel)) {
	di._PathCell_WillDisplayOpenPanel = f
}

func (di *PathCellDelegateImpl) PathCell_WillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel) {
	di._PathCell_WillDisplayOpenPanel(pathCell, openPanel)
}
func (di *PathCellDelegateImpl) ImplementsPathCell_WillPopUpMenu() bool {
	return di._PathCell_WillPopUpMenu != nil
}

func (di *PathCellDelegateImpl) SetPathCell_WillPopUpMenu(f func(pathCell PathCell, menu Menu)) {
	di._PathCell_WillPopUpMenu = f
}

func (di *PathCellDelegateImpl) PathCell_WillPopUpMenu(pathCell PathCell, menu Menu) {
	di._PathCell_WillPopUpMenu(pathCell, menu)
}

type PathCellDelegateWrapper struct {
	objc.Object
}

func (p_ *PathCellDelegateWrapper) ImplementsPathCell_WillDisplayOpenPanel() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathCell:willDisplayOpenPanel:"))
}

func (p_ PathCellDelegateWrapper) PathCell_WillDisplayOpenPanel(pathCell IPathCell, openPanel IOpenPanel) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathCell:willDisplayOpenPanel:"), objc.ExtractPtr(pathCell), objc.ExtractPtr(openPanel))
}

func (p_ *PathCellDelegateWrapper) ImplementsPathCell_WillPopUpMenu() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathCell:willPopUpMenu:"))
}

func (p_ PathCellDelegateWrapper) PathCell_WillPopUpMenu(pathCell IPathCell, menu IMenu) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathCell:willPopUpMenu:"), objc.ExtractPtr(pathCell), objc.ExtractPtr(menu))
}
