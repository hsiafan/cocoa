package coface

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
)

// NewWindow create a common window with close/minimize buttons
func NewWindow(width, height float64) appkit.Window {
	return appkit.WindowClass.Alloc().InitWithContentRect_StyleMask_Backing_Defer(
		foundation.Rect{Size: foundation.Size{Width: width, Height: height}},
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable|appkit.WindowStyleMaskResizable|appkit.WindowStyleMaskMiniaturizable,
		appkit.BackingStoreBuffered,
		false,
	)
}

// NewWindowWithStyle create a common window with styles
func NewWindowWithStyle(width, height float64, style appkit.WindowStyleMask) appkit.Window {
	return appkit.WindowClass.Alloc().InitWithContentRect_StyleMask_Backing_Defer(
		foundation.Rect{Size: foundation.Size{Width: width, Height: height}},
		style,
		appkit.BackingStoreBuffered,
		false,
	)
}
