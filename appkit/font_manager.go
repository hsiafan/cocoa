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
	rv := objc.CallMethod[FontManager](fc, "alloc")
	return rv
}

func (fc _FontManagerClass) New() FontManager {
	rv := objc.CallMethod[FontManager](fc, "new")
	rv.Autorelease()
	return rv
}

func NewFontManager() FontManager {
	return FontManagerClass.New()
}

func (f_ FontManager) Init() FontManager {
	rv := objc.CallMethod[FontManager](f_, "init")
	return rv
}

func (fc _FontManagerClass) SetFontManagerFactory(factoryId objc.IClass) {
	objc.CallMethod[objc.Void](fc, "setFontManagerFactory:", factoryId)
}

func (fc _FontManagerClass) SetFontPanelFactory(factoryId objc.IClass) {
	objc.CallMethod[objc.Void](fc, "setFontPanelFactory:", factoryId)
}

func (f_ FontManager) AvailableFontNamesWithTraits(someTraits FontTraitMask) []string {
	rv := objc.CallMethod[[]string](f_, "availableFontNamesWithTraits:", someTraits)
	return rv
}

func (f_ FontManager) AvailableMembersOfFontFamily(fam string) [][]objc.Object {
	rv := objc.CallMethod[[][]objc.Object](f_, "availableMembersOfFontFamily:", fam)
	return rv
}

func (f_ FontManager) SetSelectedFont_IsMultiple(fontObj IFont, flag bool) {
	objc.CallMethod[objc.Void](f_, "setSelectedFont:isMultiple:", fontObj, flag)
}

func (f_ FontManager) SendAction() bool {
	rv := objc.CallMethod[bool](f_, "sendAction")
	return rv
}

func (f_ FontManager) LocalizedNameForFamily_Face(family string, faceKey string) string {
	rv := objc.CallMethod[string](f_, "localizedNameForFamily:face:", family, faceKey)
	return rv
}

func (f_ FontManager) AddFontTrait(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, "addFontTrait:", sender)
}

func (f_ FontManager) RemoveFontTrait(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, "removeFontTrait:", sender)
}

func (f_ FontManager) ModifyFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, "modifyFont:", sender)
}

func (f_ FontManager) ModifyFontViaPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, "modifyFontViaPanel:", sender)
}

func (f_ FontManager) OrderFrontStylesPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, "orderFrontStylesPanel:", sender)
}

func (f_ FontManager) OrderFrontFontPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, "orderFrontFontPanel:", sender)
}

func (f_ FontManager) ConvertFont(fontObj IFont) Font {
	rv := objc.CallMethod[Font](f_, "convertFont:", fontObj)
	return rv
}

// deprecated
func (f_ FontManager) ChangeFont(sender objc.IObject) {
	objc.CallMethod[objc.Void](f_, "changeFont:", sender)
}

func (f_ FontManager) ConvertFont_ToFace(fontObj IFont, typeface string) Font {
	rv := objc.CallMethod[Font](f_, "convertFont:toFace:", fontObj, typeface)
	return rv
}

func (f_ FontManager) ConvertFont_ToFamily(fontObj IFont, family string) Font {
	rv := objc.CallMethod[Font](f_, "convertFont:toFamily:", fontObj, family)
	return rv
}

func (f_ FontManager) ConvertFont_ToHaveTrait(fontObj IFont, trait FontTraitMask) Font {
	rv := objc.CallMethod[Font](f_, "convertFont:toHaveTrait:", fontObj, trait)
	return rv
}

func (f_ FontManager) ConvertFont_ToNotHaveTrait(fontObj IFont, trait FontTraitMask) Font {
	rv := objc.CallMethod[Font](f_, "convertFont:toNotHaveTrait:", fontObj, trait)
	return rv
}

func (f_ FontManager) ConvertFont_ToSize(fontObj IFont, size float64) Font {
	rv := objc.CallMethod[Font](f_, "convertFont:toSize:", fontObj, size)
	return rv
}

func (f_ FontManager) ConvertWeight_OfFont(upFlag bool, fontObj IFont) Font {
	rv := objc.CallMethod[Font](f_, "convertWeight:ofFont:", upFlag, fontObj)
	return rv
}

func (f_ FontManager) ConvertFontTraits(traits FontTraitMask) FontTraitMask {
	rv := objc.CallMethod[FontTraitMask](f_, "convertFontTraits:", traits)
	return rv
}

func (f_ FontManager) FontWithFamily_Traits_Weight_Size(family string, traits FontTraitMask, weight int, size float64) Font {
	rv := objc.CallMethod[Font](f_, "fontWithFamily:traits:weight:size:", family, traits, weight, size)
	return rv
}

func (f_ FontManager) TraitsOfFont(fontObj IFont) FontTraitMask {
	rv := objc.CallMethod[FontTraitMask](f_, "traitsOfFont:", fontObj)
	return rv
}

func (f_ FontManager) FontNamed_HasTraits(fName string, someTraits FontTraitMask) bool {
	rv := objc.CallMethod[bool](f_, "fontNamed:hasTraits:", fName, someTraits)
	return rv
}

func (f_ FontManager) WeightOfFont(fontObj IFont) int {
	rv := objc.CallMethod[int](f_, "weightOfFont:", fontObj)
	return rv
}

func (f_ FontManager) FontPanel(create bool) FontPanel {
	rv := objc.CallMethod[FontPanel](f_, "fontPanel:", create)
	return rv
}

func (f_ FontManager) SetFontMenu(newMenu IMenu) {
	objc.CallMethod[objc.Void](f_, "setFontMenu:", newMenu)
}

func (f_ FontManager) FontMenu(create bool) Menu {
	rv := objc.CallMethod[Menu](f_, "fontMenu:", create)
	return rv
}

func (f_ FontManager) SetSelectedAttributes_IsMultiple(attributes map[string]objc.IObject, flag bool) {
	objc.CallMethod[objc.Void](f_, "setSelectedAttributes:isMultiple:", attributes, flag)
}

func (f_ FontManager) ConvertAttributes(attributes map[string]objc.IObject) map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](f_, "convertAttributes:", attributes)
	return rv
}

// deprecated
func (f_ FontManager) AvailableFontNamesMatchingFontDescriptor(descriptor IFontDescriptor) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](f_, "availableFontNamesMatchingFontDescriptor:", descriptor)
	return rv
}

// deprecated
func (f_ FontManager) FontDescriptorsInCollection(collectionNames string) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](f_, "fontDescriptorsInCollection:", collectionNames)
	return rv
}

// deprecated
func (f_ FontManager) AddCollection_Options(collectionName string, collectionOptions FontCollectionOptions) bool {
	rv := objc.CallMethod[bool](f_, "addCollection:options:", collectionName, collectionOptions)
	return rv
}

// deprecated
func (f_ FontManager) RemoveCollection(collectionName string) bool {
	rv := objc.CallMethod[bool](f_, "removeCollection:", collectionName)
	return rv
}

// deprecated
func (f_ FontManager) AddFontDescriptors_ToCollection(descriptors []objc.IObject, collectionName string) {
	objc.CallMethod[objc.Void](f_, "addFontDescriptors:toCollection:", descriptors, collectionName)
}

// deprecated
func (f_ FontManager) RemoveFontDescriptor_FromCollection(descriptor IFontDescriptor, collection string) {
	objc.CallMethod[objc.Void](f_, "removeFontDescriptor:fromCollection:", descriptor, collection)
}

// deprecated
func (f_ FontManager) FontManager_WillIncludeFont(sender objc.IObject, fontName string) bool {
	rv := objc.CallMethod[bool](f_, "fontManager:willIncludeFont:", sender, fontName)
	return rv
}

func (fc _FontManagerClass) SharedFontManager() FontManager {
	rv := objc.CallMethod[FontManager](fc, "sharedFontManager")
	return rv
}

func (f_ FontManager) AvailableFonts() []string {
	rv := objc.CallMethod[[]string](f_, "availableFonts")
	return rv
}

func (f_ FontManager) AvailableFontFamilies() []string {
	rv := objc.CallMethod[[]string](f_, "availableFontFamilies")
	return rv
}

func (f_ FontManager) SelectedFont() Font {
	rv := objc.CallMethod[Font](f_, "selectedFont")
	return rv
}

func (f_ FontManager) IsMultiple() bool {
	rv := objc.CallMethod[bool](f_, "isMultiple")
	return rv
}

func (f_ FontManager) CurrentFontAction() FontAction {
	rv := objc.CallMethod[FontAction](f_, "currentFontAction")
	return rv
}

func (f_ FontManager) IsEnabled() bool {
	rv := objc.CallMethod[bool](f_, "isEnabled")
	return rv
}

func (f_ FontManager) SetEnabled(value bool) {
	objc.CallMethod[objc.Void](f_, "setEnabled:", value)
}

func (f_ FontManager) Action() objc.Selector {
	rv := objc.CallMethod[objc.Selector](f_, "action")
	return rv
}

func (f_ FontManager) SetAction(value objc.Selector) {
	objc.CallMethod[objc.Void](f_, "setAction:", value)
}

func (f_ FontManager) Target() objc.Object {
	rv := objc.CallMethod[objc.Object](f_, "target")
	return rv
}

func (f_ FontManager) SetTarget(value objc.IObject) {
	objc.CallMethod[objc.Void](f_, "setTarget:", value)
}

// deprecated
func (f_ FontManager) CollectionNames() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](f_, "collectionNames")
	return rv
}

// deprecated
func (f_ FontManager) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](f_, "delegate")
	return rv
}

// deprecated
func (f_ FontManager) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](f_, "setDelegate:", value)
}
