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

type SharingServicePickerDelegateCreator struct {
	className string
	class     objc.Class
}

func NewSharingServicePickerDelegateCreator(name string) *SharingServicePickerDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &SharingServicePickerDelegateCreator{className: name, class: class}
}

func (c *SharingServicePickerDelegateCreator) SetSharingServicePicker_SharingServicesForItems_ProposedSharingServices(handle func(o objc.ProtocolBase, sharingServicePicker SharingServicePicker, items []objc.Object, proposedServices []SharingService) []ISharingService) {
	objc.AddMethod(c.class, objc.SEL("sharingServicePicker:sharingServicesForItems:proposedSharingServices:"), handle)
}

func (c *SharingServicePickerDelegateCreator) SetSharingServicePicker_DidChooseSharingService(handle func(o objc.ProtocolBase, sharingServicePicker SharingServicePicker, service SharingService)) {
	objc.AddMethod(c.class, objc.SEL("sharingServicePicker:didChooseSharingService:"), handle)
}

func (c *SharingServicePickerDelegateCreator) SetSharingServicePicker_DelegateForSharingService(handle func(o objc.ProtocolBase, sharingServicePicker SharingServicePicker, sharingService SharingService) objc.IObject) {
	objc.AddMethod(c.class, objc.SEL("sharingServicePicker:delegateForSharingService:"), handle)
}

func (c *SharingServicePickerDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
