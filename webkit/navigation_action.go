// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var NavigationActionClass = _NavigationActionClass{objc.GetClass("WKNavigationAction")}

type _NavigationActionClass struct {
	objc.Class
}

type INavigationAction interface {
	objc.IObject
	NavigationType() NavigationType
	Request() foundation.URLRequest
	SourceFrame() FrameInfo
	TargetFrame() FrameInfo
	ButtonNumber() int
	ModifierFlags() appkit.EventModifierFlags
	ShouldPerformDownload() bool
}

type NavigationAction struct {
	objc.Object
}

func MakeNavigationAction(ptr unsafe.Pointer) NavigationAction {
	return NavigationAction{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NavigationActionClass) Alloc() NavigationAction {
	rv := ffi.CallMethod[NavigationAction](nc, "alloc")
	return rv
}

func (nc _NavigationActionClass) New() NavigationAction {
	rv := ffi.CallMethod[NavigationAction](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNavigationAction() NavigationAction {
	return NavigationActionClass.New()
}

func (n_ NavigationAction) Init() NavigationAction {
	rv := ffi.CallMethod[NavigationAction](n_, "init")
	return rv
}

func (n_ NavigationAction) NavigationType() NavigationType {
	rv := ffi.CallMethod[NavigationType](n_, "navigationType")
	return rv
}

func (n_ NavigationAction) Request() foundation.URLRequest {
	rv := ffi.CallMethod[foundation.URLRequest](n_, "request")
	return rv
}

func (n_ NavigationAction) SourceFrame() FrameInfo {
	rv := ffi.CallMethod[FrameInfo](n_, "sourceFrame")
	return rv
}

func (n_ NavigationAction) TargetFrame() FrameInfo {
	rv := ffi.CallMethod[FrameInfo](n_, "targetFrame")
	return rv
}

func (n_ NavigationAction) ButtonNumber() int {
	rv := ffi.CallMethod[int](n_, "buttonNumber")
	return rv
}

func (n_ NavigationAction) ModifierFlags() appkit.EventModifierFlags {
	rv := ffi.CallMethod[appkit.EventModifierFlags](n_, "modifierFlags")
	return rv
}

func (n_ NavigationAction) ShouldPerformDownload() bool {
	rv := ffi.CallMethod[bool](n_, "shouldPerformDownload")
	return rv
}