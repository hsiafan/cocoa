// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var FormatterClass = _FormatterClass{objc.GetClass("NSFormatter")}

type _FormatterClass struct {
	objc.Class
}

type IFormatter interface {
	objc.IObject
	StringForObjectValue(obj objc.IObject) string
	AttributedStringForObjectValue_WithDefaultAttributes(obj objc.IObject, attrs map[AttributedStringKey]objc.IObject) AttributedString
	EditingStringForObjectValue(obj objc.IObject) string
	IsPartialStringValid_NewEditingString_ErrorDescription(partialString string, newString *String, error *String) bool
	IsPartialStringValid_ProposedSelectedRange_OriginalString_OriginalSelectedRange_ErrorDescription(partialStringPtr *String, proposedSelRangePtr *Range, origString string, origSelRange Range, error *String) bool
}

type Formatter struct {
	objc.Object
}

func MakeFormatter(ptr unsafe.Pointer) Formatter {
	return Formatter{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FormatterClass) Alloc() Formatter {
	rv := objc.CallMethod[Formatter](fc, objc.GetSelector("alloc"))
	return rv
}

func (fc _FormatterClass) New() Formatter {
	rv := objc.CallMethod[Formatter](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFormatter() Formatter {
	return FormatterClass.New()
}

func (f_ Formatter) Init() Formatter {
	rv := objc.CallMethod[Formatter](f_, objc.GetSelector("init"))
	return rv
}

func (f_ Formatter) StringForObjectValue(obj objc.IObject) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("stringForObjectValue:"), obj)
	return rv
}

func (f_ Formatter) AttributedStringForObjectValue_WithDefaultAttributes(obj objc.IObject, attrs map[AttributedStringKey]objc.IObject) AttributedString {
	rv := objc.CallMethod[AttributedString](f_, objc.GetSelector("attributedStringForObjectValue:withDefaultAttributes:"), obj, attrs)
	return rv
}

func (f_ Formatter) EditingStringForObjectValue(obj objc.IObject) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("editingStringForObjectValue:"), obj)
	return rv
}

func (f_ Formatter) IsPartialStringValid_NewEditingString_ErrorDescription(partialString string, newString *String, error *String) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isPartialStringValid:newEditingString:errorDescription:"), partialString, newString, error)
	return rv
}

func (f_ Formatter) IsPartialStringValid_ProposedSelectedRange_OriginalString_OriginalSelectedRange_ErrorDescription(partialStringPtr *String, proposedSelRangePtr *Range, origString string, origSelRange Range, error *String) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isPartialStringValid:proposedSelectedRange:originalString:originalSelectedRange:errorDescription:"), partialStringPtr, proposedSelRangePtr, origString, origSelRange, error)
	return rv
}
