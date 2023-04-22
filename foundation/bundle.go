// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var BundleClass = _BundleClass{objc.GetClass("NSBundle")}

type _BundleClass struct {
	objc.Class
}

type IBundle interface {
	objc.IObject
	URLForResource_WithExtension_Subdirectory(name string, ext string, subpath string) URL
	URLForResource_WithExtension(name string, ext string) URL
	URLsForResourcesWithExtension_Subdirectory(ext string, subpath string) []URL
	URLForResource_WithExtension_Subdirectory_Localization(name string, ext string, subpath string, localizationName string) URL
	URLsForResourcesWithExtension_Subdirectory_Localization(ext string, subpath string, localizationName string) []URL
	PathForResource_OfType(name string, ext string) string
	PathForResource_OfType_InDirectory(name string, ext string, subpath string) string
	PathForResource_OfType_InDirectory_ForLocalization(name string, ext string, subpath string, localizationName string) string
	PathsForResourcesOfType_InDirectory(ext string, subpath string) []string
	PathsForResourcesOfType_InDirectory_ForLocalization(ext string, subpath string, localizationName string) []string
	LocalizedStringForKey_Value_Table(key string, value string, tableName string) string
	URLForAuxiliaryExecutable(executableName string) URL
	PathForAuxiliaryExecutable(executableName string) string
	ObjectForInfoDictionaryKey(key string) objc.Object
	ClassNamed(className string) objc.Class
	PreflightAndReturnError(error *Error) bool
	Load() bool
	LoadAndReturnError(error *Error) bool
	Unload() bool
	LocalizedAttributedStringForKey_Value_Table(key string, value string, tableName string) AttributedString
	ResourceURL() URL
	ExecutableURL() URL
	PrivateFrameworksURL() URL
	SharedFrameworksURL() URL
	BuiltInPlugInsURL() URL
	SharedSupportURL() URL
	AppStoreReceiptURL() URL
	ResourcePath() string
	ExecutablePath() string
	PrivateFrameworksPath() string
	SharedFrameworksPath() string
	BuiltInPlugInsPath() string
	SharedSupportPath() string
	BundleURL() URL
	BundlePath() string
	BundleIdentifier() string
	InfoDictionary() map[string]objc.Object
	Localizations() []string
	PreferredLocalizations() []string
	DevelopmentLocalization() string
	LocalizedInfoDictionary() map[string]objc.Object
	PrincipalClass() objc.Class
	ExecutableArchitectures() []Number
	IsLoaded() bool
}

type Bundle struct {
	objc.Object
}

func MakeBundle(ptr unsafe.Pointer) Bundle {
	return Bundle{
		Object: objc.MakeObject(ptr),
	}
}

func (bc _BundleClass) BundleWithURL(url IURL) Bundle {
	rv := objc.CallMethod[Bundle](bc, objc.GetSelector("bundleWithURL:"), url)
	return rv
}

func (bc _BundleClass) BundleWithPath(path string) Bundle {
	rv := objc.CallMethod[Bundle](bc, objc.GetSelector("bundleWithPath:"), path)
	return rv
}

func (b_ Bundle) InitWithURL(url IURL) Bundle {
	rv := objc.CallMethod[Bundle](b_, objc.GetSelector("initWithURL:"), url)
	return rv
}

func (b_ Bundle) InitWithPath(path string) Bundle {
	rv := objc.CallMethod[Bundle](b_, objc.GetSelector("initWithPath:"), path)
	return rv
}

func (bc _BundleClass) Alloc() Bundle {
	rv := objc.CallMethod[Bundle](bc, objc.GetSelector("alloc"))
	return rv
}

func (bc _BundleClass) New() Bundle {
	rv := objc.CallMethod[Bundle](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBundle() Bundle {
	return BundleClass.New()
}

func (b_ Bundle) Init() Bundle {
	rv := objc.CallMethod[Bundle](b_, objc.GetSelector("init"))
	return rv
}

func (bc _BundleClass) BundleForClass(aClass objc.IClass) Bundle {
	rv := objc.CallMethod[Bundle](bc, objc.GetSelector("bundleForClass:"), aClass)
	return rv
}

func (bc _BundleClass) BundleWithIdentifier(identifier string) Bundle {
	rv := objc.CallMethod[Bundle](bc, objc.GetSelector("bundleWithIdentifier:"), identifier)
	return rv
}

func (b_ Bundle) URLForResource_WithExtension_Subdirectory(name string, ext string, subpath string) URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("URLForResource:withExtension:subdirectory:"), name, ext, subpath)
	return rv
}

func (b_ Bundle) URLForResource_WithExtension(name string, ext string) URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("URLForResource:withExtension:"), name, ext)
	return rv
}

func (b_ Bundle) URLsForResourcesWithExtension_Subdirectory(ext string, subpath string) []URL {
	rv := objc.CallMethod[[]URL](b_, objc.GetSelector("URLsForResourcesWithExtension:subdirectory:"), ext, subpath)
	return rv
}

func (b_ Bundle) URLForResource_WithExtension_Subdirectory_Localization(name string, ext string, subpath string, localizationName string) URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("URLForResource:withExtension:subdirectory:localization:"), name, ext, subpath, localizationName)
	return rv
}

func (b_ Bundle) URLsForResourcesWithExtension_Subdirectory_Localization(ext string, subpath string, localizationName string) []URL {
	rv := objc.CallMethod[[]URL](b_, objc.GetSelector("URLsForResourcesWithExtension:subdirectory:localization:"), ext, subpath, localizationName)
	return rv
}

func (bc _BundleClass) URLForResource_WithExtension_Subdirectory_InBundleWithURL(name string, ext string, subpath string, bundleURL IURL) URL {
	rv := objc.CallMethod[URL](bc, objc.GetSelector("URLForResource:withExtension:subdirectory:inBundleWithURL:"), name, ext, subpath, bundleURL)
	return rv
}

func (bc _BundleClass) URLsForResourcesWithExtension_Subdirectory_InBundleWithURL(ext string, subpath string, bundleURL IURL) []URL {
	rv := objc.CallMethod[[]URL](bc, objc.GetSelector("URLsForResourcesWithExtension:subdirectory:inBundleWithURL:"), ext, subpath, bundleURL)
	return rv
}

func (b_ Bundle) PathForResource_OfType(name string, ext string) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("pathForResource:ofType:"), name, ext)
	return rv
}

func (b_ Bundle) PathForResource_OfType_InDirectory(name string, ext string, subpath string) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("pathForResource:ofType:inDirectory:"), name, ext, subpath)
	return rv
}

func (b_ Bundle) PathForResource_OfType_InDirectory_ForLocalization(name string, ext string, subpath string, localizationName string) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("pathForResource:ofType:inDirectory:forLocalization:"), name, ext, subpath, localizationName)
	return rv
}

func (b_ Bundle) PathsForResourcesOfType_InDirectory(ext string, subpath string) []string {
	rv := objc.CallMethod[[]string](b_, objc.GetSelector("pathsForResourcesOfType:inDirectory:"), ext, subpath)
	return rv
}

func (b_ Bundle) PathsForResourcesOfType_InDirectory_ForLocalization(ext string, subpath string, localizationName string) []string {
	rv := objc.CallMethod[[]string](b_, objc.GetSelector("pathsForResourcesOfType:inDirectory:forLocalization:"), ext, subpath, localizationName)
	return rv
}

func (b_ Bundle) LocalizedStringForKey_Value_Table(key string, value string, tableName string) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("localizedStringForKey:value:table:"), key, value, tableName)
	return rv
}

func (b_ Bundle) URLForAuxiliaryExecutable(executableName string) URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("URLForAuxiliaryExecutable:"), executableName)
	return rv
}

func (b_ Bundle) PathForAuxiliaryExecutable(executableName string) string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("pathForAuxiliaryExecutable:"), executableName)
	return rv
}

func (b_ Bundle) ObjectForInfoDictionaryKey(key string) objc.Object {
	rv := objc.CallMethod[objc.Object](b_, objc.GetSelector("objectForInfoDictionaryKey:"), key)
	return rv
}

func (bc _BundleClass) PreferredLocalizationsFromArray(localizationsArray []string) []string {
	rv := objc.CallMethod[[]string](bc, objc.GetSelector("preferredLocalizationsFromArray:"), localizationsArray)
	return rv
}

func (bc _BundleClass) PreferredLocalizationsFromArray_ForPreferences(localizationsArray []string, preferencesArray []string) []string {
	rv := objc.CallMethod[[]string](bc, objc.GetSelector("preferredLocalizationsFromArray:forPreferences:"), localizationsArray, preferencesArray)
	return rv
}

func (b_ Bundle) ClassNamed(className string) objc.Class {
	rv := objc.CallMethod[objc.Class](b_, objc.GetSelector("classNamed:"), className)
	return rv
}

func (b_ Bundle) PreflightAndReturnError(error *Error) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("preflightAndReturnError:"), unsafe.Pointer(error))
	return rv
}

func (b_ Bundle) Load() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("load"))
	return rv
}

func (b_ Bundle) LoadAndReturnError(error *Error) bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("loadAndReturnError:"), unsafe.Pointer(error))
	return rv
}

func (b_ Bundle) Unload() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("unload"))
	return rv
}

func (b_ Bundle) LocalizedAttributedStringForKey_Value_Table(key string, value string, tableName string) AttributedString {
	rv := objc.CallMethod[AttributedString](b_, objc.GetSelector("localizedAttributedStringForKey:value:table:"), key, value, tableName)
	return rv
}

func (bc _BundleClass) MainBundle() Bundle {
	rv := objc.CallMethod[Bundle](bc, objc.GetSelector("mainBundle"))
	return rv
}

func (bc _BundleClass) AllFrameworks() []Bundle {
	rv := objc.CallMethod[[]Bundle](bc, objc.GetSelector("allFrameworks"))
	return rv
}

func (bc _BundleClass) AllBundles() []Bundle {
	rv := objc.CallMethod[[]Bundle](bc, objc.GetSelector("allBundles"))
	return rv
}

func (b_ Bundle) ResourceURL() URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("resourceURL"))
	return rv
}

func (b_ Bundle) ExecutableURL() URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("executableURL"))
	return rv
}

func (b_ Bundle) PrivateFrameworksURL() URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("privateFrameworksURL"))
	return rv
}

func (b_ Bundle) SharedFrameworksURL() URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("sharedFrameworksURL"))
	return rv
}

func (b_ Bundle) BuiltInPlugInsURL() URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("builtInPlugInsURL"))
	return rv
}

func (b_ Bundle) SharedSupportURL() URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("sharedSupportURL"))
	return rv
}

func (b_ Bundle) AppStoreReceiptURL() URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("appStoreReceiptURL"))
	return rv
}

func (b_ Bundle) ResourcePath() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("resourcePath"))
	return rv
}

func (b_ Bundle) ExecutablePath() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("executablePath"))
	return rv
}

func (b_ Bundle) PrivateFrameworksPath() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("privateFrameworksPath"))
	return rv
}

func (b_ Bundle) SharedFrameworksPath() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("sharedFrameworksPath"))
	return rv
}

func (b_ Bundle) BuiltInPlugInsPath() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("builtInPlugInsPath"))
	return rv
}

func (b_ Bundle) SharedSupportPath() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("sharedSupportPath"))
	return rv
}

func (b_ Bundle) BundleURL() URL {
	rv := objc.CallMethod[URL](b_, objc.GetSelector("bundleURL"))
	return rv
}

func (b_ Bundle) BundlePath() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("bundlePath"))
	return rv
}

func (b_ Bundle) BundleIdentifier() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("bundleIdentifier"))
	return rv
}

func (b_ Bundle) InfoDictionary() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](b_, objc.GetSelector("infoDictionary"))
	return rv
}

func (b_ Bundle) Localizations() []string {
	rv := objc.CallMethod[[]string](b_, objc.GetSelector("localizations"))
	return rv
}

func (b_ Bundle) PreferredLocalizations() []string {
	rv := objc.CallMethod[[]string](b_, objc.GetSelector("preferredLocalizations"))
	return rv
}

func (b_ Bundle) DevelopmentLocalization() string {
	rv := objc.CallMethod[string](b_, objc.GetSelector("developmentLocalization"))
	return rv
}

func (b_ Bundle) LocalizedInfoDictionary() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](b_, objc.GetSelector("localizedInfoDictionary"))
	return rv
}

func (b_ Bundle) PrincipalClass() objc.Class {
	rv := objc.CallMethod[objc.Class](b_, objc.GetSelector("principalClass"))
	return rv
}

func (b_ Bundle) ExecutableArchitectures() []Number {
	rv := objc.CallMethod[[]Number](b_, objc.GetSelector("executableArchitectures"))
	return rv
}

func (b_ Bundle) IsLoaded() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isLoaded"))
	return rv
}
