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
	rv.Autorelease()
	return rv
}

func (r_ RuleEditor) Init() RuleEditor {
	rv := ffi.CallMethod[RuleEditor](r_, "init")
	rv.Autorelease()
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
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(r_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](r_, "setDelegate:", po)
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
	rv := ffi.CallMethod[objc.Object](r_, "ruleEditor:child:forCriterion:withRowType:", editor, index, criterion, rowType)
	return rv
}

func (r_ RuleEditorDelegateWrapper) RuleEditor_DisplayValueForCriterion_InRow(editor IRuleEditor, criterion objc.IObject, row int) objc.Object {
	rv := ffi.CallMethod[objc.Object](r_, "ruleEditor:displayValueForCriterion:inRow:", editor, criterion, row)
	return rv
}

func (r_ RuleEditorDelegateWrapper) RuleEditor_NumberOfChildrenForCriterion_WithRowType(editor IRuleEditor, criterion objc.IObject, rowType RuleEditorRowType) int {
	rv := ffi.CallMethod[int](r_, "ruleEditor:numberOfChildrenForCriterion:withRowType:", editor, criterion, rowType)
	return rv
}

func (r_ *RuleEditorDelegateWrapper) ImplementsRuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow() bool {
	return r_.RespondsToSelector(objc.GetSelector("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"))
}

func (r_ RuleEditorDelegateWrapper) RuleEditor_PredicatePartsForCriterion_WithDisplayValue_InRow(editor IRuleEditor, criterion objc.IObject, value objc.IObject, row int) map[RuleEditorPredicatePartKey]objc.Object {
	rv := ffi.CallMethod[map[RuleEditorPredicatePartKey]objc.Object](r_, "ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:", editor, criterion, value, row)
	return rv
}

func (r_ *RuleEditorDelegateWrapper) ImplementsRuleEditorRowsDidChange() bool {
	return r_.RespondsToSelector(objc.GetSelector("ruleEditorRowsDidChange:"))
}

func (r_ RuleEditorDelegateWrapper) RuleEditorRowsDidChange(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](r_, "ruleEditorRowsDidChange:", notification)
}

var PredicateEditorClass = _PredicateEditorClass{objc.GetClass("NSPredicateEditor")}

type _PredicateEditorClass struct {
	objc.Class
}

type IPredicateEditor interface {
	IRuleEditor
	RowTemplates() []PredicateEditorRowTemplate
	SetRowTemplates(value []IPredicateEditorRowTemplate)
}

type PredicateEditor struct {
	RuleEditor
}

func MakePredicateEditor(ptr unsafe.Pointer) PredicateEditor {
	return PredicateEditor{
		RuleEditor: MakeRuleEditor(ptr),
	}
}

func (p_ PredicateEditor) InitWithFrame(frameRect foundation.Rect) PredicateEditor {
	rv := ffi.CallMethod[PredicateEditor](p_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (p_ PredicateEditor) Init() PredicateEditor {
	rv := ffi.CallMethod[PredicateEditor](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PredicateEditorClass) Alloc() PredicateEditor {
	rv := ffi.CallMethod[PredicateEditor](pc, "alloc")
	return rv
}

func (pc _PredicateEditorClass) New() PredicateEditor {
	rv := ffi.CallMethod[PredicateEditor](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPredicateEditor() PredicateEditor {
	return PredicateEditorClass.New()
}

func (p_ PredicateEditor) RowTemplates() []PredicateEditorRowTemplate {
	rv := ffi.CallMethod[[]PredicateEditorRowTemplate](p_, "rowTemplates")
	return rv
}

func (p_ PredicateEditor) SetRowTemplates(value []IPredicateEditorRowTemplate) {
	ffi.CallMethod[ffi.Void](p_, "setRowTemplates:", value)
}

var PredicateEditorRowTemplateClass = _PredicateEditorRowTemplateClass{objc.GetClass("NSPredicateEditorRowTemplate")}

type _PredicateEditorRowTemplateClass struct {
	objc.Class
}

type IPredicateEditorRowTemplate interface {
	objc.IObject
	MatchForPredicate(predicate foundation.IPredicate) float64
	SetPredicate(predicate foundation.IPredicate)
	DisplayableSubpredicatesOfPredicate(predicate foundation.IPredicate) []foundation.Predicate
	PredicateWithSubpredicates(subpredicates []foundation.IPredicate) foundation.Predicate
	TemplateViews() []View
	LeftExpressions() []foundation.Expression
	RightExpressions() []foundation.Expression
	CompoundTypes() []foundation.Number
	Modifier() foundation.ComparisonPredicateModifier
	Operators() []foundation.Number
	Options() uint
}

type PredicateEditorRowTemplate struct {
	objc.Object
}

func MakePredicateEditorRowTemplate(ptr unsafe.Pointer) PredicateEditorRowTemplate {
	return PredicateEditorRowTemplate{
		Object: objc.MakeObject(ptr),
	}
}

func (p_ PredicateEditorRowTemplate) InitWithLeftExpressions_RightExpressions_Modifier_Operators_Options(leftExpressions []foundation.IExpression, rightExpressions []foundation.IExpression, modifier foundation.ComparisonPredicateModifier, operators []foundation.INumber, options uint) PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](p_, "initWithLeftExpressions:rightExpressions:modifier:operators:options:", leftExpressions, rightExpressions, modifier, operators, options)
	rv.Autorelease()
	return rv
}

func (p_ PredicateEditorRowTemplate) InitWithCompoundTypes(compoundTypes []foundation.INumber) PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](p_, "initWithCompoundTypes:", compoundTypes)
	rv.Autorelease()
	return rv
}

func (pc _PredicateEditorRowTemplateClass) Alloc() PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](pc, "alloc")
	return rv
}

func (p_ PredicateEditorRowTemplate) Init() PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PredicateEditorRowTemplateClass) New() PredicateEditorRowTemplate {
	rv := ffi.CallMethod[PredicateEditorRowTemplate](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPredicateEditorRowTemplate() PredicateEditorRowTemplate {
	return PredicateEditorRowTemplateClass.New()
}

func (p_ PredicateEditorRowTemplate) MatchForPredicate(predicate foundation.IPredicate) float64 {
	rv := ffi.CallMethod[float64](p_, "matchForPredicate:", predicate)
	return rv
}

func (p_ PredicateEditorRowTemplate) SetPredicate(predicate foundation.IPredicate) {
	ffi.CallMethod[ffi.Void](p_, "setPredicate:", predicate)
}

func (p_ PredicateEditorRowTemplate) DisplayableSubpredicatesOfPredicate(predicate foundation.IPredicate) []foundation.Predicate {
	rv := ffi.CallMethod[[]foundation.Predicate](p_, "displayableSubpredicatesOfPredicate:", predicate)
	return rv
}

func (p_ PredicateEditorRowTemplate) PredicateWithSubpredicates(subpredicates []foundation.IPredicate) foundation.Predicate {
	rv := ffi.CallMethod[foundation.Predicate](p_, "predicateWithSubpredicates:", subpredicates)
	return rv
}

func (p_ PredicateEditorRowTemplate) TemplateViews() []View {
	rv := ffi.CallMethod[[]View](p_, "templateViews")
	return rv
}

func (p_ PredicateEditorRowTemplate) LeftExpressions() []foundation.Expression {
	rv := ffi.CallMethod[[]foundation.Expression](p_, "leftExpressions")
	return rv
}

func (p_ PredicateEditorRowTemplate) RightExpressions() []foundation.Expression {
	rv := ffi.CallMethod[[]foundation.Expression](p_, "rightExpressions")
	return rv
}

func (p_ PredicateEditorRowTemplate) CompoundTypes() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](p_, "compoundTypes")
	return rv
}

func (p_ PredicateEditorRowTemplate) Modifier() foundation.ComparisonPredicateModifier {
	rv := ffi.CallMethod[foundation.ComparisonPredicateModifier](p_, "modifier")
	return rv
}

func (p_ PredicateEditorRowTemplate) Operators() []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](p_, "operators")
	return rv
}

func (p_ PredicateEditorRowTemplate) Options() uint {
	rv := ffi.CallMethod[uint](p_, "options")
	return rv
}
