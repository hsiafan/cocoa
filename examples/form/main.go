package main

import (
	"runtime"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/coface"
	"github.com/hsiafan/cocoa/foundation"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.ApplicationClass.SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	w := coface.NewWindow(600, 400)
	w.SetTitle("Form")

	fv := coface.NewFormView()
	fv.SetLabelWidth(100)
	fv.SetLabelAlignment(coface.LabelAlignmentTrailing)
	fv.SetLabelControlSpacing(10)
	fv.AddRow("user", coface.NewTextField())
	fv.SetLabelFont(appkit.FontClass.UserFontOfSize(13))
	fv.AddRow("password", coface.NewSecureTextField())
	cb := coface.NewCheckBox("")
	fv.AddRow("males", cb)
	fv.AddExpandRow()

	w.ContentView().AddSubview(fv)
	fv.LeftAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().LeftAnchor(), 10).SetActive(true)
	fv.TopAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().TopAnchor(), 10).SetActive(true)
	fv.RightAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().RightAnchor(), -10).SetActive(true)
	fv.BottomAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().BottomAnchor(), -10).SetActive(true)

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
