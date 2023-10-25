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
	rv := objc.CallMethod[WindowController](w_, objc.SEL("initWithWindow:"), objc.ExtractPtr(window))
	return rv
}

func (w_ WindowController) InitWithWindowNibName(windowNibName NibName) WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.SEL("initWithWindowNibName:"), windowNibName)
	return rv
}

func (w_ WindowController) InitWithWindowNibName_Owner(windowNibName NibName, owner objc.IObject) WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.SEL("initWithWindowNibName:owner:"), windowNibName, objc.ExtractPtr(owner))
	return rv
}

func (w_ WindowController) InitWithWindowNibPath_Owner(windowNibPath string, owner objc.IObject) WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.SEL("initWithWindowNibPath:owner:"), windowNibPath, objc.ExtractPtr(owner))
	return rv
}

func (w_ WindowController) Init() WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.SEL("init"))
	return rv
}

func (wc _WindowControllerClass) Alloc() WindowController {
	rv := objc.CallMethod[WindowController](wc, objc.SEL("alloc"))
	return rv
}

func (wc _WindowControllerClass) New() WindowController {
	rv := objc.CallMethod[WindowController](wc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewWindowController() WindowController {
	return WindowControllerClass.New()
}

func (w_ WindowController) LoadWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("loadWindow"))
}

func (w_ WindowController) ShowWindow(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("showWindow:"), objc.ExtractPtr(sender))
}

func (w_ WindowController) WindowDidLoad() {
	objc.CallMethod[objc.Void](w_, objc.SEL("windowDidLoad"))
}

func (w_ WindowController) WindowWillLoad() {
	objc.CallMethod[objc.Void](w_, objc.SEL("windowWillLoad"))
}

func (w_ WindowController) SetDocumentEdited(dirtyFlag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setDocumentEdited:"), dirtyFlag)
}

func (w_ WindowController) Close() {
	objc.CallMethod[objc.Void](w_, objc.SEL("close"))
}

func (w_ WindowController) SynchronizeWindowTitleWithDocumentName() {
	objc.CallMethod[objc.Void](w_, objc.SEL("synchronizeWindowTitleWithDocumentName"))
}

func (w_ WindowController) WindowTitleForDocumentDisplayName(displayName string) string {
	rv := objc.CallMethod[string](w_, objc.SEL("windowTitleForDocumentDisplayName:"), displayName)
	return rv
}

func (w_ WindowController) DismissController(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("dismissController:"), objc.ExtractPtr(sender))
}

func (w_ WindowController) IsWindowLoaded() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isWindowLoaded"))
	return rv
}

func (w_ WindowController) Window() Window {
	rv := objc.CallMethod[Window](w_, objc.SEL("window"))
	return rv
}

func (w_ WindowController) SetWindow(value IWindow) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setWindow:"), objc.ExtractPtr(value))
}

// weak property
func (w_ WindowController) Document() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("document"))
	return rv
}

// weak property
func (w_ WindowController) SetDocument(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setDocument:"), objc.ExtractPtr(value))
}

func (w_ WindowController) ShouldCloseDocument() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("shouldCloseDocument"))
	return rv
}

func (w_ WindowController) SetShouldCloseDocument(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setShouldCloseDocument:"), value)
}

// weak property
func (w_ WindowController) Owner() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("owner"))
	return rv
}

func (w_ WindowController) Storyboard() Storyboard {
	rv := objc.CallMethod[Storyboard](w_, objc.SEL("storyboard"))
	return rv
}

func (w_ WindowController) WindowNibName() NibName {
	rv := objc.CallMethod[NibName](w_, objc.SEL("windowNibName"))
	return rv
}

func (w_ WindowController) WindowNibPath() string {
	rv := objc.CallMethod[string](w_, objc.SEL("windowNibPath"))
	return rv
}

func (w_ WindowController) ShouldCascadeWindows() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("shouldCascadeWindows"))
	return rv
}

func (w_ WindowController) SetShouldCascadeWindows(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setShouldCascadeWindows:"), value)
}

func (w_ WindowController) WindowFrameAutosaveName() WindowFrameAutosaveName {
	rv := objc.CallMethod[WindowFrameAutosaveName](w_, objc.SEL("windowFrameAutosaveName"))
	return rv
}

func (w_ WindowController) SetWindowFrameAutosaveName(value WindowFrameAutosaveName) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setWindowFrameAutosaveName:"), value)
}

func (w_ WindowController) ContentViewController() ViewController {
	rv := objc.CallMethod[ViewController](w_, objc.SEL("contentViewController"))
	return rv
}

func (w_ WindowController) SetContentViewController(value IViewController) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setContentViewController:"), objc.ExtractPtr(value))
}
