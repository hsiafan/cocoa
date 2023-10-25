// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[ScriptCommand](s_, objc.SEL("initWithCommandDescription:"), objc.ExtractPtr(commandDef))
	return rv
}

func (sc _ScriptCommandClass) Alloc() ScriptCommand {
	rv := objc.CallMethod[ScriptCommand](sc, objc.SEL("alloc"))
	return rv
}

func (sc _ScriptCommandClass) New() ScriptCommand {
	rv := objc.CallMethod[ScriptCommand](sc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewScriptCommand() ScriptCommand {
	return ScriptCommandClass.New()
}

func (s_ ScriptCommand) Init() ScriptCommand {
	rv := objc.CallMethod[ScriptCommand](s_, objc.SEL("init"))
	return rv
}

func (sc _ScriptCommandClass) CurrentCommand() ScriptCommand {
	rv := objc.CallMethod[ScriptCommand](sc, objc.SEL("currentCommand"))
	return rv
}

func (s_ ScriptCommand) ExecuteCommand() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("executeCommand"))
	return rv
}

func (s_ ScriptCommand) PerformDefaultImplementation() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("performDefaultImplementation"))
	return rv
}

func (s_ ScriptCommand) SuspendExecution() {
	objc.CallMethod[objc.Void](s_, objc.SEL("suspendExecution"))
}

func (s_ ScriptCommand) ResumeExecutionWithResult(result objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("resumeExecutionWithResult:"), objc.ExtractPtr(result))
}

func (s_ ScriptCommand) AppleEvent() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](s_, objc.SEL("appleEvent"))
	return rv
}

func (s_ ScriptCommand) EvaluatedReceivers() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("evaluatedReceivers"))
	return rv
}

func (s_ ScriptCommand) ReceiversSpecifier() ScriptObjectSpecifier {
	rv := objc.CallMethod[ScriptObjectSpecifier](s_, objc.SEL("receiversSpecifier"))
	return rv
}

func (s_ ScriptCommand) SetReceiversSpecifier(value IScriptObjectSpecifier) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setReceiversSpecifier:"), objc.ExtractPtr(value))
}

func (s_ ScriptCommand) Arguments() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](s_, objc.SEL("arguments"))
	return rv
}

func (s_ ScriptCommand) SetArguments(value map[string]objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setArguments:"), value)
}

func (s_ ScriptCommand) EvaluatedArguments() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](s_, objc.SEL("evaluatedArguments"))
	return rv
}

func (s_ ScriptCommand) DirectParameter() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.SEL("directParameter"))
	return rv
}

func (s_ ScriptCommand) SetDirectParameter(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setDirectParameter:"), objc.ExtractPtr(value))
}

func (s_ ScriptCommand) CommandDescription() ScriptCommandDescription {
	rv := objc.CallMethod[ScriptCommandDescription](s_, objc.SEL("commandDescription"))
	return rv
}

func (s_ ScriptCommand) ScriptErrorExpectedTypeDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](s_, objc.SEL("scriptErrorExpectedTypeDescriptor"))
	return rv
}

func (s_ ScriptCommand) SetScriptErrorExpectedTypeDescriptor(value IAppleEventDescriptor) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setScriptErrorExpectedTypeDescriptor:"), objc.ExtractPtr(value))
}

func (s_ ScriptCommand) ScriptErrorNumber() int {
	rv := objc.CallMethod[int](s_, objc.SEL("scriptErrorNumber"))
	return rv
}

func (s_ ScriptCommand) SetScriptErrorNumber(value int) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setScriptErrorNumber:"), value)
}

func (s_ ScriptCommand) ScriptErrorOffendingObjectDescriptor() AppleEventDescriptor {
	rv := objc.CallMethod[AppleEventDescriptor](s_, objc.SEL("scriptErrorOffendingObjectDescriptor"))
	return rv
}

func (s_ ScriptCommand) SetScriptErrorOffendingObjectDescriptor(value IAppleEventDescriptor) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setScriptErrorOffendingObjectDescriptor:"), objc.ExtractPtr(value))
}

func (s_ ScriptCommand) ScriptErrorString() string {
	rv := objc.CallMethod[string](s_, objc.SEL("scriptErrorString"))
	return rv
}

func (s_ ScriptCommand) SetScriptErrorString(value string) {
	objc.CallMethod[objc.Void](s_, objc.SEL("setScriptErrorString:"), value)
}

func (s_ ScriptCommand) IsWellFormed() bool {
	rv := objc.CallMethod[bool](s_, objc.SEL("isWellFormed"))
	return rv
}
