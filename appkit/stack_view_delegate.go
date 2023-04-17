// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
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

type StackViewDelegateImpl struct {
	_StackView_DidReattachViews func(stackView StackView, views []View)
	_StackView_WillDetachViews  func(stackView StackView, views []View)
}

func (di *StackViewDelegateImpl) ImplementsStackView_DidReattachViews() bool {
	return di._StackView_DidReattachViews != nil
}

func (di *StackViewDelegateImpl) SetStackView_DidReattachViews(f func(stackView StackView, views []View)) {
	di._StackView_DidReattachViews = f
}

func (di *StackViewDelegateImpl) StackView_DidReattachViews(stackView StackView, views []View) {
	di._StackView_DidReattachViews(stackView, views)
}
func (di *StackViewDelegateImpl) ImplementsStackView_WillDetachViews() bool {
	return di._StackView_WillDetachViews != nil
}

func (di *StackViewDelegateImpl) SetStackView_WillDetachViews(f func(stackView StackView, views []View)) {
	di._StackView_WillDetachViews = f
}

func (di *StackViewDelegateImpl) StackView_WillDetachViews(stackView StackView, views []View) {
	di._StackView_WillDetachViews(stackView, views)
}

type StackViewDelegateWrapper struct {
	objc.Object
}

func (s_ *StackViewDelegateWrapper) ImplementsStackView_DidReattachViews() bool {
	return s_.RespondsToSelector(objc.GetSelector("stackView:didReattachViews:"))
}

func (s_ StackViewDelegateWrapper) StackView_DidReattachViews(stackView IStackView, views []IView) {
	ffi.CallMethod[ffi.Void](s_, "stackView:didReattachViews:", stackView, views)
}

func (s_ *StackViewDelegateWrapper) ImplementsStackView_WillDetachViews() bool {
	return s_.RespondsToSelector(objc.GetSelector("stackView:willDetachViews:"))
}

func (s_ StackViewDelegateWrapper) StackView_WillDetachViews(stackView IStackView, views []IView) {
	ffi.CallMethod[ffi.Void](s_, "stackView:willDetachViews:", stackView, views)
}
