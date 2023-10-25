package main

import (
	"runtime"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
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
	w.SetTitle("Form")

	fv := widgets.NewFormView()
	fv.SetLabelWidth(100)
	fv.SetLabelAlignment(widgets.LabelAlignmentTrailing)
	fv.SetLabelControlSpacing(10)
	fv.AddRow("user", appkit.NewTextField())
	fv.SetLabelFont(appkit.FontClass.UserFontOfSize(13))
	fv.AddRow("password", appkit.NewSecureTextField())
	cb := appkit.NewCheckBox("")
	fv.AddRow("males", cb)
	fv.AddExpandRow()

	w.ContentView().AddSubview(fv)
	layout.PinEdgesToSuperView(fv, foundation.EdgeInsets{Top: 10, Left: 10, Bottom: 10, Right: 10})

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	creator := appkit.NewApplicationDelegateCreator("MyApplicationDelegate")
	creator.SetApplicationDidFinishLaunching(func(o objc.Object, notification foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})
	creator.SetApplicationShouldTerminateAfterLastWindowClosed(func(o objc.Object, sender appkit.Application) bool {
		return true
	})
	app.SetDelegate(creator.Create())

	app.Run()
}

func main() {
	initAndRun()
}
