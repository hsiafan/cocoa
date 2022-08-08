// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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

var UserInterfaceCompressionOptionsClass = _UserInterfaceCompressionOptionsClass{objc.GetClass("NSUserInterfaceCompressionOptions")}

type _UserInterfaceCompressionOptionsClass struct {
	objc.Class
}

type IUserInterfaceCompressionOptions interface {
	objc.IObject
	ContainsOptions(options IUserInterfaceCompressionOptions) bool
	IntersectsOptions(options IUserInterfaceCompressionOptions) bool
	OptionsByAddingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions
	OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions
	IsEmpty() bool
}

type UserInterfaceCompressionOptions struct {
	objc.Object
}

func MakeUserInterfaceCompressionOptions(ptr unsafe.Pointer) UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptions{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ UserInterfaceCompressionOptions) Init() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](u_, "init")
	return rv
}

func (u_ UserInterfaceCompressionOptions) InitWithCompressionOptions(options foundation.ISet) UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](u_, "initWithCompressionOptions:", options)
	return rv
}

func (u_ UserInterfaceCompressionOptions) InitWithIdentifier(identifier string) UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](u_, "initWithIdentifier:", identifier)
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) Alloc() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](uc, "alloc")
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) New() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](uc, "new")
	rv.Autorelease()
	return rv
}

func NewUserInterfaceCompressionOptions() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.New()
}

func (u_ UserInterfaceCompressionOptions) ContainsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := ffi.CallMethod[bool](u_, "containsOptions:", options)
	return rv
}

func (u_ UserInterfaceCompressionOptions) IntersectsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := ffi.CallMethod[bool](u_, "intersectsOptions:", options)
	return rv
}

func (u_ UserInterfaceCompressionOptions) OptionsByAddingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](u_, "optionsByAddingOptions:", options)
	return rv
}

func (u_ UserInterfaceCompressionOptions) OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](u_, "optionsByRemovingOptions:", options)
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) HideImagesOption() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](uc, "hideImagesOption")
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) HideTextOption() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](uc, "hideTextOption")
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) ReduceMetricsOption() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](uc, "reduceMetricsOption")
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](uc, "breakEqualWidthsOption")
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) StandardOptions() UserInterfaceCompressionOptions {
	rv := ffi.CallMethod[UserInterfaceCompressionOptions](uc, "standardOptions")
	return rv
}

func (u_ UserInterfaceCompressionOptions) IsEmpty() bool {
	rv := ffi.CallMethod[bool](u_, "isEmpty")
	return rv
}
