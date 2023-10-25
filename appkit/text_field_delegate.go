// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type TextFieldDelegate interface {
	ControlTextEditingDelegate
	ImplementsTextField_TextView_Candidates_ForSelectedRange() bool
	// optional
	TextField_TextView_Candidates_ForSelectedRange(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	ImplementsTextField_TextView_CandidatesForSelectedRange() bool
	// optional
	TextField_TextView_CandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject
	ImplementsTextField_TextView_ShouldSelectCandidateAtIndex() bool
	// optional
	TextField_TextView_ShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool
}

func WrapTextFieldDelegate(v TextFieldDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTextFieldDelegate", v)
}

type TextFieldDelegateBase struct {
	ControlTextEditingDelegateBase
}

func (p *TextFieldDelegateBase) ImplementsTextField_TextView_Candidates_ForSelectedRange() bool {
	return false
}

func (p *TextFieldDelegateBase) TextField_TextView_Candidates_ForSelectedRange(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult {
	panic("unimpemented protocol method")
}

func (p *TextFieldDelegateBase) ImplementsTextField_TextView_CandidatesForSelectedRange() bool {
	return false
}

func (p *TextFieldDelegateBase) TextField_TextView_CandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject {
	panic("unimpemented protocol method")
}

func (p *TextFieldDelegateBase) ImplementsTextField_TextView_ShouldSelectCandidateAtIndex() bool {
	return false
}

func (p *TextFieldDelegateBase) TextField_TextView_ShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool {
	panic("unimpemented protocol method")
}

type TextFieldDelegateCreator struct {
	className string
	class     objc.Class
}

func NewTextFieldDelegateCreator(name string) *TextFieldDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &TextFieldDelegateCreator{className: name, class: class}
}

func (c *TextFieldDelegateCreator) SetTextField_TextView_Candidates_ForSelectedRange(handle func(o objc.Object, textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult) {
	objc.AddMethod(c.class, objc.GetSelector("textField:textView:candidates:forSelectedRange:"), handle)
}

func (c *TextFieldDelegateCreator) SetTextField_TextView_CandidatesForSelectedRange(handle func(o objc.Object, textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject) {
	objc.AddMethod(c.class, objc.GetSelector("textField:textView:candidatesForSelectedRange:"), handle)
}

func (c *TextFieldDelegateCreator) SetTextField_TextView_ShouldSelectCandidateAtIndex(handle func(o objc.Object, textField TextField, textView TextView, index uint) bool) {
	objc.AddMethod(c.class, objc.GetSelector("textField:textView:shouldSelectCandidateAtIndex:"), handle)
}

func (c *TextFieldDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
