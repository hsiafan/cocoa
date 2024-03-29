// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[UndoManager](uc, objc.SEL("alloc"))
	return rv
}

func (uc _UndoManagerClass) New() UndoManager {
	rv := objc.CallMethod[UndoManager](uc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewUndoManager() UndoManager {
	return UndoManagerClass.New()
}

func (u_ UndoManager) Init() UndoManager {
	rv := objc.CallMethod[UndoManager](u_, objc.SEL("init"))
	return rv
}

func (u_ UndoManager) RegisterUndoWithTarget_Handler(target objc.IObject, undoHandler func(target objc.Object)) {
	objc.CallMethod[objc.Void](u_, objc.SEL("registerUndoWithTarget:handler:"), objc.ExtractPtr(target), undoHandler)
}

func (u_ UndoManager) RegisterUndoWithTarget_Selector_Object(target objc.IObject, selector objc.Selector, anObject objc.IObject) {
	objc.CallMethod[objc.Void](u_, objc.SEL("registerUndoWithTarget:selector:object:"), objc.ExtractPtr(target), selector, objc.ExtractPtr(anObject))
}

func (u_ UndoManager) PrepareWithInvocationTarget(target objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](u_, objc.SEL("prepareWithInvocationTarget:"), objc.ExtractPtr(target))
	return rv
}

func (u_ UndoManager) Undo() {
	objc.CallMethod[objc.Void](u_, objc.SEL("undo"))
}

func (u_ UndoManager) UndoNestedGroup() {
	objc.CallMethod[objc.Void](u_, objc.SEL("undoNestedGroup"))
}

func (u_ UndoManager) Redo() {
	objc.CallMethod[objc.Void](u_, objc.SEL("redo"))
}

func (u_ UndoManager) BeginUndoGrouping() {
	objc.CallMethod[objc.Void](u_, objc.SEL("beginUndoGrouping"))
}

func (u_ UndoManager) EndUndoGrouping() {
	objc.CallMethod[objc.Void](u_, objc.SEL("endUndoGrouping"))
}

func (u_ UndoManager) DisableUndoRegistration() {
	objc.CallMethod[objc.Void](u_, objc.SEL("disableUndoRegistration"))
}

func (u_ UndoManager) EnableUndoRegistration() {
	objc.CallMethod[objc.Void](u_, objc.SEL("enableUndoRegistration"))
}

func (u_ UndoManager) RemoveAllActions() {
	objc.CallMethod[objc.Void](u_, objc.SEL("removeAllActions"))
}

func (u_ UndoManager) RemoveAllActionsWithTarget(target objc.IObject) {
	objc.CallMethod[objc.Void](u_, objc.SEL("removeAllActionsWithTarget:"), objc.ExtractPtr(target))
}

func (u_ UndoManager) SetActionName(actionName string) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setActionName:"), actionName)
}

func (u_ UndoManager) UndoMenuTitleForUndoActionName(actionName string) string {
	rv := objc.CallMethod[string](u_, objc.SEL("undoMenuTitleForUndoActionName:"), actionName)
	return rv
}

func (u_ UndoManager) RedoMenuTitleForUndoActionName(actionName string) string {
	rv := objc.CallMethod[string](u_, objc.SEL("redoMenuTitleForUndoActionName:"), actionName)
	return rv
}

func (u_ UndoManager) SetActionIsDiscardable(discardable bool) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setActionIsDiscardable:"), discardable)
}

func (u_ UndoManager) CanUndo() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("canUndo"))
	return rv
}

func (u_ UndoManager) CanRedo() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("canRedo"))
	return rv
}

func (u_ UndoManager) LevelsOfUndo() uint {
	rv := objc.CallMethod[uint](u_, objc.SEL("levelsOfUndo"))
	return rv
}

func (u_ UndoManager) SetLevelsOfUndo(value uint) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setLevelsOfUndo:"), value)
}

func (u_ UndoManager) GroupsByEvent() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("groupsByEvent"))
	return rv
}

func (u_ UndoManager) SetGroupsByEvent(value bool) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setGroupsByEvent:"), value)
}

func (u_ UndoManager) GroupingLevel() int {
	rv := objc.CallMethod[int](u_, objc.SEL("groupingLevel"))
	return rv
}

func (u_ UndoManager) IsUndoRegistrationEnabled() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("isUndoRegistrationEnabled"))
	return rv
}

func (u_ UndoManager) IsUndoing() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("isUndoing"))
	return rv
}

func (u_ UndoManager) IsRedoing() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("isRedoing"))
	return rv
}

func (u_ UndoManager) UndoActionName() string {
	rv := objc.CallMethod[string](u_, objc.SEL("undoActionName"))
	return rv
}

func (u_ UndoManager) RedoActionName() string {
	rv := objc.CallMethod[string](u_, objc.SEL("redoActionName"))
	return rv
}

func (u_ UndoManager) UndoMenuItemTitle() string {
	rv := objc.CallMethod[string](u_, objc.SEL("undoMenuItemTitle"))
	return rv
}

func (u_ UndoManager) RedoMenuItemTitle() string {
	rv := objc.CallMethod[string](u_, objc.SEL("redoMenuItemTitle"))
	return rv
}

func (u_ UndoManager) RunLoopModes() []RunLoopMode {
	rv := objc.CallMethod[[]RunLoopMode](u_, objc.SEL("runLoopModes"))
	return rv
}

func (u_ UndoManager) SetRunLoopModes(value []RunLoopMode) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setRunLoopModes:"), value)
}

func (u_ UndoManager) UndoActionIsDiscardable() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("undoActionIsDiscardable"))
	return rv
}

func (u_ UndoManager) RedoActionIsDiscardable() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("redoActionIsDiscardable"))
	return rv
}
