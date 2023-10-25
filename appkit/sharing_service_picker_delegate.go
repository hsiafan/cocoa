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
