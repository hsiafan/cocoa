// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[SharingService](s_, objc.SEL("initWithTitle:image:alternateImage:handler:"), title, objc.ExtractPtr(image), objc.ExtractPtr(alternateImage), block)
	return rv
}

func (sc _SharingServiceClass) Alloc() SharingService {
	rv := objc.CallMethod[SharingService](sc, objc.SEL("alloc"))
	return rv
}

func (sc _SharingServiceClass) New() SharingService {
	rv := objc.CallMethod[SharingService](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewSharingService() SharingService {
	return SharingServiceClass.New()
}

func (s_ SharingService) Init() SharingService {
	rv := objc.CallMethod[SharingService](s_, objc.SEL("init"))
	return rv
}

func (sc _SharingServiceClass) SharingServiceNamed(serviceName SharingServiceName) SharingService {
	rv := objc.CallMethod[SharingService](sc, objc.SEL("sharingServiceNamed:"), serviceName)
	return rv
}

// deprecated
func (sc _SharingServiceClass) SharingServicesForItems(items []objc.IObject) []SharingService {
	rv := objc.CallMethod[[]SharingService](sc, objc.SEL("sharingServicesForItems:"), items)
	return rv
}

func (s_ SharingService) CanPerformWithItems(items []objc.IObject) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("canPerformWithItems:"), items)
	return rv
}

func (s_ SharingService) PerformWithItems(items []objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("performWithItems:"), items)
}

// weak property
func (s_ SharingService) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("delegate"))
	return rv
}

// weak property
func (s_ SharingService) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ SharingService) AccountName() string {
	rv := objc.CallMethod[string](s_, objc.SEL("accountName"))
	return rv
}

func (s_ SharingService) AlternateImage() Image {
	rv := objc.CallMethod[Image](s_, objc.SEL("alternateImage"))
	return rv
}

func (s_ SharingService) Image() Image {
	rv := objc.CallMethod[Image](s_, objc.SEL("image"))
	return rv
}

func (s_ SharingService) Title() string {
	rv := objc.CallMethod[string](s_, objc.SEL("title"))
	return rv
}

func (s_ SharingService) MenuItemTitle() string {
	rv := objc.CallMethod[string](s_, objc.SEL("menuItemTitle"))
	return rv
}

func (s_ SharingService) SetMenuItemTitle(value string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setMenuItemTitle:"), value)
}

func (s_ SharingService) Recipients() []string {
	rv := objc.CallMethod[[]string](s_, objc.SEL("recipients"))
	return rv
}

func (s_ SharingService) SetRecipients(value []string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setRecipients:"), value)
}

func (s_ SharingService) Subject() string {
	rv := objc.CallMethod[string](s_, objc.SEL("subject"))
	return rv
}

func (s_ SharingService) SetSubject(value string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setSubject:"), value)
}

func (s_ SharingService) AttachmentFileURLs() []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](s_, objc.SEL("attachmentFileURLs"))
	return rv
}

func (s_ SharingService) MessageBody() string {
	rv := objc.CallMethod[string](s_, objc.SEL("messageBody"))
	return rv
}

func (s_ SharingService) PermanentLink() foundation.URL {
	rv := objc.CallMethod[foundation.URL](s_, objc.SEL("permanentLink"))
	return rv
}
