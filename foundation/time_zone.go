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
	rv := objc.CallMethod[TimeZone](tc, objc.GetSelector("timeZoneWithName:"), tzName)
	return rv
}

func (tc _TimeZoneClass) TimeZoneWithName_Data(tzName string, aData []byte) TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.GetSelector("timeZoneWithName:data:"), tzName, aData)
	return rv
}

func (t_ TimeZone) InitWithName(tzName string) TimeZone {
	rv := objc.CallMethod[TimeZone](t_, objc.GetSelector("initWithName:"), tzName)
	return rv
}

func (t_ TimeZone) InitWithName_Data(tzName string, aData []byte) TimeZone {
	rv := objc.CallMethod[TimeZone](t_, objc.GetSelector("initWithName:data:"), tzName, aData)
	return rv
}

func (tc _TimeZoneClass) TimeZoneWithAbbreviation(abbreviation string) TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.GetSelector("timeZoneWithAbbreviation:"), abbreviation)
	return rv
}

func (tc _TimeZoneClass) TimeZoneForSecondsFromGMT(seconds int) TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.GetSelector("timeZoneForSecondsFromGMT:"), seconds)
	return rv
}

func (tc _TimeZoneClass) Alloc() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.GetSelector("alloc"))
	return rv
}

func (tc _TimeZoneClass) New() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTimeZone() TimeZone {
	return TimeZoneClass.New()
}

func (t_ TimeZone) Init() TimeZone {
	rv := objc.CallMethod[TimeZone](t_, objc.GetSelector("init"))
	return rv
}

func (tc _TimeZoneClass) ResetSystemTimeZone() {
	objc.CallMethod[objc.Void](tc, objc.GetSelector("resetSystemTimeZone"))
}

func (t_ TimeZone) AbbreviationForDate(aDate IDate) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("abbreviationForDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) SecondsFromGMTForDate(aDate IDate) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("secondsFromGMTForDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) IsDaylightSavingTimeForDate(aDate IDate) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isDaylightSavingTimeForDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) DaylightSavingTimeOffsetForDate(aDate IDate) TimeInterval {
	rv := objc.CallMethod[TimeInterval](t_, objc.GetSelector("daylightSavingTimeOffsetForDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) NextDaylightSavingTimeTransitionAfterDate(aDate IDate) Date {
	rv := objc.CallMethod[Date](t_, objc.GetSelector("nextDaylightSavingTimeTransitionAfterDate:"), objc.ExtractPtr(aDate))
	return rv
}

func (t_ TimeZone) IsEqualToTimeZone(aTimeZone ITimeZone) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isEqualToTimeZone:"), objc.ExtractPtr(aTimeZone))
	return rv
}

func (t_ TimeZone) LocalizedName_Locale(style TimeZoneNameStyle, locale ILocale) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("localizedName:locale:"), style, objc.ExtractPtr(locale))
	return rv
}

func (tc _TimeZoneClass) LocalTimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.GetSelector("localTimeZone"))
	return rv
}

func (tc _TimeZoneClass) SystemTimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.GetSelector("systemTimeZone"))
	return rv
}

func (tc _TimeZoneClass) DefaultTimeZone() TimeZone {
	rv := objc.CallMethod[TimeZone](tc, objc.GetSelector("defaultTimeZone"))
	return rv
}

func (tc _TimeZoneClass) SetDefaultTimeZone(value ITimeZone) {
	objc.CallMethod[objc.Void](tc, objc.GetSelector("setDefaultTimeZone:"), objc.ExtractPtr(value))
}

func (tc _TimeZoneClass) KnownTimeZoneNames() []string {
	rv := objc.CallMethod[[]string](tc, objc.GetSelector("knownTimeZoneNames"))
	return rv
}

func (tc _TimeZoneClass) AbbreviationDictionary() map[string]string {
	rv := objc.CallMethod[map[string]string](tc, objc.GetSelector("abbreviationDictionary"))
	return rv
}

func (tc _TimeZoneClass) SetAbbreviationDictionary(value map[string]string) {
	objc.CallMethod[objc.Void](tc, objc.GetSelector("setAbbreviationDictionary:"), value)
}

func (t_ TimeZone) Name() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("name"))
	return rv
}

func (t_ TimeZone) Abbreviation() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("abbreviation"))
	return rv
}

func (t_ TimeZone) SecondsFromGMT() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("secondsFromGMT"))
	return rv
}

func (t_ TimeZone) Data() []byte {
	rv := objc.CallMethod[[]byte](t_, objc.GetSelector("data"))
	return rv
}

func (tc _TimeZoneClass) TimeZoneDataVersion() string {
	rv := objc.CallMethod[string](tc, objc.GetSelector("timeZoneDataVersion"))
	return rv
}

func (t_ TimeZone) IsDaylightSavingTime() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isDaylightSavingTime"))
	return rv
}

func (t_ TimeZone) DaylightSavingTimeOffset() TimeInterval {
	rv := objc.CallMethod[TimeInterval](t_, objc.GetSelector("daylightSavingTimeOffset"))
	return rv
}

func (t_ TimeZone) NextDaylightSavingTimeTransition() Date {
	rv := objc.CallMethod[Date](t_, objc.GetSelector("nextDaylightSavingTimeTransition"))
	return rv
}

func (t_ TimeZone) Description() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("description"))
	return rv
}
