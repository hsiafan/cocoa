package objc

// #import <stdint.h>
// void* C_NewDeallocListener(uintptr_t id);
// void Run_WithAutoreleasePool(uintptr_t ptr);
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

// WithAutoreleasePool run code in a new auto release pool.
func WithAutoreleasePool(task func()) {
	id := cgo.NewHandle(task)
	C.Run_WithAutoreleasePool(C.uintptr_t(id))
}

//export runTaskAndDeleteHandle
func runTaskAndDeleteHandle(p C.uintptr_t) {
	h := cgo.Handle(p)
	task := h.Value().(func())
	task()
	h.Delete()
}
