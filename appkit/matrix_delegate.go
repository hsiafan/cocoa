// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type MatrixDelegate interface {
	ControlTextEditingDelegate
}

func WrapMatrixDelegate(v MatrixDelegate) objc.Object {
	return objc.WrapAsProtocol("NSMatrixDelegate", v)
}

type MatrixDelegateBase struct {
	ControlTextEditingDelegateBase
}
type MatrixDelegateCreator struct {
	className string
	class     objc.Class
}

func NewMatrixDelegateCreator(name string) *MatrixDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &MatrixDelegateCreator{className: name, class: class}
}

func (c *MatrixDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
