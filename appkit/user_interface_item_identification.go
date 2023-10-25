// auto-generated code, do not modify
package appkit

import (
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

func WrapUserInterfaceItemIdentification(v UserInterfaceItemIdentification) objc.Object {
	return objc.WrapAsProtocol("NSUserInterfaceItemIdentification", v)
}

type UserInterfaceItemIdentificationBase struct {
}

func (p *UserInterfaceItemIdentificationBase) ImplementsSetIdentifier() bool {
	return false
}

func (p *UserInterfaceItemIdentificationBase) SetIdentifier(value UserInterfaceItemIdentifier) {
	panic("unimpemented protocol method")
}

func (p *UserInterfaceItemIdentificationBase) ImplementsIdentifier() bool {
	return false
}

func (p *UserInterfaceItemIdentificationBase) Identifier() UserInterfaceItemIdentifier {
	panic("unimpemented protocol method")
}
