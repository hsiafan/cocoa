// auto-generated code, do not modify
package appkit

import (
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
