// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type DraggingDestinationWrapper struct {
	objc.Object
}

func (d_ *DraggingDestinationWrapper) ImplementsDraggingEntered() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingEntered:"))
}

func (d_ DraggingDestinationWrapper) DraggingEntered(sender DraggingInfo) DragOperation {
	po := ffi.CreateProtocol("NSDraggingInfo", sender)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](d_, "draggingEntered:", po)
	return rv
}

func (d_ *DraggingDestinationWrapper) ImplementsWantsPeriodicDraggingUpdates() bool {
	return d_.RespondsToSelector(objc.GetSelector("wantsPeriodicDraggingUpdates"))
}

func (d_ DraggingDestinationWrapper) WantsPeriodicDraggingUpdates() bool {
	rv := ffi.CallMethod[bool](d_, "wantsPeriodicDraggingUpdates")
	return rv
}

func (d_ *DraggingDestinationWrapper) ImplementsDraggingUpdated() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingUpdated:"))
}

func (d_ DraggingDestinationWrapper) DraggingUpdated(sender DraggingInfo) DragOperation {
	po := ffi.CreateProtocol("NSDraggingInfo", sender)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](d_, "draggingUpdated:", po)
	return rv
}

func (d_ *DraggingDestinationWrapper) ImplementsDraggingEnded() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingEnded:"))
}

func (d_ DraggingDestinationWrapper) DraggingEnded(sender DraggingInfo) {
	po := ffi.CreateProtocol("NSDraggingInfo", sender)
	defer po.Release()
	ffi.CallMethod[ffi.Void](d_, "draggingEnded:", po)
}

func (d_ *DraggingDestinationWrapper) ImplementsDraggingExited() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingExited:"))
}

func (d_ DraggingDestinationWrapper) DraggingExited(sender DraggingInfo) {
	po := ffi.CreateProtocol("NSDraggingInfo", sender)
	defer po.Release()
	ffi.CallMethod[ffi.Void](d_, "draggingExited:", po)
}

func (d_ *DraggingDestinationWrapper) ImplementsPrepareForDragOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("prepareForDragOperation:"))
}

func (d_ DraggingDestinationWrapper) PrepareForDragOperation(sender DraggingInfo) bool {
	po := ffi.CreateProtocol("NSDraggingInfo", sender)
	defer po.Release()
	rv := ffi.CallMethod[bool](d_, "prepareForDragOperation:", po)
	return rv
}

func (d_ *DraggingDestinationWrapper) ImplementsPerformDragOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("performDragOperation:"))
}

func (d_ DraggingDestinationWrapper) PerformDragOperation(sender DraggingInfo) bool {
	po := ffi.CreateProtocol("NSDraggingInfo", sender)
	defer po.Release()
	rv := ffi.CallMethod[bool](d_, "performDragOperation:", po)
	return rv
}

func (d_ *DraggingDestinationWrapper) ImplementsConcludeDragOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("concludeDragOperation:"))
}

func (d_ DraggingDestinationWrapper) ConcludeDragOperation(sender DraggingInfo) {
	po := ffi.CreateProtocol("NSDraggingInfo", sender)
	defer po.Release()
	ffi.CallMethod[ffi.Void](d_, "concludeDragOperation:", po)
}

func (d_ *DraggingDestinationWrapper) ImplementsUpdateDraggingItemsForDrag() bool {
	return d_.RespondsToSelector(objc.GetSelector("updateDraggingItemsForDrag:"))
}

func (d_ DraggingDestinationWrapper) UpdateDraggingItemsForDrag(sender DraggingInfo) {
	po := ffi.CreateProtocol("NSDraggingInfo", sender)
	defer po.Release()
	ffi.CallMethod[ffi.Void](d_, "updateDraggingItemsForDrag:", po)
}
