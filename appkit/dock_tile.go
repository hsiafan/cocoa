// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[DockTile](dc, objc.SEL("alloc"))
	return rv
}

func (dc _DockTileClass) New() DockTile {
	rv := objc.CallMethod[DockTile](dc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewDockTile() DockTile {
	return DockTileClass.New()
}

func (d_ DockTile) Init() DockTile {
	rv := objc.CallMethod[DockTile](d_, objc.SEL("init"))
	return rv
}

func (d_ DockTile) Display() {
	objc.CallMethod[objc.Void](d_, objc.SEL("display"))
}

func (d_ DockTile) ContentView() View {
	rv := objc.CallMethod[View](d_, objc.SEL("contentView"))
	return rv
}

func (d_ DockTile) SetContentView(value IView) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setContentView:"), objc.ExtractPtr(value))
}

func (d_ DockTile) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](d_, objc.SEL("size"))
	return rv
}

// weak property
func (d_ DockTile) Owner() objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.SEL("owner"))
	return rv
}

func (d_ DockTile) ShowsApplicationBadge() bool {
	rv := objc.CallMethod[bool](d_, objc.SEL("showsApplicationBadge"))
	return rv
}

func (d_ DockTile) SetShowsApplicationBadge(value bool) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setShowsApplicationBadge:"), value)
}

func (d_ DockTile) BadgeLabel() string {
	rv := objc.CallMethod[string](d_, objc.SEL("badgeLabel"))
	return rv
}

func (d_ DockTile) SetBadgeLabel(value string) {
	objc.CallMethod[objc.Void](d_, objc.SEL("setBadgeLabel:"), value)
}
