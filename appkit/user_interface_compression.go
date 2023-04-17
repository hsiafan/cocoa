// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type UserInterfaceCompressionWrapper struct {
	objc.Object
}

func (u_ UserInterfaceCompressionWrapper) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	ffi.CallMethod[ffi.Void](u_, "compressWithPrioritizedCompressionOptions:", prioritizedOptions)
}

func (u_ UserInterfaceCompressionWrapper) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](u_, "minimumSizeWithPrioritizedCompressionOptions:", prioritizedOptions)
	return rv
}

func (u_ *UserInterfaceCompressionWrapper) ImplementsActiveCompressionOptions() bool {
	return u_.RespondsToSelector(objc.GetSelector("activeCompressionOptions"))
}

func (u_ UserInterfaceCompressionWrapper) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](u_, "activeCompressionOptions")
	return rv
}
