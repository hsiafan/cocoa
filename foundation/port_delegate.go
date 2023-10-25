// auto-generated code, do not modify
package foundation

import (
	"github.com/hsiafan/cocoa/objc"
)

type PortDelegate interface {
}

func WrapPortDelegate(v PortDelegate) objc.Object {
	return objc.WrapAsProtocol("NSPortDelegate", v)
}

type PortDelegateBase struct {
}
type PortDelegateCreator struct {
	className string
	class     objc.Class
}

func NewPortDelegateCreator(name string) *PortDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &PortDelegateCreator{className: name, class: class}
}

func (c *PortDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
