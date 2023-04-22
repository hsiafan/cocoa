// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var PathControlClass = _PathControlClass{objc.GetClass("NSPathControl")}

type _PathControlClass struct {
	objc.Class
}

type IPathControl interface {
	IControl
	// deprecated
	ClickedPathComponentCell() PathComponentCell
	// deprecated
	PathComponentCells() []PathComponentCell
	// deprecated
	SetPathComponentCells(cells []IPathComponentCell)
	SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	Delegate() PathControlDelegateWrapper
	SetDelegate(value PathControlDelegate)
	SetDelegate0(value objc.IObject)
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	ClickedPathItem() PathControlItem
	IsEditable() bool
	SetEditable(value bool)
	PathItems() []PathControlItem
	SetPathItems(value []IPathControlItem)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
}

type PathControl struct {
	Control
}

func MakePathControl(ptr unsafe.Pointer) PathControl {
	return PathControl{
		Control: MakeControl(ptr),
	}
}

func (p_ PathControl) InitWithFrame(frameRect foundation.Rect) PathControl {
	rv := objc.CallMethod[PathControl](p_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (p_ PathControl) Init() PathControl {
	rv := objc.CallMethod[PathControl](p_, objc.GetSelector("init"))
	return rv
}

func (pc _PathControlClass) Alloc() PathControl {
	rv := objc.CallMethod[PathControl](pc, objc.GetSelector("alloc"))
	return rv
}

func (pc _PathControlClass) New() PathControl {
	rv := objc.CallMethod[PathControl](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPathControl() PathControl {
	return PathControlClass.New()
}

// deprecated
func (p_ PathControl) ClickedPathComponentCell() PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.GetSelector("clickedPathComponentCell"))
	return rv
}

// deprecated
func (p_ PathControl) PathComponentCells() []PathComponentCell {
	rv := objc.CallMethod[[]PathComponentCell](p_, objc.GetSelector("pathComponentCells"))
	return rv
}

// deprecated
func (p_ PathControl) SetPathComponentCells(cells []IPathComponentCell) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPathComponentCells:"), cells)
}

func (p_ PathControl) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDraggingSourceOperationMask:forLocal:"), mask, isLocal)
}

func (p_ PathControl) PathStyle() PathStyle {
	rv := objc.CallMethod[PathStyle](p_, objc.GetSelector("pathStyle"))
	return rv
}

func (p_ PathControl) SetPathStyle(value PathStyle) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPathStyle:"), value)
}

func (p_ PathControl) BackgroundColor() Color {
	rv := objc.CallMethod[Color](p_, objc.GetSelector("backgroundColor"))
	return rv
}

func (p_ PathControl) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setBackgroundColor:"), value)
}

func (p_ PathControl) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](p_, objc.GetSelector("doubleAction"))
	return rv
}

func (p_ PathControl) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDoubleAction:"), value)
}

func (p_ PathControl) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](p_, objc.GetSelector("URL"))
	return rv
}

func (p_ PathControl) SetURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setURL:"), value)
}

func (p_ PathControl) Delegate() PathControlDelegateWrapper {
	rv := objc.CallMethod[PathControlDelegateWrapper](p_, objc.GetSelector("delegate"))
	return rv
}

func (p_ PathControl) SetDelegate(value PathControlDelegate) {
	po := objc.CreateProtocol("NSPathControlDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(p_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDelegate:"), po)
}

func (p_ PathControl) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDelegate:"), value)
}

func (p_ PathControl) AllowedTypes() []string {
	rv := objc.CallMethod[[]string](p_, objc.GetSelector("allowedTypes"))
	return rv
}

func (p_ PathControl) SetAllowedTypes(value []string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setAllowedTypes:"), value)
}

func (p_ PathControl) ClickedPathItem() PathControlItem {
	rv := objc.CallMethod[PathControlItem](p_, objc.GetSelector("clickedPathItem"))
	return rv
}

func (p_ PathControl) IsEditable() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isEditable"))
	return rv
}

func (p_ PathControl) SetEditable(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setEditable:"), value)
}

func (p_ PathControl) PathItems() []PathControlItem {
	rv := objc.CallMethod[[]PathControlItem](p_, objc.GetSelector("pathItems"))
	return rv
}

func (p_ PathControl) SetPathItems(value []IPathControlItem) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPathItems:"), value)
}

func (p_ PathControl) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](p_, objc.GetSelector("placeholderAttributedString"))
	return rv
}

func (p_ PathControl) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPlaceholderAttributedString:"), value)
}

func (p_ PathControl) PlaceholderString() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("placeholderString"))
	return rv
}

func (p_ PathControl) SetPlaceholderString(value string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPlaceholderString:"), value)
}
