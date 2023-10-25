// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var TimeZoneClass = _TimeZoneClass{objc.GetClass("NSTimeZone")}

type _TimeZoneClass struct {
	objc.Class
}

type ITimeZone interface {
	objc.IObject
	AbbreviationForDate(aDate IDate) string
	SecondsFromGMTForDate(aDate IDate) int
	IsDaylightSavingTimeForDate(aDate IDate) bool
	DaylightSavingTimeOffsetForDate(aDate IDate) TimeInterval
	NextDaylightSavingTimeTransitionAfterDate(aDate IDate) Date
	IsEqualToTimeZone(aTimeZone ITimeZone) bool
	LocalizedName_Locale(style TimeZoneNameStyle, locale ILocale) string
	Name() string
	Abbreviation() string
	SecondsFromGMT() int
	Data() []byte
	IsDaylightSavingTime() bool
	DaylightSavingTimeOffset() TimeInterval
	NextDaylightSavingTimeTransition() Date
	Description() string
}

type TimeZone struct {
	objc.Object
}

func MakeTimeZone(ptr unsafe.Pointer) TimeZone {
	return TimeZone{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TimeZoneClass) TimeZoneWithName(tzName string) TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.SEL("timeZoneWithName:"), tzName)
	return rv
}

func (tc _TimeZoneClass) TimeZoneWithName_Data(tzName string, aData []byte) TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.SEL("timeZoneWithName:data:"), tzName, aData)
	return rv
}

func (t_ TimeZone) InitWithName(tzName string) TimeZone {
	rv := objc.CallMethod[TimeZone](t_, objc.SEL("initWithName:"), tzName)
	return rv
}

func (t_ TimeZone) InitWithName_Data(tzName string, aData []byte) TimeZone {
	rv := objc.CallMethod[TimeZone](t_, objc.SEL("initWithName:data:"), tzName, aData)
	return rv
}

func (tc _TimeZoneClass) TimeZoneWithAbbreviation(abbreviation string) TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.SEL("timeZoneWithAbbreviation:"), abbreviation)
	return rv
}

func (tc _TimeZoneClass) TimeZoneForSecondsFromGMT(seconds int) TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.SEL("timeZoneForSecondsFromGMT:"), seconds)
	return rv
}

func (tc _TimeZoneClass) Alloc() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TimeZoneClass) New() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTimeZone() TimeZone {
	return TimeZoneClass.New()
}

func (t_ TimeZone) Init() TimeZone {
	rv := objc.CallMethod[TimeZone](t_, objc.SEL("init"))
	return rv
}

func (tc _TimeZoneClass) ResetSystemTimeZone() {
	objc.CallMethod[objc.Void](tc, objc.SEL("resetSystemTimeZone"))
}

func (t_ TimeZone) AbbreviationForDate(aDate IDate) string {
	rv := objc.CallMethod[string](t_, objc.SEL("abbreviationForDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) SecondsFromGMTForDate(aDate IDate) int {
	rv := objc.CallMethod[int](t_, objc.SEL("secondsFromGMTForDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) IsDaylightSavingTimeForDate(aDate IDate) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isDaylightSavingTimeForDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) DaylightSavingTimeOffsetForDate(aDate IDate) TimeInterval {
	rv := objc.CallMethod[TimeInterval](t_, objc.SEL("daylightSavingTimeOffsetForDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) NextDaylightSavingTimeTransitionAfterDate(aDate IDate) Date {
	rv := objc.CallMethod[Date](t_, objc.SEL("nextDaylightSavingTimeTransitionAfterDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) IsEqualToTimeZone(aTimeZone ITimeZone) bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isEqualToTimeZone:"), objc.ExtractPtr(aTimeZone))
	return rv
}

func (t_ TimeZone) LocalizedName_Locale(style TimeZoneNameStyle, locale ILocale) string {
	rv := objc.CallMethod[string](t_, objc.SEL("localizedName:locale:"), style, objc.ExtractPtr(locale))
	return rv
}

func (tc _TimeZoneClass) LocalTimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.SEL("localTimeZone"))
	return rv
}

func (tc _TimeZoneClass) SystemTimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.SEL("systemTimeZone"))
	return rv
}

func (tc _TimeZoneClass) DefaultTimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.SEL("defaultTimeZone"))
	return rv
}

func (tc _TimeZoneClass) SetDefaultTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](tc, objc.SEL("setDefaultTimeZone:"), objc.ExtractPtr(value))
}

func (tc _TimeZoneClass) KnownTimeZoneNames() []string {
	rv := objc.CallMethod[[]string](tc, objc.SEL("knownTimeZoneNames"))
	return rv
}

func (tc _TimeZoneClass) AbbreviationDictionary() map[string]string {
	rv := objc.CallMethod[map[string]string](tc, objc.SEL("abbreviationDictionary"))
	return rv
}

func (tc _TimeZoneClass) SetAbbreviationDictionary(value map[string]string) {
	objc.CallMethod[objc.Void](tc, objc.SEL("setAbbreviationDictionary:"), value)
}

func (t_ TimeZone) Name() string {
	rv := objc.CallMethod[string](t_, objc.SEL("name"))
	return rv
}

func (t_ TimeZone) Abbreviation() string {
	rv := objc.CallMethod[string](t_, objc.SEL("abbreviation"))
	return rv
}

func (t_ TimeZone) SecondsFromGMT() int {
	rv := objc.CallMethod[int](t_, objc.SEL("secondsFromGMT"))
	return rv
}

func (t_ TimeZone) Data() []byte {
	rv := objc.CallMethod[[]byte](t_, objc.SEL("data"))
	return rv
}

func (tc _TimeZoneClass) TimeZoneDataVersion() string {
	rv := objc.CallMethod[string](tc, objc.SEL("timeZoneDataVersion"))
	return rv
}

func (t_ TimeZone) IsDaylightSavingTime() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("isDaylightSavingTime"))
	return rv
}

func (t_ TimeZone) DaylightSavingTimeOffset() TimeInterval {
	rv := objc.CallMethod[TimeInterval](t_, objc.SEL("daylightSavingTimeOffset"))
	return rv
}

func (t_ TimeZone) NextDaylightSavingTimeTransition() Date {
	rv := objc.CallMethod[Date](t_, objc.SEL("nextDaylightSavingTimeTransition"))
	return rv
}

func (t_ TimeZone) Description() string {
	rv := objc.CallMethod[string](t_, objc.SEL("description"))
	return rv
}
