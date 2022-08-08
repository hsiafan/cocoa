// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var EventClass = _EventClass{objc.GetClass("NSEvent")}

type _EventClass struct {
	objc.Class
}

type IEvent interface {
	objc.IObject
	CharactersByApplyingModifiers(modifiers EventModifierFlags) string
	TrackSwipeEventWithOptions_DampenAmountThresholdMin_Max_UsingHandler(options EventSwipeTrackingOptions, minDampenThreshold float64, maxDampenThreshold float64, trackingHandler func(gestureAmount float64, phase EventPhase, isComplete bool, stop *bool))
	TouchesMatchingPhase_InView(phase TouchPhase, view IView) foundation.Set
	AllTouches() foundation.Set
	TouchesForView(view IView) foundation.Set
	CoalescedTouchesForTouch(touch ITouch) []Touch
	Type() EventType
	Subtype() EventSubtype
	LocationInWindow() foundation.Point
	Timestamp() foundation.TimeInterval
	Window() Window
	WindowNumber() int
	EventRef() unsafe.Pointer
	CGEvent() coregraphics.EventRef
	ModifierFlags() EventModifierFlags
	Characters() string
	CharactersIgnoringModifiers() string
	KeyCode() uint16
	IsARepeat() bool
	ButtonNumber() int
	ClickCount() int
	AssociatedEventsMask() EventMask
	DeltaX() float64
	DeltaY() float64
	DeltaZ() float64
	HasPreciseScrollingDeltas() bool
	ScrollingDeltaX() float64
	ScrollingDeltaY() float64
	MomentumPhase() EventPhase
	IsDirectionInvertedFromDevice() bool
	Phase() EventPhase
	Magnification() float64
	Pressure() float32
	Stage() int
	StageTransition() float64
	PressureBehavior() PressureBehavior
	CapabilityMask() uint
	DeviceID() uint
	IsEnteringProximity() bool
	PointingDeviceID() uint
	PointingDeviceSerialNumber() uint
	PointingDeviceType() PointingDeviceType
	SystemTabletID() uint
	TabletID() uint
	UniqueID() uint64
	VendorID() uint
	VendorPointingDeviceType() uint
	AbsoluteX() int
	AbsoluteY() int
	AbsoluteZ() int
	ButtonMask() EventButtonMask
	Rotation() float32
	TangentialPressure() float32
	Tilt() foundation.Point
	VendorDefined() objc.Object
	EventNumber() int
	TrackingNumber() int
	TrackingArea() TrackingArea
	UserData() unsafe.Pointer
	Data1() int
	Data2() int
	// deprecated
	Context() GraphicsContext
}

type Event struct {
	objc.Object
}

func MakeEvent(ptr unsafe.Pointer) Event {
	return Event{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _EventClass) Alloc() Event {
	rv := ffi.CallMethod[Event](ec, "alloc")
	return rv
}

func (e_ Event) Init() Event {
	rv := ffi.CallMethod[Event](e_, "init")
	return rv
}

func (ec _EventClass) New() Event {
	rv := ffi.CallMethod[Event](ec, "new")
	rv.Autorelease()
	return rv
}

func NewEvent() Event {
	return EventClass.New()
}

func (ec _EventClass) KeyEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_Characters_CharactersIgnoringModifiers_IsARepeat_KeyCode(_type EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, keys string, ukeys string, flag bool, code uint16) Event {
	rv := ffi.CallMethod[Event](ec, "keyEventWithType:location:modifierFlags:timestamp:windowNumber:context:characters:charactersIgnoringModifiers:isARepeat:keyCode:", _type, location, flags, time, wNum, unusedPassNil, keys, ukeys, flag, code)
	return rv
}

func (ec _EventClass) MouseEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_ClickCount_Pressure(_type EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, eNum int, cNum int, pressure float32) Event {
	rv := ffi.CallMethod[Event](ec, "mouseEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:clickCount:pressure:", _type, location, flags, time, wNum, unusedPassNil, eNum, cNum, pressure)
	return rv
}

func (ec _EventClass) EnterExitEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_TrackingNumber_UserData(_type EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, eNum int, tNum int, data unsafe.Pointer) Event {
	rv := ffi.CallMethod[Event](ec, "enterExitEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:trackingNumber:userData:", _type, location, flags, time, wNum, unusedPassNil, eNum, tNum, data)
	return rv
}

func (ec _EventClass) OtherEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_Subtype_Data1_Data2(_type EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, subtype int16, d1 int, d2 int) Event {
	rv := ffi.CallMethod[Event](ec, "otherEventWithType:location:modifierFlags:timestamp:windowNumber:context:subtype:data1:data2:", _type, location, flags, time, wNum, unusedPassNil, subtype, d1, d2)
	return rv
}

func (ec _EventClass) EventWithEventRef(eventRef unsafe.Pointer) Event {
	rv := ffi.CallMethod[Event](ec, "eventWithEventRef:", eventRef)
	return rv
}

func (ec _EventClass) EventWithCGEvent(cgEvent coregraphics.EventRef) Event {
	rv := ffi.CallMethod[Event](ec, "eventWithCGEvent:", cgEvent)
	return rv
}

func (e_ Event) CharactersByApplyingModifiers(modifiers EventModifierFlags) string {
	rv := ffi.CallMethod[string](e_, "charactersByApplyingModifiers:", modifiers)
	return rv
}

func (e_ Event) TrackSwipeEventWithOptions_DampenAmountThresholdMin_Max_UsingHandler(options EventSwipeTrackingOptions, minDampenThreshold float64, maxDampenThreshold float64, trackingHandler func(gestureAmount float64, phase EventPhase, isComplete bool, stop *bool)) {
	ffi.CallMethod[ffi.Void](e_, "trackSwipeEventWithOptions:dampenAmountThresholdMin:max:usingHandler:", options, minDampenThreshold, maxDampenThreshold, trackingHandler)
}

func (e_ Event) TouchesMatchingPhase_InView(phase TouchPhase, view IView) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](e_, "touchesMatchingPhase:inView:", phase, view)
	return rv
}

func (e_ Event) AllTouches() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](e_, "allTouches")
	return rv
}

func (e_ Event) TouchesForView(view IView) foundation.Set {
	rv := ffi.CallMethod[foundation.Set](e_, "touchesForView:", view)
	return rv
}

func (e_ Event) CoalescedTouchesForTouch(touch ITouch) []Touch {
	rv := ffi.CallMethod[[]Touch](e_, "coalescedTouchesForTouch:", touch)
	return rv
}

func (ec _EventClass) StartPeriodicEventsAfterDelay_WithPeriod(delay foundation.TimeInterval, period foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](ec, "startPeriodicEventsAfterDelay:withPeriod:", delay, period)
}

func (ec _EventClass) StopPeriodicEvents() {
	ffi.CallMethod[ffi.Void](ec, "stopPeriodicEvents")
}

func (ec _EventClass) AddGlobalMonitorForEventsMatchingMask_Handler(mask EventMask, block func(event Event)) objc.Object {
	rv := ffi.CallMethod[objc.Object](ec, "addGlobalMonitorForEventsMatchingMask:handler:", mask, block)
	return rv
}

func (ec _EventClass) AddLocalMonitorForEventsMatchingMask_Handler(mask EventMask, block func(event Event) IEvent) objc.Object {
	rv := ffi.CallMethod[objc.Object](ec, "addLocalMonitorForEventsMatchingMask:handler:", mask, block)
	return rv
}

func (ec _EventClass) RemoveMonitor(eventMonitor objc.IObject) {
	ffi.CallMethod[ffi.Void](ec, "removeMonitor:", eventMonitor)
}

func (e_ Event) Type() EventType {
	rv := ffi.CallMethod[EventType](e_, "type")
	return rv
}

func (e_ Event) Subtype() EventSubtype {
	rv := ffi.CallMethod[EventSubtype](e_, "subtype")
	return rv
}

func (e_ Event) LocationInWindow() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](e_, "locationInWindow")
	return rv
}

func (e_ Event) Timestamp() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](e_, "timestamp")
	return rv
}

func (e_ Event) Window() Window {
	rv := ffi.CallMethod[Window](e_, "window")
	return rv
}

func (e_ Event) WindowNumber() int {
	rv := ffi.CallMethod[int](e_, "windowNumber")
	return rv
}

func (e_ Event) EventRef() unsafe.Pointer {
	rv := ffi.CallMethod[unsafe.Pointer](e_, "eventRef")
	return rv
}

func (e_ Event) CGEvent() coregraphics.EventRef {
	rv := ffi.CallMethod[coregraphics.EventRef](e_, "CGEvent")
	return rv
}

func (e_ Event) ModifierFlags() EventModifierFlags {
	rv := ffi.CallMethod[EventModifierFlags](e_, "modifierFlags")
	return rv
}

func (e_ Event) Characters() string {
	rv := ffi.CallMethod[string](e_, "characters")
	return rv
}

func (e_ Event) CharactersIgnoringModifiers() string {
	rv := ffi.CallMethod[string](e_, "charactersIgnoringModifiers")
	return rv
}

func (e_ Event) KeyCode() uint16 {
	rv := ffi.CallMethod[uint16](e_, "keyCode")
	return rv
}

func (ec _EventClass) KeyRepeatDelay() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](ec, "keyRepeatDelay")
	return rv
}

func (ec _EventClass) KeyRepeatInterval() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](ec, "keyRepeatInterval")
	return rv
}

func (e_ Event) IsARepeat() bool {
	rv := ffi.CallMethod[bool](e_, "isARepeat")
	return rv
}

func (ec _EventClass) PressedMouseButtons() uint {
	rv := ffi.CallMethod[uint](ec, "pressedMouseButtons")
	return rv
}

func (ec _EventClass) DoubleClickInterval() foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](ec, "doubleClickInterval")
	return rv
}

func (ec _EventClass) MouseLocation() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](ec, "mouseLocation")
	return rv
}

func (e_ Event) ButtonNumber() int {
	rv := ffi.CallMethod[int](e_, "buttonNumber")
	return rv
}

func (e_ Event) ClickCount() int {
	rv := ffi.CallMethod[int](e_, "clickCount")
	return rv
}

func (e_ Event) AssociatedEventsMask() EventMask {
	rv := ffi.CallMethod[EventMask](e_, "associatedEventsMask")
	return rv
}

func (e_ Event) DeltaX() float64 {
	rv := ffi.CallMethod[float64](e_, "deltaX")
	return rv
}

func (e_ Event) DeltaY() float64 {
	rv := ffi.CallMethod[float64](e_, "deltaY")
	return rv
}

func (e_ Event) DeltaZ() float64 {
	rv := ffi.CallMethod[float64](e_, "deltaZ")
	return rv
}

func (e_ Event) HasPreciseScrollingDeltas() bool {
	rv := ffi.CallMethod[bool](e_, "hasPreciseScrollingDeltas")
	return rv
}

func (e_ Event) ScrollingDeltaX() float64 {
	rv := ffi.CallMethod[float64](e_, "scrollingDeltaX")
	return rv
}

func (e_ Event) ScrollingDeltaY() float64 {
	rv := ffi.CallMethod[float64](e_, "scrollingDeltaY")
	return rv
}

func (e_ Event) MomentumPhase() EventPhase {
	rv := ffi.CallMethod[EventPhase](e_, "momentumPhase")
	return rv
}

func (e_ Event) IsDirectionInvertedFromDevice() bool {
	rv := ffi.CallMethod[bool](e_, "isDirectionInvertedFromDevice")
	return rv
}

func (ec _EventClass) IsSwipeTrackingFromScrollEventsEnabled() bool {
	rv := ffi.CallMethod[bool](ec, "isSwipeTrackingFromScrollEventsEnabled")
	return rv
}

func (e_ Event) Phase() EventPhase {
	rv := ffi.CallMethod[EventPhase](e_, "phase")
	return rv
}

func (e_ Event) Magnification() float64 {
	rv := ffi.CallMethod[float64](e_, "magnification")
	return rv
}

func (ec _EventClass) IsMouseCoalescingEnabled() bool {
	rv := ffi.CallMethod[bool](ec, "isMouseCoalescingEnabled")
	return rv
}

func (ec _EventClass) SetMouseCoalescingEnabled(value bool) {
	ffi.CallMethod[ffi.Void](ec, "setMouseCoalescingEnabled:", value)
}

func (e_ Event) Pressure() float32 {
	rv := ffi.CallMethod[float32](e_, "pressure")
	return rv
}

func (e_ Event) Stage() int {
	rv := ffi.CallMethod[int](e_, "stage")
	return rv
}

func (e_ Event) StageTransition() float64 {
	rv := ffi.CallMethod[float64](e_, "stageTransition")
	return rv
}

func (e_ Event) PressureBehavior() PressureBehavior {
	rv := ffi.CallMethod[PressureBehavior](e_, "pressureBehavior")
	return rv
}

func (e_ Event) CapabilityMask() uint {
	rv := ffi.CallMethod[uint](e_, "capabilityMask")
	return rv
}

func (e_ Event) DeviceID() uint {
	rv := ffi.CallMethod[uint](e_, "deviceID")
	return rv
}

func (e_ Event) IsEnteringProximity() bool {
	rv := ffi.CallMethod[bool](e_, "isEnteringProximity")
	return rv
}

func (e_ Event) PointingDeviceID() uint {
	rv := ffi.CallMethod[uint](e_, "pointingDeviceID")
	return rv
}

func (e_ Event) PointingDeviceSerialNumber() uint {
	rv := ffi.CallMethod[uint](e_, "pointingDeviceSerialNumber")
	return rv
}

func (e_ Event) PointingDeviceType() PointingDeviceType {
	rv := ffi.CallMethod[PointingDeviceType](e_, "pointingDeviceType")
	return rv
}

func (e_ Event) SystemTabletID() uint {
	rv := ffi.CallMethod[uint](e_, "systemTabletID")
	return rv
}

func (e_ Event) TabletID() uint {
	rv := ffi.CallMethod[uint](e_, "tabletID")
	return rv
}

func (e_ Event) UniqueID() uint64 {
	rv := ffi.CallMethod[uint64](e_, "uniqueID")
	return rv
}

func (e_ Event) VendorID() uint {
	rv := ffi.CallMethod[uint](e_, "vendorID")
	return rv
}

func (e_ Event) VendorPointingDeviceType() uint {
	rv := ffi.CallMethod[uint](e_, "vendorPointingDeviceType")
	return rv
}

func (e_ Event) AbsoluteX() int {
	rv := ffi.CallMethod[int](e_, "absoluteX")
	return rv
}

func (e_ Event) AbsoluteY() int {
	rv := ffi.CallMethod[int](e_, "absoluteY")
	return rv
}

func (e_ Event) AbsoluteZ() int {
	rv := ffi.CallMethod[int](e_, "absoluteZ")
	return rv
}

func (e_ Event) ButtonMask() EventButtonMask {
	rv := ffi.CallMethod[EventButtonMask](e_, "buttonMask")
	return rv
}

func (e_ Event) Rotation() float32 {
	rv := ffi.CallMethod[float32](e_, "rotation")
	return rv
}

func (e_ Event) TangentialPressure() float32 {
	rv := ffi.CallMethod[float32](e_, "tangentialPressure")
	return rv
}

func (e_ Event) Tilt() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](e_, "tilt")
	return rv
}

func (e_ Event) VendorDefined() objc.Object {
	rv := ffi.CallMethod[objc.Object](e_, "vendorDefined")
	return rv
}

func (e_ Event) EventNumber() int {
	rv := ffi.CallMethod[int](e_, "eventNumber")
	return rv
}

func (e_ Event) TrackingNumber() int {
	rv := ffi.CallMethod[int](e_, "trackingNumber")
	return rv
}

func (e_ Event) TrackingArea() TrackingArea {
	rv := ffi.CallMethod[TrackingArea](e_, "trackingArea")
	return rv
}

func (e_ Event) UserData() unsafe.Pointer {
	rv := ffi.CallMethod[unsafe.Pointer](e_, "userData")
	return rv
}

func (e_ Event) Data1() int {
	rv := ffi.CallMethod[int](e_, "data1")
	return rv
}

func (e_ Event) Data2() int {
	rv := ffi.CallMethod[int](e_, "data2")
	return rv
}

// deprecated
func (e_ Event) Context() GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](e_, "context")
	return rv
}
