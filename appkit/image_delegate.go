// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type ImageDelegate interface {
	ImplementsImageDidNotDraw_InRect() bool
	// optional
	ImageDidNotDraw_InRect(sender Image, rect foundation.Rect) IImage
	ImplementsImage_DidLoadPartOfRepresentation_WithValidRows() bool
	// optional
	Image_DidLoadPartOfRepresentation_WithValidRows(image Image, rep ImageRep, rows int)
	ImplementsImage_DidLoadRepresentation_WithStatus() bool
	// optional
	Image_DidLoadRepresentation_WithStatus(image Image, rep ImageRep, status ImageLoadStatus)
	ImplementsImage_DidLoadRepresentationHeader() bool
	// optional
	Image_DidLoadRepresentationHeader(image Image, rep ImageRep)
	ImplementsImage_WillLoadRepresentation() bool
	// optional
	Image_WillLoadRepresentation(image Image, rep ImageRep)
}

type ImageDelegateImpl struct {
	_ImageDidNotDraw_InRect                          func(sender Image, rect foundation.Rect) IImage
	_Image_DidLoadPartOfRepresentation_WithValidRows func(image Image, rep ImageRep, rows int)
	_Image_DidLoadRepresentation_WithStatus          func(image Image, rep ImageRep, status ImageLoadStatus)
	_Image_DidLoadRepresentationHeader               func(image Image, rep ImageRep)
	_Image_WillLoadRepresentation                    func(image Image, rep ImageRep)
}

func (di *ImageDelegateImpl) ImplementsImageDidNotDraw_InRect() bool {
	return di._ImageDidNotDraw_InRect != nil
}

func (di *ImageDelegateImpl) SetImageDidNotDraw_InRect(f func(sender Image, rect foundation.Rect) IImage) {
	di._ImageDidNotDraw_InRect = f
}

func (di *ImageDelegateImpl) ImageDidNotDraw_InRect(sender Image, rect foundation.Rect) IImage {
	return di._ImageDidNotDraw_InRect(sender, rect)
}
func (di *ImageDelegateImpl) ImplementsImage_DidLoadPartOfRepresentation_WithValidRows() bool {
	return di._Image_DidLoadPartOfRepresentation_WithValidRows != nil
}

func (di *ImageDelegateImpl) SetImage_DidLoadPartOfRepresentation_WithValidRows(f func(image Image, rep ImageRep, rows int)) {
	di._Image_DidLoadPartOfRepresentation_WithValidRows = f
}

func (di *ImageDelegateImpl) Image_DidLoadPartOfRepresentation_WithValidRows(image Image, rep ImageRep, rows int) {
	di._Image_DidLoadPartOfRepresentation_WithValidRows(image, rep, rows)
}
func (di *ImageDelegateImpl) ImplementsImage_DidLoadRepresentation_WithStatus() bool {
	return di._Image_DidLoadRepresentation_WithStatus != nil
}

func (di *ImageDelegateImpl) SetImage_DidLoadRepresentation_WithStatus(f func(image Image, rep ImageRep, status ImageLoadStatus)) {
	di._Image_DidLoadRepresentation_WithStatus = f
}

func (di *ImageDelegateImpl) Image_DidLoadRepresentation_WithStatus(image Image, rep ImageRep, status ImageLoadStatus) {
	di._Image_DidLoadRepresentation_WithStatus(image, rep, status)
}
func (di *ImageDelegateImpl) ImplementsImage_DidLoadRepresentationHeader() bool {
	return di._Image_DidLoadRepresentationHeader != nil
}

func (di *ImageDelegateImpl) SetImage_DidLoadRepresentationHeader(f func(image Image, rep ImageRep)) {
	di._Image_DidLoadRepresentationHeader = f
}

func (di *ImageDelegateImpl) Image_DidLoadRepresentationHeader(image Image, rep ImageRep) {
	di._Image_DidLoadRepresentationHeader(image, rep)
}
func (di *ImageDelegateImpl) ImplementsImage_WillLoadRepresentation() bool {
	return di._Image_WillLoadRepresentation != nil
}

func (di *ImageDelegateImpl) SetImage_WillLoadRepresentation(f func(image Image, rep ImageRep)) {
	di._Image_WillLoadRepresentation = f
}

func (di *ImageDelegateImpl) Image_WillLoadRepresentation(image Image, rep ImageRep) {
	di._Image_WillLoadRepresentation(image, rep)
}

type ImageDelegateWrapper struct {
	objc.Object
}

func (i_ *ImageDelegateWrapper) ImplementsImageDidNotDraw_InRect() bool {
	return i_.RespondsToSelector(objc.GetSelector("imageDidNotDraw:inRect:"))
}

func (i_ ImageDelegateWrapper) ImageDidNotDraw_InRect(sender IImage, rect foundation.Rect) Image {
	rv := objc.CallMethod[Image](i_, "imageDidNotDraw:inRect:", sender, rect)
	return rv
}

func (i_ *ImageDelegateWrapper) ImplementsImage_DidLoadPartOfRepresentation_WithValidRows() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:didLoadPartOfRepresentation:withValidRows:"))
}

func (i_ ImageDelegateWrapper) Image_DidLoadPartOfRepresentation_WithValidRows(image IImage, rep IImageRep, rows int) {
	objc.CallMethod[objc.Void](i_, "image:didLoadPartOfRepresentation:withValidRows:", image, rep, rows)
}

func (i_ *ImageDelegateWrapper) ImplementsImage_DidLoadRepresentation_WithStatus() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:didLoadRepresentation:withStatus:"))
}

func (i_ ImageDelegateWrapper) Image_DidLoadRepresentation_WithStatus(image IImage, rep IImageRep, status ImageLoadStatus) {
	objc.CallMethod[objc.Void](i_, "image:didLoadRepresentation:withStatus:", image, rep, status)
}

func (i_ *ImageDelegateWrapper) ImplementsImage_DidLoadRepresentationHeader() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:didLoadRepresentationHeader:"))
}

func (i_ ImageDelegateWrapper) Image_DidLoadRepresentationHeader(image IImage, rep IImageRep) {
	objc.CallMethod[objc.Void](i_, "image:didLoadRepresentationHeader:", image, rep)
}

func (i_ *ImageDelegateWrapper) ImplementsImage_WillLoadRepresentation() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:willLoadRepresentation:"))
}

func (i_ ImageDelegateWrapper) Image_WillLoadRepresentation(image IImage, rep IImageRep) {
	objc.CallMethod[objc.Void](i_, "image:willLoadRepresentation:", image, rep)
}
