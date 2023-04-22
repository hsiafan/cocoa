// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type AlertDelegate interface {
	ImplementsAlertShowHelp() bool
	// optional
	AlertShowHelp(alert Alert) bool
}

type AlertDelegateImpl struct {
	_AlertShowHelp func(alert Alert) bool
}

func (di *AlertDelegateImpl) ImplementsAlertShowHelp() bool {
	return di._AlertShowHelp != nil
}

func (di *AlertDelegateImpl) SetAlertShowHelp(f func(alert Alert) bool) {
	di._AlertShowHelp = f
}

func (di *AlertDelegateImpl) AlertShowHelp(alert Alert) bool {
	return di._AlertShowHelp(alert)
}

type AlertDelegateWrapper struct {
	objc.Object
}

func (a_ *AlertDelegateWrapper) ImplementsAlertShowHelp() bool {
	return a_.RespondsToSelector(objc.GetSelector("alertShowHelp:"))
}

func (a_ AlertDelegateWrapper) AlertShowHelp(alert IAlert) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("alertShowHelp:"), alert)
	return rv
}
