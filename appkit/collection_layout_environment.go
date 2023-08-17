// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type CollectionLayoutEnvironment interface {
	ImplementsContainer() bool
	// optional
	Container() objc.IObject
}

func WrapCollectionLayoutEnvironment(v CollectionLayoutEnvironment) objc.Object {
	return objc.WrapAsProtocol("NSCollectionLayoutEnvironment", v)
}

type CollectionLayoutEnvironmentBase struct {
}

func (p *CollectionLayoutEnvironmentBase) ImplementsContainer() bool {
	return false
}

func (p *CollectionLayoutEnvironmentBase) Container() objc.IObject {
	panic("unimpemented protocol method")
}

type CollectionLayoutEnvironmentWrapper struct {
	objc.Object
}

func (c_ CollectionLayoutEnvironmentWrapper) Container() CollectionLayoutContainerWrapper {
	rv := objc.CallMethod[CollectionLayoutContainerWrapper](c_, objc.GetSelector("container"))
	return rv
}
