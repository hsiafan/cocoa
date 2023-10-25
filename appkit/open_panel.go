// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var OpenPanelClass = _OpenPanelClass{objc.GetClass("NSOpenPanel")}

type _OpenPanelClass struct {
	objc.Class
}

type IOpenPanel interface {
	ISavePanel
	// deprecated
	Filenames() []objc.Object
	// deprecated
	BeginForDirectory_File_Types_ModelessDelegate_DidEndSelector_ContextInfo(path string, name string, fileTypes []objc.IObject, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	// deprecated
	BeginSheetForDirectory_File_Types_ModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(path string, name string, fileTypes []objc.IObject, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	// deprecated
	RunModalForDirectory_File_Types(path string, name string, fileTypes []objc.IObject) int
	// deprecated
	RunModalForTypes(fileTypes []objc.IObject) int
	CanChooseFiles() bool
	SetCanChooseFiles(value bool)
	CanChooseDirectories() bool
	SetCanChooseDirectories(value bool)
	ResolvesAliases() bool
	SetResolvesAliases(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	IsAccessoryViewDisclosed() bool
	SetAccessoryViewDisclosed(value bool)
	URLs() []foundation.URL
	CanDownloadUbiquitousContents() bool
	SetCanDownloadUbiquitousContents(value bool)
	CanResolveUbiquitousConflicts() bool
	SetCanResolveUbiquitousConflicts(value bool)
}

type OpenPanel struct {
	SavePanel
}

func MakeOpenPanel(ptr unsafe.Pointer) OpenPanel {
	return OpenPanel{
		SavePanel: MakeSavePanel(ptr),
	}
}

func (oc _OpenPanelClass) WindowWithContentViewController(contentViewController IViewController) OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, objc.SEL("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func (o_ OpenPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) OpenPanel {
	rv := objc.CallMethod[OpenPanel](o_, objc.SEL("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func (o_ OpenPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) OpenPanel {
	rv := objc.CallMethod[OpenPanel](o_, objc.SEL("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func (o_ OpenPanel) Init() OpenPanel {
	rv := objc.CallMethod[OpenPanel](o_, objc.SEL("init"))
	return rv
}

func (oc _OpenPanelClass) Alloc() OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, objc.SEL("alloc"))
	return rv
}

func (oc _OpenPanelClass) New() OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewOpenPanel() OpenPanel {
	return OpenPanelClass.New()
}

func (oc _OpenPanelClass) OpenPanel() OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, objc.SEL("openPanel"))
	return rv
}

// deprecated
func (o_ OpenPanel) Filenames() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](o_, objc.SEL("filenames"))
	return rv
}

// deprecated
func (o_ OpenPanel) BeginForDirectory_File_Types_ModelessDelegate_DidEndSelector_ContextInfo(path string, name string, fileTypes []objc.IObject, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](o_, objc.SEL("beginForDirectory:file:types:modelessDelegate:didEndSelector:contextInfo:"), path, name, fileTypes, objc.ExtractPtr(delegate), didEndSelector, contextInfo)
}

// deprecated
func (o_ OpenPanel) BeginSheetForDirectory_File_Types_ModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(path string, name string, fileTypes []objc.IObject, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](o_, objc.SEL("beginSheetForDirectory:file:types:modalForWindow:modalDelegate:didEndSelector:contextInfo:"), path, name, fileTypes, objc.ExtractPtr(docWindow), objc.ExtractPtr(delegate), didEndSelector, contextInfo)
}

// deprecated
func (o_ OpenPanel) RunModalForDirectory_File_Types(path string, name string, fileTypes []objc.IObject) int {
	rv := objc.CallMethod[int](o_, objc.SEL("runModalForDirectory:file:types:"), path, name, fileTypes)
	return rv
}

// deprecated
func (o_ OpenPanel) RunModalForTypes(fileTypes []objc.IObject) int {
	rv := objc.CallMethod[int](o_, objc.SEL("runModalForTypes:"), fileTypes)
	return rv
}

func (o_ OpenPanel) CanChooseFiles() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("canChooseFiles"))
	return rv
}

func (o_ OpenPanel) SetCanChooseFiles(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setCanChooseFiles:"), value)
}

func (o_ OpenPanel) CanChooseDirectories() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("canChooseDirectories"))
	return rv
}

func (o_ OpenPanel) SetCanChooseDirectories(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setCanChooseDirectories:"), value)
}

func (o_ OpenPanel) ResolvesAliases() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("resolvesAliases"))
	return rv
}

func (o_ OpenPanel) SetResolvesAliases(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setResolvesAliases:"), value)
}

func (o_ OpenPanel) AllowsMultipleSelection() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("allowsMultipleSelection"))
	return rv
}

func (o_ OpenPanel) SetAllowsMultipleSelection(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setAllowsMultipleSelection:"), value)
}

func (o_ OpenPanel) IsAccessoryViewDisclosed() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("isAccessoryViewDisclosed"))
	return rv
}

func (o_ OpenPanel) SetAccessoryViewDisclosed(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setAccessoryViewDisclosed:"), value)
}

func (o_ OpenPanel) URLs() []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](o_, objc.SEL("URLs"))
	return rv
}

func (o_ OpenPanel) CanDownloadUbiquitousContents() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("canDownloadUbiquitousContents"))
	return rv
}

func (o_ OpenPanel) SetCanDownloadUbiquitousContents(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setCanDownloadUbiquitousContents:"), value)
}

func (o_ OpenPanel) CanResolveUbiquitousConflicts() bool {
	rv := objc.CallMethod[bool](o_, objc.SEL("canResolveUbiquitousConflicts"))
	return rv
}

func (o_ OpenPanel) SetCanResolveUbiquitousConflicts(value bool) {
	objc.CallMethod[objc.Void](o_, objc.SEL("setCanResolveUbiquitousConflicts:"), value)
}
