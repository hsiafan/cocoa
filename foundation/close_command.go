// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

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

func (cc _CloseCommandClass) New() CloseCommand {
	rv := ffi.CallMethod[CloseCommand](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCloseCommand() CloseCommand {
	return CloseCommandClass.New()
}

func (c_ CloseCommand) Init() CloseCommand {
	rv := ffi.CallMethod[CloseCommand](c_, "init")
	return rv
}

func (c_ CloseCommand) SaveOptions() SaveOptions {
	rv := ffi.CallMethod[SaveOptions](c_, "saveOptions")
	return rv
}