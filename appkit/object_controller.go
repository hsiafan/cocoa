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
	rv := objc.CallMethod[ObjectController](o_, objc.SEL("initWithContent:"), objc.ExtractPtr(content))
	return rv
}

func (o_ ObjectController) Init() ObjectController {
	rv := objc.CallMethod[ObjectController](o_, objc.SEL("init"))
	return rv
}

func (oc _ObjectControllerClass) Alloc() ObjectController {
	rv := objc.CallMethod[ObjectController](oc, objc.SEL("alloc"))
	return rv
}

func (oc _ObjectControllerClass) New() ObjectController {
	rv := objc.CallMethod[ObjectController](oc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewObjectController() ObjectController {
	return ObjectControllerClass.New()
}

func (o_ ObjectController) PrepareContent() {
	objc.CallMethod[objc.Void](o_, objc.SEL("prepareContent"))
}

func (o_ ObjectController) NewObject() objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.SEL("newObject"))
	rv.Autorelease()
	return rv
}

func (o_ ObjectController) AddObject(object objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.SEL("addObject:"), objc.ExtractPtr(object))
}

func (o_ ObjectController) RemoveObject(object objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.SEL("removeObject:"), objc.ExtractPtr(object))
}

func (o_ ObjectController) Add(sender objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.SEL("add:"), objc.ExtractPtr(sender))
}

func (o_ ObjectController) Remove(sender objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.SEL("remove:"), objc.ExtractPtr(sender))
}

func (o_ ObjectController) Fetch(sender objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.SEL("fetch:"), objc.ExtractPtr(sender))
}

func (o_ ObjectController) ValidateUserInterfaceItem(item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("validateUserInterfaceItem:"), objc.ExtractPtr(item))
	return rv
}

func (o_ ObjectController) Content() objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.SEL("content"))
	return rv
}

func (o_ ObjectController) SetContent(value objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setContent:"), objc.ExtractPtr(value))
}

func (o_ ObjectController) AutomaticallyPreparesContent() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("automaticallyPreparesContent"))
	return rv
}

func (o_ ObjectController) SetAutomaticallyPreparesContent(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setAutomaticallyPreparesContent:"), value)
}

// weak property
func (o_ ObjectController) ObjectClass() objc.Class {
	rv := objc.CallMethod[objc.Class](o_, objc.SEL("objectClass"))
	return rv
}

// weak property
func (o_ ObjectController) SetObjectClass(value objc.IClass) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setObjectClass:"), objc.ExtractPtr(value))
}

func (o_ ObjectController) CanAdd() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("canAdd"))
	return rv
}

func (o_ ObjectController) CanRemove() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("canRemove"))
	return rv
}

func (o_ ObjectController) IsEditable() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("isEditable"))
	return rv
}

func (o_ ObjectController) SetEditable(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setEditable:"), value)
}

func (o_ ObjectController) EntityName() string {
	rv := objc.CallMethod[string](o_, objc.SEL("entityName"))
	return rv
}

func (o_ ObjectController) SetEntityName(value string) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setEntityName:"), value)
}

func (o_ ObjectController) UsesLazyFetching() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("usesLazyFetching"))
	return rv
}

func (o_ ObjectController) SetUsesLazyFetching(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setUsesLazyFetching:"), value)
}

func (o_ ObjectController) FetchPredicate() foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](o_, objc.SEL("fetchPredicate"))
	return rv
}

func (o_ ObjectController) SetFetchPredicate(value foundation.IPredicate) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setFetchPredicate:"), objc.ExtractPtr(value))
}

func (o_ ObjectController) SelectedObjects() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](o_, objc.SEL("selectedObjects"))
	return rv
}

func (o_ ObjectController) Selection() objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.SEL("selection"))
	return rv
}
