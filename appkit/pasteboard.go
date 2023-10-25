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
	rv := objc.CallMethod[Pasteboard](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PasteboardClass) New() Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPasteboard() Pasteboard {
	return PasteboardClass.New()
}

func (p_ Pasteboard) Init() Pasteboard {
	rv := objc.CallMethod[Pasteboard](p_, objc.SEL("init"))
	return rv
}

func (pc _PasteboardClass) PasteboardByFilteringData_OfType(data []byte, type_ PasteboardType) Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, objc.SEL("pasteboardByFilteringData:ofType:"), data, type_)
	return rv
}

func (pc _PasteboardClass) PasteboardByFilteringFile(filename string) Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, objc.SEL("pasteboardByFilteringFile:"), filename)
	return rv
}

func (pc _PasteboardClass) PasteboardByFilteringTypesInPasteboard(pboard IPasteboard) Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, objc.SEL("pasteboardByFilteringTypesInPasteboard:"), objc.ExtractPtr(pboard))
	return rv
}

func (pc _PasteboardClass) PasteboardWithName(name PasteboardName) Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, objc.SEL("pasteboardWithName:"), name)
	return rv
}

func (pc _PasteboardClass) PasteboardWithUniqueName() Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, objc.SEL("pasteboardWithUniqueName"))
	return rv
}

func (p_ Pasteboard) ClearContents() int {
	rv := objc.CallMethod[int](p_, objc.SEL("clearContents"))
	return rv
}

func (p_ Pasteboard) SetData_ForType(data []byte, dataType PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("setData:forType:"), data, dataType)
	return rv
}

func (p_ Pasteboard) SetPropertyList_ForType(plist objc.IObject, dataType PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("setPropertyList:forType:"), objc.ExtractPtr(plist), dataType)
	return rv
}

func (p_ Pasteboard) SetString_ForType(string_ string, dataType PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("setString:forType:"), string_, dataType)
	return rv
}

func (p_ Pasteboard) ReadObjectsForClasses_Options(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](p_, objc.SEL("readObjectsForClasses:options:"), classArray, options)
	return rv
}

func (p_ Pasteboard) IndexOfPasteboardItem(pasteboardItem IPasteboardItem) uint {
	rv := objc.CallMethod[uint](p_, objc.SEL("indexOfPasteboardItem:"), objc.ExtractPtr(pasteboardItem))
	return rv
}

func (p_ Pasteboard) DataForType(dataType PasteboardType) []byte {
	rv := objc.CallMethod[[]byte](p_, objc.SEL("dataForType:"), dataType)
	return rv
}

func (p_ Pasteboard) PropertyListForType(dataType PasteboardType) objc.Object {
	rv := objc.CallMethod[objc.Object](p_, objc.SEL("propertyListForType:"), dataType)
	return rv
}

func (p_ Pasteboard) StringForType(dataType PasteboardType) string {
	rv := objc.CallMethod[string](p_, objc.SEL("stringForType:"), dataType)
	return rv
}

func (p_ Pasteboard) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	rv := objc.CallMethod[PasteboardType](p_, objc.SEL("availableTypeFromArray:"), types)
	return rv
}

func (p_ Pasteboard) CanReadItemWithDataConformingToTypes(types []string) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("canReadItemWithDataConformingToTypes:"), types)
	return rv
}

func (p_ Pasteboard) CanReadObjectForClasses_Options(classArray []objc.IClass, options map[PasteboardReadingOptionKey]objc.IObject) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("canReadObjectForClasses:options:"), classArray, options)
	return rv
}

func (pc _PasteboardClass) TypesFilterableTo(type_ PasteboardType) []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](pc, objc.SEL("typesFilterableTo:"), type_)
	return rv
}

func (p_ Pasteboard) PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int {
	rv := objc.CallMethod[int](p_, objc.SEL("prepareForNewContentsWithOptions:"), options)
	return rv
}

func (p_ Pasteboard) DeclareTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int {
	rv := objc.CallMethod[int](p_, objc.SEL("declareTypes:owner:"), newTypes, objc.ExtractPtr(newOwner))
	return rv
}

func (p_ Pasteboard) AddTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int {
	rv := objc.CallMethod[int](p_, objc.SEL("addTypes:owner:"), newTypes, objc.ExtractPtr(newOwner))
	return rv
}

func (p_ Pasteboard) WriteFileContents(filename string) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("writeFileContents:"), filename)
	return rv
}

func (p_ Pasteboard) WriteFileWrapper(wrapper foundation.IFileWrapper) bool {
	rv := objc.CallMethod[bool](p_, objc.SEL("writeFileWrapper:"), objc.ExtractPtr(wrapper))
	return rv
}

func (p_ Pasteboard) ReadFileContentsType_ToFile(type_ PasteboardType, filename string) string {
	rv := objc.CallMethod[string](p_, objc.SEL("readFileContentsType:toFile:"), type_, filename)
	return rv
}

func (p_ Pasteboard) ReadFileWrapper() foundation.FileWrapper {
	rv := objc.CallMethod[foundation.FileWrapper](p_, objc.SEL("readFileWrapper"))
	return rv
}

func (pc _PasteboardClass) GeneralPasteboard() Pasteboard {
	rv := objc.CallMethod[Pasteboard](pc, objc.SEL("generalPasteboard"))
	return rv
}

func (p_ Pasteboard) PasteboardItems() []PasteboardItem {
	rv := objc.CallMethod[[]PasteboardItem](p_, objc.SEL("pasteboardItems"))
	return rv
}

func (p_ Pasteboard) Types() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](p_, objc.SEL("types"))
	return rv
}

func (p_ Pasteboard) Name() PasteboardName {
	rv := objc.CallMethod[PasteboardName](p_, objc.SEL("name"))
	return rv
}

func (p_ Pasteboard) ChangeCount() int {
	rv := objc.CallMethod[int](p_, objc.SEL("changeCount"))
	return rv
}
