// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[TimeZone](tc, "timeZoneWithName:", tzName)
	return rv
}

func (tc _TimeZoneClass) TimeZoneWithName_Data(tzName string, aData []byte) TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "timeZoneWithName:data:", tzName, aData)
	return rv
}

func (t_ TimeZone) InitWithName(tzName string) TimeZone {
	rv := ffi.CallMethod[TimeZone](t_, "initWithName:", tzName)
	return rv
}

func (t_ TimeZone) InitWithName_Data(tzName string, aData []byte) TimeZone {
	rv := ffi.CallMethod[TimeZone](t_, "initWithName:data:", tzName, aData)
	return rv
}

func (tc _TimeZoneClass) TimeZoneWithAbbreviation(abbreviation string) TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "timeZoneWithAbbreviation:", abbreviation)
	return rv
}

func (tc _TimeZoneClass) TimeZoneForSecondsFromGMT(seconds int) TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "timeZoneForSecondsFromGMT:", seconds)
	return rv
}

func (tc _TimeZoneClass) Alloc() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "alloc")
	return rv
}

func (tc _TimeZoneClass) New() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTimeZone() TimeZone {
	return TimeZoneClass.New()
}

func (t_ TimeZone) Init() TimeZone {
	rv := ffi.CallMethod[TimeZone](t_, "init")
	return rv
}

func (tc _TimeZoneClass) ResetSystemTimeZone() {
	ffi.CallMethod[ffi.Void](tc, "resetSystemTimeZone")
}

func (t_ TimeZone) AbbreviationForDate(aDate IDate) string {
	rv := ffi.CallMethod[string](t_, "abbreviationForDate:", aDate)
	return rv
}

func (t_ TimeZone) SecondsFromGMTForDate(aDate IDate) int {
	rv := ffi.CallMethod[int](t_, "secondsFromGMTForDate:", aDate)
	return rv
}

func (t_ TimeZone) IsDaylightSavingTimeForDate(aDate IDate) bool {
	rv := ffi.CallMethod[bool](t_, "isDaylightSavingTimeForDate:", aDate)
	return rv
}

func (t_ TimeZone) DaylightSavingTimeOffsetForDate(aDate IDate) TimeInterval {
	rv := ffi.CallMethod[TimeInterval](t_, "daylightSavingTimeOffsetForDate:", aDate)
	return rv
}

func (t_ TimeZone) NextDaylightSavingTimeTransitionAfterDate(aDate IDate) Date {
	rv := ffi.CallMethod[Date](t_, "nextDaylightSavingTimeTransitionAfterDate:", aDate)
	return rv
}

func (t_ TimeZone) IsEqualToTimeZone(aTimeZone ITimeZone) bool {
	rv := ffi.CallMethod[bool](t_, "isEqualToTimeZone:", aTimeZone)
	return rv
}

func (t_ TimeZone) LocalizedName_Locale(style TimeZoneNameStyle, locale ILocale) string {
	rv := ffi.CallMethod[string](t_, "localizedName:locale:", style, locale)
	return rv
}

func (tc _TimeZoneClass) LocalTimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "localTimeZone")
	return rv
}

func (tc _TimeZoneClass) SystemTimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "systemTimeZone")
	return rv
}

func (tc _TimeZoneClass) DefaultTimeZone() TimeZone {
	rv := ffi.CallMethod[TimeZone](tc, "defaultTimeZone")
	return rv
}

func (tc _TimeZoneClass) SetDefaultTimeZone(value ITimeZone) {
	ffi.CallMethod[ffi.Void](tc, "setDefaultTimeZone:", value)
}

func (tc _TimeZoneClass) KnownTimeZoneNames() []string {
	rv := ffi.CallMethod[[]string](tc, "knownTimeZoneNames")
	return rv
}

func (tc _TimeZoneClass) AbbreviationDictionary() map[string]string {
	rv := ffi.CallMethod[map[string]string](tc, "abbreviationDictionary")
	return rv
}

func (tc _TimeZoneClass) SetAbbreviationDictionary(value map[string]string) {
	ffi.CallMethod[ffi.Void](tc, "setAbbreviationDictionary:", value)
}

func (t_ TimeZone) Name() string {
	rv := ffi.CallMethod[string](t_, "name")
	return rv
}

func (t_ TimeZone) Abbreviation() string {
	rv := ffi.CallMethod[string](t_, "abbreviation")
	return rv
}

func (t_ TimeZone) SecondsFromGMT() int {
	rv := ffi.CallMethod[int](t_, "secondsFromGMT")
	return rv
}

func (t_ TimeZone) Data() []byte {
	rv := ffi.CallMethod[[]byte](t_, "data")
	return rv
}

func (tc _TimeZoneClass) TimeZoneDataVersion() string {
	rv := ffi.CallMethod[string](tc, "timeZoneDataVersion")
	return rv
}

func (t_ TimeZone) IsDaylightSavingTime() bool {
	rv := ffi.CallMethod[bool](t_, "isDaylightSavingTime")
	return rv
}

func (t_ TimeZone) DaylightSavingTimeOffset() TimeInterval {
	rv := ffi.CallMethod[TimeInterval](t_, "daylightSavingTimeOffset")
	return rv
}

func (t_ TimeZone) NextDaylightSavingTimeTransition() Date {
	rv := ffi.CallMethod[Date](t_, "nextDaylightSavingTimeTransition")
	return rv
}

func (t_ TimeZone) Description() string {
	rv := ffi.CallMethod[string](t_, "description")
	return rv
}
