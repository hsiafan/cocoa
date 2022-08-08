// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var RunningApplicationClass = _RunningApplicationClass{objc.GetClass("NSRunningApplication")}

type _RunningApplicationClass struct {
	objc.Class
}

type IRunningApplication interface {
	objc.IObject
	ActivateWithOptions(options ApplicationActivationOptions) bool
	Hide() bool
	Unhide() bool
	ForceTerminate() bool
	Terminate() bool
	IsActive() bool
	ActivationPolicy() ApplicationActivationPolicy
	IsHidden() bool
	LocalizedName() string
	Icon() Image
	BundleIdentifier() string
	BundleURL() foundation.URL
	ExecutableArchitecture() int
	ExecutableURL() foundation.URL
	LaunchDate() foundation.Date
	IsFinishedLaunching() bool
	ProcessIdentifier() int32
	OwnsMenuBar() bool
	IsTerminated() bool
}

type RunningApplication struct {
	objc.Object
}

func MakeRunningApplication(ptr unsafe.Pointer) RunningApplication {
	return RunningApplication{
		Object: objc.MakeObject(ptr),
	}
}

func (rc _RunningApplicationClass) RunningApplicationWithProcessIdentifier(pid int32) RunningApplication {
	rv := ffi.CallMethod[RunningApplication](rc, "runningApplicationWithProcessIdentifier:", pid)
	return rv
}

func (rc _RunningApplicationClass) Alloc() RunningApplication {
	rv := ffi.CallMethod[RunningApplication](rc, "alloc")
	return rv
}

func (r_ RunningApplication) Init() RunningApplication {
	rv := ffi.CallMethod[RunningApplication](r_, "init")
	return rv
}

func (rc _RunningApplicationClass) New() RunningApplication {
	rv := ffi.CallMethod[RunningApplication](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRunningApplication() RunningApplication {
	return RunningApplicationClass.New()
}

func (rc _RunningApplicationClass) RunningApplicationsWithBundleIdentifier(bundleIdentifier string) []RunningApplication {
	rv := ffi.CallMethod[[]RunningApplication](rc, "runningApplicationsWithBundleIdentifier:", bundleIdentifier)
	return rv
}

func (r_ RunningApplication) ActivateWithOptions(options ApplicationActivationOptions) bool {
	rv := ffi.CallMethod[bool](r_, "activateWithOptions:", options)
	return rv
}

func (r_ RunningApplication) Hide() bool {
	rv := ffi.CallMethod[bool](r_, "hide")
	return rv
}

func (r_ RunningApplication) Unhide() bool {
	rv := ffi.CallMethod[bool](r_, "unhide")
	return rv
}

func (r_ RunningApplication) ForceTerminate() bool {
	rv := ffi.CallMethod[bool](r_, "forceTerminate")
	return rv
}

func (r_ RunningApplication) Terminate() bool {
	rv := ffi.CallMethod[bool](r_, "terminate")
	return rv
}

func (rc _RunningApplicationClass) TerminateAutomaticallyTerminableApplications() {
	ffi.CallMethod[ffi.Void](rc, "terminateAutomaticallyTerminableApplications")
}

func (rc _RunningApplicationClass) CurrentApplication() RunningApplication {
	rv := ffi.CallMethod[RunningApplication](rc, "currentApplication")
	return rv
}

func (r_ RunningApplication) IsActive() bool {
	rv := ffi.CallMethod[bool](r_, "isActive")
	return rv
}

func (r_ RunningApplication) ActivationPolicy() ApplicationActivationPolicy {
	rv := ffi.CallMethod[ApplicationActivationPolicy](r_, "activationPolicy")
	return rv
}

func (r_ RunningApplication) IsHidden() bool {
	rv := ffi.CallMethod[bool](r_, "isHidden")
	return rv
}

func (r_ RunningApplication) LocalizedName() string {
	rv := ffi.CallMethod[string](r_, "localizedName")
	return rv
}

func (r_ RunningApplication) Icon() Image {
	rv := ffi.CallMethod[Image](r_, "icon")
	return rv
}

func (r_ RunningApplication) BundleIdentifier() string {
	rv := ffi.CallMethod[string](r_, "bundleIdentifier")
	return rv
}

func (r_ RunningApplication) BundleURL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](r_, "bundleURL")
	return rv
}

func (r_ RunningApplication) ExecutableArchitecture() int {
	rv := ffi.CallMethod[int](r_, "executableArchitecture")
	return rv
}

func (r_ RunningApplication) ExecutableURL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](r_, "executableURL")
	return rv
}

func (r_ RunningApplication) LaunchDate() foundation.Date {
	rv := ffi.CallMethod[foundation.Date](r_, "launchDate")
	return rv
}

func (r_ RunningApplication) IsFinishedLaunching() bool {
	rv := ffi.CallMethod[bool](r_, "isFinishedLaunching")
	return rv
}

func (r_ RunningApplication) ProcessIdentifier() int32 {
	rv := ffi.CallMethod[int32](r_, "processIdentifier")
	return rv
}

func (r_ RunningApplication) OwnsMenuBar() bool {
	rv := ffi.CallMethod[bool](r_, "ownsMenuBar")
	return rv
}

func (r_ RunningApplication) IsTerminated() bool {
	rv := ffi.CallMethod[bool](r_, "isTerminated")
	return rv
}
