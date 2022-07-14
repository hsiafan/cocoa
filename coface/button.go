package coface

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/objc"
)

// NewButton return a new common used Button
func NewButton(title string) appkit.Button {
	btn := appkit.ButtonClass.ButtonWithTitle_Target_Action(title, objc.Object{}, objc.Selector{})
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return btn
}

// NewButtonWithImage return a new common used Button with image
func NewButtonWithImage(image appkit.Image) appkit.Button {
	btn := appkit.ButtonClass.ButtonWithImage_Target_Action(image, objc.Object{}, objc.Selector{})
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return btn
}

// NewCheckBox return a new common used checkbox Button
func NewCheckBox(title string) appkit.Button {
	btn := appkit.ButtonClass.CheckboxWithTitle_Target_Action(title, objc.Object{}, objc.Selector{})
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return btn
}

// NewRadioButton return a new common used radio Button
func NewRadioButton(title string) appkit.Button {
	btn := appkit.ButtonClass.RadioButtonWithTitle_Target_Action(title, objc.Object{}, objc.Selector{})
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return btn
}

// NewPushButton return a button that switches between on and off states with each click.
func NewPushButton(title string) appkit.Button {
	btn := appkit.ButtonClass.New()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetButtonType(appkit.ButtonTypePushOnPushOff)
	btn.SetTitle(title)
	return btn
}
