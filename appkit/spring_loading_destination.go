package appkit

// #include "spring_loading_destination.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type SpringLoadingDestination struct {
	SpringLoadingActivated_DraggingInfo func(activated bool, draggingInfo objc.Object) // required
	SpringLoadingHighlightChanged       func(draggingInfo objc.Object)                 // required
	SpringLoadingEntered                func(draggingInfo objc.Object) SpringLoadingOptions
	SpringLoadingUpdated                func(draggingInfo objc.Object) SpringLoadingOptions
	SpringLoadingExited                 func(draggingInfo objc.Object)
	DraggingEnded                       func(draggingInfo objc.Object)
}

func WrapSpringLoadingDestination(delegate *SpringLoadingDestination) objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapSpringLoadingDestination(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export springLoadingDestination_SpringLoadingActivated_DraggingInfo
func springLoadingDestination_SpringLoadingActivated_DraggingInfo(hp C.uintptr_t, activated C.bool, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	delegate.SpringLoadingActivated_DraggingInfo(bool(activated), objc.MakeObject(draggingInfo))
}

//export springLoadingDestination_SpringLoadingHighlightChanged
func springLoadingDestination_SpringLoadingHighlightChanged(hp C.uintptr_t, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	delegate.SpringLoadingHighlightChanged(objc.MakeObject(draggingInfo))
}

//export springLoadingDestination_SpringLoadingEntered
func springLoadingDestination_SpringLoadingEntered(hp C.uintptr_t, draggingInfo unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	result := delegate.SpringLoadingEntered(objc.MakeObject(draggingInfo))
	return C.uint(uint(result))
}

//export springLoadingDestination_SpringLoadingUpdated
func springLoadingDestination_SpringLoadingUpdated(hp C.uintptr_t, draggingInfo unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	result := delegate.SpringLoadingUpdated(objc.MakeObject(draggingInfo))
	return C.uint(uint(result))
}

//export springLoadingDestination_SpringLoadingExited
func springLoadingDestination_SpringLoadingExited(hp C.uintptr_t, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	delegate.SpringLoadingExited(objc.MakeObject(draggingInfo))
}

//export springLoadingDestination_DraggingEnded
func springLoadingDestination_DraggingEnded(hp C.uintptr_t, draggingInfo unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	delegate.DraggingEnded(objc.MakeObject(draggingInfo))
}

//export SpringLoadingDestination_RespondsTo
func SpringLoadingDestination_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*SpringLoadingDestination)
	switch selName {
	case "springLoadingActivated:draggingInfo:":
		return delegate.SpringLoadingActivated_DraggingInfo != nil
	case "springLoadingHighlightChanged:":
		return delegate.SpringLoadingHighlightChanged != nil
	case "springLoadingEntered:":
		return delegate.SpringLoadingEntered != nil
	case "springLoadingUpdated:":
		return delegate.SpringLoadingUpdated != nil
	case "springLoadingExited:":
		return delegate.SpringLoadingExited != nil
	case "draggingEnded:":
		return delegate.DraggingEnded != nil
	default:
		return false
	}
}

//export deleteSpringLoadingDestination
func deleteSpringLoadingDestination(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
