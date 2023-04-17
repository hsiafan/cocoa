// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[BrowserCell](b_, "initImageCell:", image)
	return rv
}

func (b_ BrowserCell) InitTextCell(string_ string) BrowserCell {
	rv := ffi.CallMethod[BrowserCell](b_, "initTextCell:", string_)
	return rv
}

func (b_ BrowserCell) Init() BrowserCell {
	rv := ffi.CallMethod[BrowserCell](b_, "init")
	return rv
}

func (bc _BrowserCellClass) Alloc() BrowserCell {
	rv := ffi.CallMethod[BrowserCell](bc, "alloc")
	return rv
}

func (bc _BrowserCellClass) New() BrowserCell {
	rv := ffi.CallMethod[BrowserCell](bc, "new")
	rv.Autorelease()
	return rv
}

func NewBrowserCell() BrowserCell {
	return BrowserCellClass.New()
}

func (b_ BrowserCell) Reset() {
	ffi.CallMethod[ffi.Void](b_, "reset")
}

func (b_ BrowserCell) Set() {
	ffi.CallMethod[ffi.Void](b_, "set")
}

func (b_ BrowserCell) HighlightColorInView(controlView IView) Color {
	rv := ffi.CallMethod[Color](b_, "highlightColorInView:", controlView)
	return rv
}

func (bc _BrowserCellClass) BranchImage() Image {
	rv := ffi.CallMethod[Image](bc, "branchImage")
	return rv
}

func (bc _BrowserCellClass) HighlightedBranchImage() Image {
	rv := ffi.CallMethod[Image](bc, "highlightedBranchImage")
	return rv
}

func (b_ BrowserCell) AlternateImage() Image {
	rv := ffi.CallMethod[Image](b_, "alternateImage")
	return rv
}

func (b_ BrowserCell) SetAlternateImage(value IImage) {
	ffi.CallMethod[ffi.Void](b_, "setAlternateImage:", value)
}

func (b_ BrowserCell) IsLeaf() bool {
	rv := ffi.CallMethod[bool](b_, "isLeaf")
	return rv
}

func (b_ BrowserCell) SetLeaf(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setLeaf:", value)
}

func (b_ BrowserCell) IsLoaded() bool {
	rv := ffi.CallMethod[bool](b_, "isLoaded")
	return rv
}

func (b_ BrowserCell) SetLoaded(value bool) {
	ffi.CallMethod[ffi.Void](b_, "setLoaded:", value)
}
