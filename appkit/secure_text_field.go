// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var SecureTextFieldClass = _SecureTextFieldClass{objc.GetClass("NSSecureTextField")}

type _SecureTextFieldClass struct {
	objc.Class
}

type ISecureTextField interface {
	ITextField
}

type SecureTextField struct {
	TextField
}

func MakeSecureTextField(ptr unsafe.Pointer) SecureTextField {
	return SecureTextField{
		TextField: MakeTextField(ptr),
	}
}

func (sc _SecureTextFieldClass) LabelWithAttributedString(attributedStringValue foundation.IAttributedString) SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, "labelWithAttributedString:", attributedStringValue)
	return rv
}

func (sc _SecureTextFieldClass) LabelWithString(stringValue string) SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, "labelWithString:", stringValue)
	return rv
}

func (sc _SecureTextFieldClass) TextFieldWithString(stringValue string) SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, "textFieldWithString:", stringValue)
	return rv
}

func (sc _SecureTextFieldClass) WrappingLabelWithString(stringValue string) SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, "wrappingLabelWithString:", stringValue)
	return rv
}

func (s_ SecureTextField) InitWithFrame(frameRect foundation.Rect) SecureTextField {
	rv := objc.CallMethod[SecureTextField](s_, "initWithFrame:", frameRect)
	return rv
}

func (s_ SecureTextField) Init() SecureTextField {
	rv := objc.CallMethod[SecureTextField](s_, "init")
	return rv
}

func (sc _SecureTextFieldClass) Alloc() SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, "alloc")
	return rv
}

func (sc _SecureTextFieldClass) New() SecureTextField {
	rv := objc.CallMethod[SecureTextField](sc, "new")
	rv.Autorelease()
	return rv
}

func NewSecureTextField() SecureTextField {
	return SecureTextFieldClass.New()
}
