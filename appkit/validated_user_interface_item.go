// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type ValidatedUserInterfaceItem interface {
	ImplementsAction() bool
	// optional
	Action() objc.Selector
	ImplementsTag() bool
	// optional
	Tag() int
}

func WrapValidatedUserInterfaceItem(v ValidatedUserInterfaceItem) objc.Object {
	return objc.WrapAsProtocol("NSValidatedUserInterfaceItem", v)
}

type ValidatedUserInterfaceItemBase struct {
}

func (p *ValidatedUserInterfaceItemBase) ImplementsAction() bool {
	return false
}

func (p *ValidatedUserInterfaceItemBase) Action() objc.Selector {
	panic("unimpemented protocol method")
}

func (p *ValidatedUserInterfaceItemBase) ImplementsTag() bool {
	return false
}

func (p *ValidatedUserInterfaceItemBase) Tag() int {
	panic("unimpemented protocol method")
}
