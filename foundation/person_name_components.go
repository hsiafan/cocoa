// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[PersonNameComponents](pc, "alloc")
	return rv
}

func (pc _PersonNameComponentsClass) New() PersonNameComponents {
	rv := ffi.CallMethod[PersonNameComponents](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPersonNameComponents() PersonNameComponents {
	return PersonNameComponentsClass.New()
}

func (p_ PersonNameComponents) Init() PersonNameComponents {
	rv := ffi.CallMethod[PersonNameComponents](p_, "init")
	return rv
}

func (p_ PersonNameComponents) NamePrefix() string {
	rv := ffi.CallMethod[string](p_, "namePrefix")
	return rv
}

func (p_ PersonNameComponents) SetNamePrefix(value string) {
	ffi.CallMethod[ffi.Void](p_, "setNamePrefix:", value)
}

func (p_ PersonNameComponents) GivenName() string {
	rv := ffi.CallMethod[string](p_, "givenName")
	return rv
}

func (p_ PersonNameComponents) SetGivenName(value string) {
	ffi.CallMethod[ffi.Void](p_, "setGivenName:", value)
}

func (p_ PersonNameComponents) MiddleName() string {
	rv := ffi.CallMethod[string](p_, "middleName")
	return rv
}

func (p_ PersonNameComponents) SetMiddleName(value string) {
	ffi.CallMethod[ffi.Void](p_, "setMiddleName:", value)
}

func (p_ PersonNameComponents) FamilyName() string {
	rv := ffi.CallMethod[string](p_, "familyName")
	return rv
}

func (p_ PersonNameComponents) SetFamilyName(value string) {
	ffi.CallMethod[ffi.Void](p_, "setFamilyName:", value)
}

func (p_ PersonNameComponents) NameSuffix() string {
	rv := ffi.CallMethod[string](p_, "nameSuffix")
	return rv
}

func (p_ PersonNameComponents) SetNameSuffix(value string) {
	ffi.CallMethod[ffi.Void](p_, "setNameSuffix:", value)
}

func (p_ PersonNameComponents) Nickname() string {
	rv := ffi.CallMethod[string](p_, "nickname")
	return rv
}

func (p_ PersonNameComponents) SetNickname(value string) {
	ffi.CallMethod[ffi.Void](p_, "setNickname:", value)
}

func (p_ PersonNameComponents) PhoneticRepresentation() PersonNameComponents {
	rv := ffi.CallMethod[PersonNameComponents](p_, "phoneticRepresentation")
	return rv
}

func (p_ PersonNameComponents) SetPhoneticRepresentation(value IPersonNameComponents) {
	ffi.CallMethod[ffi.Void](p_, "setPhoneticRepresentation:", value)
}
