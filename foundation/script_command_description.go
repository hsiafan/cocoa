// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ScriptCommandDescriptionClass = _ScriptCommandDescriptionClass{objc.GetClass("NSScriptCommandDescription")}

type _ScriptCommandDescriptionClass struct {
	objc.Class
}

type IScriptCommandDescription interface {
	objc.IObject
	IsOptionalArgumentWithName(argumentName string) bool
	TypeForArgumentWithName(argumentName string) string
	CreateCommandInstance() ScriptCommand
	CommandClassName() string
	CommandName() string
	SuiteName() string
	ArgumentNames() []string
	ReturnType() string
}

type ScriptCommandDescription struct {
	objc.Object
}

func MakeScriptCommandDescription(ptr unsafe.Pointer) ScriptCommandDescription {
	return ScriptCommandDescription{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _ScriptCommandDescriptionClass) Alloc() ScriptCommandDescription {
	rv := objc.CallMethod[ScriptCommandDescription](sc, objc.SEL("alloc"))
	return rv
}

func (sc _ScriptCommandDescriptionClass) New() ScriptCommandDescription {
	rv := objc.CallMethod[ScriptCommandDescription](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewScriptCommandDescription() ScriptCommandDescription {
	return ScriptCommandDescriptionClass.New()
}

func (s_ ScriptCommandDescription) Init() ScriptCommandDescription {
	rv := objc.CallMethod[ScriptCommandDescription](s_, objc.SEL("init"))
	return rv
}

func (s_ ScriptCommandDescription) IsOptionalArgumentWithName(argumentName string) bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isOptionalArgumentWithName:"), argumentName)
	return rv
}

func (s_ ScriptCommandDescription) TypeForArgumentWithName(argumentName string) string {
	rv := objc.CallMethod[string](s_, objc.SEL("typeForArgumentWithName:"), argumentName)
	return rv
}

func (s_ ScriptCommandDescription) CreateCommandInstance() ScriptCommand {
	rv := objc.CallMethod[ScriptCommand](s_, objc.SEL("createCommandInstance"))
	return rv
}

func (s_ ScriptCommandDescription) CommandClassName() string {
	rv := objc.CallMethod[string](s_, objc.SEL("commandClassName"))
	return rv
}

func (s_ ScriptCommandDescription) CommandName() string {
	rv := objc.CallMethod[string](s_, objc.SEL("commandName"))
	return rv
}

func (s_ ScriptCommandDescription) SuiteName() string {
	rv := objc.CallMethod[string](s_, objc.SEL("suiteName"))
	return rv
}

func (s_ ScriptCommandDescription) ArgumentNames() []string {
	rv := objc.CallMethod[[]string](s_, objc.SEL("argumentNames"))
	return rv
}

func (s_ ScriptCommandDescription) ReturnType() string {
	rv := objc.CallMethod[string](s_, objc.SEL("returnType"))
	return rv
}
