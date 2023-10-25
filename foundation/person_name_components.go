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
	rv := objc.CallMethod[PersonNameComponents](pc, objc.SEL("alloc"))
	return rv
}

func (pc _PersonNameComponentsClass) New() PersonNameComponents {
	rv := objc.CallMethod[PersonNameComponents](pc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewPersonNameComponents() PersonNameComponents {
	return PersonNameComponentsClass.New()
}

func (p_ PersonNameComponents) Init() PersonNameComponents {
	rv := objc.CallMethod[PersonNameComponents](p_, objc.SEL("init"))
	return rv
}

func (p_ PersonNameComponents) NamePrefix() string {
	rv := objc.CallMethod[string](p_, objc.SEL("namePrefix"))
	return rv
}

func (p_ PersonNameComponents) SetNamePrefix(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setNamePrefix:"), value)
}

func (p_ PersonNameComponents) GivenName() string {
	rv := objc.CallMethod[string](p_, objc.SEL("givenName"))
	return rv
}

func (p_ PersonNameComponents) SetGivenName(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setGivenName:"), value)
}

func (p_ PersonNameComponents) MiddleName() string {
	rv := objc.CallMethod[string](p_, objc.SEL("middleName"))
	return rv
}

func (p_ PersonNameComponents) SetMiddleName(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setMiddleName:"), value)
}

func (p_ PersonNameComponents) FamilyName() string {
	rv := objc.CallMethod[string](p_, objc.SEL("familyName"))
	return rv
}

func (p_ PersonNameComponents) SetFamilyName(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setFamilyName:"), value)
}

func (p_ PersonNameComponents) NameSuffix() string {
	rv := objc.CallMethod[string](p_, objc.SEL("nameSuffix"))
	return rv
}

func (p_ PersonNameComponents) SetNameSuffix(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setNameSuffix:"), value)
}

func (p_ PersonNameComponents) Nickname() string {
	rv := objc.CallMethod[string](p_, objc.SEL("nickname"))
	return rv
}

func (p_ PersonNameComponents) SetNickname(value string) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setNickname:"), value)
}

func (p_ PersonNameComponents) PhoneticRepresentation() PersonNameComponents {
	rv := objc.CallMethod[PersonNameComponents](p_, objc.SEL("phoneticRepresentation"))
	return rv
}

func (p_ PersonNameComponents) SetPhoneticRepresentation(value IPersonNameComponents) {
	objc.CallMethod[objc.Void](p_, objc.SEL("setPhoneticRepresentation:"), objc.ExtractPtr(value))
}
