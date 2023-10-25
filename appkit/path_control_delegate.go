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

type PathControlDelegateCreator struct {
	className string
	class     objc.Class
}

func NewPathControlDelegateCreator(name string) *PathControlDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &PathControlDelegateCreator{className: name, class: class}
}

func (c *PathControlDelegateCreator) SetPathControl_ShouldDragPathComponentCell_WithPasteboard(handle func(o objc.ProtocolBase, pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool) {
	objc.AddMethod(c.class, objc.SEL("pathControl:shouldDragPathComponentCell:withPasteboard:"), handle)
}

func (c *PathControlDelegateCreator) SetPathControl_ValidateDrop(handle func(o objc.ProtocolBase, pathControl PathControl, info objc.Object) DragOperation) {
	objc.AddMethod(c.class, objc.SEL("pathControl:validateDrop:"), handle)
}

func (c *PathControlDelegateCreator) SetPathControl_AcceptDrop(handle func(o objc.ProtocolBase, pathControl PathControl, info objc.Object) bool) {
	objc.AddMethod(c.class, objc.SEL("pathControl:acceptDrop:"), handle)
}

func (c *PathControlDelegateCreator) SetPathControl_WillDisplayOpenPanel(handle func(o objc.ProtocolBase, pathControl PathControl, openPanel OpenPanel)) {
	objc.AddMethod(c.class, objc.SEL("pathControl:willDisplayOpenPanel:"), handle)
}

func (c *PathControlDelegateCreator) SetPathControl_WillPopUpMenu(handle func(o objc.ProtocolBase, pathControl PathControl, menu Menu)) {
	objc.AddMethod(c.class, objc.SEL("pathControl:willPopUpMenu:"), handle)
}

func (c *PathControlDelegateCreator) SetPathControl_ShouldDragItem_WithPasteboard(handle func(o objc.ProtocolBase, pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool) {
	objc.AddMethod(c.class, objc.SEL("pathControl:shouldDragItem:withPasteboard:"), handle)
}

func (c *PathControlDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
