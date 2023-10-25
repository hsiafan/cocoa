// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type DraggingInfo interface {
	// required
	// deprecated
	NamesOfPromisedFilesDroppedAtDestination(dropDestination foundation.URL) []string
	// required
	SlideDraggedImageTo(screenPoint foundation.Point)
	// required
	EnumerateDraggingItemsWithOptions_ForView_Classes_SearchOptions_UsingBlock(enumOpts DraggingItemEnumerationOptions, view View, classArray []objc.Class, searchOptions map[PasteboardReadingOptionKey]objc.Object, block func(draggingItem IDraggingItem, idx int, stop *bool))
	// required
	ResetSpringLoading()
	ImplementsDraggingPasteboard() bool
	// optional
	DraggingPasteboard() IPasteboard
	ImplementsDraggingSequenceNumber() bool
	// optional
	DraggingSequenceNumber() int
	ImplementsDraggingSource() bool
	// optional
	DraggingSource() objc.IObject
	ImplementsDraggingSourceOperationMask() bool
	// optional
	DraggingSourceOperationMask() DragOperation
	ImplementsDraggingLocation() bool
	// optional
	DraggingLocation() foundation.Point
	ImplementsDraggingDestinationWindow() bool
	// optional
	DraggingDestinationWindow() IWindow
	ImplementsSetNumberOfValidItemsForDrop() bool
	// optional
	SetNumberOfValidItemsForDrop(value int)
	ImplementsNumberOfValidItemsForDrop() bool
	// optional
	NumberOfValidItemsForDrop() int
	ImplementsDraggedImage() bool
	// optional
	// deprecated
	DraggedImage() IImage
	ImplementsDraggedImageLocation() bool
	// optional
	DraggedImageLocation() foundation.Point
	ImplementsSetAnimatesToDestination() bool
	// optional
	SetAnimatesToDestination(value bool)
	ImplementsAnimatesToDestination() bool
	// optional
	AnimatesToDestination() bool
	ImplementsSetDraggingFormation() bool
	// optional
	SetDraggingFormation(value DraggingFormation)
	ImplementsDraggingFormation() bool
	// optional
	DraggingFormation() DraggingFormation
	ImplementsSpringLoadingHighlight() bool
	// optional
	SpringLoadingHighlight() SpringLoadingHighlight
}

func WrapDraggingInfo(v DraggingInfo) objc.Object {
	return objc.WrapAsProtocol("NSDraggingInfo", v)
}

type DraggingInfoBase struct {
}

func (p *DraggingInfoBase) ImplementsDraggingPasteboard() bool {
	return false
}

func (p *DraggingInfoBase) DraggingPasteboard() IPasteboard {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsDraggingSequenceNumber() bool {
	return false
}

func (p *DraggingInfoBase) DraggingSequenceNumber() int {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsDraggingSource() bool {
	return false
}

func (p *DraggingInfoBase) DraggingSource() objc.IObject {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsDraggingSourceOperationMask() bool {
	return false
}

func (p *DraggingInfoBase) DraggingSourceOperationMask() DragOperation {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsDraggingLocation() bool {
	return false
}

func (p *DraggingInfoBase) DraggingLocation() foundation.Point {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsDraggingDestinationWindow() bool {
	return false
}

func (p *DraggingInfoBase) DraggingDestinationWindow() IWindow {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsSetNumberOfValidItemsForDrop() bool {
	return false
}

func (p *DraggingInfoBase) SetNumberOfValidItemsForDrop(value int) {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsNumberOfValidItemsForDrop() bool {
	return false
}

func (p *DraggingInfoBase) NumberOfValidItemsForDrop() int {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsDraggedImage() bool {
	return false
}

// deprecated
func (p *DraggingInfoBase) DraggedImage() IImage {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsDraggedImageLocation() bool {
	return false
}

func (p *DraggingInfoBase) DraggedImageLocation() foundation.Point {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsSetAnimatesToDestination() bool {
	return false
}

func (p *DraggingInfoBase) SetAnimatesToDestination(value bool) {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsAnimatesToDestination() bool {
	return false
}

func (p *DraggingInfoBase) AnimatesToDestination() bool {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsSetDraggingFormation() bool {
	return false
}

func (p *DraggingInfoBase) SetDraggingFormation(value DraggingFormation) {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsDraggingFormation() bool {
	return false
}

func (p *DraggingInfoBase) DraggingFormation() DraggingFormation {
	panic("unimpemented protocol method")
}

func (p *DraggingInfoBase) ImplementsSpringLoadingHighlight() bool {
	return false
}

func (p *DraggingInfoBase) SpringLoadingHighlight() SpringLoadingHighlight {
	panic("unimpemented protocol method")
}
