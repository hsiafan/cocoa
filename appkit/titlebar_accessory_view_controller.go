// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TitlebarAccessoryViewControllerClass = _TitlebarAccessoryViewControllerClass{objc.GetClass("NSTitlebarAccessoryViewController")}

type _TitlebarAccessoryViewControllerClass struct {
	objc.Class
}

type ITitlebarAccessoryViewController interface {
	IViewController
	FullScreenMinHeight() float64
	SetFullScreenMinHeight(value float64)
	LayoutAttribute() LayoutAttribute
	SetLayoutAttribute(value LayoutAttribute)
	AutomaticallyAdjustsSize() bool
	SetAutomaticallyAdjustsSize(value bool)
	IsHidden() bool
	SetHidden(value bool)
}

type TitlebarAccessoryViewController struct {
	ViewController
}

func MakeTitlebarAccessoryViewController(ptr unsafe.Pointer) TitlebarAccessoryViewController {
	return TitlebarAccessoryViewController{
		ViewController: MakeViewController(ptr),
	}
}

func (t_ TitlebarAccessoryViewController) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) TitlebarAccessoryViewController {
	rv := objc.CallMethod[TitlebarAccessoryViewController](t_, "initWithNibName:bundle:", nibNameOrNil, nibBundleOrNil)
	return rv
}

func (t_ TitlebarAccessoryViewController) Init() TitlebarAccessoryViewController {
	rv := objc.CallMethod[TitlebarAccessoryViewController](t_, "init")
	return rv
}

func (tc _TitlebarAccessoryViewControllerClass) Alloc() TitlebarAccessoryViewController {
	rv := objc.CallMethod[TitlebarAccessoryViewController](tc, "alloc")
	return rv
}

func (tc _TitlebarAccessoryViewControllerClass) New() TitlebarAccessoryViewController {
	rv := objc.CallMethod[TitlebarAccessoryViewController](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTitlebarAccessoryViewController() TitlebarAccessoryViewController {
	return TitlebarAccessoryViewControllerClass.New()
}

func (t_ TitlebarAccessoryViewController) FullScreenMinHeight() float64 {
	rv := objc.CallMethod[float64](t_, "fullScreenMinHeight")
	return rv
}

func (t_ TitlebarAccessoryViewController) SetFullScreenMinHeight(value float64) {
	objc.CallMethod[objc.Void](t_, "setFullScreenMinHeight:", value)
}

func (t_ TitlebarAccessoryViewController) LayoutAttribute() LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](t_, "layoutAttribute")
	return rv
}

func (t_ TitlebarAccessoryViewController) SetLayoutAttribute(value LayoutAttribute) {
	objc.CallMethod[objc.Void](t_, "setLayoutAttribute:", value)
}

func (t_ TitlebarAccessoryViewController) AutomaticallyAdjustsSize() bool {
	rv := objc.CallMethod[bool](t_, "automaticallyAdjustsSize")
	return rv
}

func (t_ TitlebarAccessoryViewController) SetAutomaticallyAdjustsSize(value bool) {
	objc.CallMethod[objc.Void](t_, "setAutomaticallyAdjustsSize:", value)
}

func (t_ TitlebarAccessoryViewController) IsHidden() bool {
	rv := objc.CallMethod[bool](t_, "isHidden")
	return rv
}

func (t_ TitlebarAccessoryViewController) SetHidden(value bool) {
	objc.CallMethod[objc.Void](t_, "setHidden:", value)
}
