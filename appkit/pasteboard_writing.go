// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type PasteboardWriting interface {
	// required
	WritableTypesForPasteboard(pasteboard Pasteboard) []PasteboardType
	ImplementsWritingOptionsForType_Pasteboard() bool
	// optional
	WritingOptionsForType_Pasteboard(type_ PasteboardType, pasteboard Pasteboard) PasteboardWritingOptions
	// required
	PasteboardPropertyListForType(type_ PasteboardType) objc.IObject
}

func WrapPasteboardWriting(v PasteboardWriting) objc.Object {
	return objc.WrapAsProtocol("NSPasteboardWriting", v)
}

type PasteboardWritingBase struct {
}

func (p *PasteboardWritingBase) ImplementsWritingOptionsForType_Pasteboard() bool {
	return false
}

func (p *PasteboardWritingBase) WritingOptionsForType_Pasteboard(type_ PasteboardType, pasteboard Pasteboard) PasteboardWritingOptions {
	panic("unimpemented protocol method")
}
