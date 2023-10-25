// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	ImplementsApplication_WillEncodeRestorableState() bool
	// optional
	Application_WillEncodeRestorableState(app Application, coder foundation.Coder)
	ImplementsApplication_DidDecodeRestorableState() bool
	// optional
	Application_DidDecodeRestorableState(app Application, coder foundation.Coder)
	ImplementsApplicationDidChangeOcclusionState() bool
	// optional
	ApplicationDidChangeOcclusionState(notification foundation.Notification)
	ImplementsApplication_DelegateHandlesKey() bool
	// optional
	Application_DelegateHandlesKey(sender Application, key string) bool
}

func WrapApplicationDelegate(v ApplicationDelegate) objc.Object {
	return objc.WrapAsProtocol("NSApplicationDelegate", v)
}

type ApplicationDelegateBase struct {
}

func (p *ApplicationDelegateBase) ImplementsApplicationWillFinishLaunching() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationWillFinishLaunching(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationDidFinishLaunching() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationDidFinishLaunching(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationWillBecomeActive() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationWillBecomeActive(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationDidBecomeActive() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationDidBecomeActive(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationWillResignActive() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationWillResignActive(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationDidResignActive() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationDidResignActive(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationShouldTerminate() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationShouldTerminate(sender Application) ApplicationTerminateReply {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationShouldTerminateAfterLastWindowClosed(sender Application) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationWillTerminate() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationWillTerminate(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationWillHide() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationWillHide(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationDidHide() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationDidHide(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationWillUnhide() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationWillUnhide(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationDidUnhide() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationDidUnhide(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationWillUpdate() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationWillUpdate(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationDidUpdate() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationDidUpdate(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationShouldHandleReopen_HasVisibleWindows() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationShouldHandleReopen_HasVisibleWindows(sender Application, flag bool) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationDockMenu() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationDockMenu(sender Application) IMenu {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application Application) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_WillPresentError() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_WillPresentError(application Application, error foundation.Error) foundation.IError {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationDidChangeScreenParameters() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationDidChangeScreenParameters(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_WillContinueUserActivityWithType() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_WillContinueUserActivityWithType(application Application, userActivityType string) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_DidFailToContinueUserActivityWithType_Error() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_DidFailToContinueUserActivityWithType_Error(application Application, userActivityType string, error foundation.Error) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_DidUpdateUserActivity() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_DidUpdateUserActivity(application Application, userActivity foundation.UserActivity) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_DidRegisterForRemoteNotificationsWithDeviceToken() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_DidRegisterForRemoteNotificationsWithDeviceToken(application Application, deviceToken []byte) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_DidFailToRegisterForRemoteNotificationsWithError() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_DidFailToRegisterForRemoteNotificationsWithError(application Application, error foundation.Error) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_DidReceiveRemoteNotification() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_DidReceiveRemoteNotification(application Application, userInfo map[string]objc.Object) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_OpenURLs() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_OpenURLs(application Application, urls []foundation.URL) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_OpenFile() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_OpenFile(sender Application, filename string) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_OpenFileWithoutUI() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_OpenFileWithoutUI(sender objc.Object, filename string) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_OpenTempFile() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_OpenTempFile(sender Application, filename string) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_OpenFiles() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_OpenFiles(sender Application, filenames []string) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationShouldOpenUntitledFile() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationShouldOpenUntitledFile(sender Application) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationOpenUntitledFile() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationOpenUntitledFile(sender Application) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_PrintFile() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_PrintFile(sender Application, filename string) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_PrintFiles_WithSettings_ShowPrintPanels() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_PrintFiles_WithSettings_ShowPrintPanels(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationSupportsSecureRestorableState() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationSupportsSecureRestorableState(app Application) bool {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationProtectedDataDidBecomeAvailable() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationProtectedDataDidBecomeAvailable(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationProtectedDataWillBecomeUnavailable() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_WillEncodeRestorableState() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_WillEncodeRestorableState(app Application, coder foundation.Coder) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_DidDecodeRestorableState() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_DidDecodeRestorableState(app Application, coder foundation.Coder) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplicationDidChangeOcclusionState() bool {
	return false
}

func (p *ApplicationDelegateBase) ApplicationDidChangeOcclusionState(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *ApplicationDelegateBase) ImplementsApplication_DelegateHandlesKey() bool {
	return false
}

func (p *ApplicationDelegateBase) Application_DelegateHandlesKey(sender Application, key string) bool {
	panic("unimpemented protocol method")
}

type ApplicationDelegateCreator struct {
	className string
	class     objc.Class
}

func NewApplicationDelegateCreator(name string) *ApplicationDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &ApplicationDelegateCreator{className: name, class: class}
}

func (c *ApplicationDelegateCreator) SetApplicationWillFinishLaunching(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationWillFinishLaunching:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationDidFinishLaunching(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationDidFinishLaunching:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationWillBecomeActive(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationWillBecomeActive:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationDidBecomeActive(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationDidBecomeActive:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationWillResignActive(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationWillResignActive:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationDidResignActive(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationDidResignActive:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationShouldTerminate(handle func(o objc.Object, sender Application) ApplicationTerminateReply) {
	objc.AddMethod(c.class, objc.GetSelector("applicationShouldTerminate:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationShouldTerminateAfterLastWindowClosed(handle func(o objc.Object, sender Application) bool) {
	objc.AddMethod(c.class, objc.GetSelector("applicationShouldTerminateAfterLastWindowClosed:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationWillTerminate(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationWillTerminate:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationWillHide(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationWillHide:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationDidHide(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationDidHide:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationWillUnhide(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationWillUnhide:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationDidUnhide(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationDidUnhide:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationWillUpdate(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationWillUpdate:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationDidUpdate(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationDidUpdate:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationShouldHandleReopen_HasVisibleWindows(handle func(o objc.Object, sender Application, flag bool) bool) {
	objc.AddMethod(c.class, objc.GetSelector("applicationShouldHandleReopen:hasVisibleWindows:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationDockMenu(handle func(o objc.Object, sender Application) IMenu) {
	objc.AddMethod(c.class, objc.GetSelector("applicationDockMenu:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationShouldAutomaticallyLocalizeKeyEquivalents(handle func(o objc.Object, application Application) bool) {
	objc.AddMethod(c.class, objc.GetSelector("applicationShouldAutomaticallyLocalizeKeyEquivalents:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_WillPresentError(handle func(o objc.Object, application Application, error foundation.Error) foundation.IError) {
	objc.AddMethod(c.class, objc.GetSelector("application:willPresentError:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationDidChangeScreenParameters(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationDidChangeScreenParameters:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_WillContinueUserActivityWithType(handle func(o objc.Object, application Application, userActivityType string) bool) {
	objc.AddMethod(c.class, objc.GetSelector("application:willContinueUserActivityWithType:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_DidFailToContinueUserActivityWithType_Error(handle func(o objc.Object, application Application, userActivityType string, error foundation.Error)) {
	objc.AddMethod(c.class, objc.GetSelector("application:didFailToContinueUserActivityWithType:error:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_DidUpdateUserActivity(handle func(o objc.Object, application Application, userActivity foundation.UserActivity)) {
	objc.AddMethod(c.class, objc.GetSelector("application:didUpdateUserActivity:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_DidRegisterForRemoteNotificationsWithDeviceToken(handle func(o objc.Object, application Application, deviceToken []byte)) {
	objc.AddMethod(c.class, objc.GetSelector("application:didRegisterForRemoteNotificationsWithDeviceToken:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_DidFailToRegisterForRemoteNotificationsWithError(handle func(o objc.Object, application Application, error foundation.Error)) {
	objc.AddMethod(c.class, objc.GetSelector("application:didFailToRegisterForRemoteNotificationsWithError:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_DidReceiveRemoteNotification(handle func(o objc.Object, application Application, userInfo map[string]objc.Object)) {
	objc.AddMethod(c.class, objc.GetSelector("application:didReceiveRemoteNotification:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_OpenURLs(handle func(o objc.Object, application Application, urls []foundation.URL)) {
	objc.AddMethod(c.class, objc.GetSelector("application:openURLs:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_OpenFile(handle func(o objc.Object, sender Application, filename string) bool) {
	objc.AddMethod(c.class, objc.GetSelector("application:openFile:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_OpenFileWithoutUI(handle func(o objc.Object, sender objc.Object, filename string) bool) {
	objc.AddMethod(c.class, objc.GetSelector("application:openFileWithoutUI:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_OpenTempFile(handle func(o objc.Object, sender Application, filename string) bool) {
	objc.AddMethod(c.class, objc.GetSelector("application:openTempFile:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_OpenFiles(handle func(o objc.Object, sender Application, filenames []string)) {
	objc.AddMethod(c.class, objc.GetSelector("application:openFiles:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationShouldOpenUntitledFile(handle func(o objc.Object, sender Application) bool) {
	objc.AddMethod(c.class, objc.GetSelector("applicationShouldOpenUntitledFile:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationOpenUntitledFile(handle func(o objc.Object, sender Application) bool) {
	objc.AddMethod(c.class, objc.GetSelector("applicationOpenUntitledFile:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_PrintFile(handle func(o objc.Object, sender Application, filename string) bool) {
	objc.AddMethod(c.class, objc.GetSelector("application:printFile:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_PrintFiles_WithSettings_ShowPrintPanels(handle func(o objc.Object, application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply) {
	objc.AddMethod(c.class, objc.GetSelector("application:printFiles:withSettings:showPrintPanels:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationSupportsSecureRestorableState(handle func(o objc.Object, app Application) bool) {
	objc.AddMethod(c.class, objc.GetSelector("applicationSupportsSecureRestorableState:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationProtectedDataDidBecomeAvailable(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationProtectedDataDidBecomeAvailable:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationProtectedDataWillBecomeUnavailable(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationProtectedDataWillBecomeUnavailable:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_WillEncodeRestorableState(handle func(o objc.Object, app Application, coder foundation.Coder)) {
	objc.AddMethod(c.class, objc.GetSelector("application:willEncodeRestorableState:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_DidDecodeRestorableState(handle func(o objc.Object, app Application, coder foundation.Coder)) {
	objc.AddMethod(c.class, objc.GetSelector("application:didDecodeRestorableState:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplicationDidChangeOcclusionState(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.GetSelector("applicationDidChangeOcclusionState:"), handle)
}

func (c *ApplicationDelegateCreator) SetApplication_DelegateHandlesKey(handle func(o objc.Object, sender Application, key string) bool) {
	objc.AddMethod(c.class, objc.GetSelector("application:delegateHandlesKey:"), handle)
}

func (c *ApplicationDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
