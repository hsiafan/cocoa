// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type DraggingSource interface {
	// required
	DraggingSession_SourceOperationMaskForDraggingContext(session DraggingSession, context DraggingContext) DragOperation
	ImplementsDraggingSession_EndedAtPoint_Operation() bool
	// optional
	DraggingSession_EndedAtPoint_Operation(session DraggingSession, screenPoint foundation.Point, operation DragOperation)
	ImplementsDraggingSession_MovedToPoint() bool
	// optional
	DraggingSession_MovedToPoint(session DraggingSession, screenPoint foundation.Point)
	ImplementsDraggingSession_WillBeginAtPoint() bool
	// optional
	DraggingSession_WillBeginAtPoint(session DraggingSession, screenPoint foundation.Point)
	ImplementsIgnoreModifierKeysForDraggingSession() bool
	// optional
	IgnoreModifierKeysForDraggingSession(session DraggingSession) bool
}

type DraggingSourceWrapper struct {
	objc.Object
}

func (d_ DraggingSourceWrapper) DraggingSession_SourceOperationMaskForDraggingContext(session IDraggingSession, context DraggingContext) DragOperation {
	rv := objc.CallMethod[DragOperation](d_, objc.GetSelector("draggingSession:sourceOperationMaskForDraggingContext:"), session, context)
	return rv
}

func (d_ *DraggingSourceWrapper) ImplementsDraggingSession_EndedAtPoint_Operation() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSession:endedAtPoint:operation:"))
}

func (d_ DraggingSourceWrapper) DraggingSession_EndedAtPoint_Operation(session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingSession:endedAtPoint:operation:"), session, screenPoint, operation)
}

func (d_ *DraggingSourceWrapper) ImplementsDraggingSession_MovedToPoint() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSession:movedToPoint:"))
}

func (d_ DraggingSourceWrapper) DraggingSession_MovedToPoint(session IDraggingSession, screenPoint foundation.Point) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingSession:movedToPoint:"), session, screenPoint)
}

func (d_ *DraggingSourceWrapper) ImplementsDraggingSession_WillBeginAtPoint() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSession:willBeginAtPoint:"))
}

func (d_ DraggingSourceWrapper) DraggingSession_WillBeginAtPoint(session IDraggingSession, screenPoint foundation.Point) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingSession:willBeginAtPoint:"), session, screenPoint)
}

func (d_ *DraggingSourceWrapper) ImplementsIgnoreModifierKeysForDraggingSession() bool {
	return d_.RespondsToSelector(objc.GetSelector("ignoreModifierKeysForDraggingSession:"))
}

func (d_ DraggingSourceWrapper) IgnoreModifierKeysForDraggingSession(session IDraggingSession) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("ignoreModifierKeysForDraggingSession:"), session)
	return rv
}
