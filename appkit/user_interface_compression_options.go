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
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.SEL("init"))
	return rv
}

func (u_ UserInterfaceCompressionOptions) InitWithCompressionOptions(options foundation.ISet) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.SEL("initWithCompressionOptions:"), objc.ExtractPtr(options))
	return rv
}

func (u_ UserInterfaceCompressionOptions) InitWithIdentifier(identifier string) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.SEL("initWithIdentifier:"), identifier)
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) Alloc() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.SEL("alloc"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) New() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewUserInterfaceCompressionOptions() UserInterfaceCompressionOptions {
	return UserInterfaceCompressionOptionsClass.New()
}

func (u_ UserInterfaceCompressionOptions) ContainsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("containsOptions:"), objc.ExtractPtr(options))
	return rv
}

func (u_ UserInterfaceCompressionOptions) IntersectsOptions(options IUserInterfaceCompressionOptions) bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("intersectsOptions:"), objc.ExtractPtr(options))
	return rv
}

func (u_ UserInterfaceCompressionOptions) OptionsByAddingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.SEL("optionsByAddingOptions:"), objc.ExtractPtr(options))
	return rv
}

func (u_ UserInterfaceCompressionOptions) OptionsByRemovingOptions(options IUserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](u_, objc.SEL("optionsByRemovingOptions:"), objc.ExtractPtr(options))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) HideImagesOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.SEL("hideImagesOption"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) HideTextOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.SEL("hideTextOption"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) ReduceMetricsOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.SEL("reduceMetricsOption"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.SEL("breakEqualWidthsOption"))
	return rv
}

func (uc _UserInterfaceCompressionOptionsClass) StandardOptions() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](uc, objc.SEL("standardOptions"))
	return rv
}

func (u_ UserInterfaceCompressionOptions) IsEmpty() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("isEmpty"))
	return rv
}
