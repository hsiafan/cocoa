// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[SharingService](s_, "initWithTitle:image:alternateImage:handler:", title, image, alternateImage, block)
	return rv
}

func (sc _SharingServiceClass) Alloc() SharingService {
	rv := objc.CallMethod[SharingService](sc, "alloc")
	return rv
}

func (sc _SharingServiceClass) New() SharingService {
	rv := objc.CallMethod[SharingService](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSharingService() SharingService {
	return SharingServiceClass.New()
}

func (s_ SharingService) Init() SharingService {
	rv := objc.CallMethod[SharingService](s_, "init")
	return rv
}

func (sc _SharingServiceClass) SharingServiceNamed(serviceName SharingServiceName) SharingService {
	rv := objc.CallMethod[SharingService](sc, "sharingServiceNamed:", serviceName)
	return rv
}

// deprecated
func (sc _SharingServiceClass) SharingServicesForItems(items []objc.IObject) []SharingService {
	rv := objc.CallMethod[[]SharingService](sc, "sharingServicesForItems:", items)
	return rv
}

func (s_ SharingService) CanPerformWithItems(items []objc.IObject) bool {
	rv := objc.CallMethod[bool](s_, "canPerformWithItems:", items)
	return rv
}

func (s_ SharingService) PerformWithItems(items []objc.IObject) {
	objc.CallMethod[objc.Void](s_, "performWithItems:", items)
}

func (s_ SharingService) Delegate() SharingServiceDelegateWrapper {
	rv := objc.CallMethod[SharingServiceDelegateWrapper](s_, "delegate")
	return rv
}

func (s_ SharingService) SetDelegate(value SharingServiceDelegate) {
	po := objc.CreateProtocol("NSSharingServiceDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(s_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, "setDelegate:", po)
}

func (s_ SharingService) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, "setDelegate:", value)
}

func (s_ SharingService) AccountName() string {
	rv := objc.CallMethod[string](s_, "accountName")
	return rv
}

func (s_ SharingService) AlternateImage() Image {
	rv := objc.CallMethod[Image](s_, "alternateImage")
	return rv
}

func (s_ SharingService) Image() Image {
	rv := objc.CallMethod[Image](s_, "image")
	return rv
}

func (s_ SharingService) Title() string {
	rv := objc.CallMethod[string](s_, "title")
	return rv
}

func (s_ SharingService) MenuItemTitle() string {
	rv := objc.CallMethod[string](s_, "menuItemTitle")
	return rv
}

func (s_ SharingService) SetMenuItemTitle(value string) {
	objc.CallMethod[objc.Void](s_, "setMenuItemTitle:", value)
}

func (s_ SharingService) Recipients() []string {
	rv := objc.CallMethod[[]string](s_, "recipients")
	return rv
}

func (s_ SharingService) SetRecipients(value []string) {
	objc.CallMethod[objc.Void](s_, "setRecipients:", value)
}

func (s_ SharingService) Subject() string {
	rv := objc.CallMethod[string](s_, "subject")
	return rv
}

func (s_ SharingService) SetSubject(value string) {
	objc.CallMethod[objc.Void](s_, "setSubject:", value)
}

func (s_ SharingService) AttachmentFileURLs() []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](s_, "attachmentFileURLs")
	return rv
}

func (s_ SharingService) MessageBody() string {
	rv := objc.CallMethod[string](s_, "messageBody")
	return rv
}

func (s_ SharingService) PermanentLink() foundation.URL {
	rv := objc.CallMethod[foundation.URL](s_, "permanentLink")
	return rv
}
