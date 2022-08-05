package objc

// #import <stdint.h>
// void* C_NewDeallocListener(uintptr_t id);
import "C"
import (
	"runtime/cgo"

	"github.com/hsiafan/cocoa/internal"
)

//export deleteHandle
func deleteHandle(hp C.uintptr_t) {
	h := cgo.Handle(hp)
	h.Delete()
}

// SetDeallocListener set a listener to be invoked when object ref count is decreased to zero(so the object is dealloced).
// Call dealloc method directly, but ref count is still large than 0 will not trigger the listener.
// Call multi times will remove previouse listener.
func SetDeallocListener(o IObject, listener func()) {
	h := cgo.NewHandle(listener)
	lo := MakeObject(C.C_NewDeallocListener(C.uintptr_t(h)))
	defer lo.Release()
	SetAssociatedObject(o, internal.AssociationKey("deallocListener"), lo, ASSOCIATION_RETAIN)
}

func TryRetain(objects ...IObject) {
	for _, o := range objects {
		if o != nil && !o.IsNil() {
			o.Retain()
		}
	}
}

func TryRelease(objects ...IObject) {
	for _, o := range objects {
		if o != nil && o.Ptr() != nil {
			o.Release()
		}
	}
}
