package layout

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
)

func AliginAnchors(anchor appkit.ILayoutAnchor, targetAncor appkit.ILayoutAnchor) {
	anchor.ConstraintEqualToAnchor(targetAncor).SetActive(true)
}

func PinAnchorTo(anchor appkit.ILayoutAnchor, targetAncor appkit.ILayoutAnchor, offset float64) {
	anchor.ConstraintEqualToAnchor_Constant(targetAncor, offset).SetActive(true)
}

// PinEdgesToSuperView set view's insets to it's super view.
func PinEdgesToSuperView(view appkit.IView, edgeInsets foundation.EdgeInsets) {
	superView := view.Superview()
	if superView.IsNil() {
		return
	}
	view.TopAnchor().ConstraintEqualToAnchor_Constant(superView.TopAnchor(), edgeInsets.Top).SetActive(true)
	view.RightAnchor().ConstraintEqualToAnchor_Constant(superView.RightAnchor(), -edgeInsets.Right).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor_Constant(superView.BottomAnchor(), -edgeInsets.Bottom).SetActive(true)
	view.LeftAnchor().ConstraintEqualToAnchor_Constant(superView.LeftAnchor(), edgeInsets.Left).SetActive(true)
}

func SetWidth(view appkit.IView, width float64) {
	view.WidthAnchor().ConstraintEqualToConstant(width).SetActive(true)
}

func SetHeight(view appkit.IView, height float64) {
	view.HeightAnchor().ConstraintEqualToConstant(height).SetActive(true)
}

func SetMaxWidth(view appkit.IView, width float64) {
	view.WidthAnchor().ConstraintLessThanOrEqualToConstant(width).SetActive(true)
}

func SetMaxHeight(view appkit.IView, height float64) {
	view.HeightAnchor().ConstraintLessThanOrEqualToConstant(height).SetActive(true)
}

func SetMinWidth(view appkit.IView, width float64) {
	view.WidthAnchor().Init().ConstraintGreaterThanOrEqualToConstant(width).SetActive(true)
}

func SetMinHeight(view appkit.IView, height float64) {
	view.HeightAnchor().ConstraintGreaterThanOrEqualToConstant(height).SetActive(true)
}

func AlignWidthWith(view appkit.IView, targetView appkit.IView) {
	view.WidthAnchor().ConstraintEqualToAnchor(targetView.WidthAnchor()).SetActive(true)
}

func AlginHeightWith(view appkit.IView, targetView appkit.IView) {
	view.HeightAnchor().ConstraintEqualToAnchor(targetView.HeightAnchor()).SetActive(true)
}

// AutoAlignControlsWidth set width anchor for multi controls, using the max width for all controls' width
func AutoAlignControlsWidth(controls ...appkit.IControl) {
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
