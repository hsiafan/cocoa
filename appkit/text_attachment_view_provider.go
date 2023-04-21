// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var TextAttachmentViewProviderClass = _TextAttachmentViewProviderClass{objc.GetClass("NSTextAttachmentViewProvider")}

type _TextAttachmentViewProviderClass struct {
	objc.Class
}

type ITextAttachmentViewProvider interface {
	objc.IObject
	LoadView()
	TextAttachment() TextAttachment
	TextLayoutManager() TextLayoutManager
	TracksTextAttachmentViewBounds() bool
	SetTracksTextAttachmentViewBounds(value bool)
	View() View
	SetView(value IView)
}

type TextAttachmentViewProvider struct {
	objc.Object
}

func MakeTextAttachmentViewProvider(ptr unsafe.Pointer) TextAttachmentViewProvider {
	return TextAttachmentViewProvider{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TextAttachmentViewProviderClass) Alloc() TextAttachmentViewProvider {
	rv := objc.CallMethod[TextAttachmentViewProvider](tc, "alloc")
	return rv
}

func (tc _TextAttachmentViewProviderClass) New() TextAttachmentViewProvider {
	rv := objc.CallMethod[TextAttachmentViewProvider](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextAttachmentViewProvider() TextAttachmentViewProvider {
	return TextAttachmentViewProviderClass.New()
}

func (t_ TextAttachmentViewProvider) Init() TextAttachmentViewProvider {
	rv := objc.CallMethod[TextAttachmentViewProvider](t_, "init")
	return rv
}

func (t_ TextAttachmentViewProvider) LoadView() {
	objc.CallMethod[objc.Void](t_, "loadView")
}

func (t_ TextAttachmentViewProvider) TextAttachment() TextAttachment {
	rv := objc.CallMethod[TextAttachment](t_, "textAttachment")
	return rv
}

func (t_ TextAttachmentViewProvider) TextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, "textLayoutManager")
	return rv
}

func (t_ TextAttachmentViewProvider) TracksTextAttachmentViewBounds() bool {
	rv := objc.CallMethod[bool](t_, "tracksTextAttachmentViewBounds")
	return rv
}

func (t_ TextAttachmentViewProvider) SetTracksTextAttachmentViewBounds(value bool) {
	objc.CallMethod[objc.Void](t_, "setTracksTextAttachmentViewBounds:", value)
}

func (t_ TextAttachmentViewProvider) View() View {
	rv := objc.CallMethod[View](t_, "view")
	return rv
}

func (t_ TextAttachmentViewProvider) SetView(value IView) {
	objc.CallMethod[objc.Void](t_, "setView:", value)
}
