// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var NibClass = _NibClass{objc.GetClass("NSNib")}

type _NibClass struct {
	objc.Class
}

type INib interface {
	objc.IObject
	// deprecated
	InitWithContentsOfURL(nibFileURL foundation.IURL) objc.Object
	InstantiateWithOwner_TopLevelObjects(owner objc.IObject, topLevelObjects *foundation.Array) bool
}

type Nib struct {
	objc.Object
}

func MakeNib(ptr unsafe.Pointer) Nib {
	return Nib{
		Object: objc.MakeObject(ptr),
	}
}

func (n_ Nib) InitWithNibNamed_Bundle(nibName NibName, bundle foundation.IBundle) Nib {
	rv := ffi.CallMethod[Nib](n_, "initWithNibNamed:bundle:", nibName, bundle)
	rv.Autorelease()
	return rv
}

func (n_ Nib) InitWithNibData_Bundle(nibData []byte, bundle foundation.IBundle) Nib {
	rv := ffi.CallMethod[Nib](n_, "initWithNibData:bundle:", nibData, bundle)
	rv.Autorelease()
	return rv
}

func (nc _NibClass) Alloc() Nib {
	rv := ffi.CallMethod[Nib](nc, "alloc")
	return rv
}

func (n_ Nib) Init() Nib {
	rv := ffi.CallMethod[Nib](n_, "init")
	rv.Autorelease()
	return rv
}

func (nc _NibClass) New() Nib {
	rv := ffi.CallMethod[Nib](nc, "new")
	rv.Autorelease()
	return rv
}

func NewNib() Nib {
	return NibClass.New()
}

// deprecated
func (n_ Nib) InitWithContentsOfURL(nibFileURL foundation.IURL) objc.Object {
	rv := ffi.CallMethod[objc.Object](n_, "initWithContentsOfURL:", nibFileURL)
	rv.Autorelease()
	return rv
}

func (n_ Nib) InstantiateWithOwner_TopLevelObjects(owner objc.IObject, topLevelObjects *foundation.Array) bool {
	rv := ffi.CallMethod[bool](n_, "instantiateWithOwner:topLevelObjects:", owner, topLevelObjects)
	return rv
}