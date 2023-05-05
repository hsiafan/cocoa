// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var FontManagerClass = _FontManagerClass{objc.GetClass("NSFontManager")}

type _FontManagerClass struct {
	objc.Class
}

type IFontManager interface {
	objc.IObject
	AvailableFontNamesWithTraits(someTraits FontTraitMask) []string
	AvailableMembersOfFontFamily(fam string) [][]objc.Object
	SetSelectedFont_IsMultiple(fontObj IFont, flag bool)
	SendAction() bool
	LocalizedNameForFamily_Face(family string, faceKey string) string
	AddFontTrait(sender objc.IObject)
	RemoveFontTrait(sender objc.IObject)
	ModifyFont(sender objc.IObject)
	ModifyFontViaPanel(sender objc.IObject)
	OrderFrontStylesPanel(sender objc.IObject)
	OrderFrontFontPanel(sender objc.IObject)
	ConvertFont(fontObj IFont) Font
	// deprecated
	ChangeFont(sender objc.IObject)
	ConvertFont_ToFace(fontObj IFont, typeface string) Font
	ConvertFont_ToFamily(fontObj IFont, family string) Font
	ConvertFont_ToHaveTrait(fontObj IFont, trait FontTraitMask) Font
	ConvertFont_ToNotHaveTrait(fontObj IFont, trait FontTraitMask) Font
	ConvertFont_ToSize(fontObj IFont, size float64) Font
	ConvertWeight_OfFont(upFlag bool, fontObj IFont) Font
	ConvertFontTraits(traits FontTraitMask) FontTraitMask
	FontWithFamily_Traits_Weight_Size(family string, traits FontTraitMask, weight int, size float64) Font
	TraitsOfFont(fontObj IFont) FontTraitMask
	FontNamed_HasTraits(fName string, someTraits FontTraitMask) bool
	WeightOfFont(fontObj IFont) int
	FontPanel(create bool) FontPanel
	SetFontMenu(newMenu IMenu)
	FontMenu(create bool) Menu
	SetSelectedAttributes_IsMultiple(attributes map[string]objc.IObject, flag bool)
	ConvertAttributes(attributes map[string]objc.IObject) map[string]objc.Object
	// deprecated
	AvailableFontNamesMatchingFontDescriptor(descriptor IFontDescriptor) []objc.Object
	// deprecated
	FontDescriptorsInCollection(collectionNames string) []objc.Object
	// deprecated
	AddCollection_Options(collectionName string, collectionOptions FontCollectionOptions) bool
	// deprecated
	RemoveCollection(collectionName string) bool
	// deprecated
	AddFontDescriptors_ToCollection(descriptors []objc.IObject, collectionName string)
	// deprecated
	RemoveFontDescriptor_FromCollection(descriptor IFontDescriptor, collection string)
	// deprecated
	FontManager_WillIncludeFont(sender objc.IObject, fontName string) bool
	AvailableFonts() []string
	AvailableFontFamilies() []string
	SelectedFont() Font
	IsMultiple() bool
	CurrentFontAction() FontAction
	IsEnabled() bool
	SetEnabled(value bool)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.IObject)
	// deprecated
	CollectionNames() []objc.Object
	// deprecated
	Delegate() objc.Object
	// deprecated
	SetDelegate(value objc.IObject)
}

type FontManager struct {
	objc.Object
}

func MakeFontManager(ptr unsafe.Pointer) FontManager {
	return FontManager{
		Object: objc.MakeObject(ptr),
	}
}

func (fc _FontManagerClass) Alloc() FontManager {
	rv := objc.CallMethod[FontManager](fc, objc.GetSelector("alloc"))
	return rv
}

func (fc _FontManagerClass) New() FontManager {
	rv := objc.CallMethod[FontManager](fc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewFontManager() FontManager {
	return FontManagerClass.New()
}

func (f_ FontManager) Init() FontManager {
	rv := objc.CallMethod[FontManager](f_, objc.GetSelector("init"))
	return rv
}

func (fc _FontManagerClass) SetFontManagerFactory(factoryId objc.IClass) {
	objc.CallMethod[objc.Void](fc, objc.GetSelector("setFontManagerFactory:"), objc.ExtractPtr(factoryId))
}

func (fc _FontManagerClass) SetFontPanelFactory(factoryId objc.IClass) {
	objc.CallMethod[objc.Void](fc, objc.GetSelector("setFontPanelFactory:"), objc.ExtractPtr(factoryId))
}

func (f_ FontManager) AvailableFontNamesWithTraits(someTraits FontTraitMask) []string {
	rv := objc.CallMethod[[]string](f_, objc.GetSelector("availableFontNamesWithTraits:"), someTraits)
	return rv
}

func (f_ FontManager) AvailableMembersOfFontFamily(fam string) [][]objc.Object {
	rv := objc.CallMethod[[][]objc.Object](f_, objc.GetSelector("availableMembersOfFontFamily:"), fam)
	return rv
}

func (f_ FontManager) SetSelectedFont_IsMultiple(fontObj IFont, flag bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setSelectedFont:isMultiple:"), objc.ExtractPtr(fontObj), flag)
}

func (f_ FontManager) SendAction() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("sendAction"))
	return rv
}

func (f_ FontManager) LocalizedNameForFamily_Face(family string, faceKey string) string {
	rv := objc.CallMethod[string](f_, objc.GetSelector("localizedNameForFamily:face:"), family, faceKey)
	return rv
}

func (f_ FontManager) AddFontTrait(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("addFontTrait:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) RemoveFontTrait(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("removeFontTrait:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) ModifyFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("modifyFont:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) ModifyFontViaPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("modifyFontViaPanel:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) OrderFrontStylesPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("orderFrontStylesPanel:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) OrderFrontFontPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("orderFrontFontPanel:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) ConvertFont(fontObj IFont) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:"), objc.ExtractPtr(fontObj))
	return rv
}

// deprecated
func (f_ FontManager) ChangeFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("changeFont:"), objc.ExtractPtr(sender))
}

func (f_ FontManager) ConvertFont_ToFace(fontObj IFont, typeface string) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toFace:"), objc.ExtractPtr(fontObj), typeface)
	return rv
}

func (f_ FontManager) ConvertFont_ToFamily(fontObj IFont, family string) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toFamily:"), objc.ExtractPtr(fontObj), family)
	return rv
}

func (f_ FontManager) ConvertFont_ToHaveTrait(fontObj IFont, trait FontTraitMask) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toHaveTrait:"), objc.ExtractPtr(fontObj), trait)
	return rv
}

func (f_ FontManager) ConvertFont_ToNotHaveTrait(fontObj IFont, trait FontTraitMask) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toNotHaveTrait:"), objc.ExtractPtr(fontObj), trait)
	return rv
}

func (f_ FontManager) ConvertFont_ToSize(fontObj IFont, size float64) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertFont:toSize:"), objc.ExtractPtr(fontObj), size)
	return rv
}

func (f_ FontManager) ConvertWeight_OfFont(upFlag bool, fontObj IFont) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("convertWeight:ofFont:"), upFlag, objc.ExtractPtr(fontObj))
	return rv
}

func (f_ FontManager) ConvertFontTraits(traits FontTraitMask) FontTraitMask {
	rv := objc.CallMethod[FontTraitMask](f_, objc.GetSelector("convertFontTraits:"), traits)
	return rv
}

func (f_ FontManager) FontWithFamily_Traits_Weight_Size(family string, traits FontTraitMask, weight int, size float64) Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("fontWithFamily:traits:weight:size:"), family, traits, weight, size)
	return rv
}

func (f_ FontManager) TraitsOfFont(fontObj IFont) FontTraitMask {
	rv := objc.CallMethod[FontTraitMask](f_, objc.GetSelector("traitsOfFont:"), objc.ExtractPtr(fontObj))
	return rv
}

func (f_ FontManager) FontNamed_HasTraits(fName string, someTraits FontTraitMask) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("fontNamed:hasTraits:"), fName, someTraits)
	return rv
}

func (f_ FontManager) WeightOfFont(fontObj IFont) int {
	rv := objc.CallMethod[int](f_, objc.GetSelector("weightOfFont:"), objc.ExtractPtr(fontObj))
	return rv
}

func (f_ FontManager) FontPanel(create bool) FontPanel {
	rv := objc.CallMethod[FontPanel](f_, objc.GetSelector("fontPanel:"), create)
	return rv
}

func (f_ FontManager) SetFontMenu(newMenu IMenu) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setFontMenu:"), objc.ExtractPtr(newMenu))
}

func (f_ FontManager) FontMenu(create bool) Menu {
	rv := objc.CallMethod[Menu](f_, objc.GetSelector("fontMenu:"), create)
	return rv
}

func (f_ FontManager) SetSelectedAttributes_IsMultiple(attributes map[string]objc.IObject, flag bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setSelectedAttributes:isMultiple:"), attributes, flag)
}

func (f_ FontManager) ConvertAttributes(attributes map[string]objc.IObject) map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](f_, objc.GetSelector("convertAttributes:"), attributes)
	return rv
}

// deprecated
func (f_ FontManager) AvailableFontNamesMatchingFontDescriptor(descriptor IFontDescriptor) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](f_, objc.GetSelector("availableFontNamesMatchingFontDescriptor:"), objc.ExtractPtr(descriptor))
	return rv
}

// deprecated
func (f_ FontManager) FontDescriptorsInCollection(collectionNames string) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](f_, objc.GetSelector("fontDescriptorsInCollection:"), collectionNames)
	return rv
}

// deprecated
func (f_ FontManager) AddCollection_Options(collectionName string, collectionOptions FontCollectionOptions) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("addCollection:options:"), collectionName, collectionOptions)
	return rv
}

// deprecated
func (f_ FontManager) RemoveCollection(collectionName string) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("removeCollection:"), collectionName)
	return rv
}

// deprecated
func (f_ FontManager) AddFontDescriptors_ToCollection(descriptors []objc.IObject, collectionName string) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("addFontDescriptors:toCollection:"), descriptors, collectionName)
}

// deprecated
func (f_ FontManager) RemoveFontDescriptor_FromCollection(descriptor IFontDescriptor, collection string) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("removeFontDescriptor:fromCollection:"), objc.ExtractPtr(descriptor), collection)
}

// deprecated
func (f_ FontManager) FontManager_WillIncludeFont(sender objc.IObject, fontName string) bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("fontManager:willIncludeFont:"), objc.ExtractPtr(sender), fontName)
	return rv
}

func (fc _FontManagerClass) SharedFontManager() FontManager {
	rv := objc.CallMethod[FontManager](fc, objc.GetSelector("sharedFontManager"))
	return rv
}

func (f_ FontManager) AvailableFonts() []string {
	rv := objc.CallMethod[[]string](f_, objc.GetSelector("availableFonts"))
	return rv
}

func (f_ FontManager) AvailableFontFamilies() []string {
	rv := objc.CallMethod[[]string](f_, objc.GetSelector("availableFontFamilies"))
	return rv
}

func (f_ FontManager) SelectedFont() Font {
	rv := objc.CallMethod[Font](f_, objc.GetSelector("selectedFont"))
	return rv
}

func (f_ FontManager) IsMultiple() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isMultiple"))
	return rv
}

func (f_ FontManager) CurrentFontAction() FontAction {
	rv := objc.CallMethod[FontAction](f_, objc.GetSelector("currentFontAction"))
	return rv
}

func (f_ FontManager) IsEnabled() bool {
	rv := objc.CallMethod[bool](f_, objc.GetSelector("isEnabled"))
	return rv
}

func (f_ FontManager) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setEnabled:"), value)
}

func (f_ FontManager) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](f_, objc.GetSelector("action"))
	return rv
}

func (f_ FontManager) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setAction:"), value)
}

func (f_ FontManager) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](f_, objc.GetSelector("target"))
	return rv
}

func (f_ FontManager) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setTarget:"), objc.ExtractPtr(value))
}

// deprecated
func (f_ FontManager) CollectionNames() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](f_, objc.GetSelector("collectionNames"))
	return rv
}

// deprecated
func (f_ FontManager) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](f_, objc.GetSelector("delegate"))
	return rv
}

// deprecated
func (f_ FontManager) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}
