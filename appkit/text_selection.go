// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextSelectionClass = _TextSelectionClass{objc.GetClass("NSTextSelection")}

type _TextSelectionClass struct {
	objc.Class
}

type ITextSelection interface {
	objc.IObject
	TextSelectionWithTextRanges(textRanges []ITextRange) TextSelection
	Affinity() TextSelectionAffinity
	AnchorPositionOffset() float64
	SetAnchorPositionOffset(value float64)
	Granularity() TextSelectionGranularity
	IsLogical() bool
	SetLogical(value bool)
	IsTransient() bool
	TextRanges() []TextRange
	TypingAttributes() map[foundation.AttributedStringKey]objc.Object
	SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject)
}

type TextSelection struct {
	objc.Object
}

func MakeTextSelection(ptr unsafe.Pointer) TextSelection {
	return TextSelection{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextSelection) InitWithRange_Affinity_Granularity(range_ ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.SEL("initWithRange:affinity:granularity:"), objc.ExtractPtr(range_), affinity, granularity)
	return rv
}

func (t_ TextSelection) InitWithRanges_Affinity_Granularity(textRanges []ITextRange, affinity TextSelectionAffinity, granularity TextSelectionGranularity) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.SEL("initWithRanges:affinity:granularity:"), textRanges, affinity, granularity)
	return rv
}

func (tc _TextSelectionClass) Alloc() TextSelection {
	rv := objc.CallMethod[TextSelection](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextSelectionClass) New() TextSelection {
	rv := objc.CallMethod[TextSelection](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextSelection() TextSelection {
	return TextSelectionClass.New()
}

func (t_ TextSelection) Init() TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.SEL("init"))
	return rv
}

func (t_ TextSelection) TextSelectionWithTextRanges(textRanges []ITextRange) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, objc.SEL("textSelectionWithTextRanges:"), textRanges)
	return rv
}

func (t_ TextSelection) Affinity() TextSelectionAffinity {
	rv := objc.CallMethod[TextSelectionAffinity](t_, objc.SEL("affinity"))
	return rv
}

func (t_ TextSelection) AnchorPositionOffset() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("anchorPositionOffset"))
	return rv
}

func (t_ TextSelection) SetAnchorPositionOffset(value float64) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAnchorPositionOffset:"), value)
}

func (t_ TextSelection) Granularity() TextSelectionGranularity {
	rv := objc.CallMethod[TextSelectionGranularity](t_, objc.SEL("granularity"))
	return rv
}

func (t_ TextSelection) IsLogical() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isLogical"))
	return rv
}

func (t_ TextSelection) SetLogical(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLogical:"), value)
}

func (t_ TextSelection) IsTransient() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isTransient"))
	return rv
}

func (t_ TextSelection) TextRanges() []TextRange {
	rv := objc.CallMethod[[]TextRange](t_, objc.SEL("textRanges"))
	return rv
}

func (t_ TextSelection) TypingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.SEL("typingAttributes"))
	return rv
}

func (t_ TextSelection) SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setTypingAttributes:"), value)
}
