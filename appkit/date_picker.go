// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var DatePickerClass = _DatePickerClass{objc.GetClass("NSDatePicker")}

type _DatePickerClass struct {
	objc.Class
}

type IDatePicker interface {
	IControl
	IsBezeled() bool
	SetBezeled(value bool)
	IsBordered() bool
	SetBordered(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	TextColor() Color
	SetTextColor(value IColor)
	DatePickerStyle() DatePickerStyle
	SetDatePickerStyle(value DatePickerStyle)
	PresentsCalendarOverlay() bool
	SetPresentsCalendarOverlay(value bool)
	Delegate() DatePickerCellDelegateWrapper
	SetDelegate(value DatePickerCellDelegate)
	SetDelegate0(value objc.IObject)
	DatePickerElements() DatePickerElementFlags
	SetDatePickerElements(value DatePickerElementFlags)
	Calendar() foundation.Calendar
	SetCalendar(value foundation.ICalendar)
	Locale() foundation.Locale
	SetLocale(value foundation.ILocale)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(value DatePickerMode)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.ITimeZone)
	DateValue() foundation.Date
	SetDateValue(value foundation.IDate)
	TimeInterval() foundation.TimeInterval
	SetTimeInterval(value foundation.TimeInterval)
	MinDate() foundation.Date
	SetMinDate(value foundation.IDate)
	MaxDate() foundation.Date
	SetMaxDate(value foundation.IDate)
}

type DatePicker struct {
	Control
}

func MakeDatePicker(ptr unsafe.Pointer) DatePicker {
	return DatePicker{
		Control: MakeControl(ptr),
	}
}

func (d_ DatePicker) InitWithFrame(frameRect foundation.Rect) DatePicker {
	rv := ffi.CallMethod[DatePicker](d_, "initWithFrame:", frameRect)
	return rv
}

func (d_ DatePicker) Init() DatePicker {
	rv := ffi.CallMethod[DatePicker](d_, "init")
	return rv
}

func (dc _DatePickerClass) Alloc() DatePicker {
	rv := ffi.CallMethod[DatePicker](dc, "alloc")
	return rv
}

func (dc _DatePickerClass) New() DatePicker {
	rv := ffi.CallMethod[DatePicker](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDatePicker() DatePicker {
	return DatePickerClass.New()
}

func (d_ DatePicker) IsBezeled() bool {
	rv := ffi.CallMethod[bool](d_, "isBezeled")
	return rv
}

func (d_ DatePicker) SetBezeled(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setBezeled:", value)
}

func (d_ DatePicker) IsBordered() bool {
	rv := ffi.CallMethod[bool](d_, "isBordered")
	return rv
}

func (d_ DatePicker) SetBordered(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setBordered:", value)
}

func (d_ DatePicker) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](d_, "backgroundColor")
	return rv
}

func (d_ DatePicker) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](d_, "setBackgroundColor:", value)
}

func (d_ DatePicker) DrawsBackground() bool {
	rv := ffi.CallMethod[bool](d_, "drawsBackground")
	return rv
}

func (d_ DatePicker) SetDrawsBackground(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setDrawsBackground:", value)
}

func (d_ DatePicker) TextColor() Color {
	rv := ffi.CallMethod[Color](d_, "textColor")
	return rv
}

func (d_ DatePicker) SetTextColor(value IColor) {
	ffi.CallMethod[ffi.Void](d_, "setTextColor:", value)
}

func (d_ DatePicker) DatePickerStyle() DatePickerStyle {
	rv := ffi.CallMethod[DatePickerStyle](d_, "datePickerStyle")
	return rv
}

func (d_ DatePicker) SetDatePickerStyle(value DatePickerStyle) {
	ffi.CallMethod[ffi.Void](d_, "setDatePickerStyle:", value)
}

func (d_ DatePicker) PresentsCalendarOverlay() bool {
	rv := ffi.CallMethod[bool](d_, "presentsCalendarOverlay")
	return rv
}

func (d_ DatePicker) SetPresentsCalendarOverlay(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setPresentsCalendarOverlay:", value)
}

func (d_ DatePicker) Delegate() DatePickerCellDelegateWrapper {
	rv := ffi.CallMethod[DatePickerCellDelegateWrapper](d_, "delegate")
	return rv
}

func (d_ DatePicker) SetDelegate(value DatePickerCellDelegate) {
	po := ffi.CreateProtocol("NSDatePickerCellDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(d_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](d_, "setDelegate:", po)
}

func (d_ DatePicker) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "setDelegate:", value)
}

func (d_ DatePicker) DatePickerElements() DatePickerElementFlags {
	rv := ffi.CallMethod[DatePickerElementFlags](d_, "datePickerElements")
	return rv
}

func (d_ DatePicker) SetDatePickerElements(value DatePickerElementFlags) {
	ffi.CallMethod[ffi.Void](d_, "setDatePickerElements:", value)
}

func (d_ DatePicker) Calendar() foundation.Calendar {
	rv := ffi.CallMethod[foundation.Calendar](d_, "calendar")
	return rv
}

func (d_ DatePicker) SetCalendar(value foundation.ICalendar) {
	ffi.CallMethod[ffi.Void](d_, "setCalendar:", value)
}

func (d_ DatePicker) Locale() foundation.Locale {
	rv := ffi.CallMethod[foundation.Locale](d_, "locale")
	return rv
}

func (d_ DatePicker) SetLocale(value foundation.ILocale) {
	ffi.CallMethod[ffi.Void](d_, "setLocale:", value)
}

func (d_ DatePicker) DatePickerMode() DatePickerMode {
	rv := ffi.CallMethod[DatePickerMode](d_, "datePickerMode")
	return rv
}

func (d_ DatePicker) SetDatePickerMode(value DatePickerMode) {
	ffi.CallMethod[ffi.Void](d_, "setDatePickerMode:", value)
}

func (d_ DatePicker) TimeZone() foundation.TimeZone {
	rv := ffi.CallMethod[foundation.TimeZone](d_, "timeZone")
	return rv
}

func (d_ DatePicker) SetTimeZone(value foundation.ITimeZone) {
	ffi.CallMethod[ffi.Void](d_, "setTimeZone:", value)
}

func (d_ DatePicker) DateValue() foundation.Date {
	rv := ffi.CallMethod[foundation.Date](d_, "dateValue")
	return rv
}

func (d_ DatePicker) SetDateValue(value foundation.IDate) {
	ffi.CallMethod[ffi.Void](d_, "setDateValue:", value)
}

func (d_ DatePicker) TimeInterval() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](d_, "timeInterval")
	return rv
}

func (d_ DatePicker) SetTimeInterval(value foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](d_, "setTimeInterval:", value)
}

func (d_ DatePicker) MinDate() foundation.Date {
	rv := ffi.CallMethod[foundation.Date](d_, "minDate")
	return rv
}

func (d_ DatePicker) SetMinDate(value foundation.IDate) {
	ffi.CallMethod[ffi.Void](d_, "setMinDate:", value)
}

func (d_ DatePicker) MaxDate() foundation.Date {
	rv := ffi.CallMethod[foundation.Date](d_, "maxDate")
	return rv
}

func (d_ DatePicker) SetMaxDate(value foundation.IDate) {
	ffi.CallMethod[ffi.Void](d_, "setMaxDate:", value)
}
