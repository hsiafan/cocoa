package objc

// #import <stdint.h>
// void* C_NewAction(uintptr_t hp);
// void* C_NewDeallocListener(uintptr_t id);
import "C"
import (
	"runtime/cgo"
	"unsafe"

	"github.com/hsiafan/cocoa/internal"
)

// ActionHandler is a callback function for an ActionTarget.
type ActionHandler func(sender IObject)

// CanSetAction is an interface for objc instance which can set a target and action
type CanSetAction interface {
	IObject
	SetTarget(target IObject)
	SetAction(sel Selector)
}

// ActionTarget hold an objc ActionTarget target instance
type ActionTarget struct {
	Object
}

// WrapAction create a new objc ActionTarget instance, from ActionHandler
func WrapAction(handler ActionHandler) (target ActionTarget, selector Selector) {
	if handler == nil {
		panic("handler is nil")
	}
	h := cgo.NewHandle(handler)
	return ActionTarget{
		Object: MakeObject(C.C_NewAction(C.uintptr_t(h))),
	}, SelectorRegisterName("onAction:")
}

// SetAction set action for an ojbc instance, if it has hasTarget and setAction method.
func SetAction(instance CanSetAction, handler ActionHandler) {
	target, selector := WrapAction(handler)
	defer target.Release()
	SetAssociatedObject(instance, internal.AssociationKey("target"), target, ASSOCIATION_RETAIN)
	instance.SetTarget(target)
	instance.SetAction(selector)
}

//export callAction
func callAction(hp C.uintptr_t, senderPtr unsafe.Pointer) {
	h := cgo.Handle(hp)
	handler := h.Value().(ActionHandler)
	handler(MakeObject(senderPtr))
}

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
