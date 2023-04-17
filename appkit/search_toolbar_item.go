// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[SearchToolbarItem](s_, "initWithItemIdentifier:", itemIdentifier)
	return rv
}

func (sc _SearchToolbarItemClass) Alloc() SearchToolbarItem {
	rv := ffi.CallMethod[SearchToolbarItem](sc, "alloc")
	return rv
}

func (sc _SearchToolbarItemClass) New() SearchToolbarItem {
	rv := ffi.CallMethod[SearchToolbarItem](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSearchToolbarItem() SearchToolbarItem {
	return SearchToolbarItemClass.New()
}

func (s_ SearchToolbarItem) Init() SearchToolbarItem {
	rv := ffi.CallMethod[SearchToolbarItem](s_, "init")
	return rv
}

func (s_ SearchToolbarItem) BeginSearchInteraction() {
	ffi.CallMethod[ffi.Void](s_, "beginSearchInteraction")
}

func (s_ SearchToolbarItem) EndSearchInteraction() {
	ffi.CallMethod[ffi.Void](s_, "endSearchInteraction")
}

func (s_ SearchToolbarItem) PreferredWidthForSearchField() float64 {
	rv := ffi.CallMethod[float64](s_, "preferredWidthForSearchField")
	return rv
}

func (s_ SearchToolbarItem) SetPreferredWidthForSearchField(value float64) {
	ffi.CallMethod[ffi.Void](s_, "setPreferredWidthForSearchField:", value)
}

func (s_ SearchToolbarItem) ResignsFirstResponderWithCancel() bool {
	rv := ffi.CallMethod[bool](s_, "resignsFirstResponderWithCancel")
	return rv
}

func (s_ SearchToolbarItem) SetResignsFirstResponderWithCancel(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setResignsFirstResponderWithCancel:", value)
}

func (s_ SearchToolbarItem) SearchField() SearchField {
	rv := ffi.CallMethod[SearchField](s_, "searchField")
	return rv
}

func (s_ SearchToolbarItem) SetSearchField(value ISearchField) {
	ffi.CallMethod[ffi.Void](s_, "setSearchField:", value)
}
