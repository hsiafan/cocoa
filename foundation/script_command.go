// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ScriptCommandClass = _ScriptCommandClass{objc.GetClass("NSScriptCommand")}

type _ScriptCommandClass struct {
	objc.Class
}

type IScriptCommand interface {
	objc.IObject
	ExecuteCommand() objc.Object
	PerformDefaultImplementation() objc.Object
	SuspendExecution()
	ResumeExecutionWithResult(result objc.IObject)
	AppleEvent() AppleEventDescriptor
	EvaluatedReceivers() objc.Object
	ReceiversSpecifier() ScriptObjectSpecifier
	SetReceiversSpecifier(value IScriptObjectSpecifier)
	Arguments() map[string]objc.Object
	SetArguments(value map[string]objc.IObject)
	EvaluatedArguments() map[string]objc.Object
	DirectParameter() objc.Object
	SetDirectParameter(value objc.IObject)
	CommandDescription() ScriptCommandDescription
	ScriptErrorExpectedTypeDescriptor() AppleEventDescriptor
	SetScriptErrorExpectedTypeDescriptor(value IAppleEventDescriptor)
	ScriptErrorNumber() int
	SetScriptErrorNumber(value int)
	ScriptErrorOffendingObjectDescriptor() AppleEventDescriptor
	SetScriptErrorOffendingObjectDescriptor(value IAppleEventDescriptor)
	ScriptErrorString() string
	SetScriptErrorString(value string)
	IsWellFormed() bool
}

type ScriptCommand struct {
	objc.Object
}

func MakeScriptCommand(ptr unsafe.Pointer) ScriptCommand {
	return ScriptCommand{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ ScriptCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) ScriptCommand {
	rv := ffi.CallMethod[ScriptCommand](s_, "initWithCommandDescription:", commandDef)
	return rv
}

func (sc _ScriptCommandClass) Alloc() ScriptCommand {
	rv := ffi.CallMethod[ScriptCommand](sc, "alloc")
	return rv
}

func (sc _ScriptCommandClass) New() ScriptCommand {
	rv := ffi.CallMethod[ScriptCommand](sc, "new")
	rv.Autorelease()
	return rv
}

func NewScriptCommand() ScriptCommand {
	return ScriptCommandClass.New()
}

func (s_ ScriptCommand) Init() ScriptCommand {
	rv := ffi.CallMethod[ScriptCommand](s_, "init")
	return rv
}

func (sc _ScriptCommandClass) CurrentCommand() ScriptCommand {
	rv := ffi.CallMethod[ScriptCommand](sc, "currentCommand")
	return rv
}

func (s_ ScriptCommand) ExecuteCommand() objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "executeCommand")
	return rv
}

func (s_ ScriptCommand) PerformDefaultImplementation() objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "performDefaultImplementation")
	return rv
}

func (s_ ScriptCommand) SuspendExecution() {
	ffi.CallMethod[ffi.Void](s_, "suspendExecution")
}

func (s_ ScriptCommand) ResumeExecutionWithResult(result objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "resumeExecutionWithResult:", result)
}

func (s_ ScriptCommand) AppleEvent() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](s_, "appleEvent")
	return rv
}

func (s_ ScriptCommand) EvaluatedReceivers() objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "evaluatedReceivers")
	return rv
}

func (s_ ScriptCommand) ReceiversSpecifier() ScriptObjectSpecifier {
	rv := ffi.CallMethod[ScriptObjectSpecifier](s_, "receiversSpecifier")
	return rv
}

func (s_ ScriptCommand) SetReceiversSpecifier(value IScriptObjectSpecifier) {
	ffi.CallMethod[ffi.Void](s_, "setReceiversSpecifier:", value)
}

func (s_ ScriptCommand) Arguments() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](s_, "arguments")
	return rv
}

func (s_ ScriptCommand) SetArguments(value map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "setArguments:", value)
}

func (s_ ScriptCommand) EvaluatedArguments() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](s_, "evaluatedArguments")
	return rv
}

func (s_ ScriptCommand) DirectParameter() objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "directParameter")
	return rv
}

func (s_ ScriptCommand) SetDirectParameter(value objc.IObject) {
	ffi.CallMethod[ffi.Void](s_, "setDirectParameter:", value)
}

func (s_ ScriptCommand) CommandDescription() ScriptCommandDescription {
	rv := ffi.CallMethod[ScriptCommandDescription](s_, "commandDescription")
	return rv
}

func (s_ ScriptCommand) ScriptErrorExpectedTypeDescriptor() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](s_, "scriptErrorExpectedTypeDescriptor")
	return rv
}

func (s_ ScriptCommand) SetScriptErrorExpectedTypeDescriptor(value IAppleEventDescriptor) {
	ffi.CallMethod[ffi.Void](s_, "setScriptErrorExpectedTypeDescriptor:", value)
}

func (s_ ScriptCommand) ScriptErrorNumber() int {
	rv := ffi.CallMethod[int](s_, "scriptErrorNumber")
	return rv
}

func (s_ ScriptCommand) SetScriptErrorNumber(value int) {
	ffi.CallMethod[ffi.Void](s_, "setScriptErrorNumber:", value)
}

func (s_ ScriptCommand) ScriptErrorOffendingObjectDescriptor() AppleEventDescriptor {
	rv := ffi.CallMethod[AppleEventDescriptor](s_, "scriptErrorOffendingObjectDescriptor")
	return rv
}

func (s_ ScriptCommand) SetScriptErrorOffendingObjectDescriptor(value IAppleEventDescriptor) {
	ffi.CallMethod[ffi.Void](s_, "setScriptErrorOffendingObjectDescriptor:", value)
}

func (s_ ScriptCommand) ScriptErrorString() string {
	rv := ffi.CallMethod[string](s_, "scriptErrorString")
	return rv
}

func (s_ ScriptCommand) SetScriptErrorString(value string) {
	ffi.CallMethod[ffi.Void](s_, "setScriptErrorString:", value)
}

func (s_ ScriptCommand) IsWellFormed() bool {
	rv := ffi.CallMethod[bool](s_, "isWellFormed")
	return rv
}
