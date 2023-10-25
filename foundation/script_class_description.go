// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[ScriptClassDescription](sc, objc.SEL("alloc"))
	return rv
}

func (sc _ScriptClassDescriptionClass) New() ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewScriptClassDescription() ScriptClassDescription {
	return ScriptClassDescriptionClass.New()
}

func (s_ ScriptClassDescription) Init() ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](s_, objc.SEL("init"))
	return rv
}

func (s_ ScriptClassDescription) ClassDescriptionForKey(key string) ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](s_, objc.SEL("classDescriptionForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isLocationRequiredToCreateForKey:"), toManyRelationshipKey)
	return rv
}

func (s_ ScriptClassDescription) HasOrderedToManyRelationshipForKey(key string) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("hasOrderedToManyRelationshipForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) HasPropertyForKey(key string) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("hasPropertyForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) HasReadablePropertyForKey(key string) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("hasReadablePropertyForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) HasWritablePropertyForKey(key string) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("hasWritablePropertyForKey:"), key)
	return rv
}

// deprecated
func (s_ ScriptClassDescription) IsReadOnlyKey(key string) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isReadOnlyKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) TypeForKey(key string) string {
	rv := objc.CallMethod[string](s_, objc.SEL("typeForKey:"), key)
	return rv
}

func (s_ ScriptClassDescription) SelectorForCommand(commandDescription IScriptCommandDescription) objc.Selector {
	rv := objc.CallMethod[objc.Selector](s_, objc.SEL("selectorForCommand:"), objc.ExtractPtr(commandDescription))
	return rv
}

func (s_ ScriptClassDescription) SupportsCommand(commandDescription IScriptCommandDescription) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("supportsCommand:"), objc.ExtractPtr(commandDescription))
	return rv
}

func (s_ ScriptClassDescription) SuperclassDescription() ScriptClassDescription {
	rv := objc.CallMethod[ScriptClassDescription](s_, objc.SEL("superclassDescription"))
	return rv
}

func (s_ ScriptClassDescription) ClassName() string {
	rv := objc.CallMethod[string](s_, objc.SEL("className"))
	return rv
}

func (s_ ScriptClassDescription) DefaultSubcontainerAttributeKey() string {
	rv := objc.CallMethod[string](s_, objc.SEL("defaultSubcontainerAttributeKey"))
	return rv
}

func (s_ ScriptClassDescription) ImplementationClassName() string {
	rv := objc.CallMethod[string](s_, objc.SEL("implementationClassName"))
	return rv
}

func (s_ ScriptClassDescription) SuiteName() string {
	rv := objc.CallMethod[string](s_, objc.SEL("suiteName"))
	return rv
}
