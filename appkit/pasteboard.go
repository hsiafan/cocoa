// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PasteboardClass = _PasteboardClass{objc.GetClass("NSPasteboard")}

type _PasteboardClass struct {
	objc.Class
}

type IPasteboard interface {
	objc.IObject
	ClearContents() int
	SetData_ForType(data []byte, dataType PasteboardType) bool
	SetPropertyList_ForType(plist objc.IObject, dataType PasteboardType) bool
	SetString_ForType(string_ string, dataType PasteboardType) bool
	ReadObjectsForClasses_Options(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) []objc.Object
	IndexOfPasteboardItem(pasteboardItem IPasteboardItem) uint
	DataForType(dataType PasteboardType) []byte
	PropertyListForType(dataType PasteboardType) objc.Object
	StringForType(dataType PasteboardType) string
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	CanReadItemWithDataConformingToTypes(types []string) bool
	CanReadObjectForClasses_Options(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) bool
	PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int
	DeclareTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int
	AddTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int
	WriteFileContents(filename string) bool
	WriteFileWrapper(wrapper foundation.IFileWrapper) bool
	ReadFileContentsType_ToFile(type_ PasteboardType, filename string) string
	ReadFileWrapper() foundation.FileWrapper
	PasteboardItems() []PasteboardItem
	Types() []PasteboardType
	Name() PasteboardName
	ChangeCount() int
}

type Pasteboard struct {
	objc.Object
}

func MakePasteboard(ptr unsafe.Pointer) Pasteboard {
	return Pasteboard{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PasteboardClass) Alloc() Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, "alloc")
	return rv
}

func (pc _PasteboardClass) New() Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPasteboard() Pasteboard {
	return PasteboardClass.New()
}

func (p_ Pasteboard) Init() Pasteboard {
	rv := objc.CallMethod[Pasteboard](p_, "init")
	return rv
}

func (pc _PasteboardClass) PasteboardByFilteringData_OfType(data []byte, type_ PasteboardType) Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, "pasteboardByFilteringData:ofType:", data, type_)
	return rv
}

func (pc _PasteboardClass) PasteboardByFilteringFile(filename string) Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, "pasteboardByFilteringFile:", filename)
	return rv
}

func (pc _PasteboardClass) PasteboardByFilteringTypesInPasteboard(pboard IPasteboard) Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, "pasteboardByFilteringTypesInPasteboard:", pboard)
	return rv
}

func (pc _PasteboardClass) PasteboardWithName(name PasteboardName) Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, "pasteboardWithName:", name)
	return rv
}

func (pc _PasteboardClass) PasteboardWithUniqueName() Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, "pasteboardWithUniqueName")
	return rv
}

func (p_ Pasteboard) ClearContents() int {
	rv := objc.CallMethod[int](p_, "clearContents")
	return rv
}

func (p_ Pasteboard) SetData_ForType(data []byte, dataType PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, "setData:forType:", data, dataType)
	return rv
}

func (p_ Pasteboard) SetPropertyList_ForType(plist objc.IObject, dataType PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, "setPropertyList:forType:", plist, dataType)
	return rv
}

func (p_ Pasteboard) SetString_ForType(string_ string, dataType PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, "setString:forType:", string_, dataType)
	return rv
}

func (p_ Pasteboard) ReadObjectsForClasses_Options(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](p_, "readObjectsForClasses:options:", classArray, options)
	return rv
}

func (p_ Pasteboard) IndexOfPasteboardItem(pasteboardItem IPasteboardItem) uint {
	rv := objc.CallMethod[uint](p_, "indexOfPasteboardItem:", pasteboardItem)
	return rv
}

func (p_ Pasteboard) DataForType(dataType PasteboardType) []byte {
	rv := objc.CallMethod[[]byte](p_, "dataForType:", dataType)
	return rv
}

func (p_ Pasteboard) PropertyListForType(dataType PasteboardType) objc.Object {
	rv := objc.CallMethod[objc.Object](p_, "propertyListForType:", dataType)
	return rv
}

func (p_ Pasteboard) StringForType(dataType PasteboardType) string {
	rv := objc.CallMethod[string](p_, "stringForType:", dataType)
	return rv
}

func (p_ Pasteboard) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	rv := objc.CallMethod[PasteboardType](p_, "availableTypeFromArray:", types)
	return rv
}

func (p_ Pasteboard) CanReadItemWithDataConformingToTypes(types []string) bool {
	rv := objc.CallMethod[bool](p_, "canReadItemWithDataConformingToTypes:", types)
	return rv
}

func (p_ Pasteboard) CanReadObjectForClasses_Options(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) bool {
	rv := objc.CallMethod[bool](p_, "canReadObjectForClasses:options:", classArray, options)
	return rv
}

func (pc _PasteboardClass) TypesFilterableTo(type_ PasteboardType) []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](pc, "typesFilterableTo:", type_)
	return rv
}

func (p_ Pasteboard) PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int {
	rv := objc.CallMethod[int](p_, "prepareForNewContentsWithOptions:", options)
	return rv
}

func (p_ Pasteboard) DeclareTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int {
	rv := objc.CallMethod[int](p_, "declareTypes:owner:", newTypes, newOwner)
	return rv
}

func (p_ Pasteboard) AddTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int {
	rv := objc.CallMethod[int](p_, "addTypes:owner:", newTypes, newOwner)
	return rv
}

func (p_ Pasteboard) WriteFileContents(filename string) bool {
	rv := objc.CallMethod[bool](p_, "writeFileContents:", filename)
	return rv
}

func (p_ Pasteboard) WriteFileWrapper(wrapper foundation.IFileWrapper) bool {
	rv := objc.CallMethod[bool](p_, "writeFileWrapper:", wrapper)
	return rv
}

func (p_ Pasteboard) ReadFileContentsType_ToFile(type_ PasteboardType, filename string) string {
	rv := objc.CallMethod[string](p_, "readFileContentsType:toFile:", type_, filename)
	return rv
}

func (p_ Pasteboard) ReadFileWrapper() foundation.FileWrapper {
	rv := objc.CallMethod[foundation.FileWrapper](p_, "readFileWrapper")
	return rv
}

func (pc _PasteboardClass) GeneralPasteboard() Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, "generalPasteboard")
	return rv
}

func (p_ Pasteboard) PasteboardItems() []PasteboardItem {
	rv := objc.CallMethod[[]PasteboardItem](p_, "pasteboardItems")
	return rv
}

func (p_ Pasteboard) Types() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](p_, "types")
	return rv
}

func (p_ Pasteboard) Name() PasteboardName {
	rv := objc.CallMethod[PasteboardName](p_, "name")
	return rv
}

func (p_ Pasteboard) ChangeCount() int {
	rv := objc.CallMethod[int](p_, "changeCount")
	return rv
}
