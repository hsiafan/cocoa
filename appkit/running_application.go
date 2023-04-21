// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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

func (rc _RunningApplicationClass) Alloc() RunningApplication {
	rv := objc.CallMethod[RunningApplication](rc, "alloc")
	return rv
}

func (rc _RunningApplicationClass) New() RunningApplication {
	rv := objc.CallMethod[RunningApplication](rc, "new")
	rv.Autorelease()
	return rv
}

func NewRunningApplication() RunningApplication {
	return RunningApplicationClass.New()
}

func (r_ RunningApplication) Init() RunningApplication {
	rv := objc.CallMethod[RunningApplication](r_, "init")
	return rv
}

func (rc _RunningApplicationClass) RunningApplicationsWithBundleIdentifier(bundleIdentifier string) []RunningApplication {
	rv := objc.CallMethod[[]RunningApplication](rc, "runningApplicationsWithBundleIdentifier:", bundleIdentifier)
	return rv
}

func (r_ RunningApplication) ActivateWithOptions(options ApplicationActivationOptions) bool {
	rv := objc.CallMethod[bool](r_, "activateWithOptions:", options)
	return rv
}

func (r_ RunningApplication) Hide() bool {
	rv := objc.CallMethod[bool](r_, "hide")
	return rv
}

func (r_ RunningApplication) Unhide() bool {
	rv := objc.CallMethod[bool](r_, "unhide")
	return rv
}

func (r_ RunningApplication) ForceTerminate() bool {
	rv := objc.CallMethod[bool](r_, "forceTerminate")
	return rv
}

func (r_ RunningApplication) Terminate() bool {
	rv := objc.CallMethod[bool](r_, "terminate")
	return rv
}

func (rc _RunningApplicationClass) TerminateAutomaticallyTerminableApplications() {
	objc.CallMethod[objc.Void](rc, "terminateAutomaticallyTerminableApplications")
}

func (rc _RunningApplicationClass) CurrentApplication() RunningApplication {
	rv := objc.CallMethod[RunningApplication](rc, "currentApplication")
	return rv
}

func (r_ RunningApplication) IsActive() bool {
	rv := objc.CallMethod[bool](r_, "isActive")
	return rv
}

func (r_ RunningApplication) ActivationPolicy() ApplicationActivationPolicy {
	rv := objc.CallMethod[ApplicationActivationPolicy](r_, "activationPolicy")
	return rv
}

func (r_ RunningApplication) IsHidden() bool {
	rv := objc.CallMethod[bool](r_, "isHidden")
	return rv
}

func (r_ RunningApplication) LocalizedName() string {
	rv := objc.CallMethod[string](r_, "localizedName")
	return rv
}

func (r_ RunningApplication) Icon() Image {
	rv := objc.CallMethod[Image](r_, "icon")
	return rv
}

func (r_ RunningApplication) BundleIdentifier() string {
	rv := objc.CallMethod[string](r_, "bundleIdentifier")
	return rv
}

func (r_ RunningApplication) BundleURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](r_, "bundleURL")
	return rv
}

func (r_ RunningApplication) ExecutableArchitecture() int {
	rv := objc.CallMethod[int](r_, "executableArchitecture")
	return rv
}

func (r_ RunningApplication) ExecutableURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](r_, "executableURL")
	return rv
}

func (r_ RunningApplication) LaunchDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](r_, "launchDate")
	return rv
}

func (r_ RunningApplication) IsFinishedLaunching() bool {
	rv := objc.CallMethod[bool](r_, "isFinishedLaunching")
	return rv
}

func (r_ RunningApplication) OwnsMenuBar() bool {
	rv := objc.CallMethod[bool](r_, "ownsMenuBar")
	return rv
}

func (r_ RunningApplication) IsTerminated() bool {
	rv := objc.CallMethod[bool](r_, "isTerminated")
	return rv
}
