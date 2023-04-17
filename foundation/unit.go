// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var UnitClass = _UnitClass{objc.GetClass("NSUnit")}

type _UnitClass struct {
	objc.Class
}

type IUnit interface {
	objc.IObject
	Symbol() string
}

type Unit struct {
	objc.Object
}

func MakeUnit(ptr unsafe.Pointer) Unit {
	return Unit{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ Unit) InitWithSymbol(symbol string) Unit {
	rv := ffi.CallMethod[Unit](u_, "initWithSymbol:", symbol)
	return rv
}

func (uc _UnitClass) Alloc() Unit {
	rv := ffi.CallMethod[Unit](uc, "alloc")
	return rv
}

func (uc _UnitClass) New() Unit {
	rv := ffi.CallMethod[Unit](uc, "new")
	rv.Autorelease()
	return rv
}

func NewUnit() Unit {
	return UnitClass.New()
}

func (u_ Unit) Init() Unit {
	rv := ffi.CallMethod[Unit](u_, "init")
	return rv
}

func (u_ Unit) Symbol() string {
	rv := ffi.CallMethod[string](u_, "symbol")
	return rv
}
