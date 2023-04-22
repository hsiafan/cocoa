// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	GetInfoForFile_Application_Type(fullPath string, appName *foundation.String, type_ *foundation.String) bool
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
	NoteFileSystemChanged1(path string)
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
	rv := objc.CallMethod[Workspace](wc, objc.GetSelector("alloc"))
	return rv
}

func (wc _WorkspaceClass) New() Workspace {
	rv := objc.CallMethod[Workspace](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWorkspace() Workspace {
	return WorkspaceClass.New()
}

func (w_ Workspace) Init() Workspace {
	rv := objc.CallMethod[Workspace](w_, objc.GetSelector("init"))
	return rv
}

func (w_ Workspace) OpenURL(url foundation.IURL) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("openURL:"), url)
	return rv
}

func (w_ Workspace) HideOtherApplications() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("hideOtherApplications"))
}

func (w_ Workspace) ActivateFileViewerSelectingURLs(fileURLs []foundation.IURL) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("activateFileViewerSelectingURLs:"), fileURLs)
}

func (w_ Workspace) SelectFile_InFileViewerRootedAtPath(fullPath string, rootFullPath string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("selectFile:inFileViewerRootedAtPath:"), fullPath, rootFullPath)
	return rv
}

// deprecated
func (w_ Workspace) TypeOfFile_Error(absoluteFilePath string, outError *foundation.Error) string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("typeOfFile:error:"), absoluteFilePath, unsafe.Pointer(outError))
	return rv
}

// deprecated
func (w_ Workspace) LocalizedDescriptionForType(typeName string) string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("localizedDescriptionForType:"), typeName)
	return rv
}

// deprecated
func (w_ Workspace) PreferredFilenameExtensionForType(typeName string) string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("preferredFilenameExtensionForType:"), typeName)
	return rv
}

// deprecated
func (w_ Workspace) FilenameExtension_IsValidForType(filenameExtension string, typeName string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("filenameExtension:isValidForType:"), filenameExtension, typeName)
	return rv
}

// deprecated
func (w_ Workspace) Type_ConformsToType(firstTypeName string, secondTypeName string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("type:conformsToType:"), firstTypeName, secondTypeName)
	return rv
}

func (w_ Workspace) URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("URLForApplicationWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

// deprecated
func (w_ Workspace) GetInfoForFile_Application_Type(fullPath string, appName *foundation.String, type_ *foundation.String) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("getInfoForFile:application:type:"), fullPath, appName, type_)
	return rv
}

func (w_ Workspace) URLForApplicationToOpenURL(url foundation.IURL) foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("URLForApplicationToOpenURL:"), url)
	return rv
}

func (w_ Workspace) GetFileSystemInfoForPath_IsRemovable_IsWritable_IsUnmountable_Description_Type(fullPath string, removableFlag *bool, writableFlag *bool, unmountableFlag *bool, description *foundation.String, fileSystemType *foundation.String) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("getFileSystemInfoForPath:isRemovable:isWritable:isUnmountable:description:type:"), fullPath, removableFlag, writableFlag, unmountableFlag, description, fileSystemType)
	return rv
}

func (w_ Workspace) IsFilePackageAtPath(fullPath string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isFilePackageAtPath:"), fullPath)
	return rv
}

func (w_ Workspace) IconForFile(fullPath string) Image {
	rv := objc.CallMethod[Image](w_, objc.GetSelector("iconForFile:"), fullPath)
	return rv
}

func (w_ Workspace) IconForFiles(fullPaths []string) Image {
	rv := objc.CallMethod[Image](w_, objc.GetSelector("iconForFiles:"), fullPaths)
	return rv
}

func (w_ Workspace) IconForContentType(contentType uniformtypeidentifiers.IType) Image {
	rv := objc.CallMethod[Image](w_, objc.GetSelector("iconForContentType:"), contentType)
	return rv
}

func (w_ Workspace) SetIcon_ForFile_Options(image IImage, fullPath string, options WorkspaceIconCreationOptions) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("setIcon:forFile:options:"), image, fullPath, options)
	return rv
}

func (w_ Workspace) UnmountAndEjectDeviceAtPath(path string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("unmountAndEjectDeviceAtPath:"), path)
	return rv
}

func (w_ Workspace) UnmountAndEjectDeviceAtURL_Error(url foundation.IURL, error *foundation.Error) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("unmountAndEjectDeviceAtURL:error:"), url, unsafe.Pointer(error))
	return rv
}

func (w_ Workspace) DesktopImageURLForScreen(screen IScreen) foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("desktopImageURLForScreen:"), screen)
	return rv
}

func (w_ Workspace) SetDesktopImageURL_ForScreen_Options_Error(url foundation.IURL, screen IScreen, options map[WorkspaceDesktopImageOptionKey]objc.IObject, error *foundation.Error) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("setDesktopImageURL:forScreen:options:error:"), url, screen, options, unsafe.Pointer(error))
	return rv
}

func (w_ Workspace) DesktopImageOptionsForScreen(screen IScreen) map[WorkspaceDesktopImageOptionKey]objc.Object {
	rv := objc.CallMethod[map[WorkspaceDesktopImageOptionKey]objc.Object](w_, objc.GetSelector("desktopImageOptionsForScreen:"), screen)
	return rv
}

func (w_ Workspace) ShowSearchResultsForQueryString(queryString string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("showSearchResultsForQueryString:"), queryString)
	return rv
}

func (w_ Workspace) NoteFileSystemChanged1(path string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("noteFileSystemChanged:"), path)
}

func (w_ Workspace) ExtendPowerOffBy(requested int) int {
	rv := objc.CallMethod[int](w_, objc.GetSelector("extendPowerOffBy:"), requested)
	return rv
}

func (w_ Workspace) SetDefaultApplicationAtURL_ToOpenContentType_CompletionHandler(applicationURL foundation.IURL, contentType uniformtypeidentifiers.IType, completionHandler func(error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultApplicationAtURL:toOpenContentType:completionHandler:"), applicationURL, contentType, completionHandler)
}

func (w_ Workspace) SetDefaultApplicationAtURL_ToOpenContentTypeOfFileAtURL_CompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultApplicationAtURL:toOpenContentTypeOfFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}

func (w_ Workspace) SetDefaultApplicationAtURL_ToOpenFileAtURL_CompletionHandler(applicationURL foundation.IURL, url foundation.IURL, completionHandler func(error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultApplicationAtURL:toOpenFileAtURL:completionHandler:"), applicationURL, url, completionHandler)
}

func (w_ Workspace) SetDefaultApplicationAtURL_ToOpenURLsWithScheme_CompletionHandler(applicationURL foundation.IURL, urlScheme string, completionHandler func(error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultApplicationAtURL:toOpenURLsWithScheme:completionHandler:"), applicationURL, urlScheme, completionHandler)
}

func (w_ Workspace) URLForApplicationToOpenContentType(contentType uniformtypeidentifiers.IType) foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("URLForApplicationToOpenContentType:"), contentType)
	return rv
}

func (w_ Workspace) URLsForApplicationsToOpenContentType(contentType uniformtypeidentifiers.IType) []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](w_, objc.GetSelector("URLsForApplicationsToOpenContentType:"), contentType)
	return rv
}

func (w_ Workspace) URLsForApplicationsToOpenURL(url foundation.IURL) []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](w_, objc.GetSelector("URLsForApplicationsToOpenURL:"), url)
	return rv
}

func (w_ Workspace) URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](w_, objc.GetSelector("URLsForApplicationsWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

// deprecated
func (w_ Workspace) OpenURL_Options_Configuration_Error(url foundation.IURL, options WorkspaceLaunchOptions, configuration map[WorkspaceLaunchConfigurationKey]objc.IObject, error *foundation.Error) RunningApplication {
	rv := objc.CallMethod[RunningApplication](w_, objc.GetSelector("openURL:options:configuration:error:"), url, options, configuration, unsafe.Pointer(error))
	return rv
}

// deprecated
func (w_ Workspace) OpenURLs_WithApplicationAtURL_Options_Configuration_Error(urls []foundation.IURL, applicationURL foundation.IURL, options WorkspaceLaunchOptions, configuration map[WorkspaceLaunchConfigurationKey]objc.IObject, error *foundation.Error) RunningApplication {
	rv := objc.CallMethod[RunningApplication](w_, objc.GetSelector("openURLs:withApplicationAtURL:options:configuration:error:"), urls, applicationURL, options, configuration, unsafe.Pointer(error))
	return rv
}

// deprecated
func (w_ Workspace) OpenFile(fullPath string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("openFile:"), fullPath)
	return rv
}

// deprecated
func (w_ Workspace) OpenFile_WithApplication(fullPath string, appName string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("openFile:withApplication:"), fullPath, appName)
	return rv
}

// deprecated
func (w_ Workspace) OpenFile_WithApplication_AndDeactivate(fullPath string, appName string, flag bool) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("openFile:withApplication:andDeactivate:"), fullPath, appName, flag)
	return rv
}

// deprecated
func (w_ Workspace) OpenFile_FromImage_At_InView(fullPath string, image IImage, point foundation.Point, view IView) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("openFile:fromImage:at:inView:"), fullPath, image, point, view)
	return rv
}

// deprecated
func (w_ Workspace) LaunchApplication(appName string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("launchApplication:"), appName)
	return rv
}

// deprecated
func (w_ Workspace) LaunchApplication_ShowIcon_Autolaunch(appName string, showIcon bool, autolaunch bool) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("launchApplication:showIcon:autolaunch:"), appName, showIcon, autolaunch)
	return rv
}

// deprecated
func (w_ Workspace) LaunchApplicationAtURL_Options_Configuration_Error(url foundation.IURL, options WorkspaceLaunchOptions, configuration map[WorkspaceLaunchConfigurationKey]objc.IObject, error *foundation.Error) RunningApplication {
	rv := objc.CallMethod[RunningApplication](w_, objc.GetSelector("launchApplicationAtURL:options:configuration:error:"), url, options, configuration, unsafe.Pointer(error))
	return rv
}

// deprecated
func (w_ Workspace) PerformFileOperation_Source_Destination_Files_Tag(operation WorkspaceFileOperationName, source string, destination string, files []objc.IObject, tag *int) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("performFileOperation:source:destination:files:tag:"), operation, source, destination, files, tag)
	return rv
}

// deprecated
func (w_ Workspace) FullPathForApplication(appName string) string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("fullPathForApplication:"), appName)
	return rv
}

// deprecated
func (w_ Workspace) AbsolutePathForAppBundleWithIdentifier(bundleIdentifier string) string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("absolutePathForAppBundleWithIdentifier:"), bundleIdentifier)
	return rv
}

// deprecated
func (w_ Workspace) LaunchAppWithBundleIdentifier_Options_AdditionalEventParamDescriptor_LaunchIdentifier(bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifier *foundation.Number) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("launchAppWithBundleIdentifier:options:additionalEventParamDescriptor:launchIdentifier:"), bundleIdentifier, options, descriptor, unsafe.Pointer(identifier))
	return rv
}

// deprecated
func (w_ Workspace) OpenURLs_WithAppBundleIdentifier_Options_AdditionalEventParamDescriptor_LaunchIdentifiers(urls []foundation.IURL, bundleIdentifier string, options WorkspaceLaunchOptions, descriptor foundation.IAppleEventDescriptor, identifiers *foundation.Array) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("openURLs:withAppBundleIdentifier:options:additionalEventParamDescriptor:launchIdentifiers:"), urls, bundleIdentifier, options, descriptor, identifiers)
	return rv
}

// deprecated
func (w_ Workspace) OpenTempFile(fullPath string) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("openTempFile:"), fullPath)
	return rv
}

// deprecated
func (w_ Workspace) FindApplications() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("findApplications"))
}

// deprecated
func (w_ Workspace) NoteUserDefaultsChanged() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("noteUserDefaultsChanged"))
}

// deprecated
func (w_ Workspace) SlideImage_From_To(image IImage, fromPoint foundation.Point, toPoint foundation.Point) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("slideImage:from:to:"), image, fromPoint, toPoint)
}

// deprecated
func (w_ Workspace) CheckForRemovableMedia() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("checkForRemovableMedia"))
}

// deprecated
func (w_ Workspace) NoteFileSystemChanged() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("noteFileSystemChanged"))
}

// deprecated
func (w_ Workspace) FileSystemChanged() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("fileSystemChanged"))
	return rv
}

// deprecated
func (w_ Workspace) UserDefaultsChanged() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("userDefaultsChanged"))
	return rv
}

// deprecated
func (w_ Workspace) MountNewRemovableMedia() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](w_, objc.GetSelector("mountNewRemovableMedia"))
	return rv
}

// deprecated
func (w_ Workspace) MountedRemovableMedia() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](w_, objc.GetSelector("mountedRemovableMedia"))
	return rv
}

// deprecated
func (w_ Workspace) MountedLocalVolumePaths() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](w_, objc.GetSelector("mountedLocalVolumePaths"))
	return rv
}

// deprecated
func (w_ Workspace) LaunchedApplications() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](w_, objc.GetSelector("launchedApplications"))
	return rv
}

// deprecated
func (w_ Workspace) IconForFileType(fileType string) Image {
	rv := objc.CallMethod[Image](w_, objc.GetSelector("iconForFileType:"), fileType)
	return rv
}

func (wc _WorkspaceClass) SharedWorkspace() Workspace {
	rv := objc.CallMethod[Workspace](wc, objc.GetSelector("sharedWorkspace"))
	return rv
}

func (w_ Workspace) NotificationCenter() foundation.NotificationCenter {
	rv := objc.CallMethod[foundation.NotificationCenter](w_, objc.GetSelector("notificationCenter"))
	return rv
}

func (w_ Workspace) FrontmostApplication() RunningApplication {
	rv := objc.CallMethod[RunningApplication](w_, objc.GetSelector("frontmostApplication"))
	return rv
}

func (w_ Workspace) RunningApplications() []RunningApplication {
	rv := objc.CallMethod[[]RunningApplication](w_, objc.GetSelector("runningApplications"))
	return rv
}

func (w_ Workspace) MenuBarOwningApplication() RunningApplication {
	rv := objc.CallMethod[RunningApplication](w_, objc.GetSelector("menuBarOwningApplication"))
	return rv
}

func (w_ Workspace) FileLabels() []string {
	rv := objc.CallMethod[[]string](w_, objc.GetSelector("fileLabels"))
	return rv
}

func (w_ Workspace) FileLabelColors() []Color {
	rv := objc.CallMethod[[]Color](w_, objc.GetSelector("fileLabelColors"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldDifferentiateWithoutColor"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldIncreaseContrast"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldReduceTransparency() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldReduceTransparency"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldInvertColors() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldInvertColors"))
	return rv
}

func (w_ Workspace) AccessibilityDisplayShouldReduceMotion() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("accessibilityDisplayShouldReduceMotion"))
	return rv
}

func (w_ Workspace) IsSwitchControlEnabled() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isSwitchControlEnabled"))
	return rv
}

func (w_ Workspace) IsVoiceOverEnabled() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isVoiceOverEnabled"))
	return rv
}
