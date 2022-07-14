// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type UserInterfaceItemIdentification interface {
	ImplementsSetIdentifier() bool
	// optional
	SetIdentifier(value UserInterfaceItemIdentifier)
	ImplementsIdentifier() bool
	// optional
	Identifier() UserInterfaceItemIdentifier
}

type UserInterfaceItemIdentificationWrapper struct {
	objc.Object
}

func (u_ *UserInterfaceItemIdentificationWrapper) ImplementsSetIdentifier() bool {
	return u_.RespondsToSelector(objc.GetSelector("setIdentifier:"))
}

func (u_ UserInterfaceItemIdentificationWrapper) SetIdentifier(value UserInterfaceItemIdentifier) {
	ffi.CallMethod[ffi.Void](u_, "setIdentifier:", value)
}

func (u_ *UserInterfaceItemIdentificationWrapper) ImplementsIdentifier() bool {
	return u_.RespondsToSelector(objc.GetSelector("identifier"))
}

func (u_ UserInterfaceItemIdentificationWrapper) Identifier() UserInterfaceItemIdentifier {
	rv := ffi.CallMethod[UserInterfaceItemIdentifier](u_, "identifier")
	return rv
}
