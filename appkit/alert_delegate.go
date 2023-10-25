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

type AlertDelegateCreator struct {
	className string
	class     objc.Class
}

func NewAlertDelegateCreator(name string) *AlertDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &AlertDelegateCreator{className: name, class: class}
}

func (c *AlertDelegateCreator) SetAlertShowHelp(handle func(o objc.Object, alert Alert) bool) {
	objc.AddMethod(c.class, objc.SEL("alertShowHelp:"), handle)
}

func (c *AlertDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
