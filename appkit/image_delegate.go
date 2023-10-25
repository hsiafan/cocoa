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

func WrapImageDelegate(v ImageDelegate) objc.Object {
	return objc.WrapAsProtocol("NSImageDelegate", v)
}

type ImageDelegateBase struct {
}

func (p *ImageDelegateBase) ImplementsImageDidNotDraw_InRect() bool {
	return false
}

func (p *ImageDelegateBase) ImageDidNotDraw_InRect(sender Image, rect foundation.Rect) IImage {
	panic("unimpemented protocol method")
}

func (p *ImageDelegateBase) ImplementsImage_DidLoadPartOfRepresentation_WithValidRows() bool {
	return false
}

func (p *ImageDelegateBase) Image_DidLoadPartOfRepresentation_WithValidRows(image Image, rep ImageRep, rows int) {
	panic("unimpemented protocol method")
}

func (p *ImageDelegateBase) ImplementsImage_DidLoadRepresentation_WithStatus() bool {
	return false
}

func (p *ImageDelegateBase) Image_DidLoadRepresentation_WithStatus(image Image, rep ImageRep, status ImageLoadStatus) {
	panic("unimpemented protocol method")
}

func (p *ImageDelegateBase) ImplementsImage_DidLoadRepresentationHeader() bool {
	return false
}

func (p *ImageDelegateBase) Image_DidLoadRepresentationHeader(image Image, rep ImageRep) {
	panic("unimpemented protocol method")
}

func (p *ImageDelegateBase) ImplementsImage_WillLoadRepresentation() bool {
	return false
}

func (p *ImageDelegateBase) Image_WillLoadRepresentation(image Image, rep ImageRep) {
	panic("unimpemented protocol method")
}
