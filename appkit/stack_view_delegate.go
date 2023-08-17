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

type StackViewDelegateWrapper struct {
	objc.Object
}

func (s_ StackViewDelegateWrapper) StackView_DidReattachViews(stackView IStackView, views []IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("stackView:didReattachViews:"), objc.ExtractPtr(stackView), views)
}

func (s_ StackViewDelegateWrapper) StackView_WillDetachViews(stackView IStackView, views []IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("stackView:willDetachViews:"), objc.ExtractPtr(stackView), views)
}
