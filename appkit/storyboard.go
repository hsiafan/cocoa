// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var StoryboardClass = _StoryboardClass{objc.GetClass("NSStoryboard")}

type _StoryboardClass struct {
	objc.Class
}

type IStoryboard interface {
	objc.IObject
	InstantiateInitialController() objc.Object
	InstantiateControllerWithIdentifier(identifier StoryboardSceneIdentifier) objc.Object
}

type Storyboard struct {
	objc.Object
}

func MakeStoryboard(ptr unsafe.Pointer) Storyboard {
	return Storyboard{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _StoryboardClass) StoryboardWithName_Bundle(name StoryboardName, storyboardBundleOrNil foundation.IBundle) Storyboard {
	rv := ffi.CallMethod[Storyboard](sc, "storyboardWithName:bundle:", name, storyboardBundleOrNil)
	return rv
}

func (sc _StoryboardClass) Alloc() Storyboard {
	rv := ffi.CallMethod[Storyboard](sc, "alloc")
	return rv
}

func (s_ Storyboard) Init() Storyboard {
	rv := ffi.CallMethod[Storyboard](s_, "init")
	return rv
}

func (sc _StoryboardClass) New() Storyboard {
	rv := ffi.CallMethod[Storyboard](sc, "new")
	rv.Autorelease()
	return rv
}

func NewStoryboard() Storyboard {
	return StoryboardClass.New()
}

func (s_ Storyboard) InstantiateInitialController() objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "instantiateInitialController")
	return rv
}

func (s_ Storyboard) InstantiateControllerWithIdentifier(identifier StoryboardSceneIdentifier) objc.Object {
	rv := ffi.CallMethod[objc.Object](s_, "instantiateControllerWithIdentifier:", identifier)
	return rv
}

func (sc _StoryboardClass) MainStoryboard() Storyboard {
	rv := ffi.CallMethod[Storyboard](sc, "mainStoryboard")
	return rv
}
