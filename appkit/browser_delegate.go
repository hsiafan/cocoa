// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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
	Browser_ValidateDrop_ProposedRow_Column_DropOperation(browser Browser, info DraggingInfoWrapper, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation
	ImplementsBrowser_AcceptDrop_AtRow_Column_DropOperation() bool
	// optional
	Browser_AcceptDrop_AtRow_Column_DropOperation(browser Browser, info DraggingInfoWrapper, row int, column int, dropOperation BrowserDropOperation) bool
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

type BrowserDelegateImpl struct {
	_Browser_IsColumnValid                                                               func(sender Browser, column int) bool
	_Browser_NumberOfRowsInColumn                                                        func(sender Browser, column int) int
	_Browser_NumberOfChildrenOfItem                                                      func(browser Browser, item objc.Object) int
	_Browser_TitleOfColumn                                                               func(sender Browser, column int) string
	_Browser_ShouldTypeSelectForEvent_WithCurrentSearchString                            func(browser Browser, event Event, searchString string) bool
	_Browser_TypeSelectStringForRow_InColumn                                             func(browser Browser, row int, column int) string
	_Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString                         func(browser Browser, startRow int, endRow int, column int, searchString string) int
	_Browser_SelectCellWithString_InColumn                                               func(sender Browser, title string, column int) bool
	_Browser_SelectRow_InColumn                                                          func(sender Browser, row int, column int) bool
	_Browser_SelectionIndexesForProposedSelection_InColumn                               func(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet
	_Browser_Child_OfItem                                                                func(browser Browser, index int, item objc.Object) objc.IObject
	_Browser_IsLeafItem                                                                  func(browser Browser, item objc.Object) bool
	_Browser_ShouldEditItem                                                              func(browser Browser, item objc.Object) bool
	_Browser_ObjectValueForItem                                                          func(browser Browser, item objc.Object) objc.IObject
	_Browser_SetObjectValue_ForItem                                                      func(browser Browser, object objc.Object, item objc.Object)
	_RootItemForBrowser                                                                  func(browser Browser) objc.IObject
	_Browser_PreviewViewControllerForLeafItem                                            func(browser Browser, item objc.Object) IViewController
	_Browser_HeaderViewControllerForItem                                                 func(browser Browser, item objc.Object) IViewController
	_Browser_CreateRowsForColumn_InMatrix                                                func(sender Browser, column int, matrix Matrix)
	_Browser_WillDisplayCell_AtRow_Column                                                func(sender Browser, cell objc.Object, row int, column int)
	_Browser_DidChangeLastColumn_ToColumn                                                func(browser Browser, oldLastColumn int, column int)
	_BrowserWillScroll                                                                   func(sender Browser)
	_BrowserDidScroll                                                                    func(sender Browser)
	_Browser_CanDragRowsWithIndexes_InColumn_WithEvent                                   func(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool
	_Browser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset                   func(browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage
	_Browser_ValidateDrop_ProposedRow_Column_DropOperation                               func(browser Browser, info DraggingInfoWrapper, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation
	_Browser_AcceptDrop_AtRow_Column_DropOperation                                       func(browser Browser, info DraggingInfoWrapper, row int, column int, dropOperation BrowserDropOperation) bool
	_Browser_WriteRowsWithIndexes_InColumn_ToPasteboard                                  func(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool
	_Browser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn func(browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string
	_Browser_ShouldSizeColumn_ForUserResize_ToWidth                                      func(browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64
	_Browser_SizeToFitWidthOfColumn                                                      func(browser Browser, columnIndex int) float64
	_BrowserColumnConfigurationDidChange                                                 func(notification foundation.Notification)
	_Browser_HeightOfRow_InColumn                                                        func(browser Browser, row int, columnIndex int) float64
	_Browser_ShouldShowCellExpansionForRow_Column                                        func(browser Browser, row int, column int) bool
}

func (di *BrowserDelegateImpl) ImplementsBrowser_IsColumnValid() bool {
	return di._Browser_IsColumnValid != nil
}

func (di *BrowserDelegateImpl) SetBrowser_IsColumnValid(f func(sender Browser, column int) bool) {
	di._Browser_IsColumnValid = f
}

func (di *BrowserDelegateImpl) Browser_IsColumnValid(sender Browser, column int) bool {
	return di._Browser_IsColumnValid(sender, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_NumberOfRowsInColumn() bool {
	return di._Browser_NumberOfRowsInColumn != nil
}

func (di *BrowserDelegateImpl) SetBrowser_NumberOfRowsInColumn(f func(sender Browser, column int) int) {
	di._Browser_NumberOfRowsInColumn = f
}

func (di *BrowserDelegateImpl) Browser_NumberOfRowsInColumn(sender Browser, column int) int {
	return di._Browser_NumberOfRowsInColumn(sender, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_NumberOfChildrenOfItem() bool {
	return di._Browser_NumberOfChildrenOfItem != nil
}

func (di *BrowserDelegateImpl) SetBrowser_NumberOfChildrenOfItem(f func(browser Browser, item objc.Object) int) {
	di._Browser_NumberOfChildrenOfItem = f
}

func (di *BrowserDelegateImpl) Browser_NumberOfChildrenOfItem(browser Browser, item objc.Object) int {
	return di._Browser_NumberOfChildrenOfItem(browser, item)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_TitleOfColumn() bool {
	return di._Browser_TitleOfColumn != nil
}

func (di *BrowserDelegateImpl) SetBrowser_TitleOfColumn(f func(sender Browser, column int) string) {
	di._Browser_TitleOfColumn = f
}

func (di *BrowserDelegateImpl) Browser_TitleOfColumn(sender Browser, column int) string {
	return di._Browser_TitleOfColumn(sender, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return di._Browser_ShouldTypeSelectForEvent_WithCurrentSearchString != nil
}

func (di *BrowserDelegateImpl) SetBrowser_ShouldTypeSelectForEvent_WithCurrentSearchString(f func(browser Browser, event Event, searchString string) bool) {
	di._Browser_ShouldTypeSelectForEvent_WithCurrentSearchString = f
}

func (di *BrowserDelegateImpl) Browser_ShouldTypeSelectForEvent_WithCurrentSearchString(browser Browser, event Event, searchString string) bool {
	return di._Browser_ShouldTypeSelectForEvent_WithCurrentSearchString(browser, event, searchString)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_TypeSelectStringForRow_InColumn() bool {
	return di._Browser_TypeSelectStringForRow_InColumn != nil
}

func (di *BrowserDelegateImpl) SetBrowser_TypeSelectStringForRow_InColumn(f func(browser Browser, row int, column int) string) {
	di._Browser_TypeSelectStringForRow_InColumn = f
}

func (di *BrowserDelegateImpl) Browser_TypeSelectStringForRow_InColumn(browser Browser, row int, column int) string {
	return di._Browser_TypeSelectStringForRow_InColumn(browser, row, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString() bool {
	return di._Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString != nil
}

func (di *BrowserDelegateImpl) SetBrowser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString(f func(browser Browser, startRow int, endRow int, column int, searchString string) int) {
	di._Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString = f
}

func (di *BrowserDelegateImpl) Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString(browser Browser, startRow int, endRow int, column int, searchString string) int {
	return di._Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString(browser, startRow, endRow, column, searchString)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_SelectCellWithString_InColumn() bool {
	return di._Browser_SelectCellWithString_InColumn != nil
}

func (di *BrowserDelegateImpl) SetBrowser_SelectCellWithString_InColumn(f func(sender Browser, title string, column int) bool) {
	di._Browser_SelectCellWithString_InColumn = f
}

func (di *BrowserDelegateImpl) Browser_SelectCellWithString_InColumn(sender Browser, title string, column int) bool {
	return di._Browser_SelectCellWithString_InColumn(sender, title, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_SelectRow_InColumn() bool {
	return di._Browser_SelectRow_InColumn != nil
}

func (di *BrowserDelegateImpl) SetBrowser_SelectRow_InColumn(f func(sender Browser, row int, column int) bool) {
	di._Browser_SelectRow_InColumn = f
}

func (di *BrowserDelegateImpl) Browser_SelectRow_InColumn(sender Browser, row int, column int) bool {
	return di._Browser_SelectRow_InColumn(sender, row, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_SelectionIndexesForProposedSelection_InColumn() bool {
	return di._Browser_SelectionIndexesForProposedSelection_InColumn != nil
}

func (di *BrowserDelegateImpl) SetBrowser_SelectionIndexesForProposedSelection_InColumn(f func(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet) {
	di._Browser_SelectionIndexesForProposedSelection_InColumn = f
}

func (di *BrowserDelegateImpl) Browser_SelectionIndexesForProposedSelection_InColumn(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IIndexSet {
	return di._Browser_SelectionIndexesForProposedSelection_InColumn(browser, proposedSelectionIndexes, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_Child_OfItem() bool {
	return di._Browser_Child_OfItem != nil
}

func (di *BrowserDelegateImpl) SetBrowser_Child_OfItem(f func(browser Browser, index int, item objc.Object) objc.IObject) {
	di._Browser_Child_OfItem = f
}

func (di *BrowserDelegateImpl) Browser_Child_OfItem(browser Browser, index int, item objc.Object) objc.IObject {
	return di._Browser_Child_OfItem(browser, index, item)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_IsLeafItem() bool {
	return di._Browser_IsLeafItem != nil
}

func (di *BrowserDelegateImpl) SetBrowser_IsLeafItem(f func(browser Browser, item objc.Object) bool) {
	di._Browser_IsLeafItem = f
}

func (di *BrowserDelegateImpl) Browser_IsLeafItem(browser Browser, item objc.Object) bool {
	return di._Browser_IsLeafItem(browser, item)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_ShouldEditItem() bool {
	return di._Browser_ShouldEditItem != nil
}

func (di *BrowserDelegateImpl) SetBrowser_ShouldEditItem(f func(browser Browser, item objc.Object) bool) {
	di._Browser_ShouldEditItem = f
}

func (di *BrowserDelegateImpl) Browser_ShouldEditItem(browser Browser, item objc.Object) bool {
	return di._Browser_ShouldEditItem(browser, item)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_ObjectValueForItem() bool {
	return di._Browser_ObjectValueForItem != nil
}

func (di *BrowserDelegateImpl) SetBrowser_ObjectValueForItem(f func(browser Browser, item objc.Object) objc.IObject) {
	di._Browser_ObjectValueForItem = f
}

func (di *BrowserDelegateImpl) Browser_ObjectValueForItem(browser Browser, item objc.Object) objc.IObject {
	return di._Browser_ObjectValueForItem(browser, item)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_SetObjectValue_ForItem() bool {
	return di._Browser_SetObjectValue_ForItem != nil
}

func (di *BrowserDelegateImpl) SetBrowser_SetObjectValue_ForItem(f func(browser Browser, object objc.Object, item objc.Object)) {
	di._Browser_SetObjectValue_ForItem = f
}

func (di *BrowserDelegateImpl) Browser_SetObjectValue_ForItem(browser Browser, object objc.Object, item objc.Object) {
	di._Browser_SetObjectValue_ForItem(browser, object, item)
}
func (di *BrowserDelegateImpl) ImplementsRootItemForBrowser() bool {
	return di._RootItemForBrowser != nil
}

func (di *BrowserDelegateImpl) SetRootItemForBrowser(f func(browser Browser) objc.IObject) {
	di._RootItemForBrowser = f
}

func (di *BrowserDelegateImpl) RootItemForBrowser(browser Browser) objc.IObject {
	return di._RootItemForBrowser(browser)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_PreviewViewControllerForLeafItem() bool {
	return di._Browser_PreviewViewControllerForLeafItem != nil
}

func (di *BrowserDelegateImpl) SetBrowser_PreviewViewControllerForLeafItem(f func(browser Browser, item objc.Object) IViewController) {
	di._Browser_PreviewViewControllerForLeafItem = f
}

func (di *BrowserDelegateImpl) Browser_PreviewViewControllerForLeafItem(browser Browser, item objc.Object) IViewController {
	return di._Browser_PreviewViewControllerForLeafItem(browser, item)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_HeaderViewControllerForItem() bool {
	return di._Browser_HeaderViewControllerForItem != nil
}

func (di *BrowserDelegateImpl) SetBrowser_HeaderViewControllerForItem(f func(browser Browser, item objc.Object) IViewController) {
	di._Browser_HeaderViewControllerForItem = f
}

func (di *BrowserDelegateImpl) Browser_HeaderViewControllerForItem(browser Browser, item objc.Object) IViewController {
	return di._Browser_HeaderViewControllerForItem(browser, item)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_CreateRowsForColumn_InMatrix() bool {
	return di._Browser_CreateRowsForColumn_InMatrix != nil
}

func (di *BrowserDelegateImpl) SetBrowser_CreateRowsForColumn_InMatrix(f func(sender Browser, column int, matrix Matrix)) {
	di._Browser_CreateRowsForColumn_InMatrix = f
}

func (di *BrowserDelegateImpl) Browser_CreateRowsForColumn_InMatrix(sender Browser, column int, matrix Matrix) {
	di._Browser_CreateRowsForColumn_InMatrix(sender, column, matrix)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_WillDisplayCell_AtRow_Column() bool {
	return di._Browser_WillDisplayCell_AtRow_Column != nil
}

func (di *BrowserDelegateImpl) SetBrowser_WillDisplayCell_AtRow_Column(f func(sender Browser, cell objc.Object, row int, column int)) {
	di._Browser_WillDisplayCell_AtRow_Column = f
}

func (di *BrowserDelegateImpl) Browser_WillDisplayCell_AtRow_Column(sender Browser, cell objc.Object, row int, column int) {
	di._Browser_WillDisplayCell_AtRow_Column(sender, cell, row, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_DidChangeLastColumn_ToColumn() bool {
	return di._Browser_DidChangeLastColumn_ToColumn != nil
}

func (di *BrowserDelegateImpl) SetBrowser_DidChangeLastColumn_ToColumn(f func(browser Browser, oldLastColumn int, column int)) {
	di._Browser_DidChangeLastColumn_ToColumn = f
}

func (di *BrowserDelegateImpl) Browser_DidChangeLastColumn_ToColumn(browser Browser, oldLastColumn int, column int) {
	di._Browser_DidChangeLastColumn_ToColumn(browser, oldLastColumn, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowserWillScroll() bool {
	return di._BrowserWillScroll != nil
}

func (di *BrowserDelegateImpl) SetBrowserWillScroll(f func(sender Browser)) {
	di._BrowserWillScroll = f
}

func (di *BrowserDelegateImpl) BrowserWillScroll(sender Browser) {
	di._BrowserWillScroll(sender)
}
func (di *BrowserDelegateImpl) ImplementsBrowserDidScroll() bool {
	return di._BrowserDidScroll != nil
}

func (di *BrowserDelegateImpl) SetBrowserDidScroll(f func(sender Browser)) {
	di._BrowserDidScroll = f
}

func (di *BrowserDelegateImpl) BrowserDidScroll(sender Browser) {
	di._BrowserDidScroll(sender)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_CanDragRowsWithIndexes_InColumn_WithEvent() bool {
	return di._Browser_CanDragRowsWithIndexes_InColumn_WithEvent != nil
}

func (di *BrowserDelegateImpl) SetBrowser_CanDragRowsWithIndexes_InColumn_WithEvent(f func(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool) {
	di._Browser_CanDragRowsWithIndexes_InColumn_WithEvent = f
}

func (di *BrowserDelegateImpl) Browser_CanDragRowsWithIndexes_InColumn_WithEvent(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool {
	return di._Browser_CanDragRowsWithIndexes_InColumn_WithEvent(browser, rowIndexes, column, event)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset() bool {
	return di._Browser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset != nil
}

func (di *BrowserDelegateImpl) SetBrowser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(f func(browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage) {
	di._Browser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset = f
}

func (di *BrowserDelegateImpl) Browser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(browser Browser, rowIndexes foundation.IndexSet, column int, event Event, dragImageOffset *foundation.Point) IImage {
	return di._Browser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(browser, rowIndexes, column, event, dragImageOffset)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_ValidateDrop_ProposedRow_Column_DropOperation() bool {
	return di._Browser_ValidateDrop_ProposedRow_Column_DropOperation != nil
}

func (di *BrowserDelegateImpl) SetBrowser_ValidateDrop_ProposedRow_Column_DropOperation(f func(browser Browser, info DraggingInfoWrapper, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation) {
	di._Browser_ValidateDrop_ProposedRow_Column_DropOperation = f
}

func (di *BrowserDelegateImpl) Browser_ValidateDrop_ProposedRow_Column_DropOperation(browser Browser, info DraggingInfoWrapper, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation {
	return di._Browser_ValidateDrop_ProposedRow_Column_DropOperation(browser, info, row, column, dropOperation)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_AcceptDrop_AtRow_Column_DropOperation() bool {
	return di._Browser_AcceptDrop_AtRow_Column_DropOperation != nil
}

func (di *BrowserDelegateImpl) SetBrowser_AcceptDrop_AtRow_Column_DropOperation(f func(browser Browser, info DraggingInfoWrapper, row int, column int, dropOperation BrowserDropOperation) bool) {
	di._Browser_AcceptDrop_AtRow_Column_DropOperation = f
}

func (di *BrowserDelegateImpl) Browser_AcceptDrop_AtRow_Column_DropOperation(browser Browser, info DraggingInfoWrapper, row int, column int, dropOperation BrowserDropOperation) bool {
	return di._Browser_AcceptDrop_AtRow_Column_DropOperation(browser, info, row, column, dropOperation)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_WriteRowsWithIndexes_InColumn_ToPasteboard() bool {
	return di._Browser_WriteRowsWithIndexes_InColumn_ToPasteboard != nil
}

func (di *BrowserDelegateImpl) SetBrowser_WriteRowsWithIndexes_InColumn_ToPasteboard(f func(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool) {
	di._Browser_WriteRowsWithIndexes_InColumn_ToPasteboard = f
}

func (di *BrowserDelegateImpl) Browser_WriteRowsWithIndexes_InColumn_ToPasteboard(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool {
	return di._Browser_WriteRowsWithIndexes_InColumn_ToPasteboard(browser, rowIndexes, column, pasteboard)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn() bool {
	return di._Browser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn != nil
}

// deprecated
func (di *BrowserDelegateImpl) SetBrowser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn(f func(browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string) {
	di._Browser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn = f
}

// deprecated
func (di *BrowserDelegateImpl) Browser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn(browser Browser, dropDestination foundation.URL, rowIndexes foundation.IndexSet, column int) []string {
	return di._Browser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn(browser, dropDestination, rowIndexes, column)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_ShouldSizeColumn_ForUserResize_ToWidth() bool {
	return di._Browser_ShouldSizeColumn_ForUserResize_ToWidth != nil
}

func (di *BrowserDelegateImpl) SetBrowser_ShouldSizeColumn_ForUserResize_ToWidth(f func(browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64) {
	di._Browser_ShouldSizeColumn_ForUserResize_ToWidth = f
}

func (di *BrowserDelegateImpl) Browser_ShouldSizeColumn_ForUserResize_ToWidth(browser Browser, columnIndex int, forUserResize bool, suggestedWidth float64) float64 {
	return di._Browser_ShouldSizeColumn_ForUserResize_ToWidth(browser, columnIndex, forUserResize, suggestedWidth)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_SizeToFitWidthOfColumn() bool {
	return di._Browser_SizeToFitWidthOfColumn != nil
}

func (di *BrowserDelegateImpl) SetBrowser_SizeToFitWidthOfColumn(f func(browser Browser, columnIndex int) float64) {
	di._Browser_SizeToFitWidthOfColumn = f
}

func (di *BrowserDelegateImpl) Browser_SizeToFitWidthOfColumn(browser Browser, columnIndex int) float64 {
	return di._Browser_SizeToFitWidthOfColumn(browser, columnIndex)
}
func (di *BrowserDelegateImpl) ImplementsBrowserColumnConfigurationDidChange() bool {
	return di._BrowserColumnConfigurationDidChange != nil
}

func (di *BrowserDelegateImpl) SetBrowserColumnConfigurationDidChange(f func(notification foundation.Notification)) {
	di._BrowserColumnConfigurationDidChange = f
}

func (di *BrowserDelegateImpl) BrowserColumnConfigurationDidChange(notification foundation.Notification) {
	di._BrowserColumnConfigurationDidChange(notification)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_HeightOfRow_InColumn() bool {
	return di._Browser_HeightOfRow_InColumn != nil
}

func (di *BrowserDelegateImpl) SetBrowser_HeightOfRow_InColumn(f func(browser Browser, row int, columnIndex int) float64) {
	di._Browser_HeightOfRow_InColumn = f
}

func (di *BrowserDelegateImpl) Browser_HeightOfRow_InColumn(browser Browser, row int, columnIndex int) float64 {
	return di._Browser_HeightOfRow_InColumn(browser, row, columnIndex)
}
func (di *BrowserDelegateImpl) ImplementsBrowser_ShouldShowCellExpansionForRow_Column() bool {
	return di._Browser_ShouldShowCellExpansionForRow_Column != nil
}

func (di *BrowserDelegateImpl) SetBrowser_ShouldShowCellExpansionForRow_Column(f func(browser Browser, row int, column int) bool) {
	di._Browser_ShouldShowCellExpansionForRow_Column = f
}

func (di *BrowserDelegateImpl) Browser_ShouldShowCellExpansionForRow_Column(browser Browser, row int, column int) bool {
	return di._Browser_ShouldShowCellExpansionForRow_Column(browser, row, column)
}

type BrowserDelegateWrapper struct {
	objc.Object
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_IsColumnValid() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:isColumnValid:"))
}

func (b_ BrowserDelegateWrapper) Browser_IsColumnValid(sender IBrowser, column int) bool {
	rv := ffi.CallMethod[bool](b_, "browser:isColumnValid:", sender, column)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_NumberOfRowsInColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:numberOfRowsInColumn:"))
}

func (b_ BrowserDelegateWrapper) Browser_NumberOfRowsInColumn(sender IBrowser, column int) int {
	rv := ffi.CallMethod[int](b_, "browser:numberOfRowsInColumn:", sender, column)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_NumberOfChildrenOfItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:numberOfChildrenOfItem:"))
}

func (b_ BrowserDelegateWrapper) Browser_NumberOfChildrenOfItem(browser IBrowser, item objc.IObject) int {
	rv := ffi.CallMethod[int](b_, "browser:numberOfChildrenOfItem:", browser, item)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_TitleOfColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:titleOfColumn:"))
}

func (b_ BrowserDelegateWrapper) Browser_TitleOfColumn(sender IBrowser, column int) string {
	rv := ffi.CallMethod[string](b_, "browser:titleOfColumn:", sender, column)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:shouldTypeSelectForEvent:withCurrentSearchString:"))
}

func (b_ BrowserDelegateWrapper) Browser_ShouldTypeSelectForEvent_WithCurrentSearchString(browser IBrowser, event IEvent, searchString string) bool {
	rv := ffi.CallMethod[bool](b_, "browser:shouldTypeSelectForEvent:withCurrentSearchString:", browser, event, searchString)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_TypeSelectStringForRow_InColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:typeSelectStringForRow:inColumn:"))
}

func (b_ BrowserDelegateWrapper) Browser_TypeSelectStringForRow_InColumn(browser IBrowser, row int, column int) string {
	rv := ffi.CallMethod[string](b_, "browser:typeSelectStringForRow:inColumn:", browser, row, column)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:nextTypeSelectMatchFromRow:toRow:inColumn:forString:"))
}

func (b_ BrowserDelegateWrapper) Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString(browser IBrowser, startRow int, endRow int, column int, searchString string) int {
	rv := ffi.CallMethod[int](b_, "browser:nextTypeSelectMatchFromRow:toRow:inColumn:forString:", browser, startRow, endRow, column, searchString)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_SelectCellWithString_InColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:selectCellWithString:inColumn:"))
}

func (b_ BrowserDelegateWrapper) Browser_SelectCellWithString_InColumn(sender IBrowser, title string, column int) bool {
	rv := ffi.CallMethod[bool](b_, "browser:selectCellWithString:inColumn:", sender, title, column)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_SelectRow_InColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:selectRow:inColumn:"))
}

func (b_ BrowserDelegateWrapper) Browser_SelectRow_InColumn(sender IBrowser, row int, column int) bool {
	rv := ffi.CallMethod[bool](b_, "browser:selectRow:inColumn:", sender, row, column)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_SelectionIndexesForProposedSelection_InColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:selectionIndexesForProposedSelection:inColumn:"))
}

func (b_ BrowserDelegateWrapper) Browser_SelectionIndexesForProposedSelection_InColumn(browser IBrowser, proposedSelectionIndexes foundation.IIndexSet, column int) foundation.IndexSet {
	rv := ffi.CallMethod[foundation.IndexSet](b_, "browser:selectionIndexesForProposedSelection:inColumn:", browser, proposedSelectionIndexes, column)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_Child_OfItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:child:ofItem:"))
}

func (b_ BrowserDelegateWrapper) Browser_Child_OfItem(browser IBrowser, index int, item objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "browser:child:ofItem:", browser, index, item)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_IsLeafItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:isLeafItem:"))
}

func (b_ BrowserDelegateWrapper) Browser_IsLeafItem(browser IBrowser, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](b_, "browser:isLeafItem:", browser, item)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_ShouldEditItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:shouldEditItem:"))
}

func (b_ BrowserDelegateWrapper) Browser_ShouldEditItem(browser IBrowser, item objc.IObject) bool {
	rv := ffi.CallMethod[bool](b_, "browser:shouldEditItem:", browser, item)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_ObjectValueForItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:objectValueForItem:"))
}

func (b_ BrowserDelegateWrapper) Browser_ObjectValueForItem(browser IBrowser, item objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "browser:objectValueForItem:", browser, item)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_SetObjectValue_ForItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:setObjectValue:forItem:"))
}

func (b_ BrowserDelegateWrapper) Browser_SetObjectValue_ForItem(browser IBrowser, object objc.IObject, item objc.IObject) {
	ffi.CallMethod[ffi.Void](b_, "browser:setObjectValue:forItem:", browser, object, item)
}

func (b_ *BrowserDelegateWrapper) ImplementsRootItemForBrowser() bool {
	return b_.RespondsToSelector(objc.GetSelector("rootItemForBrowser:"))
}

func (b_ BrowserDelegateWrapper) RootItemForBrowser(browser IBrowser) objc.Object {
	rv := ffi.CallMethod[objc.Object](b_, "rootItemForBrowser:", browser)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_PreviewViewControllerForLeafItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:previewViewControllerForLeafItem:"))
}

func (b_ BrowserDelegateWrapper) Browser_PreviewViewControllerForLeafItem(browser IBrowser, item objc.IObject) ViewController {
	rv := ffi.CallMethod[ViewController](b_, "browser:previewViewControllerForLeafItem:", browser, item)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_HeaderViewControllerForItem() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:headerViewControllerForItem:"))
}

func (b_ BrowserDelegateWrapper) Browser_HeaderViewControllerForItem(browser IBrowser, item objc.IObject) ViewController {
	rv := ffi.CallMethod[ViewController](b_, "browser:headerViewControllerForItem:", browser, item)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_CreateRowsForColumn_InMatrix() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:createRowsForColumn:inMatrix:"))
}

func (b_ BrowserDelegateWrapper) Browser_CreateRowsForColumn_InMatrix(sender IBrowser, column int, matrix IMatrix) {
	ffi.CallMethod[ffi.Void](b_, "browser:createRowsForColumn:inMatrix:", sender, column, matrix)
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_WillDisplayCell_AtRow_Column() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:willDisplayCell:atRow:column:"))
}

func (b_ BrowserDelegateWrapper) Browser_WillDisplayCell_AtRow_Column(sender IBrowser, cell objc.IObject, row int, column int) {
	ffi.CallMethod[ffi.Void](b_, "browser:willDisplayCell:atRow:column:", sender, cell, row, column)
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_DidChangeLastColumn_ToColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:didChangeLastColumn:toColumn:"))
}

func (b_ BrowserDelegateWrapper) Browser_DidChangeLastColumn_ToColumn(browser IBrowser, oldLastColumn int, column int) {
	ffi.CallMethod[ffi.Void](b_, "browser:didChangeLastColumn:toColumn:", browser, oldLastColumn, column)
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowserWillScroll() bool {
	return b_.RespondsToSelector(objc.GetSelector("browserWillScroll:"))
}

func (b_ BrowserDelegateWrapper) BrowserWillScroll(sender IBrowser) {
	ffi.CallMethod[ffi.Void](b_, "browserWillScroll:", sender)
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowserDidScroll() bool {
	return b_.RespondsToSelector(objc.GetSelector("browserDidScroll:"))
}

func (b_ BrowserDelegateWrapper) BrowserDidScroll(sender IBrowser) {
	ffi.CallMethod[ffi.Void](b_, "browserDidScroll:", sender)
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_CanDragRowsWithIndexes_InColumn_WithEvent() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:canDragRowsWithIndexes:inColumn:withEvent:"))
}

func (b_ BrowserDelegateWrapper) Browser_CanDragRowsWithIndexes_InColumn_WithEvent(browser IBrowser, rowIndexes foundation.IIndexSet, column int, event IEvent) bool {
	rv := ffi.CallMethod[bool](b_, "browser:canDragRowsWithIndexes:inColumn:withEvent:", browser, rowIndexes, column, event)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:draggingImageForRowsWithIndexes:inColumn:withEvent:offset:"))
}

func (b_ BrowserDelegateWrapper) Browser_DraggingImageForRowsWithIndexes_InColumn_WithEvent_Offset(browser IBrowser, rowIndexes foundation.IIndexSet, column int, event IEvent, dragImageOffset *foundation.Point) Image {
	rv := ffi.CallMethod[Image](b_, "browser:draggingImageForRowsWithIndexes:inColumn:withEvent:offset:", browser, rowIndexes, column, event, dragImageOffset)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_ValidateDrop_ProposedRow_Column_DropOperation() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:validateDrop:proposedRow:column:dropOperation:"))
}

func (b_ BrowserDelegateWrapper) Browser_ValidateDrop_ProposedRow_Column_DropOperation(browser IBrowser, info DraggingInfo, row *int, column *int, dropOperation *BrowserDropOperation) DragOperation {
	po := ffi.CreateProtocol("NSDraggingInfo", info)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](b_, "browser:validateDrop:proposedRow:column:dropOperation:", browser, po, row, column, dropOperation)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_AcceptDrop_AtRow_Column_DropOperation() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:acceptDrop:atRow:column:dropOperation:"))
}

func (b_ BrowserDelegateWrapper) Browser_AcceptDrop_AtRow_Column_DropOperation(browser IBrowser, info DraggingInfo, row int, column int, dropOperation BrowserDropOperation) bool {
	po := ffi.CreateProtocol("NSDraggingInfo", info)
	defer po.Release()
	rv := ffi.CallMethod[bool](b_, "browser:acceptDrop:atRow:column:dropOperation:", browser, po, row, column, dropOperation)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_WriteRowsWithIndexes_InColumn_ToPasteboard() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:writeRowsWithIndexes:inColumn:toPasteboard:"))
}

func (b_ BrowserDelegateWrapper) Browser_WriteRowsWithIndexes_InColumn_ToPasteboard(browser IBrowser, rowIndexes foundation.IIndexSet, column int, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](b_, "browser:writeRowsWithIndexes:inColumn:toPasteboard:", browser, rowIndexes, column, pasteboard)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:inColumn:"))
}

// deprecated
func (b_ BrowserDelegateWrapper) Browser_NamesOfPromisedFilesDroppedAtDestination_ForDraggedRowsWithIndexes_InColumn(browser IBrowser, dropDestination foundation.IURL, rowIndexes foundation.IIndexSet, column int) []string {
	rv := ffi.CallMethod[[]string](b_, "browser:namesOfPromisedFilesDroppedAtDestination:forDraggedRowsWithIndexes:inColumn:", browser, dropDestination, rowIndexes, column)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_ShouldSizeColumn_ForUserResize_ToWidth() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:shouldSizeColumn:forUserResize:toWidth:"))
}

func (b_ BrowserDelegateWrapper) Browser_ShouldSizeColumn_ForUserResize_ToWidth(browser IBrowser, columnIndex int, forUserResize bool, suggestedWidth float64) float64 {
	rv := ffi.CallMethod[float64](b_, "browser:shouldSizeColumn:forUserResize:toWidth:", browser, columnIndex, forUserResize, suggestedWidth)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_SizeToFitWidthOfColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:sizeToFitWidthOfColumn:"))
}

func (b_ BrowserDelegateWrapper) Browser_SizeToFitWidthOfColumn(browser IBrowser, columnIndex int) float64 {
	rv := ffi.CallMethod[float64](b_, "browser:sizeToFitWidthOfColumn:", browser, columnIndex)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowserColumnConfigurationDidChange() bool {
	return b_.RespondsToSelector(objc.GetSelector("browserColumnConfigurationDidChange:"))
}

func (b_ BrowserDelegateWrapper) BrowserColumnConfigurationDidChange(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](b_, "browserColumnConfigurationDidChange:", notification)
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_HeightOfRow_InColumn() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:heightOfRow:inColumn:"))
}

func (b_ BrowserDelegateWrapper) Browser_HeightOfRow_InColumn(browser IBrowser, row int, columnIndex int) float64 {
	rv := ffi.CallMethod[float64](b_, "browser:heightOfRow:inColumn:", browser, row, columnIndex)
	return rv
}

func (b_ *BrowserDelegateWrapper) ImplementsBrowser_ShouldShowCellExpansionForRow_Column() bool {
	return b_.RespondsToSelector(objc.GetSelector("browser:shouldShowCellExpansionForRow:column:"))
}

func (b_ BrowserDelegateWrapper) Browser_ShouldShowCellExpansionForRow_Column(browser IBrowser, row int, column int) bool {
	rv := ffi.CallMethod[bool](b_, "browser:shouldShowCellExpansionForRow:column:", browser, row, column)
	return rv
}
