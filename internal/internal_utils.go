package internal

import "C"
import (
	"unsafe"
)

func ForceCast[S any, T any](v S) T {
	var zeroS S
	var zeroT T
	if unsafe.Sizeof(zeroS) != unsafe.Sizeof(zeroT) {
		panic("types should have same size")
	}
	return *(*T)(unsafe.Pointer(&v))
}

var associationKeyCache = SyncCache[string, unsafe.Pointer]{}

// AssociationKey return key for  AssociatedObject
func AssociationKey(name string) unsafe.Pointer {
	return associationKeyCache.Load(name, func(name string) unsafe.Pointer {
		associationKey := unsafe.Pointer(C.CString(name))
		return associationKey
	})
}
