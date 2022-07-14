// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := ffi.CallMethod[DragOperation](d_, "draggingSession:sourceOperationMaskForDraggingContext:", session, context)
	return rv
}

func (d_ *DraggingSourceWrapper) ImplementsDraggingSession_EndedAtPoint_Operation() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSession:endedAtPoint:operation:"))
}

func (d_ DraggingSourceWrapper) DraggingSession_EndedAtPoint_Operation(session IDraggingSession, screenPoint foundation.Point, operation DragOperation) {
	ffi.CallMethod[ffi.Void](d_, "draggingSession:endedAtPoint:operation:", session, screenPoint, operation)
}

func (d_ *DraggingSourceWrapper) ImplementsDraggingSession_MovedToPoint() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSession:movedToPoint:"))
}

func (d_ DraggingSourceWrapper) DraggingSession_MovedToPoint(session IDraggingSession, screenPoint foundation.Point) {
	ffi.CallMethod[ffi.Void](d_, "draggingSession:movedToPoint:", session, screenPoint)
}

func (d_ *DraggingSourceWrapper) ImplementsDraggingSession_WillBeginAtPoint() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingSession:willBeginAtPoint:"))
}

func (d_ DraggingSourceWrapper) DraggingSession_WillBeginAtPoint(session IDraggingSession, screenPoint foundation.Point) {
	ffi.CallMethod[ffi.Void](d_, "draggingSession:willBeginAtPoint:", session, screenPoint)
}

func (d_ *DraggingSourceWrapper) ImplementsIgnoreModifierKeysForDraggingSession() bool {
	return d_.RespondsToSelector(objc.GetSelector("ignoreModifierKeysForDraggingSession:"))
}

func (d_ DraggingSourceWrapper) IgnoreModifierKeysForDraggingSession(session IDraggingSession) bool {
	rv := ffi.CallMethod[bool](d_, "ignoreModifierKeysForDraggingSession:", session)
	return rv
}

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
	po := ffi.CreateProtocol(sender)
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
	po := ffi.CreateProtocol(sender)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](d_, "draggingUpdated:", po)
	return rv
}

func (d_ *DraggingDestinationWrapper) ImplementsDraggingEnded() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingEnded:"))
}

func (d_ DraggingDestinationWrapper) DraggingEnded(sender DraggingInfo) {
	po := ffi.CreateProtocol(sender)
	defer po.Release()
	ffi.CallMethod[ffi.Void](d_, "draggingEnded:", po)
}

func (d_ *DraggingDestinationWrapper) ImplementsDraggingExited() bool {
	return d_.RespondsToSelector(objc.GetSelector("draggingExited:"))
}

func (d_ DraggingDestinationWrapper) DraggingExited(sender DraggingInfo) {
	po := ffi.CreateProtocol(sender)
	defer po.Release()
	ffi.CallMethod[ffi.Void](d_, "draggingExited:", po)
}

func (d_ *DraggingDestinationWrapper) ImplementsPrepareForDragOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("prepareForDragOperation:"))
}

func (d_ DraggingDestinationWrapper) PrepareForDragOperation(sender DraggingInfo) bool {
	po := ffi.CreateProtocol(sender)
	defer po.Release()
	rv := ffi.CallMethod[bool](d_, "prepareForDragOperation:", po)
	return rv
}

func (d_ *DraggingDestinationWrapper) ImplementsPerformDragOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("performDragOperation:"))
}

func (d_ DraggingDestinationWrapper) PerformDragOperation(sender DraggingInfo) bool {
	po := ffi.CreateProtocol(sender)
	defer po.Release()
	rv := ffi.CallMethod[bool](d_, "performDragOperation:", po)
	return rv
}

func (d_ *DraggingDestinationWrapper) ImplementsConcludeDragOperation() bool {
	return d_.RespondsToSelector(objc.GetSelector("concludeDragOperation:"))
}

func (d_ DraggingDestinationWrapper) ConcludeDragOperation(sender DraggingInfo) {
	po := ffi.CreateProtocol(sender)
	defer po.Release()
	ffi.CallMethod[ffi.Void](d_, "concludeDragOperation:", po)
}

func (d_ *DraggingDestinationWrapper) ImplementsUpdateDraggingItemsForDrag() bool {
	return d_.RespondsToSelector(objc.GetSelector("updateDraggingItemsForDrag:"))
}

func (d_ DraggingDestinationWrapper) UpdateDraggingItemsForDrag(sender DraggingInfo) {
	po := ffi.CreateProtocol(sender)
	defer po.Release()
	ffi.CallMethod[ffi.Void](d_, "updateDraggingItemsForDrag:", po)
}

var DraggingItemClass = _DraggingItemClass{objc.GetClass("NSDraggingItem")}

type _DraggingItemClass struct {
	objc.Class
}

type IDraggingItem interface {
	objc.IObject
	SetDraggingFrame_Contents(frame foundation.Rect, contents objc.IObject)
	DraggingFrame() foundation.Rect
	SetDraggingFrame(value foundation.Rect)
	ImageComponents() []DraggingImageComponent
	ImageComponentsProvider() func() []DraggingImageComponent
	SetImageComponentsProvider(value func() []IDraggingImageComponent)
	Item() objc.Object
}

type DraggingItem struct {
	objc.Object
}

func MakeDraggingItem(ptr unsafe.Pointer) DraggingItem {
	return DraggingItem{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ DraggingItem) InitWithPasteboardWriter(pasteboardWriter PasteboardWriting) DraggingItem {
	po := ffi.CreateProtocol(pasteboardWriter)
	defer po.Release()
	rv := ffi.CallMethod[DraggingItem](d_, "initWithPasteboardWriter:", po)
	rv.Autorelease()
	return rv
}

func (dc _DraggingItemClass) Alloc() DraggingItem {
	rv := ffi.CallMethod[DraggingItem](dc, "alloc")
	return rv
}

func (d_ DraggingItem) Init() DraggingItem {
	rv := ffi.CallMethod[DraggingItem](d_, "init")
	rv.Autorelease()
	return rv
}

func (dc _DraggingItemClass) New() DraggingItem {
	rv := ffi.CallMethod[DraggingItem](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDraggingItem() DraggingItem {
	return DraggingItemClass.New()
}

func (d_ DraggingItem) SetDraggingFrame_Contents(frame foundation.Rect, contents objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "setDraggingFrame:contents:", frame, contents)
}

func (d_ DraggingItem) DraggingFrame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](d_, "draggingFrame")
	return rv
}

func (d_ DraggingItem) SetDraggingFrame(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](d_, "setDraggingFrame:", value)
}

func (d_ DraggingItem) ImageComponents() []DraggingImageComponent {
	rv := ffi.CallMethod[[]DraggingImageComponent](d_, "imageComponents")
	return rv
}

func (d_ DraggingItem) ImageComponentsProvider() func() []DraggingImageComponent {
	rv := ffi.CallMethod[func() []DraggingImageComponent](d_, "imageComponentsProvider")
	return rv
}

func (d_ DraggingItem) SetImageComponentsProvider(value func() []IDraggingImageComponent) {
	ffi.CallMethod[ffi.Void](d_, "setImageComponentsProvider:", value)
}

func (d_ DraggingItem) Item() objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "item")
	return rv
}

var DraggingImageComponentClass = _DraggingImageComponentClass{objc.GetClass("NSDraggingImageComponent")}

type _DraggingImageComponentClass struct {
	objc.Class
}

type IDraggingImageComponent interface {
	objc.IObject
	Key() DraggingImageComponentKey
	SetKey(value DraggingImageComponentKey)
	Contents() objc.Object
	SetContents(value objc.IObject)
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
}

type DraggingImageComponent struct {
	objc.Object
}

func MakeDraggingImageComponent(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{
		Object: objc.MakeObject(ptr),
	}
}

func (d_ DraggingImageComponent) InitWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	rv := ffi.CallMethod[DraggingImageComponent](d_, "initWithKey:", key)
	rv.Autorelease()
	return rv
}

func (dc _DraggingImageComponentClass) Alloc() DraggingImageComponent {
	rv := ffi.CallMethod[DraggingImageComponent](dc, "alloc")
	return rv
}

func (d_ DraggingImageComponent) Init() DraggingImageComponent {
	rv := ffi.CallMethod[DraggingImageComponent](d_, "init")
	rv.Autorelease()
	return rv
}

func (dc _DraggingImageComponentClass) New() DraggingImageComponent {
	rv := ffi.CallMethod[DraggingImageComponent](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDraggingImageComponent() DraggingImageComponent {
	return DraggingImageComponentClass.New()
}

func (dc _DraggingImageComponentClass) DraggingImageComponentWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	rv := ffi.CallMethod[DraggingImageComponent](dc, "draggingImageComponentWithKey:", key)
	return rv
}

func (d_ DraggingImageComponent) Key() DraggingImageComponentKey {
	rv := ffi.CallMethod[DraggingImageComponentKey](d_, "key")
	return rv
}

func (d_ DraggingImageComponent) SetKey(value DraggingImageComponentKey) {
	ffi.CallMethod[ffi.Void](d_, "setKey:", value)
}

func (d_ DraggingImageComponent) Contents() objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "contents")
	return rv
}

func (d_ DraggingImageComponent) SetContents(value objc.IObject) {
	ffi.CallMethod[ffi.Void](d_, "setContents:", value)
}

func (d_ DraggingImageComponent) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](d_, "frame")
	return rv
}

func (d_ DraggingImageComponent) SetFrame(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](d_, "setFrame:", value)
}

var DraggingSessionClass = _DraggingSessionClass{objc.GetClass("NSDraggingSession")}

type _DraggingSessionClass struct {
	objc.Class
}

type IDraggingSession interface {
	objc.IObject
	DraggingPasteboard() Pasteboard
	AnimatesToStartingPositionsOnCancelOrFail() bool
	SetAnimatesToStartingPositionsOnCancelOrFail(value bool)
	DraggingFormation() DraggingFormation
	SetDraggingFormation(value DraggingFormation)
	DraggingSequenceNumber() int
	DraggingLocation() foundation.Point
	DraggingLeaderIndex() int
	SetDraggingLeaderIndex(value int)
}

type DraggingSession struct {
	objc.Object
}

func MakeDraggingSession(ptr unsafe.Pointer) DraggingSession {
	return DraggingSession{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DraggingSessionClass) Alloc() DraggingSession {
	rv := ffi.CallMethod[DraggingSession](dc, "alloc")
	return rv
}

func (d_ DraggingSession) Init() DraggingSession {
	rv := ffi.CallMethod[DraggingSession](d_, "init")
	rv.Autorelease()
	return rv
}

func (dc _DraggingSessionClass) New() DraggingSession {
	rv := ffi.CallMethod[DraggingSession](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDraggingSession() DraggingSession {
	return DraggingSessionClass.New()
}

func (d_ DraggingSession) DraggingPasteboard() Pasteboard {
	rv := ffi.CallMethod[Pasteboard](d_, "draggingPasteboard")
	return rv
}

func (d_ DraggingSession) AnimatesToStartingPositionsOnCancelOrFail() bool {
	rv := ffi.CallMethod[bool](d_, "animatesToStartingPositionsOnCancelOrFail")
	return rv
}

func (d_ DraggingSession) SetAnimatesToStartingPositionsOnCancelOrFail(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setAnimatesToStartingPositionsOnCancelOrFail:", value)
}

func (d_ DraggingSession) DraggingFormation() DraggingFormation {
	rv := ffi.CallMethod[DraggingFormation](d_, "draggingFormation")
	return rv
}

func (d_ DraggingSession) SetDraggingFormation(value DraggingFormation) {
	ffi.CallMethod[ffi.Void](d_, "setDraggingFormation:", value)
}

func (d_ DraggingSession) DraggingSequenceNumber() int {
	rv := ffi.CallMethod[int](d_, "draggingSequenceNumber")
	return rv
}

func (d_ DraggingSession) DraggingLocation() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](d_, "draggingLocation")
	return rv
}

func (d_ DraggingSession) DraggingLeaderIndex() int {
	rv := ffi.CallMethod[int](d_, "draggingLeaderIndex")
	return rv
}

func (d_ DraggingSession) SetDraggingLeaderIndex(value int) {
	ffi.CallMethod[ffi.Void](d_, "setDraggingLeaderIndex:", value)
}
