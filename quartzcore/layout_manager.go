// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type LayoutManager interface {
	ImplementsInvalidateLayoutOfLayer() bool
	// optional
	InvalidateLayoutOfLayer(layer Layer)
	ImplementsLayoutSublayersOfLayer() bool
	// optional
	LayoutSublayersOfLayer(layer Layer)
	ImplementsPreferredSizeOfLayer() bool
	// optional
	PreferredSizeOfLayer(layer Layer) coregraphics.Size
}

type LayoutManagerWrapper struct {
	objc.Object
}

func (l_ *LayoutManagerWrapper) ImplementsInvalidateLayoutOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("invalidateLayoutOfLayer:"))
}

func (l_ LayoutManagerWrapper) InvalidateLayoutOfLayer(layer ILayer) {
	ffi.CallMethod[ffi.Void](l_, "invalidateLayoutOfLayer:", layer)
}

func (l_ *LayoutManagerWrapper) ImplementsLayoutSublayersOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutSublayersOfLayer:"))
}

func (l_ LayoutManagerWrapper) LayoutSublayersOfLayer(layer ILayer) {
	ffi.CallMethod[ffi.Void](l_, "layoutSublayersOfLayer:", layer)
}

func (l_ *LayoutManagerWrapper) ImplementsPreferredSizeOfLayer() bool {
	return l_.RespondsToSelector(objc.GetSelector("preferredSizeOfLayer:"))
}

func (l_ LayoutManagerWrapper) PreferredSizeOfLayer(layer ILayer) coregraphics.Size {
	rv := ffi.CallMethod[coregraphics.Size](l_, "preferredSizeOfLayer:", layer)
	return rv
}

var ConstraintClass = _ConstraintClass{objc.GetClass("CAConstraint")}

type _ConstraintClass struct {
	objc.Class
}

type IConstraint interface {
	objc.IObject
	Attribute() ConstraintAttribute
	Offset() float64
	Scale() float64
	SourceAttribute() ConstraintAttribute
	SourceName() string
}

type Constraint struct {
	objc.Object
}

func MakeConstraint(ptr unsafe.Pointer) Constraint {
	return Constraint{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ConstraintClass) ConstraintWithAttribute_RelativeTo_Attribute_Scale_Offset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, m float64, c float64) Constraint {
	rv := ffi.CallMethod[Constraint](cc, "constraintWithAttribute:relativeTo:attribute:scale:offset:", attr, srcId, srcAttr, m, c)
	return rv
}

func (cc _ConstraintClass) ConstraintWithAttribute_RelativeTo_Attribute_Offset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, c float64) Constraint {
	rv := ffi.CallMethod[Constraint](cc, "constraintWithAttribute:relativeTo:attribute:offset:", attr, srcId, srcAttr, c)
	return rv
}

func (cc _ConstraintClass) ConstraintWithAttribute_RelativeTo_Attribute(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute) Constraint {
	rv := ffi.CallMethod[Constraint](cc, "constraintWithAttribute:relativeTo:attribute:", attr, srcId, srcAttr)
	return rv
}

func (c_ Constraint) InitWithAttribute_RelativeTo_Attribute_Scale_Offset(attr ConstraintAttribute, srcId string, srcAttr ConstraintAttribute, m float64, c float64) Constraint {
	rv := ffi.CallMethod[Constraint](c_, "initWithAttribute:relativeTo:attribute:scale:offset:", attr, srcId, srcAttr, m, c)
	return rv
}

func (cc _ConstraintClass) Alloc() Constraint {
	rv := ffi.CallMethod[Constraint](cc, "alloc")
	return rv
}

func (c_ Constraint) Init() Constraint {
	rv := ffi.CallMethod[Constraint](c_, "init")
	return rv
}

func (cc _ConstraintClass) New() Constraint {
	rv := ffi.CallMethod[Constraint](cc, "new")
	rv.Autorelease()
	return rv
}

func NewConstraint() Constraint {
	return ConstraintClass.New()
}

func (c_ Constraint) Attribute() ConstraintAttribute {
	rv := ffi.CallMethod[ConstraintAttribute](c_, "attribute")
	return rv
}

func (c_ Constraint) Offset() float64 {
	rv := ffi.CallMethod[float64](c_, "offset")
	return rv
}

func (c_ Constraint) Scale() float64 {
	rv := ffi.CallMethod[float64](c_, "scale")
	return rv
}

func (c_ Constraint) SourceAttribute() ConstraintAttribute {
	rv := ffi.CallMethod[ConstraintAttribute](c_, "sourceAttribute")
	return rv
}

func (c_ Constraint) SourceName() string {
	rv := ffi.CallMethod[string](c_, "sourceName")
	return rv
}

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
	rv := ffi.CallMethod[ConstraintLayoutManager](cc, "layoutManager")
	return rv
}

func (cc _ConstraintLayoutManagerClass) Alloc() ConstraintLayoutManager {
	rv := ffi.CallMethod[ConstraintLayoutManager](cc, "alloc")
	return rv
}

func (c_ ConstraintLayoutManager) Init() ConstraintLayoutManager {
	rv := ffi.CallMethod[ConstraintLayoutManager](c_, "init")
	return rv
}

func (cc _ConstraintLayoutManagerClass) New() ConstraintLayoutManager {
	rv := ffi.CallMethod[ConstraintLayoutManager](cc, "new")
	rv.Autorelease()
	return rv
}

func NewConstraintLayoutManager() ConstraintLayoutManager {
	return ConstraintLayoutManagerClass.New()
}
