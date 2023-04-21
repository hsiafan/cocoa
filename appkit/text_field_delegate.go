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

type TextFieldDelegateImpl struct {
	ControlTextEditingDelegateImpl
	_TextField_TextView_Candidates_ForSelectedRange  func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	_TextField_TextView_CandidatesForSelectedRange   func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject
	_TextField_TextView_ShouldSelectCandidateAtIndex func(textField TextField, textView TextView, index uint) bool
}

func (di *TextFieldDelegateImpl) ImplementsTextField_TextView_Candidates_ForSelectedRange() bool {
	return di._TextField_TextView_Candidates_ForSelectedRange != nil
}

func (di *TextFieldDelegateImpl) SetTextField_TextView_Candidates_ForSelectedRange(f func(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult) {
	di._TextField_TextView_Candidates_ForSelectedRange = f
}

func (di *TextFieldDelegateImpl) TextField_TextView_Candidates_ForSelectedRange(textField TextField, textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult {
	return di._TextField_TextView_Candidates_ForSelectedRange(textField, textView, candidates, selectedRange)
}
func (di *TextFieldDelegateImpl) ImplementsTextField_TextView_CandidatesForSelectedRange() bool {
	return di._TextField_TextView_CandidatesForSelectedRange != nil
}

func (di *TextFieldDelegateImpl) SetTextField_TextView_CandidatesForSelectedRange(f func(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject) {
	di._TextField_TextView_CandidatesForSelectedRange = f
}

func (di *TextFieldDelegateImpl) TextField_TextView_CandidatesForSelectedRange(textField TextField, textView TextView, selectedRange foundation.Range) []objc.IObject {
	return di._TextField_TextView_CandidatesForSelectedRange(textField, textView, selectedRange)
}
func (di *TextFieldDelegateImpl) ImplementsTextField_TextView_ShouldSelectCandidateAtIndex() bool {
	return di._TextField_TextView_ShouldSelectCandidateAtIndex != nil
}

func (di *TextFieldDelegateImpl) SetTextField_TextView_ShouldSelectCandidateAtIndex(f func(textField TextField, textView TextView, index uint) bool) {
	di._TextField_TextView_ShouldSelectCandidateAtIndex = f
}

func (di *TextFieldDelegateImpl) TextField_TextView_ShouldSelectCandidateAtIndex(textField TextField, textView TextView, index uint) bool {
	return di._TextField_TextView_ShouldSelectCandidateAtIndex(textField, textView, index)
}

type TextFieldDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (t_ *TextFieldDelegateWrapper) ImplementsTextField_TextView_Candidates_ForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textField:textView:candidates:forSelectedRange:"))
}

func (t_ TextFieldDelegateWrapper) TextField_TextView_Candidates_ForSelectedRange(textField ITextField, textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, "textField:textView:candidates:forSelectedRange:", textField, textView, candidates, selectedRange)
	return rv
}

func (t_ *TextFieldDelegateWrapper) ImplementsTextField_TextView_CandidatesForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textField:textView:candidatesForSelectedRange:"))
}

func (t_ TextFieldDelegateWrapper) TextField_TextView_CandidatesForSelectedRange(textField ITextField, textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, "textField:textView:candidatesForSelectedRange:", textField, textView, selectedRange)
	return rv
}

func (t_ *TextFieldDelegateWrapper) ImplementsTextField_TextView_ShouldSelectCandidateAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textField:textView:shouldSelectCandidateAtIndex:"))
}

func (t_ TextFieldDelegateWrapper) TextField_TextView_ShouldSelectCandidateAtIndex(textField ITextField, textView ITextView, index uint) bool {
	rv := objc.CallMethod[bool](t_, "textField:textView:shouldSelectCandidateAtIndex:", textField, textView, index)
	return rv
}
