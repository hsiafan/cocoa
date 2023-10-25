// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/appkit"
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
	rv := objc.CallMethod[NavigationAction](nc, objc.SEL("alloc"))
	return rv
}

func (nc _NavigationActionClass) New() NavigationAction {
	rv := objc.CallMethod[NavigationAction](nc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewNavigationAction() NavigationAction {
	return NavigationActionClass.New()
}

func (n_ NavigationAction) Init() NavigationAction {
	rv := objc.CallMethod[NavigationAction](n_, objc.SEL("init"))
	return rv
}

func (n_ NavigationAction) NavigationType() NavigationType {
	rv := objc.CallMethod[NavigationType](n_, objc.SEL("navigationType"))
	return rv
}

func (n_ NavigationAction) Request() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](n_, objc.SEL("request"))
	return rv
}

func (n_ NavigationAction) SourceFrame() FrameInfo {
	rv := objc.CallMethod[FrameInfo](n_, objc.SEL("sourceFrame"))
	return rv
}

func (n_ NavigationAction) TargetFrame() FrameInfo {
	rv := objc.CallMethod[FrameInfo](n_, objc.SEL("targetFrame"))
	return rv
}

func (n_ NavigationAction) ButtonNumber() int {
	rv := objc.CallMethod[int](n_, objc.SEL("buttonNumber"))
	return rv
}

func (n_ NavigationAction) ModifierFlags() appkit.EventModifierFlags {
	rv := objc.CallMethod[appkit.EventModifierFlags](n_, objc.SEL("modifierFlags"))
	return rv
}

func (n_ NavigationAction) ShouldPerformDownload() bool {
	rv := objc.CallMethod[bool](n_, objc.SEL("shouldPerformDownload"))
	return rv
}
