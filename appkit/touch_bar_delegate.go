// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type TouchBarDelegate interface {
	ImplementsTouchBar_MakeItemForIdentifier() bool
	// optional
	TouchBar_MakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem
}

func WrapTouchBarDelegate(v TouchBarDelegate) objc.Object {
	return objc.WrapAsProtocol("NSTouchBarDelegate", v)
}

type TouchBarDelegateBase struct {
}

func (p *TouchBarDelegateBase) ImplementsTouchBar_MakeItemForIdentifier() bool {
	return false
}

func (p *TouchBarDelegateBase) TouchBar_MakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem {
	panic("unimpemented protocol method")
}

type TouchBarDelegateCreator struct {
	className string
	class     objc.Class
}

func NewTouchBarDelegateCreator(name string) *TouchBarDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("ProtocolBase"), name, 0)
	objc.RegisterClassPair(class)
	return &TouchBarDelegateCreator{className: name, class: class}
}

func (c *TouchBarDelegateCreator) SetTouchBar_MakeItemForIdentifier(handle func(o objc.ProtocolBase, touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem) {
	objc.AddMethod(c.class, objc.SEL("touchBar:makeItemForIdentifier:"), handle)
}

func (c *TouchBarDelegateCreator) Create() objc.ProtocolBase {
	return objc.ProtocolBase{Object: c.class.CreateInstance(0)}
}
