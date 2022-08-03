package layout

import "github.com/hsiafan/cocoa/appkit"

// SetMargin add layout sub View with paddings.
// The padding value sequences by top, right, bottom, left, the right/bottom/left values can be omitted.
func SetMargin(parentView appkit.IView, view appkit.IView, top float64, values ...float64) {
	if len(values) > 3 {
		panic("too many values")
	}

	view.TopAnchor().ConstraintEqualToAnchor_Constant(parentView.TopAnchor(), top).SetActive(true)
	var right, bottom, left = top, top, top
	if len(values) >= 1 {
		right = values[0]
		left = right
		if len(values) >= 2 {
			bottom = values[1]
		}
		if len(values) >= 3 {
			left = values[2]
		}
	}

	view.RightAnchor().ConstraintEqualToAnchor_Constant(parentView.RightAnchor(), -right).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor_Constant(parentView.BottomAnchor(), -bottom).SetActive(true)
	view.LeftAnchor().ConstraintEqualToAnchor_Constant(parentView.LeftAnchor(), left).SetActive(true)
}

// UseSameMaxWidth set width anchor for multi controls, using the max width for all controls' width
func UseSameMaxWidth(controls ...appkit.IControl) {
	var maxWidth float64
	for _, control := range controls {
		control.SizeToFit()
		w := control.Bounds().Size.Width
		if maxWidth < w {
			maxWidth = w
		}
	}
	for _, control := range controls {
		control.WidthAnchor().ConstraintEqualToConstant(maxWidth).SetActive(true)
	}
}
