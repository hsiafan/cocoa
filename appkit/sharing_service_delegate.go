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

type SharingServiceDelegateWrapper struct {
	objc.Object
}

func (s_ SharingServiceDelegateWrapper) SharingService_WillShareItems(sharingService ISharingService, items []objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sharingService:willShareItems:"), objc.ExtractPtr(sharingService), items)
}

func (s_ SharingServiceDelegateWrapper) SharingService_DidShareItems(sharingService ISharingService, items []objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sharingService:didShareItems:"), objc.ExtractPtr(sharingService), items)
}

func (s_ SharingServiceDelegateWrapper) SharingService_DidFailToShareItems_Error(sharingService ISharingService, items []objc.IObject, error foundation.IError) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sharingService:didFailToShareItems:error:"), objc.ExtractPtr(sharingService), items, objc.ExtractPtr(error))
}

func (s_ SharingServiceDelegateWrapper) SharingService_SourceFrameOnScreenForShareItem(sharingService ISharingService, item objc.IObject) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("sharingService:sourceFrameOnScreenForShareItem:"), objc.ExtractPtr(sharingService), objc.ExtractPtr(item))
	return rv
}

func (s_ SharingServiceDelegateWrapper) SharingService_TransitionImageForShareItem_ContentRect(sharingService ISharingService, item objc.IObject, contentRect *foundation.Rect) Image {
	rv := objc.CallMethod[Image](s_, objc.GetSelector("sharingService:transitionImageForShareItem:contentRect:"), objc.ExtractPtr(sharingService), objc.ExtractPtr(item), contentRect)
	return rv
}

func (s_ SharingServiceDelegateWrapper) SharingService_SourceWindowForShareItems_SharingContentScope(sharingService ISharingService, items []objc.IObject, sharingContentScope *SharingContentScope) Window {
	rv := objc.CallMethod[Window](s_, objc.GetSelector("sharingService:sourceWindowForShareItems:sharingContentScope:"), objc.ExtractPtr(sharingService), items, sharingContentScope)
	return rv
}

func (s_ SharingServiceDelegateWrapper) AnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge(sharingService ISharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) View {
	rv := objc.CallMethod[View](s_, objc.GetSelector("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"), objc.ExtractPtr(sharingService), positioningRect, preferredEdge)
	return rv
}
