// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type RuleEditorDelegate interface {
	// required
	RuleEditor_Child_ForCriterion_WithRowType(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.IObject
	// required
	RuleEditor_DisplayValueForCriterion_InRow(editor RuleEditor, criterion objc.Object, row int) objc.IObject
	// required
	RuleEditor_NumberOfChildrenForCriterion_WithRowType(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int
	ImplementsRuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow() bool
	// optional
	RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.IObject
	ImplementsRuleEditorRowsDidChange() bool
	// optional
	RuleEditorRowsDidChange(notification foundation.Notification)
}

type RuleEditorDelegateImpl struct {
	_RuleEditor_Child_ForCriterion_WithRowType                    func(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.IObject
	_RuleEditor_DisplayValueForCriterion_InRow                    func(editor RuleEditor, criterion objc.Object, row int) objc.IObject
	_RuleEditor_NumberOfChildrenForCriterion_WithRowType          func(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int
	_RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow func(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.IObject
	_RuleEditorRowsDidChange                                      func(notification foundation.Notification)
}

func (di *RuleEditorDelegateImpl) SetRuleEditor_Child_ForCriterion_WithRowType(f func(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.IObject) {
	di._RuleEditor_Child_ForCriterion_WithRowType = f
}

// required

func (di *RuleEditorDelegateImpl) RuleEditor_Child_ForCriterion_WithRowType(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.IObject {
	return di._RuleEditor_Child_ForCriterion_WithRowType(editor, index, criterion, rowType)
}

func (di *RuleEditorDelegateImpl) SetRuleEditor_DisplayValueForCriterion_InRow(f func(editor RuleEditor, criterion objc.Object, row int) objc.IObject) {
	di._RuleEditor_DisplayValueForCriterion_InRow = f
}

// required

func (di *RuleEditorDelegateImpl) RuleEditor_DisplayValueForCriterion_InRow(editor RuleEditor, criterion objc.Object, row int) objc.IObject {
	return di._RuleEditor_DisplayValueForCriterion_InRow(editor, criterion, row)
}

func (di *RuleEditorDelegateImpl) SetRuleEditor_NumberOfChildrenForCriterion_WithRowType(f func(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int) {
	di._RuleEditor_NumberOfChildrenForCriterion_WithRowType = f
}

// required

func (di *RuleEditorDelegateImpl) RuleEditor_NumberOfChildrenForCriterion_WithRowType(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int {
	return di._RuleEditor_NumberOfChildrenForCriterion_WithRowType(editor, criterion, rowType)
}
func (di *RuleEditorDelegateImpl) ImplementsRuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow() bool {
	return di._RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow != nil
}

func (di *RuleEditorDelegateImpl) SetRuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow(f func(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.IObject) {
	di._RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow = f
}

func (di *RuleEditorDelegateImpl) RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.IObject {
	return di._RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow(editor, criterion, value, row)
}
func (di *RuleEditorDelegateImpl) ImplementsRuleEditorRowsDidChange() bool {
	return di._RuleEditorRowsDidChange != nil
}

func (di *RuleEditorDelegateImpl) SetRuleEditorRowsDidChange(f func(notification foundation.Notification)) {
	di._RuleEditorRowsDidChange = f
}

func (di *RuleEditorDelegateImpl) RuleEditorRowsDidChange(notification foundation.Notification) {
	di._RuleEditorRowsDidChange(notification)
}

type RuleEditorDelegateWrapper struct {
	objc.Object
}

func (r_ RuleEditorDelegateWrapper) RuleEditor_Child_ForCriterion_WithRowType(editor IRuleEditor, index int, criterion objc.IObject, rowType RuleEditorRowType) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.GetSelector("ruleEditor:child:forCriterion:withRowType:"), objc.ExtractPtr(editor), index, objc.ExtractPtr(criterion), rowType)
	return rv
}

func (r_ RuleEditorDelegateWrapper) RuleEditor_DisplayValueForCriterion_InRow(editor IRuleEditor, criterion objc.IObject, row int) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.GetSelector("ruleEditor:displayValueForCriterion:inRow:"), objc.ExtractPtr(editor), objc.ExtractPtr(criterion), row)
	return rv
}

func (r_ RuleEditorDelegateWrapper) RuleEditor_NumberOfChildrenForCriterion_WithRowType(editor IRuleEditor, criterion objc.IObject, rowType RuleEditorRowType) int {
	rv := objc.CallMethod[int](r_, objc.GetSelector("ruleEditor:numberOfChildrenForCriterion:withRowType:"), objc.ExtractPtr(editor), objc.ExtractPtr(criterion), rowType)
	return rv
}

func (r_ *RuleEditorDelegateWrapper) ImplementsRuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow() bool {
	return r_.RespondsToSelector(objc.GetSelector("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"))
}

func (r_ RuleEditorDelegateWrapper) RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow(editor IRuleEditor, criterion objc.IObject, value objc.IObject, row int) map[RuleEditorPredicatePartKey]objc.Object {
	rv := objc.CallMethod[map[RuleEditorPredicatePartKey]objc.Object](r_, objc.GetSelector("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"), objc.ExtractPtr(editor), objc.ExtractPtr(criterion), objc.ExtractPtr(value), row)
	return rv
}

func (r_ *RuleEditorDelegateWrapper) ImplementsRuleEditorRowsDidChange() bool {
	return r_.RespondsToSelector(objc.GetSelector("ruleEditorRowsDidChange:"))
}

func (r_ RuleEditorDelegateWrapper) RuleEditorRowsDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("ruleEditorRowsDidChange:"), objc.ExtractPtr(notification))
}
