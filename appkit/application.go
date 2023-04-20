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
	RegisterUserInterfaceItemSearchHandler(handler UserInterfaceItemSearching)
	RegisterUserInterfaceItemSearchHandler0(handler objc.IObject)
	SearchString_InUserInterfaceItemString_SearchRange_FoundRange(searchString string, stringToSearch string, searchRange foundation.Range, foundRange *foundation.Range) bool
	UnregisterUserInterfaceItemSearchHandler(handler UserInterfaceItemSearching)
	UnregisterUserInterfaceItemSearchHandler0(handler objc.IObject)
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
	RestoreWindowWithIdentifier_State_CompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(param1 Window, param2 foundation.Error)) bool
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
	AddWindowsItem_Title_Filename(win IWindow, string_ string, isFilename bool)
	ChangeWindowsItem_Title_Filename(win IWindow, string_ string, isFilename bool)
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
	SetDelegate0(value objc.IObject)
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

func (a_ Application) RegisterUserInterfaceItemSearchHandler(handler UserInterfaceItemSearching) {
	po := ffi.CreateProtocol("NSUserInterfaceItemSearching", handler)
	defer po.Release()
	ffi.CallMethod[ffi.Void](a_, "registerUserInterfaceItemSearchHandler:", po)
}

func (a_ Application) RegisterUserInterfaceItemSearchHandler0(handler objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "registerUserInterfaceItemSearchHandler:", handler)
}

func (a_ Application) SearchString_InUserInterfaceItemString_SearchRange_FoundRange(searchString string, stringToSearch string, searchRange foundation.Range, foundRange *foundation.Range) bool {
	rv := ffi.CallMethod[bool](a_, "searchString:inUserInterfaceItemString:searchRange:foundRange:", searchString, stringToSearch, searchRange, foundRange)
	return rv
}

func (a_ Application) UnregisterUserInterfaceItemSearchHandler(handler UserInterfaceItemSearching) {
	po := ffi.CreateProtocol("NSUserInterfaceItemSearching", handler)
	defer po.Release()
	ffi.CallMethod[ffi.Void](a_, "unregisterUserInterfaceItemSearchHandler:", po)
}

func (a_ Application) UnregisterUserInterfaceItemSearchHandler0(handler objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "unregisterUserInterfaceItemSearchHandler:", handler)
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

func (a_ Application) RestoreWindowWithIdentifier_State_CompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(param1 Window, param2 foundation.Error)) bool {
	rv := ffi.CallMethod[bool](a_, "restoreWindowWithIdentifier:state:completionHandler:", identifier, state, completionHandler)
	return rv
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

func (a_ Application) AddWindowsItem_Title_Filename(win IWindow, string_ string, isFilename bool) {
	ffi.CallMethod[ffi.Void](a_, "addWindowsItem:title:filename:", win, string_, isFilename)
}

func (a_ Application) ChangeWindowsItem_Title_Filename(win IWindow, string_ string, isFilename bool) {
	ffi.CallMethod[ffi.Void](a_, "changeWindowsItem:title:filename:", win, string_, isFilename)
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
	po := ffi.CreateProtocol("NSApplicationDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(a_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](a_, "setDelegate:", po)
}

func (a_ Application) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](a_, "setDelegate:", value)
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
