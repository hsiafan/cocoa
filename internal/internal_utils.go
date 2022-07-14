package internal

import "C"
import (
	"strings"
	"sync"
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

var associationKeys = map[string]unsafe.Pointer{}
var associationKeysLock sync.Mutex

// AssociationKey return key for  AssociatedObject
func AssociationKey(name string) unsafe.Pointer {
	associationKeysLock.Lock()
	defer associationKeysLock.Unlock()
	if key, ok := associationKeys[name]; ok {
		return key
	}
	key := unsafe.Pointer(C.CString(name))
	associationKeys[name] = key
	return key
}
