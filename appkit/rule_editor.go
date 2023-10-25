// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[RuleEditor](r_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (r_ RuleEditor) Init() RuleEditor {
	rv := objc.CallMethod[RuleEditor](r_, objc.SEL("init"))
	return rv
}

func (rc _RuleEditorClass) Alloc() RuleEditor {
	rv := objc.CallMethod[RuleEditor](rc, objc.SEL("alloc"))
	return rv
}

func (rc _RuleEditorClass) New() RuleEditor {
	rv := objc.CallMethod[RuleEditor](rc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewRuleEditor() RuleEditor {
	return RuleEditorClass.New()
}

func (r_ RuleEditor) ReloadCriteria() {
	objc.CallMethod[objc.Void](r_, objc.SEL("reloadCriteria"))
}

func (r_ RuleEditor) SetCriteria_AndDisplayValues_ForRowAtIndex(criteria []objc.IObject, values []objc.IObject, rowIndex int) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setCriteria:andDisplayValues:forRowAtIndex:"), criteria, values, rowIndex)
}

func (r_ RuleEditor) CriteriaForRow(row int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](r_, objc.SEL("criteriaForRow:"), row)
	return rv
}

func (r_ RuleEditor) DisplayValuesForRow(row int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](r_, objc.SEL("displayValuesForRow:"), row)
	return rv
}

func (r_ RuleEditor) ParentRowForRow(rowIndex int) int {
	rv := objc.CallMethod[int](r_, objc.SEL("parentRowForRow:"), rowIndex)
	return rv
}

func (r_ RuleEditor) RowForDisplayValue(displayValue objc.IObject) int {
	rv := objc.CallMethod[int](r_, objc.SEL("rowForDisplayValue:"), objc.ExtractPtr(displayValue))
	return rv
}

func (r_ RuleEditor) RowTypeForRow(rowIndex int) RuleEditorRowType {
	rv := objc.CallMethod[RuleEditorRowType](r_, objc.SEL("rowTypeForRow:"), rowIndex)
	return rv
}

func (r_ RuleEditor) SubrowIndexesForRow(rowIndex int) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](r_, objc.SEL("subrowIndexesForRow:"), rowIndex)
	return rv
}

func (r_ RuleEditor) SelectRowIndexes_ByExtendingSelection(indexes foundation.IIndexSet, extend bool) {
	objc.CallMethod[objc.Void](r_, objc.SEL("selectRowIndexes:byExtendingSelection:"), objc.ExtractPtr(indexes), extend)
}

func (r_ RuleEditor) AddRow(sender objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.SEL("addRow:"), objc.ExtractPtr(sender))
}

func (r_ RuleEditor) InsertRowAtIndex_WithType_AsSubrowOfRow_Animate(rowIndex int, rowType RuleEditorRowType, parentRow int, shouldAnimate bool) {
	objc.CallMethod[objc.Void](r_, objc.SEL("insertRowAtIndex:withType:asSubrowOfRow:animate:"), rowIndex, rowType, parentRow, shouldAnimate)
}

func (r_ RuleEditor) RemoveRowAtIndex(rowIndex int) {
	objc.CallMethod[objc.Void](r_, objc.SEL("removeRowAtIndex:"), rowIndex)
}

func (r_ RuleEditor) RemoveRowsAtIndexes_IncludeSubrows(rowIndexes foundation.IIndexSet, includeSubrows bool) {
	objc.CallMethod[objc.Void](r_, objc.SEL("removeRowsAtIndexes:includeSubrows:"), objc.ExtractPtr(rowIndexes), includeSubrows)
}

func (r_ RuleEditor) ReloadPredicate() {
	objc.CallMethod[objc.Void](r_, objc.SEL("reloadPredicate"))
}

func (r_ RuleEditor) PredicateForRow(row int) foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](r_, objc.SEL("predicateForRow:"), row)
	return rv
}

// weak property
func (r_ RuleEditor) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.SEL("delegate"))
	return rv
}

// weak property
func (r_ RuleEditor) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (r_ RuleEditor) IsEditable() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("isEditable"))
	return rv
}

func (r_ RuleEditor) SetEditable(value bool) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setEditable:"), value)
}

func (r_ RuleEditor) NestingMode() RuleEditorNestingMode {
	rv := objc.CallMethod[RuleEditorNestingMode](r_, objc.SEL("nestingMode"))
	return rv
}

func (r_ RuleEditor) SetNestingMode(value RuleEditorNestingMode) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setNestingMode:"), value)
}

func (r_ RuleEditor) CanRemoveAllRows() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("canRemoveAllRows"))
	return rv
}

func (r_ RuleEditor) SetCanRemoveAllRows(value bool) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setCanRemoveAllRows:"), value)
}

func (r_ RuleEditor) RowHeight() float64 {
	rv := objc.CallMethod[float64](r_, objc.SEL("rowHeight"))
	return rv
}

func (r_ RuleEditor) SetRowHeight(value float64) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setRowHeight:"), value)
}

func (r_ RuleEditor) FormattingDictionary() map[string]string {
	rv := objc.CallMethod[map[string]string](r_, objc.SEL("formattingDictionary"))
	return rv
}

func (r_ RuleEditor) SetFormattingDictionary(value map[string]string) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setFormattingDictionary:"), value)
}

func (r_ RuleEditor) FormattingStringsFilename() string {
	rv := objc.CallMethod[string](r_, objc.SEL("formattingStringsFilename"))
	return rv
}

func (r_ RuleEditor) SetFormattingStringsFilename(value string) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setFormattingStringsFilename:"), value)
}

func (r_ RuleEditor) NumberOfRows() int {
	rv := objc.CallMethod[int](r_, objc.SEL("numberOfRows"))
	return rv
}

func (r_ RuleEditor) SelectedRowIndexes() foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](r_, objc.SEL("selectedRowIndexes"))
	return rv
}

func (r_ RuleEditor) Predicate() foundation.Predicate {
	rv := objc.CallMethod[foundation.Predicate](r_, objc.SEL("predicate"))
	return rv
}

// weak property
func (r_ RuleEditor) RowClass() objc.Class {
	rv := objc.CallMethod[objc.Class](r_, objc.SEL("rowClass"))
	return rv
}

// weak property
func (r_ RuleEditor) SetRowClass(value objc.IClass) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setRowClass:"), objc.ExtractPtr(value))
}

func (r_ RuleEditor) RowTypeKeyPath() string {
	rv := objc.CallMethod[string](r_, objc.SEL("rowTypeKeyPath"))
	return rv
}

func (r_ RuleEditor) SetRowTypeKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setRowTypeKeyPath:"), value)
}

func (r_ RuleEditor) SubrowsKeyPath() string {
	rv := objc.CallMethod[string](r_, objc.SEL("subrowsKeyPath"))
	return rv
}

func (r_ RuleEditor) SetSubrowsKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setSubrowsKeyPath:"), value)
}

func (r_ RuleEditor) CriteriaKeyPath() string {
	rv := objc.CallMethod[string](r_, objc.SEL("criteriaKeyPath"))
	return rv
}

func (r_ RuleEditor) SetCriteriaKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setCriteriaKeyPath:"), value)
}

func (r_ RuleEditor) DisplayValuesKeyPath() string {
	rv := objc.CallMethod[string](r_, objc.SEL("displayValuesKeyPath"))
	return rv
}

func (r_ RuleEditor) SetDisplayValuesKeyPath(value string) {
	objc.CallMethod[objc.Void](r_, objc.SEL("setDisplayValuesKeyPath:"), value)
}
