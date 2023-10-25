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
