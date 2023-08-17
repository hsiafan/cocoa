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

type ApplicationDelegateWrapper struct {
	objc.Object
}

func (a_ ApplicationDelegateWrapper) ApplicationWillFinishLaunching(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillFinishLaunching:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidFinishLaunching(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidFinishLaunching:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillBecomeActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillBecomeActive:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidBecomeActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidBecomeActive:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillResignActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillResignActive:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidResignActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidResignActive:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminate(sender IApplication) ApplicationTerminateReply {
	rv := objc.CallMethod[ApplicationTerminateReply](a_, objc.GetSelector("applicationShouldTerminate:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminateAfterLastWindowClosed(sender IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldTerminateAfterLastWindowClosed:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) ApplicationWillTerminate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillTerminate:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillHide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillHide:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidHide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidHide:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillUnhide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillUnhide:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidUnhide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidUnhide:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillUpdate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillUpdate:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidUpdate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidUpdate:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldHandleReopen_HasVisibleWindows(sender IApplication, flag bool) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldHandleReopen:hasVisibleWindows:"), objc.ExtractPtr(sender), flag)
	return rv
}

func (a_ ApplicationDelegateWrapper) ApplicationDockMenu(sender IApplication) Menu {
	rv := objc.CallMethod[Menu](a_, objc.GetSelector("applicationDockMenu:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldAutomaticallyLocalizeKeyEquivalents:"), objc.ExtractPtr(application))
	return rv
}

func (a_ ApplicationDelegateWrapper) Application_WillPresentError(application IApplication, error foundation.IError) foundation.Error {
	rv := objc.CallMethod[foundation.Error](a_, objc.GetSelector("application:willPresentError:"), objc.ExtractPtr(application), objc.ExtractPtr(error))
	return rv
}

func (a_ ApplicationDelegateWrapper) ApplicationDidChangeScreenParameters(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidChangeScreenParameters:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) Application_WillContinueUserActivityWithType(application IApplication, userActivityType string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:willContinueUserActivityWithType:"), objc.ExtractPtr(application), userActivityType)
	return rv
}

func (a_ ApplicationDelegateWrapper) Application_DidFailToContinueUserActivityWithType_Error(application IApplication, userActivityType string, error foundation.IError) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didFailToContinueUserActivityWithType:error:"), objc.ExtractPtr(application), userActivityType, objc.ExtractPtr(error))
}

func (a_ ApplicationDelegateWrapper) Application_DidUpdateUserActivity(application IApplication, userActivity foundation.IUserActivity) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didUpdateUserActivity:"), objc.ExtractPtr(application), objc.ExtractPtr(userActivity))
}

func (a_ ApplicationDelegateWrapper) Application_DidRegisterForRemoteNotificationsWithDeviceToken(application IApplication, deviceToken []byte) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didRegisterForRemoteNotificationsWithDeviceToken:"), objc.ExtractPtr(application), deviceToken)
}

func (a_ ApplicationDelegateWrapper) Application_DidFailToRegisterForRemoteNotificationsWithError(application IApplication, error foundation.IError) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didFailToRegisterForRemoteNotificationsWithError:"), objc.ExtractPtr(application), objc.ExtractPtr(error))
}

func (a_ ApplicationDelegateWrapper) Application_DidReceiveRemoteNotification(application IApplication, userInfo map[string]objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didReceiveRemoteNotification:"), objc.ExtractPtr(application), userInfo)
}

func (a_ ApplicationDelegateWrapper) Application_OpenURLs(application IApplication, urls []foundation.IURL) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:openURLs:"), objc.ExtractPtr(application), urls)
}

func (a_ ApplicationDelegateWrapper) Application_OpenFile(sender IApplication, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:openFile:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ ApplicationDelegateWrapper) Application_OpenFileWithoutUI(sender objc.IObject, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:openFileWithoutUI:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ ApplicationDelegateWrapper) Application_OpenTempFile(sender IApplication, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:openTempFile:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ ApplicationDelegateWrapper) Application_OpenFiles(sender IApplication, filenames []string) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:openFiles:"), objc.ExtractPtr(sender), filenames)
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldOpenUntitledFile(sender IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldOpenUntitledFile:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) ApplicationOpenUntitledFile(sender IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationOpenUntitledFile:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) Application_PrintFile(sender IApplication, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:printFile:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ ApplicationDelegateWrapper) Application_PrintFiles_WithSettings_ShowPrintPanels(application IApplication, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanels bool) ApplicationPrintReply {
	rv := objc.CallMethod[ApplicationPrintReply](a_, objc.GetSelector("application:printFiles:withSettings:showPrintPanels:"), objc.ExtractPtr(application), fileNames, printSettings, showPrintPanels)
	return rv
}

func (a_ ApplicationDelegateWrapper) ApplicationSupportsSecureRestorableState(app IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationSupportsSecureRestorableState:"), objc.ExtractPtr(app))
	return rv
}

func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataDidBecomeAvailable(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationProtectedDataDidBecomeAvailable:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationProtectedDataWillBecomeUnavailable:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) Application_WillEncodeRestorableState(app IApplication, coder foundation.ICoder) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:willEncodeRestorableState:"), objc.ExtractPtr(app), objc.ExtractPtr(coder))
}

func (a_ ApplicationDelegateWrapper) Application_DidDecodeRestorableState(app IApplication, coder foundation.ICoder) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didDecodeRestorableState:"), objc.ExtractPtr(app), objc.ExtractPtr(coder))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidChangeOcclusionState(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidChangeOcclusionState:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) Application_DelegateHandlesKey(sender IApplication, key string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:delegateHandlesKey:"), objc.ExtractPtr(sender), key)
	return rv
}
