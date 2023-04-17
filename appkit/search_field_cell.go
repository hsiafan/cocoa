// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var SearchFieldCellClass = _SearchFieldCellClass{objc.GetClass("NSSearchFieldCell")}

type _SearchFieldCellClass struct {
	objc.Class
}

type ISearchFieldCell interface {
	ITextFieldCell
	ResetSearchButtonCell()
	ResetCancelButtonCell()
	SearchTextRectForBounds(rect foundation.Rect) foundation.Rect
	SearchButtonRectForBounds(rect foundation.Rect) foundation.Rect
	CancelButtonRectForBounds(rect foundation.Rect) foundation.Rect
	SearchButtonCell() ButtonCell
	SetSearchButtonCell(value IButtonCell)
	CancelButtonCell() ButtonCell
	SetCancelButtonCell(value IButtonCell)
	SearchMenuTemplate() Menu
	SetSearchMenuTemplate(value IMenu)
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	MaximumRecents() int
	SetMaximumRecents(value int)
	RecentSearches() []string
	SetRecentSearches(value []string)
	RecentsAutosaveName() SearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName)
}

type SearchFieldCell struct {
	TextFieldCell
}

func MakeSearchFieldCell(ptr unsafe.Pointer) SearchFieldCell {
	return SearchFieldCell{
		TextFieldCell: MakeTextFieldCell(ptr),
	}
}

func (s_ SearchFieldCell) InitTextCell(string_ string) SearchFieldCell {
	rv := ffi.CallMethod[SearchFieldCell](s_, "initTextCell:", string_)
	return rv
}

func (s_ SearchFieldCell) InitImageCell(image IImage) SearchFieldCell {
	rv := ffi.CallMethod[SearchFieldCell](s_, "initImageCell:", image)
	return rv
}

func (s_ SearchFieldCell) Init() SearchFieldCell {
	rv := ffi.CallMethod[SearchFieldCell](s_, "init")
	return rv
}

func (sc _SearchFieldCellClass) Alloc() SearchFieldCell {
	rv := ffi.CallMethod[SearchFieldCell](sc, "alloc")
	return rv
}

func (sc _SearchFieldCellClass) New() SearchFieldCell {
	rv := ffi.CallMethod[SearchFieldCell](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSearchFieldCell() SearchFieldCell {
	return SearchFieldCellClass.New()
}

func (s_ SearchFieldCell) ResetSearchButtonCell() {
	ffi.CallMethod[ffi.Void](s_, "resetSearchButtonCell")
}

func (s_ SearchFieldCell) ResetCancelButtonCell() {
	ffi.CallMethod[ffi.Void](s_, "resetCancelButtonCell")
}

func (s_ SearchFieldCell) SearchTextRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "searchTextRectForBounds:", rect)
	return rv
}

func (s_ SearchFieldCell) SearchButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "searchButtonRectForBounds:", rect)
	return rv
}

func (s_ SearchFieldCell) CancelButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "cancelButtonRectForBounds:", rect)
	return rv
}

func (s_ SearchFieldCell) SearchButtonCell() ButtonCell {
	rv := ffi.CallMethod[ButtonCell](s_, "searchButtonCell")
	return rv
}

func (s_ SearchFieldCell) SetSearchButtonCell(value IButtonCell) {
	ffi.CallMethod[ffi.Void](s_, "setSearchButtonCell:", value)
}

func (s_ SearchFieldCell) CancelButtonCell() ButtonCell {
	rv := ffi.CallMethod[ButtonCell](s_, "cancelButtonCell")
	return rv
}

func (s_ SearchFieldCell) SetCancelButtonCell(value IButtonCell) {
	ffi.CallMethod[ffi.Void](s_, "setCancelButtonCell:", value)
}

func (s_ SearchFieldCell) SearchMenuTemplate() Menu {
	rv := ffi.CallMethod[Menu](s_, "searchMenuTemplate")
	return rv
}

func (s_ SearchFieldCell) SetSearchMenuTemplate(value IMenu) {
	ffi.CallMethod[ffi.Void](s_, "setSearchMenuTemplate:", value)
}

func (s_ SearchFieldCell) SendsWholeSearchString() bool {
	rv := ffi.CallMethod[bool](s_, "sendsWholeSearchString")
	return rv
}

func (s_ SearchFieldCell) SetSendsWholeSearchString(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setSendsWholeSearchString:", value)
}

func (s_ SearchFieldCell) SendsSearchStringImmediately() bool {
	rv := ffi.CallMethod[bool](s_, "sendsSearchStringImmediately")
	return rv
}

func (s_ SearchFieldCell) SetSendsSearchStringImmediately(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setSendsSearchStringImmediately:", value)
}

func (s_ SearchFieldCell) MaximumRecents() int {
	rv := ffi.CallMethod[int](s_, "maximumRecents")
	return rv
}

func (s_ SearchFieldCell) SetMaximumRecents(value int) {
	ffi.CallMethod[ffi.Void](s_, "setMaximumRecents:", value)
}

func (s_ SearchFieldCell) RecentSearches() []string {
	rv := ffi.CallMethod[[]string](s_, "recentSearches")
	return rv
}

func (s_ SearchFieldCell) SetRecentSearches(value []string) {
	ffi.CallMethod[ffi.Void](s_, "setRecentSearches:", value)
}

func (s_ SearchFieldCell) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	rv := ffi.CallMethod[SearchFieldRecentsAutosaveName](s_, "recentsAutosaveName")
	return rv
}

func (s_ SearchFieldCell) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName) {
	ffi.CallMethod[ffi.Void](s_, "setRecentsAutosaveName:", value)
}
