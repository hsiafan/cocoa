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

func WrapDraggingSource(v DraggingSource) objc.Object {
	return objc.WrapAsProtocol("NSDraggingSource", v)
}

type DraggingSourceBase struct {
}

func (p *DraggingSourceBase) ImplementsDraggingSession_EndedAtPoint_Operation() bool {
	return false
}

func (p *DraggingSourceBase) DraggingSession_EndedAtPoint_Operation(session DraggingSession, screenPoint foundation.Point, operation DragOperation) {
	panic("unimpemented protocol method")
}

func (p *DraggingSourceBase) ImplementsDraggingSession_MovedToPoint() bool {
	return false
}

func (p *DraggingSourceBase) DraggingSession_MovedToPoint(session DraggingSession, screenPoint foundation.Point) {
	panic("unimpemented protocol method")
}

func (p *DraggingSourceBase) ImplementsDraggingSession_WillBeginAtPoint() bool {
	return false
}

func (p *DraggingSourceBase) DraggingSession_WillBeginAtPoint(session DraggingSession, screenPoint foundation.Point) {
	panic("unimpemented protocol method")
}

func (p *DraggingSourceBase) ImplementsIgnoreModifierKeysForDraggingSession() bool {
	return false
}

func (p *DraggingSourceBase) IgnoreModifierKeysForDraggingSession(session DraggingSession) bool {
	panic("unimpemented protocol method")
}

type DraggingSourceWrapper struct {
	objc.Object
}

func (d_ DraggingSourceWrapper) DraggingSession_SourceOperationMaskForDraggingContext(session IDraggingSession, context DraggingContext) DragOperation {
	rv := objc.CallMethod[DragOperation](d_, objc.GetSelector("draggingSession:sourceOperationMaskForDraggingContext:"), objc.ExtractPtr(session), context)
	return rv
}

func (d_ DraggingSourceWrapper) DraggingSession_EndedAtPoint_Operation(session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingSession:endedAtPoint:operation:"), objc.ExtractPtr(session), screenPoint, operation)
}

func (d_ DraggingSourceWrapper) DraggingSession_MovedToPoint(session IDraggingSession, screenPoint foundation.Point) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingSession:movedToPoint:"), objc.ExtractPtr(session), screenPoint)
}

func (d_ DraggingSourceWrapper) DraggingSession_WillBeginAtPoint(session IDraggingSession, screenPoint foundation.Point) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("draggingSession:willBeginAtPoint:"), objc.ExtractPtr(session), screenPoint)
}

func (d_ DraggingSourceWrapper) IgnoreModifierKeysForDraggingSession(session IDraggingSession) bool {
	rv := objc.CallMethod[bool](d_, objc.GetSelector("ignoreModifierKeysForDraggingSession:"), objc.ExtractPtr(session))
	return rv
}
