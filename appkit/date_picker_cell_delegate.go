// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type DatePickerCellDelegate interface {
	ImplementsDatePickerCell_ValidateProposedDateValue_TimeInterval() bool
	// optional
	DatePickerCell_ValidateProposedDateValue_TimeInterval(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval)
}

type DatePickerCellDelegateImpl struct {
	_DatePickerCell_ValidateProposedDateValue_TimeInterval func(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval)
}

func (di *DatePickerCellDelegateImpl) ImplementsDatePickerCell_ValidateProposedDateValue_TimeInterval() bool {
	return di._DatePickerCell_ValidateProposedDateValue_TimeInterval != nil
}

func (di *DatePickerCellDelegateImpl) SetDatePickerCell_ValidateProposedDateValue_TimeInterval(f func(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval)) {
	di._DatePickerCell_ValidateProposedDateValue_TimeInterval = f
}

func (di *DatePickerCellDelegateImpl) DatePickerCell_ValidateProposedDateValue_TimeInterval(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval) {
	di._DatePickerCell_ValidateProposedDateValue_TimeInterval(datePickerCell, proposedDateValue, proposedTimeInterval)
}

type DatePickerCellDelegateWrapper struct {
	objc.Object
}

func (d_ *DatePickerCellDelegateWrapper) ImplementsDatePickerCell_ValidateProposedDateValue_TimeInterval() bool {
	return d_.RespondsToSelector(objc.GetSelector("datePickerCell:validateProposedDateValue:timeInterval:"))
}

func (d_ DatePickerCellDelegateWrapper) DatePickerCell_ValidateProposedDateValue_TimeInterval(datePickerCell IDatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](d_, "datePickerCell:validateProposedDateValue:timeInterval:", datePickerCell, unsafe.Pointer(proposedDateValue), proposedTimeInterval)
}
