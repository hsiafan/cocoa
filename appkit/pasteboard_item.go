// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var PasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}

type _PasteboardItemClass struct {
	objc.Class
}

type IPasteboardItem interface {
	objc.IObject
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	SetDataProvider_ForTypes(dataProvider objc.IObject, types []PasteboardType) bool
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
	rv := objc.CallMethod[PasteboardItem](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PasteboardItemClass) New() PasteboardItem {
	rv := objc.CallMethod[PasteboardItem](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPasteboardItem() PasteboardItem {
	return PasteboardItemClass.New()
}

func (p_ PasteboardItem) Init() PasteboardItem {
	rv := objc.CallMethod[PasteboardItem](p_, objc.SEL("init"))
	return rv
}

func (p_ PasteboardItem) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	rv := objc.CallMethod[PasteboardType](p_, objc.SEL("availableTypeFromArray:"), types)
	return rv
}

func (p_ PasteboardItem) SetDataProvider_ForTypes(dataProvider objc.IObject, types []PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("setDataProvider:forTypes:"), objc.ExtractPtr(dataProvider), types)
	return rv
}

func (p_ PasteboardItem) SetData_ForType(data []byte, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("setData:forType:"), data, type_)
	return rv
}

func (p_ PasteboardItem) SetString_ForType(string_ string, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("setString:forType:"), string_, type_)
	return rv
}

func (p_ PasteboardItem) SetPropertyList_ForType(propertyList objc.IObject, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("setPropertyList:forType:"), objc.ExtractPtr(propertyList), type_)
	return rv
}

func (p_ PasteboardItem) DataForType(type_ PasteboardType) []byte {
	rv := objc.CallMethod[[]byte](p_, objc.SEL("dataForType:"), type_)
	return rv
}

func (p_ PasteboardItem) StringForType(type_ PasteboardType) string {
	rv := objc.CallMethod[string](p_, objc.SEL("stringForType:"), type_)
	return rv
}

func (p_ PasteboardItem) PropertyListForType(type_ PasteboardType) objc.Object {
	rv := objc.CallMethod[objc.Object](p_, objc.SEL("propertyListForType:"), type_)
	return rv
}

func (p_ PasteboardItem) Types() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](p_, objc.SEL("types"))
	return rv
}
