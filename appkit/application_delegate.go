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
	_Application_WillEncodeRestorableState                        func(app Application, coder foundation.Coder)
	_Application_DidDecodeRestorableState                         func(app Application, coder foundation.Coder)
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
func (di *ApplicationDelegateImpl) ImplementsApplication_WillEncodeRestorableState() bool {
	return di._Application_WillEncodeRestorableState != nil
}

func (di *ApplicationDelegateImpl) SetApplication_WillEncodeRestorableState(f func(app Application, coder foundation.Coder)) {
	di._Application_WillEncodeRestorableState = f
}

func (di *ApplicationDelegateImpl) Application_WillEncodeRestorableState(app Application, coder foundation.Coder) {
	di._Application_WillEncodeRestorableState(app, coder)
}
func (di *ApplicationDelegateImpl) ImplementsApplication_DidDecodeRestorableState() bool {
	return di._Application_DidDecodeRestorableState != nil
}

func (di *ApplicationDelegateImpl) SetApplication_DidDecodeRestorableState(f func(app Application, coder foundation.Coder)) {
	di._Application_DidDecodeRestorableState = f
}

func (di *ApplicationDelegateImpl) Application_DidDecodeRestorableState(app Application, coder foundation.Coder) {
	di._Application_DidDecodeRestorableState(app, coder)
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
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillFinishLaunching:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidFinishLaunching() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidFinishLaunching:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidFinishLaunching(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidFinishLaunching:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillBecomeActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillBecomeActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillBecomeActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillBecomeActive:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidBecomeActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidBecomeActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidBecomeActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidBecomeActive:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillResignActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillResignActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillResignActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillResignActive:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidResignActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidResignActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidResignActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidResignActive:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldTerminate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldTerminate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminate(sender IApplication) ApplicationTerminateReply {
	rv := objc.CallMethod[ApplicationTerminateReply](a_, objc.GetSelector("applicationShouldTerminate:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldTerminateAfterLastWindowClosed:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminateAfterLastWindowClosed(sender IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldTerminateAfterLastWindowClosed:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillTerminate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillTerminate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillTerminate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillTerminate:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillHide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillHide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillHide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillHide:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidHide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidHide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidHide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidHide:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillUnhide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillUnhide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillUnhide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillUnhide:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidUnhide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidUnhide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidUnhide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidUnhide:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationWillUpdate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillUpdate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillUpdate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillUpdate:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidUpdate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidUpdate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidUpdate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidUpdate:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldHandleReopen_HasVisibleWindows() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldHandleReopen:hasVisibleWindows:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldHandleReopen_HasVisibleWindows(sender IApplication, flag bool) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldHandleReopen:hasVisibleWindows:"), objc.ExtractPtr(sender), flag)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDockMenu() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDockMenu:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDockMenu(sender IApplication) Menu {
	rv := objc.CallMethod[Menu](a_, objc.GetSelector("applicationDockMenu:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldAutomaticallyLocalizeKeyEquivalents:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldAutomaticallyLocalizeKeyEquivalents:"), objc.ExtractPtr(application))
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_WillPresentError() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:willPresentError:"))
}

func (a_ ApplicationDelegateWrapper) Application_WillPresentError(application IApplication, error foundation.IError) foundation.Error {
	rv := objc.CallMethod[foundation.Error](a_, objc.GetSelector("application:willPresentError:"), objc.ExtractPtr(application), objc.ExtractPtr(error))
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidChangeScreenParameters() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidChangeScreenParameters:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidChangeScreenParameters(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidChangeScreenParameters:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_WillContinueUserActivityWithType() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:willContinueUserActivityWithType:"))
}

func (a_ ApplicationDelegateWrapper) Application_WillContinueUserActivityWithType(application IApplication, userActivityType string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:willContinueUserActivityWithType:"), objc.ExtractPtr(application), userActivityType)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidFailToContinueUserActivityWithType_Error() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didFailToContinueUserActivityWithType:error:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidFailToContinueUserActivityWithType_Error(application IApplication, userActivityType string, error foundation.IError) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didFailToContinueUserActivityWithType:error:"), objc.ExtractPtr(application), userActivityType, objc.ExtractPtr(error))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidUpdateUserActivity() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didUpdateUserActivity:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidUpdateUserActivity(application IApplication, userActivity foundation.IUserActivity) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didUpdateUserActivity:"), objc.ExtractPtr(application), objc.ExtractPtr(userActivity))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidRegisterForRemoteNotificationsWithDeviceToken() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didRegisterForRemoteNotificationsWithDeviceToken:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidRegisterForRemoteNotificationsWithDeviceToken(application IApplication, deviceToken []byte) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didRegisterForRemoteNotificationsWithDeviceToken:"), objc.ExtractPtr(application), deviceToken)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidFailToRegisterForRemoteNotificationsWithError() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didFailToRegisterForRemoteNotificationsWithError:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidFailToRegisterForRemoteNotificationsWithError(application IApplication, error foundation.IError) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didFailToRegisterForRemoteNotificationsWithError:"), objc.ExtractPtr(application), objc.ExtractPtr(error))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidReceiveRemoteNotification() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didReceiveRemoteNotification:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidReceiveRemoteNotification(application IApplication, userInfo map[string]objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didReceiveRemoteNotification:"), objc.ExtractPtr(application), userInfo)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenURLs() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openURLs:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenURLs(application IApplication, urls []foundation.IURL) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:openURLs:"), objc.ExtractPtr(application), urls)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openFile:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenFile(sender IApplication, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:openFile:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenFileWithoutUI() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openFileWithoutUI:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenFileWithoutUI(sender objc.IObject, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:openFileWithoutUI:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenTempFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openTempFile:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenTempFile(sender IApplication, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:openTempFile:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_OpenFiles() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openFiles:"))
}

func (a_ ApplicationDelegateWrapper) Application_OpenFiles(sender IApplication, filenames []string) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:openFiles:"), objc.ExtractPtr(sender), filenames)
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationShouldOpenUntitledFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldOpenUntitledFile:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldOpenUntitledFile(sender IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldOpenUntitledFile:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationOpenUntitledFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationOpenUntitledFile:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationOpenUntitledFile(sender IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationOpenUntitledFile:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_PrintFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:printFile:"))
}

func (a_ ApplicationDelegateWrapper) Application_PrintFile(sender IApplication, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:printFile:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_PrintFiles_WithSettings_ShowPrintPanels() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:printFiles:withSettings:showPrintPanels:"))
}

func (a_ ApplicationDelegateWrapper) Application_PrintFiles_WithSettings_ShowPrintPanels(application IApplication, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanels bool) ApplicationPrintReply {
	rv := objc.CallMethod[ApplicationPrintReply](a_, objc.GetSelector("application:printFiles:withSettings:showPrintPanels:"), objc.ExtractPtr(application), fileNames, printSettings, showPrintPanels)
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationSupportsSecureRestorableState() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationSupportsSecureRestorableState:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationSupportsSecureRestorableState(app IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationSupportsSecureRestorableState:"), objc.ExtractPtr(app))
	return rv
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationProtectedDataDidBecomeAvailable() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationProtectedDataDidBecomeAvailable:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataDidBecomeAvailable(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationProtectedDataDidBecomeAvailable:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationProtectedDataWillBecomeUnavailable() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationProtectedDataWillBecomeUnavailable:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationProtectedDataWillBecomeUnavailable:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_WillEncodeRestorableState() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:willEncodeRestorableState:"))
}

func (a_ ApplicationDelegateWrapper) Application_WillEncodeRestorableState(app IApplication, coder foundation.ICoder) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:willEncodeRestorableState:"), objc.ExtractPtr(app), objc.ExtractPtr(coder))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DidDecodeRestorableState() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didDecodeRestorableState:"))
}

func (a_ ApplicationDelegateWrapper) Application_DidDecodeRestorableState(app IApplication, coder foundation.ICoder) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didDecodeRestorableState:"), objc.ExtractPtr(app), objc.ExtractPtr(coder))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplicationDidChangeOcclusionState() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidChangeOcclusionState:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidChangeOcclusionState(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidChangeOcclusionState:"), objc.ExtractPtr(notification))
}

func (a_ *ApplicationDelegateWrapper) ImplementsApplication_DelegateHandlesKey() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:delegateHandlesKey:"))
}

func (a_ ApplicationDelegateWrapper) Application_DelegateHandlesKey(sender IApplication, key string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:delegateHandlesKey:"), objc.ExtractPtr(sender), key)
	return rv
}
