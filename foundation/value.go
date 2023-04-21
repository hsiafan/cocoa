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
	rv := objc.CallMethod[Value](v_, "initWithBytes:objCType:", value, type_)
	return rv
}

func (vc _ValueClass) Alloc() Value {
	rv := objc.CallMethod[Value](vc, "alloc")
	return rv
}

func (vc _ValueClass) New() Value {
	rv := objc.CallMethod[Value](vc, "new")
	rv.Autorelease()
	return rv
}

func NewValue() Value {
	return ValueClass.New()
}

func (v_ Value) Init() Value {
	rv := objc.CallMethod[Value](v_, "init")
	return rv
}

func (vc _ValueClass) ValueWithBytes_ObjCType(value unsafe.Pointer, type_ *byte) Value {
	rv := objc.CallMethod[Value](vc, "valueWithBytes:objCType:", value, type_)
	return rv
}

func (vc _ValueClass) Value_WithObjCType(value unsafe.Pointer, type_ *byte) Value {
	rv := objc.CallMethod[Value](vc, "value:withObjCType:", value, type_)
	return rv
}

// deprecated
func (v_ Value) GetValue(value unsafe.Pointer) {
	objc.CallMethod[objc.Void](v_, "getValue:", value)
}

func (vc _ValueClass) ValueWithPointer(pointer unsafe.Pointer) Value {
	rv := objc.CallMethod[Value](vc, "valueWithPointer:", pointer)
	return rv
}

func (vc _ValueClass) ValueWithNonretainedObject(anObject objc.IObject) Value {
	rv := objc.CallMethod[Value](vc, "valueWithNonretainedObject:", anObject)
	return rv
}

func (vc _ValueClass) ValueWithRange(range_ Range) Value {
	rv := objc.CallMethod[Value](vc, "valueWithRange:", range_)
	return rv
}

func (vc _ValueClass) ValueWithPoint(point Point) Value {
	rv := objc.CallMethod[Value](vc, "valueWithPoint:", point)
	return rv
}

func (vc _ValueClass) ValueWithSize(size Size) Value {
	rv := objc.CallMethod[Value](vc, "valueWithSize:", size)
	return rv
}

func (vc _ValueClass) ValueWithRect(rect Rect) Value {
	rv := objc.CallMethod[Value](vc, "valueWithRect:", rect)
	return rv
}

func (v_ Value) IsEqualToValue(value IValue) bool {
	rv := objc.CallMethod[bool](v_, "isEqualToValue:", value)
	return rv
}

func (vc _ValueClass) ValueWithEdgeInsets(insets EdgeInsets) Value {
	rv := objc.CallMethod[Value](vc, "valueWithEdgeInsets:", insets)
	return rv
}

func (v_ Value) GetValue_Size(value unsafe.Pointer, size uint) {
	objc.CallMethod[objc.Void](v_, "getValue:size:", value, size)
}

func (v_ Value) ObjCType() *byte {
	rv := objc.CallMethod[*byte](v_, "objCType")
	return rv
}

func (v_ Value) PointerValue() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](v_, "pointerValue")
	return rv
}

func (v_ Value) NonretainedObjectValue() objc.Object {
	rv := objc.CallMethod[objc.Object](v_, "nonretainedObjectValue")
	return rv
}

func (v_ Value) RangeValue() Range {
	rv := objc.CallMethod[Range](v_, "rangeValue")
	return rv
}

func (v_ Value) PointValue() Point {
	rv := objc.CallMethod[Point](v_, "pointValue")
	return rv
}

func (v_ Value) SizeValue() Size {
	rv := objc.CallMethod[Size](v_, "sizeValue")
	return rv
}

func (v_ Value) RectValue() Rect {
	rv := objc.CallMethod[Rect](v_, "rectValue")
	return rv
}

func (v_ Value) EdgeInsetsValue() EdgeInsets {
	rv := objc.CallMethod[EdgeInsets](v_, "edgeInsetsValue")
	return rv
}
