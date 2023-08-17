// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type PasteboardItemDataProvider interface {
	// required
	Pasteboard_Item_ProvideDataForType(pasteboard Pasteboard, item PasteboardItem, type_ PasteboardType)
	ImplementsPasteboardFinishedWithDataProvider() bool
	// optional
	PasteboardFinishedWithDataProvider(pasteboard Pasteboard)
}

func WrapPasteboardItemDataProvider(v PasteboardItemDataProvider) objc.Object {
	return objc.WrapAsProtocol("NSPasteboardItemDataProvider", v)
}

type PasteboardItemDataProviderBase struct {
}

func (p *PasteboardItemDataProviderBase) ImplementsPasteboardFinishedWithDataProvider() bool {
	return false
}

func (p *PasteboardItemDataProviderBase) PasteboardFinishedWithDataProvider(pasteboard Pasteboard) {
	panic("unimpemented protocol method")
}

type PasteboardItemDataProviderWrapper struct {
	objc.Object
}

func (p_ PasteboardItemDataProviderWrapper) Pasteboard_Item_ProvideDataForType(pasteboard IPasteboard, item IPasteboardItem, type_ PasteboardType) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pasteboard:item:provideDataForType:"), objc.ExtractPtr(pasteboard), objc.ExtractPtr(item), type_)
}

func (p_ PasteboardItemDataProviderWrapper) PasteboardFinishedWithDataProvider(pasteboard IPasteboard) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("pasteboardFinishedWithDataProvider:"), objc.ExtractPtr(pasteboard))
}
