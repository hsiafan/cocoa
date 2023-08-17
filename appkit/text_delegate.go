// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

type TextDelegate interface {
	ImplementsTextDidChange() bool
	// optional
	TextDidChange(notification foundation.Notification)
	ImplementsTextShouldBeginEditing() bool
	// optional
	TextShouldBeginEditing(textObject Text) bool
	ImplementsTextDidBeginEditing() bool
	// optional
	TextDidBeginEditing(notification foundation.Notification)
	ImplementsTextShouldEndEditing() bool
	// optional
	TextShouldEndEditing(textObject Text) bool
	ImplementsTextDidEndEditing() bool
	// optional
	TextDidEndEditing(notification foundation.Notification)
}

func WrapTextDelegate(v TextDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTextDelegate", v)
}

type TextDelegateBase struct {
}

func (p *TextDelegateBase) ImplementsTextDidChange() bool {
	return false
}

func (p *TextDelegateBase) TextDidChange(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *TextDelegateBase) ImplementsTextShouldBeginEditing() bool {
	return false
}

func (p *TextDelegateBase) TextShouldBeginEditing(textObject Text) bool {
	panic("unimpemented protocol method")
}

func (p *TextDelegateBase) ImplementsTextDidBeginEditing() bool {
	return false
}

func (p *TextDelegateBase) TextDidBeginEditing(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

func (p *TextDelegateBase) ImplementsTextShouldEndEditing() bool {
	return false
}

func (p *TextDelegateBase) TextShouldEndEditing(textObject Text) bool {
	panic("unimpemented protocol method")
}

func (p *TextDelegateBase) ImplementsTextDidEndEditing() bool {
	return false
}

func (p *TextDelegateBase) TextDidEndEditing(notification foundation.Notification) {
	panic("unimpemented protocol method")
}

type TextDelegateWrapper struct {
	objc.Object
}

func (t_ TextDelegateWrapper) TextDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidChange:"), objc.ExtractPtr(notification))
}

func (t_ TextDelegateWrapper) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldBeginEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextDelegateWrapper) TextDidBeginEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidBeginEditing:"), objc.ExtractPtr(notification))
}

func (t_ TextDelegateWrapper) TextShouldEndEditing(textObject IText) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textShouldEndEditing:"), objc.ExtractPtr(textObject))
	return rv
}

func (t_ TextDelegateWrapper) TextDidEndEditing(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textDidEndEditing:"), objc.ExtractPtr(notification))
}
