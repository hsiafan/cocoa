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

type TextDelegateCreator struct {
	className string
	class     objc.Class
}

func NewTextDelegateCreator(name string) *TextDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &TextDelegateCreator{className: name, class: class}
}

func (c *TextDelegateCreator) SetTextDidChange(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("textDidChange:"), handle)
}

func (c *TextDelegateCreator) SetTextShouldBeginEditing(handle func(o objc.Object, textObject Text) bool) {
	objc.AddMethod(c.class, objc.SEL("textShouldBeginEditing:"), handle)
}

func (c *TextDelegateCreator) SetTextDidBeginEditing(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("textDidBeginEditing:"), handle)
}

func (c *TextDelegateCreator) SetTextShouldEndEditing(handle func(o objc.Object, textObject Text) bool) {
	objc.AddMethod(c.class, objc.SEL("textShouldEndEditing:"), handle)
}

func (c *TextDelegateCreator) SetTextDidEndEditing(handle func(o objc.Object, notification foundation.Notification)) {
	objc.AddMethod(c.class, objc.SEL("textDidEndEditing:"), handle)
}

func (c *TextDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
