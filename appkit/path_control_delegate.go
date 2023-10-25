// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type PathControlDelegate interface {
	ImplementsPathControl_ShouldDragPathComponentCell_WithPasteboard() bool
	// optional
	PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	ImplementsPathControl_ValidateDrop() bool
	// optional
	PathControl_ValidateDrop(pathControl PathControl, info objc.Object) DragOperation
	ImplementsPathControl_AcceptDrop() bool
	// optional
	PathControl_AcceptDrop(pathControl PathControl, info objc.Object) bool
	ImplementsPathControl_WillDisplayOpenPanel() bool
	// optional
	PathControl_WillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel)
	ImplementsPathControl_WillPopUpMenu() bool
	// optional
	PathControl_WillPopUpMenu(pathControl PathControl, menu Menu)
	ImplementsPathControl_ShouldDragItem_WithPasteboard() bool
	// optional
	PathControl_ShouldDragItem_WithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
}

func WrapPathControlDelegate(v PathControlDelegate) objc.Object {
	return objc.WrapAsProtocol("NSPathControlDelegate", v)
}

type PathControlDelegateBase struct {
}

func (p *PathControlDelegateBase) ImplementsPathControl_ShouldDragPathComponentCell_WithPasteboard() bool {
	return false
}

func (p *PathControlDelegateBase) PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

func (p *PathControlDelegateBase) ImplementsPathControl_ValidateDrop() bool {
	return false
}

func (p *PathControlDelegateBase) PathControl_ValidateDrop(pathControl PathControl, info objc.Object) DragOperation {
	panic("unimpemented protocol method")
}

func (p *PathControlDelegateBase) ImplementsPathControl_AcceptDrop() bool {
	return false
}

func (p *PathControlDelegateBase) PathControl_AcceptDrop(pathControl PathControl, info objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *PathControlDelegateBase) ImplementsPathControl_WillDisplayOpenPanel() bool {
	return false
}

func (p *PathControlDelegateBase) PathControl_WillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel) {
	panic("unimpemented protocol method")
}

func (p *PathControlDelegateBase) ImplementsPathControl_WillPopUpMenu() bool {
	return false
}

func (p *PathControlDelegateBase) PathControl_WillPopUpMenu(pathControl PathControl, menu Menu) {
	panic("unimpemented protocol method")
}

func (p *PathControlDelegateBase) ImplementsPathControl_ShouldDragItem_WithPasteboard() bool {
	return false
}

func (p *PathControlDelegateBase) PathControl_ShouldDragItem_WithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool {
	panic("unimpemented protocol method")
}
