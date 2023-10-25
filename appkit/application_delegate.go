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
