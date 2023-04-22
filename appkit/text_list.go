// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var TextListClass = _TextListClass{objc.GetClass("NSTextList")}

type _TextListClass struct {
	objc.Class
}

type ITextList interface {
	objc.IObject
	MarkerForItemNumber(itemNumber int) string
	MarkerFormat() TextListMarkerFormat
	IsOrdered() bool
	ListOptions() TextListOptions
	StartingItemNumber() int
	SetStartingItemNumber(value int)
}

type TextList struct {
	objc.Object
}

func MakeTextList(ptr unsafe.Pointer) TextList {
	return TextList{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextList) InitWithMarkerFormat_Options(markerFormat TextListMarkerFormat, options uint) TextList {
	rv := objc.CallMethod[TextList](t_, objc.GetSelector("initWithMarkerFormat:options:"), markerFormat, options)
	return rv
}

func (t_ TextList) InitWithMarkerFormat_Options_StartingItemNumber(markerFormat TextListMarkerFormat, options TextListOptions, startingItemNumber int) TextList {
	rv := objc.CallMethod[TextList](t_, objc.GetSelector("initWithMarkerFormat:options:startingItemNumber:"), markerFormat, options, startingItemNumber)
	return rv
}

func (tc _TextListClass) Alloc() TextList {
	rv := objc.CallMethod[TextList](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TextListClass) New() TextList {
	rv := objc.CallMethod[TextList](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextList() TextList {
	return TextListClass.New()
}

func (t_ TextList) Init() TextList {
	rv := objc.CallMethod[TextList](t_, objc.GetSelector("init"))
	return rv
}

func (t_ TextList) MarkerForItemNumber(itemNumber int) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("markerForItemNumber:"), itemNumber)
	return rv
}

func (t_ TextList) MarkerFormat() TextListMarkerFormat {
	rv := objc.CallMethod[TextListMarkerFormat](t_, objc.GetSelector("markerFormat"))
	return rv
}

func (t_ TextList) IsOrdered() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isOrdered"))
	return rv
}

func (t_ TextList) ListOptions() TextListOptions {
	rv := objc.CallMethod[TextListOptions](t_, objc.GetSelector("listOptions"))
	return rv
}

func (t_ TextList) StartingItemNumber() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("startingItemNumber"))
	return rv
}

func (t_ TextList) SetStartingItemNumber(value int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setStartingItemNumber:"), value)
}
