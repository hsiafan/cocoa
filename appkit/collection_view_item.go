// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewItemClass = _CollectionViewItemClass{objc.GetClass("NSCollectionViewItem")}

type _CollectionViewItemClass struct {
	objc.Class
}

type ICollectionViewItem interface {
	IViewController
	ImageView() ImageView
	SetImageView(value IImageView)
	TextField() TextField
	SetTextField(value ITextField)
	IsSelected() bool
	SetSelected(value bool)
	HighlightState() CollectionViewItemHighlightState
	SetHighlightState(value CollectionViewItemHighlightState)
	CollectionView() CollectionView
	DraggingImageComponents() []DraggingImageComponent
}

type CollectionViewItem struct {
	ViewController
}

func MakeCollectionViewItem(ptr unsafe.Pointer) CollectionViewItem {
	return CollectionViewItem{
		ViewController: MakeViewController(ptr),
	}
}

func (c_ CollectionViewItem) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "initWithNibName:bundle:", nibNameOrNil, nibBundleOrNil)
	return rv
}

func (c_ CollectionViewItem) Init() CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](c_, "init")
	return rv
}

func (cc _CollectionViewItemClass) Alloc() CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](cc, "alloc")
	return rv
}

func (cc _CollectionViewItemClass) New() CollectionViewItem {
	rv := ffi.CallMethod[CollectionViewItem](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewItem() CollectionViewItem {
	return CollectionViewItemClass.New()
}

func (c_ CollectionViewItem) ImageView() ImageView {
	rv := ffi.CallMethod[ImageView](c_, "imageView")
	return rv
}

func (c_ CollectionViewItem) SetImageView(value IImageView) {
	ffi.CallMethod[ffi.Void](c_, "setImageView:", value)
}

func (c_ CollectionViewItem) TextField() TextField {
	rv := ffi.CallMethod[TextField](c_, "textField")
	return rv
}

func (c_ CollectionViewItem) SetTextField(value ITextField) {
	ffi.CallMethod[ffi.Void](c_, "setTextField:", value)
}

func (c_ CollectionViewItem) IsSelected() bool {
	rv := ffi.CallMethod[bool](c_, "isSelected")
	return rv
}

func (c_ CollectionViewItem) SetSelected(value bool) {
	ffi.CallMethod[ffi.Void](c_, "setSelected:", value)
}

func (c_ CollectionViewItem) HighlightState() CollectionViewItemHighlightState {
	rv := ffi.CallMethod[CollectionViewItemHighlightState](c_, "highlightState")
	return rv
}

func (c_ CollectionViewItem) SetHighlightState(value CollectionViewItemHighlightState) {
	ffi.CallMethod[ffi.Void](c_, "setHighlightState:", value)
}

func (c_ CollectionViewItem) CollectionView() CollectionView {
	rv := ffi.CallMethod[CollectionView](c_, "collectionView")
	return rv
}

func (c_ CollectionViewItem) DraggingImageComponents() []DraggingImageComponent {
	rv := ffi.CallMethod[[]DraggingImageComponent](c_, "draggingImageComponents")
	return rv
}
