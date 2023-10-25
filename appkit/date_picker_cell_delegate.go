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

type DatePickerCellDelegateCreator struct {
	className string
	class     objc.Class
}

func NewDatePickerCellDelegateCreator(name string) *DatePickerCellDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &DatePickerCellDelegateCreator{className: name, class: class}
}

func (c *DatePickerCellDelegateCreator) SetDatePickerCell_ValidateProposedDateValue_TimeInterval(handle func(o objc.Object, datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval)) {
	objc.AddMethod(c.class, objc.GetSelector("datePickerCell:validateProposedDateValue:timeInterval:"), handle)
}

func (c *DatePickerCellDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
