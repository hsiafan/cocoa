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

type DraggingInfoWrapper struct {
	objc.Object
}

// deprecated
func (d_ DraggingInfoWrapper) NamesOfPromisedFilesDroppedAtDestination(dropDestination foundation.IURL) []string {
	rv := objc.CallMethod[[]string](d_, objc.GetSelector("namesOfPromisedFilesDroppedAtDestination:"), objc.ExtractPtr(dropDestination))
	return rv
}

func (d_ DraggingInfoWrapper) SlideDraggedImageTo(screenPoint foundation.Point) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("slideDraggedImageTo:"), screenPoint)
}

func (d_ DraggingInfoWrapper) EnumerateDraggingItemsWithOptions_ForView_Classes_SearchOptions_UsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool)) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:"), enumOpts, objc.ExtractPtr(view), classArray, searchOptions, block)
}

func (d_ DraggingInfoWrapper) ResetSpringLoading() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("resetSpringLoading"))
}

func (d_ DraggingInfoWrapper) DraggingPasteboard() Pasteboard {
	rv := objc.CallMethod[Pasteboard](d_, objc.GetSelector("draggingPasteboard"))
	return rv
}

func (d_ DraggingInfoWrapper) DraggingSequenceNumber() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("draggingSequenceNumber"))
	return rv
}

func (d_ DraggingInfoWrapper) DraggingSource() objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("draggingSource"))
	return rv
}

func (d_ DraggingInfoWrapper) DraggingSourceOperationMask() DragOperation {
	rv := objc.CallMethod[DragOperation](d_, objc.GetSelector("draggingSourceOperationMask"))
	return rv
}

func (d_ DraggingInfoWrapper) DraggingLocation() foundation.Point {
	rv := objc.CallMethod[foundation.Point](d_, objc.GetSelector("draggingLocation"))
	return rv
}

func (d_ DraggingInfoWrapper) DraggingDestinationWindow() Window {
	rv := objc.CallMethod[Window](d_, objc.GetSelector("draggingDestinationWindow"))
	return rv
}

func (d_ DraggingInfoWrapper) SetNumberOfValidItemsForDrop(value int) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setNumberOfValidItemsForDrop:"), value)
}

func (d_ DraggingInfoWrapper) NumberOfValidItemsForDrop() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("numberOfValidItemsForDrop"))
	return rv
}

// deprecated
func (d_ DraggingInfoWrapper) DraggedImage() Image {
	rv := objc.CallMethod[Image](d_, objc.GetSelector("draggedImage"))
	return rv
}

func (d_ DraggingInfoWrapper) DraggedImageLocation() foundation.Point {
	rv := objc.CallMethod[foundation.Point](d_, objc.GetSelector("draggedImageLocation"))
	return rv
}

func (d_ DraggingInfoWrapper) SetAnimatesToDestination(value bool) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setAnimatesToDestination:"), value)
}

func (d_ DraggingInfoWrapper) AnimatesToDestination() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("animatesToDestination"))
	return rv
}

func (d_ DraggingInfoWrapper) SetDraggingFormation(value DraggingFormation) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("setDraggingFormation:"), value)
}

func (d_ DraggingInfoWrapper) DraggingFormation() DraggingFormation {
	rv := objc.CallMethod[DraggingFormation](d_, objc.GetSelector("draggingFormation"))
	return rv
}

func (d_ DraggingInfoWrapper) SpringLoadingHighlight() SpringLoadingHighlight {
	rv := objc.CallMethod[SpringLoadingHighlight](d_, objc.GetSelector("springLoadingHighlight"))
	return rv
}
