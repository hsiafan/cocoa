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

type TouchBarDelegateWrapper struct {
	objc.Object
}

func (t_ TouchBarDelegateWrapper) TouchBar_MakeItemForIdentifier(touchBar ITouchBar, identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, objc.GetSelector("touchBar:makeItemForIdentifier:"), objc.ExtractPtr(touchBar), identifier)
	return rv
}
