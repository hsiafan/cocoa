// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

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
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("init"))
	return rv
}

func (u_ UserInterfaceCompressionOptions) InitWithCompressionOptions(options foundation.ISet) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("initWithCompressionOptions:"), options)
	return rv
}

func (u_ UserInterfaceCompressionOptions) InitWithIdentifier(identifier string) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("initWithIdentifier:"), identifier)
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) Alloc() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("alloc"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) New() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewUserInterfaceCompressionOptions() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.New()
}

func (u_ UserInterfaceCompressionOptions) ContainsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("containsOptions:"), options)
	return rv
}

func (u_ UserInterfaceCompressionOptions) IntersectsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("intersectsOptions:"), options)
	return rv
}

func (u_ UserInterfaceCompressionOptions) OptionsByAddingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("optionsByAddingOptions:"), options)
	return rv
}

func (u_ UserInterfaceCompressionOptions) OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.GetSelector("optionsByRemovingOptions:"), options)
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) HideImagesOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("hideImagesOption"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) HideTextOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("hideTextOption"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) ReduceMetricsOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("reduceMetricsOption"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("breakEqualWidthsOption"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) StandardOptions() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.GetSelector("standardOptions"))
	return rv
}

func (u_ UserInterfaceCompressionOptions) IsEmpty() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isEmpty"))
	return rv
}
