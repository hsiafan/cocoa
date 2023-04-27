package objc

import "reflect"

type UInteger = uint64
type Integer = int64

// for void type
type Void struct{}

var voidType = reflect.TypeOf((*Void)(nil)).Elem()
