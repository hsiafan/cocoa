// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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
