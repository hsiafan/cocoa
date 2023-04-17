// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[PathControl](p_, "initWithFrame:", frameRect)
	return rv
}

func (p_ PathControl) Init() PathControl {
	rv := ffi.CallMethod[PathControl](p_, "init")
	return rv
}

func (pc _PathControlClass) Alloc() PathControl {
	rv := ffi.CallMethod[PathControl](pc, "alloc")
	return rv
}

func (pc _PathControlClass) New() PathControl {
	rv := ffi.CallMethod[PathControl](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPathControl() PathControl {
	return PathControlClass.New()
}

// deprecated
func (p_ PathControl) ClickedPathComponentCell() PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "clickedPathComponentCell")
	return rv
}

// deprecated
func (p_ PathControl) PathComponentCells() []PathComponentCell {
	rv := ffi.CallMethod[[]PathComponentCell](p_, "pathComponentCells")
	return rv
}

// deprecated
func (p_ PathControl) SetPathComponentCells(cells []IPathComponentCell) {
	ffi.CallMethod[ffi.Void](p_, "setPathComponentCells:", cells)
}

func (p_ PathControl) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	ffi.CallMethod[ffi.Void](p_, "setDraggingSourceOperationMask:forLocal:", mask, isLocal)
}

func (p_ PathControl) PathStyle() PathStyle {
	rv := ffi.CallMethod[PathStyle](p_, "pathStyle")
	return rv
}

func (p_ PathControl) SetPathStyle(value PathStyle) {
	ffi.CallMethod[ffi.Void](p_, "setPathStyle:", value)
}

func (p_ PathControl) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](p_, "backgroundColor")
	return rv
}

func (p_ PathControl) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](p_, "setBackgroundColor:", value)
}

func (p_ PathControl) DoubleAction() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](p_, "doubleAction")
	return rv
}

func (p_ PathControl) SetDoubleAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](p_, "setDoubleAction:", value)
}

func (p_ PathControl) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](p_, "URL")
	return rv
}

func (p_ PathControl) SetURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](p_, "setURL:", value)
}

func (p_ PathControl) Delegate() PathControlDelegateWrapper {
	rv := ffi.CallMethod[PathControlDelegateWrapper](p_, "delegate")
	return rv
}

func (p_ PathControl) SetDelegate(value PathControlDelegate) {
	po := ffi.CreateProtocol("NSPathControlDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(p_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](p_, "setDelegate:", po)
}

func (p_ PathControl) AllowedTypes() []string {
	rv := ffi.CallMethod[[]string](p_, "allowedTypes")
	return rv
}

func (p_ PathControl) SetAllowedTypes(value []string) {
	ffi.CallMethod[ffi.Void](p_, "setAllowedTypes:", value)
}

func (p_ PathControl) ClickedPathItem() PathControlItem {
	rv := ffi.CallMethod[PathControlItem](p_, "clickedPathItem")
	return rv
}

func (p_ PathControl) IsEditable() bool {
	rv := ffi.CallMethod[bool](p_, "isEditable")
	return rv
}

func (p_ PathControl) SetEditable(value bool) {
	ffi.CallMethod[ffi.Void](p_, "setEditable:", value)
}

func (p_ PathControl) PathItems() []PathControlItem {
	rv := ffi.CallMethod[[]PathControlItem](p_, "pathItems")
	return rv
}

func (p_ PathControl) SetPathItems(value []IPathControlItem) {
	ffi.CallMethod[ffi.Void](p_, "setPathItems:", value)
}

func (p_ PathControl) PlaceholderAttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](p_, "placeholderAttributedString")
	return rv
}

func (p_ PathControl) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](p_, "setPlaceholderAttributedString:", value)
}

func (p_ PathControl) PlaceholderString() string {
	rv := ffi.CallMethod[string](p_, "placeholderString")
	return rv
}

func (p_ PathControl) SetPlaceholderString(value string) {
	ffi.CallMethod[ffi.Void](p_, "setPlaceholderString:", value)
}
