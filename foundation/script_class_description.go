// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ScriptClassDescriptionClass = _ScriptClassDescriptionClass{objc.GetClass("NSScriptClassDescription")}

type _ScriptClassDescriptionClass struct {
	objc.Class
}

type IScriptClassDescription interface {
	IClassDescription
	ClassDescriptionForKey(key string) ScriptClassDescription
	IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool
	HasOrderedToManyRelationshipForKey(key string) bool
	HasPropertyForKey(key string) bool
	HasReadablePropertyForKey(key string) bool
	HasWritablePropertyForKey(key string) bool
	// deprecated
	IsReadOnlyKey(key string) bool
	TypeForKey(key string) string
	SelectorForCommand(commandDescription IScriptCommandDescription) objc.Selector
	SupportsCommand(commandDescription IScriptCommandDescription) bool
	SuperclassDescription() ScriptClassDescription
	ClassName() string
	DefaultSubcontainerAttributeKey() string
	ImplementationClassName() string
	SuiteName() string
}

type ScriptClassDescription struct {
	ClassDescription
}

func MakeScriptClassDescription(ptr unsafe.Pointer) ScriptClassDescription {
	return ScriptClassDescription{
		ClassDescription: MakeClassDescription(ptr),
	}
}

func (sc _ScriptClassDescriptionClass) Alloc() ScriptClassDescription {
	rv := ffi.CallMethod[ScriptClassDescription](sc, "alloc")
	return rv
}

func (sc _ScriptClassDescriptionClass) New() ScriptClassDescription {
	rv := ffi.CallMethod[ScriptClassDescription](sc, "new")
	rv.Autorelease()
	return rv
}

func NewScriptClassDescription() ScriptClassDescription {
	return ScriptClassDescriptionClass.New()
}

func (s_ ScriptClassDescription) Init() ScriptClassDescription {
	rv := ffi.CallMethod[ScriptClassDescription](s_, "init")
	return rv
}

func (s_ ScriptClassDescription) ClassDescriptionForKey(key string) ScriptClassDescription {
	rv := ffi.CallMethod[ScriptClassDescription](s_, "classDescriptionForKey:", key)
	return rv
}

func (s_ ScriptClassDescription) IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool {
	rv := ffi.CallMethod[bool](s_, "isLocationRequiredToCreateForKey:", toManyRelationshipKey)
	return rv
}

func (s_ ScriptClassDescription) HasOrderedToManyRelationshipForKey(key string) bool {
	rv := ffi.CallMethod[bool](s_, "hasOrderedToManyRelationshipForKey:", key)
	return rv
}

func (s_ ScriptClassDescription) HasPropertyForKey(key string) bool {
	rv := ffi.CallMethod[bool](s_, "hasPropertyForKey:", key)
	return rv
}

func (s_ ScriptClassDescription) HasReadablePropertyForKey(key string) bool {
	rv := ffi.CallMethod[bool](s_, "hasReadablePropertyForKey:", key)
	return rv
}

func (s_ ScriptClassDescription) HasWritablePropertyForKey(key string) bool {
	rv := ffi.CallMethod[bool](s_, "hasWritablePropertyForKey:", key)
	return rv
}

// deprecated
func (s_ ScriptClassDescription) IsReadOnlyKey(key string) bool {
	rv := ffi.CallMethod[bool](s_, "isReadOnlyKey:", key)
	return rv
}

func (s_ ScriptClassDescription) TypeForKey(key string) string {
	rv := ffi.CallMethod[string](s_, "typeForKey:", key)
	return rv
}

func (s_ ScriptClassDescription) SelectorForCommand(commandDescription IScriptCommandDescription) objc.Selector {
	rv := ffi.CallMethod[objc.Selector](s_, "selectorForCommand:", commandDescription)
	return rv
}

func (s_ ScriptClassDescription) SupportsCommand(commandDescription IScriptCommandDescription) bool {
	rv := ffi.CallMethod[bool](s_, "supportsCommand:", commandDescription)
	return rv
}

func (s_ ScriptClassDescription) SuperclassDescription() ScriptClassDescription {
	rv := ffi.CallMethod[ScriptClassDescription](s_, "superclassDescription")
	return rv
}

func (s_ ScriptClassDescription) ClassName() string {
	rv := ffi.CallMethod[string](s_, "className")
	return rv
}

func (s_ ScriptClassDescription) DefaultSubcontainerAttributeKey() string {
	rv := ffi.CallMethod[string](s_, "defaultSubcontainerAttributeKey")
	return rv
}

func (s_ ScriptClassDescription) ImplementationClassName() string {
	rv := ffi.CallMethod[string](s_, "implementationClassName")
	return rv
}

func (s_ ScriptClassDescription) SuiteName() string {
	rv := ffi.CallMethod[string](s_, "suiteName")
	return rv
}
