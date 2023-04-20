// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var PasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}

type _PasteboardItemClass struct {
	objc.Class
}

type IPasteboardItem interface {
	objc.IObject
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	SetDataProvider_ForTypes(dataProvider PasteboardItemDataProvider, types []PasteboardType) bool
	SetDataProvider0_ForTypes(dataProvider objc.IObject, types []PasteboardType) bool
	SetData_ForType(data []byte, type_ PasteboardType) bool
	SetString_ForType(string_ string, type_ PasteboardType) bool
	SetPropertyList_ForType(propertyList objc.IObject, type_ PasteboardType) bool
	DataForType(type_ PasteboardType) []byte
	StringForType(type_ PasteboardType) string
	PropertyListForType(type_ PasteboardType) objc.Object
	Types() []PasteboardType
}

type PasteboardItem struct {
	objc.Object
}

func MakePasteboardItem(ptr unsafe.Pointer) PasteboardItem {
	return PasteboardItem{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PasteboardItemClass) Alloc() PasteboardItem {
	rv := ffi.CallMethod[PasteboardItem](pc, "alloc")
	return rv
}

func (pc _PasteboardItemClass) New() PasteboardItem {
	rv := ffi.CallMethod[PasteboardItem](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPasteboardItem() PasteboardItem {
	return PasteboardItemClass.New()
}

func (p_ PasteboardItem) Init() PasteboardItem {
	rv := ffi.CallMethod[PasteboardItem](p_, "init")
	return rv
}

func (p_ PasteboardItem) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	rv := ffi.CallMethod[PasteboardType](p_, "availableTypeFromArray:", types)
	return rv
}

func (p_ PasteboardItem) SetDataProvider_ForTypes(dataProvider PasteboardItemDataProvider, types []PasteboardType) bool {
	po := ffi.CreateProtocol("NSPasteboardItemDataProvider", dataProvider)
	defer po.Release()
	rv := ffi.CallMethod[bool](p_, "setDataProvider:forTypes:", po, types)
	return rv
}

func (p_ PasteboardItem) SetDataProvider0_ForTypes(dataProvider objc.IObject, types []PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setDataProvider:forTypes:", dataProvider, types)
	return rv
}

func (p_ PasteboardItem) SetData_ForType(data []byte, type_ PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setData:forType:", data, type_)
	return rv
}

func (p_ PasteboardItem) SetString_ForType(string_ string, type_ PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setString:forType:", string_, type_)
	return rv
}

func (p_ PasteboardItem) SetPropertyList_ForType(propertyList objc.IObject, type_ PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setPropertyList:forType:", propertyList, type_)
	return rv
}

func (p_ PasteboardItem) DataForType(type_ PasteboardType) []byte {
	rv := ffi.CallMethod[[]byte](p_, "dataForType:", type_)
	return rv
}

func (p_ PasteboardItem) StringForType(type_ PasteboardType) string {
	rv := ffi.CallMethod[string](p_, "stringForType:", type_)
	return rv
}

func (p_ PasteboardItem) PropertyListForType(type_ PasteboardType) objc.Object {
	rv := ffi.CallMethod[objc.Object](p_, "propertyListForType:", type_)
	return rv
}

func (p_ PasteboardItem) Types() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](p_, "types")
	return rv
}
