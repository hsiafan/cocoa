// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[SearchFieldCell](s_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (s_ SearchFieldCell) InitImageCell(image IImage) SearchFieldCell {
	rv := objc.CallMethod[SearchFieldCell](s_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (s_ SearchFieldCell) Init() SearchFieldCell {
	rv := objc.CallMethod[SearchFieldCell](s_, objc.SEL("init"))
	return rv
}

func (sc _SearchFieldCellClass) Alloc() SearchFieldCell {
	rv := objc.CallMethod[SearchFieldCell](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SearchFieldCellClass) New() SearchFieldCell {
	rv := objc.CallMethod[SearchFieldCell](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSearchFieldCell() SearchFieldCell {
	return SearchFieldCellClass.New()
}

func (s_ SearchFieldCell) ResetSearchButtonCell() {
	objc.CallMethod[objc.Void](s_, objc.SEL("resetSearchButtonCell"))
}

func (s_ SearchFieldCell) ResetCancelButtonCell() {
	objc.CallMethod[objc.Void](s_, objc.SEL("resetCancelButtonCell"))
}

func (s_ SearchFieldCell) SearchTextRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("searchTextRectForBounds:"), rect)
	return rv
}

func (s_ SearchFieldCell) SearchButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("searchButtonRectForBounds:"), rect)
	return rv
}

func (s_ SearchFieldCell) CancelButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("cancelButtonRectForBounds:"), rect)
	return rv
}

func (s_ SearchFieldCell) SearchButtonCell() ButtonCell {
	rv := objc.CallMethod[ButtonCell](s_, objc.SEL("searchButtonCell"))
	return rv
}

func (s_ SearchFieldCell) SetSearchButtonCell(value IButtonCell) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSearchButtonCell:"), objc.ExtractPtr(value))
}

func (s_ SearchFieldCell) CancelButtonCell() ButtonCell {
	rv := objc.CallMethod[ButtonCell](s_, objc.SEL("cancelButtonCell"))
	return rv
}

func (s_ SearchFieldCell) SetCancelButtonCell(value IButtonCell) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setCancelButtonCell:"), objc.ExtractPtr(value))
}

func (s_ SearchFieldCell) SearchMenuTemplate() Menu {
	rv := objc.CallMethod[Menu](s_, objc.SEL("searchMenuTemplate"))
	return rv
}

func (s_ SearchFieldCell) SetSearchMenuTemplate(value IMenu) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSearchMenuTemplate:"), objc.ExtractPtr(value))
}

func (s_ SearchFieldCell) SendsWholeSearchString() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("sendsWholeSearchString"))
	return rv
}

func (s_ SearchFieldCell) SetSendsWholeSearchString(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSendsWholeSearchString:"), value)
}

func (s_ SearchFieldCell) SendsSearchStringImmediately() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("sendsSearchStringImmediately"))
	return rv
}

func (s_ SearchFieldCell) SetSendsSearchStringImmediately(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSendsSearchStringImmediately:"), value)
}

func (s_ SearchFieldCell) MaximumRecents() int {
	rv := objc.CallMethod[int](s_, objc.SEL("maximumRecents"))
	return rv
}

func (s_ SearchFieldCell) SetMaximumRecents(value int) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMaximumRecents:"), value)
}

func (s_ SearchFieldCell) RecentSearches() []string {
	rv := objc.CallMethod[[]string](s_, objc.SEL("recentSearches"))
	return rv
}

func (s_ SearchFieldCell) SetRecentSearches(value []string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setRecentSearches:"), value)
}

func (s_ SearchFieldCell) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	rv := objc.CallMethod[SearchFieldRecentsAutosaveName](s_, objc.SEL("recentsAutosaveName"))
	return rv
}

func (s_ SearchFieldCell) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setRecentsAutosaveName:"), value)
}
