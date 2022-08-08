// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var ValueFunctionClass = _ValueFunctionClass{objc.GetClass("CAValueFunction")}

type _ValueFunctionClass struct {
	objc.Class
}

type IValueFunction interface {
	objc.IObject
	Name() ValueFunctionName
}

type ValueFunction struct {
	objc.Object
}

func MakeValueFunction(ptr unsafe.Pointer) ValueFunction {
	return ValueFunction{
		Object: objc.MakeObject(ptr),
	}
}

func (vc _ValueFunctionClass) FunctionWithName(name ValueFunctionName) ValueFunction {
	rv := ffi.CallMethod[ValueFunction](vc, "functionWithName:", name)
	return rv
}

func (vc _ValueFunctionClass) Alloc() ValueFunction {
	rv := ffi.CallMethod[ValueFunction](vc, "alloc")
	return rv
}

func (v_ ValueFunction) Init() ValueFunction {
	rv := ffi.CallMethod[ValueFunction](v_, "init")
	return rv
}

func (vc _ValueFunctionClass) New() ValueFunction {
	rv := ffi.CallMethod[ValueFunction](vc, "new")
	rv.Autorelease()
	return rv
}

func NewValueFunction() ValueFunction {
	return ValueFunctionClass.New()
}

func (v_ ValueFunction) Name() ValueFunctionName {
	rv := ffi.CallMethod[ValueFunctionName](v_, "name")
	return rv
}
