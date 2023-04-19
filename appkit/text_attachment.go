// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextAttachmentClass = _TextAttachmentClass{objc.GetClass("NSTextAttachment")}

type _TextAttachmentClass struct {
	objc.Class
}

type ITextAttachment interface {
	objc.IObject
	Bounds() coregraphics.Rect
	SetBounds(value coregraphics.Rect)
	Contents() []byte
	SetContents(value []byte)
	FileType() string
	SetFileType(value string)
	Image() Image
	SetImage(value IImage)
	FileWrapper() foundation.FileWrapper
	SetFileWrapper(value foundation.IFileWrapper)
	AllowsTextAttachmentView() bool
	SetAllowsTextAttachmentView(value bool)
	UsesTextAttachmentView() bool
	LineLayoutPadding() float64
	SetLineLayoutPadding(value float64)
}

type TextAttachment struct {
	objc.Object
}

func MakeTextAttachment(ptr unsafe.Pointer) TextAttachment {
	return TextAttachment{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextAttachment) InitWithFileWrapper(fileWrapper foundation.IFileWrapper) TextAttachment {
	rv := ffi.CallMethod[TextAttachment](t_, "initWithFileWrapper:", fileWrapper)
	return rv
}

func (t_ TextAttachment) InitWithData_OfType(contentData []byte, uti string) TextAttachment {
	rv := ffi.CallMethod[TextAttachment](t_, "initWithData:ofType:", contentData, uti)
	return rv
}

func (tc _TextAttachmentClass) Alloc() TextAttachment {
	rv := ffi.CallMethod[TextAttachment](tc, "alloc")
	return rv
}

func (tc _TextAttachmentClass) New() TextAttachment {
	rv := ffi.CallMethod[TextAttachment](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextAttachment() TextAttachment {
	return TextAttachmentClass.New()
}

func (t_ TextAttachment) Init() TextAttachment {
	rv := ffi.CallMethod[TextAttachment](t_, "init")
	return rv
}

func (tc _TextAttachmentClass) RegisterTextAttachmentViewProviderClass_ForFileType(textAttachmentViewProviderClass objc.IClass, fileType string) {
	ffi.CallMethod[ffi.Void](tc, "registerTextAttachmentViewProviderClass:forFileType:", textAttachmentViewProviderClass, fileType)
}

func (tc _TextAttachmentClass) TextAttachmentViewProviderClassForFileType(fileType string) objc.Class {
	rv := ffi.CallMethod[objc.Class](tc, "textAttachmentViewProviderClassForFileType:", fileType)
	return rv
}

func (t_ TextAttachment) Bounds() coregraphics.Rect {
	rv := ffi.CallMethod[coregraphics.Rect](t_, "bounds")
	return rv
}

func (t_ TextAttachment) SetBounds(value coregraphics.Rect) {
	ffi.CallMethod[ffi.Void](t_, "setBounds:", value)
}

func (t_ TextAttachment) Contents() []byte {
	rv := ffi.CallMethod[[]byte](t_, "contents")
	return rv
}

func (t_ TextAttachment) SetContents(value []byte) {
	ffi.CallMethod[ffi.Void](t_, "setContents:", value)
}

func (t_ TextAttachment) FileType() string {
	rv := ffi.CallMethod[string](t_, "fileType")
	return rv
}

func (t_ TextAttachment) SetFileType(value string) {
	ffi.CallMethod[ffi.Void](t_, "setFileType:", value)
}

func (t_ TextAttachment) Image() Image {
	rv := ffi.CallMethod[Image](t_, "image")
	return rv
}

func (t_ TextAttachment) SetImage(value IImage) {
	ffi.CallMethod[ffi.Void](t_, "setImage:", value)
}

func (t_ TextAttachment) FileWrapper() foundation.FileWrapper {
	rv := ffi.CallMethod[foundation.FileWrapper](t_, "fileWrapper")
	return rv
}

func (t_ TextAttachment) SetFileWrapper(value foundation.IFileWrapper) {
	ffi.CallMethod[ffi.Void](t_, "setFileWrapper:", value)
}

func (t_ TextAttachment) AllowsTextAttachmentView() bool {
	rv := ffi.CallMethod[bool](t_, "allowsTextAttachmentView")
	return rv
}

func (t_ TextAttachment) SetAllowsTextAttachmentView(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAllowsTextAttachmentView:", value)
}

func (t_ TextAttachment) UsesTextAttachmentView() bool {
	rv := ffi.CallMethod[bool](t_, "usesTextAttachmentView")
	return rv
}

func (t_ TextAttachment) LineLayoutPadding() float64 {
	rv := ffi.CallMethod[float64](t_, "lineLayoutPadding")
	return rv
}

func (t_ TextAttachment) SetLineLayoutPadding(value float64) {
	ffi.CallMethod[ffi.Void](t_, "setLineLayoutPadding:", value)
}
