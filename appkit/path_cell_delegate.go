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

type PathCellDelegateCreator struct {
	className string
	class     objc.Class
}

func NewPathCellDelegateCreator(name string) *PathCellDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &PathCellDelegateCreator{className: name, class: class}
}

func (c *PathCellDelegateCreator) SetPathCell_WillDisplayOpenPanel(handle func(o objc.Object, pathCell PathCell, openPanel OpenPanel)) {
	objc.AddMethod(c.class, objc.GetSelector("pathCell:willDisplayOpenPanel:"), handle)
}

func (c *PathCellDelegateCreator) SetPathCell_WillPopUpMenu(handle func(o objc.Object, pathCell PathCell, menu Menu)) {
	objc.AddMethod(c.class, objc.GetSelector("pathCell:willPopUpMenu:"), handle)
}

func (c *PathCellDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
