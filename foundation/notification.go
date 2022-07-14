// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var NotificationClass = _NotificationClass{objc.GetClass("NSNotification")}

type _NotificationClass struct {
	objc.Class
}

type INotification interface {
	objc.IObject
	Name() NotificationName
	GetObject() objc.Object
}

type Notification struct {
	objc.Object
}

func MakeNotification(ptr unsafe.Pointer) Notification {
	return Notification{
		Object: objc.MakeObject(ptr),
	}
}

func (n_ Notification) Init() Notification {
	rv := ffi.CallMethod[Notification](n_, "init")
	rv.Autorelease()
	return rv
}

func (nc _NotificationClass) NotificationWithName_Object(aName NotificationName, anObject objc.IObject) Notification {
	rv := ffi.CallMethod[Notification](nc, "notificationWithName:object:", aName, anObject)
	return rv
}

func (nc _NotificationClass) Alloc() Notification {
	rv := ffi.CallMethod[Notification](nc, "alloc")
	return rv
}

func (nc _NotificationClass) New() Notification {
	rv := ffi.CallMethod[Notification](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNotification() Notification {
	return NotificationClass.New()
}

func (n_ Notification) Name() NotificationName {
	rv := ffi.CallMethod[NotificationName](n_, "name")
	return rv
}

func (n_ Notification) GetObject() objc.Object {
	rv := ffi.CallMethod[objc.Object](n_, "object")
	return rv
}

var NotificationCenterClass = _NotificationCenterClass{objc.GetClass("NSNotificationCenter")}

type _NotificationCenterClass struct {
	objc.Class
}

type INotificationCenter interface {
	objc.IObject
	AddObserverForName_Object_Queue_UsingBlock(name NotificationName, obj objc.IObject, queue IOperationQueue, block func(note Notification)) objc.Object
	AddObserver_Selector_Name_Object(observer objc.IObject, aSelector objc.Selector, aName NotificationName, anObject objc.IObject)
	RemoveObserver_Name_Object(observer objc.IObject, aName NotificationName, anObject objc.IObject)
	RemoveObserver(observer objc.IObject)
	PostNotification(notification INotification)
	PostNotificationName_Object(aName NotificationName, anObject objc.IObject)
}

type NotificationCenter struct {
	objc.Object
}

func MakeNotificationCenter(ptr unsafe.Pointer) NotificationCenter {
	return NotificationCenter{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NotificationCenterClass) Alloc() NotificationCenter {
	rv := ffi.CallMethod[NotificationCenter](nc, "alloc")
	return rv
}

func (n_ NotificationCenter) Init() NotificationCenter {
	rv := ffi.CallMethod[NotificationCenter](n_, "init")
	rv.Autorelease()
	return rv
}

func (nc _NotificationCenterClass) New() NotificationCenter {
	rv := ffi.CallMethod[NotificationCenter](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNotificationCenter() NotificationCenter {
	return NotificationCenterClass.New()
}

func (n_ NotificationCenter) AddObserverForName_Object_Queue_UsingBlock(name NotificationName, obj objc.IObject, queue IOperationQueue, block func(note Notification)) objc.Object {
	rv := ffi.CallMethod[objc.Object](n_, "addObserverForName:object:queue:usingBlock:", name, obj, queue, block)
	return rv
}

func (n_ NotificationCenter) AddObserver_Selector_Name_Object(observer objc.IObject, aSelector objc.Selector, aName NotificationName, anObject objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "addObserver:selector:name:object:", observer, aSelector, aName, anObject)
}

func (n_ NotificationCenter) RemoveObserver_Name_Object(observer objc.IObject, aName NotificationName, anObject objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "removeObserver:name:object:", observer, aName, anObject)
}

func (n_ NotificationCenter) RemoveObserver(observer objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "removeObserver:", observer)
}

func (n_ NotificationCenter) PostNotification(notification INotification) {
	ffi.CallMethod[ffi.Void](n_, "postNotification:", notification)
}

func (n_ NotificationCenter) PostNotificationName_Object(aName NotificationName, anObject objc.IObject) {
	ffi.CallMethod[ffi.Void](n_, "postNotificationName:object:", aName, anObject)
}

func (nc _NotificationCenterClass) DefaultCenter() NotificationCenter {
	rv := ffi.CallMethod[NotificationCenter](nc, "defaultCenter")
	return rv
}
