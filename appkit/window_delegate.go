// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type WindowDelegate interface {
	ImplementsWindow_WillPositionSheet_UsingRect() bool
	// optional
	Window_WillPositionSheet_UsingRect(window Window, sheet Window, rect foundation.Rect) foundation.Rect
	ImplementsWindowWillBeginSheet() bool
	// optional
	WindowWillBeginSheet(notification foundation.Notification)
	ImplementsWindowDidEndSheet() bool
	// optional
	WindowDidEndSheet(notification foundation.Notification)
	ImplementsWindowWillResize_ToSize() bool
	// optional
	WindowWillResize_ToSize(sender Window, frameSize foundation.Size) foundation.Size
	ImplementsWindowDidResize() bool
	// optional
	WindowDidResize(notification foundation.Notification)
	ImplementsWindowWillStartLiveResize() bool
	// optional
	WindowWillStartLiveResize(notification foundation.Notification)
	ImplementsWindowDidEndLiveResize() bool
	// optional
	WindowDidEndLiveResize(notification foundation.Notification)
	ImplementsWindowWillMiniaturize() bool
	// optional
	WindowWillMiniaturize(notification foundation.Notification)
	ImplementsWindowDidMiniaturize() bool
	// optional
	WindowDidMiniaturize(notification foundation.Notification)
	ImplementsWindowDidDeminiaturize() bool
	// optional
	WindowDidDeminiaturize(notification foundation.Notification)
	ImplementsWindowWillUseStandardFrame_DefaultFrame() bool
	// optional
	WindowWillUseStandardFrame_DefaultFrame(window Window, newFrame foundation.Rect) foundation.Rect
	ImplementsWindowShouldZoom_ToFrame() bool
	// optional
	WindowShouldZoom_ToFrame(window Window, newFrame foundation.Rect) bool
	ImplementsWindow_WillUseFullScreenContentSize() bool
	// optional
	Window_WillUseFullScreenContentSize(window Window, proposedSize foundation.Size) foundation.Size
	ImplementsWindow_WillUseFullScreenPresentationOptions() bool
	// optional
	Window_WillUseFullScreenPresentationOptions(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	ImplementsWindowWillEnterFullScreen() bool
	// optional
	WindowWillEnterFullScreen(notification foundation.Notification)
	ImplementsWindowDidEnterFullScreen() bool
	// optional
	WindowDidEnterFullScreen(notification foundation.Notification)
	ImplementsWindowWillExitFullScreen() bool
	// optional
	WindowWillExitFullScreen(notification foundation.Notification)
	ImplementsWindowDidExitFullScreen() bool
	// optional
	WindowDidExitFullScreen(notification foundation.Notification)
	ImplementsCustomWindowsToEnterFullScreenForWindow() bool
	// optional
	CustomWindowsToEnterFullScreenForWindow(window Window) []IWindow
	ImplementsCustomWindowsToEnterFullScreenForWindow_OnScreen() bool
	// optional
	CustomWindowsToEnterFullScreenForWindow_OnScreen(window Window, screen Screen) []IWindow
	ImplementsWindow_StartCustomAnimationToEnterFullScreenWithDuration() bool
	// optional
	Window_StartCustomAnimationToEnterFullScreenWithDuration(window Window, duration foundation.TimeInterval)
	ImplementsWindow_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration() bool
	// optional
	Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(window Window, screen Screen, duration foundation.TimeInterval)
	ImplementsWindowDidFailToEnterFullScreen() bool
	// optional
	WindowDidFailToEnterFullScreen(window Window)
	ImplementsCustomWindowsToExitFullScreenForWindow() bool
	// optional
	CustomWindowsToExitFullScreenForWindow(window Window) []IWindow
	ImplementsWindow_StartCustomAnimationToExitFullScreenWithDuration() bool
	// optional
	Window_StartCustomAnimationToExitFullScreenWithDuration(window Window, duration foundation.TimeInterval)
	ImplementsWindowDidFailToExitFullScreen() bool
	// optional
	WindowDidFailToExitFullScreen(window Window)
	ImplementsWindowWillMove() bool
	// optional
	WindowWillMove(notification foundation.Notification)
	ImplementsWindowDidMove() bool
	// optional
	WindowDidMove(notification foundation.Notification)
	ImplementsWindowDidChangeScreen() bool
	// optional
	WindowDidChangeScreen(notification foundation.Notification)
	ImplementsWindowDidChangeScreenProfile() bool
	// optional
	WindowDidChangeScreenProfile(notification foundation.Notification)
	ImplementsWindowDidChangeBackingProperties() bool
	// optional
	WindowDidChangeBackingProperties(notification foundation.Notification)
	ImplementsWindowShouldClose() bool
	// optional
	WindowShouldClose(sender Window) bool
	ImplementsWindowWillClose() bool
	// optional
	WindowWillClose(notification foundation.Notification)
	ImplementsWindowDidBecomeKey() bool
	// optional
	WindowDidBecomeKey(notification foundation.Notification)
	ImplementsWindowDidResignKey() bool
	// optional
	WindowDidResignKey(notification foundation.Notification)
	ImplementsWindowDidBecomeMain() bool
	// optional
	WindowDidBecomeMain(notification foundation.Notification)
	ImplementsWindowDidResignMain() bool
	// optional
	WindowDidResignMain(notification foundation.Notification)
	ImplementsWindowWillReturnFieldEditor_ToObject() bool
	// optional
	WindowWillReturnFieldEditor_ToObject(sender Window, client objc.Object) objc.IObject
	ImplementsWindowDidUpdate() bool
	// optional
	WindowDidUpdate(notification foundation.Notification)
	ImplementsWindowDidExpose() bool
	// optional
	WindowDidExpose(notification foundation.Notification)
	ImplementsWindowDidChangeOcclusionState() bool
	// optional
	WindowDidChangeOcclusionState(notification foundation.Notification)
	ImplementsWindow_ShouldDragDocumentWithEvent_From_WithPasteboard() bool
	// optional
	Window_ShouldDragDocumentWithEvent_From_WithPasteboard(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool
	ImplementsWindowWillReturnUndoManager() bool
	// optional
	WindowWillReturnUndoManager(window Window) foundation.IUndoManager
	ImplementsWindow_ShouldPopUpDocumentPathMenu() bool
	// optional
	Window_ShouldPopUpDocumentPathMenu(window Window, menu Menu) bool
	ImplementsWindow_WillEncodeRestorableState() bool
	// optional
	Window_WillEncodeRestorableState(window Window, state foundation.Coder)
	ImplementsWindow_DidDecodeRestorableState() bool
	// optional
	Window_DidDecodeRestorableState(window Window, state foundation.Coder)
	ImplementsWindow_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize() bool
	// optional
	Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size
	ImplementsWindowWillEnterVersionBrowser() bool
	// optional
	WindowWillEnterVersionBrowser(notification foundation.Notification)
	ImplementsWindowDidEnterVersionBrowser() bool
	// optional
	WindowDidEnterVersionBrowser(notification foundation.Notification)
	ImplementsWindowWillExitVersionBrowser() bool
	// optional
	WindowWillExitVersionBrowser(notification foundation.Notification)
	ImplementsWindowDidExitVersionBrowser() bool
	// optional
	WindowDidExitVersionBrowser(notification foundation.Notification)
}

type WindowDelegateImpl struct {
	_Window_WillPositionSheet_UsingRect                                    func(window Window, sheet Window, rect foundation.Rect) foundation.Rect
	_WindowWillBeginSheet                                                  func(notification foundation.Notification)
	_WindowDidEndSheet                                                     func(notification foundation.Notification)
	_WindowWillResize_ToSize                                               func(sender Window, frameSize foundation.Size) foundation.Size
	_WindowDidResize                                                       func(notification foundation.Notification)
	_WindowWillStartLiveResize                                             func(notification foundation.Notification)
	_WindowDidEndLiveResize                                                func(notification foundation.Notification)
	_WindowWillMiniaturize                                                 func(notification foundation.Notification)
	_WindowDidMiniaturize                                                  func(notification foundation.Notification)
	_WindowDidDeminiaturize                                                func(notification foundation.Notification)
	_WindowWillUseStandardFrame_DefaultFrame                               func(window Window, newFrame foundation.Rect) foundation.Rect
	_WindowShouldZoom_ToFrame                                              func(window Window, newFrame foundation.Rect) bool
	_Window_WillUseFullScreenContentSize                                   func(window Window, proposedSize foundation.Size) foundation.Size
	_Window_WillUseFullScreenPresentationOptions                           func(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	_WindowWillEnterFullScreen                                             func(notification foundation.Notification)
	_WindowDidEnterFullScreen                                              func(notification foundation.Notification)
	_WindowWillExitFullScreen                                              func(notification foundation.Notification)
	_WindowDidExitFullScreen                                               func(notification foundation.Notification)
	_CustomWindowsToEnterFullScreenForWindow                               func(window Window) []IWindow
	_CustomWindowsToEnterFullScreenForWindow_OnScreen                      func(window Window, screen Screen) []IWindow
	_Window_StartCustomAnimationToEnterFullScreenWithDuration              func(window Window, duration foundation.TimeInterval)
	_Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration     func(window Window, screen Screen, duration foundation.TimeInterval)
	_WindowDidFailToEnterFullScreen                                        func(window Window)
	_CustomWindowsToExitFullScreenForWindow                                func(window Window) []IWindow
	_Window_StartCustomAnimationToExitFullScreenWithDuration               func(window Window, duration foundation.TimeInterval)
	_WindowDidFailToExitFullScreen                                         func(window Window)
	_WindowWillMove                                                        func(notification foundation.Notification)
	_WindowDidMove                                                         func(notification foundation.Notification)
	_WindowDidChangeScreen                                                 func(notification foundation.Notification)
	_WindowDidChangeScreenProfile                                          func(notification foundation.Notification)
	_WindowDidChangeBackingProperties                                      func(notification foundation.Notification)
	_WindowShouldClose                                                     func(sender Window) bool
	_WindowWillClose                                                       func(notification foundation.Notification)
	_WindowDidBecomeKey                                                    func(notification foundation.Notification)
	_WindowDidResignKey                                                    func(notification foundation.Notification)
	_WindowDidBecomeMain                                                   func(notification foundation.Notification)
	_WindowDidResignMain                                                   func(notification foundation.Notification)
	_WindowWillReturnFieldEditor_ToObject                                  func(sender Window, client objc.Object) objc.IObject
	_WindowDidUpdate                                                       func(notification foundation.Notification)
	_WindowDidExpose                                                       func(notification foundation.Notification)
	_WindowDidChangeOcclusionState                                         func(notification foundation.Notification)
	_Window_ShouldDragDocumentWithEvent_From_WithPasteboard                func(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool
	_WindowWillReturnUndoManager                                           func(window Window) foundation.IUndoManager
	_Window_ShouldPopUpDocumentPathMenu                                    func(window Window, menu Menu) bool
	_Window_WillEncodeRestorableState                                      func(window Window, state foundation.Coder)
	_Window_DidDecodeRestorableState                                       func(window Window, state foundation.Coder)
	_Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize func(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size
	_WindowWillEnterVersionBrowser                                         func(notification foundation.Notification)
	_WindowDidEnterVersionBrowser                                          func(notification foundation.Notification)
	_WindowWillExitVersionBrowser                                          func(notification foundation.Notification)
	_WindowDidExitVersionBrowser                                           func(notification foundation.Notification)
}

func (di *WindowDelegateImpl) ImplementsWindow_WillPositionSheet_UsingRect() bool {
	return di._Window_WillPositionSheet_UsingRect != nil
}

func (di *WindowDelegateImpl) SetWindow_WillPositionSheet_UsingRect(f func(window Window, sheet Window, rect foundation.Rect) foundation.Rect) {
	di._Window_WillPositionSheet_UsingRect = f
}

func (di *WindowDelegateImpl) Window_WillPositionSheet_UsingRect(window Window, sheet Window, rect foundation.Rect) foundation.Rect {
	return di._Window_WillPositionSheet_UsingRect(window, sheet, rect)
}
func (di *WindowDelegateImpl) ImplementsWindowWillBeginSheet() bool {
	return di._WindowWillBeginSheet != nil
}

func (di *WindowDelegateImpl) SetWindowWillBeginSheet(f func(notification foundation.Notification)) {
	di._WindowWillBeginSheet = f
}

func (di *WindowDelegateImpl) WindowWillBeginSheet(notification foundation.Notification) {
	di._WindowWillBeginSheet(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidEndSheet() bool {
	return di._WindowDidEndSheet != nil
}

func (di *WindowDelegateImpl) SetWindowDidEndSheet(f func(notification foundation.Notification)) {
	di._WindowDidEndSheet = f
}

func (di *WindowDelegateImpl) WindowDidEndSheet(notification foundation.Notification) {
	di._WindowDidEndSheet(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillResize_ToSize() bool {
	return di._WindowWillResize_ToSize != nil
}

func (di *WindowDelegateImpl) SetWindowWillResize_ToSize(f func(sender Window, frameSize foundation.Size) foundation.Size) {
	di._WindowWillResize_ToSize = f
}

func (di *WindowDelegateImpl) WindowWillResize_ToSize(sender Window, frameSize foundation.Size) foundation.Size {
	return di._WindowWillResize_ToSize(sender, frameSize)
}
func (di *WindowDelegateImpl) ImplementsWindowDidResize() bool {
	return di._WindowDidResize != nil
}

func (di *WindowDelegateImpl) SetWindowDidResize(f func(notification foundation.Notification)) {
	di._WindowDidResize = f
}

func (di *WindowDelegateImpl) WindowDidResize(notification foundation.Notification) {
	di._WindowDidResize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillStartLiveResize() bool {
	return di._WindowWillStartLiveResize != nil
}

func (di *WindowDelegateImpl) SetWindowWillStartLiveResize(f func(notification foundation.Notification)) {
	di._WindowWillStartLiveResize = f
}

func (di *WindowDelegateImpl) WindowWillStartLiveResize(notification foundation.Notification) {
	di._WindowWillStartLiveResize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidEndLiveResize() bool {
	return di._WindowDidEndLiveResize != nil
}

func (di *WindowDelegateImpl) SetWindowDidEndLiveResize(f func(notification foundation.Notification)) {
	di._WindowDidEndLiveResize = f
}

func (di *WindowDelegateImpl) WindowDidEndLiveResize(notification foundation.Notification) {
	di._WindowDidEndLiveResize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillMiniaturize() bool {
	return di._WindowWillMiniaturize != nil
}

func (di *WindowDelegateImpl) SetWindowWillMiniaturize(f func(notification foundation.Notification)) {
	di._WindowWillMiniaturize = f
}

func (di *WindowDelegateImpl) WindowWillMiniaturize(notification foundation.Notification) {
	di._WindowWillMiniaturize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidMiniaturize() bool {
	return di._WindowDidMiniaturize != nil
}

func (di *WindowDelegateImpl) SetWindowDidMiniaturize(f func(notification foundation.Notification)) {
	di._WindowDidMiniaturize = f
}

func (di *WindowDelegateImpl) WindowDidMiniaturize(notification foundation.Notification) {
	di._WindowDidMiniaturize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidDeminiaturize() bool {
	return di._WindowDidDeminiaturize != nil
}

func (di *WindowDelegateImpl) SetWindowDidDeminiaturize(f func(notification foundation.Notification)) {
	di._WindowDidDeminiaturize = f
}

func (di *WindowDelegateImpl) WindowDidDeminiaturize(notification foundation.Notification) {
	di._WindowDidDeminiaturize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillUseStandardFrame_DefaultFrame() bool {
	return di._WindowWillUseStandardFrame_DefaultFrame != nil
}

func (di *WindowDelegateImpl) SetWindowWillUseStandardFrame_DefaultFrame(f func(window Window, newFrame foundation.Rect) foundation.Rect) {
	di._WindowWillUseStandardFrame_DefaultFrame = f
}

func (di *WindowDelegateImpl) WindowWillUseStandardFrame_DefaultFrame(window Window, newFrame foundation.Rect) foundation.Rect {
	return di._WindowWillUseStandardFrame_DefaultFrame(window, newFrame)
}
func (di *WindowDelegateImpl) ImplementsWindowShouldZoom_ToFrame() bool {
	return di._WindowShouldZoom_ToFrame != nil
}

func (di *WindowDelegateImpl) SetWindowShouldZoom_ToFrame(f func(window Window, newFrame foundation.Rect) bool) {
	di._WindowShouldZoom_ToFrame = f
}

func (di *WindowDelegateImpl) WindowShouldZoom_ToFrame(window Window, newFrame foundation.Rect) bool {
	return di._WindowShouldZoom_ToFrame(window, newFrame)
}
func (di *WindowDelegateImpl) ImplementsWindow_WillUseFullScreenContentSize() bool {
	return di._Window_WillUseFullScreenContentSize != nil
}

func (di *WindowDelegateImpl) SetWindow_WillUseFullScreenContentSize(f func(window Window, proposedSize foundation.Size) foundation.Size) {
	di._Window_WillUseFullScreenContentSize = f
}

func (di *WindowDelegateImpl) Window_WillUseFullScreenContentSize(window Window, proposedSize foundation.Size) foundation.Size {
	return di._Window_WillUseFullScreenContentSize(window, proposedSize)
}
func (di *WindowDelegateImpl) ImplementsWindow_WillUseFullScreenPresentationOptions() bool {
	return di._Window_WillUseFullScreenPresentationOptions != nil
}

func (di *WindowDelegateImpl) SetWindow_WillUseFullScreenPresentationOptions(f func(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions) {
	di._Window_WillUseFullScreenPresentationOptions = f
}

func (di *WindowDelegateImpl) Window_WillUseFullScreenPresentationOptions(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions {
	return di._Window_WillUseFullScreenPresentationOptions(window, proposedOptions)
}
func (di *WindowDelegateImpl) ImplementsWindowWillEnterFullScreen() bool {
	return di._WindowWillEnterFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowWillEnterFullScreen(f func(notification foundation.Notification)) {
	di._WindowWillEnterFullScreen = f
}

func (di *WindowDelegateImpl) WindowWillEnterFullScreen(notification foundation.Notification) {
	di._WindowWillEnterFullScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidEnterFullScreen() bool {
	return di._WindowDidEnterFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidEnterFullScreen(f func(notification foundation.Notification)) {
	di._WindowDidEnterFullScreen = f
}

func (di *WindowDelegateImpl) WindowDidEnterFullScreen(notification foundation.Notification) {
	di._WindowDidEnterFullScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillExitFullScreen() bool {
	return di._WindowWillExitFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowWillExitFullScreen(f func(notification foundation.Notification)) {
	di._WindowWillExitFullScreen = f
}

func (di *WindowDelegateImpl) WindowWillExitFullScreen(notification foundation.Notification) {
	di._WindowWillExitFullScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidExitFullScreen() bool {
	return di._WindowDidExitFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidExitFullScreen(f func(notification foundation.Notification)) {
	di._WindowDidExitFullScreen = f
}

func (di *WindowDelegateImpl) WindowDidExitFullScreen(notification foundation.Notification) {
	di._WindowDidExitFullScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsCustomWindowsToEnterFullScreenForWindow() bool {
	return di._CustomWindowsToEnterFullScreenForWindow != nil
}

func (di *WindowDelegateImpl) SetCustomWindowsToEnterFullScreenForWindow(f func(window Window) []IWindow) {
	di._CustomWindowsToEnterFullScreenForWindow = f
}

func (di *WindowDelegateImpl) CustomWindowsToEnterFullScreenForWindow(window Window) []IWindow {
	return di._CustomWindowsToEnterFullScreenForWindow(window)
}
func (di *WindowDelegateImpl) ImplementsCustomWindowsToEnterFullScreenForWindow_OnScreen() bool {
	return di._CustomWindowsToEnterFullScreenForWindow_OnScreen != nil
}

func (di *WindowDelegateImpl) SetCustomWindowsToEnterFullScreenForWindow_OnScreen(f func(window Window, screen Screen) []IWindow) {
	di._CustomWindowsToEnterFullScreenForWindow_OnScreen = f
}

func (di *WindowDelegateImpl) CustomWindowsToEnterFullScreenForWindow_OnScreen(window Window, screen Screen) []IWindow {
	return di._CustomWindowsToEnterFullScreenForWindow_OnScreen(window, screen)
}
func (di *WindowDelegateImpl) ImplementsWindow_StartCustomAnimationToEnterFullScreenWithDuration() bool {
	return di._Window_StartCustomAnimationToEnterFullScreenWithDuration != nil
}

func (di *WindowDelegateImpl) SetWindow_StartCustomAnimationToEnterFullScreenWithDuration(f func(window Window, duration foundation.TimeInterval)) {
	di._Window_StartCustomAnimationToEnterFullScreenWithDuration = f
}

func (di *WindowDelegateImpl) Window_StartCustomAnimationToEnterFullScreenWithDuration(window Window, duration foundation.TimeInterval) {
	di._Window_StartCustomAnimationToEnterFullScreenWithDuration(window, duration)
}
func (di *WindowDelegateImpl) ImplementsWindow_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration() bool {
	return di._Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration != nil
}

func (di *WindowDelegateImpl) SetWindow_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(f func(window Window, screen Screen, duration foundation.TimeInterval)) {
	di._Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration = f
}

func (di *WindowDelegateImpl) Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(window Window, screen Screen, duration foundation.TimeInterval) {
	di._Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(window, screen, duration)
}
func (di *WindowDelegateImpl) ImplementsWindowDidFailToEnterFullScreen() bool {
	return di._WindowDidFailToEnterFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidFailToEnterFullScreen(f func(window Window)) {
	di._WindowDidFailToEnterFullScreen = f
}

func (di *WindowDelegateImpl) WindowDidFailToEnterFullScreen(window Window) {
	di._WindowDidFailToEnterFullScreen(window)
}
func (di *WindowDelegateImpl) ImplementsCustomWindowsToExitFullScreenForWindow() bool {
	return di._CustomWindowsToExitFullScreenForWindow != nil
}

func (di *WindowDelegateImpl) SetCustomWindowsToExitFullScreenForWindow(f func(window Window) []IWindow) {
	di._CustomWindowsToExitFullScreenForWindow = f
}

func (di *WindowDelegateImpl) CustomWindowsToExitFullScreenForWindow(window Window) []IWindow {
	return di._CustomWindowsToExitFullScreenForWindow(window)
}
func (di *WindowDelegateImpl) ImplementsWindow_StartCustomAnimationToExitFullScreenWithDuration() bool {
	return di._Window_StartCustomAnimationToExitFullScreenWithDuration != nil
}

func (di *WindowDelegateImpl) SetWindow_StartCustomAnimationToExitFullScreenWithDuration(f func(window Window, duration foundation.TimeInterval)) {
	di._Window_StartCustomAnimationToExitFullScreenWithDuration = f
}

func (di *WindowDelegateImpl) Window_StartCustomAnimationToExitFullScreenWithDuration(window Window, duration foundation.TimeInterval) {
	di._Window_StartCustomAnimationToExitFullScreenWithDuration(window, duration)
}
func (di *WindowDelegateImpl) ImplementsWindowDidFailToExitFullScreen() bool {
	return di._WindowDidFailToExitFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidFailToExitFullScreen(f func(window Window)) {
	di._WindowDidFailToExitFullScreen = f
}

func (di *WindowDelegateImpl) WindowDidFailToExitFullScreen(window Window) {
	di._WindowDidFailToExitFullScreen(window)
}
func (di *WindowDelegateImpl) ImplementsWindowWillMove() bool {
	return di._WindowWillMove != nil
}

func (di *WindowDelegateImpl) SetWindowWillMove(f func(notification foundation.Notification)) {
	di._WindowWillMove = f
}

func (di *WindowDelegateImpl) WindowWillMove(notification foundation.Notification) {
	di._WindowWillMove(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidMove() bool {
	return di._WindowDidMove != nil
}

func (di *WindowDelegateImpl) SetWindowDidMove(f func(notification foundation.Notification)) {
	di._WindowDidMove = f
}

func (di *WindowDelegateImpl) WindowDidMove(notification foundation.Notification) {
	di._WindowDidMove(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidChangeScreen() bool {
	return di._WindowDidChangeScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidChangeScreen(f func(notification foundation.Notification)) {
	di._WindowDidChangeScreen = f
}

func (di *WindowDelegateImpl) WindowDidChangeScreen(notification foundation.Notification) {
	di._WindowDidChangeScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidChangeScreenProfile() bool {
	return di._WindowDidChangeScreenProfile != nil
}

func (di *WindowDelegateImpl) SetWindowDidChangeScreenProfile(f func(notification foundation.Notification)) {
	di._WindowDidChangeScreenProfile = f
}

func (di *WindowDelegateImpl) WindowDidChangeScreenProfile(notification foundation.Notification) {
	di._WindowDidChangeScreenProfile(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidChangeBackingProperties() bool {
	return di._WindowDidChangeBackingProperties != nil
}

func (di *WindowDelegateImpl) SetWindowDidChangeBackingProperties(f func(notification foundation.Notification)) {
	di._WindowDidChangeBackingProperties = f
}

func (di *WindowDelegateImpl) WindowDidChangeBackingProperties(notification foundation.Notification) {
	di._WindowDidChangeBackingProperties(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowShouldClose() bool {
	return di._WindowShouldClose != nil
}

func (di *WindowDelegateImpl) SetWindowShouldClose(f func(sender Window) bool) {
	di._WindowShouldClose = f
}

func (di *WindowDelegateImpl) WindowShouldClose(sender Window) bool {
	return di._WindowShouldClose(sender)
}
func (di *WindowDelegateImpl) ImplementsWindowWillClose() bool {
	return di._WindowWillClose != nil
}

func (di *WindowDelegateImpl) SetWindowWillClose(f func(notification foundation.Notification)) {
	di._WindowWillClose = f
}

func (di *WindowDelegateImpl) WindowWillClose(notification foundation.Notification) {
	di._WindowWillClose(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidBecomeKey() bool {
	return di._WindowDidBecomeKey != nil
}

func (di *WindowDelegateImpl) SetWindowDidBecomeKey(f func(notification foundation.Notification)) {
	di._WindowDidBecomeKey = f
}

func (di *WindowDelegateImpl) WindowDidBecomeKey(notification foundation.Notification) {
	di._WindowDidBecomeKey(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidResignKey() bool {
	return di._WindowDidResignKey != nil
}

func (di *WindowDelegateImpl) SetWindowDidResignKey(f func(notification foundation.Notification)) {
	di._WindowDidResignKey = f
}

func (di *WindowDelegateImpl) WindowDidResignKey(notification foundation.Notification) {
	di._WindowDidResignKey(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidBecomeMain() bool {
	return di._WindowDidBecomeMain != nil
}

func (di *WindowDelegateImpl) SetWindowDidBecomeMain(f func(notification foundation.Notification)) {
	di._WindowDidBecomeMain = f
}

func (di *WindowDelegateImpl) WindowDidBecomeMain(notification foundation.Notification) {
	di._WindowDidBecomeMain(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidResignMain() bool {
	return di._WindowDidResignMain != nil
}

func (di *WindowDelegateImpl) SetWindowDidResignMain(f func(notification foundation.Notification)) {
	di._WindowDidResignMain = f
}

func (di *WindowDelegateImpl) WindowDidResignMain(notification foundation.Notification) {
	di._WindowDidResignMain(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillReturnFieldEditor_ToObject() bool {
	return di._WindowWillReturnFieldEditor_ToObject != nil
}

func (di *WindowDelegateImpl) SetWindowWillReturnFieldEditor_ToObject(f func(sender Window, client objc.Object) objc.IObject) {
	di._WindowWillReturnFieldEditor_ToObject = f
}

func (di *WindowDelegateImpl) WindowWillReturnFieldEditor_ToObject(sender Window, client objc.Object) objc.IObject {
	return di._WindowWillReturnFieldEditor_ToObject(sender, client)
}
func (di *WindowDelegateImpl) ImplementsWindowDidUpdate() bool {
	return di._WindowDidUpdate != nil
}

func (di *WindowDelegateImpl) SetWindowDidUpdate(f func(notification foundation.Notification)) {
	di._WindowDidUpdate = f
}

func (di *WindowDelegateImpl) WindowDidUpdate(notification foundation.Notification) {
	di._WindowDidUpdate(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidExpose() bool {
	return di._WindowDidExpose != nil
}

func (di *WindowDelegateImpl) SetWindowDidExpose(f func(notification foundation.Notification)) {
	di._WindowDidExpose = f
}

func (di *WindowDelegateImpl) WindowDidExpose(notification foundation.Notification) {
	di._WindowDidExpose(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidChangeOcclusionState() bool {
	return di._WindowDidChangeOcclusionState != nil
}

func (di *WindowDelegateImpl) SetWindowDidChangeOcclusionState(f func(notification foundation.Notification)) {
	di._WindowDidChangeOcclusionState = f
}

func (di *WindowDelegateImpl) WindowDidChangeOcclusionState(notification foundation.Notification) {
	di._WindowDidChangeOcclusionState(notification)
}
func (di *WindowDelegateImpl) ImplementsWindow_ShouldDragDocumentWithEvent_From_WithPasteboard() bool {
	return di._Window_ShouldDragDocumentWithEvent_From_WithPasteboard != nil
}

func (di *WindowDelegateImpl) SetWindow_ShouldDragDocumentWithEvent_From_WithPasteboard(f func(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool) {
	di._Window_ShouldDragDocumentWithEvent_From_WithPasteboard = f
}

func (di *WindowDelegateImpl) Window_ShouldDragDocumentWithEvent_From_WithPasteboard(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool {
	return di._Window_ShouldDragDocumentWithEvent_From_WithPasteboard(window, event, dragImageLocation, pasteboard)
}
func (di *WindowDelegateImpl) ImplementsWindowWillReturnUndoManager() bool {
	return di._WindowWillReturnUndoManager != nil
}

func (di *WindowDelegateImpl) SetWindowWillReturnUndoManager(f func(window Window) foundation.IUndoManager) {
	di._WindowWillReturnUndoManager = f
}

func (di *WindowDelegateImpl) WindowWillReturnUndoManager(window Window) foundation.IUndoManager {
	return di._WindowWillReturnUndoManager(window)
}
func (di *WindowDelegateImpl) ImplementsWindow_ShouldPopUpDocumentPathMenu() bool {
	return di._Window_ShouldPopUpDocumentPathMenu != nil
}

func (di *WindowDelegateImpl) SetWindow_ShouldPopUpDocumentPathMenu(f func(window Window, menu Menu) bool) {
	di._Window_ShouldPopUpDocumentPathMenu = f
}

func (di *WindowDelegateImpl) Window_ShouldPopUpDocumentPathMenu(window Window, menu Menu) bool {
	return di._Window_ShouldPopUpDocumentPathMenu(window, menu)
}
func (di *WindowDelegateImpl) ImplementsWindow_WillEncodeRestorableState() bool {
	return di._Window_WillEncodeRestorableState != nil
}

func (di *WindowDelegateImpl) SetWindow_WillEncodeRestorableState(f func(window Window, state foundation.Coder)) {
	di._Window_WillEncodeRestorableState = f
}

func (di *WindowDelegateImpl) Window_WillEncodeRestorableState(window Window, state foundation.Coder) {
	di._Window_WillEncodeRestorableState(window, state)
}
func (di *WindowDelegateImpl) ImplementsWindow_DidDecodeRestorableState() bool {
	return di._Window_DidDecodeRestorableState != nil
}

func (di *WindowDelegateImpl) SetWindow_DidDecodeRestorableState(f func(window Window, state foundation.Coder)) {
	di._Window_DidDecodeRestorableState = f
}

func (di *WindowDelegateImpl) Window_DidDecodeRestorableState(window Window, state foundation.Coder) {
	di._Window_DidDecodeRestorableState(window, state)
}
func (di *WindowDelegateImpl) ImplementsWindow_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize() bool {
	return di._Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize != nil
}

func (di *WindowDelegateImpl) SetWindow_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(f func(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size) {
	di._Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize = f
}

func (di *WindowDelegateImpl) Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size {
	return di._Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(window, maxPreferredFrameSize, maxAllowedFrameSize)
}
func (di *WindowDelegateImpl) ImplementsWindowWillEnterVersionBrowser() bool {
	return di._WindowWillEnterVersionBrowser != nil
}

func (di *WindowDelegateImpl) SetWindowWillEnterVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowWillEnterVersionBrowser = f
}

func (di *WindowDelegateImpl) WindowWillEnterVersionBrowser(notification foundation.Notification) {
	di._WindowWillEnterVersionBrowser(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidEnterVersionBrowser() bool {
	return di._WindowDidEnterVersionBrowser != nil
}

func (di *WindowDelegateImpl) SetWindowDidEnterVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowDidEnterVersionBrowser = f
}

func (di *WindowDelegateImpl) WindowDidEnterVersionBrowser(notification foundation.Notification) {
	di._WindowDidEnterVersionBrowser(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillExitVersionBrowser() bool {
	return di._WindowWillExitVersionBrowser != nil
}

func (di *WindowDelegateImpl) SetWindowWillExitVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowWillExitVersionBrowser = f
}

func (di *WindowDelegateImpl) WindowWillExitVersionBrowser(notification foundation.Notification) {
	di._WindowWillExitVersionBrowser(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidExitVersionBrowser() bool {
	return di._WindowDidExitVersionBrowser != nil
}

func (di *WindowDelegateImpl) SetWindowDidExitVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowDidExitVersionBrowser = f
}

func (di *WindowDelegateImpl) WindowDidExitVersionBrowser(notification foundation.Notification) {
	di._WindowDidExitVersionBrowser(notification)
}

type WindowDelegateWrapper struct {
	objc.Object
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_WillPositionSheet_UsingRect() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willPositionSheet:usingRect:"))
}

func (w_ WindowDelegateWrapper) Window_WillPositionSheet_UsingRect(window IWindow, sheet IWindow, rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "window:willPositionSheet:usingRect:", window, sheet, rect)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillBeginSheet() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillBeginSheet:"))
}

func (w_ WindowDelegateWrapper) WindowWillBeginSheet(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillBeginSheet:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidEndSheet() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEndSheet:"))
}

func (w_ WindowDelegateWrapper) WindowDidEndSheet(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidEndSheet:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillResize_ToSize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillResize:toSize:"))
}

func (w_ WindowDelegateWrapper) WindowWillResize_ToSize(sender IWindow, frameSize foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "windowWillResize:toSize:", sender, frameSize)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidResize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidResize:"))
}

func (w_ WindowDelegateWrapper) WindowDidResize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidResize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillStartLiveResize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillStartLiveResize:"))
}

func (w_ WindowDelegateWrapper) WindowWillStartLiveResize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillStartLiveResize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidEndLiveResize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEndLiveResize:"))
}

func (w_ WindowDelegateWrapper) WindowDidEndLiveResize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidEndLiveResize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillMiniaturize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillMiniaturize:"))
}

func (w_ WindowDelegateWrapper) WindowWillMiniaturize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillMiniaturize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidMiniaturize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidMiniaturize:"))
}

func (w_ WindowDelegateWrapper) WindowDidMiniaturize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidMiniaturize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidDeminiaturize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidDeminiaturize:"))
}

func (w_ WindowDelegateWrapper) WindowDidDeminiaturize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidDeminiaturize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillUseStandardFrame_DefaultFrame() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillUseStandardFrame:defaultFrame:"))
}

func (w_ WindowDelegateWrapper) WindowWillUseStandardFrame_DefaultFrame(window IWindow, newFrame foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "windowWillUseStandardFrame:defaultFrame:", window, newFrame)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowShouldZoom_ToFrame() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowShouldZoom:toFrame:"))
}

func (w_ WindowDelegateWrapper) WindowShouldZoom_ToFrame(window IWindow, newFrame foundation.Rect) bool {
	rv := ffi.CallMethod[bool](w_, "windowShouldZoom:toFrame:", window, newFrame)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_WillUseFullScreenContentSize() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willUseFullScreenContentSize:"))
}

func (w_ WindowDelegateWrapper) Window_WillUseFullScreenContentSize(window IWindow, proposedSize foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "window:willUseFullScreenContentSize:", window, proposedSize)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_WillUseFullScreenPresentationOptions() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willUseFullScreenPresentationOptions:"))
}

func (w_ WindowDelegateWrapper) Window_WillUseFullScreenPresentationOptions(window IWindow, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions {
	rv := ffi.CallMethod[ApplicationPresentationOptions](w_, "window:willUseFullScreenPresentationOptions:", window, proposedOptions)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillEnterFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowWillEnterFullScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillEnterFullScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEnterFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidEnterFullScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidEnterFullScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillExitFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillExitFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowWillExitFullScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillExitFullScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidExitFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidExitFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidExitFullScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidExitFullScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsCustomWindowsToEnterFullScreenForWindow() bool {
	return w_.RespondsToSelector(objc.GetSelector("customWindowsToEnterFullScreenForWindow:"))
}

func (w_ WindowDelegateWrapper) CustomWindowsToEnterFullScreenForWindow(window IWindow) []Window {
	rv := ffi.CallMethod[[]Window](w_, "customWindowsToEnterFullScreenForWindow:", window)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsCustomWindowsToEnterFullScreenForWindow_OnScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("customWindowsToEnterFullScreenForWindow:onScreen:"))
}

func (w_ WindowDelegateWrapper) CustomWindowsToEnterFullScreenForWindow_OnScreen(window IWindow, screen IScreen) []Window {
	rv := ffi.CallMethod[[]Window](w_, "customWindowsToEnterFullScreenForWindow:onScreen:", window, screen)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_StartCustomAnimationToEnterFullScreenWithDuration() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:startCustomAnimationToEnterFullScreenWithDuration:"))
}

func (w_ WindowDelegateWrapper) Window_StartCustomAnimationToEnterFullScreenWithDuration(window IWindow, duration foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](w_, "window:startCustomAnimationToEnterFullScreenWithDuration:", window, duration)
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:"))
}

func (w_ WindowDelegateWrapper) Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(window IWindow, screen IScreen, duration foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](w_, "window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:", window, screen, duration)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidFailToEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidFailToEnterFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidFailToEnterFullScreen(window IWindow) {
	ffi.CallMethod[ffi.Void](w_, "windowDidFailToEnterFullScreen:", window)
}

func (w_ *WindowDelegateWrapper) ImplementsCustomWindowsToExitFullScreenForWindow() bool {
	return w_.RespondsToSelector(objc.GetSelector("customWindowsToExitFullScreenForWindow:"))
}

func (w_ WindowDelegateWrapper) CustomWindowsToExitFullScreenForWindow(window IWindow) []Window {
	rv := ffi.CallMethod[[]Window](w_, "customWindowsToExitFullScreenForWindow:", window)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_StartCustomAnimationToExitFullScreenWithDuration() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:startCustomAnimationToExitFullScreenWithDuration:"))
}

func (w_ WindowDelegateWrapper) Window_StartCustomAnimationToExitFullScreenWithDuration(window IWindow, duration foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](w_, "window:startCustomAnimationToExitFullScreenWithDuration:", window, duration)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidFailToExitFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidFailToExitFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidFailToExitFullScreen(window IWindow) {
	ffi.CallMethod[ffi.Void](w_, "windowDidFailToExitFullScreen:", window)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillMove() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillMove:"))
}

func (w_ WindowDelegateWrapper) WindowWillMove(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillMove:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidMove() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidMove:"))
}

func (w_ WindowDelegateWrapper) WindowDidMove(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidMove:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidChangeScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidChangeScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidChangeScreenProfile() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeScreenProfile:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeScreenProfile(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidChangeScreenProfile:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidChangeBackingProperties() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeBackingProperties:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeBackingProperties(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidChangeBackingProperties:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowShouldClose() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowShouldClose:"))
}

func (w_ WindowDelegateWrapper) WindowShouldClose(sender IWindow) bool {
	rv := ffi.CallMethod[bool](w_, "windowShouldClose:", sender)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillClose() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillClose:"))
}

func (w_ WindowDelegateWrapper) WindowWillClose(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillClose:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidBecomeKey() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidBecomeKey:"))
}

func (w_ WindowDelegateWrapper) WindowDidBecomeKey(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidBecomeKey:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidResignKey() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidResignKey:"))
}

func (w_ WindowDelegateWrapper) WindowDidResignKey(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidResignKey:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidBecomeMain() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidBecomeMain:"))
}

func (w_ WindowDelegateWrapper) WindowDidBecomeMain(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidBecomeMain:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidResignMain() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidResignMain:"))
}

func (w_ WindowDelegateWrapper) WindowDidResignMain(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidResignMain:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillReturnFieldEditor_ToObject() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillReturnFieldEditor:toObject:"))
}

func (w_ WindowDelegateWrapper) WindowWillReturnFieldEditor_ToObject(sender IWindow, client objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](w_, "windowWillReturnFieldEditor:toObject:", sender, client)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidUpdate() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidUpdate:"))
}

func (w_ WindowDelegateWrapper) WindowDidUpdate(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidUpdate:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidExpose() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidExpose:"))
}

func (w_ WindowDelegateWrapper) WindowDidExpose(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidExpose:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidChangeOcclusionState() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeOcclusionState:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeOcclusionState(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidChangeOcclusionState:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_ShouldDragDocumentWithEvent_From_WithPasteboard() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:shouldDragDocumentWithEvent:from:withPasteboard:"))
}

func (w_ WindowDelegateWrapper) Window_ShouldDragDocumentWithEvent_From_WithPasteboard(window IWindow, event IEvent, dragImageLocation foundation.Point, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](w_, "window:shouldDragDocumentWithEvent:from:withPasteboard:", window, event, dragImageLocation, pasteboard)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillReturnUndoManager() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillReturnUndoManager:"))
}

func (w_ WindowDelegateWrapper) WindowWillReturnUndoManager(window IWindow) foundation.UndoManager {
	rv := ffi.CallMethod[foundation.UndoManager](w_, "windowWillReturnUndoManager:", window)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_ShouldPopUpDocumentPathMenu() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:shouldPopUpDocumentPathMenu:"))
}

func (w_ WindowDelegateWrapper) Window_ShouldPopUpDocumentPathMenu(window IWindow, menu IMenu) bool {
	rv := ffi.CallMethod[bool](w_, "window:shouldPopUpDocumentPathMenu:", window, menu)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_WillEncodeRestorableState() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willEncodeRestorableState:"))
}

func (w_ WindowDelegateWrapper) Window_WillEncodeRestorableState(window IWindow, state foundation.ICoder) {
	ffi.CallMethod[ffi.Void](w_, "window:willEncodeRestorableState:", window, state)
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_DidDecodeRestorableState() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:didDecodeRestorableState:"))
}

func (w_ WindowDelegateWrapper) Window_DidDecodeRestorableState(window IWindow, state foundation.ICoder) {
	ffi.CallMethod[ffi.Void](w_, "window:didDecodeRestorableState:", window, state)
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:"))
}

func (w_ WindowDelegateWrapper) Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(window IWindow, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "window:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:", window, maxPreferredFrameSize, maxAllowedFrameSize)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillEnterVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillEnterVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowWillEnterVersionBrowser(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillEnterVersionBrowser:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidEnterVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEnterVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowDidEnterVersionBrowser(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidEnterVersionBrowser:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillExitVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillExitVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowWillExitVersionBrowser(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillExitVersionBrowser:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidExitVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidExitVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowDidExitVersionBrowser(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidExitVersionBrowser:", notification)
}
