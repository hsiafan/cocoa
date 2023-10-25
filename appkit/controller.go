// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ControllerClass = _ControllerClass{objc.GetClass("NSController")}

type _ControllerClass struct {
	objc.Class
}

type IController interface {
	objc.IObject
	ObjectDidBeginEditing(editor objc.IObject)
	ObjectDidEndEditing(editor objc.IObject)
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
	rv := objc.CallMethod[Controller](c_, objc.SEL("init"))
	return rv
}

func (cc _ControllerClass) Alloc() Controller {
	rv := objc.CallMethod[Controller](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ControllerClass) New() Controller {
	rv := objc.CallMethod[Controller](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewController() Controller {
	return ControllerClass.New()
}

func (c_ Controller) ObjectDidBeginEditing(editor objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("objectDidBeginEditing:"), objc.ExtractPtr(editor))
}

func (c_ Controller) ObjectDidEndEditing(editor objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("objectDidEndEditing:"), objc.ExtractPtr(editor))
}

func (c_ Controller) CommitEditing() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("commitEditing"))
	return rv
}

func (c_ Controller) CommitEditingWithDelegate_DidCommitSelector_ContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](c_, objc.SEL("commitEditingWithDelegate:didCommitSelector:contextInfo:"), objc.ExtractPtr(delegate), didCommitSelector, contextInfo)
}

func (c_ Controller) DiscardEditing() {
	objc.CallMethod[objc.Void](c_, objc.SEL("discardEditing"))
}

func (c_ Controller) IsEditing() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isEditing"))
	return rv
}
