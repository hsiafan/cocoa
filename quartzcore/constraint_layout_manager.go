// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ConstraintLayoutManagerClass = _ConstraintLayoutManagerClass{objc.GetClass("CAConstraintLayoutManager")}

type _ConstraintLayoutManagerClass struct {
	objc.Class
}

type IConstraintLayoutManager interface {
	objc.IObject
}

type ConstraintLayoutManager struct {
	objc.Object
}

func MakeConstraintLayoutManager(ptr unsafe.Pointer) ConstraintLayoutManager {
	return ConstraintLayoutManager{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ConstraintLayoutManagerClass) LayoutManager() ConstraintLayoutManager {
	rv := objc.CallMethod[ConstraintLayoutManager](cc, "layoutManager")
	return rv
}

func (cc _ConstraintLayoutManagerClass) Alloc() ConstraintLayoutManager {
	rv := objc.CallMethod[ConstraintLayoutManager](cc, "alloc")
	return rv
}

func (cc _ConstraintLayoutManagerClass) New() ConstraintLayoutManager {
	rv := objc.CallMethod[ConstraintLayoutManager](cc, "new")
	rv.Autorelease()
	return rv
}

func NewConstraintLayoutManager() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.New()
}

func (c_ ConstraintLayoutManager) Init() ConstraintLayoutManager {
	rv := objc.CallMethod[ConstraintLayoutManager](c_, "init")
	return rv
}
