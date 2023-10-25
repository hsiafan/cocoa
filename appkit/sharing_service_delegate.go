// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type SharingServiceDelegate interface {
	ImplementsSharingService_WillShareItems() bool
	// optional
	SharingService_WillShareItems(sharingService SharingService, items []objc.Object)
	ImplementsSharingService_DidShareItems() bool
	// optional
	SharingService_DidShareItems(sharingService SharingService, items []objc.Object)
	ImplementsSharingService_DidFailToShareItems_Error() bool
	// optional
	SharingService_DidFailToShareItems_Error(sharingService SharingService, items []objc.Object, error foundation.Error)
	ImplementsSharingService_SourceFrameOnScreenForShareItem() bool
	// optional
	SharingService_SourceFrameOnScreenForShareItem(sharingService SharingService, item objc.Object) foundation.Rect
	ImplementsSharingService_TransitionImageForShareItem_ContentRect() bool
	// optional
	SharingService_TransitionImageForShareItem_ContentRect(sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage
	ImplementsSharingService_SourceWindowForShareItems_SharingContentScope() bool
	// optional
	SharingService_SourceWindowForShareItems_SharingContentScope(sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow
	ImplementsAnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge() bool
	// optional
	AnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView
}

func WrapSharingServiceDelegate(v SharingServiceDelegate) objc.Object {
	return objc.WrapAsProtocol("NSSharingServiceDelegate", v)
}

type SharingServiceDelegateBase struct {
}

func (p *SharingServiceDelegateBase) ImplementsSharingService_WillShareItems() bool {
	return false
}

func (p *SharingServiceDelegateBase) SharingService_WillShareItems(sharingService SharingService, items []objc.Object) {
	panic("unimpemented protocol method")
}

func (p *SharingServiceDelegateBase) ImplementsSharingService_DidShareItems() bool {
	return false
}

func (p *SharingServiceDelegateBase) SharingService_DidShareItems(sharingService SharingService, items []objc.Object) {
	panic("unimpemented protocol method")
}

func (p *SharingServiceDelegateBase) ImplementsSharingService_DidFailToShareItems_Error() bool {
	return false
}

func (p *SharingServiceDelegateBase) SharingService_DidFailToShareItems_Error(sharingService SharingService, items []objc.Object, error foundation.Error) {
	panic("unimpemented protocol method")
}

func (p *SharingServiceDelegateBase) ImplementsSharingService_SourceFrameOnScreenForShareItem() bool {
	return false
}

func (p *SharingServiceDelegateBase) SharingService_SourceFrameOnScreenForShareItem(sharingService SharingService, item objc.Object) foundation.Rect {
	panic("unimpemented protocol method")
}

func (p *SharingServiceDelegateBase) ImplementsSharingService_TransitionImageForShareItem_ContentRect() bool {
	return false
}

func (p *SharingServiceDelegateBase) SharingService_TransitionImageForShareItem_ContentRect(sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage {
	panic("unimpemented protocol method")
}

func (p *SharingServiceDelegateBase) ImplementsSharingService_SourceWindowForShareItems_SharingContentScope() bool {
	return false
}

func (p *SharingServiceDelegateBase) SharingService_SourceWindowForShareItems_SharingContentScope(sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow {
	panic("unimpemented protocol method")
}

func (p *SharingServiceDelegateBase) ImplementsAnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge() bool {
	return false
}

func (p *SharingServiceDelegateBase) AnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView {
	panic("unimpemented protocol method")
}

type SharingServiceDelegateCreator struct {
	className string
	class     objc.Class
}

func NewSharingServiceDelegateCreator(name string) *SharingServiceDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &SharingServiceDelegateCreator{className: name, class: class}
}

func (c *SharingServiceDelegateCreator) SetSharingService_WillShareItems(handle func(o objc.Object, sharingService SharingService, items []objc.Object)) {
	objc.AddMethod(c.class, objc.GetSelector("sharingService:willShareItems:"), handle)
}

func (c *SharingServiceDelegateCreator) SetSharingService_DidShareItems(handle func(o objc.Object, sharingService SharingService, items []objc.Object)) {
	objc.AddMethod(c.class, objc.GetSelector("sharingService:didShareItems:"), handle)
}

func (c *SharingServiceDelegateCreator) SetSharingService_DidFailToShareItems_Error(handle func(o objc.Object, sharingService SharingService, items []objc.Object, error foundation.Error)) {
	objc.AddMethod(c.class, objc.GetSelector("sharingService:didFailToShareItems:error:"), handle)
}

func (c *SharingServiceDelegateCreator) SetSharingService_SourceFrameOnScreenForShareItem(handle func(o objc.Object, sharingService SharingService, item objc.Object) foundation.Rect) {
	objc.AddMethod(c.class, objc.GetSelector("sharingService:sourceFrameOnScreenForShareItem:"), handle)
}

func (c *SharingServiceDelegateCreator) SetSharingService_TransitionImageForShareItem_ContentRect(handle func(o objc.Object, sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage) {
	objc.AddMethod(c.class, objc.GetSelector("sharingService:transitionImageForShareItem:contentRect:"), handle)
}

func (c *SharingServiceDelegateCreator) SetSharingService_SourceWindowForShareItems_SharingContentScope(handle func(o objc.Object, sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow) {
	objc.AddMethod(c.class, objc.GetSelector("sharingService:sourceWindowForShareItems:sharingContentScope:"), handle)
}

func (c *SharingServiceDelegateCreator) SetAnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge(handle func(o objc.Object, sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView) {
	objc.AddMethod(c.class, objc.GetSelector("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"), handle)
}

func (c *SharingServiceDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
