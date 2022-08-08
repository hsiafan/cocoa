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

func (s_ ScriptCommand) Init() ScriptCommand {
	rv := ffi.CallMethod[ScriptCommand](s_, "init")
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

var ScriptCommandDescriptionClass = _ScriptCommandDescriptionClass{objc.GetClass("NSScriptCommandDescription")}

type _ScriptCommandDescriptionClass struct {
	objc.Class
}

type IScriptCommandDescription interface {
	objc.IObject
	AppleEventCodeForArgumentWithName(argumentName string) uint32
	IsOptionalArgumentWithName(argumentName string) bool
	TypeForArgumentWithName(argumentName string) string
	CreateCommandInstance() ScriptCommand
	AppleEventClassCode() uint32
	AppleEventCode() uint32
	CommandClassName() string
	CommandName() string
	SuiteName() string
	ArgumentNames() []string
	AppleEventCodeForReturnType() uint32
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

func (s_ ScriptCommandDescription) Init() ScriptCommandDescription {
	rv := ffi.CallMethod[ScriptCommandDescription](s_, "init")
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

func (s_ ScriptCommandDescription) AppleEventCodeForArgumentWithName(argumentName string) uint32 {
	rv := ffi.CallMethod[uint32](s_, "appleEventCodeForArgumentWithName:", argumentName)
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

func (s_ ScriptCommandDescription) AppleEventClassCode() uint32 {
	rv := ffi.CallMethod[uint32](s_, "appleEventClassCode")
	return rv
}

func (s_ ScriptCommandDescription) AppleEventCode() uint32 {
	rv := ffi.CallMethod[uint32](s_, "appleEventCode")
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

func (s_ ScriptCommandDescription) AppleEventCodeForReturnType() uint32 {
	rv := ffi.CallMethod[uint32](s_, "appleEventCodeForReturnType")
	return rv
}

func (s_ ScriptCommandDescription) ReturnType() string {
	rv := ffi.CallMethod[string](s_, "returnType")
	return rv
}

var CloseCommandClass = _CloseCommandClass{objc.GetClass("NSCloseCommand")}

type _CloseCommandClass struct {
	objc.Class
}

type ICloseCommand interface {
	IScriptCommand
	SaveOptions() SaveOptions
}

type CloseCommand struct {
	ScriptCommand
}

func MakeCloseCommand(ptr unsafe.Pointer) CloseCommand {
	return CloseCommand{
		ScriptCommand: MakeScriptCommand(ptr),
	}
}

func (c_ CloseCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) CloseCommand {
	rv := ffi.CallMethod[CloseCommand](c_, "initWithCommandDescription:", commandDef)
	return rv
}

func (cc _CloseCommandClass) Alloc() CloseCommand {
	rv := ffi.CallMethod[CloseCommand](cc, "alloc")
	return rv
}

func (c_ CloseCommand) Init() CloseCommand {
	rv := ffi.CallMethod[CloseCommand](c_, "init")
	return rv
}

func (cc _CloseCommandClass) New() CloseCommand {
	rv := ffi.CallMethod[CloseCommand](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCloseCommand() CloseCommand {
	return CloseCommandClass.New()
}

func (c_ CloseCommand) SaveOptions() SaveOptions {
	rv := ffi.CallMethod[SaveOptions](c_, "saveOptions")
	return rv
}
