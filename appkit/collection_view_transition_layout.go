// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var CollectionViewTransitionLayoutClass = _CollectionViewTransitionLayoutClass{objc.GetClass("NSCollectionViewTransitionLayout")}

type _CollectionViewTransitionLayoutClass struct {
	objc.Class
}

type ICollectionViewTransitionLayout interface {
	ICollectionViewLayout
	UpdateValue_ForAnimatedKey(value float64, key CollectionViewTransitionLayoutAnimatedKey)
	ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey) float64
	TransitionProgress() float64
	SetTransitionProgress(value float64)
	CurrentLayout() CollectionViewLayout
	NextLayout() CollectionViewLayout
}

type CollectionViewTransitionLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewTransitionLayout(ptr unsafe.Pointer) CollectionViewTransitionLayout {
	return CollectionViewTransitionLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (c_ CollectionViewTransitionLayout) InitWithCurrentLayout_NextLayout(currentLayout ICollectionViewLayout, newLayout ICollectionViewLayout) CollectionViewTransitionLayout {
	rv := ffi.CallMethod[CollectionViewTransitionLayout](c_, "initWithCurrentLayout:nextLayout:", currentLayout, newLayout)
	return rv
}

func (cc _CollectionViewTransitionLayoutClass) Alloc() CollectionViewTransitionLayout {
	rv := ffi.CallMethod[CollectionViewTransitionLayout](cc, "alloc")
	return rv
}

func (cc _CollectionViewTransitionLayoutClass) New() CollectionViewTransitionLayout {
	rv := ffi.CallMethod[CollectionViewTransitionLayout](cc, "new")
	rv.Autorelease()
	return rv
}

func NewCollectionViewTransitionLayout() CollectionViewTransitionLayout {
	return CollectionViewTransitionLayoutClass.New()
}

func (c_ CollectionViewTransitionLayout) Init() CollectionViewTransitionLayout {
	rv := ffi.CallMethod[CollectionViewTransitionLayout](c_, "init")
	return rv
}

func (c_ CollectionViewTransitionLayout) UpdateValue_ForAnimatedKey(value float64, key CollectionViewTransitionLayoutAnimatedKey) {
	ffi.CallMethod[ffi.Void](c_, "updateValue:forAnimatedKey:", value, key)
}

func (c_ CollectionViewTransitionLayout) ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey) float64 {
	rv := ffi.CallMethod[float64](c_, "valueForAnimatedKey:", key)
	return rv
}

func (c_ CollectionViewTransitionLayout) TransitionProgress() float64 {
	rv := ffi.CallMethod[float64](c_, "transitionProgress")
	return rv
}

func (c_ CollectionViewTransitionLayout) SetTransitionProgress(value float64) {
	ffi.CallMethod[ffi.Void](c_, "setTransitionProgress:", value)
}

func (c_ CollectionViewTransitionLayout) CurrentLayout() CollectionViewLayout {
	rv := ffi.CallMethod[CollectionViewLayout](c_, "currentLayout")
	return rv
}

func (c_ CollectionViewTransitionLayout) NextLayout() CollectionViewLayout {
	rv := ffi.CallMethod[CollectionViewLayout](c_, "nextLayout")
	return rv
}
