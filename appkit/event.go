// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
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
	rv := objc.CallMethod[Event](ec, objc.SEL("alloc"))
	return rv
}

func (ec _EventClass) New() Event {
	rv := objc.CallMethod[Event](ec, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewEvent() Event {
	return EventClass.New()
}

func (e_ Event) Init() Event {
	rv := objc.CallMethod[Event](e_, objc.SEL("init"))
	return rv
}

func (ec _EventClass) KeyEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_Characters_CharactersIgnoringModifiers_IsARepeat_KeyCode(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, keys string, ukeys string, flag bool, code uint16) Event {
	rv := objc.CallMethod[Event](ec, objc.SEL("keyEventWithType:location:modifierFlags:timestamp:windowNumber:context:characters:charactersIgnoringModifiers:isARepeat:keyCode:"), type_, location, flags, time, wNum, objc.ExtractPtr(unusedPassNil), keys, ukeys, flag, code)
	return rv
}

func (ec _EventClass) MouseEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_ClickCount_Pressure(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, eNum int, cNum int, pressure float32) Event {
	rv := objc.CallMethod[Event](ec, objc.SEL("mouseEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:clickCount:pressure:"), type_, location, flags, time, wNum, objc.ExtractPtr(unusedPassNil), eNum, cNum, pressure)
	return rv
}

func (ec _EventClass) EnterExitEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_TrackingNumber_UserData(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, eNum int, tNum int, data unsafe.Pointer) Event {
	rv := objc.CallMethod[Event](ec, objc.SEL("enterExitEventWithType:location:modifierFlags:timestamp:windowNumber:context:eventNumber:trackingNumber:userData:"), type_, location, flags, time, wNum, objc.ExtractPtr(unusedPassNil), eNum, tNum, data)
	return rv
}

func (ec _EventClass) OtherEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_Subtype_Data1_Data2(type_ EventType, location foundation.Point, flags EventModifierFlags, time foundation.TimeInterval, wNum int, unusedPassNil IGraphicsContext, subtype int16, d1 int, d2 int) Event {
	rv := objc.CallMethod[Event](ec, objc.SEL("otherEventWithType:location:modifierFlags:timestamp:windowNumber:context:subtype:data1:data2:"), type_, location, flags, time, wNum, objc.ExtractPtr(unusedPassNil), subtype, d1, d2)
	return rv
}

func (ec _EventClass) EventWithEventRef(eventRef unsafe.Pointer) Event {
	rv := objc.CallMethod[Event](ec, objc.SEL("eventWithEventRef:"), eventRef)
	return rv
}

func (ec _EventClass) EventWithCGEvent(cgEvent coregraphics.EventRef) Event {
	rv := objc.CallMethod[Event](ec, objc.SEL("eventWithCGEvent:"), cgEvent)
	return rv
}

func (e_ Event) CharactersByApplyingModifiers(modifiers EventModifierFlags) string {
	rv := objc.CallMethod[string](e_, objc.SEL("charactersByApplyingModifiers:"), modifiers)
	return rv
}

func (e_ Event) TrackSwipeEventWithOptions_DampenAmountThresholdMin_Max_UsingHandler(options EventSwipeTrackingOptions, minDampenThreshold float64, maxDampenThreshold float64, trackingHandler func(gestureAmount float64, phase EventPhase, isComplete bool, stop *bool)) {
	objc.CallMethod[objc.Void](e_, objc.SEL("trackSwipeEventWithOptions:dampenAmountThresholdMin:max:usingHandler:"), options, minDampenThreshold, maxDampenThreshold, trackingHandler)
}

func (e_ Event) TouchesMatchingPhase_InView(phase TouchPhase, view IView) foundation.Set {
	rv := objc.CallMethod[foundation.Set](e_, objc.SEL("touchesMatchingPhase:inView:"), phase, objc.ExtractPtr(view))
	return rv
}

func (e_ Event) AllTouches() foundation.Set {
	rv := objc.CallMethod[foundation.Set](e_, objc.SEL("allTouches"))
	return rv
}

func (e_ Event) TouchesForView(view IView) foundation.Set {
	rv := objc.CallMethod[foundation.Set](e_, objc.SEL("touchesForView:"), objc.ExtractPtr(view))
	return rv
}

func (e_ Event) CoalescedTouchesForTouch(touch ITouch) []Touch {
	rv := objc.CallMethod[[]Touch](e_, objc.SEL("coalescedTouchesForTouch:"), objc.ExtractPtr(touch))
	return rv
}

func (ec _EventClass) StartPeriodicEventsAfterDelay_WithPeriod(delay foundation.TimeInterval, period foundation.TimeInterval) {
	objc.CallMethod[objc.Void](ec, objc.SEL("startPeriodicEventsAfterDelay:withPeriod:"), delay, period)
}

func (ec _EventClass) StopPeriodicEvents() {
	objc.CallMethod[objc.Void](ec, objc.SEL("stopPeriodicEvents"))
}

func (ec _EventClass) AddGlobalMonitorForEventsMatchingMask_Handler(mask EventMask, block func(event Event)) objc.Object {
	rv := objc.CallMethod[objc.Object](ec, objc.SEL("addGlobalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}

func (ec _EventClass) AddLocalMonitorForEventsMatchingMask_Handler(mask EventMask, block func(event Event) IEvent) objc.Object {
	rv := objc.CallMethod[objc.Object](ec, objc.SEL("addLocalMonitorForEventsMatchingMask:handler:"), mask, block)
	return rv
}

func (ec _EventClass) RemoveMonitor(eventMonitor objc.IObject) {
	objc.CallMethod[objc.Void](ec, objc.SEL("removeMonitor:"), objc.ExtractPtr(eventMonitor))
}

func (e_ Event) Type() EventType {
	rv := objc.CallMethod[EventType](e_, objc.SEL("type"))
	return rv
}

func (e_ Event) Subtype() EventSubtype {
	rv := objc.CallMethod[EventSubtype](e_, objc.SEL("subtype"))
	return rv
}

func (e_ Event) LocationInWindow() foundation.Point {
	rv := objc.CallMethod[foundation.Point](e_, objc.SEL("locationInWindow"))
	return rv
}

func (e_ Event) Timestamp() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](e_, objc.SEL("timestamp"))
	return rv
}

// weak property
func (e_ Event) Window() Window {
	rv := objc.CallMethod[Window](e_, objc.SEL("window"))
	return rv
}

func (e_ Event) WindowNumber() int {
	rv := objc.CallMethod[int](e_, objc.SEL("windowNumber"))
	return rv
}

func (e_ Event) EventRef() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](e_, objc.SEL("eventRef"))
	return rv
}

func (e_ Event) CGEvent() coregraphics.EventRef {
	rv := objc.CallMethod[coregraphics.EventRef](e_, objc.SEL("CGEvent"))
	return rv
}

func (e_ Event) ModifierFlags() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](e_, objc.SEL("modifierFlags"))
	return rv
}

func (e_ Event) Characters() string {
	rv := objc.CallMethod[string](e_, objc.SEL("characters"))
	return rv
}

func (e_ Event) CharactersIgnoringModifiers() string {
	rv := objc.CallMethod[string](e_, objc.SEL("charactersIgnoringModifiers"))
	return rv
}

func (e_ Event) KeyCode() uint16 {
	rv := objc.CallMethod[uint16](e_, objc.SEL("keyCode"))
	return rv
}

func (ec _EventClass) KeyRepeatDelay() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](ec, objc.SEL("keyRepeatDelay"))
	return rv
}

func (ec _EventClass) KeyRepeatInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](ec, objc.SEL("keyRepeatInterval"))
	return rv
}

func (e_ Event) IsARepeat() bool {
	rv := objc.CallMethod[bool](e_, objc.SEL("isARepeat"))
	return rv
}

func (ec _EventClass) PressedMouseButtons() uint {
	rv := objc.CallMethod[uint](ec, objc.SEL("pressedMouseButtons"))
	return rv
}

func (ec _EventClass) DoubleClickInterval() foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](ec, objc.SEL("doubleClickInterval"))
	return rv
}

func (ec _EventClass) MouseLocation() foundation.Point {
	rv := objc.CallMethod[foundation.Point](ec, objc.SEL("mouseLocation"))
	return rv
}

func (e_ Event) ButtonNumber() int {
	rv := objc.CallMethod[int](e_, objc.SEL("buttonNumber"))
	return rv
}

func (e_ Event) ClickCount() int {
	rv := objc.CallMethod[int](e_, objc.SEL("clickCount"))
	return rv
}

func (e_ Event) AssociatedEventsMask() EventMask {
	rv := objc.CallMethod[EventMask](e_, objc.SEL("associatedEventsMask"))
	return rv
}

func (e_ Event) DeltaX() float64 {
	rv := objc.CallMethod[float64](e_, objc.SEL("deltaX"))
	return rv
}

func (e_ Event) DeltaY() float64 {
	rv := objc.CallMethod[float64](e_, objc.SEL("deltaY"))
	return rv
}

func (e_ Event) DeltaZ() float64 {
	rv := objc.CallMethod[float64](e_, objc.SEL("deltaZ"))
	return rv
}

func (e_ Event) HasPreciseScrollingDeltas() bool {
	rv := objc.CallMethod[bool](e_, objc.SEL("hasPreciseScrollingDeltas"))
	return rv
}

func (e_ Event) ScrollingDeltaX() float64 {
	rv := objc.CallMethod[float64](e_, objc.SEL("scrollingDeltaX"))
	return rv
}

func (e_ Event) ScrollingDeltaY() float64 {
	rv := objc.CallMethod[float64](e_, objc.SEL("scrollingDeltaY"))
	return rv
}

func (e_ Event) MomentumPhase() EventPhase {
	rv := objc.CallMethod[EventPhase](e_, objc.SEL("momentumPhase"))
	return rv
}

func (e_ Event) IsDirectionInvertedFromDevice() bool {
	rv := objc.CallMethod[bool](e_, objc.SEL("isDirectionInvertedFromDevice"))
	return rv
}

func (ec _EventClass) IsSwipeTrackingFromScrollEventsEnabled() bool {
	rv := objc.CallMethod[bool](ec, objc.SEL("isSwipeTrackingFromScrollEventsEnabled"))
	return rv
}

func (e_ Event) Phase() EventPhase {
	rv := objc.CallMethod[EventPhase](e_, objc.SEL("phase"))
	return rv
}

func (e_ Event) Magnification() float64 {
	rv := objc.CallMethod[float64](e_, objc.SEL("magnification"))
	return rv
}

func (ec _EventClass) IsMouseCoalescingEnabled() bool {
	rv := objc.CallMethod[bool](ec, objc.SEL("isMouseCoalescingEnabled"))
	return rv
}

func (ec _EventClass) SetMouseCoalescingEnabled(value bool) {
	objc.CallMethod[objc.Void](ec, objc.SEL("setMouseCoalescingEnabled:"), value)
}

func (e_ Event) Pressure() float32 {
	rv := objc.CallMethod[float32](e_, objc.SEL("pressure"))
	return rv
}

func (e_ Event) Stage() int {
	rv := objc.CallMethod[int](e_, objc.SEL("stage"))
	return rv
}

func (e_ Event) StageTransition() float64 {
	rv := objc.CallMethod[float64](e_, objc.SEL("stageTransition"))
	return rv
}

func (e_ Event) PressureBehavior() PressureBehavior {
	rv := objc.CallMethod[PressureBehavior](e_, objc.SEL("pressureBehavior"))
	return rv
}

func (e_ Event) CapabilityMask() uint {
	rv := objc.CallMethod[uint](e_, objc.SEL("capabilityMask"))
	return rv
}

func (e_ Event) DeviceID() uint {
	rv := objc.CallMethod[uint](e_, objc.SEL("deviceID"))
	return rv
}

func (e_ Event) IsEnteringProximity() bool {
	rv := objc.CallMethod[bool](e_, objc.SEL("isEnteringProximity"))
	return rv
}

func (e_ Event) PointingDeviceID() uint {
	rv := objc.CallMethod[uint](e_, objc.SEL("pointingDeviceID"))
	return rv
}

func (e_ Event) PointingDeviceSerialNumber() uint {
	rv := objc.CallMethod[uint](e_, objc.SEL("pointingDeviceSerialNumber"))
	return rv
}

func (e_ Event) PointingDeviceType() PointingDeviceType {
	rv := objc.CallMethod[PointingDeviceType](e_, objc.SEL("pointingDeviceType"))
	return rv
}

func (e_ Event) SystemTabletID() uint {
	rv := objc.CallMethod[uint](e_, objc.SEL("systemTabletID"))
	return rv
}

func (e_ Event) TabletID() uint {
	rv := objc.CallMethod[uint](e_, objc.SEL("tabletID"))
	return rv
}

func (e_ Event) UniqueID() uint64 {
	rv := objc.CallMethod[uint64](e_, objc.SEL("uniqueID"))
	return rv
}

func (e_ Event) VendorID() uint {
	rv := objc.CallMethod[uint](e_, objc.SEL("vendorID"))
	return rv
}

func (e_ Event) VendorPointingDeviceType() uint {
	rv := objc.CallMethod[uint](e_, objc.SEL("vendorPointingDeviceType"))
	return rv
}

func (e_ Event) AbsoluteX() int {
	rv := objc.CallMethod[int](e_, objc.SEL("absoluteX"))
	return rv
}

func (e_ Event) AbsoluteY() int {
	rv := objc.CallMethod[int](e_, objc.SEL("absoluteY"))
	return rv
}

func (e_ Event) AbsoluteZ() int {
	rv := objc.CallMethod[int](e_, objc.SEL("absoluteZ"))
	return rv
}

func (e_ Event) ButtonMask() EventButtonMask {
	rv := objc.CallMethod[EventButtonMask](e_, objc.SEL("buttonMask"))
	return rv
}

func (e_ Event) Rotation() float32 {
	rv := objc.CallMethod[float32](e_, objc.SEL("rotation"))
	return rv
}

func (e_ Event) TangentialPressure() float32 {
	rv := objc.CallMethod[float32](e_, objc.SEL("tangentialPressure"))
	return rv
}

func (e_ Event) Tilt() foundation.Point {
	rv := objc.CallMethod[foundation.Point](e_, objc.SEL("tilt"))
	return rv
}

func (e_ Event) VendorDefined() objc.Object {
	rv := objc.CallMethod[objc.Object](e_, objc.SEL("vendorDefined"))
	return rv
}

func (e_ Event) EventNumber() int {
	rv := objc.CallMethod[int](e_, objc.SEL("eventNumber"))
	return rv
}

func (e_ Event) TrackingNumber() int {
	rv := objc.CallMethod[int](e_, objc.SEL("trackingNumber"))
	return rv
}

func (e_ Event) TrackingArea() TrackingArea {
	rv := objc.CallMethod[TrackingArea](e_, objc.SEL("trackingArea"))
	return rv
}

func (e_ Event) UserData() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](e_, objc.SEL("userData"))
	return rv
}

func (e_ Event) Data1() int {
	rv := objc.CallMethod[int](e_, objc.SEL("data1"))
	return rv
}

func (e_ Event) Data2() int {
	rv := objc.CallMethod[int](e_, objc.SEL("data2"))
	return rv
}

// deprecated
func (e_ Event) Context() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](e_, objc.SEL("context"))
	return rv
}
