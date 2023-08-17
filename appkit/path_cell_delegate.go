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

func WrapPathCellDelegate(v PathCellDelegate) objc.Object {
	return objc.WrapAsProtocol("NSPathCellDelegate", v)
}

type PathCellDelegateBase struct {
}

func (p *PathCellDelegateBase) ImplementsPathCell_WillDisplayOpenPanel() bool {
	return false
}

func (p *PathCellDelegateBase) PathCell_WillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel) {
	panic("unimpemented protocol method")
}

func (p *PathCellDelegateBase) ImplementsPathCell_WillPopUpMenu() bool {
	return false
}

func (p *PathCellDelegateBase) PathCell_WillPopUpMenu(pathCell PathCell, menu Menu) {
	panic("unimpemented protocol method")
}

type PathCellDelegateWrapper struct {
	objc.Object
}

func (p_ PathCellDelegateWrapper) PathCell_WillDisplayOpenPanel(pathCell IPathCell, openPanel IOpenPanel) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathCell:willDisplayOpenPanel:"), objc.ExtractPtr(pathCell), objc.ExtractPtr(openPanel))
}

func (p_ PathCellDelegateWrapper) PathCell_WillPopUpMenu(pathCell IPathCell, menu IMenu) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathCell:willPopUpMenu:"), objc.ExtractPtr(pathCell), objc.ExtractPtr(menu))
}
