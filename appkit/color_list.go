// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ColorListClass = _ColorListClass{objc.GetClass("NSColorList")}

type _ColorListClass struct {
	objc.Class
}

type IColorList interface {
	objc.IObject
	ColorWithKey(key ColorName) Color
	InsertColor_Key_AtIndex(color IColor, key ColorName, loc uint)
	RemoveColorWithKey(key ColorName)
	SetColor_ForKey(color IColor, key ColorName)
	WriteToURL_Error(url foundation.IURL, errPtr *foundation.Error) bool
	RemoveFile()
	// deprecated
	WriteToFile(path string) bool
	Name() ColorListName
	IsEditable() bool
	AllKeys() []ColorName
}

type ColorList struct {
	objc.Object
}

func MakeColorList(ptr unsafe.Pointer) ColorList {
	return ColorList{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ ColorList) InitWithName(name ColorListName) ColorList {
	rv := objc.CallMethod[ColorList](c_, objc.SEL("initWithName:"), name)
	return rv
}

func (c_ ColorList) InitWithName_FromFile(name ColorListName, path string) ColorList {
	rv := objc.CallMethod[ColorList](c_, objc.SEL("initWithName:fromFile:"), name, path)
	return rv
}

func (cc _ColorListClass) Alloc() ColorList {
	rv := objc.CallMethod[ColorList](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ColorListClass) New() ColorList {
	rv := objc.CallMethod[ColorList](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewColorList() ColorList {
	return ColorListClass.New()
}

func (c_ ColorList) Init() ColorList {
	rv := objc.CallMethod[ColorList](c_, objc.SEL("init"))
	return rv
}

func (cc _ColorListClass) ColorListNamed(name ColorListName) ColorList {
	rv := objc.CallMethod[ColorList](cc, objc.SEL("colorListNamed:"), name)
	return rv
}

func (c_ ColorList) ColorWithKey(key ColorName) Color {
	rv := objc.CallMethod[Color](c_, objc.SEL("colorWithKey:"), key)
	return rv
}

func (c_ ColorList) InsertColor_Key_AtIndex(color IColor, key ColorName, loc uint) {
	objc.CallMethod[objc.Void](c_, objc.SEL("insertColor:key:atIndex:"), objc.ExtractPtr(color), key, loc)
}

func (c_ ColorList) RemoveColorWithKey(key ColorName) {
	objc.CallMethod[objc.Void](c_, objc.SEL("removeColorWithKey:"), key)
}

func (c_ ColorList) SetColor_ForKey(color IColor, key ColorName) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setColor:forKey:"), objc.ExtractPtr(color), key)
}

func (c_ ColorList) WriteToURL_Error(url foundation.IURL, errPtr *foundation.Error) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("writeToURL:error:"), objc.ExtractPtr(url), unsafe.Pointer(errPtr))
	return rv
}

func (c_ ColorList) RemoveFile() {
	objc.CallMethod[objc.Void](c_, objc.SEL("removeFile"))
}

// deprecated
func (c_ ColorList) WriteToFile(path string) bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("writeToFile:"), path)
	return rv
}

func (cc _ColorListClass) AvailableColorLists() []ColorList {
	rv := objc.CallMethod[[]ColorList](cc, objc.SEL("availableColorLists"))
	return rv
}

func (c_ ColorList) Name() ColorListName {
	rv := objc.CallMethod[ColorListName](c_, objc.SEL("name"))
	return rv
}

func (c_ ColorList) IsEditable() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isEditable"))
	return rv
}

func (c_ ColorList) AllKeys() []ColorName {
	rv := objc.CallMethod[[]ColorName](c_, objc.SEL("allKeys"))
	return rv
}
