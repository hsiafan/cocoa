// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var UndoManagerClass = _UndoManagerClass{objc.GetClass("NSUndoManager")}

type _UndoManagerClass struct {
	objc.Class
}

type IUndoManager interface {
	objc.IObject
	RegisterUndoWithTarget_Handler(target objc.IObject, undoHandler func(target objc.Object))
	RegisterUndoWithTarget_Selector_Object(target objc.IObject, selector objc.Selector, anObject objc.IObject)
	PrepareWithInvocationTarget(target objc.IObject) objc.Object
	Undo()
	UndoNestedGroup()
	Redo()
	BeginUndoGrouping()
	EndUndoGrouping()
	DisableUndoRegistration()
	EnableUndoRegistration()
	RemoveAllActions()
	RemoveAllActionsWithTarget(target objc.IObject)
	SetActionName(actionName string)
	UndoMenuTitleForUndoActionName(actionName string) string
	RedoMenuTitleForUndoActionName(actionName string) string
	SetActionIsDiscardable(discardable bool)
	CanUndo() bool
	CanRedo() bool
	LevelsOfUndo() uint
	SetLevelsOfUndo(value uint)
	GroupsByEvent() bool
	SetGroupsByEvent(value bool)
	GroupingLevel() int
	IsUndoRegistrationEnabled() bool
	IsUndoing() bool
	IsRedoing() bool
	UndoActionName() string
	RedoActionName() string
	UndoMenuItemTitle() string
	RedoMenuItemTitle() string
	RunLoopModes() []RunLoopMode
	SetRunLoopModes(value []RunLoopMode)
	UndoActionIsDiscardable() bool
	RedoActionIsDiscardable() bool
}

type UndoManager struct {
	objc.Object
}

func MakeUndoManager(ptr unsafe.Pointer) UndoManager {
	return UndoManager{
		Object: objc.MakeObject(ptr),
	}
}

func (uc _UndoManagerClass) Alloc() UndoManager {
	rv := ffi.CallMethod[UndoManager](uc, "alloc")
	return rv
}

func (uc _UndoManagerClass) New() UndoManager {
	rv := ffi.CallMethod[UndoManager](uc, "new")
	rv.Autorelease()
	return rv
}

func NewUndoManager() UndoManager {
	return UndoManagerClass.New()
}

func (u_ UndoManager) Init() UndoManager {
	rv := ffi.CallMethod[UndoManager](u_, "init")
	return rv
}

func (u_ UndoManager) RegisterUndoWithTarget_Handler(target objc.IObject, undoHandler func(target objc.Object)) {
	ffi.CallMethod[ffi.Void](u_, "registerUndoWithTarget:handler:", target, undoHandler)
}

func (u_ UndoManager) RegisterUndoWithTarget_Selector_Object(target objc.IObject, selector objc.Selector, anObject objc.IObject) {
	ffi.CallMethod[ffi.Void](u_, "registerUndoWithTarget:selector:object:", target, selector, anObject)
}

func (u_ UndoManager) PrepareWithInvocationTarget(target objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](u_, "prepareWithInvocationTarget:", target)
	return rv
}

func (u_ UndoManager) Undo() {
	ffi.CallMethod[ffi.Void](u_, "undo")
}

func (u_ UndoManager) UndoNestedGroup() {
	ffi.CallMethod[ffi.Void](u_, "undoNestedGroup")
}

func (u_ UndoManager) Redo() {
	ffi.CallMethod[ffi.Void](u_, "redo")
}

func (u_ UndoManager) BeginUndoGrouping() {
	ffi.CallMethod[ffi.Void](u_, "beginUndoGrouping")
}

func (u_ UndoManager) EndUndoGrouping() {
	ffi.CallMethod[ffi.Void](u_, "endUndoGrouping")
}

func (u_ UndoManager) DisableUndoRegistration() {
	ffi.CallMethod[ffi.Void](u_, "disableUndoRegistration")
}

func (u_ UndoManager) EnableUndoRegistration() {
	ffi.CallMethod[ffi.Void](u_, "enableUndoRegistration")
}

func (u_ UndoManager) RemoveAllActions() {
	ffi.CallMethod[ffi.Void](u_, "removeAllActions")
}

func (u_ UndoManager) RemoveAllActionsWithTarget(target objc.IObject) {
	ffi.CallMethod[ffi.Void](u_, "removeAllActionsWithTarget:", target)
}

func (u_ UndoManager) SetActionName(actionName string) {
	ffi.CallMethod[ffi.Void](u_, "setActionName:", actionName)
}

func (u_ UndoManager) UndoMenuTitleForUndoActionName(actionName string) string {
	rv := ffi.CallMethod[string](u_, "undoMenuTitleForUndoActionName:", actionName)
	return rv
}

func (u_ UndoManager) RedoMenuTitleForUndoActionName(actionName string) string {
	rv := ffi.CallMethod[string](u_, "redoMenuTitleForUndoActionName:", actionName)
	return rv
}

func (u_ UndoManager) SetActionIsDiscardable(discardable bool) {
	ffi.CallMethod[ffi.Void](u_, "setActionIsDiscardable:", discardable)
}

func (u_ UndoManager) CanUndo() bool {
	rv := ffi.CallMethod[bool](u_, "canUndo")
	return rv
}

func (u_ UndoManager) CanRedo() bool {
	rv := ffi.CallMethod[bool](u_, "canRedo")
	return rv
}

func (u_ UndoManager) LevelsOfUndo() uint {
	rv := ffi.CallMethod[uint](u_, "levelsOfUndo")
	return rv
}

func (u_ UndoManager) SetLevelsOfUndo(value uint) {
	ffi.CallMethod[ffi.Void](u_, "setLevelsOfUndo:", value)
}

func (u_ UndoManager) GroupsByEvent() bool {
	rv := ffi.CallMethod[bool](u_, "groupsByEvent")
	return rv
}

func (u_ UndoManager) SetGroupsByEvent(value bool) {
	ffi.CallMethod[ffi.Void](u_, "setGroupsByEvent:", value)
}

func (u_ UndoManager) GroupingLevel() int {
	rv := ffi.CallMethod[int](u_, "groupingLevel")
	return rv
}

func (u_ UndoManager) IsUndoRegistrationEnabled() bool {
	rv := ffi.CallMethod[bool](u_, "isUndoRegistrationEnabled")
	return rv
}

func (u_ UndoManager) IsUndoing() bool {
	rv := ffi.CallMethod[bool](u_, "isUndoing")
	return rv
}

func (u_ UndoManager) IsRedoing() bool {
	rv := ffi.CallMethod[bool](u_, "isRedoing")
	return rv
}

func (u_ UndoManager) UndoActionName() string {
	rv := ffi.CallMethod[string](u_, "undoActionName")
	return rv
}

func (u_ UndoManager) RedoActionName() string {
	rv := ffi.CallMethod[string](u_, "redoActionName")
	return rv
}

func (u_ UndoManager) UndoMenuItemTitle() string {
	rv := ffi.CallMethod[string](u_, "undoMenuItemTitle")
	return rv
}

func (u_ UndoManager) RedoMenuItemTitle() string {
	rv := ffi.CallMethod[string](u_, "redoMenuItemTitle")
	return rv
}

func (u_ UndoManager) RunLoopModes() []RunLoopMode {
	rv := ffi.CallMethod[[]RunLoopMode](u_, "runLoopModes")
	return rv
}

func (u_ UndoManager) SetRunLoopModes(value []RunLoopMode) {
	ffi.CallMethod[ffi.Void](u_, "setRunLoopModes:", value)
}

func (u_ UndoManager) UndoActionIsDiscardable() bool {
	rv := ffi.CallMethod[bool](u_, "undoActionIsDiscardable")
	return rv
}

func (u_ UndoManager) RedoActionIsDiscardable() bool {
	rv := ffi.CallMethod[bool](u_, "redoActionIsDiscardable")
	return rv
}
