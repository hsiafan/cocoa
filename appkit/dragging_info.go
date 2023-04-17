// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type DraggingInfoWrapper struct {
	objc.Object
}

// deprecated
func (d_ DraggingInfoWrapper) NamesOfPromisedFilesDroppedAtDestination(dropDestination foundation.IURL) []string {
	rv := ffi.CallMethod[[]string](d_, "namesOfPromisedFilesDroppedAtDestination:", dropDestination)
	return rv
}

func (d_ DraggingInfoWrapper) SlideDraggedImageTo(screenPoint foundation.Point) {
	ffi.CallMethod[ffi.Void](d_, "slideDraggedImageTo:", screenPoint)
}

func (d_ DraggingInfoWrapper) EnumerateDraggingItemsWithOptions_ForView_Classes_SearchOptions_UsingBlock(enumOpts DraggingItemEnumerationOptions, view IView, classArray []objc.IClass, searchOptions map[PasteboardReadingOptionKey]objc.IObject, block func(draggingItem DraggingItem, idx int, stop *bool)) {
	ffi.CallMethod[ffi.Void](d_, "enumerateDraggingItemsWithOptions:forView:classes:searchOptions:usingBlock:", enumOpts, view, classArray, searchOptions, block)
}

func (d_ DraggingInfoWrapper) ResetSpringLoading() {
	ffi.CallMethod[ffi.Void](d_, "resetSpringLoading")
}

func (d_ *DraggingInfoWrapper) ImplementsDraggingPasteboard() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingPasteboard"))
}

func (d_ DraggingInfoWrapper) DraggingPasteboard() Pasteboard {
	rv := ffi.CallMethod[Pasteboard](d_, "draggingPasteboard")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsDraggingSequenceNumber() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSequenceNumber"))
}

func (d_ DraggingInfoWrapper) DraggingSequenceNumber() int {
	rv := ffi.CallMethod[int](d_, "draggingSequenceNumber")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsDraggingSource() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSource"))
}

func (d_ DraggingInfoWrapper) DraggingSource() objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "draggingSource")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsDraggingSourceOperationMask() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSourceOperationMask"))
}

func (d_ DraggingInfoWrapper) DraggingSourceOperationMask() DragOperation {
	rv := ffi.CallMethod[DragOperation](d_, "draggingSourceOperationMask")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsDraggingLocation() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingLocation"))
}

func (d_ DraggingInfoWrapper) DraggingLocation() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](d_, "draggingLocation")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsDraggingDestinationWindow() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingDestinationWindow"))
}

func (d_ DraggingInfoWrapper) DraggingDestinationWindow() Window {
	rv := ffi.CallMethod[Window](d_, "draggingDestinationWindow")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsSetNumberOfValidItemsForDrop() bool {
	return d_.RespondsToSelector(objc.GetSelector("setNumberOfValidItemsForDrop:"))
}

func (d_ DraggingInfoWrapper) SetNumberOfValidItemsForDrop(value int) {
	ffi.CallMethod[ffi.Void](d_, "setNumberOfValidItemsForDrop:", value)
}

func (d_ *DraggingInfoWrapper) ImplementsNumberOfValidItemsForDrop() bool {
	return d_.RespondsToSelector(objc.GetSelector("numberOfValidItemsForDrop"))
}

func (d_ DraggingInfoWrapper) NumberOfValidItemsForDrop() int {
	rv := ffi.CallMethod[int](d_, "numberOfValidItemsForDrop")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsDraggedImage() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggedImage"))
}

// deprecated
func (d_ DraggingInfoWrapper) DraggedImage() Image {
	rv := ffi.CallMethod[Image](d_, "draggedImage")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsDraggedImageLocation() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggedImageLocation"))
}

func (d_ DraggingInfoWrapper) DraggedImageLocation() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](d_, "draggedImageLocation")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsSetAnimatesToDestination() bool {
	return d_.RespondsToSelector(objc.GetSelector("setAnimatesToDestination:"))
}

func (d_ DraggingInfoWrapper) SetAnimatesToDestination(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setAnimatesToDestination:", value)
}

func (d_ *DraggingInfoWrapper) ImplementsAnimatesToDestination() bool {
	return d_.RespondsToSelector(objc.GetSelector("animatesToDestination"))
}

func (d_ DraggingInfoWrapper) AnimatesToDestination() bool {
	rv := ffi.CallMethod[bool](d_, "animatesToDestination")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsSetDraggingFormation() bool {
	return d_.RespondsToSelector(objc.GetSelector("setDraggingFormation:"))
}

func (d_ DraggingInfoWrapper) SetDraggingFormation(value DraggingFormation) {
	ffi.CallMethod[ffi.Void](d_, "setDraggingFormation:", value)
}

func (d_ *DraggingInfoWrapper) ImplementsDraggingFormation() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingFormation"))
}

func (d_ DraggingInfoWrapper) DraggingFormation() DraggingFormation {
	rv := ffi.CallMethod[DraggingFormation](d_, "draggingFormation")
	return rv
}

func (d_ *DraggingInfoWrapper) ImplementsSpringLoadingHighlight() bool {
	return d_.RespondsToSelector(objc.GetSelector("springLoadingHighlight"))
}

func (d_ DraggingInfoWrapper) SpringLoadingHighlight() SpringLoadingHighlight {
	rv := ffi.CallMethod[SpringLoadingHighlight](d_, "springLoadingHighlight")
	return rv
}
