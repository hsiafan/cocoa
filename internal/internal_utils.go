package internal

import "C"
import (
	"strings"
	"unsafe"
)

func ForceCast[S any, T any](v S) T {
	return *(*T)(unsafe.Pointer(&v))
}

// menuWillOpen: -> MenuWillOpen
// menu:updateItem:atIndex:shouldCancel: -> Menu_UpdateItem_AtIndex_ShouldCancel
func SelectorToGoName(sel string) string {
	var sb strings.Builder
	sb.Grow(len(sel))
	for i, item := range strings.Split(sel, ":") {
		if len(item) == 0 {
			continue
		}
		if i > 0 {
			sb.WriteByte('_')
		}
		sb.WriteString(strings.ToUpper(item[:1]))
		sb.WriteString(item[1:])
	}
	return sb.String()
}

var associationKeyCache = SyncCache[string, unsafe.Pointer]{}

// AssociationKey return key for  AssociatedObject
func AssociationKey(name string) unsafe.Pointer {
	return associationKeyCache.Load(name, func(name string) unsafe.Pointer {
		associationKey := unsafe.Pointer(C.CString(name))
		return associationKey
	})
}
