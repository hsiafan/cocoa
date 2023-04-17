// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type PathControlDelegateImpl struct {
	_PathControl_ShouldDragPathComponentCell_WithPasteboard func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	_PathControl_ValidateDrop                               func(pathControl PathControl, info DraggingInfoWrapper) DragOperation
	_PathControl_AcceptDrop                                 func(pathControl PathControl, info DraggingInfoWrapper) bool
	_PathControl_WillDisplayOpenPanel                       func(pathControl PathControl, openPanel OpenPanel)
	_PathControl_WillPopUpMenu                              func(pathControl PathControl, menu Menu)
	_PathControl_ShouldDragItem_WithPasteboard              func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
}

func (di *PathControlDelegateImpl) ImplementsPathControl_ShouldDragPathComponentCell_WithPasteboard() bool {
	return di._PathControl_ShouldDragPathComponentCell_WithPasteboard != nil
}

func (di *PathControlDelegateImpl) SetPathControl_ShouldDragPathComponentCell_WithPasteboard(f func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool) {
	di._PathControl_ShouldDragPathComponentCell_WithPasteboard = f
}

func (di *PathControlDelegateImpl) PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool {
	return di._PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl, pathComponentCell, pasteboard)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_ValidateDrop() bool {
	return di._PathControl_ValidateDrop != nil
}

func (di *PathControlDelegateImpl) SetPathControl_ValidateDrop(f func(pathControl PathControl, info DraggingInfoWrapper) DragOperation) {
	di._PathControl_ValidateDrop = f
}

func (di *PathControlDelegateImpl) PathControl_ValidateDrop(pathControl PathControl, info DraggingInfoWrapper) DragOperation {
	return di._PathControl_ValidateDrop(pathControl, info)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_AcceptDrop() bool {
	return di._PathControl_AcceptDrop != nil
}

func (di *PathControlDelegateImpl) SetPathControl_AcceptDrop(f func(pathControl PathControl, info DraggingInfoWrapper) bool) {
	di._PathControl_AcceptDrop = f
}

func (di *PathControlDelegateImpl) PathControl_AcceptDrop(pathControl PathControl, info DraggingInfoWrapper) bool {
	return di._PathControl_AcceptDrop(pathControl, info)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_WillDisplayOpenPanel() bool {
	return di._PathControl_WillDisplayOpenPanel != nil
}

func (di *PathControlDelegateImpl) SetPathControl_WillDisplayOpenPanel(f func(pathControl PathControl, openPanel OpenPanel)) {
	di._PathControl_WillDisplayOpenPanel = f
}

func (di *PathControlDelegateImpl) PathControl_WillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel) {
	di._PathControl_WillDisplayOpenPanel(pathControl, openPanel)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_WillPopUpMenu() bool {
	return di._PathControl_WillPopUpMenu != nil
}

func (di *PathControlDelegateImpl) SetPathControl_WillPopUpMenu(f func(pathControl PathControl, menu Menu)) {
	di._PathControl_WillPopUpMenu = f
}

func (di *PathControlDelegateImpl) PathControl_WillPopUpMenu(pathControl PathControl, menu Menu) {
	di._PathControl_WillPopUpMenu(pathControl, menu)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_ShouldDragItem_WithPasteboard() bool {
	return di._PathControl_ShouldDragItem_WithPasteboard != nil
}

func (di *PathControlDelegateImpl) SetPathControl_ShouldDragItem_WithPasteboard(f func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool) {
	di._PathControl_ShouldDragItem_WithPasteboard = f
}

func (di *PathControlDelegateImpl) PathControl_ShouldDragItem_WithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool {
	return di._PathControl_ShouldDragItem_WithPasteboard(pathControl, pathItem, pasteboard)
}

type PathControlDelegateWrapper struct {
	objc.Object
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_ShouldDragPathComponentCell_WithPasteboard() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:shouldDragPathComponentCell:withPasteboard:"))
}

func (p_ PathControlDelegateWrapper) PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl IPathControl, pathComponentCell IPathComponentCell, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](p_, "pathControl:shouldDragPathComponentCell:withPasteboard:", pathControl, pathComponentCell, pasteboard)
	return rv
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_ValidateDrop() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:validateDrop:"))
}

func (p_ PathControlDelegateWrapper) PathControl_ValidateDrop(pathControl IPathControl, info DraggingInfo) DragOperation {
	po := ffi.CreateProtocol("NSDraggingInfo", info)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](p_, "pathControl:validateDrop:", pathControl, po)
	return rv
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_AcceptDrop() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:acceptDrop:"))
}

func (p_ PathControlDelegateWrapper) PathControl_AcceptDrop(pathControl IPathControl, info DraggingInfo) bool {
	po := ffi.CreateProtocol("NSDraggingInfo", info)
	defer po.Release()
	rv := ffi.CallMethod[bool](p_, "pathControl:acceptDrop:", pathControl, po)
	return rv
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_WillDisplayOpenPanel() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:willDisplayOpenPanel:"))
}

func (p_ PathControlDelegateWrapper) PathControl_WillDisplayOpenPanel(pathControl IPathControl, openPanel IOpenPanel) {
	ffi.CallMethod[ffi.Void](p_, "pathControl:willDisplayOpenPanel:", pathControl, openPanel)
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_WillPopUpMenu() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:willPopUpMenu:"))
}

func (p_ PathControlDelegateWrapper) PathControl_WillPopUpMenu(pathControl IPathControl, menu IMenu) {
	ffi.CallMethod[ffi.Void](p_, "pathControl:willPopUpMenu:", pathControl, menu)
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_ShouldDragItem_WithPasteboard() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:shouldDragItem:withPasteboard:"))
}

func (p_ PathControlDelegateWrapper) PathControl_ShouldDragItem_WithPasteboard(pathControl IPathControl, pathItem IPathControlItem, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](p_, "pathControl:shouldDragItem:withPasteboard:", pathControl, pathItem, pasteboard)
	return rv
}
