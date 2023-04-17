// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/uniformtypeidentifiers"
)

var SavePanelClass = _SavePanelClass{objc.GetClass("NSSavePanel")}

type _SavePanelClass struct {
	objc.Class
}

type ISavePanel interface {
	IPanel
	BeginSheetModalForWindow_CompletionHandler(window IWindow, handler func(result ModalResponse))
	BeginWithCompletionHandler(handler func(result ModalResponse))
	RunModal() ModalResponse
	ValidateVisibleColumns()
	Ok(sender objc.IObject)
	Cancel(sender objc.IObject)
	// deprecated
	RequiredFileType() string
	// deprecated
	SetRequiredFileType(type_ string)
	// deprecated
	SetDirectory(path string)
	// deprecated
	BeginSheetForDirectory_File_ModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(path string, name string, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	// deprecated
	RunModalForDirectory_File(path string, name string) int
	// deprecated
	Panel_CompareFilename_With_CaseSensitive(sender objc.IObject, name1 string, name2 string, caseSensitive bool) foundation.ComparisonResult
	// deprecated
	Panel_IsValidFilename(sender objc.IObject, filename string) bool
	// deprecated
	Panel_ShouldShowFilename(sender objc.IObject, filename string) bool
	// deprecated
	Panel_DirectoryDidChange(sender objc.IObject, path string)
	// deprecated
	Directory() string
	// deprecated
	Filename() string
	// deprecated
	SelectText(sender objc.IObject)
	URL() foundation.URL
	Prompt() string
	SetPrompt(value string)
	Message() string
	SetMessage(value string)
	NameFieldLabel() string
	SetNameFieldLabel(value string)
	NameFieldStringValue() string
	SetNameFieldStringValue(value string)
	DirectoryURL() foundation.URL
	SetDirectoryURL(value foundation.IURL)
	AccessoryView() View
	SetAccessoryView(value IView)
	ShowsTagField() bool
	SetShowsTagField(value bool)
	TagNames() []string
	SetTagNames(value []string)
	CanCreateDirectories() bool
	SetCanCreateDirectories(value bool)
	CanSelectHiddenExtension() bool
	SetCanSelectHiddenExtension(value bool)
	ShowsHiddenFiles() bool
	SetShowsHiddenFiles(value bool)
	IsExtensionHidden() bool
	SetExtensionHidden(value bool)
	IsExpanded() bool
	AllowedContentTypes() []uniformtypeidentifiers.Type
	SetAllowedContentTypes(value []uniformtypeidentifiers.IType)
	AllowsOtherFileTypes() bool
	SetAllowsOtherFileTypes(value bool)
	TreatsFilePackagesAsDirectories() bool
	SetTreatsFilePackagesAsDirectories(value bool)
	// deprecated
	AllowedFileTypes() []string
	// deprecated
	SetAllowedFileTypes(value []string)
}

type SavePanel struct {
	Panel
}

func MakeSavePanel(ptr unsafe.Pointer) SavePanel {
	return SavePanel{
		Panel: MakePanel(ptr),
	}
}

func (sc _SavePanelClass) WindowWithContentViewController(contentViewController IViewController) SavePanel {
	rv := ffi.CallMethod[SavePanel](sc, "windowWithContentViewController:", contentViewController)
	return rv
}

func (s_ SavePanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) SavePanel {
	rv := ffi.CallMethod[SavePanel](s_, "initWithContentRect:styleMask:backing:defer:", contentRect, style, backingStoreType, flag)
	return rv
}

func (s_ SavePanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) SavePanel {
	rv := ffi.CallMethod[SavePanel](s_, "initWithContentRect:styleMask:backing:defer:screen:", contentRect, style, backingStoreType, flag, screen)
	return rv
}

func (s_ SavePanel) Init() SavePanel {
	rv := ffi.CallMethod[SavePanel](s_, "init")
	return rv
}

func (sc _SavePanelClass) Alloc() SavePanel {
	rv := ffi.CallMethod[SavePanel](sc, "alloc")
	return rv
}

func (sc _SavePanelClass) New() SavePanel {
	rv := ffi.CallMethod[SavePanel](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSavePanel() SavePanel {
	return SavePanelClass.New()
}

func (sc _SavePanelClass) SavePanel() SavePanel {
	rv := ffi.CallMethod[SavePanel](sc, "savePanel")
	return rv
}

func (s_ SavePanel) BeginSheetModalForWindow_CompletionHandler(window IWindow, handler func(result ModalResponse)) {
	ffi.CallMethod[ffi.Void](s_, "beginSheetModalForWindow:completionHandler:", window, handler)
}

func (s_ SavePanel) BeginWithCompletionHandler(handler func(result ModalResponse)) {
	ffi.CallMethod[ffi.Void](s_, "beginWithCompletionHandler:", handler)
}

func (s_ SavePanel) RunModal() ModalResponse {
	rv := ffi.CallMethod[ModalResponse](s_, "runModal")
	return rv
}

func (s_ SavePanel) ValidateVisibleColumns() {
	ffi.CallMethod[ffi.Void](s_, "validateVisibleColumns")
}

func (s_ SavePanel) Ok(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "ok:", sender)
}

func (s_ SavePanel) Cancel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "cancel:", sender)
}

// deprecated
func (s_ SavePanel) RequiredFileType() string {
	rv := ffi.CallMethod[string](s_, "requiredFileType")
	return rv
}

// deprecated
func (s_ SavePanel) SetRequiredFileType(type_ string) {
	ffi.CallMethod[ffi.Void](s_, "setRequiredFileType:", type_)
}

// deprecated
func (s_ SavePanel) SetDirectory(path string) {
	ffi.CallMethod[ffi.Void](s_, "setDirectory:", path)
}

// deprecated
func (s_ SavePanel) BeginSheetForDirectory_File_ModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(path string, name string, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](s_, "beginSheetForDirectory:file:modalForWindow:modalDelegate:didEndSelector:contextInfo:", path, name, docWindow, delegate, didEndSelector, contextInfo)
}

// deprecated
func (s_ SavePanel) RunModalForDirectory_File(path string, name string) int {
	rv := ffi.CallMethod[int](s_, "runModalForDirectory:file:", path, name)
	return rv
}

// deprecated
func (s_ SavePanel) Panel_CompareFilename_With_CaseSensitive(sender objc.IObject, name1 string, name2 string, caseSensitive bool) foundation.ComparisonResult {
	rv := ffi.CallMethod[foundation.ComparisonResult](s_, "panel:compareFilename:with:caseSensitive:", sender, name1, name2, caseSensitive)
	return rv
}

// deprecated
func (s_ SavePanel) Panel_IsValidFilename(sender objc.IObject, filename string) bool {
	rv := ffi.CallMethod[bool](s_, "panel:isValidFilename:", sender, filename)
	return rv
}

// deprecated
func (s_ SavePanel) Panel_ShouldShowFilename(sender objc.IObject, filename string) bool {
	rv := ffi.CallMethod[bool](s_, "panel:shouldShowFilename:", sender, filename)
	return rv
}

// deprecated
func (s_ SavePanel) Panel_DirectoryDidChange(sender objc.IObject, path string) {
	ffi.CallMethod[ffi.Void](s_, "panel:directoryDidChange:", sender, path)
}

// deprecated
func (s_ SavePanel) Directory() string {
	rv := ffi.CallMethod[string](s_, "directory")
	return rv
}

// deprecated
func (s_ SavePanel) Filename() string {
	rv := ffi.CallMethod[string](s_, "filename")
	return rv
}

// deprecated
func (s_ SavePanel) SelectText(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "selectText:", sender)
}

func (s_ SavePanel) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](s_, "URL")
	return rv
}

func (s_ SavePanel) Prompt() string {
	rv := ffi.CallMethod[string](s_, "prompt")
	return rv
}

func (s_ SavePanel) SetPrompt(value string) {
	ffi.CallMethod[ffi.Void](s_, "setPrompt:", value)
}

func (s_ SavePanel) Message() string {
	rv := ffi.CallMethod[string](s_, "message")
	return rv
}

func (s_ SavePanel) SetMessage(value string) {
	ffi.CallMethod[ffi.Void](s_, "setMessage:", value)
}

func (s_ SavePanel) NameFieldLabel() string {
	rv := ffi.CallMethod[string](s_, "nameFieldLabel")
	return rv
}

func (s_ SavePanel) SetNameFieldLabel(value string) {
	ffi.CallMethod[ffi.Void](s_, "setNameFieldLabel:", value)
}

func (s_ SavePanel) NameFieldStringValue() string {
	rv := ffi.CallMethod[string](s_, "nameFieldStringValue")
	return rv
}

func (s_ SavePanel) SetNameFieldStringValue(value string) {
	ffi.CallMethod[ffi.Void](s_, "setNameFieldStringValue:", value)
}

func (s_ SavePanel) DirectoryURL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](s_, "directoryURL")
	return rv
}

func (s_ SavePanel) SetDirectoryURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](s_, "setDirectoryURL:", value)
}

func (s_ SavePanel) AccessoryView() View {
	rv := ffi.CallMethod[View](s_, "accessoryView")
	return rv
}

func (s_ SavePanel) SetAccessoryView(value IView) {
	ffi.CallMethod[ffi.Void](s_, "setAccessoryView:", value)
}

func (s_ SavePanel) ShowsTagField() bool {
	rv := ffi.CallMethod[bool](s_, "showsTagField")
	return rv
}

func (s_ SavePanel) SetShowsTagField(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setShowsTagField:", value)
}

func (s_ SavePanel) TagNames() []string {
	rv := ffi.CallMethod[[]string](s_, "tagNames")
	return rv
}

func (s_ SavePanel) SetTagNames(value []string) {
	ffi.CallMethod[ffi.Void](s_, "setTagNames:", value)
}

func (s_ SavePanel) CanCreateDirectories() bool {
	rv := ffi.CallMethod[bool](s_, "canCreateDirectories")
	return rv
}

func (s_ SavePanel) SetCanCreateDirectories(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setCanCreateDirectories:", value)
}

func (s_ SavePanel) CanSelectHiddenExtension() bool {
	rv := ffi.CallMethod[bool](s_, "canSelectHiddenExtension")
	return rv
}

func (s_ SavePanel) SetCanSelectHiddenExtension(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setCanSelectHiddenExtension:", value)
}

func (s_ SavePanel) ShowsHiddenFiles() bool {
	rv := ffi.CallMethod[bool](s_, "showsHiddenFiles")
	return rv
}

func (s_ SavePanel) SetShowsHiddenFiles(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setShowsHiddenFiles:", value)
}

func (s_ SavePanel) IsExtensionHidden() bool {
	rv := ffi.CallMethod[bool](s_, "isExtensionHidden")
	return rv
}

func (s_ SavePanel) SetExtensionHidden(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setExtensionHidden:", value)
}

func (s_ SavePanel) IsExpanded() bool {
	rv := ffi.CallMethod[bool](s_, "isExpanded")
	return rv
}

func (s_ SavePanel) AllowedContentTypes() []uniformtypeidentifiers.Type {
	rv := ffi.CallMethod[[]uniformtypeidentifiers.Type](s_, "allowedContentTypes")
	return rv
}

func (s_ SavePanel) SetAllowedContentTypes(value []uniformtypeidentifiers.IType) {
	ffi.CallMethod[ffi.Void](s_, "setAllowedContentTypes:", value)
}

func (s_ SavePanel) AllowsOtherFileTypes() bool {
	rv := ffi.CallMethod[bool](s_, "allowsOtherFileTypes")
	return rv
}

func (s_ SavePanel) SetAllowsOtherFileTypes(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setAllowsOtherFileTypes:", value)
}

func (s_ SavePanel) TreatsFilePackagesAsDirectories() bool {
	rv := ffi.CallMethod[bool](s_, "treatsFilePackagesAsDirectories")
	return rv
}

func (s_ SavePanel) SetTreatsFilePackagesAsDirectories(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setTreatsFilePackagesAsDirectories:", value)
}

// deprecated
func (s_ SavePanel) AllowedFileTypes() []string {
	rv := ffi.CallMethod[[]string](s_, "allowedFileTypes")
	return rv
}

// deprecated
func (s_ SavePanel) SetAllowedFileTypes(value []string) {
	ffi.CallMethod[ffi.Void](s_, "setAllowedFileTypes:", value)
}