// auto-generated code, do not modify
package appkit

import (
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

func WrapWindowDelegate(v WindowDelegate) objc.Object {
	return objc.WrapAsProtocol("NSWindowDelegate", v)
}

type WindowDelegateBase struct {
}

func (p *WindowDelegateBase) ImplementsWindow_WillPositionSheet_UsingRect() bool {
	return false
}

func (p *WindowDelegateBase) Window_WillPositionSheet_UsingRect(window Window, sheet Window, rect foundation.Rect) foundation.Rect {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillBeginSheet() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillBeginSheet(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidEndSheet() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidEndSheet(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillResize_ToSize() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillResize_ToSize(sender Window, frameSize foundation.Size) foundation.Size {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidResize() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidResize(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillStartLiveResize() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillStartLiveResize(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidEndLiveResize() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidEndLiveResize(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillMiniaturize() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillMiniaturize(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidMiniaturize() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidMiniaturize(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidDeminiaturize() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidDeminiaturize(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillUseStandardFrame_DefaultFrame() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillUseStandardFrame_DefaultFrame(window Window, newFrame foundation.Rect) foundation.Rect {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowShouldZoom_ToFrame() bool {
	return false
}

func (p *WindowDelegateBase) WindowShouldZoom_ToFrame(window Window, newFrame foundation.Rect) bool {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_WillUseFullScreenContentSize() bool {
	return false
}

func (p *WindowDelegateBase) Window_WillUseFullScreenContentSize(window Window, proposedSize foundation.Size) foundation.Size {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_WillUseFullScreenPresentationOptions() bool {
	return false
}

func (p *WindowDelegateBase) Window_WillUseFullScreenPresentationOptions(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillEnterFullScreen() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillEnterFullScreen(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidEnterFullScreen() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidEnterFullScreen(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillExitFullScreen() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillExitFullScreen(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidExitFullScreen() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidExitFullScreen(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsCustomWindowsToEnterFullScreenForWindow() bool {
	return false
}

func (p *WindowDelegateBase) CustomWindowsToEnterFullScreenForWindow(window Window) []IWindow {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsCustomWindowsToEnterFullScreenForWindow_OnScreen() bool {
	return false
}

func (p *WindowDelegateBase) CustomWindowsToEnterFullScreenForWindow_OnScreen(window Window, screen Screen) []IWindow {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_StartCustomAnimationToEnterFullScreenWithDuration() bool {
	return false
}

func (p *WindowDelegateBase) Window_StartCustomAnimationToEnterFullScreenWithDuration(window Window, duration foundation.TimeInterval) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration() bool {
	return false
}

func (p *WindowDelegateBase) Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(window Window, screen Screen, duration foundation.TimeInterval) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidFailToEnterFullScreen() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidFailToEnterFullScreen(window Window) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsCustomWindowsToExitFullScreenForWindow() bool {
	return false
}

func (p *WindowDelegateBase) CustomWindowsToExitFullScreenForWindow(window Window) []IWindow {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_StartCustomAnimationToExitFullScreenWithDuration() bool {
	return false
}

func (p *WindowDelegateBase) Window_StartCustomAnimationToExitFullScreenWithDuration(window Window, duration foundation.TimeInterval) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidFailToExitFullScreen() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidFailToExitFullScreen(window Window) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillMove() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillMove(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidMove() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidMove(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidChangeScreen() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidChangeScreen(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidChangeScreenProfile() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidChangeScreenProfile(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidChangeBackingProperties() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidChangeBackingProperties(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowShouldClose() bool {
	return false
}

func (p *WindowDelegateBase) WindowShouldClose(sender Window) bool {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillClose() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillClose(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidBecomeKey() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidBecomeKey(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidResignKey() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidResignKey(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidBecomeMain() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidBecomeMain(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidResignMain() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidResignMain(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillReturnFieldEditor_ToObject() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillReturnFieldEditor_ToObject(sender Window, client objc.Object) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidUpdate() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidUpdate(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidExpose() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidExpose(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidChangeOcclusionState() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidChangeOcclusionState(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_ShouldDragDocumentWithEvent_From_WithPasteboard() bool {
	return false
}

func (p *WindowDelegateBase) Window_ShouldDragDocumentWithEvent_From_WithPasteboard(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillReturnUndoManager() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillReturnUndoManager(window Window) foundation.IUndoManager {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_ShouldPopUpDocumentPathMenu() bool {
	return false
}

func (p *WindowDelegateBase) Window_ShouldPopUpDocumentPathMenu(window Window, menu Menu) bool {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_WillEncodeRestorableState() bool {
	return false
}

func (p *WindowDelegateBase) Window_WillEncodeRestorableState(window Window, state foundation.Coder) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_DidDecodeRestorableState() bool {
	return false
}

func (p *WindowDelegateBase) Window_DidDecodeRestorableState(window Window, state foundation.Coder) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindow_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize() bool {
	return false
}

func (p *WindowDelegateBase) Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillEnterVersionBrowser() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillEnterVersionBrowser(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidEnterVersionBrowser() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidEnterVersionBrowser(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowWillExitVersionBrowser() bool {
	return false
}

func (p *WindowDelegateBase) WindowWillExitVersionBrowser(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *WindowDelegateBase) ImplementsWindowDidExitVersionBrowser() bool {
	return false
}

func (p *WindowDelegateBase) WindowDidExitVersionBrowser(notification foundation.Notification) {
	panic("unimpemented protocol method")
}
