// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	rv := objc.CallMethod[Nib](n_, objc.GetSelector("initWithNibNamed:bundle:"), nibName, objc.ExtractPtr(bundle))
	return rv
}

func (n_ Nib) InitWithNibData_Bundle(nibData []byte, bundle foundation.IBundle) Nib {
	rv := objc.CallMethod[Nib](n_, objc.GetSelector("initWithNibData:bundle:"), nibData, objc.ExtractPtr(bundle))
	return rv
}

func (nc _NibClass) Alloc() Nib {
	rv := objc.CallMethod[Nib](nc, objc.GetSelector("alloc"))
	return rv
}

func (nc _NibClass) New() Nib {
	rv := objc.CallMethod[Nib](nc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewNib() Nib {
	return NibClass.New()
}

func (n_ Nib) Init() Nib {
	rv := objc.CallMethod[Nib](n_, objc.GetSelector("init"))
	return rv
}

// deprecated
func (n_ Nib) InitWithContentsOfURL(nibFileURL foundation.IURL) objc.Object {
	rv := objc.CallMethod[objc.Object](n_, objc.GetSelector("initWithContentsOfURL:"), objc.ExtractPtr(nibFileURL))
	return rv
}

func (n_ Nib) InstantiateWithOwner_TopLevelObjects(owner objc.IObject, topLevelObjects *foundation.Array) bool {
	rv := objc.CallMethod[bool](n_, objc.GetSelector("instantiateWithOwner:topLevelObjects:"), objc.ExtractPtr(owner), topLevelObjects)
	return rv
}
