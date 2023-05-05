// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
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
	rv := objc.CallMethod[SnapshotConfiguration](sc, objc.GetSelector("alloc"))
	return rv
}

func (sc _SnapshotConfigurationClass) New() SnapshotConfiguration {
	rv := objc.CallMethod[SnapshotConfiguration](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSnapshotConfiguration() SnapshotConfiguration {
	return SnapshotConfigurationClass.New()
}

func (s_ SnapshotConfiguration) Init() SnapshotConfiguration {
	rv := objc.CallMethod[SnapshotConfiguration](s_, objc.GetSelector("init"))
	return rv
}

func (s_ SnapshotConfiguration) Rect() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](s_, objc.GetSelector("rect"))
	return rv
}

func (s_ SnapshotConfiguration) SetRect(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setRect:"), value)
}

func (s_ SnapshotConfiguration) SnapshotWidth() foundation.Number {
	rv := objc.CallMethod[foundation.Number](s_, objc.GetSelector("snapshotWidth"))
	return rv
}

func (s_ SnapshotConfiguration) SetSnapshotWidth(value foundation.INumber) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSnapshotWidth:"), objc.ExtractPtr(value))
}

func (s_ SnapshotConfiguration) AfterScreenUpdates() bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("afterScreenUpdates"))
	return rv
}

func (s_ SnapshotConfiguration) SetAfterScreenUpdates(value bool) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setAfterScreenUpdates:"), value)
}
