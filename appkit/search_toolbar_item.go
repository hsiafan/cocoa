// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var SearchToolbarItemClass = _SearchToolbarItemClass{objc.GetClass("NSSearchToolbarItem")}

type _SearchToolbarItemClass struct {
	objc.Class
}

type ISearchToolbarItem interface {
	IToolbarItem
	BeginSearchInteraction()
	EndSearchInteraction()
	PreferredWidthForSearchField() float64
	SetPreferredWidthForSearchField(value float64)
	ResignsFirstResponderWithCancel() bool
	SetResignsFirstResponderWithCancel(value bool)
	SearchField() SearchField
	SetSearchField(value ISearchField)
}

type SearchToolbarItem struct {
	ToolbarItem
}

func MakeSearchToolbarItem(ptr unsafe.Pointer) SearchToolbarItem {
	return SearchToolbarItem{
		ToolbarItem: MakeToolbarItem(ptr),
	}
}

func (s_ SearchToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) SearchToolbarItem {
	rv := objc.CallMethod[SearchToolbarItem](s_, objc.SEL("initWithItemIdentifier:"), itemIdentifier)
	return rv
}

func (sc _SearchToolbarItemClass) Alloc() SearchToolbarItem {
	rv := objc.CallMethod[SearchToolbarItem](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SearchToolbarItemClass) New() SearchToolbarItem {
	rv := objc.CallMethod[SearchToolbarItem](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSearchToolbarItem() SearchToolbarItem {
	return SearchToolbarItemClass.New()
}

func (s_ SearchToolbarItem) Init() SearchToolbarItem {
	rv := objc.CallMethod[SearchToolbarItem](s_, objc.SEL("init"))
	return rv
}

func (s_ SearchToolbarItem) BeginSearchInteraction() {
	objc.CallMethod[objc.Void](s_, objc.SEL("beginSearchInteraction"))
}

func (s_ SearchToolbarItem) EndSearchInteraction() {
	objc.CallMethod[objc.Void](s_, objc.SEL("endSearchInteraction"))
}

func (s_ SearchToolbarItem) PreferredWidthForSearchField() float64 {
	rv := objc.CallMethod[float64](s_, objc.SEL("preferredWidthForSearchField"))
	return rv
}

func (s_ SearchToolbarItem) SetPreferredWidthForSearchField(value float64) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setPreferredWidthForSearchField:"), value)
}

func (s_ SearchToolbarItem) ResignsFirstResponderWithCancel() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("resignsFirstResponderWithCancel"))
	return rv
}

func (s_ SearchToolbarItem) SetResignsFirstResponderWithCancel(value bool) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setResignsFirstResponderWithCancel:"), value)
}

func (s_ SearchToolbarItem) SearchField() SearchField {
	rv := objc.CallMethod[SearchField](s_, objc.SEL("searchField"))
	return rv
}

func (s_ SearchToolbarItem) SetSearchField(value ISearchField) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSearchField:"), objc.ExtractPtr(value))
}
