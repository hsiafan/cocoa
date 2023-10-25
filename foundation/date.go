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
	rv := objc.CallMethod[Date](dc, objc.SEL("date"))
	return rv
}

func (dc _DateClass) DateWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](dc, objc.SEL("dateWithTimeIntervalSinceNow:"), secs)
	return rv
}

func (dc _DateClass) DateWithTimeInterval_SinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := objc.CallMethod[Date](dc, objc.SEL("dateWithTimeInterval:sinceDate:"), secsToBeAdded, objc.ExtractPtr(date))
	return rv
}

func (dc _DateClass) DateWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := objc.CallMethod[Date](dc, objc.SEL("dateWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}

func (dc _DateClass) DateWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](dc, objc.SEL("dateWithTimeIntervalSince1970:"), secs)
	return rv
}

func (d_ Date) Init() Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("init"))
	return rv
}

func (d_ Date) InitWithTimeIntervalSinceNow(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("initWithTimeIntervalSinceNow:"), secs)
	return rv
}

func (d_ Date) InitWithTimeInterval_SinceDate(secsToBeAdded TimeInterval, date IDate) Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("initWithTimeInterval:sinceDate:"), secsToBeAdded, objc.ExtractPtr(date))
	return rv
}

func (d_ Date) InitWithTimeIntervalSinceReferenceDate(ti TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("initWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}

func (d_ Date) InitWithTimeIntervalSince1970(secs TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("initWithTimeIntervalSince1970:"), secs)
	return rv
}

func (d_ Date) DateByAddingTimeInterval(ti TimeInterval) Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("dateByAddingTimeInterval:"), ti)
	return rv
}

func (dc _DateClass) Alloc() Date {
	rv := objc.CallMethod[Date](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DateClass) New() Date {
	rv := objc.CallMethod[Date](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDate() Date {
	return DateClass.New()
}

func (d_ Date) IsEqualToDate(otherDate IDate) bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("isEqualToDate:"), objc.ExtractPtr(otherDate))
	return rv
}

func (d_ Date) EarlierDate(anotherDate IDate) Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("earlierDate:"), objc.ExtractPtr(anotherDate))
	return rv
}

func (d_ Date) LaterDate(anotherDate IDate) Date {
	rv := objc.CallMethod[Date](d_, objc.SEL("laterDate:"), objc.ExtractPtr(anotherDate))
	return rv
}

func (d_ Date) Compare(other IDate) ComparisonResult {
	rv := objc.CallMethod[ComparisonResult](d_, objc.SEL("compare:"), objc.ExtractPtr(other))
	return rv
}

func (d_ Date) TimeIntervalSinceDate(anotherDate IDate) TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.SEL("timeIntervalSinceDate:"), objc.ExtractPtr(anotherDate))
	return rv
}

func (d_ Date) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.CallMethod[string](d_, objc.SEL("descriptionWithLocale:"), objc.ExtractPtr(locale))
	return rv
}

// deprecated
func (dc _DateClass) DateWithNaturalLanguageString(string_ string) objc.Object {
	rv := objc.CallMethod[objc.Object](dc, objc.SEL("dateWithNaturalLanguageString:"), string_)
	return rv
}

// deprecated
func (dc _DateClass) DateWithNaturalLanguageString_Locale(string_ string, locale objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](dc, objc.SEL("dateWithNaturalLanguageString:locale:"), string_, objc.ExtractPtr(locale))
	return rv
}

// deprecated
func (dc _DateClass) DateWithString(aString string) objc.Object {
	rv := objc.CallMethod[objc.Object](dc, objc.SEL("dateWithString:"), aString)
	return rv
}

// deprecated
func (d_ Date) InitWithString(description string) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.SEL("initWithString:"), description)
	return rv
}

// deprecated
func (d_ Date) AddTimeInterval(seconds TimeInterval) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.SEL("addTimeInterval:"), seconds)
	return rv
}

// deprecated
func (d_ Date) DescriptionWithCalendarFormat_TimeZone_Locale(format string, aTimeZone ITimeZone, locale objc.IObject) string {
	rv := objc.CallMethod[string](d_, objc.SEL("descriptionWithCalendarFormat:timeZone:locale:"), format, objc.ExtractPtr(aTimeZone), objc.ExtractPtr(locale))
	return rv
}

func (dc _DateClass) DistantFuture() Date {
	rv := objc.CallMethod[Date](dc, objc.SEL("distantFuture"))
	return rv
}

func (dc _DateClass) DistantPast() Date {
	rv := objc.CallMethod[Date](dc, objc.SEL("distantPast"))
	return rv
}

func (dc _DateClass) Now() Date {
	rv := objc.CallMethod[Date](dc, objc.SEL("now"))
	return rv
}

func (d_ Date) TimeIntervalSinceNow() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.SEL("timeIntervalSinceNow"))
	return rv
}

func (d_ Date) TimeIntervalSinceReferenceDate() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.SEL("timeIntervalSinceReferenceDate"))
	return rv
}

func (d_ Date) TimeIntervalSince1970() TimeInterval {
	rv := objc.CallMethod[TimeInterval](d_, objc.SEL("timeIntervalSince1970"))
	return rv
}

func (d_ Date) Description() string {
	rv := objc.CallMethod[string](d_, objc.SEL("description"))
	return rv
}
