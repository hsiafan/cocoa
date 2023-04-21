// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var PersonNameComponentsClass = _PersonNameComponentsClass{objc.GetClass("NSPersonNameComponents")}

type _PersonNameComponentsClass struct {
	objc.Class
}

type IPersonNameComponents interface {
	objc.IObject
	NamePrefix() string
	SetNamePrefix(value string)
	GivenName() string
	SetGivenName(value string)
	MiddleName() string
	SetMiddleName(value string)
	FamilyName() string
	SetFamilyName(value string)
	NameSuffix() string
	SetNameSuffix(value string)
	Nickname() string
	SetNickname(value string)
	PhoneticRepresentation() PersonNameComponents
	SetPhoneticRepresentation(value IPersonNameComponents)
}

type PersonNameComponents struct {
	objc.Object
}

func MakePersonNameComponents(ptr unsafe.Pointer) PersonNameComponents {
	return PersonNameComponents{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PersonNameComponentsClass) Alloc() PersonNameComponents {
	rv := objc.CallMethod[PersonNameComponents](pc, "alloc")
	return rv
}

func (pc _PersonNameComponentsClass) New() PersonNameComponents {
	rv := objc.CallMethod[PersonNameComponents](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPersonNameComponents() PersonNameComponents {
	return PersonNameComponentsClass.New()
}

func (p_ PersonNameComponents) Init() PersonNameComponents {
	rv := objc.CallMethod[PersonNameComponents](p_, "init")
	return rv
}

func (p_ PersonNameComponents) NamePrefix() string {
	rv := objc.CallMethod[string](p_, "namePrefix")
	return rv
}

func (p_ PersonNameComponents) SetNamePrefix(value string) {
	objc.CallMethod[objc.Void](p_, "setNamePrefix:", value)
}

func (p_ PersonNameComponents) GivenName() string {
	rv := objc.CallMethod[string](p_, "givenName")
	return rv
}

func (p_ PersonNameComponents) SetGivenName(value string) {
	objc.CallMethod[objc.Void](p_, "setGivenName:", value)
}

func (p_ PersonNameComponents) MiddleName() string {
	rv := objc.CallMethod[string](p_, "middleName")
	return rv
}

func (p_ PersonNameComponents) SetMiddleName(value string) {
	objc.CallMethod[objc.Void](p_, "setMiddleName:", value)
}

func (p_ PersonNameComponents) FamilyName() string {
	rv := objc.CallMethod[string](p_, "familyName")
	return rv
}

func (p_ PersonNameComponents) SetFamilyName(value string) {
	objc.CallMethod[objc.Void](p_, "setFamilyName:", value)
}

func (p_ PersonNameComponents) NameSuffix() string {
	rv := objc.CallMethod[string](p_, "nameSuffix")
	return rv
}

func (p_ PersonNameComponents) SetNameSuffix(value string) {
	objc.CallMethod[objc.Void](p_, "setNameSuffix:", value)
}

func (p_ PersonNameComponents) Nickname() string {
	rv := objc.CallMethod[string](p_, "nickname")
	return rv
}

func (p_ PersonNameComponents) SetNickname(value string) {
	objc.CallMethod[objc.Void](p_, "setNickname:", value)
}

func (p_ PersonNameComponents) PhoneticRepresentation() PersonNameComponents {
	rv := objc.CallMethod[PersonNameComponents](p_, "phoneticRepresentation")
	return rv
}

func (p_ PersonNameComponents) SetPhoneticRepresentation(value IPersonNameComponents) {
	objc.CallMethod[objc.Void](p_, "setPhoneticRepresentation:", value)
}
