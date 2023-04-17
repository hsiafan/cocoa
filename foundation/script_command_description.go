// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[ScriptCommandDescription](sc, "alloc")
	return rv
}

func (sc _ScriptCommandDescriptionClass) New() ScriptCommandDescription {
	rv := ffi.CallMethod[ScriptCommandDescription](sc, "new")
	rv.Autorelease()
	return rv
}

func NewScriptCommandDescription() ScriptCommandDescription {
	return ScriptCommandDescriptionClass.New()
}

func (s_ ScriptCommandDescription) Init() ScriptCommandDescription {
	rv := ffi.CallMethod[ScriptCommandDescription](s_, "init")
	return rv
}

func (s_ ScriptCommandDescription) IsOptionalArgumentWithName(argumentName string) bool {
	rv := ffi.CallMethod[bool](s_, "isOptionalArgumentWithName:", argumentName)
	return rv
}

func (s_ ScriptCommandDescription) TypeForArgumentWithName(argumentName string) string {
	rv := ffi.CallMethod[string](s_, "typeForArgumentWithName:", argumentName)
	return rv
}

func (s_ ScriptCommandDescription) CreateCommandInstance() ScriptCommand {
	rv := ffi.CallMethod[ScriptCommand](s_, "createCommandInstance")
	return rv
}

func (s_ ScriptCommandDescription) CommandClassName() string {
	rv := ffi.CallMethod[string](s_, "commandClassName")
	return rv
}

func (s_ ScriptCommandDescription) CommandName() string {
	rv := ffi.CallMethod[string](s_, "commandName")
	return rv
}

func (s_ ScriptCommandDescription) SuiteName() string {
	rv := ffi.CallMethod[string](s_, "suiteName")
	return rv
}

func (s_ ScriptCommandDescription) ArgumentNames() []string {
	rv := ffi.CallMethod[[]string](s_, "argumentNames")
	return rv
}

func (s_ ScriptCommandDescription) ReturnType() string {
	rv := ffi.CallMethod[string](s_, "returnType")
	return rv
}
