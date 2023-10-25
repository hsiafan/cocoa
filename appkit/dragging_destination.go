// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type DraggingDestination interface {
	ImplementsDraggingEntered() bool
	// optional
	DraggingEntered(sender objc.Object) DragOperation
	ImplementsWantsPeriodicDraggingUpdates() bool
	// optional
	WantsPeriodicDraggingUpdates() bool
	ImplementsDraggingUpdated() bool
	// optional
	DraggingUpdated(sender objc.Object) DragOperation
	ImplementsDraggingEnded() bool
	// optional
	DraggingEnded(sender objc.Object)
	ImplementsDraggingExited() bool
	// optional
	DraggingExited(sender objc.Object)
	ImplementsPrepareForDragOperation() bool
	// optional
	PrepareForDragOperation(sender objc.Object) bool
	ImplementsPerformDragOperation() bool
	// optional
	PerformDragOperation(sender objc.Object) bool
	ImplementsConcludeDragOperation() bool
	// optional
	ConcludeDragOperation(sender objc.Object)
	ImplementsUpdateDraggingItemsForDrag() bool
	// optional
	UpdateDraggingItemsForDrag(sender objc.Object)
}

func WrapDraggingDestination(v DraggingDestination) objc.Object {
	return objc.WrapAsProtocol("NSDraggingDestination", v)
}

type DraggingDestinationBase struct {
}

func (p *DraggingDestinationBase) ImplementsDraggingEntered() bool {
	return false
}

func (p *DraggingDestinationBase) DraggingEntered(sender objc.Object) DragOperation {
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

func (p *DraggingDestinationBase) DraggingUpdated(sender objc.Object) DragOperation {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsDraggingEnded() bool {
	return false
}

func (p *DraggingDestinationBase) DraggingEnded(sender objc.Object) {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsDraggingExited() bool {
	return false
}

func (p *DraggingDestinationBase) DraggingExited(sender objc.Object) {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsPrepareForDragOperation() bool {
	return false
}

func (p *DraggingDestinationBase) PrepareForDragOperation(sender objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsPerformDragOperation() bool {
	return false
}

func (p *DraggingDestinationBase) PerformDragOperation(sender objc.Object) bool {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsConcludeDragOperation() bool {
	return false
}

func (p *DraggingDestinationBase) ConcludeDragOperation(sender objc.Object) {
	panic("unimpemented protocol method")
}

func (p *DraggingDestinationBase) ImplementsUpdateDraggingItemsForDrag() bool {
	return false
}

func (p *DraggingDestinationBase) UpdateDraggingItemsForDrag(sender objc.Object) {
	panic("unimpemented protocol method")
}
