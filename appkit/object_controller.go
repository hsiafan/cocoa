// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	ValidateUserInterfaceItem(item objc.IObject) bool
	Content() objc.Object
	SetContent(value objc.IObject)
	AutomaticallyPreparesContent() bool
	SetAutomaticallyPreparesContent(value bool)
	ObjectClass() objc.Class
	SetObjectClass(value objc.IClass)
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
	rv := objc.CallMethod[ObjectController](o_, objc.GetSelector("initWithContent:"), objc.ExtractPtr(content))
	return rv
}

func (o_ ObjectController) Init() ObjectController {
	rv := objc.CallMethod[ObjectController](o_, objc.GetSelector("init"))
	return rv
}

func (oc _ObjectControllerClass) Alloc() ObjectController {
	rv := objc.CallMethod[ObjectController](oc, objc.GetSelector("alloc"))
	return rv
}

func (oc _ObjectControllerClass) New() ObjectController {
	rv := objc.CallMethod[ObjectController](oc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewObjectController() ObjectController {
	return ObjectControllerClass.New()
}

func (o_ ObjectController) PrepareContent() {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("prepareContent"))
}

func (o_ ObjectController) NewObject() objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("newObject"))
	rv.Autorelease()
	return rv
}

func (o_ ObjectController) AddObject(object objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("addObject:"), objc.ExtractPtr(object))
}

func (o_ ObjectController) RemoveObject(object objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("removeObject:"), objc.ExtractPtr(object))
}

func (o_ ObjectController) Add(sender objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("add:"), objc.ExtractPtr(sender))
}

func (o_ ObjectController) Remove(sender objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("remove:"), objc.ExtractPtr(sender))
}

func (o_ ObjectController) Fetch(sender objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("fetch:"), objc.ExtractPtr(sender))
}

func (o_ ObjectController) ValidateUserInterfaceItem(item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("validateUserInterfaceItem:"), objc.ExtractPtr(item))
	return rv
}

func (o_ ObjectController) Content() objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("content"))
	return rv
}

func (o_ ObjectController) SetContent(value objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setContent:"), objc.ExtractPtr(value))
}

func (o_ ObjectController) AutomaticallyPreparesContent() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("automaticallyPreparesContent"))
	return rv
}

func (o_ ObjectController) SetAutomaticallyPreparesContent(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setAutomaticallyPreparesContent:"), value)
}

// weak property
func (o_ ObjectController) ObjectClass() objc.Class {
	rv := objc.CallMethod[objc.Class](o_, objc.GetSelector("objectClass"))
	return rv
}

// weak property
func (o_ ObjectController) SetObjectClass(value objc.IClass) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setObjectClass:"), objc.ExtractPtr(value))
}

func (o_ ObjectController) CanAdd() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("canAdd"))
	return rv
}

func (o_ ObjectController) CanRemove() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("canRemove"))
	return rv
}

func (o_ ObjectController) IsEditable() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("isEditable"))
	return rv
}

func (o_ ObjectController) SetEditable(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setEditable:"), value)
}

func (o_ ObjectController) EntityName() string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("entityName"))
	return rv
}

func (o_ ObjectController) SetEntityName(value string) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setEntityName:"), value)
}

func (o_ ObjectController) UsesLazyFetching() bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("usesLazyFetching"))
	return rv
}

func (o_ ObjectController) SetUsesLazyFetching(value bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setUsesLazyFetching:"), value)
}

func (o_ ObjectController) FetchPredicate() foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](o_, objc.GetSelector("fetchPredicate"))
	return rv
}

func (o_ ObjectController) SetFetchPredicate(value foundation.IPredicate) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("setFetchPredicate:"), objc.ExtractPtr(value))
}

func (o_ ObjectController) SelectedObjects() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](o_, objc.GetSelector("selectedObjects"))
	return rv
}

func (o_ ObjectController) Selection() objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("selection"))
	return rv
}
