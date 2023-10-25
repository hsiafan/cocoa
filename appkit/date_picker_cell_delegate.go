// auto-generated code, do not modify
package appkit

import (
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
