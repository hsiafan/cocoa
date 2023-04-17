// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type SharingServiceDelegateImpl struct {
	_SharingService_WillShareItems                                   func(sharingService SharingService, items []objc.Object)
	_SharingService_DidShareItems                                    func(sharingService SharingService, items []objc.Object)
	_SharingService_DidFailToShareItems_Error                        func(sharingService SharingService, items []objc.Object, error foundation.Error)
	_SharingService_SourceFrameOnScreenForShareItem                  func(sharingService SharingService, item objc.Object) foundation.Rect
	_SharingService_TransitionImageForShareItem_ContentRect          func(sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage
	_SharingService_SourceWindowForShareItems_SharingContentScope    func(sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow
	_AnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge func(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView
}

func (di *SharingServiceDelegateImpl) ImplementsSharingService_WillShareItems() bool {
	return di._SharingService_WillShareItems != nil
}

func (di *SharingServiceDelegateImpl) SetSharingService_WillShareItems(f func(sharingService SharingService, items []objc.Object)) {
	di._SharingService_WillShareItems = f
}

func (di *SharingServiceDelegateImpl) SharingService_WillShareItems(sharingService SharingService, items []objc.Object) {
	di._SharingService_WillShareItems(sharingService, items)
}
func (di *SharingServiceDelegateImpl) ImplementsSharingService_DidShareItems() bool {
	return di._SharingService_DidShareItems != nil
}

func (di *SharingServiceDelegateImpl) SetSharingService_DidShareItems(f func(sharingService SharingService, items []objc.Object)) {
	di._SharingService_DidShareItems = f
}

func (di *SharingServiceDelegateImpl) SharingService_DidShareItems(sharingService SharingService, items []objc.Object) {
	di._SharingService_DidShareItems(sharingService, items)
}
func (di *SharingServiceDelegateImpl) ImplementsSharingService_DidFailToShareItems_Error() bool {
	return di._SharingService_DidFailToShareItems_Error != nil
}

func (di *SharingServiceDelegateImpl) SetSharingService_DidFailToShareItems_Error(f func(sharingService SharingService, items []objc.Object, error foundation.Error)) {
	di._SharingService_DidFailToShareItems_Error = f
}

func (di *SharingServiceDelegateImpl) SharingService_DidFailToShareItems_Error(sharingService SharingService, items []objc.Object, error foundation.Error) {
	di._SharingService_DidFailToShareItems_Error(sharingService, items, error)
}
func (di *SharingServiceDelegateImpl) ImplementsSharingService_SourceFrameOnScreenForShareItem() bool {
	return di._SharingService_SourceFrameOnScreenForShareItem != nil
}

func (di *SharingServiceDelegateImpl) SetSharingService_SourceFrameOnScreenForShareItem(f func(sharingService SharingService, item objc.Object) foundation.Rect) {
	di._SharingService_SourceFrameOnScreenForShareItem = f
}

func (di *SharingServiceDelegateImpl) SharingService_SourceFrameOnScreenForShareItem(sharingService SharingService, item objc.Object) foundation.Rect {
	return di._SharingService_SourceFrameOnScreenForShareItem(sharingService, item)
}
func (di *SharingServiceDelegateImpl) ImplementsSharingService_TransitionImageForShareItem_ContentRect() bool {
	return di._SharingService_TransitionImageForShareItem_ContentRect != nil
}

func (di *SharingServiceDelegateImpl) SetSharingService_TransitionImageForShareItem_ContentRect(f func(sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage) {
	di._SharingService_TransitionImageForShareItem_ContentRect = f
}

func (di *SharingServiceDelegateImpl) SharingService_TransitionImageForShareItem_ContentRect(sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage {
	return di._SharingService_TransitionImageForShareItem_ContentRect(sharingService, item, contentRect)
}
func (di *SharingServiceDelegateImpl) ImplementsSharingService_SourceWindowForShareItems_SharingContentScope() bool {
	return di._SharingService_SourceWindowForShareItems_SharingContentScope != nil
}

func (di *SharingServiceDelegateImpl) SetSharingService_SourceWindowForShareItems_SharingContentScope(f func(sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow) {
	di._SharingService_SourceWindowForShareItems_SharingContentScope = f
}

func (di *SharingServiceDelegateImpl) SharingService_SourceWindowForShareItems_SharingContentScope(sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow {
	return di._SharingService_SourceWindowForShareItems_SharingContentScope(sharingService, items, sharingContentScope)
}
func (di *SharingServiceDelegateImpl) ImplementsAnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge() bool {
	return di._AnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge != nil
}

func (di *SharingServiceDelegateImpl) SetAnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge(f func(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView) {
	di._AnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge = f
}

func (di *SharingServiceDelegateImpl) AnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView {
	return di._AnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge(sharingService, positioningRect, preferredEdge)
}

type SharingServiceDelegateWrapper struct {
	objc.Object
}

func (s_ *SharingServiceDelegateWrapper) ImplementsSharingService_WillShareItems() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:willShareItems:"))
}

func (s_ SharingServiceDelegateWrapper) SharingService_WillShareItems(sharingService ISharingService, items []objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "sharingService:willShareItems:", sharingService, items)
}

func (s_ *SharingServiceDelegateWrapper) ImplementsSharingService_DidShareItems() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:didShareItems:"))
}

func (s_ SharingServiceDelegateWrapper) SharingService_DidShareItems(sharingService ISharingService, items []objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "sharingService:didShareItems:", sharingService, items)
}

func (s_ *SharingServiceDelegateWrapper) ImplementsSharingService_DidFailToShareItems_Error() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:didFailToShareItems:error:"))
}

func (s_ SharingServiceDelegateWrapper) SharingService_DidFailToShareItems_Error(sharingService ISharingService, items []objc.IObject, error foundation.IError) {
	ffi.CallMethod[ffi.Void](s_, "sharingService:didFailToShareItems:error:", sharingService, items, error)
}

func (s_ *SharingServiceDelegateWrapper) ImplementsSharingService_SourceFrameOnScreenForShareItem() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:sourceFrameOnScreenForShareItem:"))
}

func (s_ SharingServiceDelegateWrapper) SharingService_SourceFrameOnScreenForShareItem(sharingService ISharingService, item objc.IObject) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](s_, "sharingService:sourceFrameOnScreenForShareItem:", sharingService, item)
	return rv
}

func (s_ *SharingServiceDelegateWrapper) ImplementsSharingService_TransitionImageForShareItem_ContentRect() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:transitionImageForShareItem:contentRect:"))
}

func (s_ SharingServiceDelegateWrapper) SharingService_TransitionImageForShareItem_ContentRect(sharingService ISharingService, item objc.IObject, contentRect *foundation.Rect) Image {
	rv := ffi.CallMethod[Image](s_, "sharingService:transitionImageForShareItem:contentRect:", sharingService, item, contentRect)
	return rv
}

func (s_ *SharingServiceDelegateWrapper) ImplementsSharingService_SourceWindowForShareItems_SharingContentScope() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:sourceWindowForShareItems:sharingContentScope:"))
}

func (s_ SharingServiceDelegateWrapper) SharingService_SourceWindowForShareItems_SharingContentScope(sharingService ISharingService, items []objc.IObject, sharingContentScope *SharingContentScope) Window {
	rv := ffi.CallMethod[Window](s_, "sharingService:sourceWindowForShareItems:sharingContentScope:", sharingService, items, sharingContentScope)
	return rv
}

func (s_ *SharingServiceDelegateWrapper) ImplementsAnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge() bool {
	return s_.RespondsToSelector(objc.GetSelector("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"))
}

func (s_ SharingServiceDelegateWrapper) AnchoringViewForSharingService_ShowRelativeToRect_PreferredEdge(sharingService ISharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) View {
	rv := ffi.CallMethod[View](s_, "anchoringViewForSharingService:showRelativeToRect:preferredEdge:", sharingService, positioningRect, preferredEdge)
	return rv
}
