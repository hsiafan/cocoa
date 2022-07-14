// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var SharingServiceClass = _SharingServiceClass{objc.GetClass("NSSharingService")}

type _SharingServiceClass struct {
	objc.Class
}

type ISharingService interface {
	objc.IObject
	CanPerformWithItems(items []objc.IObject) bool
	PerformWithItems(items []objc.IObject)
	Delegate() SharingServiceDelegateWrapper
	SetDelegate(value SharingServiceDelegate)
	AccountName() string
	AlternateImage() Image
	Image() Image
	Title() string
	MenuItemTitle() string
	SetMenuItemTitle(value string)
	Recipients() []string
	SetRecipients(value []string)
	Subject() string
	SetSubject(value string)
	AttachmentFileURLs() []foundation.URL
	MessageBody() string
	PermanentLink() foundation.URL
}

type SharingService struct {
	objc.Object
}

func MakeSharingService(ptr unsafe.Pointer) SharingService {
	return SharingService{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ SharingService) InitWithTitle_Image_AlternateImage_Handler(title string, image IImage, alternateImage IImage, block func()) SharingService {
	rv := ffi.CallMethod[SharingService](s_, "initWithTitle:image:alternateImage:handler:", title, image, alternateImage, block)
	rv.Autorelease()
	return rv
}

func (sc _SharingServiceClass) Alloc() SharingService {
	rv := ffi.CallMethod[SharingService](sc, "alloc")
	return rv
}

func (s_ SharingService) Init() SharingService {
	rv := ffi.CallMethod[SharingService](s_, "init")
	rv.Autorelease()
	return rv
}

func (sc _SharingServiceClass) New() SharingService {
	rv := ffi.CallMethod[SharingService](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSharingService() SharingService {
	return SharingServiceClass.New()
}

func (sc _SharingServiceClass) SharingServiceNamed(serviceName SharingServiceName) SharingService {
	rv := ffi.CallMethod[SharingService](sc, "sharingServiceNamed:", serviceName)
	return rv
}

// deprecated
func (sc _SharingServiceClass) SharingServicesForItems(items []objc.IObject) []SharingService {
	rv := ffi.CallMethod[[]SharingService](sc, "sharingServicesForItems:", items)
	return rv
}

func (s_ SharingService) CanPerformWithItems(items []objc.IObject) bool {
	rv := ffi.CallMethod[bool](s_, "canPerformWithItems:", items)
	return rv
}

func (s_ SharingService) PerformWithItems(items []objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "performWithItems:", items)
}

func (s_ SharingService) Delegate() SharingServiceDelegateWrapper {
	rv := ffi.CallMethod[SharingServiceDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ SharingService) SetDelegate(value SharingServiceDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](s_, "setDelegate:", po)
}

func (s_ SharingService) AccountName() string {
	rv := ffi.CallMethod[string](s_, "accountName")
	return rv
}

func (s_ SharingService) AlternateImage() Image {
	rv := ffi.CallMethod[Image](s_, "alternateImage")
	return rv
}

func (s_ SharingService) Image() Image {
	rv := ffi.CallMethod[Image](s_, "image")
	return rv
}

func (s_ SharingService) Title() string {
	rv := ffi.CallMethod[string](s_, "title")
	return rv
}

func (s_ SharingService) MenuItemTitle() string {
	rv := ffi.CallMethod[string](s_, "menuItemTitle")
	return rv
}

func (s_ SharingService) SetMenuItemTitle(value string) {
	ffi.CallMethod[ffi.Void](s_, "setMenuItemTitle:", value)
}

func (s_ SharingService) Recipients() []string {
	rv := ffi.CallMethod[[]string](s_, "recipients")
	return rv
}

func (s_ SharingService) SetRecipients(value []string) {
	ffi.CallMethod[ffi.Void](s_, "setRecipients:", value)
}

func (s_ SharingService) Subject() string {
	rv := ffi.CallMethod[string](s_, "subject")
	return rv
}

func (s_ SharingService) SetSubject(value string) {
	ffi.CallMethod[ffi.Void](s_, "setSubject:", value)
}

func (s_ SharingService) AttachmentFileURLs() []foundation.URL {
	rv := ffi.CallMethod[[]foundation.URL](s_, "attachmentFileURLs")
	return rv
}

func (s_ SharingService) MessageBody() string {
	rv := ffi.CallMethod[string](s_, "messageBody")
	return rv
}

func (s_ SharingService) PermanentLink() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](s_, "permanentLink")
	return rv
}

var SharingServicePickerClass = _SharingServicePickerClass{objc.GetClass("NSSharingServicePicker")}

type _SharingServicePickerClass struct {
	objc.Class
}

type ISharingServicePicker interface {
	objc.IObject
	ShowRelativeToRect_OfView_PreferredEdge(rect foundation.Rect, view IView, preferredEdge foundation.RectEdge)
	Delegate() SharingServicePickerDelegateWrapper
	SetDelegate(value SharingServicePickerDelegate)
	StandardShareMenuItem() MenuItem
}

type SharingServicePicker struct {
	objc.Object
}

func MakeSharingServicePicker(ptr unsafe.Pointer) SharingServicePicker {
	return SharingServicePicker{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ SharingServicePicker) InitWithItems(items []objc.IObject) SharingServicePicker {
	rv := ffi.CallMethod[SharingServicePicker](s_, "initWithItems:", items)
	rv.Autorelease()
	return rv
}

func (sc _SharingServicePickerClass) Alloc() SharingServicePicker {
	rv := ffi.CallMethod[SharingServicePicker](sc, "alloc")
	return rv
}

func (s_ SharingServicePicker) Init() SharingServicePicker {
	rv := ffi.CallMethod[SharingServicePicker](s_, "init")
	rv.Autorelease()
	return rv
}

func (sc _SharingServicePickerClass) New() SharingServicePicker {
	rv := ffi.CallMethod[SharingServicePicker](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSharingServicePicker() SharingServicePicker {
	return SharingServicePickerClass.New()
}

func (s_ SharingServicePicker) ShowRelativeToRect_OfView_PreferredEdge(rect foundation.Rect, view IView, preferredEdge foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](s_, "showRelativeToRect:ofView:preferredEdge:", rect, view, preferredEdge)
}

func (s_ SharingServicePicker) Delegate() SharingServicePickerDelegateWrapper {
	rv := ffi.CallMethod[SharingServicePickerDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ SharingServicePicker) SetDelegate(value SharingServicePickerDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](s_, "setDelegate:", po)
}

func (s_ SharingServicePicker) StandardShareMenuItem() MenuItem {
	rv := ffi.CallMethod[MenuItem](s_, "standardShareMenuItem")
	return rv
}

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

type SharingServicePickerDelegate interface {
	ImplementsSharingServicePicker_SharingServicesForItems_ProposedSharingServices() bool
	// optional
	SharingServicePicker_SharingServicesForItems_ProposedSharingServices(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService
	ImplementsSharingServicePicker_DidChooseSharingService() bool
	// optional
	SharingServicePicker_DidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService)
	ImplementsSharingServicePicker_DelegateForSharingService() bool
	// optional
	SharingServicePicker_DelegateForSharingService(sharingServicePicker SharingServicePicker, sharingService SharingService) SharingServiceDelegate
}

type SharingServicePickerDelegateImpl struct {
	_SharingServicePicker_SharingServicesForItems_ProposedSharingServices func(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService
	_SharingServicePicker_DidChooseSharingService                         func(sharingServicePicker SharingServicePicker, service SharingService)
	_SharingServicePicker_DelegateForSharingService                       func(sharingServicePicker SharingServicePicker, sharingService SharingService) SharingServiceDelegate
}

func (di *SharingServicePickerDelegateImpl) ImplementsSharingServicePicker_SharingServicesForItems_ProposedSharingServices() bool {
	return di._SharingServicePicker_SharingServicesForItems_ProposedSharingServices != nil
}

func (di *SharingServicePickerDelegateImpl) SetSharingServicePicker_SharingServicesForItems_ProposedSharingServices(f func(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService) {
	di._SharingServicePicker_SharingServicesForItems_ProposedSharingServices = f
}

func (di *SharingServicePickerDelegateImpl) SharingServicePicker_SharingServicesForItems_ProposedSharingServices(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService {
	return di._SharingServicePicker_SharingServicesForItems_ProposedSharingServices(sharingServicePicker, items, proposedServices)
}
func (di *SharingServicePickerDelegateImpl) ImplementsSharingServicePicker_DidChooseSharingService() bool {
	return di._SharingServicePicker_DidChooseSharingService != nil
}

func (di *SharingServicePickerDelegateImpl) SetSharingServicePicker_DidChooseSharingService(f func(sharingServicePicker SharingServicePicker, service SharingService)) {
	di._SharingServicePicker_DidChooseSharingService = f
}

func (di *SharingServicePickerDelegateImpl) SharingServicePicker_DidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService) {
	di._SharingServicePicker_DidChooseSharingService(sharingServicePicker, service)
}
func (di *SharingServicePickerDelegateImpl) ImplementsSharingServicePicker_DelegateForSharingService() bool {
	return di._SharingServicePicker_DelegateForSharingService != nil
}

func (di *SharingServicePickerDelegateImpl) SetSharingServicePicker_DelegateForSharingService(f func(sharingServicePicker SharingServicePicker, sharingService SharingService) SharingServiceDelegate) {
	di._SharingServicePicker_DelegateForSharingService = f
}

func (di *SharingServicePickerDelegateImpl) SharingServicePicker_DelegateForSharingService(sharingServicePicker SharingServicePicker, sharingService SharingService) SharingServiceDelegate {
	return di._SharingServicePicker_DelegateForSharingService(sharingServicePicker, sharingService)
}

type SharingServicePickerDelegateWrapper struct {
	objc.Object
}

func (s_ *SharingServicePickerDelegateWrapper) ImplementsSharingServicePicker_SharingServicesForItems_ProposedSharingServices() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingServicePicker:sharingServicesForItems:proposedSharingServices:"))
}

func (s_ SharingServicePickerDelegateWrapper) SharingServicePicker_SharingServicesForItems_ProposedSharingServices(sharingServicePicker ISharingServicePicker, items []objc.IObject, proposedServices []ISharingService) []SharingService {
	rv := ffi.CallMethod[[]SharingService](s_, "sharingServicePicker:sharingServicesForItems:proposedSharingServices:", sharingServicePicker, items, proposedServices)
	return rv
}

func (s_ *SharingServicePickerDelegateWrapper) ImplementsSharingServicePicker_DidChooseSharingService() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingServicePicker:didChooseSharingService:"))
}

func (s_ SharingServicePickerDelegateWrapper) SharingServicePicker_DidChooseSharingService(sharingServicePicker ISharingServicePicker, service ISharingService) {
	ffi.CallMethod[ffi.Void](s_, "sharingServicePicker:didChooseSharingService:", sharingServicePicker, service)
}

func (s_ *SharingServicePickerDelegateWrapper) ImplementsSharingServicePicker_DelegateForSharingService() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingServicePicker:delegateForSharingService:"))
}

func (s_ SharingServicePickerDelegateWrapper) SharingServicePicker_DelegateForSharingService(sharingServicePicker ISharingServicePicker, sharingService ISharingService) SharingServiceDelegateWrapper {
	rv := ffi.CallMethod[SharingServiceDelegateWrapper](s_, "sharingServicePicker:delegateForSharingService:", sharingServicePicker, sharingService)
	return rv
}
