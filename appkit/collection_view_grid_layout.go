// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewGridLayoutClass = _CollectionViewGridLayoutClass{objc.GetClass("NSCollectionViewGridLayout")}

type _CollectionViewGridLayoutClass struct {
	objc.Class
}

type ICollectionViewGridLayout interface {
	ICollectionViewLayout
	MaximumNumberOfRows() uint
	SetMaximumNumberOfRows(value uint)
	MaximumNumberOfColumns() uint
	SetMaximumNumberOfColumns(value uint)
	MinimumItemSize() foundation.Size
	SetMinimumItemSize(value foundation.Size)
	MaximumItemSize() foundation.Size
	SetMaximumItemSize(value foundation.Size)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	Margins() foundation.EdgeInsets
	SetMargins(value foundation.EdgeInsets)
	BackgroundColors() []Color
	SetBackgroundColors(value []IColor)
}

type CollectionViewGridLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewGridLayout(ptr unsafe.Pointer) CollectionViewGridLayout {
	return CollectionViewGridLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (cc _CollectionViewGridLayoutClass) Alloc() CollectionViewGridLayout {
	rv := ffi.CallMethod[CollectionViewGridLayout](cc, "alloc")
	return rv
}

func (cc _CollectionViewGridLayoutClass) New() CollectionViewGridLayout {
	rv := ffi.CallMethod[CollectionViewGridLayout](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewGridLayout() CollectionViewGridLayout {
	return CollectionViewGridLayoutClass.New()
}

func (c_ CollectionViewGridLayout) Init() CollectionViewGridLayout {
	rv := ffi.CallMethod[CollectionViewGridLayout](c_, "init")
	return rv
}

func (c_ CollectionViewGridLayout) MaximumNumberOfRows() uint {
	rv := ffi.CallMethod[uint](c_, "maximumNumberOfRows")
	return rv
}

func (c_ CollectionViewGridLayout) SetMaximumNumberOfRows(value uint) {
	ffi.CallMethod[ffi.Void](c_, "setMaximumNumberOfRows:", value)
}

func (c_ CollectionViewGridLayout) MaximumNumberOfColumns() uint {
	rv := ffi.CallMethod[uint](c_, "maximumNumberOfColumns")
	return rv
}

func (c_ CollectionViewGridLayout) SetMaximumNumberOfColumns(value uint) {
	ffi.CallMethod[ffi.Void](c_, "setMaximumNumberOfColumns:", value)
}

func (c_ CollectionViewGridLayout) MinimumItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "minimumItemSize")
	return rv
}

func (c_ CollectionViewGridLayout) SetMinimumItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumItemSize:", value)
}

func (c_ CollectionViewGridLayout) MaximumItemSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](c_, "maximumItemSize")
	return rv
}

func (c_ CollectionViewGridLayout) SetMaximumItemSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](c_, "setMaximumItemSize:", value)
}

func (c_ CollectionViewGridLayout) MinimumInteritemSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "minimumInteritemSpacing")
	return rv
}

func (c_ CollectionViewGridLayout) SetMinimumInteritemSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumInteritemSpacing:", value)
}

func (c_ CollectionViewGridLayout) MinimumLineSpacing() float64 {
	rv := ffi.CallMethod[float64](c_, "minimumLineSpacing")
	return rv
}

func (c_ CollectionViewGridLayout) SetMinimumLineSpacing(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setMinimumLineSpacing:", value)
}

func (c_ CollectionViewGridLayout) Margins() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](c_, "margins")
	return rv
}

func (c_ CollectionViewGridLayout) SetMargins(value foundation.EdgeInsets) {
	ffi.CallMethod[ffi.Void](c_, "setMargins:", value)
}

func (c_ CollectionViewGridLayout) BackgroundColors() []Color {
	rv := ffi.CallMethod[[]Color](c_, "backgroundColors")
	return rv
}

func (c_ CollectionViewGridLayout) SetBackgroundColors(value []IColor) {
	ffi.CallMethod[ffi.Void](c_, "setBackgroundColors:", value)
}
