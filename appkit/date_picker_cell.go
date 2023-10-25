// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var DatePickerCellClass = _DatePickerCellClass{objc.GetClass("NSDatePickerCell")}

type _DatePickerCellClass struct {
	objc.Class
}

type IDatePickerCell interface {
	IActionCell
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	TextColor() Color
	SetTextColor(value IColor)
	DatePickerStyle() DatePickerStyle
	SetDatePickerStyle(value DatePickerStyle)
	DatePickerElements() DatePickerElementFlags
	SetDatePickerElements(value DatePickerElementFlags)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(value DatePickerMode)
	DateValue() foundation.Date
	SetDateValue(value foundation.IDate)
	TimeInterval() foundation.TimeInterval
	SetTimeInterval(value foundation.TimeInterval)
	Calendar() foundation.Calendar
	SetCalendar(value foundation.ICalendar)
	Locale() foundation.Locale
	SetLocale(value foundation.ILocale)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.ITimeZone)
	MinDate() foundation.Date
	SetMinDate(value foundation.IDate)
	MaxDate() foundation.Date
	SetMaxDate(value foundation.IDate)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
}

type DatePickerCell struct {
	ActionCell
}

func MakeDatePickerCell(ptr unsafe.Pointer) DatePickerCell {
	return DatePickerCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (d_ DatePickerCell) InitTextCell(string_ string) DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](d_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (d_ DatePickerCell) InitImageCell(image IImage) DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](d_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (d_ DatePickerCell) Init() DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](d_, objc.SEL("init"))
	return rv
}

func (dc _DatePickerCellClass) Alloc() DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DatePickerCellClass) New() DatePickerCell {
	rv := objc.CallMethod[DatePickerCell](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDatePickerCell() DatePickerCell {
	return DatePickerCellClass.New()
}

func (d_ DatePickerCell) BackgroundColor() Color {
	rv := objc.CallMethod[Color](d_, objc.SEL("backgroundColor"))
	return rv
}

func (d_ DatePickerCell) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) DrawsBackground() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("drawsBackground"))
	return rv
}

func (d_ DatePickerCell) SetDrawsBackground(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDrawsBackground:"), value)
}

func (d_ DatePickerCell) TextColor() Color {
	rv := objc.CallMethod[Color](d_, objc.SEL("textColor"))
	return rv
}

func (d_ DatePickerCell) SetTextColor(value IColor) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setTextColor:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) DatePickerStyle() DatePickerStyle {
	rv := objc.CallMethod[DatePickerStyle](d_, objc.SEL("datePickerStyle"))
	return rv
}

func (d_ DatePickerCell) SetDatePickerStyle(value DatePickerStyle) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDatePickerStyle:"), value)
}

func (d_ DatePickerCell) DatePickerElements() DatePickerElementFlags {
	rv := objc.CallMethod[DatePickerElementFlags](d_, objc.SEL("datePickerElements"))
	return rv
}

func (d_ DatePickerCell) SetDatePickerElements(value DatePickerElementFlags) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDatePickerElements:"), value)
}

func (d_ DatePickerCell) DatePickerMode() DatePickerMode {
	rv := objc.CallMethod[DatePickerMode](d_, objc.SEL("datePickerMode"))
	return rv
}

func (d_ DatePickerCell) SetDatePickerMode(value DatePickerMode) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDatePickerMode:"), value)
}

func (d_ DatePickerCell) DateValue() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.SEL("dateValue"))
	return rv
}

func (d_ DatePickerCell) SetDateValue(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDateValue:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) TimeInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](d_, objc.SEL("timeInterval"))
	return rv
}

func (d_ DatePickerCell) SetTimeInterval(value foundation.TimeInterval) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setTimeInterval:"), value)
}

func (d_ DatePickerCell) Calendar() foundation.Calendar {
	rv := objc.CallMethod[foundation.Calendar](d_, objc.SEL("calendar"))
	return rv
}

func (d_ DatePickerCell) SetCalendar(value foundation.ICalendar) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setCalendar:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) Locale() foundation.Locale {
	rv := objc.CallMethod[foundation.Locale](d_, objc.SEL("locale"))
	return rv
}

func (d_ DatePickerCell) SetLocale(value foundation.ILocale) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setLocale:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) TimeZone() foundation.TimeZone {
	rv := objc.CallMethod[foundation.TimeZone](d_, objc.SEL("timeZone"))
	return rv
}

func (d_ DatePickerCell) SetTimeZone(value foundation.ITimeZone) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setTimeZone:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) MinDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.SEL("minDate"))
	return rv
}

func (d_ DatePickerCell) SetMinDate(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setMinDate:"), objc.ExtractPtr(value))
}

func (d_ DatePickerCell) MaxDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](d_, objc.SEL("maxDate"))
	return rv
}

func (d_ DatePickerCell) SetMaxDate(value foundation.IDate) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setMaxDate:"), objc.ExtractPtr(value))
}

// weak property
func (d_ DatePickerCell) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.SEL("delegate"))
	return rv
}

// weak property
func (d_ DatePickerCell) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}
