// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var ApplicationClass = _ApplicationClass{objc.GetClass("NSApplication")}

type _ApplicationClass struct {
	objc.Class
}

type IApplication interface {
	IResponder
	NextEventMatchingMask_UntilDate_InMode_Dequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event
	DiscardEventsMatchingMask_BeforeEvent(mask EventMask, lastEvent IEvent)
	Run()
	FinishLaunching()
	Stop(sender objc.IObject)
	SendEvent(event IEvent)
	PostEvent_AtStart(event IEvent, flag bool)
	SendAction_To_From(action objc.Selector, target objc.IObject, sender objc.IObject) bool
	TargetForAction(action objc.Selector) objc.Object
	TargetForAction_To_From(action objc.Selector, target objc.IObject, sender objc.IObject) objc.Object
	Terminate(sender objc.IObject)
	ReplyToApplicationShouldTerminate(shouldTerminate bool)
	ActivateIgnoringOtherApps(flag bool)
	Deactivate()
	DisableRelaunchOnLogin()
	EnableRelaunchOnLogin()
	RegisterForRemoteNotifications()
	UnregisterForRemoteNotifications()
	RegisterForRemoteNotificationTypes(types RemoteNotificationType)
	ToggleTouchBarCustomizationPalette(sender objc.IObject)
	RequestUserAttention(requestType RequestUserAttentionType) int
	CancelUserAttentionRequest(request int)
	ReplyToOpenOrPrint(reply ApplicationDelegateReply)
	SearchString_InUserInterfaceItemString_SearchRange_FoundRange(searchString string, stringToSearch string, searchRange foundation.Range, foundRange *foundation.Range) bool
	ShowHelp(sender objc.IObject)
	ActivateContextHelpMode(sender objc.IObject)
	HideOtherApplications(sender objc.IObject)
	UnhideAllApplications(sender objc.IObject)
	ReportException(exception foundation.IException)
	ActivationPolicy() ApplicationActivationPolicy
	SetActivationPolicy(activationPolicy ApplicationActivationPolicy) bool
	WindowWithWindowNumber(windowNum int) Window
	// deprecated
	MakeWindowsPerform_InOrder(selector objc.Selector, flag bool) Window
	EnumerateWindowsWithOptions_UsingBlock(options WindowListOptions, block func(window Window, stop *bool))
	MiniaturizeAll(sender objc.IObject)
	Hide(sender objc.IObject)
	Unhide(sender objc.IObject)
	UnhideWithoutActivation()
	UpdateWindows()
	SetWindowsNeedUpdate(needUpdate bool)
	PreventWindowOrdering()
	ArrangeInFront(sender objc.IObject)
	ExtendStateRestoration()
	CompleteStateRestoration()
	RunModalForWindow(window IWindow) ModalResponse
	StopModal()
	StopModalWithCode(returnCode ModalResponse)
	AbortModal()
	BeginModalSessionForWindow(window IWindow) ModalSession
	RunModalSession(session ModalSession) ModalResponse
	OrderFrontColorPanel(sender objc.IObject)
	OrderFrontStandardAboutPanel(sender objc.IObject)
	OrderFrontStandardAboutPanelWithOptions(optionsDictionary map[AboutPanelOptionKey]objc.IObject)
	OrderFrontCharacterPalette(sender objc.IObject)
	RunPageLayout(sender objc.IObject)
	AddWindowsItem_Title_Filename(win IWindow, _string string, isFilename bool)
	ChangeWindowsItem_Title_Filename(win IWindow, _string string, isFilename bool)
	RemoveWindowsItem(win IWindow)
	UpdateWindowsItem(win IWindow)
	RegisterServicesMenuSendTypes_ReturnTypes(sendTypes []PasteboardType, returnTypes []PasteboardType)
	// deprecated
	BeginModalSessionForWindow_RelativeToWindow(window IWindow, docWindow IWindow) ModalSession
	// deprecated
	RunModalForWindow_RelativeToWindow(window IWindow, docWindow IWindow) int
	EndModalSession(session ModalSession)
	// deprecated
	BeginSheet_ModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(sheet IWindow, docWindow IWindow, modalDelegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	// deprecated
	EndSheet(sheet IWindow)
	// deprecated
	EndSheet_ReturnCode(sheet IWindow, returnCode int)
	// deprecated
	Application_PrintFiles(sender IApplication, filenames []string)
	// deprecated
	Application_DelegateHandlesKey(sender IApplication, key string) bool
	Delegate() ApplicationDelegateWrapper
	SetDelegate(value ApplicationDelegate)
	CurrentEvent() Event
	IsRunning() bool
	IsActive() bool
	EnabledRemoteNotificationTypes() RemoteNotificationType
	IsRegisteredForRemoteNotifications() bool
	Appearance() Appearance
	SetAppearance(value IAppearance)
	EffectiveAppearance() Appearance
	CurrentSystemPresentationOptions() ApplicationPresentationOptions
	PresentationOptions() ApplicationPresentationOptions
	SetPresentationOptions(value ApplicationPresentationOptions)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	DockTile() DockTile
	ApplicationIconImage() Image
	SetApplicationIconImage(value IImage)
	HelpMenu() Menu
	SetHelpMenu(value IMenu)
	ServicesProvider() objc.Object
	SetServicesProvider(value objc.IObject)
	IsFullKeyboardAccessEnabled() bool
	OrderedDocuments() []Document
	OrderedWindows() []Window
	KeyWindow() Window
	MainWindow() Window
	Windows() []Window
	IsHidden() bool
	// deprecated
	Context() GraphicsContext
	OcclusionState() ApplicationOcclusionState
	IsProtectedDataAvailable() bool
	ModalWindow() Window
	MainMenu() Menu
	SetMainMenu(value IMenu)
	IsAutomaticCustomizeTouchBarMenuItemEnabled() bool
	SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool)
	WindowsMenu() Menu
	SetWindowsMenu(value IMenu)
	ServicesMenu() Menu
	SetServicesMenu(value IMenu)
}

type Application struct {
	Responder
}

func MakeApplication(ptr unsafe.Pointer) Application {
	return Application{
		Responder: MakeResponder(ptr),
	}
}

func (a_ Application) Init() Application {
	rv := ffi.CallMethod[Application](a_, "init")
	rv.Autorelease()
	return rv
}

func (ac _ApplicationClass) Alloc() Application {
	rv := ffi.CallMethod[Application](ac, "alloc")
	return rv
}

func (ac _ApplicationClass) New() Application {
	rv := ffi.CallMethod[Application](ac, "new")
	rv.Autorelease()
	return rv
}

func NewApplication() Application {
	return ApplicationClass.New()
}

func (a_ Application) NextEventMatchingMask_UntilDate_InMode_Dequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event {
	rv := ffi.CallMethod[Event](a_, "nextEventMatchingMask:untilDate:inMode:dequeue:", mask, expiration, mode, deqFlag)
	return rv
}

func (a_ Application) DiscardEventsMatchingMask_BeforeEvent(mask EventMask, lastEvent IEvent) {
	ffi.CallMethod[ffi.Void](a_, "discardEventsMatchingMask:beforeEvent:", mask, lastEvent)
}

func (a_ Application) Run() {
	ffi.CallMethod[ffi.Void](a_, "run")
}

func (a_ Application) FinishLaunching() {
	ffi.CallMethod[ffi.Void](a_, "finishLaunching")
}

func (a_ Application) Stop(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "stop:", sender)
}

func (a_ Application) SendEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](a_, "sendEvent:", event)
}

func (a_ Application) PostEvent_AtStart(event IEvent, flag bool) {
	ffi.CallMethod[ffi.Void](a_, "postEvent:atStart:", event, flag)
}

func (a_ Application) SendAction_To_From(action objc.Selector, target objc.IObject, sender objc.IObject) bool {
	rv := ffi.CallMethod[bool](a_, "sendAction:to:from:", action, target, sender)
	return rv
}

func (a_ Application) TargetForAction(action objc.Selector) objc.Object {
	rv := ffi.CallMethod[objc.Object](a_, "targetForAction:", action)
	return rv
}

func (a_ Application) TargetForAction_To_From(action objc.Selector, target objc.IObject, sender objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](a_, "targetForAction:to:from:", action, target, sender)
	return rv
}

func (a_ Application) Terminate(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "terminate:", sender)
}

func (a_ Application) ReplyToApplicationShouldTerminate(shouldTerminate bool) {
	ffi.CallMethod[ffi.Void](a_, "replyToApplicationShouldTerminate:", shouldTerminate)
}

func (a_ Application) ActivateIgnoringOtherApps(flag bool) {
	ffi.CallMethod[ffi.Void](a_, "activateIgnoringOtherApps:", flag)
}

func (a_ Application) Deactivate() {
	ffi.CallMethod[ffi.Void](a_, "deactivate")
}

func (a_ Application) DisableRelaunchOnLogin() {
	ffi.CallMethod[ffi.Void](a_, "disableRelaunchOnLogin")
}

func (a_ Application) EnableRelaunchOnLogin() {
	ffi.CallMethod[ffi.Void](a_, "enableRelaunchOnLogin")
}

func (a_ Application) RegisterForRemoteNotifications() {
	ffi.CallMethod[ffi.Void](a_, "registerForRemoteNotifications")
}

func (a_ Application) UnregisterForRemoteNotifications() {
	ffi.CallMethod[ffi.Void](a_, "unregisterForRemoteNotifications")
}

func (a_ Application) RegisterForRemoteNotificationTypes(types RemoteNotificationType) {
	ffi.CallMethod[ffi.Void](a_, "registerForRemoteNotificationTypes:", types)
}

func (a_ Application) ToggleTouchBarCustomizationPalette(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "toggleTouchBarCustomizationPalette:", sender)
}

func (a_ Application) RequestUserAttention(requestType RequestUserAttentionType) int {
	rv := ffi.CallMethod[int](a_, "requestUserAttention:", requestType)
	return rv
}

func (a_ Application) CancelUserAttentionRequest(request int) {
	ffi.CallMethod[ffi.Void](a_, "cancelUserAttentionRequest:", request)
}

func (a_ Application) ReplyToOpenOrPrint(reply ApplicationDelegateReply) {
	ffi.CallMethod[ffi.Void](a_, "replyToOpenOrPrint:", reply)
}

func (a_ Application) SearchString_InUserInterfaceItemString_SearchRange_FoundRange(searchString string, stringToSearch string, searchRange foundation.Range, foundRange *foundation.Range) bool {
	rv := ffi.CallMethod[bool](a_, "searchString:inUserInterfaceItemString:searchRange:foundRange:", searchString, stringToSearch, searchRange, foundRange)
	return rv
}

func (a_ Application) ShowHelp(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "showHelp:", sender)
}

func (a_ Application) ActivateContextHelpMode(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "activateContextHelpMode:", sender)
}

func (a_ Application) HideOtherApplications(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "hideOtherApplications:", sender)
}

func (a_ Application) UnhideAllApplications(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "unhideAllApplications:", sender)
}

func (ac _ApplicationClass) DetachDrawingThread_ToTarget_WithObject(selector objc.Selector, target objc.IObject, argument objc.IObject) {
	ffi.CallMethod[ffi.Void](ac, "detachDrawingThread:toTarget:withObject:", selector, target, argument)
}

func (a_ Application) ReportException(exception foundation.IException) {
	ffi.CallMethod[ffi.Void](a_, "reportException:", exception)
}

func (a_ Application) ActivationPolicy() ApplicationActivationPolicy {
	rv := ffi.CallMethod[ApplicationActivationPolicy](a_, "activationPolicy")
	return rv
}

func (a_ Application) SetActivationPolicy(activationPolicy ApplicationActivationPolicy) bool {
	rv := ffi.CallMethod[bool](a_, "setActivationPolicy:", activationPolicy)
	return rv
}

func (a_ Application) WindowWithWindowNumber(windowNum int) Window {
	rv := ffi.CallMethod[Window](a_, "windowWithWindowNumber:", windowNum)
	return rv
}

// deprecated
func (a_ Application) MakeWindowsPerform_InOrder(selector objc.Selector, flag bool) Window {
	rv := ffi.CallMethod[Window](a_, "makeWindowsPerform:inOrder:", selector, flag)
	return rv
}

func (a_ Application) EnumerateWindowsWithOptions_UsingBlock(options WindowListOptions, block func(window Window, stop *bool)) {
	ffi.CallMethod[ffi.Void](a_, "enumerateWindowsWithOptions:usingBlock:", options, block)
}

func (a_ Application) MiniaturizeAll(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "miniaturizeAll:", sender)
}

func (a_ Application) Hide(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "hide:", sender)
}

func (a_ Application) Unhide(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "unhide:", sender)
}

func (a_ Application) UnhideWithoutActivation() {
	ffi.CallMethod[ffi.Void](a_, "unhideWithoutActivation")
}

func (a_ Application) UpdateWindows() {
	ffi.CallMethod[ffi.Void](a_, "updateWindows")
}

func (a_ Application) SetWindowsNeedUpdate(needUpdate bool) {
	ffi.CallMethod[ffi.Void](a_, "setWindowsNeedUpdate:", needUpdate)
}

func (a_ Application) PreventWindowOrdering() {
	ffi.CallMethod[ffi.Void](a_, "preventWindowOrdering")
}

func (a_ Application) ArrangeInFront(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "arrangeInFront:", sender)
}

func (a_ Application) ExtendStateRestoration() {
	ffi.CallMethod[ffi.Void](a_, "extendStateRestoration")
}

func (a_ Application) CompleteStateRestoration() {
	ffi.CallMethod[ffi.Void](a_, "completeStateRestoration")
}

func (a_ Application) RunModalForWindow(window IWindow) ModalResponse {
	rv := ffi.CallMethod[ModalResponse](a_, "runModalForWindow:", window)
	return rv
}

func (a_ Application) StopModal() {
	ffi.CallMethod[ffi.Void](a_, "stopModal")
}

func (a_ Application) StopModalWithCode(returnCode ModalResponse) {
	ffi.CallMethod[ffi.Void](a_, "stopModalWithCode:", returnCode)
}

func (a_ Application) AbortModal() {
	ffi.CallMethod[ffi.Void](a_, "abortModal")
}

func (a_ Application) BeginModalSessionForWindow(window IWindow) ModalSession {
	rv := ffi.CallMethod[ModalSession](a_, "beginModalSessionForWindow:", window)
	return rv
}

func (a_ Application) RunModalSession(session ModalSession) ModalResponse {
	rv := ffi.CallMethod[ModalResponse](a_, "runModalSession:", session)
	return rv
}

func (a_ Application) OrderFrontColorPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "orderFrontColorPanel:", sender)
}

func (a_ Application) OrderFrontStandardAboutPanel(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "orderFrontStandardAboutPanel:", sender)
}

func (a_ Application) OrderFrontStandardAboutPanelWithOptions(optionsDictionary map[AboutPanelOptionKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "orderFrontStandardAboutPanelWithOptions:", optionsDictionary)
}

func (a_ Application) OrderFrontCharacterPalette(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "orderFrontCharacterPalette:", sender)
}

func (a_ Application) RunPageLayout(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "runPageLayout:", sender)
}

func (a_ Application) AddWindowsItem_Title_Filename(win IWindow, _string string, isFilename bool) {
	ffi.CallMethod[ffi.Void](a_, "addWindowsItem:title:filename:", win, _string, isFilename)
}

func (a_ Application) ChangeWindowsItem_Title_Filename(win IWindow, _string string, isFilename bool) {
	ffi.CallMethod[ffi.Void](a_, "changeWindowsItem:title:filename:", win, _string, isFilename)
}

func (a_ Application) RemoveWindowsItem(win IWindow) {
	ffi.CallMethod[ffi.Void](a_, "removeWindowsItem:", win)
}

func (a_ Application) UpdateWindowsItem(win IWindow) {
	ffi.CallMethod[ffi.Void](a_, "updateWindowsItem:", win)
}

func (a_ Application) RegisterServicesMenuSendTypes_ReturnTypes(sendTypes []PasteboardType, returnTypes []PasteboardType) {
	ffi.CallMethod[ffi.Void](a_, "registerServicesMenuSendTypes:returnTypes:", sendTypes, returnTypes)
}

// deprecated
func (a_ Application) BeginModalSessionForWindow_RelativeToWindow(window IWindow, docWindow IWindow) ModalSession {
	rv := ffi.CallMethod[ModalSession](a_, "beginModalSessionForWindow:relativeToWindow:", window, docWindow)
	return rv
}

// deprecated
func (a_ Application) RunModalForWindow_RelativeToWindow(window IWindow, docWindow IWindow) int {
	rv := ffi.CallMethod[int](a_, "runModalForWindow:relativeToWindow:", window, docWindow)
	return rv
}

func (a_ Application) EndModalSession(session ModalSession) {
	ffi.CallMethod[ffi.Void](a_, "endModalSession:", session)
}

// deprecated
func (a_ Application) BeginSheet_ModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(sheet IWindow, docWindow IWindow, modalDelegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	ffi.CallMethod[ffi.Void](a_, "beginSheet:modalForWindow:modalDelegate:didEndSelector:contextInfo:", sheet, docWindow, modalDelegate, didEndSelector, contextInfo)
}

// deprecated
func (a_ Application) EndSheet(sheet IWindow) {
	ffi.CallMethod[ffi.Void](a_, "endSheet:", sheet)
}

// deprecated
func (a_ Application) EndSheet_ReturnCode(sheet IWindow, returnCode int) {
	ffi.CallMethod[ffi.Void](a_, "endSheet:returnCode:", sheet, returnCode)
}

// deprecated
func (a_ Application) Application_PrintFiles(sender IApplication, filenames []string) {
	ffi.CallMethod[ffi.Void](a_, "application:printFiles:", sender, filenames)
}

// deprecated
func (a_ Application) Application_DelegateHandlesKey(sender IApplication, key string) bool {
	rv := ffi.CallMethod[bool](a_, "application:delegateHandlesKey:", sender, key)
	return rv
}

func (ac _ApplicationClass) SharedApplication() Application {
	rv := ffi.CallMethod[Application](ac, "sharedApplication")
	return rv
}

func (a_ Application) Delegate() ApplicationDelegateWrapper {
	rv := ffi.CallMethod[ApplicationDelegateWrapper](a_, "delegate")
	return rv
}

func (a_ Application) SetDelegate(value ApplicationDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(a_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](a_, "setDelegate:", po)
}

func (a_ Application) CurrentEvent() Event {
	rv := ffi.CallMethod[Event](a_, "currentEvent")
	return rv
}

func (a_ Application) IsRunning() bool {
	rv := ffi.CallMethod[bool](a_, "isRunning")
	return rv
}

func (a_ Application) IsActive() bool {
	rv := ffi.CallMethod[bool](a_, "isActive")
	return rv
}

func (a_ Application) EnabledRemoteNotificationTypes() RemoteNotificationType {
	rv := ffi.CallMethod[RemoteNotificationType](a_, "enabledRemoteNotificationTypes")
	return rv
}

func (a_ Application) IsRegisteredForRemoteNotifications() bool {
	rv := ffi.CallMethod[bool](a_, "isRegisteredForRemoteNotifications")
	return rv
}

func (a_ Application) Appearance() Appearance {
	rv := ffi.CallMethod[Appearance](a_, "appearance")
	return rv
}

func (a_ Application) SetAppearance(value IAppearance) {
	ffi.CallMethod[ffi.Void](a_, "setAppearance:", value)
}

func (a_ Application) EffectiveAppearance() Appearance {
	rv := ffi.CallMethod[Appearance](a_, "effectiveAppearance")
	return rv
}

func (a_ Application) CurrentSystemPresentationOptions() ApplicationPresentationOptions {
	rv := ffi.CallMethod[ApplicationPresentationOptions](a_, "currentSystemPresentationOptions")
	return rv
}

func (a_ Application) PresentationOptions() ApplicationPresentationOptions {
	rv := ffi.CallMethod[ApplicationPresentationOptions](a_, "presentationOptions")
	return rv
}

func (a_ Application) SetPresentationOptions(value ApplicationPresentationOptions) {
	ffi.CallMethod[ffi.Void](a_, "setPresentationOptions:", value)
}

func (a_ Application) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := ffi.CallMethod[UserInterfaceLayoutDirection](a_, "userInterfaceLayoutDirection")
	return rv
}

func (a_ Application) DockTile() DockTile {
	rv := ffi.CallMethod[DockTile](a_, "dockTile")
	return rv
}

func (a_ Application) ApplicationIconImage() Image {
	rv := ffi.CallMethod[Image](a_, "applicationIconImage")
	return rv
}

func (a_ Application) SetApplicationIconImage(value IImage) {
	ffi.CallMethod[ffi.Void](a_, "setApplicationIconImage:", value)
}

func (a_ Application) HelpMenu() Menu {
	rv := ffi.CallMethod[Menu](a_, "helpMenu")
	return rv
}

func (a_ Application) SetHelpMenu(value IMenu) {
	ffi.CallMethod[ffi.Void](a_, "setHelpMenu:", value)
}

func (a_ Application) ServicesProvider() objc.Object {
	rv := ffi.CallMethod[objc.Object](a_, "servicesProvider")
	return rv
}

func (a_ Application) SetServicesProvider(value objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "setServicesProvider:", value)
}

func (a_ Application) IsFullKeyboardAccessEnabled() bool {
	rv := ffi.CallMethod[bool](a_, "isFullKeyboardAccessEnabled")
	return rv
}

func (a_ Application) OrderedDocuments() []Document {
	rv := ffi.CallMethod[[]Document](a_, "orderedDocuments")
	return rv
}

func (a_ Application) OrderedWindows() []Window {
	rv := ffi.CallMethod[[]Window](a_, "orderedWindows")
	return rv
}

func (a_ Application) KeyWindow() Window {
	rv := ffi.CallMethod[Window](a_, "keyWindow")
	return rv
}

func (a_ Application) MainWindow() Window {
	rv := ffi.CallMethod[Window](a_, "mainWindow")
	return rv
}

func (a_ Application) Windows() []Window {
	rv := ffi.CallMethod[[]Window](a_, "windows")
	return rv
}

func (a_ Application) IsHidden() bool {
	rv := ffi.CallMethod[bool](a_, "isHidden")
	return rv
}

// deprecated
func (a_ Application) Context() GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](a_, "context")
	return rv
}

func (a_ Application) OcclusionState() ApplicationOcclusionState {
	rv := ffi.CallMethod[ApplicationOcclusionState](a_, "occlusionState")
	return rv
}

func (a_ Application) IsProtectedDataAvailable() bool {
	rv := ffi.CallMethod[bool](a_, "isProtectedDataAvailable")
	return rv
}

func (a_ Application) ModalWindow() Window {
	rv := ffi.CallMethod[Window](a_, "modalWindow")
	return rv
}

func (a_ Application) MainMenu() Menu {
	rv := ffi.CallMethod[Menu](a_, "mainMenu")
	return rv
}

func (a_ Application) SetMainMenu(value IMenu) {
	ffi.CallMethod[ffi.Void](a_, "setMainMenu:", value)
}

func (a_ Application) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := ffi.CallMethod[bool](a_, "isAutomaticCustomizeTouchBarMenuItemEnabled")
	return rv
}

func (a_ Application) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	ffi.CallMethod[ffi.Void](a_, "setAutomaticCustomizeTouchBarMenuItemEnabled:", value)
}

func (a_ Application) WindowsMenu() Menu {
	rv := ffi.CallMethod[Menu](a_, "windowsMenu")
	return rv
}

func (a_ Application) SetWindowsMenu(value IMenu) {
	ffi.CallMethod[ffi.Void](a_, "setWindowsMenu:", value)
}

func (a_ Application) ServicesMenu() Menu {
	rv := ffi.CallMethod[Menu](a_, "servicesMenu")
	return rv
}

func (a_ Application) SetServicesMenu(value IMenu) {
	ffi.CallMethod[ffi.Void](a_, "setServicesMenu:", value)
}

type ApplicationDelegate interface {
	ImplementsApplicationWillFinishLaunching() bool
	// optional
	ApplicationWillFinishLaunching(notification foundation.Notification)
	ImplementsApplicationDidFinishLaunching() bool
	// optional
	ApplicationDidFinishLaunching(notification foundation.Notification)
	ImplementsApplicationWillBecomeActive() bool
	// optional
	ApplicationWillBecomeActive(notification foundation.Notification)
	ImplementsApplicationDidBecomeActive() bool
	// optional
	ApplicationDidBecomeActive(notification foundation.Notification)
	ImplementsApplicationWillResignActive() bool
	// optional
	ApplicationWillResignActive(notification foundation.Notification)
	ImplementsApplicationDidResignActive() bool
	// optional
	ApplicationDidResignActive(notification foundation.Notification)
	ImplementsApplicationShouldTerminate() bool
	// optional
	ApplicationShouldTerminate(sender Application) ApplicationTerminateReply
	ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool
	// optional
	ApplicationShouldTerminateAfterLastWindowClosed(sender Application) bool
	ImplementsApplicationWillTerminate() bool
	// optional
	ApplicationWillTerminate(notification foundation.Notification)
	ImplementsApplicationWillHide() bool
	// optional
	ApplicationWillHide(notification foundation.Notification)
	ImplementsApplicationDidHide() bool
	// optional
	ApplicationDidHide(notification foundation.Notification)
	ImplementsApplicationWillUnhide() bool
	// optional
	ApplicationWillUnhide(notification foundation.Notification)
	ImplementsApplicationDidUnhide() bool
	// optional
	ApplicationDidUnhide(notification foundation.Notification)
	ImplementsApplicationWillUpdate() bool
	// optional
	ApplicationWillUpdate(notification foundation.Notification)
	ImplementsApplicationDidUpdate() bool
	// optional
	ApplicationDidUpdate(notification foundation.Notification)
	ImplementsApplicationShouldHandleReopen_HasVisibleWindows() bool
	// optional
	ApplicationShouldHandleReopen_HasVisibleWindows(sender Application, flag bool) bool
	ImplementsApplicationDockMenu() bool
	// optional
	ApplicationDockMenu(sender Application) IMenu
	ImplementsApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool
	// optional
	ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application Application) bool
	ImplementsApplication_WillPresentError() bool
	// optional
	Application_WillPresentError(application Application, error foundation.Error) foundation.IError
	ImplementsApplicationDidChangeScreenParameters() bool
	// optional
	ApplicationDidChangeScreenParameters(notification foundation.Notification)
	ImplementsApplication_WillContinueUserActivityWithType() bool
	// optional
	Application_WillContinueUserActivityWithType(application Application, userActivityType string) bool
	ImplementsApplication_DidFailToContinueUserActivityWithType_Error() bool
	// optional
	Application_DidFailToContinueUserActivityWithType_Error(application Application, userActivityType string, error foundation.Error)
	ImplementsApplication_DidUpdateUserActivity() bool
	// optional
	Application_DidUpdateUserActivity(application Application, userActivity foundation.UserActivity)
	ImplementsApplication_DidRegisterForRemoteNotificationsWithDeviceToken() bool
	// optional
	Application_DidRegisterForRemoteNotificationsWithDeviceToken(application Application, deviceToken []byte)
	ImplementsApplication_DidFailToRegisterForRemoteNotificationsWithError() bool
	// optional
	Application_DidFailToRegisterForRemoteNotificationsWithError(application Application, error foundation.Error)
	ImplementsApplication_DidReceiveRemoteNotification() bool
	// optional
	Application_DidReceiveRemoteNotification(application Application, userInfo map[string]objc.Object)
	ImplementsApplication_OpenURLs() bool
	// optional
	Application_OpenURLs(application Application, urls []foundation.URL)
	ImplementsApplication_OpenFile() bool
	// optional
	Application_OpenFile(sender Application, filename string) bool
	ImplementsApplication_OpenFileWithoutUI() bool
	// optional
	Application_OpenFileWithoutUI(sender objc.Object, filename string) bool
	ImplementsApplication_OpenTempFile() bool
	// optional
	Application_OpenTempFile(sender Application, filename string) bool
	ImplementsApplication_OpenFiles() bool
	// optional
	Application_OpenFiles(sender Application, filenames []string)
	ImplementsApplicationShouldOpenUntitledFile() bool
	// optional
	ApplicationShouldOpenUntitledFile(sender Application) bool
	ImplementsApplicationOpenUntitledFile() bool
	// optional
	ApplicationOpenUntitledFile(sender Application) bool
	ImplementsApplication_PrintFile() bool
	// optional
	Application_PrintFile(sender Application, filename string) bool
	ImplementsApplication_PrintFiles_WithSettings_ShowPrintPanels() bool
	// optional
	Application_PrintFiles_WithSettings_ShowPrintPanels(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply
	ImplementsApplicationSupportsSecureRestorableState() bool
	// optional
	ApplicationSupportsSecureRestorableState(app Application) bool
	ImplementsApplicationProtectedDataDidBecomeAvailable() bool
	// optional
	ApplicationProtectedDataDidBecomeAvailable(notification foundation.Notification)
	ImplementsApplicationProtectedDataWillBecomeUnavailable() bool
	// optional
	ApplicationProtectedDataWillBecomeUnavailable(notification foundation.Notification)
	ImplementsApplicationDidChangeOcclusionState() bool
	// optional
	ApplicationDidChangeOcclusionState(notification foundation.Notification)
	ImplementsApplication_DelegateHandlesKey() bool
	// optional
	Application_DelegateHandlesKey(sender Application, key string) bool
}

type ApplicationDelegateImpl struct {
	_ApplicationWillFinishLaunching                               func(notification foundation.Notification)
	_ApplicationDidFinishLaunching                                func(notification foundation.Notification)
	_ApplicationWillBecomeActive                                  func(notification foundation.Notification)
	_ApplicationDidBecomeActive                                   func(notification foundation.Notification)
	_ApplicationWillResignActive                                  func(notification foundation.Notification)
	_ApplicationDidResignActive                                   func(notification foundation.Notification)
	_ApplicationShouldTerminate                                   func(sender Application) ApplicationTerminateReply
	_ApplicationShouldTerminateAfterLastWindowClosed              func(sender Application) bool
	_ApplicationWillTerminate                                     func(notification foundation.Notification)
	_ApplicationWillHide                                          func(notification foundation.Notification)
	_ApplicationDidHide                                           func(notification foundation.Notification)
	_ApplicationWillUnhide                                        func(notification foundation.Notification)
	_ApplicationDidUnhide                                         func(notification foundation.Notification)
	_ApplicationWillUpdate                                        func(notification foundation.Notification)
	_ApplicationDidUpdate                                         func(notification foundation.Notification)
	_ApplicationShouldHandleReopen_HasVisibleWindows              func(sender Application, flag bool) bool
	_ApplicationDockMenu                                          func(sender Application) IMenu
	_ApplicationShouldAutomaticallyLocalizeKeyEquivalents         func(application Application) bool
	_Application_WillPresentError                                 func(application Application, error foundation.Error) foundation.IError
	_ApplicationDidChangeScreenParameters                         func(notification foundation.Notification)
	_Application_WillContinueUserActivityWithType                 func(application Application, userActivityType string) bool
	_Application_DidFailToContinueUserActivityWithType_Error      func(application Application, userActivityType string, error foundation.Error)
	_Application_DidUpdateUserActivity                            func(application Application, userActivity foundation.UserActivity)
	_Application_DidRegisterForRemoteNotificationsWithDeviceToken func(application Application, deviceToken []byte)
	_Application_DidFailToRegisterForRemoteNotificationsWithError func(application Application, error foundation.Error)
	_Application_DidReceiveRemoteNotification                     func(application Application, userInfo map[string]objc.Object)
	_Application_OpenURLs                                         func(application Application, urls []foundation.URL)
	_Application_OpenFile                                         func(sender Application, filename string) bool
	_Application_OpenFileWithoutUI                                func(sender objc.Object, filename string) bool
	_Application_OpenTempFile                                     func(sender Application, filename string) bool
	_Application_OpenFiles                                        func(sender Application, filenames []string)
	_ApplicationShouldOpenUntitledFile                            func(sender Application) bool
	_ApplicationOpenUntitledFile                                  func(sender Application) bool
	_Application_PrintFile                                        func(sender Application, filename string) bool
	_Application_PrintFiles_WithSettings_ShowPrintPanels          func(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply
	_ApplicationSupportsSecureRestorableState                     func(app Application) bool
	_ApplicationProtectedDataDidBecomeAvailable                   func(notification foundation.Notification)
	_ApplicationProtectedDataWillBecomeUnavailable                func(notification foundation.Notification)
	_ApplicationDidChangeOcclusionState                           func(notification foundation.Notification)
	_Application_DelegateHandlesKey                               func(sender Application, key string) bool
}

func (di *ApplicationDelegateImpl) ImplementsApplicationWillFinishLaunching() bool {
	return di._ApplicationWillFinishLaunching != nil
}

func (di *ApplicationDelegateImpl) SetApplicationWillFinishLaunching(f func(notification foundation.Notification)) {
	di._ApplicationWillFinishLaunching = f
}

func (di *ApplicationDelegateImpl) ApplicationWillFinishLaunching(notification foundation.Notification) {
	di._ApplicationWillFinishLaunching(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationDidFinishLaunching() bool {
	return di._ApplicationDidFinishLaunching != nil
}

func (di *ApplicationDelegateImpl) SetApplicationDidFinishLaunching(f func(notification foundation.Notification)) {
	di._ApplicationDidFinishLaunching = f
}

func (di *ApplicationDelegateImpl) ApplicationDidFinishLaunching(notification foundation.Notification) {
	di._ApplicationDidFinishLaunching(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationWillBecomeActive() bool {
	return di._ApplicationWillBecomeActive != nil
}

func (di *ApplicationDelegateImpl) SetApplicationWillBecomeActive(f func(notification foundation.Notification)) {
	di._ApplicationWillBecomeActive = f
}

func (di *ApplicationDelegateImpl) ApplicationWillBecomeActive(notification foundation.Notification) {
	di._ApplicationWillBecomeActive(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationDidBecomeActive() bool {
	return di._ApplicationDidBecomeActive != nil
}

func (di *ApplicationDelegateImpl) SetApplicationDidBecomeActive(f func(notification foundation.Notification)) {
	di._ApplicationDidBecomeActive = f
}

func (di *ApplicationDelegateImpl) ApplicationDidBecomeActive(notification foundation.Notification) {
	di._ApplicationDidBecomeActive(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationWillResignActive() bool {
	return di._ApplicationWillResignActive != nil
}

func (di *ApplicationDelegateImpl) SetApplicationWillResignActive(f func(notification foundation.Notification)) {
	di._ApplicationWillResignActive = f
}

func (di *ApplicationDelegateImpl) ApplicationWillResignActive(notification foundation.Notification) {
	di._ApplicationWillResignActive(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationDidResignActive() bool {
	return di._ApplicationDidResignActive != nil
}

func (di *ApplicationDelegateImpl) SetApplicationDidResignActive(f func(notification foundation.Notification)) {
	di._ApplicationDidResignActive = f
}

func (di *ApplicationDelegateImpl) ApplicationDidResignActive(notification foundation.Notification) {
	di._ApplicationDidResignActive(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationShouldTerminate() bool {
	return di._ApplicationShouldTerminate != nil
}

func (di *ApplicationDelegateImpl) SetApplicationShouldTerminate(f func(sender Application) ApplicationTerminateReply) {
	di._ApplicationShouldTerminate = f
}

func (di *ApplicationDelegateImpl) ApplicationShouldTerminate(sender Application) ApplicationTerminateReply {
	return di._ApplicationShouldTerminate(sender)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool {
	return di._ApplicationShouldTerminateAfterLastWindowClosed != nil
}

func (di *ApplicationDelegateImpl) SetApplicationShouldTerminateAfterLastWindowClosed(f func(sender Application) bool) {
	di._ApplicationShouldTerminateAfterLastWindowClosed = f
}

func (di *ApplicationDelegateImpl) ApplicationShouldTerminateAfterLastWindowClosed(sender Application) bool {
	return di._ApplicationShouldTerminateAfterLastWindowClosed(sender)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationWillTerminate() bool {
	return di._ApplicationWillTerminate != nil
}

func (di *ApplicationDelegateImpl) SetApplicationWillTerminate(f func(notification foundation.Notification)) {
	di._ApplicationWillTerminate = f
}

func (di *ApplicationDelegateImpl) ApplicationWillTerminate(notification foundation.Notification) {
	di._ApplicationWillTerminate(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationWillHide() bool {
	return di._ApplicationWillHide != nil
}

func (di *ApplicationDelegateImpl) SetApplicationWillHide(f func(notification foundation.Notification)) {
	di._ApplicationWillHide = f
}

func (di *ApplicationDelegateImpl) ApplicationWillHide(notification foundation.Notification) {
	di._ApplicationWillHide(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationDidHide() bool {
	return di._ApplicationDidHide != nil
}

func (di *ApplicationDelegateImpl) SetApplicationDidHide(f func(notification foundation.Notification)) {
	di._ApplicationDidHide = f
}

func (di *ApplicationDelegateImpl) ApplicationDidHide(notification foundation.Notification) {
	di._ApplicationDidHide(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationWillUnhide() bool {
	return di._ApplicationWillUnhide != nil
}

func (di *ApplicationDelegateImpl) SetApplicationWillUnhide(f func(notification foundation.Notification)) {
	di._ApplicationWillUnhide = f
}

func (di *ApplicationDelegateImpl) ApplicationWillUnhide(notification foundation.Notification) {
	di._ApplicationWillUnhide(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationDidUnhide() bool {
	return di._ApplicationDidUnhide != nil
}

func (di *ApplicationDelegateImpl) SetApplicationDidUnhide(f func(notification foundation.Notification)) {
	di._ApplicationDidUnhide = f
}

func (di *ApplicationDelegateImpl) ApplicationDidUnhide(notification foundation.Notification) {
	di._ApplicationDidUnhide(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationWillUpdate() bool {
	return di._ApplicationWillUpdate != nil
}

func (di *ApplicationDelegateImpl) SetApplicationWillUpdate(f func(notification foundation.Notification)) {
	di._ApplicationWillUpdate = f
}

func (di *ApplicationDelegateImpl) ApplicationWillUpdate(notification foundation.Notification) {
	di._ApplicationWillUpdate(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationDidUpdate() bool {
	return di._ApplicationDidUpdate != nil
}

func (di *ApplicationDelegateImpl) SetApplicationDidUpdate(f func(notification foundation.Notification)) {
	di._ApplicationDidUpdate = f
}

func (di *ApplicationDelegateImpl) ApplicationDidUpdate(notification foundation.Notification) {
	di._ApplicationDidUpdate(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationShouldHandleReopen_HasVisibleWindows() bool {
	return di._ApplicationShouldHandleReopen_HasVisibleWindows != nil
}

func (di *ApplicationDelegateImpl) SetApplicationShouldHandleReopen_HasVisibleWindows(f func(sender Application, flag bool) bool) {
	di._ApplicationShouldHandleReopen_HasVisibleWindows = f
}

func (di *ApplicationDelegateImpl) ApplicationShouldHandleReopen_HasVisibleWindows(sender Application, flag bool) bool {
	return di._ApplicationShouldHandleReopen_HasVisibleWindows(sender, flag)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationDockMenu() bool {
	return di._ApplicationDockMenu != nil
}

func (di *ApplicationDelegateImpl) SetApplicationDockMenu(f func(sender Application) IMenu) {
	di._ApplicationDockMenu = f
}

func (di *ApplicationDelegateImpl) ApplicationDockMenu(sender Application) IMenu {
	return di._ApplicationDockMenu(sender)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool {
	return di._ApplicationShouldAutomaticallyLocalizeKeyEquivalents != nil
}

func (di *ApplicationDelegateImpl) SetApplicationShouldAutomaticallyLocalizeKeyEquivalents(f func(application Application) bool) {
	di._ApplicationShouldAutomaticallyLocalizeKeyEquivalents = f
}

func (di *ApplicationDelegateImpl) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application Application) bool {
	return di._ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_WillPresentError() bool {
	return di._Application_WillPresentError != nil
}

func (di *ApplicationDelegateImpl) SetApplication_WillPresentError(f func(application Application, error foundation.Error) foundation.IError) {
	di._Application_WillPresentError = f
}

func (di *ApplicationDelegateImpl) Application_WillPresentError(application Application, error foundation.Error) foundation.IError {
	return di._Application_WillPresentError(application, error)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationDidChangeScreenParameters() bool {
	return di._ApplicationDidChangeScreenParameters != nil
}

func (di *ApplicationDelegateImpl) SetApplicationDidChangeScreenParameters(f func(notification foundation.Notification)) {
	di._ApplicationDidChangeScreenParameters = f
}

func (di *ApplicationDelegateImpl) ApplicationDidChangeScreenParameters(notification foundation.Notification) {
	di._ApplicationDidChangeScreenParameters(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_WillContinueUserActivityWithType() bool {
	return di._Application_WillContinueUserActivityWithType != nil
}

func (di *ApplicationDelegateImpl) SetApplication_WillContinueUserActivityWithType(f func(application Application, userActivityType string) bool) {
	di._Application_WillContinueUserActivityWithType = f
}

func (di *ApplicationDelegateImpl) Application_WillContinueUserActivityWithType(application Application, userActivityType string) bool {
	return di._Application_WillContinueUserActivityWithType(application, userActivityType)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_DidFailToContinueUserActivityWithType_Error() bool {
	return di._Application_DidFailToContinueUserActivityWithType_Error != nil
}

func (di *ApplicationDelegateImpl) SetApplication_DidFailToContinueUserActivityWithType_Error(f func(application Application, userActivityType string, error foundation.Error)) {
	di._Application_DidFailToContinueUserActivityWithType_Error = f
}

func (di *ApplicationDelegateImpl) Application_DidFailToContinueUserActivityWithType_Error(application Application, userActivityType string, error foundation.Error) {
	di._Application_DidFailToContinueUserActivityWithType_Error(application, userActivityType, error)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_DidUpdateUserActivity() bool {
	return di._Application_DidUpdateUserActivity != nil
}

func (di *ApplicationDelegateImpl) SetApplication_DidUpdateUserActivity(f func(application Application, userActivity foundation.UserActivity)) {
	di._Application_DidUpdateUserActivity = f
}

func (di *ApplicationDelegateImpl) Application_DidUpdateUserActivity(application Application, userActivity foundation.UserActivity) {
	di._Application_DidUpdateUserActivity(application, userActivity)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_DidRegisterForRemoteNotificationsWithDeviceToken() bool {
	return di._Application_DidRegisterForRemoteNotificationsWithDeviceToken != nil
}

func (di *ApplicationDelegateImpl) SetApplication_DidRegisterForRemoteNotificationsWithDeviceToken(f func(application Application, deviceToken []byte)) {
	di._Application_DidRegisterForRemoteNotificationsWithDeviceToken = f
}

func (di *ApplicationDelegateImpl) Application_DidRegisterForRemoteNotificationsWithDeviceToken(application Application, deviceToken []byte) {
	di._Application_DidRegisterForRemoteNotificationsWithDeviceToken(application, deviceToken)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_DidFailToRegisterForRemoteNotificationsWithError() bool {
	return di._Application_DidFailToRegisterForRemoteNotificationsWithError != nil
}

func (di *ApplicationDelegateImpl) SetApplication_DidFailToRegisterForRemoteNotificationsWithError(f func(application Application, error foundation.Error)) {
	di._Application_DidFailToRegisterForRemoteNotificationsWithError = f
}

func (di *ApplicationDelegateImpl) Application_DidFailToRegisterForRemoteNotificationsWithError(application Application, error foundation.Error) {
	di._Application_DidFailToRegisterForRemoteNotificationsWithError(application, error)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_DidReceiveRemoteNotification() bool {
	return di._Application_DidReceiveRemoteNotification != nil
}

func (di *ApplicationDelegateImpl) SetApplication_DidReceiveRemoteNotification(f func(application Application, userInfo map[string]objc.Object)) {
	di._Application_DidReceiveRemoteNotification = f
}

func (di *ApplicationDelegateImpl) Application_DidReceiveRemoteNotification(application Application, userInfo map[string]objc.Object) {
	di._Application_DidReceiveRemoteNotification(application, userInfo)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_OpenURLs() bool {
	return di._Application_OpenURLs != nil
}

func (di *ApplicationDelegateImpl) SetApplication_OpenURLs(f func(application Application, urls []foundation.URL)) {
	di._Application_OpenURLs = f
}

func (di *ApplicationDelegateImpl) Application_OpenURLs(application Application, urls []foundation.URL) {
	di._Application_OpenURLs(application, urls)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_OpenFile() bool {
	return di._Application_OpenFile != nil
}

func (di *ApplicationDelegateImpl) SetApplication_OpenFile(f func(sender Application, filename string) bool) {
	di._Application_OpenFile = f
}

func (di *ApplicationDelegateImpl) Application_OpenFile(sender Application, filename string) bool {
	return di._Application_OpenFile(sender, filename)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_OpenFileWithoutUI() bool {
	return di._Application_OpenFileWithoutUI != nil
}

func (di *ApplicationDelegateImpl) SetApplication_OpenFileWithoutUI(f func(sender objc.Object, filename string) bool) {
	di._Application_OpenFileWithoutUI = f
}

func (di *ApplicationDelegateImpl) Application_OpenFileWithoutUI(sender objc.Object, filename string) bool {
	return di._Application_OpenFileWithoutUI(sender, filename)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_OpenTempFile() bool {
	return di._Application_OpenTempFile != nil
}

func (di *ApplicationDelegateImpl) SetApplication_OpenTempFile(f func(sender Application, filename string) bool) {
	di._Application_OpenTempFile = f
}

func (di *ApplicationDelegateImpl) Application_OpenTempFile(sender Application, filename string) bool {
	return di._Application_OpenTempFile(sender, filename)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_OpenFiles() bool {
	return di._Application_OpenFiles != nil
}

func (di *ApplicationDelegateImpl) SetApplication_OpenFiles(f func(sender Application, filenames []string)) {
	di._Application_OpenFiles = f
}

func (di *ApplicationDelegateImpl) Application_OpenFiles(sender Application, filenames []string) {
	di._Application_OpenFiles(sender, filenames)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationShouldOpenUntitledFile() bool {
	return di._ApplicationShouldOpenUntitledFile != nil
}

func (di *ApplicationDelegateImpl) SetApplicationShouldOpenUntitledFile(f func(sender Application) bool) {
	di._ApplicationShouldOpenUntitledFile = f
}

func (di *ApplicationDelegateImpl) ApplicationShouldOpenUntitledFile(sender Application) bool {
	return di._ApplicationShouldOpenUntitledFile(sender)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationOpenUntitledFile() bool {
	return di._ApplicationOpenUntitledFile != nil
}

func (di *ApplicationDelegateImpl) SetApplicationOpenUntitledFile(f func(sender Application) bool) {
	di._ApplicationOpenUntitledFile = f
}

func (di *ApplicationDelegateImpl) ApplicationOpenUntitledFile(sender Application) bool {
	return di._ApplicationOpenUntitledFile(sender)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_PrintFile() bool {
	return di._Application_PrintFile != nil
}

func (di *ApplicationDelegateImpl) SetApplication_PrintFile(f func(sender Application, filename string) bool) {
	di._Application_PrintFile = f
}

func (di *ApplicationDelegateImpl) Application_PrintFile(sender Application, filename string) bool {
	return di._Application_PrintFile(sender, filename)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_PrintFiles_WithSettings_ShowPrintPanels() bool {
	return di._Application_PrintFiles_WithSettings_ShowPrintPanels != nil
}

func (di *ApplicationDelegateImpl) SetApplication_PrintFiles_WithSettings_ShowPrintPanels(f func(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply) {
	di._Application_PrintFiles_WithSettings_ShowPrintPanels = f
}

func (di *ApplicationDelegateImpl) Application_PrintFiles_WithSettings_ShowPrintPanels(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply {
	return di._Application_PrintFiles_WithSettings_ShowPrintPanels(application, fileNames, printSettings, showPrintPanels)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationSupportsSecureRestorableState() bool {
	return di._ApplicationSupportsSecureRestorableState != nil
}

func (di *ApplicationDelegateImpl) SetApplicationSupportsSecureRestorableState(f func(app Application) bool) {
	di._ApplicationSupportsSecureRestorableState = f
}

func (di *ApplicationDelegateImpl) ApplicationSupportsSecureRestorableState(app Application) bool {
	return di._ApplicationSupportsSecureRestorableState(app)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationProtectedDataDidBecomeAvailable() bool {
	return di._ApplicationProtectedDataDidBecomeAvailable != nil
}

func (di *ApplicationDelegateImpl) SetApplicationProtectedDataDidBecomeAvailable(f func(notification foundation.Notification)) {
	di._ApplicationProtectedDataDidBecomeAvailable = f
}

func (di *ApplicationDelegateImpl) ApplicationProtectedDataDidBecomeAvailable(notification foundation.Notification) {
	di._ApplicationProtectedDataDidBecomeAvailable(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationProtectedDataWillBecomeUnavailable() bool {
	return di._ApplicationProtectedDataWillBecomeUnavailable != nil
}

func (di *ApplicationDelegateImpl) SetApplicationProtectedDataWillBecomeUnavailable(f func(notification foundation.Notification)) {
	di._ApplicationProtectedDataWillBecomeUnavailable = f
}

func (di *ApplicationDelegateImpl) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.Notification) {
	di._ApplicationProtectedDataWillBecomeUnavailable(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplicationDidChangeOcclusionState() bool {
	return di._ApplicationDidChangeOcclusionState != nil
}

func (di *ApplicationDelegateImpl) SetApplicationDidChangeOcclusionState(f func(notification foundation.Notification)) {
	di._ApplicationDidChangeOcclusionState = f
}

func (di *ApplicationDelegateImpl) ApplicationDidChangeOcclusionState(notification foundation.Notification) {
	di._ApplicationDidChangeOcclusionState(notification)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_DelegateHandlesKey() bool {
	return di._Application_DelegateHandlesKey != nil
}

func (di *ApplicationDelegateImpl) SetApplication_DelegateHandlesKey(f func(sender Application, key string) bool) {
	di._Application_DelegateHandlesKey = f
}

func (di *ApplicationDelegateImpl) Application_DelegateHandlesKey(sender Application, key string) bool {
	return di._Application_DelegateHandlesKey(sender, key)
}

type ApplicationDelegateWrapper struct {
	objc.Object
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillFinishLaunching() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillFinishLaunching:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillFinishLaunching(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationWillFinishLaunching:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidFinishLaunching() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidFinishLaunching:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidFinishLaunching(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationDidFinishLaunching:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillBecomeActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillBecomeActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillBecomeActive(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationWillBecomeActive:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidBecomeActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidBecomeActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidBecomeActive(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationDidBecomeActive:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillResignActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillResignActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillResignActive(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationWillResignActive:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidResignActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidResignActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidResignActive(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationDidResignActive:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldTerminate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldTerminate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminate(sender IApplication) ApplicationTerminateReply {
	rv := ffi.CallMethod[ApplicationTerminateReply](a_, "applicationShouldTerminate:", sender)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldTerminateAfterLastWindowClosed:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminateAfterLastWindowClosed(sender IApplication) bool {
	rv := ffi.CallMethod[bool](a_, "applicationShouldTerminateAfterLastWindowClosed:", sender)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillTerminate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillTerminate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillTerminate(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationWillTerminate:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillHide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillHide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillHide(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationWillHide:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidHide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidHide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidHide(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationDidHide:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillUnhide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillUnhide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillUnhide(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationWillUnhide:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidUnhide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidUnhide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidUnhide(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationDidUnhide:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillUpdate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillUpdate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillUpdate(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationWillUpdate:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidUpdate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidUpdate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidUpdate(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationDidUpdate:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldHandleReopen_HasVisibleWindows() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldHandleReopen:hasVisibleWindows:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldHandleReopen_HasVisibleWindows(sender IApplication, flag bool) bool {
	rv := ffi.CallMethod[bool](a_, "applicationShouldHandleReopen:hasVisibleWindows:", sender, flag)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDockMenu() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDockMenu:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDockMenu(sender IApplication) Menu {
	rv := ffi.CallMethod[Menu](a_, "applicationDockMenu:", sender)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldAutomaticallyLocalizeKeyEquivalents:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application IApplication) bool {
	rv := ffi.CallMethod[bool](a_, "applicationShouldAutomaticallyLocalizeKeyEquivalents:", application)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_WillPresentError() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:willPresentError:"))
}

func (a_ ApplicationDelegateWrapper) Application_WillPresentError(application IApplication, error foundation.IError) foundation.Error {
	rv := ffi.CallMethod[foundation.Error](a_, "application:willPresentError:", application, error)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidChangeScreenParameters() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidChangeScreenParameters:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidChangeScreenParameters(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationDidChangeScreenParameters:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_WillContinueUserActivityWithType() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:willContinueUserActivityWithType:"))
}

func (a_ ApplicationDelegateWrapper) Application_WillContinueUserActivityWithType(application IApplication, userActivityType string) bool {
	rv := ffi.CallMethod[bool](a_, "application:willContinueUserActivityWithType:", application, userActivityType)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidFailToContinueUserActivityWithType_Error() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didFailToContinueUserActivityWithType:error:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidFailToContinueUserActivityWithType_Error(application IApplication, userActivityType string, error foundation.IError) {
	ffi.CallMethod[ffi.Void](a_, "application:didFailToContinueUserActivityWithType:error:", application, userActivityType, error)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidUpdateUserActivity() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didUpdateUserActivity:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidUpdateUserActivity(application IApplication, userActivity foundation.IUserActivity) {
	ffi.CallMethod[ffi.Void](a_, "application:didUpdateUserActivity:", application, userActivity)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidRegisterForRemoteNotificationsWithDeviceToken() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didRegisterForRemoteNotificationsWithDeviceToken:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidRegisterForRemoteNotificationsWithDeviceToken(application IApplication, deviceToken []byte) {
	ffi.CallMethod[ffi.Void](a_, "application:didRegisterForRemoteNotificationsWithDeviceToken:", application, deviceToken)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidFailToRegisterForRemoteNotificationsWithError() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didFailToRegisterForRemoteNotificationsWithError:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidFailToRegisterForRemoteNotificationsWithError(application IApplication, error foundation.IError) {
	ffi.CallMethod[ffi.Void](a_, "application:didFailToRegisterForRemoteNotificationsWithError:", application, error)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidReceiveRemoteNotification() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didReceiveRemoteNotification:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidReceiveRemoteNotification(application IApplication, userInfo map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "application:didReceiveRemoteNotification:", application, userInfo)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenURLs() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openURLs:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenURLs(application IApplication, urls []foundation.IURL) {
	ffi.CallMethod[ffi.Void](a_, "application:openURLs:", application, urls)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openFile:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenFile(sender IApplication, filename string) bool {
	rv := ffi.CallMethod[bool](a_, "application:openFile:", sender, filename)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenFileWithoutUI() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openFileWithoutUI:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenFileWithoutUI(sender objc.IObject, filename string) bool {
	rv := ffi.CallMethod[bool](a_, "application:openFileWithoutUI:", sender, filename)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenTempFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openTempFile:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenTempFile(sender IApplication, filename string) bool {
	rv := ffi.CallMethod[bool](a_, "application:openTempFile:", sender, filename)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenFiles() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openFiles:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenFiles(sender IApplication, filenames []string) {
	ffi.CallMethod[ffi.Void](a_, "application:openFiles:", sender, filenames)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldOpenUntitledFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldOpenUntitledFile:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldOpenUntitledFile(sender IApplication) bool {
	rv := ffi.CallMethod[bool](a_, "applicationShouldOpenUntitledFile:", sender)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationOpenUntitledFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationOpenUntitledFile:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationOpenUntitledFile(sender IApplication) bool {
	rv := ffi.CallMethod[bool](a_, "applicationOpenUntitledFile:", sender)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_PrintFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:printFile:"))
}

func (a_ ApplicationDelegateWrapper) Application_PrintFile(sender IApplication, filename string) bool {
	rv := ffi.CallMethod[bool](a_, "application:printFile:", sender, filename)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_PrintFiles_WithSettings_ShowPrintPanels() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:printFiles:withSettings:showPrintPanels:"))
}

func (a_ ApplicationDelegateWrapper) Application_PrintFiles_WithSettings_ShowPrintPanels(application IApplication, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanels bool) ApplicationPrintReply {
	rv := ffi.CallMethod[ApplicationPrintReply](a_, "application:printFiles:withSettings:showPrintPanels:", application, fileNames, printSettings, showPrintPanels)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationSupportsSecureRestorableState() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationSupportsSecureRestorableState:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationSupportsSecureRestorableState(app IApplication) bool {
	rv := ffi.CallMethod[bool](a_, "applicationSupportsSecureRestorableState:", app)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationProtectedDataDidBecomeAvailable() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationProtectedDataDidBecomeAvailable:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataDidBecomeAvailable(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationProtectedDataDidBecomeAvailable:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationProtectedDataWillBecomeUnavailable() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationProtectedDataWillBecomeUnavailable:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationProtectedDataWillBecomeUnavailable:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidChangeOcclusionState() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidChangeOcclusionState:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidChangeOcclusionState(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](a_, "applicationDidChangeOcclusionState:", notification)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DelegateHandlesKey() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:delegateHandlesKey:"))
}

func (a_ ApplicationDelegateWrapper) Application_DelegateHandlesKey(sender IApplication, key string) bool {
	rv := ffi.CallMethod[bool](a_, "application:delegateHandlesKey:", sender, key)
	return rv
}
