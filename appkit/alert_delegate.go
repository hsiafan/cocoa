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

func WrapAlertDelegate(v AlertDelegate) objc.Object {
	return objc.WrapAsProtocol("NSAlertDelegate", v)
}

type AlertDelegateBase struct {
}

func (p *AlertDelegateBase) ImplementsAlertShowHelp() bool {
	return false
}

func (p *AlertDelegateBase) AlertShowHelp(alert Alert) bool {
	panic("unimpemented protocol method")
}

type AlertDelegateWrapper struct {
	objc.Object
}

func (a_ AlertDelegateWrapper) AlertShowHelp(alert IAlert) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("alertShowHelp:"), objc.ExtractPtr(alert))
	return rv
}
