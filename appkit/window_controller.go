// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var WindowControllerClass = _WindowControllerClass{objc.GetClass("NSWindowController")}

type _WindowControllerClass struct {
	objc.Class
}

type IWindowController interface {
	IResponder
	LoadWindow()
	ShowWindow(sender objc.IObject)
	WindowDidLoad()
	WindowWillLoad()
	SetDocumentEdited(dirtyFlag bool)
	Close()
	SynchronizeWindowTitleWithDocumentName()
	WindowTitleForDocumentDisplayName(displayName string) string
	DismissController(sender objc.IObject)
	IsWindowLoaded() bool
	Window() Window
	SetWindow(value IWindow)
	Document() objc.Object
	SetDocument(value objc.IObject)
	ShouldCloseDocument() bool
	SetShouldCloseDocument(value bool)
	Owner() objc.Object
	Storyboard() Storyboard
	WindowNibName() NibName
	WindowNibPath() string
	ShouldCascadeWindows() bool
	SetShouldCascadeWindows(value bool)
	WindowFrameAutosaveName() WindowFrameAutosaveName
	SetWindowFrameAutosaveName(value WindowFrameAutosaveName)
	ContentViewController() ViewController
	SetContentViewController(value IViewController)
}

type WindowController struct {
	Responder
}

func MakeWindowController(ptr unsafe.Pointer) WindowController {
	return WindowController{
		Responder: MakeResponder(ptr),
	}
}

func (w_ WindowController) InitWithWindow(window IWindow) WindowController {
	rv := objc.CallMethod[WindowController](w_, "initWithWindow:", window)
	return rv
}

func (w_ WindowController) InitWithWindowNibName(windowNibName NibName) WindowController {
	rv := objc.CallMethod[WindowController](w_, "initWithWindowNibName:", windowNibName)
	return rv
}

func (w_ WindowController) InitWithWindowNibName_Owner(windowNibName NibName, owner objc.IObject) WindowController {
	rv := objc.CallMethod[WindowController](w_, "initWithWindowNibName:owner:", windowNibName, owner)
	return rv
}

func (w_ WindowController) InitWithWindowNibPath_Owner(windowNibPath string, owner objc.IObject) WindowController {
	rv := objc.CallMethod[WindowController](w_, "initWithWindowNibPath:owner:", windowNibPath, owner)
	return rv
}

func (w_ WindowController) Init() WindowController {
	rv := objc.CallMethod[WindowController](w_, "init")
	return rv
}

func (wc _WindowControllerClass) Alloc() WindowController {
	rv := objc.CallMethod[WindowController](wc, "alloc")
	return rv
}

func (wc _WindowControllerClass) New() WindowController {
	rv := objc.CallMethod[WindowController](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWindowController() WindowController {
	return WindowControllerClass.New()
}

func (w_ WindowController) LoadWindow() {
	objc.CallMethod[objc.Void](w_, "loadWindow")
}

func (w_ WindowController) ShowWindow(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "showWindow:", sender)
}

func (w_ WindowController) WindowDidLoad() {
	objc.CallMethod[objc.Void](w_, "windowDidLoad")
}

func (w_ WindowController) WindowWillLoad() {
	objc.CallMethod[objc.Void](w_, "windowWillLoad")
}

func (w_ WindowController) SetDocumentEdited(dirtyFlag bool) {
	objc.CallMethod[objc.Void](w_, "setDocumentEdited:", dirtyFlag)
}

func (w_ WindowController) Close() {
	objc.CallMethod[objc.Void](w_, "close")
}

func (w_ WindowController) SynchronizeWindowTitleWithDocumentName() {
	objc.CallMethod[objc.Void](w_, "synchronizeWindowTitleWithDocumentName")
}

func (w_ WindowController) WindowTitleForDocumentDisplayName(displayName string) string {
	rv := objc.CallMethod[string](w_, "windowTitleForDocumentDisplayName:", displayName)
	return rv
}

func (w_ WindowController) DismissController(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "dismissController:", sender)
}

func (w_ WindowController) IsWindowLoaded() bool {
	rv := objc.CallMethod[bool](w_, "isWindowLoaded")
	return rv
}

func (w_ WindowController) Window() Window {
	rv := objc.CallMethod[Window](w_, "window")
	return rv
}

func (w_ WindowController) SetWindow(value IWindow) {
	objc.CallMethod[objc.Void](w_, "setWindow:", value)
}

func (w_ WindowController) Document() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, "document")
	return rv
}

func (w_ WindowController) SetDocument(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, "setDocument:", value)
}

func (w_ WindowController) ShouldCloseDocument() bool {
	rv := objc.CallMethod[bool](w_, "shouldCloseDocument")
	return rv
}

func (w_ WindowController) SetShouldCloseDocument(value bool) {
	objc.CallMethod[objc.Void](w_, "setShouldCloseDocument:", value)
}

func (w_ WindowController) Owner() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, "owner")
	return rv
}

func (w_ WindowController) Storyboard() Storyboard {
	rv := objc.CallMethod[Storyboard](w_, "storyboard")
	return rv
}

func (w_ WindowController) WindowNibName() NibName {
	rv := objc.CallMethod[NibName](w_, "windowNibName")
	return rv
}

func (w_ WindowController) WindowNibPath() string {
	rv := objc.CallMethod[string](w_, "windowNibPath")
	return rv
}

func (w_ WindowController) ShouldCascadeWindows() bool {
	rv := objc.CallMethod[bool](w_, "shouldCascadeWindows")
	return rv
}

func (w_ WindowController) SetShouldCascadeWindows(value bool) {
	objc.CallMethod[objc.Void](w_, "setShouldCascadeWindows:", value)
}

func (w_ WindowController) WindowFrameAutosaveName() WindowFrameAutosaveName {
	rv := objc.CallMethod[WindowFrameAutosaveName](w_, "windowFrameAutosaveName")
	return rv
}

func (w_ WindowController) SetWindowFrameAutosaveName(value WindowFrameAutosaveName) {
	objc.CallMethod[objc.Void](w_, "setWindowFrameAutosaveName:", value)
}

func (w_ WindowController) ContentViewController() ViewController {
	rv := objc.CallMethod[ViewController](w_, "contentViewController")
	return rv
}

func (w_ WindowController) SetContentViewController(value IViewController) {
	objc.CallMethod[objc.Void](w_, "setContentViewController:", value)
}
