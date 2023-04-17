// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[WindowController](w_, "initWithWindow:", window)
	return rv
}

func (w_ WindowController) InitWithWindowNibName(windowNibName NibName) WindowController {
	rv := ffi.CallMethod[WindowController](w_, "initWithWindowNibName:", windowNibName)
	return rv
}

func (w_ WindowController) InitWithWindowNibName_Owner(windowNibName NibName, owner objc.IObject) WindowController {
	rv := ffi.CallMethod[WindowController](w_, "initWithWindowNibName:owner:", windowNibName, owner)
	return rv
}

func (w_ WindowController) InitWithWindowNibPath_Owner(windowNibPath string, owner objc.IObject) WindowController {
	rv := ffi.CallMethod[WindowController](w_, "initWithWindowNibPath:owner:", windowNibPath, owner)
	return rv
}

func (w_ WindowController) Init() WindowController {
	rv := ffi.CallMethod[WindowController](w_, "init")
	return rv
}

func (wc _WindowControllerClass) Alloc() WindowController {
	rv := ffi.CallMethod[WindowController](wc, "alloc")
	return rv
}

func (wc _WindowControllerClass) New() WindowController {
	rv := ffi.CallMethod[WindowController](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWindowController() WindowController {
	return WindowControllerClass.New()
}

func (w_ WindowController) LoadWindow() {
	ffi.CallMethod[ffi.Void](w_, "loadWindow")
}

func (w_ WindowController) ShowWindow(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "showWindow:", sender)
}

func (w_ WindowController) WindowDidLoad() {
	ffi.CallMethod[ffi.Void](w_, "windowDidLoad")
}

func (w_ WindowController) WindowWillLoad() {
	ffi.CallMethod[ffi.Void](w_, "windowWillLoad")
}

func (w_ WindowController) SetDocumentEdited(dirtyFlag bool) {
	ffi.CallMethod[ffi.Void](w_, "setDocumentEdited:", dirtyFlag)
}

func (w_ WindowController) Close() {
	ffi.CallMethod[ffi.Void](w_, "close")
}

func (w_ WindowController) SynchronizeWindowTitleWithDocumentName() {
	ffi.CallMethod[ffi.Void](w_, "synchronizeWindowTitleWithDocumentName")
}

func (w_ WindowController) WindowTitleForDocumentDisplayName(displayName string) string {
	rv := ffi.CallMethod[string](w_, "windowTitleForDocumentDisplayName:", displayName)
	return rv
}

func (w_ WindowController) DismissController(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "dismissController:", sender)
}

func (w_ WindowController) IsWindowLoaded() bool {
	rv := ffi.CallMethod[bool](w_, "isWindowLoaded")
	return rv
}

func (w_ WindowController) Window() Window {
	rv := ffi.CallMethod[Window](w_, "window")
	return rv
}

func (w_ WindowController) SetWindow(value IWindow) {
	ffi.CallMethod[ffi.Void](w_, "setWindow:", value)
}

func (w_ WindowController) Document() objc.Object {
	rv := ffi.CallMethod[objc.Object](w_, "document")
	return rv
}

func (w_ WindowController) SetDocument(value objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "setDocument:", value)
}

func (w_ WindowController) ShouldCloseDocument() bool {
	rv := ffi.CallMethod[bool](w_, "shouldCloseDocument")
	return rv
}

func (w_ WindowController) SetShouldCloseDocument(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setShouldCloseDocument:", value)
}

func (w_ WindowController) Owner() objc.Object {
	rv := ffi.CallMethod[objc.Object](w_, "owner")
	return rv
}

func (w_ WindowController) Storyboard() Storyboard {
	rv := ffi.CallMethod[Storyboard](w_, "storyboard")
	return rv
}

func (w_ WindowController) WindowNibName() NibName {
	rv := ffi.CallMethod[NibName](w_, "windowNibName")
	return rv
}

func (w_ WindowController) WindowNibPath() string {
	rv := ffi.CallMethod[string](w_, "windowNibPath")
	return rv
}

func (w_ WindowController) ShouldCascadeWindows() bool {
	rv := ffi.CallMethod[bool](w_, "shouldCascadeWindows")
	return rv
}

func (w_ WindowController) SetShouldCascadeWindows(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setShouldCascadeWindows:", value)
}

func (w_ WindowController) WindowFrameAutosaveName() WindowFrameAutosaveName {
	rv := ffi.CallMethod[WindowFrameAutosaveName](w_, "windowFrameAutosaveName")
	return rv
}

func (w_ WindowController) SetWindowFrameAutosaveName(value WindowFrameAutosaveName) {
	ffi.CallMethod[ffi.Void](w_, "setWindowFrameAutosaveName:", value)
}

func (w_ WindowController) ContentViewController() ViewController {
	rv := ffi.CallMethod[ViewController](w_, "contentViewController")
	return rv
}

func (w_ WindowController) SetContentViewController(value IViewController) {
	ffi.CallMethod[ffi.Void](w_, "setContentViewController:", value)
}
