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
	PathControl_ValidateDrop(pathControl PathControl, info DraggingInfoWrapper) DragOperation
	ImplementsPathControl_AcceptDrop() bool
	// optional
	PathControl_AcceptDrop(pathControl PathControl, info DraggingInfoWrapper) bool
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

func (p *PathControlDelegateBase) PathControl_ValidateDrop(pathControl PathControl, info DraggingInfoWrapper) DragOperation {
	panic("unimpemented protocol method")
}

func (p *PathControlDelegateBase) ImplementsPathControl_AcceptDrop() bool {
	return false
}

func (p *PathControlDelegateBase) PathControl_AcceptDrop(pathControl PathControl, info DraggingInfoWrapper) bool {
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

type PathControlDelegateWrapper struct {
	objc.Object
}

func (p_ PathControlDelegateWrapper) PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl IPathControl, pathComponentCell IPathComponentCell, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("pathControl:shouldDragPathComponentCell:withPasteboard:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(pathComponentCell), objc.ExtractPtr(pasteboard))
	return rv
}

func (p_ PathControlDelegateWrapper) PathControl_ValidateDrop(pathControl IPathControl, info objc.IObject) DragOperation {
	rv := objc.CallMethod[DragOperation](p_, objc.GetSelector("pathControl:validateDrop:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(info))
	return rv
}

func (p_ PathControlDelegateWrapper) PathControl_AcceptDrop(pathControl IPathControl, info objc.IObject) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("pathControl:acceptDrop:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(info))
	return rv
}

func (p_ PathControlDelegateWrapper) PathControl_WillDisplayOpenPanel(pathControl IPathControl, openPanel IOpenPanel) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathControl:willDisplayOpenPanel:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(openPanel))
}

func (p_ PathControlDelegateWrapper) PathControl_WillPopUpMenu(pathControl IPathControl, menu IMenu) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pathControl:willPopUpMenu:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(menu))
}

func (p_ PathControlDelegateWrapper) PathControl_ShouldDragItem_WithPasteboard(pathControl IPathControl, pathItem IPathControlItem, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("pathControl:shouldDragItem:withPasteboard:"), objc.ExtractPtr(pathControl), objc.ExtractPtr(pathItem), objc.ExtractPtr(pasteboard))
	return rv
}
