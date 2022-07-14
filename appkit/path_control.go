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
	rv.Autorelease()
	return rv
}

func (p_ PathControl) Init() PathControl {
	rv := ffi.CallMethod[PathControl](p_, "init")
	rv.Autorelease()
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
	po := ffi.CreateProtocol(value)
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

type PathControlDelegate interface {
	ImplementsPathControl_ShouldDragPathComponentCell_WithPasteboard() bool
	// optional
	PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	ImplementsPathControl_ValidateDrop() bool
	// optional
	PathControl_ValidateDrop(pathControl PathControl, info DraggingInfoWrapper) DragOperation
	ImplementsPathControl_AcceptDrop() bool
	// optional
	PathControl_AcceptDrop(pathControl PathControl, info DraggingInfoWrapper) bool
	ImplementsPathControl_WillDisplayOpenPanel() bool
	// optional
	PathControl_WillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel)
	ImplementsPathControl_WillPopUpMenu() bool
	// optional
	PathControl_WillPopUpMenu(pathControl PathControl, menu Menu)
	ImplementsPathControl_ShouldDragItem_WithPasteboard() bool
	// optional
	PathControl_ShouldDragItem_WithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
}

type PathControlDelegateImpl struct {
	_PathControl_ShouldDragPathComponentCell_WithPasteboard func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool
	_PathControl_ValidateDrop                               func(pathControl PathControl, info DraggingInfoWrapper) DragOperation
	_PathControl_AcceptDrop                                 func(pathControl PathControl, info DraggingInfoWrapper) bool
	_PathControl_WillDisplayOpenPanel                       func(pathControl PathControl, openPanel OpenPanel)
	_PathControl_WillPopUpMenu                              func(pathControl PathControl, menu Menu)
	_PathControl_ShouldDragItem_WithPasteboard              func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
}

func (di *PathControlDelegateImpl) ImplementsPathControl_ShouldDragPathComponentCell_WithPasteboard() bool {
	return di._PathControl_ShouldDragPathComponentCell_WithPasteboard != nil
}

func (di *PathControlDelegateImpl) SetPathControl_ShouldDragPathComponentCell_WithPasteboard(f func(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool) {
	di._PathControl_ShouldDragPathComponentCell_WithPasteboard = f
}

func (di *PathControlDelegateImpl) PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl PathControl, pathComponentCell PathComponentCell, pasteboard Pasteboard) bool {
	return di._PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl, pathComponentCell, pasteboard)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_ValidateDrop() bool {
	return di._PathControl_ValidateDrop != nil
}

func (di *PathControlDelegateImpl) SetPathControl_ValidateDrop(f func(pathControl PathControl, info DraggingInfoWrapper) DragOperation) {
	di._PathControl_ValidateDrop = f
}

func (di *PathControlDelegateImpl) PathControl_ValidateDrop(pathControl PathControl, info DraggingInfoWrapper) DragOperation {
	return di._PathControl_ValidateDrop(pathControl, info)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_AcceptDrop() bool {
	return di._PathControl_AcceptDrop != nil
}

func (di *PathControlDelegateImpl) SetPathControl_AcceptDrop(f func(pathControl PathControl, info DraggingInfoWrapper) bool) {
	di._PathControl_AcceptDrop = f
}

func (di *PathControlDelegateImpl) PathControl_AcceptDrop(pathControl PathControl, info DraggingInfoWrapper) bool {
	return di._PathControl_AcceptDrop(pathControl, info)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_WillDisplayOpenPanel() bool {
	return di._PathControl_WillDisplayOpenPanel != nil
}

func (di *PathControlDelegateImpl) SetPathControl_WillDisplayOpenPanel(f func(pathControl PathControl, openPanel OpenPanel)) {
	di._PathControl_WillDisplayOpenPanel = f
}

func (di *PathControlDelegateImpl) PathControl_WillDisplayOpenPanel(pathControl PathControl, openPanel OpenPanel) {
	di._PathControl_WillDisplayOpenPanel(pathControl, openPanel)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_WillPopUpMenu() bool {
	return di._PathControl_WillPopUpMenu != nil
}

func (di *PathControlDelegateImpl) SetPathControl_WillPopUpMenu(f func(pathControl PathControl, menu Menu)) {
	di._PathControl_WillPopUpMenu = f
}

func (di *PathControlDelegateImpl) PathControl_WillPopUpMenu(pathControl PathControl, menu Menu) {
	di._PathControl_WillPopUpMenu(pathControl, menu)
}
func (di *PathControlDelegateImpl) ImplementsPathControl_ShouldDragItem_WithPasteboard() bool {
	return di._PathControl_ShouldDragItem_WithPasteboard != nil
}

func (di *PathControlDelegateImpl) SetPathControl_ShouldDragItem_WithPasteboard(f func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool) {
	di._PathControl_ShouldDragItem_WithPasteboard = f
}

func (di *PathControlDelegateImpl) PathControl_ShouldDragItem_WithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool {
	return di._PathControl_ShouldDragItem_WithPasteboard(pathControl, pathItem, pasteboard)
}

type PathControlDelegateWrapper struct {
	objc.Object
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_ShouldDragPathComponentCell_WithPasteboard() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:shouldDragPathComponentCell:withPasteboard:"))
}

func (p_ PathControlDelegateWrapper) PathControl_ShouldDragPathComponentCell_WithPasteboard(pathControl IPathControl, pathComponentCell IPathComponentCell, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](p_, "pathControl:shouldDragPathComponentCell:withPasteboard:", pathControl, pathComponentCell, pasteboard)
	return rv
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_ValidateDrop() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:validateDrop:"))
}

func (p_ PathControlDelegateWrapper) PathControl_ValidateDrop(pathControl IPathControl, info DraggingInfo) DragOperation {
	po := ffi.CreateProtocol(info)
	defer po.Release()
	rv := ffi.CallMethod[DragOperation](p_, "pathControl:validateDrop:", pathControl, po)
	return rv
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_AcceptDrop() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:acceptDrop:"))
}

func (p_ PathControlDelegateWrapper) PathControl_AcceptDrop(pathControl IPathControl, info DraggingInfo) bool {
	po := ffi.CreateProtocol(info)
	defer po.Release()
	rv := ffi.CallMethod[bool](p_, "pathControl:acceptDrop:", pathControl, po)
	return rv
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_WillDisplayOpenPanel() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:willDisplayOpenPanel:"))
}

func (p_ PathControlDelegateWrapper) PathControl_WillDisplayOpenPanel(pathControl IPathControl, openPanel IOpenPanel) {
	ffi.CallMethod[ffi.Void](p_, "pathControl:willDisplayOpenPanel:", pathControl, openPanel)
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_WillPopUpMenu() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:willPopUpMenu:"))
}

func (p_ PathControlDelegateWrapper) PathControl_WillPopUpMenu(pathControl IPathControl, menu IMenu) {
	ffi.CallMethod[ffi.Void](p_, "pathControl:willPopUpMenu:", pathControl, menu)
}

func (p_ *PathControlDelegateWrapper) ImplementsPathControl_ShouldDragItem_WithPasteboard() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathControl:shouldDragItem:withPasteboard:"))
}

func (p_ PathControlDelegateWrapper) PathControl_ShouldDragItem_WithPasteboard(pathControl IPathControl, pathItem IPathControlItem, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](p_, "pathControl:shouldDragItem:withPasteboard:", pathControl, pathItem, pasteboard)
	return rv
}

var PathCellClass = _PathCellClass{objc.GetClass("NSPathCell")}

type _PathCellClass struct {
	objc.Class
}

type IPathCell interface {
	IActionCell
	MouseEntered_WithFrame_InView(event IEvent, frame foundation.Rect, view IView)
	MouseExited_WithFrame_InView(event IEvent, frame foundation.Rect, view IView)
	RectOfPathComponentCell_WithFrame_InView(cell IPathComponentCell, frame foundation.Rect, view IView) foundation.Rect
	PathComponentCellAtPoint_WithFrame_InView(point foundation.Point, frame foundation.Rect, view IView) PathComponentCell
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	ClickedPathComponentCell() PathComponentCell
	PathComponentCells() []PathComponentCell
	SetPathComponentCells(value []IPathComponentCell)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	Delegate() PathCellDelegateWrapper
	SetDelegate(value PathCellDelegate)
}

type PathCell struct {
	ActionCell
}

func MakePathCell(ptr unsafe.Pointer) PathCell {
	return PathCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (p_ PathCell) InitImageCell(image IImage) PathCell {
	rv := ffi.CallMethod[PathCell](p_, "initImageCell:", image)
	rv.Autorelease()
	return rv
}

func (p_ PathCell) InitTextCell(_string string) PathCell {
	rv := ffi.CallMethod[PathCell](p_, "initTextCell:", _string)
	rv.Autorelease()
	return rv
}

func (p_ PathCell) Init() PathCell {
	rv := ffi.CallMethod[PathCell](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PathCellClass) Alloc() PathCell {
	rv := ffi.CallMethod[PathCell](pc, "alloc")
	return rv
}

func (pc _PathCellClass) New() PathCell {
	rv := ffi.CallMethod[PathCell](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPathCell() PathCell {
	return PathCellClass.New()
}

func (p_ PathCell) MouseEntered_WithFrame_InView(event IEvent, frame foundation.Rect, view IView) {
	ffi.CallMethod[ffi.Void](p_, "mouseEntered:withFrame:inView:", event, frame, view)
}

func (p_ PathCell) MouseExited_WithFrame_InView(event IEvent, frame foundation.Rect, view IView) {
	ffi.CallMethod[ffi.Void](p_, "mouseExited:withFrame:inView:", event, frame, view)
}

func (p_ PathCell) RectOfPathComponentCell_WithFrame_InView(cell IPathComponentCell, frame foundation.Rect, view IView) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](p_, "rectOfPathComponentCell:withFrame:inView:", cell, frame, view)
	return rv
}

func (p_ PathCell) PathComponentCellAtPoint_WithFrame_InView(point foundation.Point, frame foundation.Rect, view IView) PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "pathComponentCellAtPoint:withFrame:inView:", point, frame, view)
	return rv
}

func (p_ PathCell) AllowedTypes() []string {
	rv := ffi.CallMethod[[]string](p_, "allowedTypes")
	return rv
}

func (p_ PathCell) SetAllowedTypes(value []string) {
	ffi.CallMethod[ffi.Void](p_, "setAllowedTypes:", value)
}

func (p_ PathCell) PathStyle() PathStyle {
	rv := ffi.CallMethod[PathStyle](p_, "pathStyle")
	return rv
}

func (p_ PathCell) SetPathStyle(value PathStyle) {
	ffi.CallMethod[ffi.Void](p_, "setPathStyle:", value)
}

func (p_ PathCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](p_, "placeholderAttributedString")
	return rv
}

func (p_ PathCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](p_, "setPlaceholderAttributedString:", value)
}

func (p_ PathCell) PlaceholderString() string {
	rv := ffi.CallMethod[string](p_, "placeholderString")
	return rv
}

func (p_ PathCell) SetPlaceholderString(value string) {
	ffi.CallMethod[ffi.Void](p_, "setPlaceholderString:", value)
}

func (p_ PathCell) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](p_, "backgroundColor")
	return rv
}

func (p_ PathCell) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](p_, "setBackgroundColor:", value)
}

func (p_ PathCell) ClickedPathComponentCell() PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "clickedPathComponentCell")
	return rv
}

func (p_ PathCell) PathComponentCells() []PathComponentCell {
	rv := ffi.CallMethod[[]PathComponentCell](p_, "pathComponentCells")
	return rv
}

func (p_ PathCell) SetPathComponentCells(value []IPathComponentCell) {
	ffi.CallMethod[ffi.Void](p_, "setPathComponentCells:", value)
}

func (p_ PathCell) DoubleAction() objc.Selector {
	rv := ffi.CallMethod[objc.Selector](p_, "doubleAction")
	return rv
}

func (p_ PathCell) SetDoubleAction(value objc.Selector) {
	ffi.CallMethod[ffi.Void](p_, "setDoubleAction:", value)
}

func (p_ PathCell) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](p_, "URL")
	return rv
}

func (p_ PathCell) SetURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](p_, "setURL:", value)
}

func (p_ PathCell) Delegate() PathCellDelegateWrapper {
	rv := ffi.CallMethod[PathCellDelegateWrapper](p_, "delegate")
	return rv
}

func (p_ PathCell) SetDelegate(value PathCellDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(p_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](p_, "setDelegate:", po)
}

type PathCellDelegate interface {
	ImplementsPathCell_WillDisplayOpenPanel() bool
	// optional
	PathCell_WillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel)
	ImplementsPathCell_WillPopUpMenu() bool
	// optional
	PathCell_WillPopUpMenu(pathCell PathCell, menu Menu)
}

type PathCellDelegateImpl struct {
	_PathCell_WillDisplayOpenPanel func(pathCell PathCell, openPanel OpenPanel)
	_PathCell_WillPopUpMenu        func(pathCell PathCell, menu Menu)
}

func (di *PathCellDelegateImpl) ImplementsPathCell_WillDisplayOpenPanel() bool {
	return di._PathCell_WillDisplayOpenPanel != nil
}

func (di *PathCellDelegateImpl) SetPathCell_WillDisplayOpenPanel(f func(pathCell PathCell, openPanel OpenPanel)) {
	di._PathCell_WillDisplayOpenPanel = f
}

func (di *PathCellDelegateImpl) PathCell_WillDisplayOpenPanel(pathCell PathCell, openPanel OpenPanel) {
	di._PathCell_WillDisplayOpenPanel(pathCell, openPanel)
}
func (di *PathCellDelegateImpl) ImplementsPathCell_WillPopUpMenu() bool {
	return di._PathCell_WillPopUpMenu != nil
}

func (di *PathCellDelegateImpl) SetPathCell_WillPopUpMenu(f func(pathCell PathCell, menu Menu)) {
	di._PathCell_WillPopUpMenu = f
}

func (di *PathCellDelegateImpl) PathCell_WillPopUpMenu(pathCell PathCell, menu Menu) {
	di._PathCell_WillPopUpMenu(pathCell, menu)
}

type PathCellDelegateWrapper struct {
	objc.Object
}

func (p_ *PathCellDelegateWrapper) ImplementsPathCell_WillDisplayOpenPanel() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathCell:willDisplayOpenPanel:"))
}

func (p_ PathCellDelegateWrapper) PathCell_WillDisplayOpenPanel(pathCell IPathCell, openPanel IOpenPanel) {
	ffi.CallMethod[ffi.Void](p_, "pathCell:willDisplayOpenPanel:", pathCell, openPanel)
}

func (p_ *PathCellDelegateWrapper) ImplementsPathCell_WillPopUpMenu() bool {
	return p_.RespondsToSelector(objc.GetSelector("pathCell:willPopUpMenu:"))
}

func (p_ PathCellDelegateWrapper) PathCell_WillPopUpMenu(pathCell IPathCell, menu IMenu) {
	ffi.CallMethod[ffi.Void](p_, "pathCell:willPopUpMenu:", pathCell, menu)
}

var PathComponentCellClass = _PathComponentCellClass{objc.GetClass("NSPathComponentCell")}

type _PathComponentCellClass struct {
	objc.Class
}

type IPathComponentCell interface {
	ITextFieldCell
	URL() foundation.URL
	SetURL(value foundation.IURL)
}

type PathComponentCell struct {
	TextFieldCell
}

func MakePathComponentCell(ptr unsafe.Pointer) PathComponentCell {
	return PathComponentCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (p_ PathComponentCell) InitTextCell(_string string) PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "initTextCell:", _string)
	rv.Autorelease()
	return rv
}

func (p_ PathComponentCell) InitImageCell(image IImage) PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "initImageCell:", image)
	rv.Autorelease()
	return rv
}

func (p_ PathComponentCell) Init() PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PathComponentCellClass) Alloc() PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](pc, "alloc")
	return rv
}

func (pc _PathComponentCellClass) New() PathComponentCell {
	rv := ffi.CallMethod[PathComponentCell](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPathComponentCell() PathComponentCell {
	return PathComponentCellClass.New()
}

func (p_ PathComponentCell) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](p_, "URL")
	return rv
}

func (p_ PathComponentCell) SetURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](p_, "setURL:", value)
}

var PathControlItemClass = _PathControlItemClass{objc.GetClass("NSPathControlItem")}

type _PathControlItemClass struct {
	objc.Class
}

type IPathControlItem interface {
	objc.IObject
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	Image() Image
	SetImage(value IImage)
	Title() string
	SetTitle(value string)
	URL() foundation.URL
}

type PathControlItem struct {
	objc.Object
}

func MakePathControlItem(ptr unsafe.Pointer) PathControlItem {
	return PathControlItem{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PathControlItemClass) Alloc() PathControlItem {
	rv := ffi.CallMethod[PathControlItem](pc, "alloc")
	return rv
}

func (p_ PathControlItem) Init() PathControlItem {
	rv := ffi.CallMethod[PathControlItem](p_, "init")
	rv.Autorelease()
	return rv
}

func (pc _PathControlItemClass) New() PathControlItem {
	rv := ffi.CallMethod[PathControlItem](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPathControlItem() PathControlItem {
	return PathControlItemClass.New()
}

func (p_ PathControlItem) AttributedTitle() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](p_, "attributedTitle")
	return rv
}

func (p_ PathControlItem) SetAttributedTitle(value foundation.IAttributedString) {
	ffi.CallMethod[ffi.Void](p_, "setAttributedTitle:", value)
}

func (p_ PathControlItem) Image() Image {
	rv := ffi.CallMethod[Image](p_, "image")
	return rv
}

func (p_ PathControlItem) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](p_, "setImage:", value)
}

func (p_ PathControlItem) Title() string {
	rv := ffi.CallMethod[string](p_, "title")
	return rv
}

func (p_ PathControlItem) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](p_, "setTitle:", value)
}

func (p_ PathControlItem) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](p_, "URL")
	return rv
}
