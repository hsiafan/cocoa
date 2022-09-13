package main

import (
	"fmt"
	"runtime"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/helper/action"
	"github.com/hsiafan/cocoa/helper/layout"
	"github.com/hsiafan/cocoa/helper/widgets"
	"github.com/hsiafan/cocoa/objc"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.ApplicationClass.SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	w := appkit.NewWindowWithSize(600, 400)
	w.SetTitle("Test Layout")

	label := appkit.NewLabel("label")
	mdButton := appkit.NewButtonWithTitle("modal dialog")
	dButton := appkit.NewButtonWithTitle("dialog")
	textView := appkit.TextViewClass.ScrollableTextView()

	action.Set(mdButton, func(sender objc.IObject) {
		d := widgets.NewDialog(400, 300)
		d.SetView(appkit.NewLabel("test modal dialog"))
		if d.RunModal() == appkit.ModalResponseOK {
			fmt.Println("ok!")
		}
	})

	action.Set(dButton, func(sender objc.IObject) {
		d := widgets.NewDialog(400, 300)
		d.SetView(appkit.NewLabel("test dialog"))
		d.Center()
		d.Show(func() {
			fmt.Println("ok!")
		})
	})

	gridView := appkit.GridViewClass.New()
	for i := 0; i < 3; i++ {
		var views []appkit.IView
		for j := 0; j < 4; j++ {
			label := appkit.NewLabel(fmt.Sprintf("label-%v-%v", i, j))
			views = append(views, label)
		}
		gridView.AddRowWithViews(views)
	}
	gridView.SetContentHuggingPriority_ForOrientation(appkit.LayoutPriorityDefaultHigh, appkit.LayoutConstraintOrientationHorizontal)
	gridView.ColumnAtIndex(0).SetXPlacement(appkit.GridCellPlacementTrailing)
	gridView.SetRowAlignment(appkit.GridRowAlignmentLastBaseline)

	stackView := appkit.StackViewClass.StackViewWithViews([]appkit.IView{label, mdButton, dButton, textView, gridView})
	stackView.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
	stackView.SetDistribution(appkit.StackViewDistributionFillEqually)
	stackView.SetAlignment(appkit.LayoutAttributeCenterX)
	stackView.SetSpacing(10)

	w.ContentView().AddSubview(stackView)
	layout.PinEdgesToSuperView(stackView, foundation.EdgeInsets{Top: 10, Bottom: 10, Left: 20, Right: 20})

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	ad := &appkit.ApplicationDelegateImpl{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})
	ad.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(ad)

	app.Run()
}

func main() {
	initAndRun()

}
