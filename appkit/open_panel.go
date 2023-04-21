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
	rv := objc.CallMethod[OpenPanel](oc, "windowWithContentViewController:", contentViewController)
	return rv
}

func (o_ OpenPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) OpenPanel {
	rv := objc.CallMethod[OpenPanel](o_, "initWithContentRect:styleMask:backing:defer:", contentRect, style, backingStoreType, flag)
	return rv
}

func (o_ OpenPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) OpenPanel {
	rv := objc.CallMethod[OpenPanel](o_, "initWithContentRect:styleMask:backing:defer:screen:", contentRect, style, backingStoreType, flag, screen)
	return rv
}

func (o_ OpenPanel) Init() OpenPanel {
	rv := objc.CallMethod[OpenPanel](o_, "init")
	return rv
}

func (oc _OpenPanelClass) Alloc() OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, "alloc")
	return rv
}

func (oc _OpenPanelClass) New() OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, "new")
	rv.Autorelease()
	return rv
}

func NewOpenPanel() OpenPanel {
	return OpenPanelClass.New()
}

func (oc _OpenPanelClass) OpenPanel() OpenPanel {
	rv := objc.CallMethod[OpenPanel](oc, "openPanel")
	return rv
}

// deprecated
func (o_ OpenPanel) Filenames() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](o_, "filenames")
	return rv
}

// deprecated
func (o_ OpenPanel) BeginForDirectory_File_Types_ModelessDelegate_DidEndSelector_ContextInfo(path string, name string, fileTypes []objc.IObject, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](o_, "beginForDirectory:file:types:modelessDelegate:didEndSelector:contextInfo:", path, name, fileTypes, delegate, didEndSelector, contextInfo)
}

// deprecated
func (o_ OpenPanel) BeginSheetForDirectory_File_Types_ModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(path string, name string, fileTypes []objc.IObject, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](o_, "beginSheetForDirectory:file:types:modalForWindow:modalDelegate:didEndSelector:contextInfo:", path, name, fileTypes, docWindow, delegate, didEndSelector, contextInfo)
}

// deprecated
func (o_ OpenPanel) RunModalForDirectory_File_Types(path string, name string, fileTypes []objc.IObject) int {
	rv := objc.CallMethod[int](o_, "runModalForDirectory:file:types:", path, name, fileTypes)
	return rv
}

// deprecated
func (o_ OpenPanel) RunModalForTypes(fileTypes []objc.IObject) int {
	rv := objc.CallMethod[int](o_, "runModalForTypes:", fileTypes)
	return rv
}

func (o_ OpenPanel) CanChooseFiles() bool {
	rv := objc.CallMethod[bool](o_, "canChooseFiles")
	return rv
}

func (o_ OpenPanel) SetCanChooseFiles(value bool) {
	objc.CallMethod[objc.Void](o_, "setCanChooseFiles:", value)
}

func (o_ OpenPanel) CanChooseDirectories() bool {
	rv := objc.CallMethod[bool](o_, "canChooseDirectories")
	return rv
}

func (o_ OpenPanel) SetCanChooseDirectories(value bool) {
	objc.CallMethod[objc.Void](o_, "setCanChooseDirectories:", value)
}

func (o_ OpenPanel) ResolvesAliases() bool {
	rv := objc.CallMethod[bool](o_, "resolvesAliases")
	return rv
}

func (o_ OpenPanel) SetResolvesAliases(value bool) {
	objc.CallMethod[objc.Void](o_, "setResolvesAliases:", value)
}

func (o_ OpenPanel) AllowsMultipleSelection() bool {
	rv := objc.CallMethod[bool](o_, "allowsMultipleSelection")
	return rv
}

func (o_ OpenPanel) SetAllowsMultipleSelection(value bool) {
	objc.CallMethod[objc.Void](o_, "setAllowsMultipleSelection:", value)
}

func (o_ OpenPanel) IsAccessoryViewDisclosed() bool {
	rv := objc.CallMethod[bool](o_, "isAccessoryViewDisclosed")
	return rv
}

func (o_ OpenPanel) SetAccessoryViewDisclosed(value bool) {
	objc.CallMethod[objc.Void](o_, "setAccessoryViewDisclosed:", value)
}

func (o_ OpenPanel) URLs() []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](o_, "URLs")
	return rv
}

func (o_ OpenPanel) CanDownloadUbiquitousContents() bool {
	rv := objc.CallMethod[bool](o_, "canDownloadUbiquitousContents")
	return rv
}

func (o_ OpenPanel) SetCanDownloadUbiquitousContents(value bool) {
	objc.CallMethod[objc.Void](o_, "setCanDownloadUbiquitousContents:", value)
}

func (o_ OpenPanel) CanResolveUbiquitousConflicts() bool {
	rv := objc.CallMethod[bool](o_, "canResolveUbiquitousConflicts")
	return rv
}

func (o_ OpenPanel) SetCanResolveUbiquitousConflicts(value bool) {
	objc.CallMethod[objc.Void](o_, "setCanResolveUbiquitousConflicts:", value)
}
