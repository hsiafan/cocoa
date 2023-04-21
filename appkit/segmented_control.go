// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var SegmentedControlClass = _SegmentedControlClass{objc.GetClass("NSSegmentedControl")}

type _SegmentedControlClass struct {
	objc.Class
}

type ISegmentedControl interface {
	IControl
	LabelForSegment(segment int) string
	SetLabel_ForSegment(label string, segment int)
	SetAlignment_ForSegment(alignment TextAlignment, segment int)
	AlignmentForSegment(segment int) TextAlignment
	SetImage_ForSegment(image IImage, segment int)
	ImageForSegment(segment int) Image
	SetImageScaling_ForSegment(scaling ImageScaling, segment int)
	ImageScalingForSegment(segment int) ImageScaling
	SetMenu_ForSegment(menu IMenu, segment int)
	MenuForSegment(segment int) Menu
	SetShowsMenuIndicator_ForSegment(showsMenuIndicator bool, segment int)
	ShowsMenuIndicatorForSegment(segment int) bool
	SelectSegmentWithTag(tag int) bool
	SetSelected_ForSegment(selected bool, segment int)
	IsSelectedForSegment(segment int) bool
	SetWidth_ForSegment(width float64, segment int)
	WidthForSegment(segment int) float64
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size
	SetEnabled_ForSegment(enabled bool, segment int)
	IsEnabledForSegment(segment int) bool
	TagForSegment(segment int) int
	SetTag_ForSegment(tag int, segment int)
	SetToolTip_ForSegment(toolTip string, segment int)
	ToolTipForSegment(segment int) string
	TrackingMode() SegmentSwitchTracking
	SetTrackingMode(value SegmentSwitchTracking)
	SegmentStyle() SegmentStyle
	SetSegmentStyle(value SegmentStyle)
	SegmentCount() int
	SetSegmentCount(value int)
	IsSpringLoaded() bool
	SetSpringLoaded(value bool)
	SelectedSegment() int
	SetSelectedSegment(value int)
	IndexOfSelectedItem() int
	SelectedSegmentBezelColor() Color
	SetSelectedSegmentBezelColor(value IColor)
	DoubleValueForSelectedSegment() float64
	SegmentDistribution() SegmentDistribution
	SetSegmentDistribution(value SegmentDistribution)
	ActiveCompressionOptions() UserInterfaceCompressionOptions
}

type SegmentedControl struct {
	Control
}

func MakeSegmentedControl(ptr unsafe.Pointer) SegmentedControl {
	return SegmentedControl{
		Control: MakeControl(ptr),
	}
}

func (sc _SegmentedControlClass) SegmentedControlWithImages_TrackingMode_Target_Action(images []IImage, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](sc, "segmentedControlWithImages:trackingMode:target:action:", images, trackingMode, target, action)
	return rv
}

func (sc _SegmentedControlClass) SegmentedControlWithLabels_TrackingMode_Target_Action(labels []string, trackingMode SegmentSwitchTracking, target objc.IObject, action objc.Selector) SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](sc, "segmentedControlWithLabels:trackingMode:target:action:", labels, trackingMode, target, action)
	return rv
}

func (s_ SegmentedControl) InitWithFrame(frameRect foundation.Rect) SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](s_, "initWithFrame:", frameRect)
	return rv
}

func (s_ SegmentedControl) Init() SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](s_, "init")
	return rv
}

func (sc _SegmentedControlClass) Alloc() SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](sc, "alloc")
	return rv
}

func (sc _SegmentedControlClass) New() SegmentedControl {
	rv := objc.CallMethod[SegmentedControl](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSegmentedControl() SegmentedControl {
	return SegmentedControlClass.New()
}

func (s_ SegmentedControl) LabelForSegment(segment int) string {
	rv := objc.CallMethod[string](s_, "labelForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) SetLabel_ForSegment(label string, segment int) {
	objc.CallMethod[objc.Void](s_, "setLabel:forSegment:", label, segment)
}

func (s_ SegmentedControl) SetAlignment_ForSegment(alignment TextAlignment, segment int) {
	objc.CallMethod[objc.Void](s_, "setAlignment:forSegment:", alignment, segment)
}

func (s_ SegmentedControl) AlignmentForSegment(segment int) TextAlignment {
	rv := objc.CallMethod[TextAlignment](s_, "alignmentForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) SetImage_ForSegment(image IImage, segment int) {
	objc.CallMethod[objc.Void](s_, "setImage:forSegment:", image, segment)
}

func (s_ SegmentedControl) ImageForSegment(segment int) Image {
	rv := objc.CallMethod[Image](s_, "imageForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) SetImageScaling_ForSegment(scaling ImageScaling, segment int) {
	objc.CallMethod[objc.Void](s_, "setImageScaling:forSegment:", scaling, segment)
}

func (s_ SegmentedControl) ImageScalingForSegment(segment int) ImageScaling {
	rv := objc.CallMethod[ImageScaling](s_, "imageScalingForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) SetMenu_ForSegment(menu IMenu, segment int) {
	objc.CallMethod[objc.Void](s_, "setMenu:forSegment:", menu, segment)
}

func (s_ SegmentedControl) MenuForSegment(segment int) Menu {
	rv := objc.CallMethod[Menu](s_, "menuForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) SetShowsMenuIndicator_ForSegment(showsMenuIndicator bool, segment int) {
	objc.CallMethod[objc.Void](s_, "setShowsMenuIndicator:forSegment:", showsMenuIndicator, segment)
}

func (s_ SegmentedControl) ShowsMenuIndicatorForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, "showsMenuIndicatorForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) SelectSegmentWithTag(tag int) bool {
	rv := objc.CallMethod[bool](s_, "selectSegmentWithTag:", tag)
	return rv
}

func (s_ SegmentedControl) SetSelected_ForSegment(selected bool, segment int) {
	objc.CallMethod[objc.Void](s_, "setSelected:forSegment:", selected, segment)
}

func (s_ SegmentedControl) IsSelectedForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, "isSelectedForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) SetWidth_ForSegment(width float64, segment int) {
	objc.CallMethod[objc.Void](s_, "setWidth:forSegment:", width, segment)
}

func (s_ SegmentedControl) WidthForSegment(segment int) float64 {
	rv := objc.CallMethod[float64](s_, "widthForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) CompressWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) {
	objc.CallMethod[objc.Void](s_, "compressWithPrioritizedCompressionOptions:", prioritizedOptions)
}

func (s_ SegmentedControl) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []IUserInterfaceCompressionOptions) foundation.Size {
	rv := objc.CallMethod[foundation.Size](s_, "minimumSizeWithPrioritizedCompressionOptions:", prioritizedOptions)
	return rv
}

func (s_ SegmentedControl) SetEnabled_ForSegment(enabled bool, segment int) {
	objc.CallMethod[objc.Void](s_, "setEnabled:forSegment:", enabled, segment)
}

func (s_ SegmentedControl) IsEnabledForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, "isEnabledForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) TagForSegment(segment int) int {
	rv := objc.CallMethod[int](s_, "tagForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) SetTag_ForSegment(tag int, segment int) {
	objc.CallMethod[objc.Void](s_, "setTag:forSegment:", tag, segment)
}

func (s_ SegmentedControl) SetToolTip_ForSegment(toolTip string, segment int) {
	objc.CallMethod[objc.Void](s_, "setToolTip:forSegment:", toolTip, segment)
}

func (s_ SegmentedControl) ToolTipForSegment(segment int) string {
	rv := objc.CallMethod[string](s_, "toolTipForSegment:", segment)
	return rv
}

func (s_ SegmentedControl) TrackingMode() SegmentSwitchTracking {
	rv := objc.CallMethod[SegmentSwitchTracking](s_, "trackingMode")
	return rv
}

func (s_ SegmentedControl) SetTrackingMode(value SegmentSwitchTracking) {
	objc.CallMethod[objc.Void](s_, "setTrackingMode:", value)
}

func (s_ SegmentedControl) SegmentStyle() SegmentStyle {
	rv := objc.CallMethod[SegmentStyle](s_, "segmentStyle")
	return rv
}

func (s_ SegmentedControl) SetSegmentStyle(value SegmentStyle) {
	objc.CallMethod[objc.Void](s_, "setSegmentStyle:", value)
}

func (s_ SegmentedControl) SegmentCount() int {
	rv := objc.CallMethod[int](s_, "segmentCount")
	return rv
}

func (s_ SegmentedControl) SetSegmentCount(value int) {
	objc.CallMethod[objc.Void](s_, "setSegmentCount:", value)
}

func (s_ SegmentedControl) IsSpringLoaded() bool {
	rv := objc.CallMethod[bool](s_, "isSpringLoaded")
	return rv
}

func (s_ SegmentedControl) SetSpringLoaded(value bool) {
	objc.CallMethod[objc.Void](s_, "setSpringLoaded:", value)
}

func (s_ SegmentedControl) SelectedSegment() int {
	rv := objc.CallMethod[int](s_, "selectedSegment")
	return rv
}

func (s_ SegmentedControl) SetSelectedSegment(value int) {
	objc.CallMethod[objc.Void](s_, "setSelectedSegment:", value)
}

func (s_ SegmentedControl) IndexOfSelectedItem() int {
	rv := objc.CallMethod[int](s_, "indexOfSelectedItem")
	return rv
}

func (s_ SegmentedControl) SelectedSegmentBezelColor() Color {
	rv := objc.CallMethod[Color](s_, "selectedSegmentBezelColor")
	return rv
}

func (s_ SegmentedControl) SetSelectedSegmentBezelColor(value IColor) {
	objc.CallMethod[objc.Void](s_, "setSelectedSegmentBezelColor:", value)
}

func (s_ SegmentedControl) DoubleValueForSelectedSegment() float64 {
	rv := objc.CallMethod[float64](s_, "doubleValueForSelectedSegment")
	return rv
}

func (s_ SegmentedControl) SegmentDistribution() SegmentDistribution {
	rv := objc.CallMethod[SegmentDistribution](s_, "segmentDistribution")
	return rv
}

func (s_ SegmentedControl) SetSegmentDistribution(value SegmentDistribution) {
	objc.CallMethod[objc.Void](s_, "setSegmentDistribution:", value)
}

func (s_ SegmentedControl) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	rv := objc.CallMethod[UserInterfaceCompressionOptions](s_, "activeCompressionOptions")
	return rv
}
