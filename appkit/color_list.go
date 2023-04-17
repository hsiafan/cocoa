// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ColorList](c_, "initWithName:", name)
	return rv
}

func (c_ ColorList) InitWithName_FromFile(name ColorListName, path string) ColorList {
	rv := ffi.CallMethod[ColorList](c_, "initWithName:fromFile:", name, path)
	return rv
}

func (cc _ColorListClass) Alloc() ColorList {
	rv := ffi.CallMethod[ColorList](cc, "alloc")
	return rv
}

func (cc _ColorListClass) New() ColorList {
	rv := ffi.CallMethod[ColorList](cc, "new")
	rv.Autorelease()
	return rv
}

func NewColorList() ColorList {
	return ColorListClass.New()
}

func (c_ ColorList) Init() ColorList {
	rv := ffi.CallMethod[ColorList](c_, "init")
	return rv
}

func (cc _ColorListClass) ColorListNamed(name ColorListName) ColorList {
	rv := ffi.CallMethod[ColorList](cc, "colorListNamed:", name)
	return rv
}

func (c_ ColorList) ColorWithKey(key ColorName) Color {
	rv := ffi.CallMethod[Color](c_, "colorWithKey:", key)
	return rv
}

func (c_ ColorList) InsertColor_Key_AtIndex(color IColor, key ColorName, loc uint) {
	ffi.CallMethod[ffi.Void](c_, "insertColor:key:atIndex:", color, key, loc)
}

func (c_ ColorList) RemoveColorWithKey(key ColorName) {
	ffi.CallMethod[ffi.Void](c_, "removeColorWithKey:", key)
}

func (c_ ColorList) SetColor_ForKey(color IColor, key ColorName) {
	ffi.CallMethod[ffi.Void](c_, "setColor:forKey:", color, key)
}

func (c_ ColorList) WriteToURL_Error(url foundation.IURL, errPtr *foundation.Error) bool {
	rv := ffi.CallMethod[bool](c_, "writeToURL:error:", url, unsafe.Pointer(errPtr))
	return rv
}

func (c_ ColorList) RemoveFile() {
	ffi.CallMethod[ffi.Void](c_, "removeFile")
}

// deprecated
func (c_ ColorList) WriteToFile(path string) bool {
	rv := ffi.CallMethod[bool](c_, "writeToFile:", path)
	return rv
}

func (cc _ColorListClass) AvailableColorLists() []ColorList {
	rv := ffi.CallMethod[[]ColorList](cc, "availableColorLists")
	return rv
}

func (c_ ColorList) Name() ColorListName {
	rv := ffi.CallMethod[ColorListName](c_, "name")
	return rv
}

func (c_ ColorList) IsEditable() bool {
	rv := ffi.CallMethod[bool](c_, "isEditable")
	return rv
}

func (c_ ColorList) AllKeys() []ColorName {
	rv := ffi.CallMethod[[]ColorName](c_, "allKeys")
	return rv
}
