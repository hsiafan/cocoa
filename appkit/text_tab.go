// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextTabClass = _TextTabClass{objc.GetClass("NSTextTab")}

type _TextTabClass struct {
	objc.Class
}

type ITextTab interface {
	objc.IObject
	Location() float64
	Alignment() TextAlignment
	Options() map[TextTabOptionKey]objc.Object
	// deprecated
	TabStopType() TextTabType
}

type TextTab struct {
	objc.Object
}

func MakeTextTab(ptr unsafe.Pointer) TextTab {
	return TextTab{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextTab) InitWithTextAlignment_Location_Options(alignment TextAlignment, loc float64, options map[TextTabOptionKey]objc.IObject) TextTab {
	rv := objc.CallMethod[TextTab](t_, "initWithTextAlignment:location:options:", alignment, loc, options)
	return rv
}

func (t_ TextTab) InitWithType_Location(type_ TextTabType, loc float64) TextTab {
	rv := objc.CallMethod[TextTab](t_, "initWithType:location:", type_, loc)
	return rv
}

func (tc _TextTabClass) Alloc() TextTab {
	rv := objc.CallMethod[TextTab](tc, "alloc")
	return rv
}

func (tc _TextTabClass) New() TextTab {
	rv := objc.CallMethod[TextTab](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextTab() TextTab {
	return TextTabClass.New()
}

func (t_ TextTab) Init() TextTab {
	rv := objc.CallMethod[TextTab](t_, "init")
	return rv
}

func (tc _TextTabClass) ColumnTerminatorsForLocale(aLocale foundation.ILocale) foundation.CharacterSet {
	rv := objc.CallMethod[foundation.CharacterSet](tc, "columnTerminatorsForLocale:", aLocale)
	return rv
}

func (t_ TextTab) Location() float64 {
	rv := objc.CallMethod[float64](t_, "location")
	return rv
}

func (t_ TextTab) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](t_, "alignment")
	return rv
}

func (t_ TextTab) Options() map[TextTabOptionKey]objc.Object {
	rv := objc.CallMethod[map[TextTabOptionKey]objc.Object](t_, "options")
	return rv
}

// deprecated
func (t_ TextTab) TabStopType() TextTabType {
	rv := objc.CallMethod[TextTabType](t_, "tabStopType")
	return rv
}
