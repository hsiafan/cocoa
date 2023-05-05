// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[CloseCommand](c_, objc.GetSelector("initWithCommandDescription:"), objc.ExtractPtr(commandDef))
	return rv
}

func (cc _CloseCommandClass) Alloc() CloseCommand {
	rv := objc.CallMethod[CloseCommand](cc, objc.GetSelector("alloc"))
	return rv
}

func (cc _CloseCommandClass) New() CloseCommand {
	rv := objc.CallMethod[CloseCommand](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCloseCommand() CloseCommand {
	return CloseCommandClass.New()
}

func (c_ CloseCommand) Init() CloseCommand {
	rv := objc.CallMethod[CloseCommand](c_, objc.GetSelector("init"))
	return rv
}

func (c_ CloseCommand) SaveOptions() SaveOptions {
	rv := objc.CallMethod[SaveOptions](c_, objc.GetSelector("saveOptions"))
	return rv
}
