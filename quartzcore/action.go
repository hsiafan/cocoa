// auto-generated code, do not modify
package quartzcore

import (
	"github.com/hsiafan/cocoa/objc"
)

type Action interface {
}

func WrapAction(v Action) objc.Object {
	return objc.WrapAsProtocol("CAAction", v)
}

type ActionBase struct {
}
