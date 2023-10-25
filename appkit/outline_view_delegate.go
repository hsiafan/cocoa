// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type OutlineViewDelegate interface {
	ControlTextEditingDelegate
	ImplementsOutlineView_ShouldExpandItem() bool
	// optional
	OutlineView_ShouldExpandItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_ShouldCollapseItem() bool
	// optional
	OutlineView_ShouldCollapseItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_TypeSelectStringForTableColumn_Item() bool
	// optional
	OutlineView_TypeSelectStringForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string
	ImplementsOutlineView_NextTypeSelectMatchFromItem_ToItem_ForString() bool
	// optional
	OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject
	ImplementsOutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool
	// optional
	OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(outlineView OutlineView, event Event, searchString string) bool
	ImplementsOutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation() bool
	// optional
	OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string
	ImplementsOutlineView_ShouldSelectTableColumn() bool
	// optional
	OutlineView_ShouldSelectTableColumn(outlineView OutlineView, tableColumn TableColumn) bool
	ImplementsOutlineView_ShouldSelectItem() bool
	// optional
	OutlineView_ShouldSelectItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_SelectionIndexesForProposedSelection() bool
	// optional
	OutlineView_SelectionIndexesForProposedSelection(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	ImplementsSelectionShouldChangeInOutlineView() bool
	// optional
	SelectionShouldChangeInOutlineView(outlineView OutlineView) bool
	ImplementsOutlineViewSelectionIsChanging() bool
	// optional
	OutlineViewSelectionIsChanging(notification foundation.Notification)
	ImplementsOutlineViewSelectionDidChange() bool
	// optional
	OutlineViewSelectionDidChange(notification foundation.Notification)
	ImplementsOutlineView_WillDisplayCell_ForTableColumn_Item() bool
	// optional
	OutlineView_WillDisplayCell_ForTableColumn_Item(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	ImplementsOutlineView_WillDisplayOutlineCell_ForTableColumn_Item() bool
	// optional
	OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	ImplementsOutlineView_DataCellForTableColumn_Item() bool
	// optional
	OutlineView_DataCellForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell
	ImplementsOutlineView_ShouldShowOutlineCellForItem() bool
	// optional
	OutlineView_ShouldShowOutlineCellForItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_ShouldShowCellExpansionForTableColumn_Item() bool
	// optional
	OutlineView_ShouldShowCellExpansionForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	ImplementsOutlineView_ShouldReorderColumn_ToColumn() bool
	// optional
	OutlineView_ShouldReorderColumn_ToColumn(outlineView OutlineView, columnIndex int, newColumnIndex int) bool
	ImplementsOutlineViewColumnDidMove() bool
	// optional
	OutlineViewColumnDidMove(notification foundation.Notification)
	ImplementsOutlineViewColumnDidResize() bool
	// optional
	OutlineViewColumnDidResize(notification foundation.Notification)
	ImplementsOutlineViewItemWillExpand() bool
	// optional
	OutlineViewItemWillExpand(notification foundation.Notification)
	ImplementsOutlineViewItemDidExpand() bool
	// optional
	OutlineViewItemDidExpand(notification foundation.Notification)
	ImplementsOutlineViewItemWillCollapse() bool
	// optional
	OutlineViewItemWillCollapse(notification foundation.Notification)
	ImplementsOutlineViewItemDidCollapse() bool
	// optional
	OutlineViewItemDidCollapse(notification foundation.Notification)
	ImplementsOutlineView_ShouldEditTableColumn_Item() bool
	// optional
	OutlineView_ShouldEditTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	ImplementsOutlineView_MouseDownInHeaderOfTableColumn() bool
	// optional
	OutlineView_MouseDownInHeaderOfTableColumn(outlineView OutlineView, tableColumn TableColumn)
	ImplementsOutlineView_DidClickTableColumn() bool
	// optional
	OutlineView_DidClickTableColumn(outlineView OutlineView, tableColumn TableColumn)
	ImplementsOutlineView_DidDragTableColumn() bool
	// optional
	OutlineView_DidDragTableColumn(outlineView OutlineView, tableColumn TableColumn)
	ImplementsOutlineView_HeightOfRowByItem() bool
	// optional
	OutlineView_HeightOfRowByItem(outlineView OutlineView, item objc.Object) float64
	ImplementsOutlineView_SizeToFitWidthOfColumn() bool
	// optional
	OutlineView_SizeToFitWidthOfColumn(outlineView OutlineView, column int) float64
	ImplementsOutlineView_TintConfigurationForItem() bool
	// optional
	OutlineView_TintConfigurationForItem(outlineView OutlineView, item objc.Object) ITintConfiguration
	ImplementsOutlineView_ShouldTrackCell_ForTableColumn_Item() bool
	// optional
	OutlineView_ShouldTrackCell_ForTableColumn_Item(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool
	ImplementsOutlineView_IsGroupItem() bool
	// optional
	OutlineView_IsGroupItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineView_DidAddRowView_ForRow() bool
	// optional
	OutlineView_DidAddRowView_ForRow(outlineView OutlineView, rowView TableRowView, row int)
	ImplementsOutlineView_DidRemoveRowView_ForRow() bool
	// optional
	OutlineView_DidRemoveRowView_ForRow(outlineView OutlineView, rowView TableRowView, row int)
	ImplementsOutlineView_RowViewForItem() bool
	// optional
	OutlineView_RowViewForItem(outlineView OutlineView, item objc.Object) ITableRowView
	ImplementsOutlineView_ViewForTableColumn_Item() bool
	// optional
	OutlineView_ViewForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView
}

func WrapOutlineViewDelegate(v OutlineViewDelegate) objc.Object {
	return objc.WrapAsProtocol("NSOutlineViewDelegate", v)
}

type OutlineViewDelegateBase struct {
	ControlTextEditingDelegateBase
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldExpandItem() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldExpandItem(outlineView OutlineView, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldCollapseItem() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldCollapseItem(outlineView OutlineView, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_TypeSelectStringForTableColumn_Item() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_TypeSelectStringForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_NextTypeSelectMatchFromItem_ToItem_ForString() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(outlineView OutlineView, event Event, searchString string) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldSelectTableColumn() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldSelectTableColumn(outlineView OutlineView, tableColumn TableColumn) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldSelectItem() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldSelectItem(outlineView OutlineView, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_SelectionIndexesForProposedSelection() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_SelectionIndexesForProposedSelection(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsSelectionShouldChangeInOutlineView() bool {
	return false
}

func (p *OutlineViewDelegateBase) SelectionShouldChangeInOutlineView(outlineView OutlineView) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineViewSelectionIsChanging() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineViewSelectionIsChanging(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineViewSelectionDidChange() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineViewSelectionDidChange(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_WillDisplayCell_ForTableColumn_Item() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_WillDisplayCell_ForTableColumn_Item(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_WillDisplayOutlineCell_ForTableColumn_Item() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_DataCellForTableColumn_Item() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_DataCellForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldShowOutlineCellForItem() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldShowOutlineCellForItem(outlineView OutlineView, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldShowCellExpansionForTableColumn_Item() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldShowCellExpansionForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldReorderColumn_ToColumn() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldReorderColumn_ToColumn(outlineView OutlineView, columnIndex int, newColumnIndex int) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineViewColumnDidMove() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineViewColumnDidMove(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineViewColumnDidResize() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineViewColumnDidResize(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineViewItemWillExpand() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineViewItemWillExpand(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineViewItemDidExpand() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineViewItemDidExpand(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineViewItemWillCollapse() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineViewItemWillCollapse(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineViewItemDidCollapse() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineViewItemDidCollapse(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldEditTableColumn_Item() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldEditTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_MouseDownInHeaderOfTableColumn() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_MouseDownInHeaderOfTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_DidClickTableColumn() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_DidClickTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_DidDragTableColumn() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_DidDragTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_HeightOfRowByItem() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_HeightOfRowByItem(outlineView OutlineView, item objc.Object) float64 {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_SizeToFitWidthOfColumn() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_SizeToFitWidthOfColumn(outlineView OutlineView, column int) float64 {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_TintConfigurationForItem() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_TintConfigurationForItem(outlineView OutlineView, item objc.Object) ITintConfiguration {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ShouldTrackCell_ForTableColumn_Item() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ShouldTrackCell_ForTableColumn_Item(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_IsGroupItem() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_IsGroupItem(outlineView OutlineView, item objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_DidAddRowView_ForRow() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_DidAddRowView_ForRow(outlineView OutlineView, rowView TableRowView, row int) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_DidRemoveRowView_ForRow() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_DidRemoveRowView_ForRow(outlineView OutlineView, rowView TableRowView, row int) {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_RowViewForItem() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_RowViewForItem(outlineView OutlineView, item objc.Object) ITableRowView {
	panic("unimpemented protocol method")
}

func (p *OutlineViewDelegateBase) ImplementsOutlineView_ViewForTableColumn_Item() bool {
	return false
}

func (p *OutlineViewDelegateBase) OutlineView_ViewForTableColumn_Item(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView {
	panic("unimpemented protocol method")
}
