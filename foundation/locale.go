// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var LocaleClass = _LocaleClass{objc.GetClass("NSLocale")}

type _LocaleClass struct {
	objc.Class
}

type ILocale interface {
	objc.IObject
	LocalizedStringForLocaleIdentifier(localeIdentifier string) string
	LocalizedStringForCountryCode(countryCode string) string
	LocalizedStringForLanguageCode(languageCode string) string
	LocalizedStringForScriptCode(scriptCode string) string
	LocalizedStringForVariantCode(variantCode string) string
	LocalizedStringForCollationIdentifier(collationIdentifier string) string
	LocalizedStringForCollatorIdentifier(collatorIdentifier string) string
	LocalizedStringForCurrencyCode(currencyCode string) string
	LocalizedStringForCalendarIdentifier(calendarIdentifier string) string
	ObjectForKey(key LocaleKey) objc.Object
	DisplayNameForKey_Value(key LocaleKey, value objc.IObject) string
	LocaleIdentifier() string
	CountryCode() string
	LanguageCode() string
	ScriptCode() string
	VariantCode() string
	ExemplarCharacterSet() CharacterSet
	CollationIdentifier() string
	CollatorIdentifier() string
	UsesMetricSystem() bool
	DecimalSeparator() string
	GroupingSeparator() string
	CurrencyCode() string
	CurrencySymbol() string
	CalendarIdentifier() string
	QuotationBeginDelimiter() string
	QuotationEndDelimiter() string
	AlternateQuotationBeginDelimiter() string
	AlternateQuotationEndDelimiter() string
}

type Locale struct {
	objc.Object
}

func MakeLocale(ptr unsafe.Pointer) Locale {
	return Locale{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LocaleClass) LocaleWithLocaleIdentifier(ident string) Locale {
	rv := ffi.CallMethod[Locale](lc, "localeWithLocaleIdentifier:", ident)
	return rv
}

func (l_ Locale) InitWithLocaleIdentifier(string_ string) Locale {
	rv := ffi.CallMethod[Locale](l_, "initWithLocaleIdentifier:", string_)
	return rv
}

func (lc _LocaleClass) Alloc() Locale {
	rv := ffi.CallMethod[Locale](lc, "alloc")
	return rv
}

func (lc _LocaleClass) New() Locale {
	rv := ffi.CallMethod[Locale](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLocale() Locale {
	return LocaleClass.New()
}

func (l_ Locale) Init() Locale {
	rv := ffi.CallMethod[Locale](l_, "init")
	return rv
}

func (lc _LocaleClass) CanonicalLocaleIdentifierFromString(string_ string) string {
	rv := ffi.CallMethod[string](lc, "canonicalLocaleIdentifierFromString:", string_)
	return rv
}

func (lc _LocaleClass) ComponentsFromLocaleIdentifier(string_ string) map[string]string {
	rv := ffi.CallMethod[map[string]string](lc, "componentsFromLocaleIdentifier:", string_)
	return rv
}

func (lc _LocaleClass) LocaleIdentifierFromComponents(dict map[string]string) string {
	rv := ffi.CallMethod[string](lc, "localeIdentifierFromComponents:", dict)
	return rv
}

func (lc _LocaleClass) CanonicalLanguageIdentifierFromString(string_ string) string {
	rv := ffi.CallMethod[string](lc, "canonicalLanguageIdentifierFromString:", string_)
	return rv
}

func (lc _LocaleClass) LocaleIdentifierFromWindowsLocaleCode(lcid uint32) string {
	rv := ffi.CallMethod[string](lc, "localeIdentifierFromWindowsLocaleCode:", lcid)
	return rv
}

func (lc _LocaleClass) WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) uint32 {
	rv := ffi.CallMethod[uint32](lc, "windowsLocaleCodeFromLocaleIdentifier:", localeIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForLocaleIdentifier(localeIdentifier string) string {
	rv := ffi.CallMethod[string](l_, "localizedStringForLocaleIdentifier:", localeIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForCountryCode(countryCode string) string {
	rv := ffi.CallMethod[string](l_, "localizedStringForCountryCode:", countryCode)
	return rv
}

func (l_ Locale) LocalizedStringForLanguageCode(languageCode string) string {
	rv := ffi.CallMethod[string](l_, "localizedStringForLanguageCode:", languageCode)
	return rv
}

func (l_ Locale) LocalizedStringForScriptCode(scriptCode string) string {
	rv := ffi.CallMethod[string](l_, "localizedStringForScriptCode:", scriptCode)
	return rv
}

func (l_ Locale) LocalizedStringForVariantCode(variantCode string) string {
	rv := ffi.CallMethod[string](l_, "localizedStringForVariantCode:", variantCode)
	return rv
}

func (l_ Locale) LocalizedStringForCollationIdentifier(collationIdentifier string) string {
	rv := ffi.CallMethod[string](l_, "localizedStringForCollationIdentifier:", collationIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForCollatorIdentifier(collatorIdentifier string) string {
	rv := ffi.CallMethod[string](l_, "localizedStringForCollatorIdentifier:", collatorIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForCurrencyCode(currencyCode string) string {
	rv := ffi.CallMethod[string](l_, "localizedStringForCurrencyCode:", currencyCode)
	return rv
}

func (l_ Locale) LocalizedStringForCalendarIdentifier(calendarIdentifier string) string {
	rv := ffi.CallMethod[string](l_, "localizedStringForCalendarIdentifier:", calendarIdentifier)
	return rv
}

func (l_ Locale) ObjectForKey(key LocaleKey) objc.Object {
	rv := ffi.CallMethod[objc.Object](l_, "objectForKey:", key)
	return rv
}

func (l_ Locale) DisplayNameForKey_Value(key LocaleKey, value objc.IObject) string {
	rv := ffi.CallMethod[string](l_, "displayNameForKey:value:", key, value)
	return rv
}

func (lc _LocaleClass) CharacterDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	rv := ffi.CallMethod[LocaleLanguageDirection](lc, "characterDirectionForLanguage:", isoLangCode)
	return rv
}

func (lc _LocaleClass) LineDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	rv := ffi.CallMethod[LocaleLanguageDirection](lc, "lineDirectionForLanguage:", isoLangCode)
	return rv
}

func (lc _LocaleClass) AutoupdatingCurrentLocale() Locale {
	rv := ffi.CallMethod[Locale](lc, "autoupdatingCurrentLocale")
	return rv
}

func (lc _LocaleClass) CurrentLocale() Locale {
	rv := ffi.CallMethod[Locale](lc, "currentLocale")
	return rv
}

func (lc _LocaleClass) SystemLocale() Locale {
	rv := ffi.CallMethod[Locale](lc, "systemLocale")
	return rv
}

func (lc _LocaleClass) AvailableLocaleIdentifiers() []string {
	rv := ffi.CallMethod[[]string](lc, "availableLocaleIdentifiers")
	return rv
}

func (lc _LocaleClass) ISOCountryCodes() []string {
	rv := ffi.CallMethod[[]string](lc, "ISOCountryCodes")
	return rv
}

func (lc _LocaleClass) ISOLanguageCodes() []string {
	rv := ffi.CallMethod[[]string](lc, "ISOLanguageCodes")
	return rv
}

func (lc _LocaleClass) ISOCurrencyCodes() []string {
	rv := ffi.CallMethod[[]string](lc, "ISOCurrencyCodes")
	return rv
}

func (lc _LocaleClass) CommonISOCurrencyCodes() []string {
	rv := ffi.CallMethod[[]string](lc, "commonISOCurrencyCodes")
	return rv
}

func (l_ Locale) LocaleIdentifier() string {
	rv := ffi.CallMethod[string](l_, "localeIdentifier")
	return rv
}

func (l_ Locale) CountryCode() string {
	rv := ffi.CallMethod[string](l_, "countryCode")
	return rv
}

func (l_ Locale) LanguageCode() string {
	rv := ffi.CallMethod[string](l_, "languageCode")
	return rv
}

func (l_ Locale) ScriptCode() string {
	rv := ffi.CallMethod[string](l_, "scriptCode")
	return rv
}

func (l_ Locale) VariantCode() string {
	rv := ffi.CallMethod[string](l_, "variantCode")
	return rv
}

func (l_ Locale) ExemplarCharacterSet() CharacterSet {
	rv := ffi.CallMethod[CharacterSet](l_, "exemplarCharacterSet")
	return rv
}

func (l_ Locale) CollationIdentifier() string {
	rv := ffi.CallMethod[string](l_, "collationIdentifier")
	return rv
}

func (l_ Locale) CollatorIdentifier() string {
	rv := ffi.CallMethod[string](l_, "collatorIdentifier")
	return rv
}

func (l_ Locale) UsesMetricSystem() bool {
	rv := ffi.CallMethod[bool](l_, "usesMetricSystem")
	return rv
}

func (l_ Locale) DecimalSeparator() string {
	rv := ffi.CallMethod[string](l_, "decimalSeparator")
	return rv
}

func (l_ Locale) GroupingSeparator() string {
	rv := ffi.CallMethod[string](l_, "groupingSeparator")
	return rv
}

func (l_ Locale) CurrencyCode() string {
	rv := ffi.CallMethod[string](l_, "currencyCode")
	return rv
}

func (l_ Locale) CurrencySymbol() string {
	rv := ffi.CallMethod[string](l_, "currencySymbol")
	return rv
}

func (l_ Locale) CalendarIdentifier() string {
	rv := ffi.CallMethod[string](l_, "calendarIdentifier")
	return rv
}

func (l_ Locale) QuotationBeginDelimiter() string {
	rv := ffi.CallMethod[string](l_, "quotationBeginDelimiter")
	return rv
}

func (l_ Locale) QuotationEndDelimiter() string {
	rv := ffi.CallMethod[string](l_, "quotationEndDelimiter")
	return rv
}

func (l_ Locale) AlternateQuotationBeginDelimiter() string {
	rv := ffi.CallMethod[string](l_, "alternateQuotationBeginDelimiter")
	return rv
}

func (l_ Locale) AlternateQuotationEndDelimiter() string {
	rv := ffi.CallMethod[string](l_, "alternateQuotationEndDelimiter")
	return rv
}

func (lc _LocaleClass) PreferredLanguages() []string {
	rv := ffi.CallMethod[[]string](lc, "preferredLanguages")
	return rv
}
