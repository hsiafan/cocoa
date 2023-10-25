// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	rv := objc.CallMethod[Unit](u_, objc.SEL("initWithSymbol:"), symbol)
	return rv
}

func (uc _UnitClass) Alloc() Unit {
	rv := objc.CallMethod[Unit](uc, objc.SEL("alloc"))
	return rv
}

func (uc _UnitClass) New() Unit {
	rv := objc.CallMethod[Unit](uc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewUnit() Unit {
	return UnitClass.New()
}

func (u_ Unit) Init() Unit {
	rv := objc.CallMethod[Unit](u_, objc.SEL("init"))
	return rv
}

func (u_ Unit) Symbol() string {
	rv := objc.CallMethod[string](u_, objc.SEL("symbol"))
	return rv
}
