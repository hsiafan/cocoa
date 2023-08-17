// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type SharingServicePickerDelegate interface {
	ImplementsSharingServicePicker_SharingServicesForItems_ProposedSharingServices() bool
	// optional
	SharingServicePicker_SharingServicesForItems_ProposedSharingServices(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService
	ImplementsSharingServicePicker_DidChooseSharingService() bool
	// optional
	SharingServicePicker_DidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService)
	ImplementsSharingServicePicker_DelegateForSharingService() bool
	// optional
	SharingServicePicker_DelegateForSharingService(sharingServicePicker SharingServicePicker, sharingService SharingService) objc.IObject
}

func WrapSharingServicePickerDelegate(v SharingServicePickerDelegate) objc.Object {
	return objc.WrapAsProtocol("NSSharingServicePickerDelegate", v)
}

type SharingServicePickerDelegateBase struct {
}

func (p *SharingServicePickerDelegateBase) ImplementsSharingServicePicker_SharingServicesForItems_ProposedSharingServices() bool {
	return false
}

func (p *SharingServicePickerDelegateBase) SharingServicePicker_SharingServicesForItems_ProposedSharingServices(sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService {
	panic("unimpemented protocol method")
}

func (p *SharingServicePickerDelegateBase) ImplementsSharingServicePicker_DidChooseSharingService() bool {
	return false
}

func (p *SharingServicePickerDelegateBase) SharingServicePicker_DidChooseSharingService(sharingServicePicker SharingServicePicker, service SharingService) {
	panic("unimpemented protocol method")
}

func (p *SharingServicePickerDelegateBase) ImplementsSharingServicePicker_DelegateForSharingService() bool {
	return false
}

func (p *SharingServicePickerDelegateBase) SharingServicePicker_DelegateForSharingService(sharingServicePicker SharingServicePicker, sharingService SharingService) objc.IObject {
	panic("unimpemented protocol method")
}

type SharingServicePickerDelegateWrapper struct {
	objc.Object
}

func (s_ SharingServicePickerDelegateWrapper) SharingServicePicker_SharingServicesForItems_ProposedSharingServices(sharingServicePicker ISharingServicePicker, items []objc.IObject, proposedServices []ISharingService) []SharingService {
	rv := objc.CallMethod[[]SharingService](s_, objc.GetSelector("sharingServicePicker:sharingServicesForItems:proposedSharingServices:"), objc.ExtractPtr(sharingServicePicker), items, proposedServices)
	return rv
}

func (s_ SharingServicePickerDelegateWrapper) SharingServicePicker_DidChooseSharingService(sharingServicePicker ISharingServicePicker, service ISharingService) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sharingServicePicker:didChooseSharingService:"), objc.ExtractPtr(sharingServicePicker), objc.ExtractPtr(service))
}

func (s_ SharingServicePickerDelegateWrapper) SharingServicePicker_DelegateForSharingService(sharingServicePicker ISharingServicePicker, sharingService ISharingService) SharingServiceDelegateWrapper {
	rv := objc.CallMethod[SharingServiceDelegateWrapper](s_, objc.GetSelector("sharingServicePicker:delegateForSharingService:"), objc.ExtractPtr(sharingServicePicker), objc.ExtractPtr(sharingService))
	return rv
}
