// auto-generated code, do not modify
package uniformtypeidentifiers

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TypeClass = _TypeClass{objc.GetClass("UTType")}

type _TypeClass struct {
	objc.Class
}

type IType interface {
	objc.IObject
	ConformsToType(_type IType) bool
	IsSubtypeOfType(_type IType) bool
	IsSupertypeOfType(_type IType) bool
	Identifier() string
	PreferredFilenameExtension() string
	PreferredMIMEType() string
	Tags() map[string][]string
	IsDeclared() bool
	IsDynamic() bool
	IsPublicType() bool
	ReferenceURL() foundation.URL
	Version() foundation.Number
	Supertypes() foundation.Set
	LocalizedDescription() string
}

type Type struct {
	objc.Object
}

func MakeType(ptr unsafe.Pointer) Type {
	return Type{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TypeClass) TypeWithIdentifier(identifier string) Type {
	rv := ffi.CallMethod[Type](tc, "typeWithIdentifier:", identifier)
	return rv
}

func (tc _TypeClass) TypeWithFilenameExtension(filenameExtension string) Type {
	rv := ffi.CallMethod[Type](tc, "typeWithFilenameExtension:", filenameExtension)
	return rv
}

func (tc _TypeClass) TypeWithFilenameExtension_ConformingToType(filenameExtension string, supertype IType) Type {
	rv := ffi.CallMethod[Type](tc, "typeWithFilenameExtension:conformingToType:", filenameExtension, supertype)
	return rv
}

func (tc _TypeClass) TypeWithMIMEType(mimeType string) Type {
	rv := ffi.CallMethod[Type](tc, "typeWithMIMEType:", mimeType)
	return rv
}

func (tc _TypeClass) TypeWithMIMEType_ConformingToType(mimeType string, supertype IType) Type {
	rv := ffi.CallMethod[Type](tc, "typeWithMIMEType:conformingToType:", mimeType, supertype)
	return rv
}

func (tc _TypeClass) TypeWithTag_TagClass_ConformingToType(tag string, tagClass string, supertype IType) Type {
	rv := ffi.CallMethod[Type](tc, "typeWithTag:tagClass:conformingToType:", tag, tagClass, supertype)
	return rv
}

func (tc _TypeClass) Alloc() Type {
	rv := ffi.CallMethod[Type](tc, "alloc")
	return rv
}

func (t_ Type) Init() Type {
	rv := ffi.CallMethod[Type](t_, "init")
	return rv
}

func (tc _TypeClass) New() Type {
	rv := ffi.CallMethod[Type](tc, "new")
	rv.Autorelease()
	return rv
}

func NewType() Type {
	return TypeClass.New()
}

func (tc _TypeClass) TypesWithTag_TagClass_ConformingToType(tag string, tagClass string, supertype IType) []Type {
	rv := ffi.CallMethod[[]Type](tc, "typesWithTag:tagClass:conformingToType:", tag, tagClass, supertype)
	return rv
}

func (tc _TypeClass) ExportedTypeWithIdentifier(identifier string) Type {
	rv := ffi.CallMethod[Type](tc, "exportedTypeWithIdentifier:", identifier)
	return rv
}

func (tc _TypeClass) ExportedTypeWithIdentifier_ConformingToType(identifier string, parentType IType) Type {
	rv := ffi.CallMethod[Type](tc, "exportedTypeWithIdentifier:conformingToType:", identifier, parentType)
	return rv
}

func (tc _TypeClass) ImportedTypeWithIdentifier(identifier string) Type {
	rv := ffi.CallMethod[Type](tc, "importedTypeWithIdentifier:", identifier)
	return rv
}

func (tc _TypeClass) ImportedTypeWithIdentifier_ConformingToType(identifier string, parentType IType) Type {
	rv := ffi.CallMethod[Type](tc, "importedTypeWithIdentifier:conformingToType:", identifier, parentType)
	return rv
}

func (t_ Type) ConformsToType(_type IType) bool {
	rv := ffi.CallMethod[bool](t_, "conformsToType:", _type)
	return rv
}

func (t_ Type) IsSubtypeOfType(_type IType) bool {
	rv := ffi.CallMethod[bool](t_, "isSubtypeOfType:", _type)
	return rv
}

func (t_ Type) IsSupertypeOfType(_type IType) bool {
	rv := ffi.CallMethod[bool](t_, "isSupertypeOfType:", _type)
	return rv
}

func (t_ Type) Identifier() string {
	rv := ffi.CallMethod[string](t_, "identifier")
	return rv
}

func (t_ Type) PreferredFilenameExtension() string {
	rv := ffi.CallMethod[string](t_, "preferredFilenameExtension")
	return rv
}

func (t_ Type) PreferredMIMEType() string {
	rv := ffi.CallMethod[string](t_, "preferredMIMEType")
	return rv
}

func (t_ Type) Tags() map[string][]string {
	rv := ffi.CallMethod[map[string][]string](t_, "tags")
	return rv
}

func (t_ Type) IsDeclared() bool {
	rv := ffi.CallMethod[bool](t_, "isDeclared")
	return rv
}

func (t_ Type) IsDynamic() bool {
	rv := ffi.CallMethod[bool](t_, "isDynamic")
	return rv
}

func (t_ Type) IsPublicType() bool {
	rv := ffi.CallMethod[bool](t_, "isPublicType")
	return rv
}

func (t_ Type) ReferenceURL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](t_, "referenceURL")
	return rv
}

func (t_ Type) Version() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](t_, "version")
	return rv
}

func (t_ Type) Supertypes() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](t_, "supertypes")
	return rv
}

func (t_ Type) LocalizedDescription() string {
	rv := ffi.CallMethod[string](t_, "localizedDescription")
	return rv
}
