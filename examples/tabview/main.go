package main

import (
	"runtime"
	"strconv"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.ApplicationClass.SharedApplication()
	w := appkit.NewWindowWithSize(720, 440)
	w.SetTitle("Decoder")

	tabView := appkit.TabViewClass.New()
	tabView.SetTranslatesAutoresizingMaskIntoConstraints(false)

	// add tabs
	tabView.AddTabViewItem(createNewView(1))
	tabView.AddTabViewItem(createNewView(2))

	w.SetContentView(tabView)

	ad := &appkit.ApplicationDelegateImpl{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})
	ad.SetApplicationWillFinishLaunching(func(foundation.Notification) {
		w.SetFrameAutosaveName("tab-test")
	})
	ad.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(ad)

	w.Center()
	w.MakeKeyAndOrderFront(nil)
	app.Run()
}

func createNewView(idx int) appkit.ITabViewItem {
	sv := appkit.StackViewClass.New()
	sv.SetTranslatesAutoresizingMaskIntoConstraints(true)
	sv.AddView_InGravity(appkit.NewButtonWithTitle("button"), appkit.StackViewGravityTop)
	sv.AddView_InGravity(appkit.TextFieldClass.New(), appkit.StackViewGravityTop)
	ti := appkit.TabViewItemClass.New()
	ti.SetLabel("Tab" + strconv.Itoa(idx))
	ti.SetView(sv)
	return ti
}

func main() {
	initAndRun()
}
