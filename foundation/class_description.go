// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ClassDescriptionClass = _ClassDescriptionClass{objc.GetClass("NSClassDescription")}

type _ClassDescriptionClass struct {
	objc.Class
}

type IClassDescription interface {
	objc.IObject
	InverseForRelationshipKey(relationshipKey string) string
	AttributeKeys() []string
	ToManyRelationshipKeys() []string
	ToOneRelationshipKeys() []string
}

type ClassDescription struct {
	objc.Object
}

func MakeClassDescription(ptr unsafe.Pointer) ClassDescription {
	return ClassDescription{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ClassDescriptionClass) Alloc() ClassDescription {
	rv := ffi.CallMethod[ClassDescription](cc, "alloc")
	return rv
}

func (c_ ClassDescription) Init() ClassDescription {
	rv := ffi.CallMethod[ClassDescription](c_, "init")
	return rv
}

func (cc _ClassDescriptionClass) New() ClassDescription {
	rv := ffi.CallMethod[ClassDescription](cc, "new")
	rv.Autorelease()
	return rv
}

func NewClassDescription() ClassDescription {
	return ClassDescriptionClass.New()
}

func (cc _ClassDescriptionClass) InvalidateClassDescriptionCache() {
	ffi.CallMethod[ffi.Void](cc, "invalidateClassDescriptionCache")
}

func (c_ ClassDescription) InverseForRelationshipKey(relationshipKey string) string {
	rv := ffi.CallMethod[string](c_, "inverseForRelationshipKey:", relationshipKey)
	return rv
}

func (c_ ClassDescription) AttributeKeys() []string {
	rv := ffi.CallMethod[[]string](c_, "attributeKeys")
	return rv
}

func (c_ ClassDescription) ToManyRelationshipKeys() []string {
	rv := ffi.CallMethod[[]string](c_, "toManyRelationshipKeys")
	return rv
}

func (c_ ClassDescription) ToOneRelationshipKeys() []string {
	rv := ffi.CallMethod[[]string](c_, "toOneRelationshipKeys")
	return rv
}
