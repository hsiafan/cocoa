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

type OutlineViewDelegateWrapper struct {
	objc.Object
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldExpandItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldExpandItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldCollapseItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldCollapseItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_TypeSelectStringForTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("outlineView:typeSelectStringForTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_NextTypeSelectMatchFromItem_ToItem_ForString(outlineView IOutlineView, startItem objc.IObject, endItem objc.IObject, searchString string) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:nextTypeSelectMatchFromItem:toItem:forString:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(startItem), objc.ExtractPtr(endItem), searchString)
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldTypeSelectForEvent_WithCurrentSearchString(outlineView IOutlineView, event IEvent, searchString string) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldTypeSelectForEvent:withCurrentSearchString:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(event), searchString)
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ToolTipForCell_Rect_TableColumn_Item_MouseLocation(outlineView IOutlineView, cell ICell, rect *foundation.Rect, tableColumn ITableColumn, item objc.IObject, mouseLocation foundation.Point) string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("outlineView:toolTipForCell:rect:tableColumn:item:mouseLocation:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(cell), rect, objc.ExtractPtr(tableColumn), objc.ExtractPtr(item), mouseLocation)
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldSelectTableColumn(outlineView IOutlineView, tableColumn ITableColumn) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldSelectTableColumn:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldSelectItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldSelectItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_SelectionIndexesForProposedSelection(outlineView IOutlineView, proposedSelectionIndexes foundation.IIndexSet) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](o_, objc.GetSelector("outlineView:selectionIndexesForProposedSelection:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(proposedSelectionIndexes))
	return rv
}

func (o_ OutlineViewDelegateWrapper) SelectionShouldChangeInOutlineView(outlineView IOutlineView) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("selectionShouldChangeInOutlineView:"), objc.ExtractPtr(outlineView))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineViewSelectionIsChanging(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewSelectionIsChanging:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewSelectionDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewSelectionDidChange:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_WillDisplayCell_ForTableColumn_Item(outlineView IOutlineView, cell objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:willDisplayCell:forTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_WillDisplayOutlineCell_ForTableColumn_Item(outlineView IOutlineView, cell objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:willDisplayOutlineCell:forTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DataCellForTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) Cell {
	rv := objc.CallMethod[Cell](o_, objc.GetSelector("outlineView:dataCellForTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldShowOutlineCellForItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldShowOutlineCellForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldShowCellExpansionForTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldShowCellExpansionForTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldReorderColumn_ToColumn(outlineView IOutlineView, columnIndex int, newColumnIndex int) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldReorderColumn:toColumn:"), objc.ExtractPtr(outlineView), columnIndex, newColumnIndex)
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineViewColumnDidMove(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewColumnDidMove:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewColumnDidResize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewColumnDidResize:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemWillExpand(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewItemWillExpand:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemDidExpand(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewItemDidExpand:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemWillCollapse(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewItemWillCollapse:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemDidCollapse(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewItemDidCollapse:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldEditTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldEditTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_MouseDownInHeaderOfTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:mouseDownInHeaderOfTableColumn:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DidClickTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:didClickTableColumn:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DidDragTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:didDragTableColumn:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn))
}

func (o_ OutlineViewDelegateWrapper) OutlineView_HeightOfRowByItem(outlineView IOutlineView, item objc.IObject) float64 {
	rv := objc.CallMethod[float64](o_, objc.GetSelector("outlineView:heightOfRowByItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_SizeToFitWidthOfColumn(outlineView IOutlineView, column int) float64 {
	rv := objc.CallMethod[float64](o_, objc.GetSelector("outlineView:sizeToFitWidthOfColumn:"), objc.ExtractPtr(outlineView), column)
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_TintConfigurationForItem(outlineView IOutlineView, item objc.IObject) TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](o_, objc.GetSelector("outlineView:tintConfigurationForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ShouldTrackCell_ForTableColumn_Item(outlineView IOutlineView, cell ICell, tableColumn ITableColumn, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldTrackCell:forTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_IsGroupItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:isGroupItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DidAddRowView_ForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:didAddRowView:forRow:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(rowView), row)
}

func (o_ OutlineViewDelegateWrapper) OutlineView_DidRemoveRowView_ForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:didRemoveRowView:forRow:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(rowView), row)
}

func (o_ OutlineViewDelegateWrapper) OutlineView_RowViewForItem(outlineView IOutlineView, item objc.IObject) TableRowView {
	rv := objc.CallMethod[TableRowView](o_, objc.GetSelector("outlineView:rowViewForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) OutlineView_ViewForTableColumn_Item(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) View {
	rv := objc.CallMethod[View](o_, objc.GetSelector("outlineView:viewForTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) Control_IsValidObject(control IControl, obj objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("control:isValidObject:"), objc.ExtractPtr(control), objc.ExtractPtr(obj))
	return rv
}

func (o_ OutlineViewDelegateWrapper) Control_DidFailToValidatePartialString_ErrorDescription(control IControl, string_ string, error string) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("control:didFailToValidatePartialString:errorDescription:"), objc.ExtractPtr(control), string_, error)
}

func (o_ OutlineViewDelegateWrapper) Control_DidFailToFormatString_ErrorDescription(control IControl, string_ string, error string) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("control:didFailToFormatString:errorDescription:"), objc.ExtractPtr(control), string_, error)
	return rv
}

func (o_ OutlineViewDelegateWrapper) Control_TextShouldBeginEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("control:textShouldBeginEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (o_ OutlineViewDelegateWrapper) Control_TextShouldEndEditing(control IControl, fieldEditor IText) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("control:textShouldEndEditing:"), objc.ExtractPtr(control), objc.ExtractPtr(fieldEditor))
	return rv
}

func (o_ OutlineViewDelegateWrapper) Control_TextView_Completions_ForPartialWordRange_IndexOfSelectedItem(control IControl, textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](o_, objc.GetSelector("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (o_ OutlineViewDelegateWrapper) Control_TextView_DoCommandBySelector(control IControl, textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("control:textView:doCommandBySelector:"), objc.ExtractPtr(control), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (o_ OutlineViewDelegateWrapper) ControlTextDidBeginEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("controlTextDidBeginEditing:"), objc.ExtractPtr(obj))
}

func (o_ OutlineViewDelegateWrapper) ControlTextDidChange(obj foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("controlTextDidChange:"), objc.ExtractPtr(obj))
}

func (o_ OutlineViewDelegateWrapper) ControlTextDidEndEditing(obj foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("controlTextDidEndEditing:"), objc.ExtractPtr(obj))
}
