// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type PasteboardItemDataProvider interface {
	// required
	Pasteboard_Item_ProvideDataForType(pasteboard Pasteboard, item PasteboardItem, type_ PasteboardType)
	ImplementsPasteboardFinishedWithDataProvider() bool
	// optional
	PasteboardFinishedWithDataProvider(pasteboard Pasteboard)
}

type PasteboardItemDataProviderWrapper struct {
	objc.Object
}

func (p_ PasteboardItemDataProviderWrapper) Pasteboard_Item_ProvideDataForType(pasteboard IPasteboard, item IPasteboardItem, type_ PasteboardType) {
	ffi.CallMethod[ffi.Void](p_, "pasteboard:item:provideDataForType:", pasteboard, item, type_)
}

func (p_ *PasteboardItemDataProviderWrapper) ImplementsPasteboardFinishedWithDataProvider() bool {
	return p_.RespondsToSelector(objc.GetSelector("pasteboardFinishedWithDataProvider:"))
}

func (p_ PasteboardItemDataProviderWrapper) PasteboardFinishedWithDataProvider(pasteboard IPasteboard) {
	ffi.CallMethod[ffi.Void](p_, "pasteboardFinishedWithDataProvider:", pasteboard)
}
