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

func WrapOpenSavePanelDelegate(v OpenSavePanelDelegate) objc.Object {
	return objc.WrapAsProtocol("NSOpenSavePanelDelegate", v)
}

type OpenSavePanelDelegateBase struct {
}

func (p *OpenSavePanelDelegateBase) ImplementsPanel_UserEnteredFilename_Confirmed() bool {
	return false
}

func (p *OpenSavePanelDelegateBase) Panel_UserEnteredFilename_Confirmed(sender objc.Object, filename string, okFlag bool) string {
	panic("unimpemented protocol method")
}

func (p *OpenSavePanelDelegateBase) ImplementsPanelSelectionDidChange() bool {
	return false
}

func (p *OpenSavePanelDelegateBase) PanelSelectionDidChange(sender objc.Object) {
	panic("unimpemented protocol method")
}

func (p *OpenSavePanelDelegateBase) ImplementsPanel_DidChangeToDirectoryURL() bool {
	return false
}

func (p *OpenSavePanelDelegateBase) Panel_DidChangeToDirectoryURL(sender objc.Object, url foundation.URL) {
	panic("unimpemented protocol method")
}

func (p *OpenSavePanelDelegateBase) ImplementsPanel_WillExpand() bool {
	return false
}

func (p *OpenSavePanelDelegateBase) Panel_WillExpand(sender objc.Object, expanding bool) {
	panic("unimpemented protocol method")
}

func (p *OpenSavePanelDelegateBase) ImplementsPanel_ShouldEnableURL() bool {
	return false
}

func (p *OpenSavePanelDelegateBase) Panel_ShouldEnableURL(sender objc.Object, url foundation.URL) bool {
	panic("unimpemented protocol method")
}

func (p *OpenSavePanelDelegateBase) ImplementsPanel_ValidateURL_Error() bool {
	return false
}

func (p *OpenSavePanelDelegateBase) Panel_ValidateURL_Error(sender objc.Object, url foundation.URL, outError *foundation.Error) bool {
	panic("unimpemented protocol method")
}

type OpenSavePanelDelegateWrapper struct {
	objc.Object
}

func (o_ OpenSavePanelDelegateWrapper) Panel_UserEnteredFilename_Confirmed(sender objc.IObject, filename string, okFlag bool) string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("panel:userEnteredFilename:confirmed:"), objc.ExtractPtr(sender), filename, okFlag)
	return rv
}

func (o_ OpenSavePanelDelegateWrapper) PanelSelectionDidChange(sender objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("panelSelectionDidChange:"), objc.ExtractPtr(sender))
}

func (o_ OpenSavePanelDelegateWrapper) Panel_DidChangeToDirectoryURL(sender objc.IObject, url foundation.IURL) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("panel:didChangeToDirectoryURL:"), objc.ExtractPtr(sender), objc.ExtractPtr(url))
}

func (o_ OpenSavePanelDelegateWrapper) Panel_WillExpand(sender objc.IObject, expanding bool) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("panel:willExpand:"), objc.ExtractPtr(sender), expanding)
}

func (o_ OpenSavePanelDelegateWrapper) Panel_ShouldEnableURL(sender objc.IObject, url foundation.IURL) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("panel:shouldEnableURL:"), objc.ExtractPtr(sender), objc.ExtractPtr(url))
	return rv
}

func (o_ OpenSavePanelDelegateWrapper) Panel_ValidateURL_Error(sender objc.IObject, url foundation.IURL, outError *foundation.Error) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("panel:validateURL:error:"), objc.ExtractPtr(sender), objc.ExtractPtr(url), unsafe.Pointer(outError))
	return rv
}
