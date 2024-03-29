// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ExceptionClass = _ExceptionClass{objc.GetClass("NSException")}

type _ExceptionClass struct {
	objc.Class
}

type IException interface {
	objc.IObject
	Raise()
	Name() ExceptionName
	Reason() string
	CallStackReturnAddresses() []Number
	CallStackSymbols() []string
}

type Exception struct {
	objc.Object
}

func MakeException(ptr unsafe.Pointer) Exception {
	return Exception{
		Object: objc.MakeObject(ptr),
	}
}

func (ec _ExceptionClass) Alloc() Exception {
	rv := objc.CallMethod[Exception](ec, objc.SEL("alloc"))
	return rv
}

func (ec _ExceptionClass) New() Exception {
	rv := objc.CallMethod[Exception](ec, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewException() Exception {
	return ExceptionClass.New()
}

func (e_ Exception) Init() Exception {
	rv := objc.CallMethod[Exception](e_, objc.SEL("init"))
	return rv
}

func (e_ Exception) Raise() {
	objc.CallMethod[objc.Void](e_, objc.SEL("raise"))
}

func (e_ Exception) Name() ExceptionName {
	rv := objc.CallMethod[ExceptionName](e_, objc.SEL("name"))
	return rv
}

func (e_ Exception) Reason() string {
	rv := objc.CallMethod[string](e_, objc.SEL("reason"))
	return rv
}

func (e_ Exception) CallStackReturnAddresses() []Number {
	rv := objc.CallMethod[[]Number](e_, objc.SEL("callStackReturnAddresses"))
	return rv
}

func (e_ Exception) CallStackSymbols() []string {
	rv := objc.CallMethod[[]string](e_, objc.SEL("callStackSymbols"))
	return rv
}
