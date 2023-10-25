// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[Locale](lc, objc.SEL("localeWithLocaleIdentifier:"), ident)
	return rv
}

func (l_ Locale) InitWithLocaleIdentifier(string_ string) Locale {
	rv := objc.CallMethod[Locale](l_, objc.SEL("initWithLocaleIdentifier:"), string_)
	return rv
}

func (lc _LocaleClass) Alloc() Locale {
	rv := objc.CallMethod[Locale](lc, objc.SEL("alloc"))
	return rv
}

func (lc _LocaleClass) New() Locale {
	rv := objc.CallMethod[Locale](lc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewLocale() Locale {
	return LocaleClass.New()
}

func (l_ Locale) Init() Locale {
	rv := objc.CallMethod[Locale](l_, objc.SEL("init"))
	return rv
}

func (lc _LocaleClass) CanonicalLocaleIdentifierFromString(string_ string) string {
	rv := objc.CallMethod[string](lc, objc.SEL("canonicalLocaleIdentifierFromString:"), string_)
	return rv
}

func (lc _LocaleClass) ComponentsFromLocaleIdentifier(string_ string) map[string]string {
	rv := objc.CallMethod[map[string]string](lc, objc.SEL("componentsFromLocaleIdentifier:"), string_)
	return rv
}

func (lc _LocaleClass) LocaleIdentifierFromComponents(dict map[string]string) string {
	rv := objc.CallMethod[string](lc, objc.SEL("localeIdentifierFromComponents:"), dict)
	return rv
}

func (lc _LocaleClass) CanonicalLanguageIdentifierFromString(string_ string) string {
	rv := objc.CallMethod[string](lc, objc.SEL("canonicalLanguageIdentifierFromString:"), string_)
	return rv
}

func (lc _LocaleClass) LocaleIdentifierFromWindowsLocaleCode(lcid uint32) string {
	rv := objc.CallMethod[string](lc, objc.SEL("localeIdentifierFromWindowsLocaleCode:"), lcid)
	return rv
}

func (lc _LocaleClass) WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) uint32 {
	rv := objc.CallMethod[uint32](lc, objc.SEL("windowsLocaleCodeFromLocaleIdentifier:"), localeIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForLocaleIdentifier(localeIdentifier string) string {
	rv := objc.CallMethod[string](l_, objc.SEL("localizedStringForLocaleIdentifier:"), localeIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForCountryCode(countryCode string) string {
	rv := objc.CallMethod[string](l_, objc.SEL("localizedStringForCountryCode:"), countryCode)
	return rv
}

func (l_ Locale) LocalizedStringForLanguageCode(languageCode string) string {
	rv := objc.CallMethod[string](l_, objc.SEL("localizedStringForLanguageCode:"), languageCode)
	return rv
}

func (l_ Locale) LocalizedStringForScriptCode(scriptCode string) string {
	rv := objc.CallMethod[string](l_, objc.SEL("localizedStringForScriptCode:"), scriptCode)
	return rv
}

func (l_ Locale) LocalizedStringForVariantCode(variantCode string) string {
	rv := objc.CallMethod[string](l_, objc.SEL("localizedStringForVariantCode:"), variantCode)
	return rv
}

func (l_ Locale) LocalizedStringForCollationIdentifier(collationIdentifier string) string {
	rv := objc.CallMethod[string](l_, objc.SEL("localizedStringForCollationIdentifier:"), collationIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForCollatorIdentifier(collatorIdentifier string) string {
	rv := objc.CallMethod[string](l_, objc.SEL("localizedStringForCollatorIdentifier:"), collatorIdentifier)
	return rv
}

func (l_ Locale) LocalizedStringForCurrencyCode(currencyCode string) string {
	rv := objc.CallMethod[string](l_, objc.SEL("localizedStringForCurrencyCode:"), currencyCode)
	return rv
}

func (l_ Locale) LocalizedStringForCalendarIdentifier(calendarIdentifier string) string {
	rv := objc.CallMethod[string](l_, objc.SEL("localizedStringForCalendarIdentifier:"), calendarIdentifier)
	return rv
}

func (l_ Locale) ObjectForKey(key LocaleKey) objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.SEL("objectForKey:"), key)
	return rv
}

func (l_ Locale) DisplayNameForKey_Value(key LocaleKey, value objc.IObject) string {
	rv := objc.CallMethod[string](l_, objc.SEL("displayNameForKey:value:"), key, objc.ExtractPtr(value))
	return rv
}

func (lc _LocaleClass) CharacterDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	rv := objc.CallMethod[LocaleLanguageDirection](lc, objc.SEL("characterDirectionForLanguage:"), isoLangCode)
	return rv
}

func (lc _LocaleClass) LineDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	rv := objc.CallMethod[LocaleLanguageDirection](lc, objc.SEL("lineDirectionForLanguage:"), isoLangCode)
	return rv
}

func (lc _LocaleClass) AutoupdatingCurrentLocale() Locale {
	rv := objc.CallMethod[Locale](lc, objc.SEL("autoupdatingCurrentLocale"))
	return rv
}

func (lc _LocaleClass) CurrentLocale() Locale {
	rv := objc.CallMethod[Locale](lc, objc.SEL("currentLocale"))
	return rv
}

func (lc _LocaleClass) SystemLocale() Locale {
	rv := objc.CallMethod[Locale](lc, objc.SEL("systemLocale"))
	return rv
}

func (lc _LocaleClass) AvailableLocaleIdentifiers() []string {
	rv := objc.CallMethod[[]string](lc, objc.SEL("availableLocaleIdentifiers"))
	return rv
}

func (lc _LocaleClass) ISOCountryCodes() []string {
	rv := objc.CallMethod[[]string](lc, objc.SEL("ISOCountryCodes"))
	return rv
}

func (lc _LocaleClass) ISOLanguageCodes() []string {
	rv := objc.CallMethod[[]string](lc, objc.SEL("ISOLanguageCodes"))
	return rv
}

func (lc _LocaleClass) ISOCurrencyCodes() []string {
	rv := objc.CallMethod[[]string](lc, objc.SEL("ISOCurrencyCodes"))
	return rv
}

func (lc _LocaleClass) CommonISOCurrencyCodes() []string {
	rv := objc.CallMethod[[]string](lc, objc.SEL("commonISOCurrencyCodes"))
	return rv
}

func (l_ Locale) LocaleIdentifier() string {
	rv := objc.CallMethod[string](l_, objc.SEL("localeIdentifier"))
	return rv
}

func (l_ Locale) CountryCode() string {
	rv := objc.CallMethod[string](l_, objc.SEL("countryCode"))
	return rv
}

func (l_ Locale) LanguageCode() string {
	rv := objc.CallMethod[string](l_, objc.SEL("languageCode"))
	return rv
}

func (l_ Locale) ScriptCode() string {
	rv := objc.CallMethod[string](l_, objc.SEL("scriptCode"))
	return rv
}

func (l_ Locale) VariantCode() string {
	rv := objc.CallMethod[string](l_, objc.SEL("variantCode"))
	return rv
}

func (l_ Locale) ExemplarCharacterSet() CharacterSet {
	rv := objc.CallMethod[CharacterSet](l_, objc.SEL("exemplarCharacterSet"))
	return rv
}

func (l_ Locale) CollationIdentifier() string {
	rv := objc.CallMethod[string](l_, objc.SEL("collationIdentifier"))
	return rv
}

func (l_ Locale) CollatorIdentifier() string {
	rv := objc.CallMethod[string](l_, objc.SEL("collatorIdentifier"))
	return rv
}

func (l_ Locale) UsesMetricSystem() bool {
	rv := objc.CallMethod[bool](l_, objc.SEL("usesMetricSystem"))
	return rv
}

func (l_ Locale) DecimalSeparator() string {
	rv := objc.CallMethod[string](l_, objc.SEL("decimalSeparator"))
	return rv
}

func (l_ Locale) GroupingSeparator() string {
	rv := objc.CallMethod[string](l_, objc.SEL("groupingSeparator"))
	return rv
}

func (l_ Locale) CurrencyCode() string {
	rv := objc.CallMethod[string](l_, objc.SEL("currencyCode"))
	return rv
}

func (l_ Locale) CurrencySymbol() string {
	rv := objc.CallMethod[string](l_, objc.SEL("currencySymbol"))
	return rv
}

func (l_ Locale) CalendarIdentifier() string {
	rv := objc.CallMethod[string](l_, objc.SEL("calendarIdentifier"))
	return rv
}

func (l_ Locale) QuotationBeginDelimiter() string {
	rv := objc.CallMethod[string](l_, objc.SEL("quotationBeginDelimiter"))
	return rv
}

func (l_ Locale) QuotationEndDelimiter() string {
	rv := objc.CallMethod[string](l_, objc.SEL("quotationEndDelimiter"))
	return rv
}

func (l_ Locale) AlternateQuotationBeginDelimiter() string {
	rv := objc.CallMethod[string](l_, objc.SEL("alternateQuotationBeginDelimiter"))
	return rv
}

func (l_ Locale) AlternateQuotationEndDelimiter() string {
	rv := objc.CallMethod[string](l_, objc.SEL("alternateQuotationEndDelimiter"))
	return rv
}

func (lc _LocaleClass) PreferredLanguages() []string {
	rv := objc.CallMethod[[]string](lc, objc.SEL("preferredLanguages"))
	return rv
}
