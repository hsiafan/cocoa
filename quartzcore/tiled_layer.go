// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

var TiledLayerClass = _TiledLayerClass{objc.GetClass("CATiledLayer")}

type _TiledLayerClass struct {
	objc.Class
}

type ITiledLayer interface {
	ILayer
	TileSize() coregraphics.Size
	SetTileSize(value coregraphics.Size)
}

type TiledLayer struct {
	Layer
}

func MakeTiledLayer(ptr unsafe.Pointer) TiledLayer {
	return TiledLayer{
		Layer: MakeLayer(ptr),
	}
}

func (tc _TiledLayerClass) Layer() TiledLayer {
	rv := ffi.CallMethod[TiledLayer](tc, "layer")
	return rv
}

func (t_ TiledLayer) Init() TiledLayer {
	rv := ffi.CallMethod[TiledLayer](t_, "init")
	return rv
}

func (t_ TiledLayer) InitWithLayer(layer objc.IObject) TiledLayer {
	rv := ffi.CallMethod[TiledLayer](t_, "initWithLayer:", layer)
	return rv
}

func (t_ TiledLayer) PresentationLayer() TiledLayer {
	rv := ffi.CallMethod[TiledLayer](t_, "presentationLayer")
	return rv
}

func (t_ TiledLayer) ModelLayer() TiledLayer {
	rv := ffi.CallMethod[TiledLayer](t_, "modelLayer")
	return rv
}

func (tc _TiledLayerClass) Alloc() TiledLayer {
	rv := ffi.CallMethod[TiledLayer](tc, "alloc")
	return rv
}

func (tc _TiledLayerClass) New() TiledLayer {
	rv := ffi.CallMethod[TiledLayer](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTiledLayer() TiledLayer {
	return TiledLayerClass.New()
}

func (t_ TiledLayer) TileSize() coregraphics.Size {
	rv := ffi.CallMethod[coregraphics.Size](t_, "tileSize")
	return rv
}

func (t_ TiledLayer) SetTileSize(value coregraphics.Size) {
	ffi.CallMethod[ffi.Void](t_, "setTileSize:", value)
}
