// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[CollectionViewItem](c_, objc.SEL("initWithNibName:bundle:"), nibNameOrNil, objc.ExtractPtr(nibBundleOrNil))
	return rv
}

func (c_ CollectionViewItem) Init() CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](c_, objc.SEL("init"))
	return rv
}

func (cc _CollectionViewItemClass) Alloc() CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CollectionViewItemClass) New() CollectionViewItem {
	rv := objc.CallMethod[CollectionViewItem](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewItem() CollectionViewItem {
	return CollectionViewItemClass.New()
}

// weak property
func (c_ CollectionViewItem) ImageView() ImageView {
	rv := objc.CallMethod[ImageView](c_, objc.SEL("imageView"))
	return rv
}

// weak property
func (c_ CollectionViewItem) SetImageView(value IImageView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setImageView:"), objc.ExtractPtr(value))
}

// weak property
func (c_ CollectionViewItem) TextField() TextField {
	rv := objc.CallMethod[TextField](c_, objc.SEL("textField"))
	return rv
}

// weak property
func (c_ CollectionViewItem) SetTextField(value ITextField) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setTextField:"), objc.ExtractPtr(value))
}

func (c_ CollectionViewItem) IsSelected() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isSelected"))
	return rv
}

func (c_ CollectionViewItem) SetSelected(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setSelected:"), value)
}

func (c_ CollectionViewItem) HighlightState() CollectionViewItemHighlightState {
	rv := objc.CallMethod[CollectionViewItemHighlightState](c_, objc.SEL("highlightState"))
	return rv
}

func (c_ CollectionViewItem) SetHighlightState(value CollectionViewItemHighlightState) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setHighlightState:"), value)
}

// weak property
func (c_ CollectionViewItem) CollectionView() CollectionView {
	rv := objc.CallMethod[CollectionView](c_, objc.SEL("collectionView"))
	return rv
}

func (c_ CollectionViewItem) DraggingImageComponents() []DraggingImageComponent {
	rv := objc.CallMethod[[]DraggingImageComponent](c_, objc.SEL("draggingImageComponents"))
	return rv
}
