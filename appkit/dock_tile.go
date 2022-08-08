// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var DockTileClass = _DockTileClass{objc.GetClass("NSDockTile")}

type _DockTileClass struct {
	objc.Class
}

type IDockTile interface {
	objc.IObject
	Display()
	ContentView() View
	SetContentView(value IView)
	Size() foundation.Size
	Owner() objc.Object
	ShowsApplicationBadge() bool
	SetShowsApplicationBadge(value bool)
	BadgeLabel() string
	SetBadgeLabel(value string)
}

type DockTile struct {
	objc.Object
}

func MakeDockTile(ptr unsafe.Pointer) DockTile {
	return DockTile{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DockTileClass) Alloc() DockTile {
	rv := ffi.CallMethod[DockTile](dc, "alloc")
	return rv
}

func (d_ DockTile) Init() DockTile {
	rv := ffi.CallMethod[DockTile](d_, "init")
	return rv
}

func (dc _DockTileClass) New() DockTile {
	rv := ffi.CallMethod[DockTile](dc, "new")
	rv.Autorelease()
	return rv
}

func NewDockTile() DockTile {
	return DockTileClass.New()
}

func (d_ DockTile) Display() {
	ffi.CallMethod[ffi.Void](d_, "display")
}

func (d_ DockTile) ContentView() View {
	rv := ffi.CallMethod[View](d_, "contentView")
	return rv
}

func (d_ DockTile) SetContentView(value IView) {
	ffi.CallMethod[ffi.Void](d_, "setContentView:", value)
}

func (d_ DockTile) Size() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](d_, "size")
	return rv
}

func (d_ DockTile) Owner() objc.Object {
	rv := ffi.CallMethod[objc.Object](d_, "owner")
	return rv
}

func (d_ DockTile) ShowsApplicationBadge() bool {
	rv := ffi.CallMethod[bool](d_, "showsApplicationBadge")
	return rv
}

func (d_ DockTile) SetShowsApplicationBadge(value bool) {
	ffi.CallMethod[ffi.Void](d_, "setShowsApplicationBadge:", value)
}

func (d_ DockTile) BadgeLabel() string {
	rv := ffi.CallMethod[string](d_, "badgeLabel")
	return rv
}

func (d_ DockTile) SetBadgeLabel(value string) {
	ffi.CallMethod[ffi.Void](d_, "setBadgeLabel:", value)
}
