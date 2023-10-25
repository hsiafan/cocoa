// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type SearchFieldDelegate interface {
	TextFieldDelegate
	ImplementsSearchFieldDidStartSearching() bool
	// optional
	SearchFieldDidStartSearching(sender SearchField)
	ImplementsSearchFieldDidEndSearching() bool
	// optional
	SearchFieldDidEndSearching(sender SearchField)
}

func WrapSearchFieldDelegate(v SearchFieldDelegate) objc.Object {
	return objc.WrapAsProtocol("NSSearchFieldDelegate", v)
}

type SearchFieldDelegateBase struct {
	TextFieldDelegateBase
}

func (p *SearchFieldDelegateBase) ImplementsSearchFieldDidStartSearching() bool {
	return false
}

func (p *SearchFieldDelegateBase) SearchFieldDidStartSearching(sender SearchField) {
	panic("unimpemented protocol method")
}

func (p *SearchFieldDelegateBase) ImplementsSearchFieldDidEndSearching() bool {
	return false
}

func (p *SearchFieldDelegateBase) SearchFieldDidEndSearching(sender SearchField) {
	panic("unimpemented protocol method")
}

type SearchFieldDelegateCreator struct {
	className string
	class     objc.Class
}

func NewSearchFieldDelegateCreator(name string) *SearchFieldDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &SearchFieldDelegateCreator{className: name, class: class}
}

func (c *SearchFieldDelegateCreator) SetSearchFieldDidStartSearching(handle func(o objc.ProtocolBase, sender SearchField)) {
	objc.AddMethod(c.class, objc.SEL("searchFieldDidStartSearching:"), handle)
}

func (c *SearchFieldDelegateCreator) SetSearchFieldDidEndSearching(handle func(o objc.ProtocolBase, sender SearchField)) {
	objc.AddMethod(c.class, objc.SEL("searchFieldDidEndSearching:"), handle)
}

func (c *SearchFieldDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
