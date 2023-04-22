// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ValueClass = _ValueClass{objc.GetClass("NSValue")}

type _ValueClass struct {
	objc.Class
}

type IValue interface {
	objc.IObject
	// deprecated
	GetValue(value unsafe.Pointer)
	IsEqualToValue(value IValue) bool
	GetValue_Size(value unsafe.Pointer, size uint)
	ObjCType() *byte
	PointerValue() unsafe.Pointer
	NonretainedObjectValue() objc.Object
	RangeValue() Range
	PointValue() Point
	SizeValue() Size
	RectValue() Rect
	EdgeInsetsValue() EdgeInsets
}

type Value struct {
	objc.Object
}

func MakeValue(ptr unsafe.Pointer) Value {
	return Value{
		Object: objc.MakeObject(ptr),
	}
}

func (v_ Value) InitWithBytes_ObjCType(value unsafe.Pointer, type_ *byte) Value {
	rv := objc.CallMethod[Value](v_, objc.GetSelector("initWithBytes:objCType:"), value, type_)
	return rv
}

func (vc _ValueClass) Alloc() Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("alloc"))
	return rv
}

func (vc _ValueClass) New() Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewValue() Value {
	return ValueClass.New()
}

func (v_ Value) Init() Value {
	rv := objc.CallMethod[Value](v_, objc.GetSelector("init"))
	return rv
}

func (vc _ValueClass) ValueWithBytes_ObjCType(value unsafe.Pointer, type_ *byte) Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("valueWithBytes:objCType:"), value, type_)
	return rv
}

func (vc _ValueClass) Value_WithObjCType(value unsafe.Pointer, type_ *byte) Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("value:withObjCType:"), value, type_)
	return rv
}

// deprecated
func (v_ Value) GetValue(value unsafe.Pointer) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("getValue:"), value)
}

func (vc _ValueClass) ValueWithPointer(pointer unsafe.Pointer) Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("valueWithPointer:"), pointer)
	return rv
}

func (vc _ValueClass) ValueWithNonretainedObject(anObject objc.IObject) Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("valueWithNonretainedObject:"), anObject)
	return rv
}

func (vc _ValueClass) ValueWithRange(range_ Range) Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("valueWithRange:"), range_)
	return rv
}

func (vc _ValueClass) ValueWithPoint(point Point) Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("valueWithPoint:"), point)
	return rv
}

func (vc _ValueClass) ValueWithSize(size Size) Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("valueWithSize:"), size)
	return rv
}

func (vc _ValueClass) ValueWithRect(rect Rect) Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("valueWithRect:"), rect)
	return rv
}

func (v_ Value) IsEqualToValue(value IValue) bool {
	rv := objc.CallMethod[bool](v_, objc.GetSelector("isEqualToValue:"), value)
	return rv
}

func (vc _ValueClass) ValueWithEdgeInsets(insets EdgeInsets) Value {
	rv := objc.CallMethod[Value](vc, objc.GetSelector("valueWithEdgeInsets:"), insets)
	return rv
}

func (v_ Value) GetValue_Size(value unsafe.Pointer, size uint) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("getValue:size:"), value, size)
}

func (v_ Value) ObjCType() *byte {
	rv := objc.CallMethod[*byte](v_, objc.GetSelector("objCType"))
	return rv
}

func (v_ Value) PointerValue() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](v_, objc.GetSelector("pointerValue"))
	return rv
}

func (v_ Value) NonretainedObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](v_, objc.GetSelector("nonretainedObjectValue"))
	return rv
}

func (v_ Value) RangeValue() Range {
	rv := objc.CallMethod[Range](v_, objc.GetSelector("rangeValue"))
	return rv
}

func (v_ Value) PointValue() Point {
	rv := objc.CallMethod[Point](v_, objc.GetSelector("pointValue"))
	return rv
}

func (v_ Value) SizeValue() Size {
	rv := objc.CallMethod[Size](v_, objc.GetSelector("sizeValue"))
	return rv
}

func (v_ Value) RectValue() Rect {
	rv := objc.CallMethod[Rect](v_, objc.GetSelector("rectValue"))
	return rv
}

func (v_ Value) EdgeInsetsValue() EdgeInsets {
	rv := objc.CallMethod[EdgeInsets](v_, objc.GetSelector("edgeInsetsValue"))
	return rv
}
