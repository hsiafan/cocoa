// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/uniformtypeidentifiers"
)

var WorkspaceClass = _WorkspaceClass{objc.GetClass("NSWorkspace")}

type _WorkspaceClass struct {
	objc.Class
}

type IWorkspace interface {
	objc.IObject
	OpenURL(url foundation.IURL) bool
	HideOtherApplications()
	DuplicateURLs_CompletionHandler(URLs []foundation.IURL, handler func(newURLs foundation.Dictionary, error foundation.Error))
	RecycleURLs_CompletionHandler(URLs []foundation.IURL, handler func(newURLs foundation.Dictionary, error foundation.Error))
	ActivateFileViewerSelectingURLs(fileURLs []foundation.IURL)
	SelectFile_InFileViewerRootedAtPath(fullPath string, rootFullPath string) bool
	// deprecated
	TypeOfFile_Error(absoluteFilePath string, outError *foundation.Error) string
	// deprecated
	LocalizedDescriptionForType(typeName string) string
	// deprecated
	PreferredFilenameExtensionForType(typeName string) string
	// deprecated
	FilenameExtension_IsValidForType(filenameExtension string, typeName string) bool
	// deprecated
	Type_ConformsToType(firstTypeName string, secondTypeName string) bool
	URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.URL
	// deprecated
	GetInfoForFile_Application_Type(fullPath string, appName *foundation.String, _type *foundation.String) bool
	URLForApplicationToOpenURL(url foundation.IURL) foundation.URL
	GetFileSystemInfoForPath_IsRemovable_IsWritable_IsUnmountable_Description_Type(fullPath string, removableFlag *bool, writableFlag *bool, unmountableFlag *bool, description *foundation.String, fileSystemType *foundation.String) bool
	IsFilePackageAtPath(fullPath string) bool
	IconForFile(fullPath string) Image
	IconForFiles(fullPaths []string) Image
	IconForContentType(contentType uniformtypeidentifiers.IType) Image
	SetIcon_ForFile_Options(image IImage, fullPath string, options WorkspaceIconCreationOptions) bool
	UnmountAndEjectDeviceAtPath(path string) bool
	UnmountAndEjectDeviceAtURL_Error(url foundation.IURL, error *foundation.Error) bool
	DesktopImageURLForScreen(screen IScreen) foundation.URL
	SetDesktopImageURL_ForScreen_Options_Error(url foundation.IURL, screen IScreen, options map[WorkspaceDesktopImageOptionKey]objc.IObject, error *foundation.Error) bool
	DesktopImageOptionsForScreen(screen IScreen) map[WorkspaceDesktopImageOptionKey]objc.Object
	ShowSearchResultsForQueryString(queryString string) bool
	NoteFileSystemChanged_(path string)
	ExtendPowerOffBy(requested int) int
	SetDefaultApplicationAtURL_ToOpenContentType_CompletionHandler(applicationURL foundation.IURL, contentType uniformtypeidentifiers.IType, completionHandler func(error foundation.Error))
	SetDefaultApplicationAtURL_ToOpenContentTypeOfFileAtURL_CompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error))
	SetDefaultApplicationAtURL_ToOpenFileAtURL_CompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error))
	SetDefaultApplicationAtURL_ToOpenURLsWithScheme_CompletionHandler(applicationURL foundation.IURL, urlScheme string, completionHandler func(error foundation.Error))
	URLForApplicationToOpenContentType(contentType uniformtypeidentifiers.IType) foundation.URL
	URLsForApplicationsToOpenContentType(contentType uniformtypeidentifiers.IType) []foundation.URL
	URLsForApplicationsToOpenURL(url foundation.IURL) []foundation.URL
	URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.URL
	// deprecated
	OpenURL_Options_Configuration_Error(url foundation.IURL, options WorkspaceLaunchOptions, configuration map[WorkspaceLaunchConfigurationKey]objc.IObject, error *foundation.Error) RunningApplication
	// deprecated
	OpenURLs_WithApplicationAtURL_Options_Configuration_Error(urls []foundation.IURL, applicationURL foundation.IURL, options WorkspaceLaunchOptions, configuration map[WorkspaceLaunchConfigurationKey]objc.IObject, error *foundation.Error) RunningApplication
	// deprecated
	OpenFile(fullPath string) bool
	// deprecated
	OpenFile_WithApplication(fullPath string, appName string) bool
	// deprecated
	OpenFile_WithApplication_AndDeactivate(fullPath string, appName string, flag bool) bool
	// deprecated
	OpenFile_FromImage_At_InView(fullPath string, image IImage, point foundation.Point, view IView) bool
	// deprecated
	LaunchApplication(appName string) bool
	// deprecated
	LaunchApplication_ShowIcon_Autolaunch(appName string, showIcon bool, autolaunch bool) bool
	// deprecated
	LaunchApplicationAtURL_Options_Configuration_Error(url foundation.IURL, options WorkspaceLaunchOptions, configuration map[WorkspaceLaunchConfigurationKey]objc.IObject, error *foundation.Error) RunningApplication
	// deprecated
	PerformFileOperation_Source_Destination_Files_Tag(operation WorkspaceFileOperationName, source string, destination string, files []objc.IObject, tag *int) bool
	// deprecated
	FullPathForApplication(appName string) string
	// deprecated
	AbsolutePathForAppBundleWithIdentifier(bundleIdentifier string) string
	// deprecated
	LaunchAppWithBundleIdentifier_Options_AdditionalEventParamDescriptor_LaunchIdentifier(bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifier *foundation.Number) bool
	// deprecated
	OpenURLs_WithAppBundleIdentifier_Options_AdditionalEventParamDescriptor_LaunchIdentifiers(urls []foundation.IURL, bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifiers *foundation.Array) bool
	// deprecated
	OpenTempFile(fullPath string) bool
	// deprecated
	FindApplications()
	// deprecated
	NoteUserDefaultsChanged()
	// deprecated
	SlideImage_From_To(image IImage, fromPoint foundation.Point, toPoint foundation.Point)
	// deprecated
	CheckForRemovableMedia()
	// deprecated
	NoteFileSystemChanged()
	// deprecated
	FileSystemChanged() bool
	// deprecated
	UserDefaultsChanged() bool
	// deprecated
	MountNewRemovableMedia() []objc.Object
	// deprecated
	MountedRemovableMedia() []objc.Object
	// deprecated
	MountedLocalVolumePaths() []objc.Object
	// deprecated
	LaunchedApplications() []objc.Object
	// deprecated
	IconForFileType(fileType string) Image
	NotificationCenter() foundation.NotificationCenter
	FrontmostApplication() RunningApplication
	RunningApplications() []RunningApplication
	MenuBarOwningApplication() RunningApplication
	FileLabels() []string
	FileLabelColors() []Color
	AccessibilityDisplayShouldDifferentiateWithoutColor() bool
	AccessibilityDisplayShouldIncreaseContrast() bool
	AccessibilityDisplayShouldReduceTransparency() bool
	AccessibilityDisplayShouldInvertColors() bool
	AccessibilityDisplayShouldReduceMotion() bool
	IsSwitchControlEnabled() bool
	IsVoiceOverEnabled() bool
}

type Workspace struct {
	objc.Object
}

func MakeWorkspace(ptr unsafe.Pointer) Workspace {
	return Workspace{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WorkspaceClass) Alloc() Workspace {
	rv := ffi.CallMethod[Workspace](wc, "alloc")
	return rv
}

func (w_ Workspace) Init() Workspace {
	rv := ffi.CallMethod[Workspace](w_, "init")
	return rv
}

func (wc _WorkspaceClass) New() Workspace {
	rv := ffi.CallMethod[Workspace](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWorkspace() Workspace {
	return WorkspaceClass.New()
}

func (w_ Workspace) OpenURL(url foundation.IURL) bool {
	rv := ffi.CallMethod[bool](w_, "openURL:", url)
	return rv
}

func (w_ Workspace) HideOtherApplications() {
	ffi.CallMethod[ffi.Void](w_, "hideOtherApplications")
}

func (w_ Workspace) DuplicateURLs_CompletionHandler(URLs []foundation.IURL, handler func(newURLs foundation.Dictionary, error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "duplicateURLs:completionHandler:", URLs, handler)
}

func (w_ Workspace) RecycleURLs_CompletionHandler(URLs []foundation.IURL, handler func(newURLs foundation.Dictionary, error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "recycleURLs:completionHandler:", URLs, handler)
}

func (w_ Workspace) ActivateFileViewerSelectingURLs(fileURLs []foundation.IURL) {
	ffi.CallMethod[ffi.Void](w_, "activateFileViewerSelectingURLs:", fileURLs)
}

func (w_ Workspace) SelectFile_InFileViewerRootedAtPath(fullPath string, rootFullPath string) bool {
	rv := ffi.CallMethod[bool](w_, "selectFile:inFileViewerRootedAtPath:", fullPath, rootFullPath)
	return rv
}

// deprecated
func (w_ Workspace) TypeOfFile_Error(absoluteFilePath string, outError *foundation.Error) string {
	rv := ffi.CallMethod[string](w_, "typeOfFile:error:", absoluteFilePath, unsafe.Pointer(outError))
	return rv
}

// deprecated
func (w_ Workspace) LocalizedDescriptionForType(typeName string) string {
	rv := ffi.CallMethod[string](w_, "localizedDescriptionForType:", typeName)
	return rv
}

// deprecated
func (w_ Workspace) PreferredFilenameExtensionForType(typeName string) string {
	rv := ffi.CallMethod[string](w_, "preferredFilenameExtensionForType:", typeName)
	return rv
}

// deprecated
func (w_ Workspace) FilenameExtension_IsValidForType(filenameExtension string, typeName string) bool {
	rv := ffi.CallMethod[bool](w_, "filenameExtension:isValidForType:", filenameExtension, typeName)
	return rv
}

// deprecated
func (w_ Workspace) Type_ConformsToType(firstTypeName string, secondTypeName string) bool {
	rv := ffi.CallMethod[bool](w_, "type:conformsToType:", firstTypeName, secondTypeName)
	return rv
}

func (w_ Workspace) URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.URL {
	rv := ffi.CallMethod[foundation.URL](w_, "URLForApplicationWithBundleIdentifier:", bundleIdentifier)
	return rv
}

// deprecated
func (w_ Workspace) GetInfoForFile_Application_Type(fullPath string, appName *foundation.String, _type *foundation.String) bool {
	rv := ffi.CallMethod[bool](w_, "getInfoForFile:application:type:", fullPath, appName, _type)
	return rv
}

func (w_ Workspace) URLForApplicationToOpenURL(url foundation.IURL) foundation.URL {
	rv := ffi.CallMethod[foundation.URL](w_, "URLForApplicationToOpenURL:", url)
	return rv
}

func (w_ Workspace) GetFileSystemInfoForPath_IsRemovable_IsWritable_IsUnmountable_Description_Type(fullPath string, removableFlag *bool, writableFlag *bool, unmountableFlag *bool, description *foundation.String, fileSystemType *foundation.String) bool {
	rv := ffi.CallMethod[bool](w_, "getFileSystemInfoForPath:isRemovable:isWritable:isUnmountable:description:type:", fullPath, removableFlag, writableFlag, unmountableFlag, description, fileSystemType)
	return rv
}

func (w_ Workspace) IsFilePackageAtPath(fullPath string) bool {
	rv := ffi.CallMethod[bool](w_, "isFilePackageAtPath:", fullPath)
	return rv
}

func (w_ Workspace) IconForFile(fullPath string) Image {
	rv := ffi.CallMethod[Image](w_, "iconForFile:", fullPath)
	return rv
}

func (w_ Workspace) IconForFiles(fullPaths []string) Image {
	rv := ffi.CallMethod[Image](w_, "iconForFiles:", fullPaths)
	return rv
}

func (w_ Workspace) IconForContentType(contentType uniformtypeidentifiers.IType) Image {
	rv := ffi.CallMethod[Image](w_, "iconForContentType:", contentType)
	return rv
}

func (w_ Workspace) SetIcon_ForFile_Options(image IImage, fullPath string, options WorkspaceIconCreationOptions) bool {
	rv := ffi.CallMethod[bool](w_, "setIcon:forFile:options:", image, fullPath, options)
	return rv
}

func (w_ Workspace) UnmountAndEjectDeviceAtPath(path string) bool {
	rv := ffi.CallMethod[bool](w_, "unmountAndEjectDeviceAtPath:", path)
	return rv
}

func (w_ Workspace) UnmountAndEjectDeviceAtURL_Error(url foundation.IURL, error *foundation.Error) bool {
	rv := ffi.CallMethod[bool](w_, "unmountAndEjectDeviceAtURL:error:", url, unsafe.Pointer(error))
	return rv
}

func (w_ Workspace) DesktopImageURLForScreen(screen IScreen) foundation.URL {
	rv := ffi.CallMethod[foundation.URL](w_, "desktopImageURLForScreen:", screen)
	return rv
}

func (w_ Workspace) SetDesktopImageURL_ForScreen_Options_Error(url foundation.IURL, screen IScreen, options map[WorkspaceDesktopImageOptionKey]objc.IObject, error *foundation.Error) bool {
	rv := ffi.CallMethod[bool](w_, "setDesktopImageURL:forScreen:options:error:", url, screen, options, unsafe.Pointer(error))
	return rv
}

func (w_ Workspace) DesktopImageOptionsForScreen(screen IScreen) map[WorkspaceDesktopImageOptionKey]objc.Object {
	rv := ffi.CallMethod[map[WorkspaceDesktopImageOptionKey]objc.Object](w_, "desktopImageOptionsForScreen:", screen)
	return rv
}

func (w_ Workspace) ShowSearchResultsForQueryString(queryString string) bool {
	rv := ffi.CallMethod[bool](w_, "showSearchResultsForQueryString:", queryString)
	return rv
}

func (w_ Workspace) NoteFileSystemChanged_(path string) {
	ffi.CallMethod[ffi.Void](w_, "noteFileSystemChanged:", path)
}

func (w_ Workspace) ExtendPowerOffBy(requested int) int {
	rv := ffi.CallMethod[int](w_, "extendPowerOffBy:", requested)
	return rv
}

func (w_ Workspace) SetDefaultApplicationAtURL_ToOpenContentType_CompletionHandler(applicationURL foundation.IURL, contentType uniformtypeidentifiers.IType, completionHandler func(error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "setDefaultApplicationAtURL:toOpenContentType:completionHandler:", applicationURL, contentType, completionHandler)
}

func (w_ Workspace) SetDefaultApplicationAtURL_ToOpenContentTypeOfFileAtURL_CompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "setDefaultApplicationAtURL:toOpenContentTypeOfFileAtURL:completionHandler:", applicationURL, url, completionHandler)
}

func (w_ Workspace) SetDefaultApplicationAtURL_ToOpenFileAtURL_CompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "setDefaultApplicationAtURL:toOpenFileAtURL:completionHandler:", applicationURL, url, completionHandler)
}

func (w_ Workspace) SetDefaultApplicationAtURL_ToOpenURLsWithScheme_CompletionHandler(applicationURL foundation.IURL, urlScheme string, completionHandler func(error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "setDefaultApplicationAtURL:toOpenURLsWithScheme:completionHandler:", applicationURL, urlScheme, completionHandler)
}

func (w_ Workspace) URLForApplicationToOpenContentType(contentType uniformtypeidentifiers.IType) foundation.URL {
	rv := ffi.CallMethod[foundation.URL](w_, "URLForApplicationToOpenContentType:", contentType)
	return rv
}

func (w_ Workspace) URLsForApplicationsToOpenContentType(contentType uniformtypeidentifiers.IType) []foundation.URL {
	rv := ffi.CallMethod[[]foundation.URL](w_, "URLsForApplicationsToOpenContentType:", contentType)
	return rv
}

func (w_ Workspace) URLsForApplicationsToOpenURL(url foundation.IURL) []foundation.URL {
	rv := ffi.CallMethod[[]foundation.URL](w_, "URLsForApplicationsToOpenURL:", url)
	return rv
}

func (w_ Workspace) URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.URL {
	rv := ffi.CallMethod[[]foundation.URL](w_, "URLsForApplicationsWithBundleIdentifier:", bundleIdentifier)
	return rv
}

// deprecated
func (w_ Workspace) OpenURL_Options_Configuration_Error(url foundation.IURL, options WorkspaceLaunchOptions, configuration map[WorkspaceLaunchConfigurationKey]objc.IObject, error *foundation.Error) RunningApplication {
	rv := ffi.CallMethod[RunningApplication](w_, "openURL:options:configuration:error:", url, options, configuration, unsafe.Pointer(error))
	return rv
}

// deprecated
func (w_ Workspace) OpenURLs_WithApplicationAtURL_Options_Configuration_Error(urls []foundation.IURL, applicationURL foundation.IURL, options WorkspaceLaunchOptions, configuration map[WorkspaceLaunchConfigurationKey]objc.IObject, error *foundation.Error) RunningApplication {
	rv := ffi.CallMethod[RunningApplication](w_, "openURLs:withApplicationAtURL:options:configuration:error:", urls, applicationURL, options, configuration, unsafe.Pointer(error))
	return rv
}

// deprecated
func (w_ Workspace) OpenFile(fullPath string) bool {
	rv := ffi.CallMethod[bool](w_, "openFile:", fullPath)
	return rv
}

// deprecated
func (w_ Workspace) OpenFile_WithApplication(fullPath string, appName string) bool {
	rv := ffi.CallMethod[bool](w_, "openFile:withApplication:", fullPath, appName)
	return rv
}

// deprecated
func (w_ Workspace) OpenFile_WithApplication_AndDeactivate(fullPath string, appName string, flag bool) bool {
	rv := ffi.CallMethod[bool](w_, "openFile:withApplication:andDeactivate:", fullPath, appName, flag)
	return rv
}

// deprecated
func (w_ Workspace) OpenFile_FromImage_At_InView(fullPath string, image IImage, point foundation.Point, view IView) bool {
	rv := ffi.CallMethod[bool](w_, "openFile:fromImage:at:inView:", fullPath, image, point, view)
	return rv
}

// deprecated
func (w_ Workspace) LaunchApplication(appName string) bool {
	rv := ffi.CallMethod[bool](w_, "launchApplication:", appName)
	return rv
}

// deprecated
func (w_ Workspace) LaunchApplication_ShowIcon_Autolaunch(appName string, showIcon bool, autolaunch bool) bool {
	rv := ffi.CallMethod[bool](w_, "launchApplication:showIcon:autolaunch:", appName, showIcon, autolaunch)
	return rv
}

// deprecated
func (w_ Workspace) LaunchApplicationAtURL_Options_Configuration_Error(url foundation.IURL, options WorkspaceLaunchOptions, configuration map[WorkspaceLaunchConfigurationKey]objc.IObject, error *foundation.Error) RunningApplication {
	rv := ffi.CallMethod[RunningApplication](w_, "launchApplicationAtURL:options:configuration:error:", url, options, configuration, unsafe.Pointer(error))
	return rv
}

// deprecated
func (w_ Workspace) PerformFileOperation_Source_Destination_Files_Tag(operation WorkspaceFileOperationName, source string, destination string, files []objc.IObject, tag *int) bool {
	rv := ffi.CallMethod[bool](w_, "performFileOperation:source:destination:files:tag:", operation, source, destination, files, tag)
	return rv
}

// deprecated
func (w_ Workspace) FullPathForApplication(appName string) string {
	rv := ffi.CallMethod[string](w_, "fullPathForApplication:", appName)
	return rv
}

// deprecated
func (w_ Workspace) AbsolutePathForAppBundleWithIdentifier(bundleIdentifier string) string {
	rv := ffi.CallMethod[string](w_, "absolutePathForAppBundleWithIdentifier:", bundleIdentifier)
	return rv
}

// deprecated
func (w_ Workspace) LaunchAppWithBundleIdentifier_Options_AdditionalEventParamDescriptor_LaunchIdentifier(bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifier *foundation.Number) bool {
	rv := ffi.CallMethod[bool](w_, "launchAppWithBundleIdentifier:options:additionalEventParamDescriptor:launchIdentifier:", bundleIdentifier, options, descriptor, unsafe.Pointer(identifier))
	return rv
}

// deprecated
func (w_ Workspace) OpenURLs_WithAppBundleIdentifier_Options_AdditionalEventParamDescriptor_LaunchIdentifiers(urls []foundation.IURL, bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifiers *foundation.Array) bool {
	rv := ffi.CallMethod[bool](w_, "openURLs:withAppBundleIdentifier:options:additionalEventParamDescriptor:launchIdentifiers:", urls, bundleIdentifier, options, descriptor, identifiers)
	return rv
}

// deprecated
func (w_ Workspace) OpenTempFile(fullPath string) bool {
	rv := ffi.CallMethod[bool](w_, "openTempFile:", fullPath)
	return rv
}

// deprecated
func (w_ Workspace) FindApplications() {
	ffi.CallMethod[ffi.Void](w_, "findApplications")
}

// deprecated
func (w_ Workspace) NoteUserDefaultsChanged() {
	ffi.CallMethod[ffi.Void](w_, "noteUserDefaultsChanged")
}

// deprecated
func (w_ Workspace) SlideImage_From_To(image IImage, fromPoint foundation.Point, toPoint foundation.Point) {
	ffi.CallMethod[ffi.Void](w_, "slideImage:from:to:", image, fromPoint, toPoint)
}

// deprecated
func (w_ Workspace) CheckForRemovableMedia() {
	ffi.CallMethod[ffi.Void](w_, "checkForRemovableMedia")
}

// deprecated
func (w_ Workspace) NoteFileSystemChanged() {
	ffi.CallMethod[ffi.Void](w_, "noteFileSystemChanged")
}

// deprecated
func (w_ Workspace) FileSystemChanged() bool {
	rv := ffi.CallMethod[bool](w_, "fileSystemChanged")
	return rv
}

// deprecated
func (w_ Workspace) UserDefaultsChanged() bool {
	rv := ffi.CallMethod[bool](w_, "userDefaultsChanged")
	return rv
}

// deprecated
func (w_ Workspace) MountNewRemovableMedia() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](w_, "mountNewRemovableMedia")
	return rv
}

// deprecated
func (w_ Workspace) MountedRemovableMedia() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](w_, "mountedRemovableMedia")
	return rv
}

// deprecated
func (w_ Workspace) MountedLocalVolumePaths() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](w_, "mountedLocalVolumePaths")
	return rv
}

// deprecated
func (w_ Workspace) LaunchedApplications() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](w_, "launchedApplications")
	return rv
}

// deprecated
func (w_ Workspace) IconForFileType(fileType string) Image {
	rv := ffi.CallMethod[Image](w_, "iconForFileType:", fileType)
	return rv
}

func (wc _WorkspaceClass) SharedWorkspace() Workspace {
	rv := ffi.CallMethod[Workspace](wc, "sharedWorkspace")
	return rv
}

func (w_ Workspace) NotificationCenter() foundation.NotificationCenter {
	rv := ffi.CallMethod[foundation.NotificationCenter](w_, "notificationCenter")
	return rv
}

func (w_ Workspace) FrontmostApplication() RunningApplication {
	rv := ffi.CallMethod[RunningApplication](w_, "frontmostApplication")
	return rv
}

func (w_ Workspace) RunningApplications() []RunningApplication {
	rv := ffi.CallMethod[[]RunningApplication](w_, "runningApplications")
	return rv
}

func (w_ Workspace) MenuBarOwningApplication() RunningApplication {
	rv := ffi.CallMethod[RunningApplication](w_, "menuBarOwningApplication")
	return rv
}

func (w_ Workspace) FileLabels() []string {
	rv := ffi.CallMethod[[]string](w_, "fileLabels")
	return rv
}

func (w_ Workspace) FileLabelColors() []Color {
	rv := ffi.CallMethod[[]Color](w_, "fileLabelColors")
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	rv := ffi.CallMethod[bool](w_, "accessibilityDisplayShouldDifferentiateWithoutColor")
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	rv := ffi.CallMethod[bool](w_, "accessibilityDisplayShouldIncreaseContrast")
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldReduceTransparency() bool {
	rv := ffi.CallMethod[bool](w_, "accessibilityDisplayShouldReduceTransparency")
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldInvertColors() bool {
	rv := ffi.CallMethod[bool](w_, "accessibilityDisplayShouldInvertColors")
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldReduceMotion() bool {
	rv := ffi.CallMethod[bool](w_, "accessibilityDisplayShouldReduceMotion")
	return rv
}

func (w_ Workspace) IsSwitchControlEnabled() bool {
	rv := ffi.CallMethod[bool](w_, "isSwitchControlEnabled")
	return rv
}

func (w_ Workspace) IsVoiceOverEnabled() bool {
	rv := ffi.CallMethod[bool](w_, "isVoiceOverEnabled")
	return rv
}
