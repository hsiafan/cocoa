// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ControllerClass = _ControllerClass{objc.GetClass("NSController")}

type _ControllerClass struct {
	objc.Class
}

type IController interface {
	objc.IObject
	ObjectDidBeginEditing(editor Editor)
	ObjectDidEndEditing(editor Editor)
	CommitEditing() bool
	CommitEditingWithDelegate_DidCommitSelector_ContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer)
	DiscardEditing()
	IsEditing() bool
}

type Controller struct {
	objc.Object
}

func MakeController(ptr unsafe.Pointer) Controller {
	return Controller{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ Controller) Init() Controller {
	rv := ffi.CallMethod[Controller](c_, "init")
	rv.Autorelease()
	return rv
}

func (cc _ControllerClass) Alloc() Controller {
	rv := ffi.CallMethod[Controller](cc, "alloc")
	return rv
}

func (cc _ControllerClass) New() Controller {
	rv := ffi.CallMethod[Controller](cc, "new")
	rv.Autorelease()
	return rv
}

func NewController() Controller {
	return ControllerClass.New()
}

func (c_ Controller) ObjectDidBeginEditing(editor Editor) {
	po := ffi.CreateProtocol(editor)
	defer po.Release()
	ffi.CallMethod[ffi.Void](c_, "objectDidBeginEditing:", po)
}

func (c_ Controller) ObjectDidEndEditing(editor Editor) {
	po := ffi.CreateProtocol(editor)
	defer po.Release()
	ffi.CallMethod[ffi.Void](c_, "objectDidEndEditing:", po)
}

func (c_ Controller) CommitEditing() bool {
	rv := ffi.CallMethod[bool](c_, "commitEditing")
	return rv
}

func (c_ Controller) CommitEditingWithDelegate_DidCommitSelector_ContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](c_, "commitEditingWithDelegate:didCommitSelector:contextInfo:", delegate, didCommitSelector, contextInfo)
}

func (c_ Controller) DiscardEditing() {
	ffi.CallMethod[ffi.Void](c_, "discardEditing")
}

func (c_ Controller) IsEditing() bool {
	rv := ffi.CallMethod[bool](c_, "isEditing")
	return rv
}

var ObjectControllerClass = _ObjectControllerClass{objc.GetClass("NSObjectController")}

type _ObjectControllerClass struct {
	objc.Class
}

type IObjectController interface {
	IController
	PrepareContent()
	NewObject() objc.Object
	AddObject(object objc.IObject)
	RemoveObject(object objc.IObject)
	Add(sender objc.IObject)
	Remove(sender objc.IObject)
	Fetch(sender objc.IObject)
	ValidateUserInterfaceItem(item ValidatedUserInterfaceItem) bool
	Content() objc.Object
	SetContent(value objc.IObject)
	AutomaticallyPreparesContent() bool
	SetAutomaticallyPreparesContent(value bool)
	CanAdd() bool
	CanRemove() bool
	IsEditable() bool
	SetEditable(value bool)
	EntityName() string
	SetEntityName(value string)
	UsesLazyFetching() bool
	SetUsesLazyFetching(value bool)
	FetchPredicate() foundation.Predicate
	SetFetchPredicate(value foundation.IPredicate)
	SelectedObjects() []objc.Object
	Selection() objc.Object
}

type ObjectController struct {
	Controller
}

func MakeObjectController(ptr unsafe.Pointer) ObjectController {
	return ObjectController{
		Controller: MakeController(ptr),
	}
}

func (o_ ObjectController) InitWithContent(content objc.IObject) ObjectController {
	rv := ffi.CallMethod[ObjectController](o_, "initWithContent:", content)
	rv.Autorelease()
	return rv
}

func (o_ ObjectController) Init() ObjectController {
	rv := ffi.CallMethod[ObjectController](o_, "init")
	rv.Autorelease()
	return rv
}

func (oc _ObjectControllerClass) Alloc() ObjectController {
	rv := ffi.CallMethod[ObjectController](oc, "alloc")
	return rv
}

func (oc _ObjectControllerClass) New() ObjectController {
	rv := ffi.CallMethod[ObjectController](oc, "new")
	rv.Autorelease()
	return rv
}

func NewObjectController() ObjectController {
	return ObjectControllerClass.New()
}

func (o_ ObjectController) PrepareContent() {
	ffi.CallMethod[ffi.Void](o_, "prepareContent")
}

func (o_ ObjectController) NewObject() objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "newObject")
	rv.Autorelease()
	return rv
}

func (o_ ObjectController) AddObject(object objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "addObject:", object)
}

func (o_ ObjectController) RemoveObject(object objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "removeObject:", object)
}

func (o_ ObjectController) Add(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "add:", sender)
}

func (o_ ObjectController) Remove(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "remove:", sender)
}

func (o_ ObjectController) Fetch(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "fetch:", sender)
}

func (o_ ObjectController) ValidateUserInterfaceItem(item ValidatedUserInterfaceItem) bool {
	po := ffi.CreateProtocol(item)
	defer po.Release()
	rv := ffi.CallMethod[bool](o_, "validateUserInterfaceItem:", po)
	return rv
}

func (o_ ObjectController) Content() objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "content")
	return rv
}

func (o_ ObjectController) SetContent(value objc.IObject) {
	ffi.CallMethod[ffi.Void](o_, "setContent:", value)
}

func (o_ ObjectController) AutomaticallyPreparesContent() bool {
	rv := ffi.CallMethod[bool](o_, "automaticallyPreparesContent")
	return rv
}

func (o_ ObjectController) SetAutomaticallyPreparesContent(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setAutomaticallyPreparesContent:", value)
}

func (o_ ObjectController) CanAdd() bool {
	rv := ffi.CallMethod[bool](o_, "canAdd")
	return rv
}

func (o_ ObjectController) CanRemove() bool {
	rv := ffi.CallMethod[bool](o_, "canRemove")
	return rv
}

func (o_ ObjectController) IsEditable() bool {
	rv := ffi.CallMethod[bool](o_, "isEditable")
	return rv
}

func (o_ ObjectController) SetEditable(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setEditable:", value)
}

func (o_ ObjectController) EntityName() string {
	rv := ffi.CallMethod[string](o_, "entityName")
	return rv
}

func (o_ ObjectController) SetEntityName(value string) {
	ffi.CallMethod[ffi.Void](o_, "setEntityName:", value)
}

func (o_ ObjectController) UsesLazyFetching() bool {
	rv := ffi.CallMethod[bool](o_, "usesLazyFetching")
	return rv
}

func (o_ ObjectController) SetUsesLazyFetching(value bool) {
	ffi.CallMethod[ffi.Void](o_, "setUsesLazyFetching:", value)
}

func (o_ ObjectController) FetchPredicate() foundation.Predicate {
	rv := ffi.CallMethod[foundation.Predicate](o_, "fetchPredicate")
	return rv
}

func (o_ ObjectController) SetFetchPredicate(value foundation.IPredicate) {
	ffi.CallMethod[ffi.Void](o_, "setFetchPredicate:", value)
}

func (o_ ObjectController) SelectedObjects() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](o_, "selectedObjects")
	return rv
}

func (o_ ObjectController) Selection() objc.Object {
	rv := ffi.CallMethod[objc.Object](o_, "selection")
	return rv
}
