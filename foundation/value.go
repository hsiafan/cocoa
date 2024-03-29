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
	rv := objc.CallMethod[Value](v_, objc.SEL("initWithBytes:objCType:"), value, type_)
	return rv
}

func (vc _ValueClass) Alloc() Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("alloc"))
	return rv
}

func (vc _ValueClass) New() Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewValue() Value {
	return ValueClass.New()
}

func (v_ Value) Init() Value {
	rv := objc.CallMethod[Value](v_, objc.SEL("init"))
	return rv
}

func (vc _ValueClass) ValueWithBytes_ObjCType(value unsafe.Pointer, type_ *byte) Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("valueWithBytes:objCType:"), value, type_)
	return rv
}

func (vc _ValueClass) Value_WithObjCType(value unsafe.Pointer, type_ *byte) Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("value:withObjCType:"), value, type_)
	return rv
}

// deprecated
func (v_ Value) GetValue(value unsafe.Pointer) {
	objc.CallMethod[objc.Void](v_, objc.SEL("getValue:"), value)
}

func (vc _ValueClass) ValueWithPointer(pointer unsafe.Pointer) Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("valueWithPointer:"), pointer)
	return rv
}

func (vc _ValueClass) ValueWithNonretainedObject(anObject objc.IObject) Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("valueWithNonretainedObject:"), objc.ExtractPtr(anObject))
	return rv
}

func (vc _ValueClass) ValueWithRange(range_ Range) Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("valueWithRange:"), range_)
	return rv
}

func (vc _ValueClass) ValueWithPoint(point Point) Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("valueWithPoint:"), point)
	return rv
}

func (vc _ValueClass) ValueWithSize(size Size) Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("valueWithSize:"), size)
	return rv
}

func (vc _ValueClass) ValueWithRect(rect Rect) Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("valueWithRect:"), rect)
	return rv
}

func (v_ Value) IsEqualToValue(value IValue) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isEqualToValue:"), objc.ExtractPtr(value))
	return rv
}

func (vc _ValueClass) ValueWithEdgeInsets(insets EdgeInsets) Value {
	rv := objc.CallMethod[Value](vc, objc.SEL("valueWithEdgeInsets:"), insets)
	return rv
}

func (v_ Value) GetValue_Size(value unsafe.Pointer, size uint) {
	objc.CallMethod[objc.Void](v_, objc.SEL("getValue:size:"), value, size)
}

func (v_ Value) ObjCType() *byte {
	rv := objc.CallMethod[*byte](v_, objc.SEL("objCType"))
	return rv
}

func (v_ Value) PointerValue() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](v_, objc.SEL("pointerValue"))
	return rv
}

func (v_ Value) NonretainedObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](v_, objc.SEL("nonretainedObjectValue"))
	return rv
}

func (v_ Value) RangeValue() Range {
	rv := objc.CallMethod[Range](v_, objc.SEL("rangeValue"))
	return rv
}

func (v_ Value) PointValue() Point {
	rv := objc.CallMethod[Point](v_, objc.SEL("pointValue"))
	return rv
}

func (v_ Value) SizeValue() Size {
	rv := objc.CallMethod[Size](v_, objc.SEL("sizeValue"))
	return rv
}

func (v_ Value) RectValue() Rect {
	rv := objc.CallMethod[Rect](v_, objc.SEL("rectValue"))
	return rv
}

func (v_ Value) EdgeInsetsValue() EdgeInsets {
	rv := objc.CallMethod[EdgeInsets](v_, objc.SEL("edgeInsetsValue"))
	return rv
}
