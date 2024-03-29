// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[SavePanel](sc, objc.SEL("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func (s_ SavePanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) SavePanel {
	rv := objc.CallMethod[SavePanel](s_, objc.SEL("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func (s_ SavePanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) SavePanel {
	rv := objc.CallMethod[SavePanel](s_, objc.SEL("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func (s_ SavePanel) Init() SavePanel {
	rv := objc.CallMethod[SavePanel](s_, objc.SEL("init"))
	return rv
}

func (sc _SavePanelClass) Alloc() SavePanel {
	rv := objc.CallMethod[SavePanel](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SavePanelClass) New() SavePanel {
	rv := objc.CallMethod[SavePanel](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSavePanel() SavePanel {
	return SavePanelClass.New()
}

func (sc _SavePanelClass) SavePanel() SavePanel {
	rv := objc.CallMethod[SavePanel](sc, objc.SEL("savePanel"))
	return rv
}

func (s_ SavePanel) BeginSheetModalForWindow_CompletionHandler(window IWindow, handler func(result ModalResponse)) {
	objc.CallMethod[objc.Void](s_, objc.SEL("beginSheetModalForWindow:completionHandler:"), objc.ExtractPtr(window), handler)
}

func (s_ SavePanel) BeginWithCompletionHandler(handler func(result ModalResponse)) {
	objc.CallMethod[objc.Void](s_, objc.SEL("beginWithCompletionHandler:"), handler)
}

func (s_ SavePanel) RunModal() ModalResponse {
	rv := objc.CallMethod[ModalResponse](s_, objc.SEL("runModal"))
	return rv
}

func (s_ SavePanel) ValidateVisibleColumns() {
	objc.CallMethod[objc.Void](s_, objc.SEL("validateVisibleColumns"))
}

func (s_ SavePanel) Ok(sender objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("ok:"), objc.ExtractPtr(sender))
}

func (s_ SavePanel) Cancel(sender objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("cancel:"), objc.ExtractPtr(sender))
}

// deprecated
func (s_ SavePanel) RequiredFileType() string {
	rv := objc.CallMethod[string](s_, objc.SEL("requiredFileType"))
	return rv
}

// deprecated
func (s_ SavePanel) SetRequiredFileType(type_ string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setRequiredFileType:"), type_)
}

// deprecated
func (s_ SavePanel) SetDirectory(path string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDirectory:"), path)
}

// deprecated
func (s_ SavePanel) BeginSheetForDirectory_File_ModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(path string, name string, docWindow IWindow, delegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](s_, objc.SEL("beginSheetForDirectory:file:modalForWindow:modalDelegate:didEndSelector:contextInfo:"), path, name, objc.ExtractPtr(docWindow), objc.ExtractPtr(delegate), didEndSelector, contextInfo)
}

// deprecated
func (s_ SavePanel) RunModalForDirectory_File(path string, name string) int {
	rv := objc.CallMethod[int](s_, objc.SEL("runModalForDirectory:file:"), path, name)
	return rv
}

// deprecated
func (s_ SavePanel) Panel_CompareFilename_With_CaseSensitive(sender objc.IObject, name1 string, name2 string, caseSensitive bool) foundation.ComparisonResult {
	rv := objc.CallMethod[foundation.ComparisonResult](s_, objc.SEL("panel:compareFilename:with:caseSensitive:"), objc.ExtractPtr(sender), name1, name2, caseSensitive)
	return rv
}

// deprecated
func (s_ SavePanel) Panel_IsValidFilename(sender objc.IObject, filename string) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("panel:isValidFilename:"), objc.ExtractPtr(sender), filename)
	return rv
}

// deprecated
func (s_ SavePanel) Panel_ShouldShowFilename(sender objc.IObject, filename string) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("panel:shouldShowFilename:"), objc.ExtractPtr(sender), filename)
	return rv
}

// deprecated
func (s_ SavePanel) Panel_DirectoryDidChange(sender objc.IObject, path string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("panel:directoryDidChange:"), objc.ExtractPtr(sender), path)
}

// deprecated
func (s_ SavePanel) Directory() string {
	rv := objc.CallMethod[string](s_, objc.SEL("directory"))
	return rv
}

// deprecated
func (s_ SavePanel) Filename() string {
	rv := objc.CallMethod[string](s_, objc.SEL("filename"))
	return rv
}

// deprecated
func (s_ SavePanel) SelectText(sender objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("selectText:"), objc.ExtractPtr(sender))
}

func (s_ SavePanel) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](s_, objc.SEL("URL"))
	return rv
}

func (s_ SavePanel) Prompt() string {
	rv := objc.CallMethod[string](s_, objc.SEL("prompt"))
	return rv
}

func (s_ SavePanel) SetPrompt(value string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setPrompt:"), value)
}

func (s_ SavePanel) Message() string {
	rv := objc.CallMethod[string](s_, objc.SEL("message"))
	return rv
}

func (s_ SavePanel) SetMessage(value string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMessage:"), value)
}

func (s_ SavePanel) NameFieldLabel() string {
	rv := objc.CallMethod[string](s_, objc.SEL("nameFieldLabel"))
	return rv
}

func (s_ SavePanel) SetNameFieldLabel(value string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setNameFieldLabel:"), value)
}

func (s_ SavePanel) NameFieldStringValue() string {
	rv := objc.CallMethod[string](s_, objc.SEL("nameFieldStringValue"))
	return rv
}

func (s_ SavePanel) SetNameFieldStringValue(value string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setNameFieldStringValue:"), value)
}

func (s_ SavePanel) DirectoryURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](s_, objc.SEL("directoryURL"))
	return rv
}

func (s_ SavePanel) SetDirectoryURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDirectoryURL:"), objc.ExtractPtr(value))
}

func (s_ SavePanel) AccessoryView() View {
	rv := objc.CallMethod[View](s_, objc.SEL("accessoryView"))
	return rv
}

func (s_ SavePanel) SetAccessoryView(value IView) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAccessoryView:"), objc.ExtractPtr(value))
}

func (s_ SavePanel) ShowsTagField() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("showsTagField"))
	return rv
}

func (s_ SavePanel) SetShowsTagField(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setShowsTagField:"), value)
}

func (s_ SavePanel) TagNames() []string {
	rv := objc.CallMethod[[]string](s_, objc.SEL("tagNames"))
	return rv
}

func (s_ SavePanel) SetTagNames(value []string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTagNames:"), value)
}

func (s_ SavePanel) CanCreateDirectories() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("canCreateDirectories"))
	return rv
}

func (s_ SavePanel) SetCanCreateDirectories(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setCanCreateDirectories:"), value)
}

func (s_ SavePanel) CanSelectHiddenExtension() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("canSelectHiddenExtension"))
	return rv
}

func (s_ SavePanel) SetCanSelectHiddenExtension(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setCanSelectHiddenExtension:"), value)
}

func (s_ SavePanel) ShowsHiddenFiles() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("showsHiddenFiles"))
	return rv
}

func (s_ SavePanel) SetShowsHiddenFiles(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setShowsHiddenFiles:"), value)
}

func (s_ SavePanel) IsExtensionHidden() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isExtensionHidden"))
	return rv
}

func (s_ SavePanel) SetExtensionHidden(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setExtensionHidden:"), value)
}

func (s_ SavePanel) IsExpanded() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isExpanded"))
	return rv
}

func (s_ SavePanel) AllowedContentTypes() []uniformtypeidentifiers.Type {
	rv := objc.CallMethod[[]uniformtypeidentifiers.Type](s_, objc.SEL("allowedContentTypes"))
	return rv
}

func (s_ SavePanel) SetAllowedContentTypes(value []uniformtypeidentifiers.IType) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAllowedContentTypes:"), value)
}

func (s_ SavePanel) AllowsOtherFileTypes() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("allowsOtherFileTypes"))
	return rv
}

func (s_ SavePanel) SetAllowsOtherFileTypes(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAllowsOtherFileTypes:"), value)
}

func (s_ SavePanel) TreatsFilePackagesAsDirectories() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("treatsFilePackagesAsDirectories"))
	return rv
}

func (s_ SavePanel) SetTreatsFilePackagesAsDirectories(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setTreatsFilePackagesAsDirectories:"), value)
}

// deprecated
func (s_ SavePanel) AllowedFileTypes() []string {
	rv := objc.CallMethod[[]string](s_, objc.SEL("allowedFileTypes"))
	return rv
}

// deprecated
func (s_ SavePanel) SetAllowedFileTypes(value []string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setAllowedFileTypes:"), value)
}
