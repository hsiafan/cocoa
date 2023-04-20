// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
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
	Delegate() DatePickerCellDelegateWrapper
	SetDelegate(value DatePickerCellDelegate)
	SetDelegate0(value objc.IObject)
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
	rv := ffi.CallMethod[DatePickerCell](d_, "initTextCell:", string_)
	return rv
}

func (d_ DatePickerCell) InitImageCell(image IImage) DatePickerCell {
	rv := ffi.CallMethod[DatePickerCell](d_, "initImageCell:", image)
	return rv
}

func (d_ DatePickerCell) Init() DatePickerCell {
	rv := ffi.CallMethod[DatePickerCell](d_, "init")
	return rv
}

func (dc _DatePickerCellClass) Alloc() DatePickerCell {
	rv := ffi.CallMethod[DatePickerCell](dc, "alloc")
	return rv
}

func (dc _DatePickerCellClass) New() DatePickerCell {
	rv := ffi.CallMethod[DatePickerCell](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDatePickerCell() DatePickerCell {
	return DatePickerCellClass.New()
}

func (d_ DatePickerCell) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](d_, "backgroundColor")
	return rv
}

func (d_ DatePickerCell) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](d_, "setBackgroundColor:", value)
}

func (d_ DatePickerCell) DrawsBackground() bool {
	rv := ffi.CallMethod[bool](d_, "drawsBackground")
	return rv
}

func (d_ DatePickerCell) SetDrawsBackground(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setDrawsBackground:", value)
}

func (d_ DatePickerCell) TextColor() Color {
	rv := ffi.CallMethod[Color](d_, "textColor")
	return rv
}

func (d_ DatePickerCell) SetTextColor(value IColor) {
	ffi.CallMethod[ffi.Void](d_, "setTextColor:", value)
}

func (d_ DatePickerCell) DatePickerStyle() DatePickerStyle {
	rv := ffi.CallMethod[DatePickerStyle](d_, "datePickerStyle")
	return rv
}

func (d_ DatePickerCell) SetDatePickerStyle(value DatePickerStyle) {
	ffi.CallMethod[ffi.Void](d_, "setDatePickerStyle:", value)
}

func (d_ DatePickerCell) DatePickerElements() DatePickerElementFlags {
	rv := ffi.CallMethod[DatePickerElementFlags](d_, "datePickerElements")
	return rv
}

func (d_ DatePickerCell) SetDatePickerElements(value DatePickerElementFlags) {
	ffi.CallMethod[ffi.Void](d_, "setDatePickerElements:", value)
}

func (d_ DatePickerCell) DatePickerMode() DatePickerMode {
	rv := ffi.CallMethod[DatePickerMode](d_, "datePickerMode")
	return rv
}

func (d_ DatePickerCell) SetDatePickerMode(value DatePickerMode) {
	ffi.CallMethod[ffi.Void](d_, "setDatePickerMode:", value)
}

func (d_ DatePickerCell) DateValue() foundation.Date {
	rv := ffi.CallMethod[foundation.Date](d_, "dateValue")
	return rv
}

func (d_ DatePickerCell) SetDateValue(value foundation.IDate) {
	ffi.CallMethod[ffi.Void](d_, "setDateValue:", value)
}

func (d_ DatePickerCell) TimeInterval() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](d_, "timeInterval")
	return rv
}

func (d_ DatePickerCell) SetTimeInterval(value foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](d_, "setTimeInterval:", value)
}

func (d_ DatePickerCell) Calendar() foundation.Calendar {
	rv := ffi.CallMethod[foundation.Calendar](d_, "calendar")
	return rv
}

func (d_ DatePickerCell) SetCalendar(value foundation.ICalendar) {
	ffi.CallMethod[ffi.Void](d_, "setCalendar:", value)
}

func (d_ DatePickerCell) Locale() foundation.Locale {
	rv := ffi.CallMethod[foundation.Locale](d_, "locale")
	return rv
}

func (d_ DatePickerCell) SetLocale(value foundation.ILocale) {
	ffi.CallMethod[ffi.Void](d_, "setLocale:", value)
}

func (d_ DatePickerCell) TimeZone() foundation.TimeZone {
	rv := ffi.CallMethod[foundation.TimeZone](d_, "timeZone")
	return rv
}

func (d_ DatePickerCell) SetTimeZone(value foundation.ITimeZone) {
	ffi.CallMethod[ffi.Void](d_, "setTimeZone:", value)
}

func (d_ DatePickerCell) MinDate() foundation.Date {
	rv := ffi.CallMethod[foundation.Date](d_, "minDate")
	return rv
}

func (d_ DatePickerCell) SetMinDate(value foundation.IDate) {
	ffi.CallMethod[ffi.Void](d_, "setMinDate:", value)
}

func (d_ DatePickerCell) MaxDate() foundation.Date {
	rv := ffi.CallMethod[foundation.Date](d_, "maxDate")
	return rv
}

func (d_ DatePickerCell) SetMaxDate(value foundation.IDate) {
	ffi.CallMethod[ffi.Void](d_, "setMaxDate:", value)
}

func (d_ DatePickerCell) Delegate() DatePickerCellDelegateWrapper {
	rv := ffi.CallMethod[DatePickerCellDelegateWrapper](d_, "delegate")
	return rv
}

func (d_ DatePickerCell) SetDelegate(value DatePickerCellDelegate) {
	po := ffi.CreateProtocol("NSDatePickerCellDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(d_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](d_, "setDelegate:", po)
}

func (d_ DatePickerCell) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "setDelegate:", value)
}
