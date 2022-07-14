// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var SearchFieldClass = _SearchFieldClass{objc.GetClass("NSSearchField")}

type _SearchFieldClass struct {
	objc.Class
}

type ISearchField interface {
	ITextField
	// deprecated
	RectForCancelButtonWhenCentered(isCentered bool) foundation.Rect
	// deprecated
	RectForSearchButtonWhenCentered(isCentered bool) foundation.Rect
	// deprecated
	RectForSearchTextWhenCentered(isCentered bool) foundation.Rect
	SearchMenuTemplate() Menu
	SetSearchMenuTemplate(value IMenu)
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
	RecentSearches() []string
	SetRecentSearches(value []string)
	MaximumRecents() int
	SetMaximumRecents(value int)
	RecentsAutosaveName() SearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName)
	CancelButtonBounds() foundation.Rect
	SearchButtonBounds() foundation.Rect
	SearchTextBounds() foundation.Rect
	// deprecated
	CentersPlaceholder() bool
	// deprecated
	SetCentersPlaceholder(value bool)
}

type SearchField struct {
	TextField
}

func MakeSearchField(ptr unsafe.Pointer) SearchField {
	return SearchField{
		TextField: MakeTextField(ptr),
	}
}

func (sc _SearchFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) SearchField {
	rv := ffi.CallMethod[SearchField](sc, "labelWithAttributedString:", attributedStringValue)
	return rv
}

func (sc _SearchFieldClass) LabelWithString(stringValue string) SearchField {
	rv := ffi.CallMethod[SearchField](sc, "labelWithString:", stringValue)
	return rv
}

func (sc _SearchFieldClass) TextFieldWithString(stringValue string) SearchField {
	rv := ffi.CallMethod[SearchField](sc, "textFieldWithString:", stringValue)
	return rv
}

func (sc _SearchFieldClass) WrappingLabelWithString(stringValue string) SearchField {
	rv := ffi.CallMethod[SearchField](sc, "wrappingLabelWithString:", stringValue)
	return rv
}

func (s_ SearchField) InitWithFrame(frameRect foundation.Rect) SearchField {
	rv := ffi.CallMethod[SearchField](s_, "initWithFrame:", frameRect)
	rv.Autorelease()
	return rv
}

func (s_ SearchField) Init() SearchField {
	rv := ffi.CallMethod[SearchField](s_, "init")
	rv.Autorelease()
	return rv
}

func (sc _SearchFieldClass) Alloc() SearchField {
	rv := ffi.CallMethod[SearchField](sc, "alloc")
	return rv
}

func (sc _SearchFieldClass) New() SearchField {
	rv := ffi.CallMethod[SearchField](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSearchField() SearchField {
	return SearchFieldClass.New()
}

// deprecated
func (s_ SearchField) RectForCancelButtonWhenCentered(isCentered bool) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "rectForCancelButtonWhenCentered:", isCentered)
	return rv
}

// deprecated
func (s_ SearchField) RectForSearchButtonWhenCentered(isCentered bool) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "rectForSearchButtonWhenCentered:", isCentered)
	return rv
}

// deprecated
func (s_ SearchField) RectForSearchTextWhenCentered(isCentered bool) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "rectForSearchTextWhenCentered:", isCentered)
	return rv
}

func (s_ SearchField) SearchMenuTemplate() Menu {
	rv := ffi.CallMethod[Menu](s_, "searchMenuTemplate")
	return rv
}

func (s_ SearchField) SetSearchMenuTemplate(value IMenu) {
	ffi.CallMethod[ffi.Void](s_, "setSearchMenuTemplate:", value)
}

func (s_ SearchField) SendsSearchStringImmediately() bool {
	rv := ffi.CallMethod[bool](s_, "sendsSearchStringImmediately")
	return rv
}

func (s_ SearchField) SetSendsSearchStringImmediately(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setSendsSearchStringImmediately:", value)
}

func (s_ SearchField) SendsWholeSearchString() bool {
	rv := ffi.CallMethod[bool](s_, "sendsWholeSearchString")
	return rv
}

func (s_ SearchField) SetSendsWholeSearchString(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setSendsWholeSearchString:", value)
}

func (s_ SearchField) RecentSearches() []string {
	rv := ffi.CallMethod[[]string](s_, "recentSearches")
	return rv
}

func (s_ SearchField) SetRecentSearches(value []string) {
	ffi.CallMethod[ffi.Void](s_, "setRecentSearches:", value)
}

func (s_ SearchField) MaximumRecents() int {
	rv := ffi.CallMethod[int](s_, "maximumRecents")
	return rv
}

func (s_ SearchField) SetMaximumRecents(value int) {
	ffi.CallMethod[ffi.Void](s_, "setMaximumRecents:", value)
}

func (s_ SearchField) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	rv := ffi.CallMethod[SearchFieldRecentsAutosaveName](s_, "recentsAutosaveName")
	return rv
}

func (s_ SearchField) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName) {
	ffi.CallMethod[ffi.Void](s_, "setRecentsAutosaveName:", value)
}

func (s_ SearchField) CancelButtonBounds() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "cancelButtonBounds")
	return rv
}

func (s_ SearchField) SearchButtonBounds() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "searchButtonBounds")
	return rv
}

func (s_ SearchField) SearchTextBounds() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "searchTextBounds")
	return rv
}

// deprecated
func (s_ SearchField) CentersPlaceholder() bool {
	rv := ffi.CallMethod[bool](s_, "centersPlaceholder")
	return rv
}

// deprecated
func (s_ SearchField) SetCentersPlaceholder(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setCentersPlaceholder:", value)
}

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

func (s_ SearchFieldCell) InitTextCell(_string string) SearchFieldCell {
	rv := ffi.CallMethod[SearchFieldCell](s_, "initTextCell:", _string)
	rv.Autorelease()
	return rv
}

func (s_ SearchFieldCell) InitImageCell(image IImage) SearchFieldCell {
	rv := ffi.CallMethod[SearchFieldCell](s_, "initImageCell:", image)
	rv.Autorelease()
	return rv
}

func (s_ SearchFieldCell) Init() SearchFieldCell {
	rv := ffi.CallMethod[SearchFieldCell](s_, "init")
	rv.Autorelease()
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

type SearchFieldDelegate interface {
	TextFieldDelegate
	ImplementsSearchFieldDidStartSearching() bool
	// optional
	SearchFieldDidStartSearching(sender SearchField)
	ImplementsSearchFieldDidEndSearching() bool
	// optional
	SearchFieldDidEndSearching(sender SearchField)
}

type SearchFieldDelegateImpl struct {
	TextFieldDelegateImpl
	_SearchFieldDidStartSearching func(sender SearchField)
	_SearchFieldDidEndSearching   func(sender SearchField)
}

func (di *SearchFieldDelegateImpl) ImplementsSearchFieldDidStartSearching() bool {
	return di._SearchFieldDidStartSearching != nil
}

func (di *SearchFieldDelegateImpl) SetSearchFieldDidStartSearching(f func(sender SearchField)) {
	di._SearchFieldDidStartSearching = f
}

func (di *SearchFieldDelegateImpl) SearchFieldDidStartSearching(sender SearchField) {
	di._SearchFieldDidStartSearching(sender)
}
func (di *SearchFieldDelegateImpl) ImplementsSearchFieldDidEndSearching() bool {
	return di._SearchFieldDidEndSearching != nil
}

func (di *SearchFieldDelegateImpl) SetSearchFieldDidEndSearching(f func(sender SearchField)) {
	di._SearchFieldDidEndSearching = f
}

func (di *SearchFieldDelegateImpl) SearchFieldDidEndSearching(sender SearchField) {
	di._SearchFieldDidEndSearching(sender)
}

type SearchFieldDelegateWrapper struct {
	TextFieldDelegateWrapper
}

func (s_ *SearchFieldDelegateWrapper) ImplementsSearchFieldDidStartSearching() bool {
	return s_.RespondsToSelector(objc.GetSelector("searchFieldDidStartSearching:"))
}

func (s_ SearchFieldDelegateWrapper) SearchFieldDidStartSearching(sender ISearchField) {
	ffi.CallMethod[ffi.Void](s_, "searchFieldDidStartSearching:", sender)
}

func (s_ *SearchFieldDelegateWrapper) ImplementsSearchFieldDidEndSearching() bool {
	return s_.RespondsToSelector(objc.GetSelector("searchFieldDidEndSearching:"))
}

func (s_ SearchFieldDelegateWrapper) SearchFieldDidEndSearching(sender ISearchField) {
	ffi.CallMethod[ffi.Void](s_, "searchFieldDidEndSearching:", sender)
}
