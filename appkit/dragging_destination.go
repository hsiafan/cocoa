// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type DraggingDestination interface {
	ImplementsDraggingEntered() bool
	// optional
	DraggingEntered(sender DraggingInfoWrapper) DragOperation
	ImplementsWantsPeriodicDraggingUpdates() bool
	// optional
	WantsPeriodicDraggingUpdates() bool
	ImplementsDraggingUpdated() bool
	// optional
	DraggingUpdated(sender DraggingInfoWrapper) DragOperation
	ImplementsDraggingEnded() bool
	// optional
	DraggingEnded(sender DraggingInfoWrapper)
	ImplementsDraggingExited() bool
	// optional
	DraggingExited(sender DraggingInfoWrapper)
	ImplementsPrepareForDragOperation() bool
	// optional
	PrepareForDragOperation(sender DraggingInfoWrapper) bool
	ImplementsPerformDragOperation() bool
	// optional
	PerformDragOperation(sender DraggingInfoWrapper) bool
	ImplementsConcludeDragOperation() bool
	// optional
	ConcludeDragOperation(sender DraggingInfoWrapper)
	ImplementsUpdateDraggingItemsForDrag() bool
	// optional
	UpdateDraggingItemsForDrag(sender DraggingInfoWrapper)
}

func WrapDraggingDestination(v DraggingDestination) objc.Object {
	return objc.WrapAsProtocol("NSDraggingDestination", v)
}

type DraggingDestinationBase struct {
}

func (p *DraggingDestinationBase) ImplementsDraggingEntered() bool {
	return false
}

func (p *DraggingDestinationBase) DraggingEntered(sender DraggingInfoWrapper) DragOperation {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsWantsPeriodicDraggingUpdates() bool {
	return false
}

func (p *DraggingDestinationBase) WantsPeriodicDraggingUpdates() bool {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsDraggingUpdated() bool {
	return false
}

func (p *DraggingDestinationBase) DraggingUpdated(sender DraggingInfoWrapper) DragOperation {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsDraggingEnded() bool {
	return false
}

func (p *DraggingDestinationBase) DraggingEnded(sender DraggingInfoWrapper) {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsDraggingExited() bool {
	return false
}

func (p *DraggingDestinationBase) DraggingExited(sender DraggingInfoWrapper) {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsPrepareForDragOperation() bool {
	return false
}

func (p *DraggingDestinationBase) PrepareForDragOperation(sender DraggingInfoWrapper) bool {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsPerformDragOperation() bool {
	return false
}

func (p *DraggingDestinationBase) PerformDragOperation(sender DraggingInfoWrapper) bool {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsConcludeDragOperation() bool {
	return false
}

func (p *DraggingDestinationBase) ConcludeDragOperation(sender DraggingInfoWrapper) {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsUpdateDraggingItemsForDrag() bool {
	return false
}

func (p *DraggingDestinationBase) UpdateDraggingItemsForDrag(sender DraggingInfoWrapper) {
	panic("unimpemented protocol method")
}

type DraggingDestinationWrapper struct {
	objc.Object
}

func (d_ DraggingDestinationWrapper) DraggingEntered(sender objc.IObject) DragOperation {
	rv := objc.CallMethod[DragOperation](d_, objc.GetSelector("draggingEntered:"), objc.ExtractPtr(sender))
	return rv
}

func (d_ DraggingDestinationWrapper) WantsPeriodicDraggingUpdates() bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("wantsPeriodicDraggingUpdates"))
	return rv
}

func (d_ DraggingDestinationWrapper) DraggingUpdated(sender objc.IObject) DragOperation {
	rv := objc.CallMethod[DragOperation](d_, objc.GetSelector("draggingUpdated:"), objc.ExtractPtr(sender))
	return rv
}

func (d_ DraggingDestinationWrapper) DraggingEnded(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingEnded:"), objc.ExtractPtr(sender))
}

func (d_ DraggingDestinationWrapper) DraggingExited(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingExited:"), objc.ExtractPtr(sender))
}

func (d_ DraggingDestinationWrapper) PrepareForDragOperation(sender objc.IObject) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("prepareForDragOperation:"), objc.ExtractPtr(sender))
	return rv
}

func (d_ DraggingDestinationWrapper) PerformDragOperation(sender objc.IObject) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("performDragOperation:"), objc.ExtractPtr(sender))
	return rv
}

func (d_ DraggingDestinationWrapper) ConcludeDragOperation(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("concludeDragOperation:"), objc.ExtractPtr(sender))
}

func (d_ DraggingDestinationWrapper) UpdateDraggingItemsForDrag(sender objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("updateDraggingItemsForDrag:"), objc.ExtractPtr(sender))
}
