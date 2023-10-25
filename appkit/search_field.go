// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[SearchField](sc, objc.SEL("labelWithAttributedString:"), objc.ExtractPtr(attributedStringValue))
	return rv
}

func (sc _SearchFieldClass) LabelWithString(stringValue string) SearchField {
	rv := objc.CallMethod[SearchField](sc, objc.SEL("labelWithString:"), stringValue)
	return rv
}

func (sc _SearchFieldClass) TextFieldWithString(stringValue string) SearchField {
	rv := objc.CallMethod[SearchField](sc, objc.SEL("textFieldWithString:"), stringValue)
	return rv
}

func (sc _SearchFieldClass) WrappingLabelWithString(stringValue string) SearchField {
	rv := objc.CallMethod[SearchField](sc, objc.SEL("wrappingLabelWithString:"), stringValue)
	return rv
}

func (s_ SearchField) InitWithFrame(frameRect foundation.Rect) SearchField {
	rv := objc.CallMethod[SearchField](s_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (s_ SearchField) Init() SearchField {
	rv := objc.CallMethod[SearchField](s_, objc.SEL("init"))
	return rv
}

func (sc _SearchFieldClass) Alloc() SearchField {
	rv := objc.CallMethod[SearchField](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SearchFieldClass) New() SearchField {
	rv := objc.CallMethod[SearchField](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSearchField() SearchField {
	return SearchFieldClass.New()
}

// deprecated
func (s_ SearchField) RectForCancelButtonWhenCentered(isCentered bool) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("rectForCancelButtonWhenCentered:"), isCentered)
	return rv
}

// deprecated
func (s_ SearchField) RectForSearchButtonWhenCentered(isCentered bool) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("rectForSearchButtonWhenCentered:"), isCentered)
	return rv
}

// deprecated
func (s_ SearchField) RectForSearchTextWhenCentered(isCentered bool) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("rectForSearchTextWhenCentered:"), isCentered)
	return rv
}

func (s_ SearchField) SearchMenuTemplate() Menu {
	rv := objc.CallMethod[Menu](s_, objc.SEL("searchMenuTemplate"))
	return rv
}

func (s_ SearchField) SetSearchMenuTemplate(value IMenu) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSearchMenuTemplate:"), objc.ExtractPtr(value))
}

func (s_ SearchField) SendsSearchStringImmediately() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("sendsSearchStringImmediately"))
	return rv
}

func (s_ SearchField) SetSendsSearchStringImmediately(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSendsSearchStringImmediately:"), value)
}

func (s_ SearchField) SendsWholeSearchString() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("sendsWholeSearchString"))
	return rv
}

func (s_ SearchField) SetSendsWholeSearchString(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSendsWholeSearchString:"), value)
}

func (s_ SearchField) RecentSearches() []string {
	rv := objc.CallMethod[[]string](s_, objc.SEL("recentSearches"))
	return rv
}

func (s_ SearchField) SetRecentSearches(value []string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setRecentSearches:"), value)
}

func (s_ SearchField) MaximumRecents() int {
	rv := objc.CallMethod[int](s_, objc.SEL("maximumRecents"))
	return rv
}

func (s_ SearchField) SetMaximumRecents(value int) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMaximumRecents:"), value)
}

func (s_ SearchField) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	rv := objc.CallMethod[SearchFieldRecentsAutosaveName](s_, objc.SEL("recentsAutosaveName"))
	return rv
}

func (s_ SearchField) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setRecentsAutosaveName:"), value)
}

func (s_ SearchField) CancelButtonBounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("cancelButtonBounds"))
	return rv
}

func (s_ SearchField) SearchButtonBounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("searchButtonBounds"))
	return rv
}

func (s_ SearchField) SearchTextBounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.SEL("searchTextBounds"))
	return rv
}

// deprecated
func (s_ SearchField) CentersPlaceholder() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("centersPlaceholder"))
	return rv
}

// deprecated
func (s_ SearchField) SetCentersPlaceholder(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setCentersPlaceholder:"), value)
}
