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

	app.SetDelegate(appkit.WrapApplicationDelegate(&myApplicationDelegate{app: app, w: w}))

	w.Center()
	w.MakeKeyAndOrderFront(nil)
	app.Run()
}

type myApplicationDelegate struct {
	appkit.ApplicationDelegateBase
	app appkit.Application
	w   appkit.Window
}

func (p *myApplicationDelegate) ImplementsApplicationDidFinishLaunching() bool {
	return true
}

func (p *myApplicationDelegate) ApplicationDidFinishLaunching(notification foundation.Notification) {
	p.app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	p.app.ActivateIgnoringOtherApps(true)
}

func (p *myApplicationDelegate) ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool {
	return true
}

func (p *myApplicationDelegate) ApplicationShouldTerminateAfterLastWindowClosed(sender appkit.Application) bool {
	return true
}

func (p *myApplicationDelegate) ImplementsApplicationWillFinishLaunching() bool {
	return true
}

func (p *myApplicationDelegate) ApplicationWillFinishLaunching(notification foundation.Notification) {
	p.w.SetFrameAutosaveName("tab-test")
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
