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

type TouchBarDelegateImpl struct {
	_TouchBar_MakeItemForIdentifier func(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem
}

func (di *TouchBarDelegateImpl) ImplementsTouchBar_MakeItemForIdentifier() bool {
	return di._TouchBar_MakeItemForIdentifier != nil
}

func (di *TouchBarDelegateImpl) SetTouchBar_MakeItemForIdentifier(f func(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem) {
	di._TouchBar_MakeItemForIdentifier = f
}

func (di *TouchBarDelegateImpl) TouchBar_MakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem {
	return di._TouchBar_MakeItemForIdentifier(touchBar, identifier)
}

type TouchBarDelegateWrapper struct {
	objc.Object
}

func (t_ *TouchBarDelegateWrapper) ImplementsTouchBar_MakeItemForIdentifier() bool {
	return t_.RespondsToSelector(objc.GetSelector("touchBar:makeItemForIdentifier:"))
}

func (t_ TouchBarDelegateWrapper) TouchBar_MakeItemForIdentifier(touchBar ITouchBar, identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, objc.GetSelector("touchBar:makeItemForIdentifier:"), objc.ExtractPtr(touchBar), identifier)
	return rv
}
