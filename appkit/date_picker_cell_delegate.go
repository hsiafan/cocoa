// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type DatePickerCellDelegate interface {
	ImplementsDatePickerCell_ValidateProposedDateValue_TimeInterval() bool
	// optional
	DatePickerCell_ValidateProposedDateValue_TimeInterval(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval)
}

func WrapDatePickerCellDelegate(v DatePickerCellDelegate) objc.Object {
	return objc.WrapAsProtocol("NSDatePickerCellDelegate", v)
}

type DatePickerCellDelegateBase struct {
}

func (p *DatePickerCellDelegateBase) ImplementsDatePickerCell_ValidateProposedDateValue_TimeInterval() bool {
	return false
}

func (p *DatePickerCellDelegateBase) DatePickerCell_ValidateProposedDateValue_TimeInterval(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval) {
	panic("unimpemented protocol method")
}

type DatePickerCellDelegateWrapper struct {
	objc.Object
}

func (d_ DatePickerCellDelegateWrapper) DatePickerCell_ValidateProposedDateValue_TimeInterval(datePickerCell IDatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("datePickerCell:validateProposedDateValue:timeInterval:"), objc.ExtractPtr(datePickerCell), unsafe.Pointer(proposedDateValue), proposedTimeInterval)
}
