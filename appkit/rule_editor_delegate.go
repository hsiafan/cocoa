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

func WrapRuleEditorDelegate(v RuleEditorDelegate) objc.Object {
	return objc.WrapAsProtocol("NSRuleEditorDelegate", v)
}

type RuleEditorDelegateBase struct {
}

func (p *RuleEditorDelegateBase) ImplementsRuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow() bool {
	return false
}

func (p *RuleEditorDelegateBase) RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.IObject {
	panic("unimpemented protocol method")
}

func (p *RuleEditorDelegateBase) ImplementsRuleEditorRowsDidChange() bool {
	return false
}

func (p *RuleEditorDelegateBase) RuleEditorRowsDidChange(notification foundation.Notification) {
	panic("unimpemented protocol method")
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

func (r_ RuleEditorDelegateWrapper) RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow(editor IRuleEditor, criterion objc.IObject, value objc.IObject, row int) map[RuleEditorPredicatePartKey]objc.Object {
	rv := objc.CallMethod[map[RuleEditorPredicatePartKey]objc.Object](r_, objc.GetSelector("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"), objc.ExtractPtr(editor), objc.ExtractPtr(criterion), objc.ExtractPtr(value), row)
	return rv
}

func (r_ RuleEditorDelegateWrapper) RuleEditorRowsDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("ruleEditorRowsDidChange:"), objc.ExtractPtr(notification))
}
