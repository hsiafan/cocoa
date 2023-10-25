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
