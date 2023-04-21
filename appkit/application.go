// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[Application](a_, "init")
	return rv
}

func (ac _ApplicationClass) Alloc() Application {
	rv := objc.CallMethod[Application](ac, "alloc")
	return rv
}

func (ac _ApplicationClass) New() Application {
	rv := objc.CallMethod[Application](ac, "new")
	rv.Autorelease()
	return rv
}

func NewApplication() Application {
	return ApplicationClass.New()
}

func (a_ Application) NextEventMatchingMask_UntilDate_InMode_Dequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event {
	rv := objc.CallMethod[Event](a_, "nextEventMatchingMask:untilDate:inMode:dequeue:", mask, expiration, mode, deqFlag)
	return rv
}

func (a_ Application) DiscardEventsMatchingMask_BeforeEvent(mask EventMask, lastEvent IEvent) {
	objc.CallMethod[objc.Void](a_, "discardEventsMatchingMask:beforeEvent:", mask, lastEvent)
}

func (a_ Application) Run() {
	objc.CallMethod[objc.Void](a_, "run")
}

func (a_ Application) FinishLaunching() {
	objc.CallMethod[objc.Void](a_, "finishLaunching")
}

func (a_ Application) Stop(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "stop:", sender)
}

func (a_ Application) SendEvent(event IEvent) {
	objc.CallMethod[objc.Void](a_, "sendEvent:", event)
}

func (a_ Application) PostEvent_AtStart(event IEvent, flag bool) {
	objc.CallMethod[objc.Void](a_, "postEvent:atStart:", event, flag)
}

func (a_ Application) SendAction_To_From(action objc.Selector, target objc.IObject, sender objc.IObject) bool {
	rv := objc.CallMethod[bool](a_, "sendAction:to:from:", action, target, sender)
	return rv
}

func (a_ Application) TargetForAction(action objc.Selector) objc.Object {
	rv := objc.CallMethod[objc.Object](a_, "targetForAction:", action)
	return rv
}

func (a_ Application) TargetForAction_To_From(action objc.Selector, target objc.IObject, sender objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](a_, "targetForAction:to:from:", action, target, sender)
	return rv
}

func (a_ Application) Terminate(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "terminate:", sender)
}

func (a_ Application) ReplyToApplicationShouldTerminate(shouldTerminate bool) {
	objc.CallMethod[objc.Void](a_, "replyToApplicationShouldTerminate:", shouldTerminate)
}

func (a_ Application) ActivateIgnoringOtherApps(flag bool) {
	objc.CallMethod[objc.Void](a_, "activateIgnoringOtherApps:", flag)
}

func (a_ Application) Deactivate() {
	objc.CallMethod[objc.Void](a_, "deactivate")
}

func (a_ Application) DisableRelaunchOnLogin() {
	objc.CallMethod[objc.Void](a_, "disableRelaunchOnLogin")
}

func (a_ Application) EnableRelaunchOnLogin() {
	objc.CallMethod[objc.Void](a_, "enableRelaunchOnLogin")
}

func (a_ Application) RegisterForRemoteNotifications() {
	objc.CallMethod[objc.Void](a_, "registerForRemoteNotifications")
}

func (a_ Application) UnregisterForRemoteNotifications() {
	objc.CallMethod[objc.Void](a_, "unregisterForRemoteNotifications")
}

func (a_ Application) RegisterForRemoteNotificationTypes(types RemoteNotificationType) {
	objc.CallMethod[objc.Void](a_, "registerForRemoteNotificationTypes:", types)
}

func (a_ Application) ToggleTouchBarCustomizationPalette(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "toggleTouchBarCustomizationPalette:", sender)
}

func (a_ Application) RequestUserAttention(requestType RequestUserAttentionType) int {
	rv := objc.CallMethod[int](a_, "requestUserAttention:", requestType)
	return rv
}

func (a_ Application) CancelUserAttentionRequest(request int) {
	objc.CallMethod[objc.Void](a_, "cancelUserAttentionRequest:", request)
}

func (a_ Application) ReplyToOpenOrPrint(reply ApplicationDelegateReply) {
	objc.CallMethod[objc.Void](a_, "replyToOpenOrPrint:", reply)
}

func (a_ Application) RegisterUserInterfaceItemSearchHandler(handler UserInterfaceItemSearching) {
	po := objc.CreateProtocol("NSUserInterfaceItemSearching", handler)
	defer po.Release()
	objc.CallMethod[objc.Void](a_, "registerUserInterfaceItemSearchHandler:", po)
}

func (a_ Application) RegisterUserInterfaceItemSearchHandler0(handler objc.IObject) {
	objc.CallMethod[objc.Void](a_, "registerUserInterfaceItemSearchHandler:", handler)
}

func (a_ Application) SearchString_InUserInterfaceItemString_SearchRange_FoundRange(searchString string, stringToSearch string, searchRange foundation.Range, foundRange *foundation.Range) bool {
	rv := objc.CallMethod[bool](a_, "searchString:inUserInterfaceItemString:searchRange:foundRange:", searchString, stringToSearch, searchRange, foundRange)
	return rv
}

func (a_ Application) UnregisterUserInterfaceItemSearchHandler(handler UserInterfaceItemSearching) {
	po := objc.CreateProtocol("NSUserInterfaceItemSearching", handler)
	defer po.Release()
	objc.CallMethod[objc.Void](a_, "unregisterUserInterfaceItemSearchHandler:", po)
}

func (a_ Application) UnregisterUserInterfaceItemSearchHandler0(handler objc.IObject) {
	objc.CallMethod[objc.Void](a_, "unregisterUserInterfaceItemSearchHandler:", handler)
}

func (a_ Application) ShowHelp(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "showHelp:", sender)
}

func (a_ Application) ActivateContextHelpMode(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "activateContextHelpMode:", sender)
}

func (a_ Application) HideOtherApplications(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "hideOtherApplications:", sender)
}

func (a_ Application) UnhideAllApplications(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "unhideAllApplications:", sender)
}

func (ac _ApplicationClass) DetachDrawingThread_ToTarget_WithObject(selector objc.Selector, target objc.IObject, argument objc.IObject) {
	objc.CallMethod[objc.Void](ac, "detachDrawingThread:toTarget:withObject:", selector, target, argument)
}

func (a_ Application) ReportException(exception foundation.IException) {
	objc.CallMethod[objc.Void](a_, "reportException:", exception)
}

func (a_ Application) ActivationPolicy() ApplicationActivationPolicy {
	rv := objc.CallMethod[ApplicationActivationPolicy](a_, "activationPolicy")
	return rv
}

func (a_ Application) SetActivationPolicy(activationPolicy ApplicationActivationPolicy) bool {
	rv := objc.CallMethod[bool](a_, "setActivationPolicy:", activationPolicy)
	return rv
}

func (a_ Application) WindowWithWindowNumber(windowNum int) Window {
	rv := objc.CallMethod[Window](a_, "windowWithWindowNumber:", windowNum)
	return rv
}

// deprecated
func (a_ Application) MakeWindowsPerform_InOrder(selector objc.Selector, flag bool) Window {
	rv := objc.CallMethod[Window](a_, "makeWindowsPerform:inOrder:", selector, flag)
	return rv
}

func (a_ Application) EnumerateWindowsWithOptions_UsingBlock(options WindowListOptions, block func(window Window, stop *bool)) {
	objc.CallMethod[objc.Void](a_, "enumerateWindowsWithOptions:usingBlock:", options, block)
}

func (a_ Application) MiniaturizeAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "miniaturizeAll:", sender)
}

func (a_ Application) Hide(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "hide:", sender)
}

func (a_ Application) Unhide(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "unhide:", sender)
}

func (a_ Application) UnhideWithoutActivation() {
	objc.CallMethod[objc.Void](a_, "unhideWithoutActivation")
}

func (a_ Application) UpdateWindows() {
	objc.CallMethod[objc.Void](a_, "updateWindows")
}

func (a_ Application) SetWindowsNeedUpdate(needUpdate bool) {
	objc.CallMethod[objc.Void](a_, "setWindowsNeedUpdate:", needUpdate)
}

func (a_ Application) PreventWindowOrdering() {
	objc.CallMethod[objc.Void](a_, "preventWindowOrdering")
}

func (a_ Application) ArrangeInFront(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "arrangeInFront:", sender)
}

func (a_ Application) ExtendStateRestoration() {
	objc.CallMethod[objc.Void](a_, "extendStateRestoration")
}

func (a_ Application) CompleteStateRestoration() {
	objc.CallMethod[objc.Void](a_, "completeStateRestoration")
}

func (a_ Application) RestoreWindowWithIdentifier_State_CompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(param1 Window, param2 foundation.Error)) bool {
	rv := objc.CallMethod[bool](a_, "restoreWindowWithIdentifier:state:completionHandler:", identifier, state, completionHandler)
	return rv
}

func (a_ Application) RunModalForWindow(window IWindow) ModalResponse {
	rv := objc.CallMethod[ModalResponse](a_, "runModalForWindow:", window)
	return rv
}

func (a_ Application) StopModal() {
	objc.CallMethod[objc.Void](a_, "stopModal")
}

func (a_ Application) StopModalWithCode(returnCode ModalResponse) {
	objc.CallMethod[objc.Void](a_, "stopModalWithCode:", returnCode)
}

func (a_ Application) AbortModal() {
	objc.CallMethod[objc.Void](a_, "abortModal")
}

func (a_ Application) BeginModalSessionForWindow(window IWindow) ModalSession {
	rv := objc.CallMethod[ModalSession](a_, "beginModalSessionForWindow:", window)
	return rv
}

func (a_ Application) RunModalSession(session ModalSession) ModalResponse {
	rv := objc.CallMethod[ModalResponse](a_, "runModalSession:", session)
	return rv
}

func (a_ Application) OrderFrontColorPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "orderFrontColorPanel:", sender)
}

func (a_ Application) OrderFrontStandardAboutPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "orderFrontStandardAboutPanel:", sender)
}

func (a_ Application) OrderFrontStandardAboutPanelWithOptions(optionsDictionary map[AboutPanelOptionKey]objc.IObject) {
	objc.CallMethod[objc.Void](a_, "orderFrontStandardAboutPanelWithOptions:", optionsDictionary)
}

func (a_ Application) OrderFrontCharacterPalette(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "orderFrontCharacterPalette:", sender)
}

func (a_ Application) RunPageLayout(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, "runPageLayout:", sender)
}

func (a_ Application) AddWindowsItem_Title_Filename(win IWindow, string_ string, isFilename bool) {
	objc.CallMethod[objc.Void](a_, "addWindowsItem:title:filename:", win, string_, isFilename)
}

func (a_ Application) ChangeWindowsItem_Title_Filename(win IWindow, string_ string, isFilename bool) {
	objc.CallMethod[objc.Void](a_, "changeWindowsItem:title:filename:", win, string_, isFilename)
}

func (a_ Application) RemoveWindowsItem(win IWindow) {
	objc.CallMethod[objc.Void](a_, "removeWindowsItem:", win)
}

func (a_ Application) UpdateWindowsItem(win IWindow) {
	objc.CallMethod[objc.Void](a_, "updateWindowsItem:", win)
}

func (a_ Application) RegisterServicesMenuSendTypes_ReturnTypes(sendTypes []PasteboardType, returnTypes []PasteboardType) {
	objc.CallMethod[objc.Void](a_, "registerServicesMenuSendTypes:returnTypes:", sendTypes, returnTypes)
}

// deprecated
func (a_ Application) BeginModalSessionForWindow_RelativeToWindow(window IWindow, docWindow IWindow) ModalSession {
	rv := objc.CallMethod[ModalSession](a_, "beginModalSessionForWindow:relativeToWindow:", window, docWindow)
	return rv
}

// deprecated
func (a_ Application) RunModalForWindow_RelativeToWindow(window IWindow, docWindow IWindow) int {
	rv := objc.CallMethod[int](a_, "runModalForWindow:relativeToWindow:", window, docWindow)
	return rv
}

func (a_ Application) EndModalSession(session ModalSession) {
	objc.CallMethod[objc.Void](a_, "endModalSession:", session)
}

// deprecated
func (a_ Application) BeginSheet_ModalForWindow_ModalDelegate_DidEndSelector_ContextInfo(sheet IWindow, docWindow IWindow, modalDelegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](a_, "beginSheet:modalForWindow:modalDelegate:didEndSelector:contextInfo:", sheet, docWindow, modalDelegate, didEndSelector, contextInfo)
}

// deprecated
func (a_ Application) EndSheet(sheet IWindow) {
	objc.CallMethod[objc.Void](a_, "endSheet:", sheet)
}

// deprecated
func (a_ Application) EndSheet_ReturnCode(sheet IWindow, returnCode int) {
	objc.CallMethod[objc.Void](a_, "endSheet:returnCode:", sheet, returnCode)
}

// deprecated
func (a_ Application) Application_PrintFiles(sender IApplication, filenames []string) {
	objc.CallMethod[objc.Void](a_, "application:printFiles:", sender, filenames)
}

// deprecated
func (a_ Application) Application_DelegateHandlesKey(sender IApplication, key string) bool {
	rv := objc.CallMethod[bool](a_, "application:delegateHandlesKey:", sender, key)
	return rv
}

func (ac _ApplicationClass) SharedApplication() Application {
	rv := objc.CallMethod[Application](ac, "sharedApplication")
	return rv
}

func (a_ Application) Delegate() ApplicationDelegateWrapper {
	rv := objc.CallMethod[ApplicationDelegateWrapper](a_, "delegate")
	return rv
}

func (a_ Application) SetDelegate(value ApplicationDelegate) {
	po := objc.CreateProtocol("NSApplicationDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(a_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](a_, "setDelegate:", po)
}

func (a_ Application) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, "setDelegate:", value)
}

func (a_ Application) CurrentEvent() Event {
	rv := objc.CallMethod[Event](a_, "currentEvent")
	return rv
}

func (a_ Application) IsRunning() bool {
	rv := objc.CallMethod[bool](a_, "isRunning")
	return rv
}

func (a_ Application) IsActive() bool {
	rv := objc.CallMethod[bool](a_, "isActive")
	return rv
}

func (a_ Application) EnabledRemoteNotificationTypes() RemoteNotificationType {
	rv := objc.CallMethod[RemoteNotificationType](a_, "enabledRemoteNotificationTypes")
	return rv
}

func (a_ Application) IsRegisteredForRemoteNotifications() bool {
	rv := objc.CallMethod[bool](a_, "isRegisteredForRemoteNotifications")
	return rv
}

func (a_ Application) Appearance() Appearance {
	rv := objc.CallMethod[Appearance](a_, "appearance")
	return rv
}

func (a_ Application) SetAppearance(value IAppearance) {
	objc.CallMethod[objc.Void](a_, "setAppearance:", value)
}

func (a_ Application) EffectiveAppearance() Appearance {
	rv := objc.CallMethod[Appearance](a_, "effectiveAppearance")
	return rv
}

func (a_ Application) CurrentSystemPresentationOptions() ApplicationPresentationOptions {
	rv := objc.CallMethod[ApplicationPresentationOptions](a_, "currentSystemPresentationOptions")
	return rv
}

func (a_ Application) PresentationOptions() ApplicationPresentationOptions {
	rv := objc.CallMethod[ApplicationPresentationOptions](a_, "presentationOptions")
	return rv
}

func (a_ Application) SetPresentationOptions(value ApplicationPresentationOptions) {
	objc.CallMethod[objc.Void](a_, "setPresentationOptions:", value)
}

func (a_ Application) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](a_, "userInterfaceLayoutDirection")
	return rv
}

func (a_ Application) DockTile() DockTile {
	rv := objc.CallMethod[DockTile](a_, "dockTile")
	return rv
}

func (a_ Application) ApplicationIconImage() Image {
	rv := objc.CallMethod[Image](a_, "applicationIconImage")
	return rv
}

func (a_ Application) SetApplicationIconImage(value IImage) {
	objc.CallMethod[objc.Void](a_, "setApplicationIconImage:", value)
}

func (a_ Application) HelpMenu() Menu {
	rv := objc.CallMethod[Menu](a_, "helpMenu")
	return rv
}

func (a_ Application) SetHelpMenu(value IMenu) {
	objc.CallMethod[objc.Void](a_, "setHelpMenu:", value)
}

func (a_ Application) ServicesProvider() objc.Object {
	rv := objc.CallMethod[objc.Object](a_, "servicesProvider")
	return rv
}

func (a_ Application) SetServicesProvider(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, "setServicesProvider:", value)
}

func (a_ Application) IsFullKeyboardAccessEnabled() bool {
	rv := objc.CallMethod[bool](a_, "isFullKeyboardAccessEnabled")
	return rv
}

func (a_ Application) OrderedDocuments() []Document {
	rv := objc.CallMethod[[]Document](a_, "orderedDocuments")
	return rv
}

func (a_ Application) OrderedWindows() []Window {
	rv := objc.CallMethod[[]Window](a_, "orderedWindows")
	return rv
}

func (a_ Application) KeyWindow() Window {
	rv := objc.CallMethod[Window](a_, "keyWindow")
	return rv
}

func (a_ Application) MainWindow() Window {
	rv := objc.CallMethod[Window](a_, "mainWindow")
	return rv
}

func (a_ Application) Windows() []Window {
	rv := objc.CallMethod[[]Window](a_, "windows")
	return rv
}

func (a_ Application) IsHidden() bool {
	rv := objc.CallMethod[bool](a_, "isHidden")
	return rv
}

// deprecated
func (a_ Application) Context() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](a_, "context")
	return rv
}

func (a_ Application) OcclusionState() ApplicationOcclusionState {
	rv := objc.CallMethod[ApplicationOcclusionState](a_, "occlusionState")
	return rv
}

func (a_ Application) IsProtectedDataAvailable() bool {
	rv := objc.CallMethod[bool](a_, "isProtectedDataAvailable")
	return rv
}

func (a_ Application) ModalWindow() Window {
	rv := objc.CallMethod[Window](a_, "modalWindow")
	return rv
}

func (a_ Application) MainMenu() Menu {
	rv := objc.CallMethod[Menu](a_, "mainMenu")
	return rv
}

func (a_ Application) SetMainMenu(value IMenu) {
	objc.CallMethod[objc.Void](a_, "setMainMenu:", value)
}

func (a_ Application) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.CallMethod[bool](a_, "isAutomaticCustomizeTouchBarMenuItemEnabled")
	return rv
}

func (a_ Application) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.CallMethod[objc.Void](a_, "setAutomaticCustomizeTouchBarMenuItemEnabled:", value)
}

func (a_ Application) WindowsMenu() Menu {
	rv := objc.CallMethod[Menu](a_, "windowsMenu")
	return rv
}

func (a_ Application) SetWindowsMenu(value IMenu) {
	objc.CallMethod[objc.Void](a_, "setWindowsMenu:", value)
}

func (a_ Application) ServicesMenu() Menu {
	rv := objc.CallMethod[Menu](a_, "servicesMenu")
	return rv
}

func (a_ Application) SetServicesMenu(value IMenu) {
	objc.CallMethod[objc.Void](a_, "setServicesMenu:", value)
}
