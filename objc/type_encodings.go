package objc

import (
	"reflect"
	"strconv"
	"strings"
)

func getMethodTypeEncoding(ft reflect.Type, classMethod bool) string {
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
	if classMethod {
		sb.WriteString("#") // class self as first parameter
	} else {
		sb.WriteString("@") // instance self as first parameter
	}
	sb.WriteString(":") // selector
	for i := 0; i < ft.NumIn(); i++ {
		sb.WriteString(getTypeEncoding(ft.In(i)))
	}
	return sb.String()
}

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
		return "^v"
	case reflect.Pointer:
		switch t.Elem().Kind() {
		case reflect.Uint8, reflect.Int8:
			return "*"
		default:
			return "^" + getTypeEncoding(t.Elem())
		}

	case reflect.String, reflect.Slice, reflect.Map:
		return "@"
	case reflect.Array:
		return "[" + strconv.Itoa(t.Len()) + getTypeEncoding(t.Elem()) + "]"
	case reflect.Struct:
		if t == classType {
			return "#"
		} else if t == selectorType {
			return ":"
		} else if t.Implements(pointerHolderType) {
			return "@"
		} else {
			return getStructTypeEncoding(t)
		}
	default:
		panic("unsupported type:" + t.Name())
	}
}

func getStructTypeEncoding(t reflect.Type) string {
	var sb strings.Builder
	sb.WriteString("{?=")
	for i := 0; i < t.NumField(); i++ {
		sb.WriteString(getTypeEncoding(t.Field(i).Type))
	}
	sb.WriteString("}")
	return sb.String()
}
