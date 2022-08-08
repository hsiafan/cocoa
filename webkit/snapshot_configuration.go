// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var SnapshotConfigurationClass = _SnapshotConfigurationClass{objc.GetClass("WKSnapshotConfiguration")}

type _SnapshotConfigurationClass struct {
	objc.Class
}

type ISnapshotConfiguration interface {
	objc.IObject
	Rect() coregraphics.Rect
	SetRect(value coregraphics.Rect)
	SnapshotWidth() foundation.Number
	SetSnapshotWidth(value foundation.INumber)
	AfterScreenUpdates() bool
	SetAfterScreenUpdates(value bool)
}

type SnapshotConfiguration struct {
	objc.Object
}

func MakeSnapshotConfiguration(ptr unsafe.Pointer) SnapshotConfiguration {
	return SnapshotConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _SnapshotConfigurationClass) Alloc() SnapshotConfiguration {
	rv := ffi.CallMethod[SnapshotConfiguration](sc, "alloc")
	return rv
}

func (s_ SnapshotConfiguration) Init() SnapshotConfiguration {
	rv := ffi.CallMethod[SnapshotConfiguration](s_, "init")
	return rv
}

func (sc _SnapshotConfigurationClass) New() SnapshotConfiguration {
	rv := ffi.CallMethod[SnapshotConfiguration](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSnapshotConfiguration() SnapshotConfiguration {
	return SnapshotConfigurationClass.New()
}

func (s_ SnapshotConfiguration) Rect() coregraphics.Rect {
	rv := ffi.CallMethod[coregraphics.Rect](s_, "rect")
	return rv
}

func (s_ SnapshotConfiguration) SetRect(value coregraphics.Rect) {
	ffi.CallMethod[ffi.Void](s_, "setRect:", value)
}

func (s_ SnapshotConfiguration) SnapshotWidth() foundation.Number {
	rv := ffi.CallMethod[foundation.Number](s_, "snapshotWidth")
	return rv
}

func (s_ SnapshotConfiguration) SetSnapshotWidth(value foundation.INumber) {
	ffi.CallMethod[ffi.Void](s_, "setSnapshotWidth:", value)
}

func (s_ SnapshotConfiguration) AfterScreenUpdates() bool {
	rv := ffi.CallMethod[bool](s_, "afterScreenUpdates")
	return rv
}

func (s_ SnapshotConfiguration) SetAfterScreenUpdates(value bool) {
	ffi.CallMethod[ffi.Void](s_, "setAfterScreenUpdates:", value)
}
