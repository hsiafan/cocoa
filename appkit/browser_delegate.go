// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type BrowserDelegate interface {
	ImplementsBrowser_IsColumnValid() bool
	// optional
	Browser_IsColumnValid(sender Browser, column int) bool
	ImplementsBrowser_NumberOfRowsInColumn() bool
	// optional
	Browser_NumberOfRowsInColumn(sender Browser, column int) int
	ImplementsBrowser_NumberOfChildrenOfItem() bool
	// optional
	Browser_NumberOfChildrenOfItem(browser Browser, item objc.Object) int
	ImplementsBrowser_TitleOfColumn() bool
	// optional
	Browser_TitleOfColumn(sender Browser, column int) string
	ImplementsBrowser_ShouldTypeSelectForEvent_WithCurrentSearchString() bool
	// optional
	Browser_ShouldTypeSelectForEvent_WithCurrentSearchString(browser Browser, event Event, searchString string) bool
	ImplementsBrowser_TypeSelectStringForRow_InColumn() bool
	// optional
	Browser_TypeSelectStringForRow_InColumn(browser Browser, row int, column int) string
	ImplementsBrowser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString() bool
	// optional
	Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString(browser Browser, startRow int, endRow int, column int, searchString string) int
	ImplementsBrowser_SelectCellWithString_InColumn() bool
	// optional
	Browser_SelectCellWithString_InColumn(sender Browser, title string, column int) bool
	ImplementsBrowser_SelectRow_InColumn() bool
	// optional
	Browser_SelectRow_InColumn(sender Browser, row int, column int) bool
	ImplementsBrowser_SelectionIndexesForProposedSelection_InColumn() bool
	// optional
	Browser_SelectionIndexesForProposedSelection_InColumn(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet
	ImplementsBrowser_Child_OfItem() bool
	// optional
	Browser_Child_OfItem(browser Browser, index int, item objc.Object) objc.IObject
	ImplementsBrowser_IsLeafItem() bool
	// optional
	Browser_IsLeafItem(browser Browser, item objc.Object) bool
	ImplementsBrowser_ShouldEditItem() bool
	// optional
	Browser_ShouldEditItem(browser Browser, item objc.Object) bool
	ImplementsBrowser_ObjectValueForItem() bool
	// optional
	Browser_ObjectValueForItem(browser Browser, item objc.Object) objc.IObject
	ImplementsBrowser_SetObjectValue_ForItem() bool
	// optional
	Browser_SetObjectValue_ForItem(browser Browser, object objc.Object, item objc.Object)
	ImplementsRootItemForBrowser() bool
	// optional
	RootItemForBrowser(browser Browser) objc.IObject
	ImplementsBrowser_PreviewViewControllerForLeafItem() bool
	// optional
	Browser_PreviewViewControllerForLeafItem(browser Browser, item objc.Object) IViewController
	ImplementsBrowser_HeaderViewControllerForItem() bool
	// optional
	Browser_HeaderViewControllerForItem(browser Browser, item objc.Object) IViewController
	ImplementsBrowser_CreateRowsForColumn_InMatrix() bool
	// optional
	Browser_CreateRowsForColumn_InMatrix(sender Browser, column int, matrix Matrix)
	ImplementsBrowser_WillDisplayCell_AtRow_Column() bool
	// optional
	Browser_WillDisplayCell_AtRow_Column(sender Browser, cell objc.Object, row int, column int)
	ImplementsBrowser_DidChangeLastColumn_ToColumn() bool
	// optional
	Browser_DidChangeLastColumn_ToColumn(browser Browser, oldLastColumn int, column int)
	ImplementsBrowserWillScroll() bool
	// optional
	BrowserWillScroll(sender Browser)
	ImplementsBrowserDidScroll() bool
	// optional
	BrowserDidScroll(sender Browser)
	ImplementsBrowser_CanDragRowsWithIndexes_InColumn_WithEvent() bool
	// optional
	Browser_CanDragRowsWithIndexes_InColumn_WithEvent(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool
	ImplementsBrowser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset() bool
	// optional
	Browser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage
	ImplementsBrowser_ValidateDrop_ProposedRow_Column_DropOperation() bool
	// optional
	Browser_ValidateDrop_ProposedRow_Column_DropOperation(browser Browser, info objc.Object, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation
	ImplementsBrowser_AcceptDrop_AtRow_Column_DropOperation() bool
	// optional
	Browser_AcceptDrop_AtRow_Column_DropOperation(browser Browser, info objc.Object, row int, column int, dropOperation BrowserDropOperation) bool
	ImplementsBrowser_WriteRowsWithIndexes_InColumn_ToPasteboard() bool
	// optional
	Browser_WriteRowsWithIndexes_InColumn_ToPasteboard(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool
	ImplementsBrowser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn() bool
	// optional
	// deprecated
	Browser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn(browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string
	ImplementsBrowser_ShouldSizeColumn_ForUserResize_ToWidth() bool
	// optional
	Browser_ShouldSizeColumn_ForUserResize_ToWidth(browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64
	ImplementsBrowser_SizeToFitWidthOfColumn() bool
	// optional
	Browser_SizeToFitWidthOfColumn(browser Browser, columnIndex int) float64
	ImplementsBrowserColumnConfigurationDidChange() bool
	// optional
	BrowserColumnConfigurationDidChange(notification foundation.Notification)
	ImplementsBrowser_HeightOfRow_InColumn() bool
	// optional
	Browser_HeightOfRow_InColumn(browser Browser, row int, columnIndex int) float64
	ImplementsBrowser_ShouldShowCellExpansionForRow_Column() bool
	// optional
	Browser_ShouldShowCellExpansionForRow_Column(browser Browser, row int, column int) bool
}

func WrapBrowserDelegate(v BrowserDelegate) objc.Object {
	return objc.WrapAsProtocol("NSBrowserDelegate", v)
}

type BrowserDelegateBase struct {
}

func (p *BrowserDelegateBase) ImplementsBrowser_IsColumnValid() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_IsColumnValid(sender Browser, column int) bool {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_NumberOfRowsInColumn() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_NumberOfRowsInColumn(sender Browser, column int) int {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_NumberOfChildrenOfItem() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_NumberOfChildrenOfItem(browser Browser, item objc.Object) int {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_TitleOfColumn() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_TitleOfColumn(sender Browser, column int) string {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_ShouldTypeSelectForEvent_WithCurrentSearchString(browser Browser, event Event, searchString string) bool {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_TypeSelectStringForRow_InColumn() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_TypeSelectStringForRow_InColumn(browser Browser, row int, column int) string {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString(browser Browser, startRow int, endRow int, column int, searchString string) int {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_SelectCellWithString_InColumn() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_SelectCellWithString_InColumn(sender Browser, title string, column int) bool {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_SelectRow_InColumn() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_SelectRow_InColumn(sender Browser, row int, column int) bool {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_SelectionIndexesForProposedSelection_InColumn() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_SelectionIndexesForProposedSelection_InColumn(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_Child_OfItem() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_Child_OfItem(browser Browser, index int, item objc.Object) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_IsLeafItem() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_IsLeafItem(browser Browser, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_ShouldEditItem() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_ShouldEditItem(browser Browser, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_ObjectValueForItem() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_ObjectValueForItem(browser Browser, item objc.Object) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_SetObjectValue_ForItem() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_SetObjectValue_ForItem(browser Browser, object objc.Object, item objc.Object) {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsRootItemForBrowser() bool {
	return false
}

func (p *BrowserDelegateBase) RootItemForBrowser(browser Browser) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_PreviewViewControllerForLeafItem() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_PreviewViewControllerForLeafItem(browser Browser, item objc.Object) IViewController {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_HeaderViewControllerForItem() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_HeaderViewControllerForItem(browser Browser, item objc.Object) IViewController {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_CreateRowsForColumn_InMatrix() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_CreateRowsForColumn_InMatrix(sender Browser, column int, matrix Matrix) {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_WillDisplayCell_AtRow_Column() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_WillDisplayCell_AtRow_Column(sender Browser, cell objc.Object, row int, column int) {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_DidChangeLastColumn_ToColumn() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_DidChangeLastColumn_ToColumn(browser Browser, oldLastColumn int, column int) {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowserWillScroll() bool {
	return false
}

func (p *BrowserDelegateBase) BrowserWillScroll(sender Browser) {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowserDidScroll() bool {
	return false
}

func (p *BrowserDelegateBase) BrowserDidScroll(sender Browser) {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_CanDragRowsWithIndexes_InColumn_WithEvent() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_CanDragRowsWithIndexes_InColumn_WithEvent(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_ValidateDrop_ProposedRow_Column_DropOperation() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_ValidateDrop_ProposedRow_Column_DropOperation(browser Browser, info objc.Object, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_AcceptDrop_AtRow_Column_DropOperation() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_AcceptDrop_AtRow_Column_DropOperation(browser Browser, info objc.Object, row int, column int, dropOperation BrowserDropOperation) bool {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_WriteRowsWithIndexes_InColumn_ToPasteboard() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_WriteRowsWithIndexes_InColumn_ToPasteboard(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn() bool {
	return false
}

// deprecated
func (p *BrowserDelegateBase) Browser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn(browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_ShouldSizeColumn_ForUserResize_ToWidth() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_ShouldSizeColumn_ForUserResize_ToWidth(browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64 {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_SizeToFitWidthOfColumn() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_SizeToFitWidthOfColumn(browser Browser, columnIndex int) float64 {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowserColumnConfigurationDidChange() bool {
	return false
}

func (p *BrowserDelegateBase) BrowserColumnConfigurationDidChange(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_HeightOfRow_InColumn() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_HeightOfRow_InColumn(browser Browser, row int, columnIndex int) float64 {
	panic("unimpemented protocol method")
}

func (p *BrowserDelegateBase) ImplementsBrowser_ShouldShowCellExpansionForRow_Column() bool {
	return false
}

func (p *BrowserDelegateBase) Browser_ShouldShowCellExpansionForRow_Column(browser Browser, row int, column int) bool {
	panic("unimpemented protocol method")
}

type BrowserDelegateCreator struct {
	className string
	class     objc.Class
}

func NewBrowserDelegateCreator(name string) *BrowserDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &BrowserDelegateCreator{className: name, class: class}
}

func (c *BrowserDelegateCreator) SetBrowser_IsColumnValid(handle func(o objc.ProtocolBase, sender Browser, column int) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:isColumnValid:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_NumberOfRowsInColumn(handle func(o objc.ProtocolBase, sender Browser, column int) int) {
	objc.AddMethod(c.class, objc.SEL("browser:numberOfRowsInColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_NumberOfChildrenOfItem(handle func(o objc.ProtocolBase, browser Browser, item objc.Object) int) {
	objc.AddMethod(c.class, objc.SEL("browser:numberOfChildrenOfItem:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_TitleOfColumn(handle func(o objc.ProtocolBase, sender Browser, column int) string) {
	objc.AddMethod(c.class, objc.SEL("browser:titleOfColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_ShouldTypeSelectForEvent_WithCurrentSearchString(handle func(o objc.ProtocolBase, browser Browser, event Event, searchString string) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:shouldTypeSelectForEvent:withCurrentSearchString:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_TypeSelectStringForRow_InColumn(handle func(o objc.ProtocolBase, browser Browser, row int, column int) string) {
	objc.AddMethod(c.class, objc.SEL("browser:typeSelectStringForRow:inColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString(handle func(o objc.ProtocolBase, browser Browser, startRow int, endRow int, column int, searchString string) int) {
	objc.AddMethod(c.class, objc.SEL("browser:nextTypeSelectMatchFromRow:toRow:inColumn:forString:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_SelectCellWithString_InColumn(handle func(o objc.ProtocolBase, sender Browser, title string, column int) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:selectCellWithString:inColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_SelectRow_InColumn(handle func(o objc.ProtocolBase, sender Browser, row int, column int) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:selectRow:inColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_SelectionIndexesForProposedSelection_InColumn(handle func(o objc.ProtocolBase, browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet) {
	objc.AddMethod(c.class, objc.SEL("browser:selectionIndexesForProposedSelection:inColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_Child_OfItem(handle func(o objc.ProtocolBase, browser Browser, index int, item objc.Object) objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("browser:child:ofItem:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_IsLeafItem(handle func(o objc.ProtocolBase, browser Browser, item objc.Object) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:isLeafItem:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_ShouldEditItem(handle func(o objc.ProtocolBase, browser Browser, item objc.Object) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:shouldEditItem:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_ObjectValueForItem(handle func(o objc.ProtocolBase, browser Browser, item objc.Object) objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("browser:objectValueForItem:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_SetObjectValue_ForItem(handle func(o objc.ProtocolBase, browser Browser, object objc.Object, item objc.Object)) {
	objc.AddMethod(c.class, objc.SEL("browser:setObjectValue:forItem:"), handle)
}

func (c *BrowserDelegateCreator) SetRootItemForBrowser(handle func(o objc.ProtocolBase, browser Browser) objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("rootItemForBrowser:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_PreviewViewControllerForLeafItem(handle func(o objc.ProtocolBase, browser Browser, item objc.Object) IViewController) {
	objc.AddMethod(c.class, objc.SEL("browser:previewViewControllerForLeafItem:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_HeaderViewControllerForItem(handle func(o objc.ProtocolBase, browser Browser, item objc.Object) IViewController) {
	objc.AddMethod(c.class, objc.SEL("browser:headerViewControllerForItem:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_CreateRowsForColumn_InMatrix(handle func(o objc.ProtocolBase, sender Browser, column int, matrix Matrix)) {
	objc.AddMethod(c.class, objc.SEL("browser:createRowsForColumn:inMatrix:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_WillDisplayCell_AtRow_Column(handle func(o objc.ProtocolBase, sender Browser, cell objc.Object, row int, column int)) {
	objc.AddMethod(c.class, objc.SEL("browser:willDisplayCell:atRow:column:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_DidChangeLastColumn_ToColumn(handle func(o objc.ProtocolBase, browser Browser, oldLastColumn int, column int)) {
	objc.AddMethod(c.class, objc.SEL("browser:didChangeLastColumn:toColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowserWillScroll(handle func(o objc.ProtocolBase, sender Browser)) {
	objc.AddMethod(c.class, objc.SEL("browserWillScroll:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowserDidScroll(handle func(o objc.ProtocolBase, sender Browser)) {
	objc.AddMethod(c.class, objc.SEL("browserDidScroll:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_CanDragRowsWithIndexes_InColumn_WithEvent(handle func(o objc.ProtocolBase, browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:canDragRowsWithIndexes:inColumn:withEvent:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(handle func(o objc.ProtocolBase, browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage) {
	objc.AddMethod(c.class, objc.SEL("browser:draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_ValidateDrop_ProposedRow_Column_DropOperation(handle func(o objc.ProtocolBase, browser Browser, info objc.Object, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation) {
	objc.AddMethod(c.class, objc.SEL("browser:validateDrop:proposedRow:column:dropOperation:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_AcceptDrop_AtRow_Column_DropOperation(handle func(o objc.ProtocolBase, browser Browser, info objc.Object, row int, column int, dropOperation BrowserDropOperation) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:acceptDrop:atRow:column:dropOperation:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_WriteRowsWithIndexes_InColumn_ToPasteboard(handle func(o objc.ProtocolBase, browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:writeRowsWithIndexes:inColumn:toPasteboard:"), handle)
}

// deprecated
func (c *BrowserDelegateCreator) SetBrowser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn(handle func(o objc.ProtocolBase, browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string) {
	objc.AddMethod(c.class, objc.SEL("browser:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:inColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_ShouldSizeColumn_ForUserResize_ToWidth(handle func(o objc.ProtocolBase, browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64) {
	objc.AddMethod(c.class, objc.SEL("browser:shouldSizeColumn:forUserResize:toWidth:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_SizeToFitWidthOfColumn(handle func(o objc.ProtocolBase, browser Browser, columnIndex int) float64) {
	objc.AddMethod(c.class, objc.SEL("browser:sizeToFitWidthOfColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowserColumnConfigurationDidChange(handle func(o objc.ProtocolBase, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("browserColumnConfigurationDidChange:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_HeightOfRow_InColumn(handle func(o objc.ProtocolBase, browser Browser, row int, columnIndex int) float64) {
	objc.AddMethod(c.class, objc.SEL("browser:heightOfRow:inColumn:"), handle)
}

func (c *BrowserDelegateCreator) SetBrowser_ShouldShowCellExpansionForRow_Column(handle func(o objc.ProtocolBase, browser Browser, row int, column int) bool) {
	objc.AddMethod(c.class, objc.SEL("browser:shouldShowCellExpansionForRow:column:"), handle)
}

func (c *BrowserDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
