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
