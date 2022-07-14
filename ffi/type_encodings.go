package ffi

import (
	"reflect"
	"strings"
)

func getBlockTypeEncoding(ft reflect.Type) string {
	if ft.Kind() != reflect.Func {
		panic("not func type")
	}
	if ft.NumOut() > 1 {
		panic("to many return values")
	}
	var sb strings.Builder
	if ft.NumOut() == 0 {
		sb.WriteByte('v')
	} else {
		sb.WriteString(getTypeEncoding(ft.Out(0)))
	}
	sb.WriteString("@?") // block self as first parameter
	for i := 0; i < ft.NumIn(); i++ {
		sb.WriteString(getTypeEncoding(ft.In(i)))
	}
	return sb.String()
}

func getTypeEncoding(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Bool:
		return "B"
	case reflect.Int8:
		return "c"
	case reflect.Int16:
		return "s"
	case reflect.Int32:
		return "i"
	case reflect.Int64, reflect.Int:
		return "q"
	case reflect.Uint8:
		return "C"
	case reflect.Uint16:
		return "S"
	case reflect.Uint32:
		return "I"
	case reflect.Uint64, reflect.Uint:
		return "Q"
	case reflect.Float32:
		return "f"
	case reflect.Float64:
		return "d"
	case reflect.UnsafePointer:
		panic("unsupported type") // shoudl use Class/Selector/Oject...
	case reflect.Pointer:
		//TODO: * for char*
		return "^" + getTypeEncoding(t.Elem())
	case reflect.String, reflect.Slice, reflect.Map:
		return "@"
	case reflect.Struct:
		if t == classType {
			return "#"
		} else if t == selectorType {
			return ":"
		} else if t.Implements(pointerHolderType) {
			return "@"
		} else {
			switch t.Name() {
			case "Size":
				return "{CGSize=dd}"
			case "Point":
				return "{CGPoint=dd}"
			case "Rect":
				return "{CGRect={CGPoint=dd}{CGSize=dd}}"
			case "EdgeInsets":
				return "{NSEdgeInsets=dddd}"
			case "Range":
				return "{NSRange=QQ}"
			case "AffineTransformStruct":
				return "{NSAffineTransformStruct=dddddd}"
			case "Decimal":
				return "{?=b8b4b1b1b18[8S]}"
			case "DirectionalEdgeInsets":
				return "{NSDirectionalEdgeInsets=dddd}"
			case "AffineTransform":
				return "{CGAffineTransform=dddddd}"
			default:
				panic("unsupported type:" + t.Name())
			}

		}
	default:
		panic("unsupported type:" + t.Name())
	}
}
