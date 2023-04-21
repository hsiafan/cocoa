// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var RuleEditorClass = _RuleEditorClass{objc.GetClass("NSRuleEditor")}

type _RuleEditorClass struct {
	objc.Class
}

type IRuleEditor interface {
	IControl
	ReloadCriteria()
	SetCriteria_AndDisplayValues_ForRowAtIndex(criteria []objc.IObject, values []objc.IObject, rowIndex int)
	CriteriaForRow(row int) []objc.Object
	DisplayValuesForRow(row int) []objc.Object
	ParentRowForRow(rowIndex int) int
	RowForDisplayValue(displayValue objc.IObject) int
	RowTypeForRow(rowIndex int) RuleEditorRowType
	SubrowIndexesForRow(rowIndex int) foundation.IndexSet
	SelectRowIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool)
	AddRow(sender objc.IObject)
	InsertRowAtIndex_WithType_AsSubrowOfRow_Animate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool)
	RemoveRowAtIndex(rowIndex int)
	RemoveRowsAtIndexes_IncludeSubrows(rowIndexes foundation.IIndexSet, includeSubrows bool)
	ReloadPredicate()
	PredicateForRow(row int) foundation.Predicate
	Delegate() RuleEditorDelegateWrapper
	SetDelegate(value RuleEditorDelegate)
	SetDelegate0(value objc.IObject)
	IsEditable() bool
	SetEditable(value bool)
	NestingMode() RuleEditorNestingMode
	SetNestingMode(value RuleEditorNestingMode)
	CanRemoveAllRows() bool
	SetCanRemoveAllRows(value bool)
	RowHeight() float64
	SetRowHeight(value float64)
	FormattingDictionary() map[string]string
	SetFormattingDictionary(value map[string]string)
	FormattingStringsFilename() string
	SetFormattingStringsFilename(value string)
	NumberOfRows() int
	SelectedRowIndexes() foundation.IndexSet
	Predicate() foundation.Predicate
	RowClass() objc.Class
	SetRowClass(value objc.IClass)
	RowTypeKeyPath() string
	SetRowTypeKeyPath(value string)
	SubrowsKeyPath() string
	SetSubrowsKeyPath(value string)
	CriteriaKeyPath() string
	SetCriteriaKeyPath(value string)
	DisplayValuesKeyPath() string
	SetDisplayValuesKeyPath(value string)
}

type RuleEditor struct {
	Control
}

func MakeRuleEditor(ptr unsafe.Pointer) RuleEditor {
	return RuleEditor{
		Control: MakeControl(ptr),
	}
}

func (r_ RuleEditor) InitWithFrame(frameRect foundation.Rect) RuleEditor {
	rv := objc.CallMethod[RuleEditor](r_, "initWithFrame:", frameRect)
	return rv
}

func (r_ RuleEditor) Init() RuleEditor {
	rv := objc.CallMethod[RuleEditor](r_, "init")
	return rv
}

func (rc _RuleEditorClass) Alloc() RuleEditor {
	rv := objc.CallMethod[RuleEditor](rc, "alloc")
	return rv
}

func (rc _RuleEditorClass) New() RuleEditor {
	rv := objc.CallMethod[RuleEditor](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRuleEditor() RuleEditor {
	return RuleEditorClass.New()
}

func (r_ RuleEditor) ReloadCriteria() {
	objc.CallMethod[objc.Void](r_, "reloadCriteria")
}

func (r_ RuleEditor) SetCriteria_AndDisplayValues_ForRowAtIndex(criteria []objc.IObject, values []objc.IObject, rowIndex int) {
	objc.CallMethod[objc.Void](r_, "setCriteria:andDisplayValues:forRowAtIndex:", criteria, values, rowIndex)
}

func (r_ RuleEditor) CriteriaForRow(row int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](r_, "criteriaForRow:", row)
	return rv
}

func (r_ RuleEditor) DisplayValuesForRow(row int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](r_, "displayValuesForRow:", row)
	return rv
}

func (r_ RuleEditor) ParentRowForRow(rowIndex int) int {
	rv := objc.CallMethod[int](r_, "parentRowForRow:", rowIndex)
	return rv
}

func (r_ RuleEditor) RowForDisplayValue(displayValue objc.IObject) int {
	rv := objc.CallMethod[int](r_, "rowForDisplayValue:", displayValue)
	return rv
}

func (r_ RuleEditor) RowTypeForRow(rowIndex int) RuleEditorRowType {
	rv := objc.CallMethod[RuleEditorRowType](r_, "rowTypeForRow:", rowIndex)
	return rv
}

func (r_ RuleEditor) SubrowIndexesForRow(rowIndex int) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](r_, "subrowIndexesForRow:", rowIndex)
	return rv
}

func (r_ RuleEditor) SelectRowIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.CallMethod[objc.Void](r_, "selectRowIndexes:byExtendingSelection:", indexes, extend)
}

func (r_ RuleEditor) AddRow(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, "addRow:", sender)
}

func (r_ RuleEditor) InsertRowAtIndex_WithType_AsSubrowOfRow_Animate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool) {
	objc.CallMethod[objc.Void](r_, "insertRowAtIndex:withType:asSubrowOfRow:animate:", rowIndex, rowType, parentRow, shouldAnimate)
}

func (r_ RuleEditor) RemoveRowAtIndex(rowIndex int) {
	objc.CallMethod[objc.Void](r_, "removeRowAtIndex:", rowIndex)
}

func (r_ RuleEditor) RemoveRowsAtIndexes_IncludeSubrows(rowIndexes foundation.IIndexSet, includeSubrows bool) {
	objc.CallMethod[objc.Void](r_, "removeRowsAtIndexes:includeSubrows:", rowIndexes, includeSubrows)
}

func (r_ RuleEditor) ReloadPredicate() {
	objc.CallMethod[objc.Void](r_, "reloadPredicate")
}

func (r_ RuleEditor) PredicateForRow(row int) foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](r_, "predicateForRow:", row)
	return rv
}

func (r_ RuleEditor) Delegate() RuleEditorDelegateWrapper {
	rv := objc.CallMethod[RuleEditorDelegateWrapper](r_, "delegate")
	return rv
}

func (r_ RuleEditor) SetDelegate(value RuleEditorDelegate) {
	po := objc.CreateProtocol("NSRuleEditorDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(r_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](r_, "setDelegate:", po)
}

func (r_ RuleEditor) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](r_, "setDelegate:", value)
}

func (r_ RuleEditor) IsEditable() bool {
	rv := objc.CallMethod[bool](r_, "isEditable")
	return rv
}

func (r_ RuleEditor) SetEditable(value bool) {
	objc.CallMethod[objc.Void](r_, "setEditable:", value)
}

func (r_ RuleEditor) NestingMode() RuleEditorNestingMode {
	rv := objc.CallMethod[RuleEditorNestingMode](r_, "nestingMode")
	return rv
}

func (r_ RuleEditor) SetNestingMode(value RuleEditorNestingMode) {
	objc.CallMethod[objc.Void](r_, "setNestingMode:", value)
}

func (r_ RuleEditor) CanRemoveAllRows() bool {
	rv := objc.CallMethod[bool](r_, "canRemoveAllRows")
	return rv
}

func (r_ RuleEditor) SetCanRemoveAllRows(value bool) {
	objc.CallMethod[objc.Void](r_, "setCanRemoveAllRows:", value)
}

func (r_ RuleEditor) RowHeight() float64 {
	rv := objc.CallMethod[float64](r_, "rowHeight")
	return rv
}

func (r_ RuleEditor) SetRowHeight(value float64) {
	objc.CallMethod[objc.Void](r_, "setRowHeight:", value)
}

func (r_ RuleEditor) FormattingDictionary() map[string]string {
	rv := objc.CallMethod[map[string]string](r_, "formattingDictionary")
	return rv
}

func (r_ RuleEditor) SetFormattingDictionary(value map[string]string) {
	objc.CallMethod[objc.Void](r_, "setFormattingDictionary:", value)
}

func (r_ RuleEditor) FormattingStringsFilename() string {
	rv := objc.CallMethod[string](r_, "formattingStringsFilename")
	return rv
}

func (r_ RuleEditor) SetFormattingStringsFilename(value string) {
	objc.CallMethod[objc.Void](r_, "setFormattingStringsFilename:", value)
}

func (r_ RuleEditor) NumberOfRows() int {
	rv := objc.CallMethod[int](r_, "numberOfRows")
	return rv
}

func (r_ RuleEditor) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](r_, "selectedRowIndexes")
	return rv
}

func (r_ RuleEditor) Predicate() foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](r_, "predicate")
	return rv
}

func (r_ RuleEditor) RowClass() objc.Class {
	rv := objc.CallMethod[objc.Class](r_, "rowClass")
	return rv
}

func (r_ RuleEditor) SetRowClass(value objc.IClass) {
	objc.CallMethod[objc.Void](r_, "setRowClass:", value)
}

func (r_ RuleEditor) RowTypeKeyPath() string {
	rv := objc.CallMethod[string](r_, "rowTypeKeyPath")
	return rv
}

func (r_ RuleEditor) SetRowTypeKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, "setRowTypeKeyPath:", value)
}

func (r_ RuleEditor) SubrowsKeyPath() string {
	rv := objc.CallMethod[string](r_, "subrowsKeyPath")
	return rv
}

func (r_ RuleEditor) SetSubrowsKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, "setSubrowsKeyPath:", value)
}

func (r_ RuleEditor) CriteriaKeyPath() string {
	rv := objc.CallMethod[string](r_, "criteriaKeyPath")
	return rv
}

func (r_ RuleEditor) SetCriteriaKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, "setCriteriaKeyPath:", value)
}

func (r_ RuleEditor) DisplayValuesKeyPath() string {
	rv := objc.CallMethod[string](r_, "displayValuesKeyPath")
	return rv
}

func (r_ RuleEditor) SetDisplayValuesKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, "setDisplayValuesKeyPath:", value)
}
