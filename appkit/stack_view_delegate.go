// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type StackViewDelegate interface {
	ImplementsStackView_DidReattachViews() bool
	// optional
	StackView_DidReattachViews(stackView StackView, views []View)
	ImplementsStackView_WillDetachViews() bool
	// optional
	StackView_WillDetachViews(stackView StackView, views []View)
}

func WrapStackViewDelegate(v StackViewDelegate) objc.Object {
	return objc.WrapAsProtocol("NSStackViewDelegate", v)
}

type StackViewDelegateBase struct {
}

func (p *StackViewDelegateBase) ImplementsStackView_DidReattachViews() bool {
	return false
}

func (p *StackViewDelegateBase) StackView_DidReattachViews(stackView StackView, views []View) {
	panic("unimpemented protocol method")
}

func (p *StackViewDelegateBase) ImplementsStackView_WillDetachViews() bool {
	return false
}

func (p *StackViewDelegateBase) StackView_WillDetachViews(stackView StackView, views []View) {
	panic("unimpemented protocol method")
}

type StackViewDelegateCreator struct {
	className string
	class     objc.Class
}

func NewStackViewDelegateCreator(name string) *StackViewDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &StackViewDelegateCreator{className: name, class: class}
}

func (c *StackViewDelegateCreator) SetStackView_DidReattachViews(handle func(o objc.Object, stackView StackView, views []View)) {
	objc.AddMethod(c.class, objc.GetSelector("stackView:didReattachViews:"), handle)
}

func (c *StackViewDelegateCreator) SetStackView_WillDetachViews(handle func(o objc.Object, stackView StackView, views []View)) {
	objc.AddMethod(c.class, objc.GetSelector("stackView:willDetachViews:"), handle)
}

func (c *StackViewDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
