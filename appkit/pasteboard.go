// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	SetString_ForType(_string string, dataType PasteboardType) bool
	IndexOfPasteboardItem(pasteboardItem IPasteboardItem) uint
	DataForType(dataType PasteboardType) []byte
	PropertyListForType(dataType PasteboardType) objc.Object
	StringForType(dataType PasteboardType) string
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	CanReadItemWithDataConformingToTypes(types []string) bool
	PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int
	DeclareTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int
	AddTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int
	WriteFileContents(filename string) bool
	WriteFileWrapper(wrapper foundation.IFileWrapper) bool
	ReadFileContentsType_ToFile(_type PasteboardType, filename string) string
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
	rv := ffi.CallMethod[Pasteboard](pc, "alloc")
	return rv
}

func (p_ Pasteboard) Init() Pasteboard {
	rv := ffi.CallMethod[Pasteboard](p_, "init")
	return rv
}

func (pc _PasteboardClass) New() Pasteboard {
	rv := ffi.CallMethod[Pasteboard](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPasteboard() Pasteboard {
	return PasteboardClass.New()
}

func (pc _PasteboardClass) PasteboardByFilteringData_OfType(data []byte, _type PasteboardType) Pasteboard {
	rv := ffi.CallMethod[Pasteboard](pc, "pasteboardByFilteringData:ofType:", data, _type)
	return rv
}

func (pc _PasteboardClass) PasteboardByFilteringFile(filename string) Pasteboard {
	rv := ffi.CallMethod[Pasteboard](pc, "pasteboardByFilteringFile:", filename)
	return rv
}

func (pc _PasteboardClass) PasteboardByFilteringTypesInPasteboard(pboard IPasteboard) Pasteboard {
	rv := ffi.CallMethod[Pasteboard](pc, "pasteboardByFilteringTypesInPasteboard:", pboard)
	return rv
}

func (pc _PasteboardClass) PasteboardWithName(name PasteboardName) Pasteboard {
	rv := ffi.CallMethod[Pasteboard](pc, "pasteboardWithName:", name)
	return rv
}

func (pc _PasteboardClass) PasteboardWithUniqueName() Pasteboard {
	rv := ffi.CallMethod[Pasteboard](pc, "pasteboardWithUniqueName")
	return rv
}

func (p_ Pasteboard) ClearContents() int {
	rv := ffi.CallMethod[int](p_, "clearContents")
	return rv
}

func (p_ Pasteboard) SetData_ForType(data []byte, dataType PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setData:forType:", data, dataType)
	return rv
}

func (p_ Pasteboard) SetPropertyList_ForType(plist objc.IObject, dataType PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setPropertyList:forType:", plist, dataType)
	return rv
}

func (p_ Pasteboard) SetString_ForType(_string string, dataType PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setString:forType:", _string, dataType)
	return rv
}

func (p_ Pasteboard) IndexOfPasteboardItem(pasteboardItem IPasteboardItem) uint {
	rv := ffi.CallMethod[uint](p_, "indexOfPasteboardItem:", pasteboardItem)
	return rv
}

func (p_ Pasteboard) DataForType(dataType PasteboardType) []byte {
	rv := ffi.CallMethod[[]byte](p_, "dataForType:", dataType)
	return rv
}

func (p_ Pasteboard) PropertyListForType(dataType PasteboardType) objc.Object {
	rv := ffi.CallMethod[objc.Object](p_, "propertyListForType:", dataType)
	return rv
}

func (p_ Pasteboard) StringForType(dataType PasteboardType) string {
	rv := ffi.CallMethod[string](p_, "stringForType:", dataType)
	return rv
}

func (p_ Pasteboard) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	rv := ffi.CallMethod[PasteboardType](p_, "availableTypeFromArray:", types)
	return rv
}

func (p_ Pasteboard) CanReadItemWithDataConformingToTypes(types []string) bool {
	rv := ffi.CallMethod[bool](p_, "canReadItemWithDataConformingToTypes:", types)
	return rv
}

func (pc _PasteboardClass) TypesFilterableTo(_type PasteboardType) []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](pc, "typesFilterableTo:", _type)
	return rv
}

func (p_ Pasteboard) PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int {
	rv := ffi.CallMethod[int](p_, "prepareForNewContentsWithOptions:", options)
	return rv
}

func (p_ Pasteboard) DeclareTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int {
	rv := ffi.CallMethod[int](p_, "declareTypes:owner:", newTypes, newOwner)
	return rv
}

func (p_ Pasteboard) AddTypes_Owner(newTypes []PasteboardType, newOwner objc.IObject) int {
	rv := ffi.CallMethod[int](p_, "addTypes:owner:", newTypes, newOwner)
	return rv
}

func (p_ Pasteboard) WriteFileContents(filename string) bool {
	rv := ffi.CallMethod[bool](p_, "writeFileContents:", filename)
	return rv
}

func (p_ Pasteboard) WriteFileWrapper(wrapper foundation.IFileWrapper) bool {
	rv := ffi.CallMethod[bool](p_, "writeFileWrapper:", wrapper)
	return rv
}

func (p_ Pasteboard) ReadFileContentsType_ToFile(_type PasteboardType, filename string) string {
	rv := ffi.CallMethod[string](p_, "readFileContentsType:toFile:", _type, filename)
	return rv
}

func (p_ Pasteboard) ReadFileWrapper() foundation.FileWrapper {
	rv := ffi.CallMethod[foundation.FileWrapper](p_, "readFileWrapper")
	return rv
}

func (pc _PasteboardClass) GeneralPasteboard() Pasteboard {
	rv := ffi.CallMethod[Pasteboard](pc, "generalPasteboard")
	return rv
}

func (p_ Pasteboard) PasteboardItems() []PasteboardItem {
	rv := ffi.CallMethod[[]PasteboardItem](p_, "pasteboardItems")
	return rv
}

func (p_ Pasteboard) Types() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](p_, "types")
	return rv
}

func (p_ Pasteboard) Name() PasteboardName {
	rv := ffi.CallMethod[PasteboardName](p_, "name")
	return rv
}

func (p_ Pasteboard) ChangeCount() int {
	rv := ffi.CallMethod[int](p_, "changeCount")
	return rv
}

var PasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}

type _PasteboardItemClass struct {
	objc.Class
}

type IPasteboardItem interface {
	objc.IObject
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	SetDataProvider_ForTypes(dataProvider PasteboardItemDataProvider, types []PasteboardType) bool
	SetData_ForType(data []byte, _type PasteboardType) bool
	SetString_ForType(_string string, _type PasteboardType) bool
	SetPropertyList_ForType(propertyList objc.IObject, _type PasteboardType) bool
	DataForType(_type PasteboardType) []byte
	StringForType(_type PasteboardType) string
	PropertyListForType(_type PasteboardType) objc.Object
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

func (p_ PasteboardItem) Init() PasteboardItem {
	rv := ffi.CallMethod[PasteboardItem](p_, "init")
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

func (p_ PasteboardItem) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	rv := ffi.CallMethod[PasteboardType](p_, "availableTypeFromArray:", types)
	return rv
}

func (p_ PasteboardItem) SetDataProvider_ForTypes(dataProvider PasteboardItemDataProvider, types []PasteboardType) bool {
	po := ffi.CreateProtocol(dataProvider)
	defer po.Release()
	rv := ffi.CallMethod[bool](p_, "setDataProvider:forTypes:", po, types)
	return rv
}

func (p_ PasteboardItem) SetData_ForType(data []byte, _type PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setData:forType:", data, _type)
	return rv
}

func (p_ PasteboardItem) SetString_ForType(_string string, _type PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setString:forType:", _string, _type)
	return rv
}

func (p_ PasteboardItem) SetPropertyList_ForType(propertyList objc.IObject, _type PasteboardType) bool {
	rv := ffi.CallMethod[bool](p_, "setPropertyList:forType:", propertyList, _type)
	return rv
}

func (p_ PasteboardItem) DataForType(_type PasteboardType) []byte {
	rv := ffi.CallMethod[[]byte](p_, "dataForType:", _type)
	return rv
}

func (p_ PasteboardItem) StringForType(_type PasteboardType) string {
	rv := ffi.CallMethod[string](p_, "stringForType:", _type)
	return rv
}

func (p_ PasteboardItem) PropertyListForType(_type PasteboardType) objc.Object {
	rv := ffi.CallMethod[objc.Object](p_, "propertyListForType:", _type)
	return rv
}

func (p_ PasteboardItem) Types() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](p_, "types")
	return rv
}

type PasteboardItemDataProvider interface {
	// required
	Pasteboard_Item_ProvideDataForType(pasteboard Pasteboard, item PasteboardItem, _type PasteboardType)
	ImplementsPasteboardFinishedWithDataProvider() bool
	// optional
	PasteboardFinishedWithDataProvider(pasteboard Pasteboard)
}

type PasteboardItemDataProviderWrapper struct {
	objc.Object
}

func (p_ PasteboardItemDataProviderWrapper) Pasteboard_Item_ProvideDataForType(pasteboard IPasteboard, item IPasteboardItem, _type PasteboardType) {
	ffi.CallMethod[ffi.Void](p_, "pasteboard:item:provideDataForType:", pasteboard, item, _type)
}

func (p_ *PasteboardItemDataProviderWrapper) ImplementsPasteboardFinishedWithDataProvider() bool {
	return p_.RespondsToSelector(objc.GetSelector("pasteboardFinishedWithDataProvider:"))
}

func (p_ PasteboardItemDataProviderWrapper) PasteboardFinishedWithDataProvider(pasteboard IPasteboard) {
	ffi.CallMethod[ffi.Void](p_, "pasteboardFinishedWithDataProvider:", pasteboard)
}

type PasteboardWriting interface {
	// required
	WritableTypesForPasteboard(pasteboard Pasteboard) []PasteboardType
	ImplementsWritingOptionsForType_Pasteboard() bool
	// optional
	WritingOptionsForType_Pasteboard(_type PasteboardType, pasteboard Pasteboard) PasteboardWritingOptions
	// required
	PasteboardPropertyListForType(_type PasteboardType) objc.IObject
}

type PasteboardWritingWrapper struct {
	objc.Object
}

func (p_ PasteboardWritingWrapper) WritableTypesForPasteboard(pasteboard IPasteboard) []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](p_, "writableTypesForPasteboard:", pasteboard)
	return rv
}

func (p_ *PasteboardWritingWrapper) ImplementsWritingOptionsForType_Pasteboard() bool {
	return p_.RespondsToSelector(objc.GetSelector("writingOptionsForType:pasteboard:"))
}

func (p_ PasteboardWritingWrapper) WritingOptionsForType_Pasteboard(_type PasteboardType, pasteboard IPasteboard) PasteboardWritingOptions {
	rv := ffi.CallMethod[PasteboardWritingOptions](p_, "writingOptionsForType:pasteboard:", _type, pasteboard)
	return rv
}

func (p_ PasteboardWritingWrapper) PasteboardPropertyListForType(_type PasteboardType) objc.Object {
	rv := ffi.CallMethod[objc.Object](p_, "pasteboardPropertyListForType:", _type)
	return rv
}
