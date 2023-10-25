// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
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
	rv := objc.CallMethod[TextAttachment](t_, objc.SEL("initWithFileWrapper:"), objc.ExtractPtr(fileWrapper))
	return rv
}

func (t_ TextAttachment) InitWithData_OfType(contentData []byte, uti string) TextAttachment {
	rv := objc.CallMethod[TextAttachment](t_, objc.SEL("initWithData:ofType:"), contentData, uti)
	return rv
}

func (tc _TextAttachmentClass) Alloc() TextAttachment {
	rv := objc.CallMethod[TextAttachment](tc, objc.SEL("alloc"))
	return rv
}

func (tc _TextAttachmentClass) New() TextAttachment {
	rv := objc.CallMethod[TextAttachment](tc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewTextAttachment() TextAttachment {
	return TextAttachmentClass.New()
}

func (t_ TextAttachment) Init() TextAttachment {
	rv := objc.CallMethod[TextAttachment](t_, objc.SEL("init"))
	return rv
}

func (tc _TextAttachmentClass) RegisterTextAttachmentViewProviderClass_ForFileType(textAttachmentViewProviderClass objc.IClass, fileType string) {
	objc.CallMethod[objc.Void](tc, objc.SEL("registerTextAttachmentViewProviderClass:forFileType:"), objc.ExtractPtr(textAttachmentViewProviderClass), fileType)
}

func (tc _TextAttachmentClass) TextAttachmentViewProviderClassForFileType(fileType string) objc.Class {
	rv := objc.CallMethod[objc.Class](tc, objc.SEL("textAttachmentViewProviderClassForFileType:"), fileType)
	return rv
}

func (t_ TextAttachment) Bounds() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.SEL("bounds"))
	return rv
}

func (t_ TextAttachment) SetBounds(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setBounds:"), value)
}

func (t_ TextAttachment) Contents() []byte {
	rv := objc.CallMethod[[]byte](t_, objc.SEL("contents"))
	return rv
}

func (t_ TextAttachment) SetContents(value []byte) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setContents:"), value)
}

func (t_ TextAttachment) FileType() string {
	rv := objc.CallMethod[string](t_, objc.SEL("fileType"))
	return rv
}

func (t_ TextAttachment) SetFileType(value string) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFileType:"), value)
}

func (t_ TextAttachment) Image() Image {
	rv := objc.CallMethod[Image](t_, objc.SEL("image"))
	return rv
}

func (t_ TextAttachment) SetImage(value IImage) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setImage:"), objc.ExtractPtr(value))
}

func (t_ TextAttachment) FileWrapper() foundation.FileWrapper {
	rv := objc.CallMethod[foundation.FileWrapper](t_, objc.SEL("fileWrapper"))
	return rv
}

func (t_ TextAttachment) SetFileWrapper(value foundation.IFileWrapper) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setFileWrapper:"), objc.ExtractPtr(value))
}

func (t_ TextAttachment) AllowsTextAttachmentView() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("allowsTextAttachmentView"))
	return rv
}

func (t_ TextAttachment) SetAllowsTextAttachmentView(value bool) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setAllowsTextAttachmentView:"), value)
}

func (t_ TextAttachment) UsesTextAttachmentView() bool {
	rv := objc.CallMethod[bool](t_, objc.SEL("usesTextAttachmentView"))
	return rv
}

func (t_ TextAttachment) LineLayoutPadding() float64 {
	rv := objc.CallMethod[float64](t_, objc.SEL("lineLayoutPadding"))
	return rv
}

func (t_ TextAttachment) SetLineLayoutPadding(value float64) {
	objc.CallMethod[objc.Void](t_, objc.SEL("setLineLayoutPadding:"), value)
}
