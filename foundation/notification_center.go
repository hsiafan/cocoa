// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var NotificationCenterClass = _NotificationCenterClass{objc.GetClass("NSNotificationCenter")}

type _NotificationCenterClass struct {
	objc.Class
}

type INotificationCenter interface {
	objc.IObject
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
	rv := objc.CallMethod[NotificationCenter](nc, objc.SEL("alloc"))
	return rv
}

func (nc _NotificationCenterClass) New() NotificationCenter {
	rv := objc.CallMethod[NotificationCenter](nc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewNotificationCenter() NotificationCenter {
	return NotificationCenterClass.New()
}

func (n_ NotificationCenter) Init() NotificationCenter {
	rv := objc.CallMethod[NotificationCenter](n_, objc.SEL("init"))
	return rv
}

func (n_ NotificationCenter) AddObserver_Selector_Name_Object(observer objc.IObject, aSelector objc.Selector, aName NotificationName, anObject objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("addObserver:selector:name:object:"), objc.ExtractPtr(observer), aSelector, aName, objc.ExtractPtr(anObject))
}

func (n_ NotificationCenter) RemoveObserver_Name_Object(observer objc.IObject, aName NotificationName, anObject objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("removeObserver:name:object:"), objc.ExtractPtr(observer), aName, objc.ExtractPtr(anObject))
}

func (n_ NotificationCenter) RemoveObserver(observer objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("removeObserver:"), objc.ExtractPtr(observer))
}

func (n_ NotificationCenter) PostNotification(notification INotification) {
	objc.CallMethod[objc.Void](n_, objc.SEL("postNotification:"), objc.ExtractPtr(notification))
}

func (n_ NotificationCenter) PostNotificationName_Object(aName NotificationName, anObject objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.SEL("postNotificationName:object:"), aName, objc.ExtractPtr(anObject))
}

func (nc _NotificationCenterClass) DefaultCenter() NotificationCenter {
	rv := objc.CallMethod[NotificationCenter](nc, objc.SEL("defaultCenter"))
	return rv
}
