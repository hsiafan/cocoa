package coface

import "github.com/hsiafan/cocoa/appkit"

// NewVerticalStackView return a new vertical StackView
func NewVerticalStackView() appkit.StackView {
	sv := appkit.StackViewClass.New()
	sv.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return sv
}

// NewHorizontalStackView return a new horizontal StackView
func NewHorizontalStackView() appkit.StackView {
	sv := appkit.StackViewClass.New()
	sv.SetOrientation(appkit.UserInterfaceLayoutOrientationHorizontal)
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return sv
}
