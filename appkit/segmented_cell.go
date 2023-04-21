// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var SegmentedCellClass = _SegmentedCellClass{objc.GetClass("NSSegmentedCell")}

type _SegmentedCellClass struct {
	objc.Class
}

type ISegmentedCell interface {
	IActionCell
	SetSelected_ForSegment(selected bool, segment int)
	SelectSegmentWithTag(tag int) bool
	MakeNextSegmentKey()
	MakePreviousSegmentKey()
	IsSelectedForSegment(segment int) bool
	SetLabel_ForSegment(label string, segment int)
	LabelForSegment(segment int) string
	SetImage_ForSegment(image IImage, segment int)
	ImageForSegment(segment int) Image
	SetImageScaling_ForSegment(scaling ImageScaling, segment int)
	ImageScalingForSegment(segment int) ImageScaling
	SetWidth_ForSegment(width float64, segment int)
	WidthForSegment(segment int) float64
	SetEnabled_ForSegment(enabled bool, segment int)
	IsEnabledForSegment(segment int) bool
	SetMenu_ForSegment(menu IMenu, segment int)
	MenuForSegment(segment int) Menu
	SetToolTip_ForSegment(toolTip string, segment int)
	ToolTipForSegment(segment int) string
	SetTag_ForSegment(tag int, segment int)
	TagForSegment(segment int) int
	DrawSegment_InFrame_WithView(segment int, frame foundation.Rect, controlView IView)
	InteriorBackgroundStyleForSegment(segment int) BackgroundStyle
	SegmentCount() int
	SetSegmentCount(value int)
	SelectedSegment() int
	SetSelectedSegment(value int)
	TrackingMode() SegmentSwitchTracking
	SetTrackingMode(value SegmentSwitchTracking)
	SegmentStyle() SegmentStyle
	SetSegmentStyle(value SegmentStyle)
}

type SegmentedCell struct {
	ActionCell
}

func MakeSegmentedCell(ptr unsafe.Pointer) SegmentedCell {
	return SegmentedCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (s_ SegmentedCell) InitImageCell(image IImage) SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](s_, "initImageCell:", image)
	return rv
}

func (s_ SegmentedCell) InitTextCell(string_ string) SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](s_, "initTextCell:", string_)
	return rv
}

func (s_ SegmentedCell) Init() SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](s_, "init")
	return rv
}

func (sc _SegmentedCellClass) Alloc() SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](sc, "alloc")
	return rv
}

func (sc _SegmentedCellClass) New() SegmentedCell {
	rv := objc.CallMethod[SegmentedCell](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSegmentedCell() SegmentedCell {
	return SegmentedCellClass.New()
}

func (s_ SegmentedCell) SetSelected_ForSegment(selected bool, segment int) {
	objc.CallMethod[objc.Void](s_, "setSelected:forSegment:", selected, segment)
}

func (s_ SegmentedCell) SelectSegmentWithTag(tag int) bool {
	rv := objc.CallMethod[bool](s_, "selectSegmentWithTag:", tag)
	return rv
}

func (s_ SegmentedCell) MakeNextSegmentKey() {
	objc.CallMethod[objc.Void](s_, "makeNextSegmentKey")
}

func (s_ SegmentedCell) MakePreviousSegmentKey() {
	objc.CallMethod[objc.Void](s_, "makePreviousSegmentKey")
}

func (s_ SegmentedCell) IsSelectedForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, "isSelectedForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) SetLabel_ForSegment(label string, segment int) {
	objc.CallMethod[objc.Void](s_, "setLabel:forSegment:", label, segment)
}

func (s_ SegmentedCell) LabelForSegment(segment int) string {
	rv := objc.CallMethod[string](s_, "labelForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) SetImage_ForSegment(image IImage, segment int) {
	objc.CallMethod[objc.Void](s_, "setImage:forSegment:", image, segment)
}

func (s_ SegmentedCell) ImageForSegment(segment int) Image {
	rv := objc.CallMethod[Image](s_, "imageForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) SetImageScaling_ForSegment(scaling ImageScaling, segment int) {
	objc.CallMethod[objc.Void](s_, "setImageScaling:forSegment:", scaling, segment)
}

func (s_ SegmentedCell) ImageScalingForSegment(segment int) ImageScaling {
	rv := objc.CallMethod[ImageScaling](s_, "imageScalingForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) SetWidth_ForSegment(width float64, segment int) {
	objc.CallMethod[objc.Void](s_, "setWidth:forSegment:", width, segment)
}

func (s_ SegmentedCell) WidthForSegment(segment int) float64 {
	rv := objc.CallMethod[float64](s_, "widthForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) SetEnabled_ForSegment(enabled bool, segment int) {
	objc.CallMethod[objc.Void](s_, "setEnabled:forSegment:", enabled, segment)
}

func (s_ SegmentedCell) IsEnabledForSegment(segment int) bool {
	rv := objc.CallMethod[bool](s_, "isEnabledForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) SetMenu_ForSegment(menu IMenu, segment int) {
	objc.CallMethod[objc.Void](s_, "setMenu:forSegment:", menu, segment)
}

func (s_ SegmentedCell) MenuForSegment(segment int) Menu {
	rv := objc.CallMethod[Menu](s_, "menuForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) SetToolTip_ForSegment(toolTip string, segment int) {
	objc.CallMethod[objc.Void](s_, "setToolTip:forSegment:", toolTip, segment)
}

func (s_ SegmentedCell) ToolTipForSegment(segment int) string {
	rv := objc.CallMethod[string](s_, "toolTipForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) SetTag_ForSegment(tag int, segment int) {
	objc.CallMethod[objc.Void](s_, "setTag:forSegment:", tag, segment)
}

func (s_ SegmentedCell) TagForSegment(segment int) int {
	rv := objc.CallMethod[int](s_, "tagForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) DrawSegment_InFrame_WithView(segment int, frame foundation.Rect, controlView IView) {
	objc.CallMethod[objc.Void](s_, "drawSegment:inFrame:withView:", segment, frame, controlView)
}

func (s_ SegmentedCell) InteriorBackgroundStyleForSegment(segment int) BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](s_, "interiorBackgroundStyleForSegment:", segment)
	return rv
}

func (s_ SegmentedCell) SegmentCount() int {
	rv := objc.CallMethod[int](s_, "segmentCount")
	return rv
}

func (s_ SegmentedCell) SetSegmentCount(value int) {
	objc.CallMethod[objc.Void](s_, "setSegmentCount:", value)
}

func (s_ SegmentedCell) SelectedSegment() int {
	rv := objc.CallMethod[int](s_, "selectedSegment")
	return rv
}

func (s_ SegmentedCell) SetSelectedSegment(value int) {
	objc.CallMethod[objc.Void](s_, "setSelectedSegment:", value)
}

func (s_ SegmentedCell) TrackingMode() SegmentSwitchTracking {
	rv := objc.CallMethod[SegmentSwitchTracking](s_, "trackingMode")
	return rv
}

func (s_ SegmentedCell) SetTrackingMode(value SegmentSwitchTracking) {
	objc.CallMethod[objc.Void](s_, "setTrackingMode:", value)
}

func (s_ SegmentedCell) SegmentStyle() SegmentStyle {
	rv := objc.CallMethod[SegmentStyle](s_, "segmentStyle")
	return rv
}

func (s_ SegmentedCell) SetSegmentStyle(value SegmentStyle) {
	objc.CallMethod[objc.Void](s_, "setSegmentStyle:", value)
}
