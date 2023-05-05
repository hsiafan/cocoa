// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	InstantiateControllerWithIdentifier_Creator(identifier StoryboardSceneIdentifier, block func(coder foundation.Coder) objc.IObject) objc.Object
	InstantiateInitialControllerWithCreator(block func(coder foundation.Coder) objc.IObject) objc.Object
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
	rv := objc.CallMethod[Storyboard](sc, objc.GetSelector("storyboardWithName:bundle:"), name, objc.ExtractPtr(storyboardBundleOrNil))
	return rv
}

func (sc _StoryboardClass) Alloc() Storyboard {
	rv := objc.CallMethod[Storyboard](sc, objc.GetSelector("alloc"))
	return rv
}

func (sc _StoryboardClass) New() Storyboard {
	rv := objc.CallMethod[Storyboard](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewStoryboard() Storyboard {
	return StoryboardClass.New()
}

func (s_ Storyboard) Init() Storyboard {
	rv := objc.CallMethod[Storyboard](s_, objc.GetSelector("init"))
	return rv
}

func (s_ Storyboard) InstantiateInitialController() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("instantiateInitialController"))
	return rv
}

func (s_ Storyboard) InstantiateControllerWithIdentifier(identifier StoryboardSceneIdentifier) objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("instantiateControllerWithIdentifier:"), identifier)
	return rv
}

func (s_ Storyboard) InstantiateControllerWithIdentifier_Creator(identifier StoryboardSceneIdentifier, block func(coder foundation.Coder) objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("instantiateControllerWithIdentifier:creator:"), identifier, block)
	return rv
}

func (s_ Storyboard) InstantiateInitialControllerWithCreator(block func(coder foundation.Coder) objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("instantiateInitialControllerWithCreator:"), block)
	return rv
}

func (sc _StoryboardClass) MainStoryboard() Storyboard {
	rv := objc.CallMethod[Storyboard](sc, objc.GetSelector("mainStoryboard"))
	return rv
}
