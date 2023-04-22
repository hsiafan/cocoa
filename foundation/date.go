// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var DateClass = _DateClass{objc.GetClass("NSDate")}

type _DateClass struct {
	objc.Class
}

type IDate interface {
	objc.IObject
	IsEqualToDate(otherDate IDate) bool
	EarlierDate(anotherDate IDate) Date
	LaterDate(anotherDate IDate) Date
	Compare(other IDate) ComparisonResult
	TimeIntervalSinceDate(anotherDate IDate) TimeInterval
	DescriptionWithLocale(locale objc.IObject) string
	// deprecated
	InitWithString(description string) objc.Object
	// deprecated
	AddTimeInterval(seconds TimeInterval) objc.Object
	// deprecated
	DescriptionWithCalendarFormat_TimeZone_Locale(format string, aTimeZone ITimeZone, locale objc.IObject) string
	TimeIntervalSinceNow() TimeInterval
	TimeIntervalSinceReferenceDate() TimeInterval
	TimeIntervalSince1970() TimeInterval
	Description() string
}

type Date struct {
	objc.Object
}

func MakeDate(ptr unsafe.Pointer) Date {
	return Date{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DateClass) Date() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("date"))
	return rv
}

func (dc _DateClass) DateWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("dateWithTimeIntervalSinceNow:"), secs)
	return rv
}

func (dc _DateClass) DateWithTimeInterval_SinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("dateWithTimeInterval:sinceDate:"), secsToBeAdded, date)
	return rv
}

func (dc _DateClass) DateWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("dateWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}

func (dc _DateClass) DateWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("dateWithTimeIntervalSince1970:"), secs)
	return rv
}

func (d_ Date) Init() Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("init"))
	return rv
}

func (d_ Date) InitWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("initWithTimeIntervalSinceNow:"), secs)
	return rv
}

func (d_ Date) InitWithTimeInterval_SinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("initWithTimeInterval:sinceDate:"), secsToBeAdded, date)
	return rv
}

func (d_ Date) InitWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("initWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}

func (d_ Date) InitWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("initWithTimeIntervalSince1970:"), secs)
	return rv
}

func (d_ Date) DateByAddingTimeInterval(ti TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("dateByAddingTimeInterval:"), ti)
	return rv
}

func (dc _DateClass) Alloc() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("alloc"))
	return rv
}

func (dc _DateClass) New() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDate() Date {
	return DateClass.New()
}

func (d_ Date) IsEqualToDate(otherDate IDate) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("isEqualToDate:"), otherDate)
	return rv
}

func (d_ Date) EarlierDate(anotherDate IDate) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("earlierDate:"), anotherDate)
	return rv
}

func (d_ Date) LaterDate(anotherDate IDate) Date {
	rv := objc.CallMethod[Date](d_, objc.GetSelector("laterDate:"), anotherDate)
	return rv
}

func (d_ Date) Compare(other IDate) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](d_, objc.GetSelector("compare:"), other)
	return rv
}

func (d_ Date) TimeIntervalSinceDate(anotherDate IDate) TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.GetSelector("timeIntervalSinceDate:"), anotherDate)
	return rv
}

func (d_ Date) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("descriptionWithLocale:"), locale)
	return rv
}

// deprecated
func (dc _DateClass) DateWithNaturalLanguageString(string_ string) objc.Object {
	rv := objc.CallMethod[objc.Object](dc, objc.GetSelector("dateWithNaturalLanguageString:"), string_)
	return rv
}

// deprecated
func (dc _DateClass) DateWithNaturalLanguageString_Locale(string_ string, locale objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](dc, objc.GetSelector("dateWithNaturalLanguageString:locale:"), string_, locale)
	return rv
}

// deprecated
func (dc _DateClass) DateWithString(aString string) objc.Object {
	rv := objc.CallMethod[objc.Object](dc, objc.GetSelector("dateWithString:"), aString)
	return rv
}

// deprecated
func (d_ Date) InitWithString(description string) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("initWithString:"), description)
	return rv
}

// deprecated
func (d_ Date) AddTimeInterval(seconds TimeInterval) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("addTimeInterval:"), seconds)
	return rv
}

// deprecated
func (d_ Date) DescriptionWithCalendarFormat_TimeZone_Locale(format string, aTimeZone ITimeZone, locale objc.IObject) string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("descriptionWithCalendarFormat:timeZone:locale:"), format, aTimeZone, locale)
	return rv
}

func (dc _DateClass) DistantFuture() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("distantFuture"))
	return rv
}

func (dc _DateClass) DistantPast() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("distantPast"))
	return rv
}

func (dc _DateClass) Now() Date {
	rv := objc.CallMethod[Date](dc, objc.GetSelector("now"))
	return rv
}

func (d_ Date) TimeIntervalSinceNow() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.GetSelector("timeIntervalSinceNow"))
	return rv
}

func (d_ Date) TimeIntervalSinceReferenceDate() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.GetSelector("timeIntervalSinceReferenceDate"))
	return rv
}

func (d_ Date) TimeIntervalSince1970() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.GetSelector("timeIntervalSince1970"))
	return rv
}

func (d_ Date) Description() string {
	rv := objc.CallMethod[string](d_, objc.GetSelector("description"))
	return rv
}
