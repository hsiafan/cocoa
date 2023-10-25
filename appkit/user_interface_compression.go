// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type UserInterfaceCompression interface {
	// required
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions)
	// required
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) foundation.Size
	ImplementsActiveCompressionOptions() bool
	// optional
	ActiveCompressionOptions() IUserInterfaceCompressionOptions
}

func WrapUserInterfaceCompression(v UserInterfaceCompression) objc.Object {
	return objc.WrapAsProtocol("NSUserInterfaceCompression", v)
}

type UserInterfaceCompressionBase struct {
}

func (p *UserInterfaceCompressionBase) ImplementsActiveCompressionOptions() bool {
	return false
}

func (p *UserInterfaceCompressionBase) ActiveCompressionOptions() IUserInterfaceCompressionOptions {
	panic("unimpemented protocol method")
}
