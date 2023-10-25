// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var BrowserCellClass = _BrowserCellClass{objc.GetClass("NSBrowserCell")}

type _BrowserCellClass struct {
	objc.Class
}

type IBrowserCell interface {
	ICell
	Reset()
	Set()
	HighlightColorInView(controlView IView) Color
	AlternateImage() Image
	SetAlternateImage(value IImage)
	IsLeaf() bool
	SetLeaf(value bool)
	IsLoaded() bool
	SetLoaded(value bool)
}

type BrowserCell struct {
	Cell
}

func MakeBrowserCell(ptr unsafe.Pointer) BrowserCell {
	return BrowserCell{
		Cell: MakeCell(ptr),
	}
}

func (b_ BrowserCell) InitImageCell(image IImage) BrowserCell {
	rv := objc.CallMethod[BrowserCell](b_, objc.SEL("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func (b_ BrowserCell) InitTextCell(string_ string) BrowserCell {
	rv := objc.CallMethod[BrowserCell](b_, objc.SEL("initTextCell:"), string_)
	return rv
}

func (b_ BrowserCell) Init() BrowserCell {
	rv := objc.CallMethod[BrowserCell](b_, objc.SEL("init"))
	return rv
}

func (bc _BrowserCellClass) Alloc() BrowserCell {
	rv := objc.CallMethod[BrowserCell](bc, objc.SEL("alloc"))
	return rv
}

func (bc _BrowserCellClass) New() BrowserCell {
	rv := objc.CallMethod[BrowserCell](bc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewBrowserCell() BrowserCell {
	return BrowserCellClass.New()
}

func (b_ BrowserCell) Reset() {
	objc.CallMethod[objc.Void](b_, objc.SEL("reset"))
}

func (b_ BrowserCell) Set() {
	objc.CallMethod[objc.Void](b_, objc.SEL("set"))
}

func (b_ BrowserCell) HighlightColorInView(controlView IView) Color {
	rv := objc.CallMethod[Color](b_, objc.SEL("highlightColorInView:"), objc.ExtractPtr(controlView))
	return rv
}

func (bc _BrowserCellClass) BranchImage() Image {
	rv := objc.CallMethod[Image](bc, objc.SEL("branchImage"))
	return rv
}

func (bc _BrowserCellClass) HighlightedBranchImage() Image {
	rv := objc.CallMethod[Image](bc, objc.SEL("highlightedBranchImage"))
	return rv
}

func (b_ BrowserCell) AlternateImage() Image {
	rv := objc.CallMethod[Image](b_, objc.SEL("alternateImage"))
	return rv
}

func (b_ BrowserCell) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setAlternateImage:"), objc.ExtractPtr(value))
}

func (b_ BrowserCell) IsLeaf() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("isLeaf"))
	return rv
}

func (b_ BrowserCell) SetLeaf(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setLeaf:"), value)
}

func (b_ BrowserCell) IsLoaded() bool {
	rv := objc.CallMethod[bool](b_, objc.SEL("isLoaded"))
	return rv
}

func (b_ BrowserCell) SetLoaded(value bool) {
	objc.CallMethod[objc.Void](b_, objc.SEL("setLoaded:"), value)
}
