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
	Object_() objc.Object
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

func (n_ Notification) Object_() objc.Object {
	rv := ffi.CallMethod[objc.Object](n_, "object")
	return rv
}
