// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[RuleEditor](r_, "initWithFrame:", frameRect)
	return rv
}

func (r_ RuleEditor) Init() RuleEditor {
	rv := ffi.CallMethod[RuleEditor](r_, "init")
	return rv
}

func (rc _RuleEditorClass) Alloc() RuleEditor {
	rv := ffi.CallMethod[RuleEditor](rc, "alloc")
	return rv
}

func (rc _RuleEditorClass) New() RuleEditor {
	rv := ffi.CallMethod[RuleEditor](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRuleEditor() RuleEditor {
	return RuleEditorClass.New()
}

func (r_ RuleEditor) ReloadCriteria() {
	ffi.CallMethod[ffi.Void](r_, "reloadCriteria")
}

func (r_ RuleEditor) SetCriteria_AndDisplayValues_ForRowAtIndex(criteria []objc.IObject, values []objc.IObject, rowIndex int) {
	ffi.CallMethod[ffi.Void](r_, "setCriteria:andDisplayValues:forRowAtIndex:", criteria, values, rowIndex)
}

func (r_ RuleEditor) CriteriaForRow(row int) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](r_, "criteriaForRow:", row)
	return rv
}

func (r_ RuleEditor) DisplayValuesForRow(row int) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](r_, "displayValuesForRow:", row)
	return rv
}

func (r_ RuleEditor) ParentRowForRow(rowIndex int) int {
	rv := ffi.CallMethod[int](r_, "parentRowForRow:", rowIndex)
	return rv
}

func (r_ RuleEditor) RowForDisplayValue(displayValue objc.IObject) int {
	rv := ffi.CallMethod[int](r_, "rowForDisplayValue:", displayValue)
	return rv
}

func (r_ RuleEditor) RowTypeForRow(rowIndex int) RuleEditorRowType {
	rv := ffi.CallMethod[RuleEditorRowType](r_, "rowTypeForRow:", rowIndex)
	return rv
}

func (r_ RuleEditor) SubrowIndexesForRow(rowIndex int) foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](r_, "subrowIndexesForRow:", rowIndex)
	return rv
}

func (r_ RuleEditor) SelectRowIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	ffi.CallMethod[ffi.Void](r_, "selectRowIndexes:byExtendingSelection:", indexes, extend)
}

func (r_ RuleEditor) AddRow(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](r_, "addRow:", sender)
}

func (r_ RuleEditor) InsertRowAtIndex_WithType_AsSubrowOfRow_Animate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool) {
	ffi.CallMethod[ffi.Void](r_, "insertRowAtIndex:withType:asSubrowOfRow:animate:", rowIndex, rowType, parentRow, shouldAnimate)
}

func (r_ RuleEditor) RemoveRowAtIndex(rowIndex int) {
	ffi.CallMethod[ffi.Void](r_, "removeRowAtIndex:", rowIndex)
}

func (r_ RuleEditor) RemoveRowsAtIndexes_IncludeSubrows(rowIndexes foundation.IIndexSet, includeSubrows bool) {
	ffi.CallMethod[ffi.Void](r_, "removeRowsAtIndexes:includeSubrows:", rowIndexes, includeSubrows)
}

func (r_ RuleEditor) ReloadPredicate() {
	ffi.CallMethod[ffi.Void](r_, "reloadPredicate")
}

func (r_ RuleEditor) PredicateForRow(row int) foundation.Predicate {
	rv := ffi.CallMethod[foundation.Predicate](r_, "predicateForRow:", row)
	return rv
}

func (r_ RuleEditor) Delegate() RuleEditorDelegateWrapper {
	rv := ffi.CallMethod[RuleEditorDelegateWrapper](r_, "delegate")
	return rv
}

func (r_ RuleEditor) SetDelegate(value RuleEditorDelegate) {
	po := ffi.CreateProtocol("NSRuleEditorDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(r_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](r_, "setDelegate:", po)
}

func (r_ RuleEditor) SetDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](r_, "setDelegate:", value)
}

func (r_ RuleEditor) IsEditable() bool {
	rv := ffi.CallMethod[bool](r_, "isEditable")
	return rv
}

func (r_ RuleEditor) SetEditable(value bool) {
	ffi.CallMethod[ffi.Void](r_, "setEditable:", value)
}

func (r_ RuleEditor) NestingMode() RuleEditorNestingMode {
	rv := ffi.CallMethod[RuleEditorNestingMode](r_, "nestingMode")
	return rv
}

func (r_ RuleEditor) SetNestingMode(value RuleEditorNestingMode) {
	ffi.CallMethod[ffi.Void](r_, "setNestingMode:", value)
}

func (r_ RuleEditor) CanRemoveAllRows() bool {
	rv := ffi.CallMethod[bool](r_, "canRemoveAllRows")
	return rv
}

func (r_ RuleEditor) SetCanRemoveAllRows(value bool) {
	ffi.CallMethod[ffi.Void](r_, "setCanRemoveAllRows:", value)
}

func (r_ RuleEditor) RowHeight() float64 {
	rv := ffi.CallMethod[float64](r_, "rowHeight")
	return rv
}

func (r_ RuleEditor) SetRowHeight(value float64) {
	ffi.CallMethod[ffi.Void](r_, "setRowHeight:", value)
}

func (r_ RuleEditor) FormattingDictionary() map[string]string {
	rv := ffi.CallMethod[map[string]string](r_, "formattingDictionary")
	return rv
}

func (r_ RuleEditor) SetFormattingDictionary(value map[string]string) {
	ffi.CallMethod[ffi.Void](r_, "setFormattingDictionary:", value)
}

func (r_ RuleEditor) FormattingStringsFilename() string {
	rv := ffi.CallMethod[string](r_, "formattingStringsFilename")
	return rv
}

func (r_ RuleEditor) SetFormattingStringsFilename(value string) {
	ffi.CallMethod[ffi.Void](r_, "setFormattingStringsFilename:", value)
}

func (r_ RuleEditor) NumberOfRows() int {
	rv := ffi.CallMethod[int](r_, "numberOfRows")
	return rv
}

func (r_ RuleEditor) SelectedRowIndexes() foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](r_, "selectedRowIndexes")
	return rv
}

func (r_ RuleEditor) Predicate() foundation.Predicate {
	rv := ffi.CallMethod[foundation.Predicate](r_, "predicate")
	return rv
}

func (r_ RuleEditor) RowClass() objc.Class {
	rv := ffi.CallMethod[objc.Class](r_, "rowClass")
	return rv
}

func (r_ RuleEditor) SetRowClass(value objc.IClass) {
	ffi.CallMethod[ffi.Void](r_, "setRowClass:", value)
}

func (r_ RuleEditor) RowTypeKeyPath() string {
	rv := ffi.CallMethod[string](r_, "rowTypeKeyPath")
	return rv
}

func (r_ RuleEditor) SetRowTypeKeyPath(value string) {
	ffi.CallMethod[ffi.Void](r_, "setRowTypeKeyPath:", value)
}

func (r_ RuleEditor) SubrowsKeyPath() string {
	rv := ffi.CallMethod[string](r_, "subrowsKeyPath")
	return rv
}

func (r_ RuleEditor) SetSubrowsKeyPath(value string) {
	ffi.CallMethod[ffi.Void](r_, "setSubrowsKeyPath:", value)
}

func (r_ RuleEditor) CriteriaKeyPath() string {
	rv := ffi.CallMethod[string](r_, "criteriaKeyPath")
	return rv
}

func (r_ RuleEditor) SetCriteriaKeyPath(value string) {
	ffi.CallMethod[ffi.Void](r_, "setCriteriaKeyPath:", value)
}

func (r_ RuleEditor) DisplayValuesKeyPath() string {
	rv := ffi.CallMethod[string](r_, "displayValuesKeyPath")
	return rv
}

func (r_ RuleEditor) SetDisplayValuesKeyPath(value string) {
	ffi.CallMethod[ffi.Void](r_, "setDisplayValuesKeyPath:", value)
}
