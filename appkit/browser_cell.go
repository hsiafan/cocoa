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
	rv := objc.CallMethod[BrowserCell](b_, "initImageCell:", image)
	return rv
}

func (b_ BrowserCell) InitTextCell(string_ string) BrowserCell {
	rv := objc.CallMethod[BrowserCell](b_, "initTextCell:", string_)
	return rv
}

func (b_ BrowserCell) Init() BrowserCell {
	rv := objc.CallMethod[BrowserCell](b_, "init")
	return rv
}

func (bc _BrowserCellClass) Alloc() BrowserCell {
	rv := objc.CallMethod[BrowserCell](bc, "alloc")
	return rv
}

func (bc _BrowserCellClass) New() BrowserCell {
	rv := objc.CallMethod[BrowserCell](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBrowserCell() BrowserCell {
	return BrowserCellClass.New()
}

func (b_ BrowserCell) Reset() {
	objc.CallMethod[objc.Void](b_, "reset")
}

func (b_ BrowserCell) Set() {
	objc.CallMethod[objc.Void](b_, "set")
}

func (b_ BrowserCell) HighlightColorInView(controlView IView) Color {
	rv := objc.CallMethod[Color](b_, "highlightColorInView:", controlView)
	return rv
}

func (bc _BrowserCellClass) BranchImage() Image {
	rv := objc.CallMethod[Image](bc, "branchImage")
	return rv
}

func (bc _BrowserCellClass) HighlightedBranchImage() Image {
	rv := objc.CallMethod[Image](bc, "highlightedBranchImage")
	return rv
}

func (b_ BrowserCell) AlternateImage() Image {
	rv := objc.CallMethod[Image](b_, "alternateImage")
	return rv
}

func (b_ BrowserCell) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](b_, "setAlternateImage:", value)
}

func (b_ BrowserCell) IsLeaf() bool {
	rv := objc.CallMethod[bool](b_, "isLeaf")
	return rv
}

func (b_ BrowserCell) SetLeaf(value bool) {
	objc.CallMethod[objc.Void](b_, "setLeaf:", value)
}

func (b_ BrowserCell) IsLoaded() bool {
	rv := objc.CallMethod[bool](b_, "isLoaded")
	return rv
}

func (b_ BrowserCell) SetLoaded(value bool) {
	objc.CallMethod[objc.Void](b_, "setLoaded:", value)
}
