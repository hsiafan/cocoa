// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var MediaTimingFunctionClass = _MediaTimingFunctionClass{objc.GetClass("CAMediaTimingFunction")}

type _MediaTimingFunctionClass struct {
	objc.Class
}

type IMediaTimingFunction interface {
	objc.IObject
}

type MediaTimingFunction struct {
	objc.Object
}

func MakeMediaTimingFunction(ptr unsafe.Pointer) MediaTimingFunction {
	return MediaTimingFunction{
		Object: objc.MakeObject(ptr),
	}
}

func (mc _MediaTimingFunctionClass) Alloc() MediaTimingFunction {
	rv := ffi.CallMethod[MediaTimingFunction](mc, "alloc")
	return rv
}

func (mc _MediaTimingFunctionClass) New() MediaTimingFunction {
	rv := ffi.CallMethod[MediaTimingFunction](mc, "new")
	rv.Autorelease()
	return rv
}

func NewMediaTimingFunction() MediaTimingFunction {
	return MediaTimingFunctionClass.New()
}

func (m_ MediaTimingFunction) Init() MediaTimingFunction {
	rv := ffi.CallMethod[MediaTimingFunction](m_, "init")
	return rv
}