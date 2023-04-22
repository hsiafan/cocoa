// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

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
	rv := objc.CallMethod[ValueFunction](vc, objc.GetSelector("functionWithName:"), name)
	return rv
}

func (vc _ValueFunctionClass) Alloc() ValueFunction {
	rv := objc.CallMethod[ValueFunction](vc, objc.GetSelector("alloc"))
	return rv
}

func (vc _ValueFunctionClass) New() ValueFunction {
	rv := objc.CallMethod[ValueFunction](vc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewValueFunction() ValueFunction {
	return ValueFunctionClass.New()
}

func (v_ ValueFunction) Init() ValueFunction {
	rv := objc.CallMethod[ValueFunction](v_, objc.GetSelector("init"))
	return rv
}

func (v_ ValueFunction) Name() ValueFunctionName {
	rv := objc.CallMethod[ValueFunctionName](v_, objc.GetSelector("name"))
	return rv
}
