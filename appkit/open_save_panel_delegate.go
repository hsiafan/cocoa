// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type OpenSavePanelDelegate interface {
	ImplementsPanel_UserEnteredFilename_Confirmed() bool
	// optional
	Panel_UserEnteredFilename_Confirmed(sender objc.Object, filename string, okFlag bool) string
	ImplementsPanelSelectionDidChange() bool
	// optional
	PanelSelectionDidChange(sender objc.Object)
	ImplementsPanel_DidChangeToDirectoryURL() bool
	// optional
	Panel_DidChangeToDirectoryURL(sender objc.Object, url foundation.URL)
	ImplementsPanel_WillExpand() bool
	// optional
	Panel_WillExpand(sender objc.Object, expanding bool)
	ImplementsPanel_ShouldEnableURL() bool
	// optional
	Panel_ShouldEnableURL(sender objc.Object, url foundation.URL) bool
	ImplementsPanel_ValidateURL_Error() bool
	// optional
	Panel_ValidateURL_Error(sender objc.Object, url foundation.URL, outError *foundation.Error) bool
}

type OpenSavePanelDelegateImpl struct {
	_Panel_UserEnteredFilename_Confirmed func(sender objc.Object, filename string, okFlag bool) string
	_PanelSelectionDidChange             func(sender objc.Object)
	_Panel_DidChangeToDirectoryURL       func(sender objc.Object, url foundation.URL)
	_Panel_WillExpand                    func(sender objc.Object, expanding bool)
	_Panel_ShouldEnableURL               func(sender objc.Object, url foundation.URL) bool
	_Panel_ValidateURL_Error             func(sender objc.Object, url foundation.URL, outError *foundation.Error) bool
}

func (di *OpenSavePanelDelegateImpl) ImplementsPanel_UserEnteredFilename_Confirmed() bool {
	return di._Panel_UserEnteredFilename_Confirmed != nil
}

func (di *OpenSavePanelDelegateImpl) SetPanel_UserEnteredFilename_Confirmed(f func(sender objc.Object, filename string, okFlag bool) string) {
	di._Panel_UserEnteredFilename_Confirmed = f
}

func (di *OpenSavePanelDelegateImpl) Panel_UserEnteredFilename_Confirmed(sender objc.Object, filename string, okFlag bool) string {
	return di._Panel_UserEnteredFilename_Confirmed(sender, filename, okFlag)
}
func (di *OpenSavePanelDelegateImpl) ImplementsPanelSelectionDidChange() bool {
	return di._PanelSelectionDidChange != nil
}

func (di *OpenSavePanelDelegateImpl) SetPanelSelectionDidChange(f func(sender objc.Object)) {
	di._PanelSelectionDidChange = f
}

func (di *OpenSavePanelDelegateImpl) PanelSelectionDidChange(sender objc.Object) {
	di._PanelSelectionDidChange(sender)
}
func (di *OpenSavePanelDelegateImpl) ImplementsPanel_DidChangeToDirectoryURL() bool {
	return di._Panel_DidChangeToDirectoryURL != nil
}

func (di *OpenSavePanelDelegateImpl) SetPanel_DidChangeToDirectoryURL(f func(sender objc.Object, url foundation.URL)) {
	di._Panel_DidChangeToDirectoryURL = f
}

func (di *OpenSavePanelDelegateImpl) Panel_DidChangeToDirectoryURL(sender objc.Object, url foundation.URL) {
	di._Panel_DidChangeToDirectoryURL(sender, url)
}
func (di *OpenSavePanelDelegateImpl) ImplementsPanel_WillExpand() bool {
	return di._Panel_WillExpand != nil
}

func (di *OpenSavePanelDelegateImpl) SetPanel_WillExpand(f func(sender objc.Object, expanding bool)) {
	di._Panel_WillExpand = f
}

func (di *OpenSavePanelDelegateImpl) Panel_WillExpand(sender objc.Object, expanding bool) {
	di._Panel_WillExpand(sender, expanding)
}
func (di *OpenSavePanelDelegateImpl) ImplementsPanel_ShouldEnableURL() bool {
	return di._Panel_ShouldEnableURL != nil
}

func (di *OpenSavePanelDelegateImpl) SetPanel_ShouldEnableURL(f func(sender objc.Object, url foundation.URL) bool) {
	di._Panel_ShouldEnableURL = f
}

func (di *OpenSavePanelDelegateImpl) Panel_ShouldEnableURL(sender objc.Object, url foundation.URL) bool {
	return di._Panel_ShouldEnableURL(sender, url)
}
func (di *OpenSavePanelDelegateImpl) ImplementsPanel_ValidateURL_Error() bool {
	return di._Panel_ValidateURL_Error != nil
}

func (di *OpenSavePanelDelegateImpl) SetPanel_ValidateURL_Error(f func(sender objc.Object, url foundation.URL, outError *foundation.Error) bool) {
	di._Panel_ValidateURL_Error = f
}

func (di *OpenSavePanelDelegateImpl) Panel_ValidateURL_Error(sender objc.Object, url foundation.URL, outError *foundation.Error) bool {
	return di._Panel_ValidateURL_Error(sender, url, outError)
}

type OpenSavePanelDelegateWrapper struct {
	objc.Object
}

func (o_ *OpenSavePanelDelegateWrapper) ImplementsPanel_UserEnteredFilename_Confirmed() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:userEnteredFilename:confirmed:"))
}

func (o_ OpenSavePanelDelegateWrapper) Panel_UserEnteredFilename_Confirmed(sender objc.IObject, filename string, okFlag bool) string {
	rv := objc.CallMethod[string](o_, "panel:userEnteredFilename:confirmed:", sender, filename, okFlag)
	return rv
}

func (o_ *OpenSavePanelDelegateWrapper) ImplementsPanelSelectionDidChange() bool {
	return o_.RespondsToSelector(objc.GetSelector("panelSelectionDidChange:"))
}

func (o_ OpenSavePanelDelegateWrapper) PanelSelectionDidChange(sender objc.IObject) {
	objc.CallMethod[objc.Void](o_, "panelSelectionDidChange:", sender)
}

func (o_ *OpenSavePanelDelegateWrapper) ImplementsPanel_DidChangeToDirectoryURL() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:didChangeToDirectoryURL:"))
}

func (o_ OpenSavePanelDelegateWrapper) Panel_DidChangeToDirectoryURL(sender objc.IObject, url foundation.IURL) {
	objc.CallMethod[objc.Void](o_, "panel:didChangeToDirectoryURL:", sender, url)
}

func (o_ *OpenSavePanelDelegateWrapper) ImplementsPanel_WillExpand() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:willExpand:"))
}

func (o_ OpenSavePanelDelegateWrapper) Panel_WillExpand(sender objc.IObject, expanding bool) {
	objc.CallMethod[objc.Void](o_, "panel:willExpand:", sender, expanding)
}

func (o_ *OpenSavePanelDelegateWrapper) ImplementsPanel_ShouldEnableURL() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:shouldEnableURL:"))
}

func (o_ OpenSavePanelDelegateWrapper) Panel_ShouldEnableURL(sender objc.IObject, url foundation.IURL) bool {
	rv := objc.CallMethod[bool](o_, "panel:shouldEnableURL:", sender, url)
	return rv
}

func (o_ *OpenSavePanelDelegateWrapper) ImplementsPanel_ValidateURL_Error() bool {
	return o_.RespondsToSelector(objc.GetSelector("panel:validateURL:error:"))
}

func (o_ OpenSavePanelDelegateWrapper) Panel_ValidateURL_Error(sender objc.IObject, url foundation.IURL, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](o_, "panel:validateURL:error:", sender, url, unsafe.Pointer(outError))
	return rv
}
