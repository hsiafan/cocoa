// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type CollectionLayoutEnvironment interface {
	ImplementsContainer() bool
	// optional
	Container() CollectionLayoutContainer
}

type CollectionLayoutEnvironmentWrapper struct {
	objc.Object
}

func (c_ *CollectionLayoutEnvironmentWrapper) ImplementsContainer() bool {
	return c_.RespondsToSelector(objc.GetSelector("container"))
}

func (c_ CollectionLayoutEnvironmentWrapper) Container() CollectionLayoutContainerWrapper {
	rv := ffi.CallMethod[CollectionLayoutContainerWrapper](c_, "container")
	return rv
}
