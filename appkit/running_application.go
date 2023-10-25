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
	rv := objc.CallMethod[RunningApplication](rc, objc.SEL("alloc"))
	return rv
}

func (rc _RunningApplicationClass) New() RunningApplication {
	rv := objc.CallMethod[RunningApplication](rc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewRunningApplication() RunningApplication {
	return RunningApplicationClass.New()
}

func (r_ RunningApplication) Init() RunningApplication {
	rv := objc.CallMethod[RunningApplication](r_, objc.SEL("init"))
	return rv
}

func (rc _RunningApplicationClass) RunningApplicationsWithBundleIdentifier(bundleIdentifier string) []RunningApplication {
	rv := objc.CallMethod[[]RunningApplication](rc, objc.SEL("runningApplicationsWithBundleIdentifier:"), bundleIdentifier)
	return rv
}

func (r_ RunningApplication) ActivateWithOptions(options ApplicationActivationOptions) bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("activateWithOptions:"), options)
	return rv
}

func (r_ RunningApplication) Hide() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("hide"))
	return rv
}

func (r_ RunningApplication) Unhide() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("unhide"))
	return rv
}

func (r_ RunningApplication) ForceTerminate() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("forceTerminate"))
	return rv
}

func (r_ RunningApplication) Terminate() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("terminate"))
	return rv
}

func (rc _RunningApplicationClass) TerminateAutomaticallyTerminableApplications() {
	objc.CallMethod[objc.Void](rc, objc.SEL("terminateAutomaticallyTerminableApplications"))
}

func (rc _RunningApplicationClass) CurrentApplication() RunningApplication {
	rv := objc.CallMethod[RunningApplication](rc, objc.SEL("currentApplication"))
	return rv
}

func (r_ RunningApplication) IsActive() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("isActive"))
	return rv
}

func (r_ RunningApplication) ActivationPolicy() ApplicationActivationPolicy {
	rv := objc.CallMethod[ApplicationActivationPolicy](r_, objc.SEL("activationPolicy"))
	return rv
}

func (r_ RunningApplication) IsHidden() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("isHidden"))
	return rv
}

func (r_ RunningApplication) LocalizedName() string {
	rv := objc.CallMethod[string](r_, objc.SEL("localizedName"))
	return rv
}

func (r_ RunningApplication) Icon() Image {
	rv := objc.CallMethod[Image](r_, objc.SEL("icon"))
	return rv
}

func (r_ RunningApplication) BundleIdentifier() string {
	rv := objc.CallMethod[string](r_, objc.SEL("bundleIdentifier"))
	return rv
}

func (r_ RunningApplication) BundleURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](r_, objc.SEL("bundleURL"))
	return rv
}

func (r_ RunningApplication) ExecutableArchitecture() int {
	rv := objc.CallMethod[int](r_, objc.SEL("executableArchitecture"))
	return rv
}

func (r_ RunningApplication) ExecutableURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](r_, objc.SEL("executableURL"))
	return rv
}

func (r_ RunningApplication) LaunchDate() foundation.Date {
	rv := objc.CallMethod[foundation.Date](r_, objc.SEL("launchDate"))
	return rv
}

func (r_ RunningApplication) IsFinishedLaunching() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("isFinishedLaunching"))
	return rv
}

func (r_ RunningApplication) OwnsMenuBar() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("ownsMenuBar"))
	return rv
}

func (r_ RunningApplication) IsTerminated() bool {
	rv := objc.CallMethod[bool](r_, objc.SEL("isTerminated"))
	return rv
}
