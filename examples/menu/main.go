package main

import (
	"runtime"

	"github.com/hsiafan/cocoa/coface"
	"github.com/hsiafan/cocoa/objc"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
)

// menus and tray

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.ApplicationClass.SharedApplication()
	w := coface.NewWindow(600, 400)
	w.SetTitle("Test")

	// text field
	textView := appkit.TextViewClass.ScrollableTextView()
	textView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tv := appkit.MakeTextView(textView.DocumentView().Ptr())
	tv.SetAllowsUndo(true)
	tv.SetRichText(false)
	w.ContentView().AddSubview(textView)
	w.ContentView().LeadingAnchor().ConstraintEqualToAnchor_Constant(textView.LeadingAnchor(), -10).SetActive(true)
	w.ContentView().TopAnchor().ConstraintEqualToAnchor_Constant(textView.TopAnchor(), -10).SetActive(true)
	w.ContentView().TrailingAnchor().ConstraintEqualToAnchor_Constant(textView.TrailingAnchor(), 10).SetActive(true)
	w.ContentView().BottomAnchor().ConstraintEqualToAnchor_Constant(textView.BottomAnchor(), 10).SetActive(true)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	ad := &appkit.ApplicationDelegateImpl{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})
	ad.SetApplicationWillFinishLaunching(func(foundation.Notification) {
		// should set menu bar at ApplicationWillFinishLaunching
		setMainMenu(app)
	})
	ad.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(ad)

	// should set system bar after window show
	setSystemBar(app)

	app.Run()
}

func setMainMenu(app appkit.Application) {
	menu := appkit.MenuClass.Alloc().InitWithTitle("main")
	app.SetMainMenu(menu)

	mainMenuItem := coface.NewMenuItem("", "", objc.Selector{})
	mainMenuMenu := appkit.MenuClass.Alloc().InitWithTitle("App")
	mainMenuMenu.AddItem(coface.NewMenuItemWithAction("Hide", "h", app.Hide))
	mainMenuMenu.AddItem(coface.NewMenuItemWithAction("Quit", "q", app.Terminate))
	mainMenuItem.SetSubmenu(mainMenuMenu)
	menu.AddItem(mainMenuItem)

	testMenuItem := coface.NewMenuItem("", "", objc.Selector{})
	testMenu := appkit.MenuClass.Alloc().InitWithTitle("Edit")
	testMenu.AddItem(coface.NewMenuItem("Select All", "a", objc.GetSelector("selectAll:")))
	testMenu.AddItem(appkit.MenuItemClass.SeparatorItem())
	testMenu.AddItem(coface.NewMenuItem("Copy", "c", objc.GetSelector("copy:")))
	testMenu.AddItem(coface.NewMenuItem("Paste", "v", objc.GetSelector("paste:")))
	testMenu.AddItem(coface.NewMenuItem("Cut", "x", objc.GetSelector("cut:")))
	testMenu.AddItem(coface.NewMenuItem("Undo", "z", objc.GetSelector("undo:")))
	testMenu.AddItem(coface.NewMenuItem("Redo", "Z", objc.GetSelector("redo:")))
	testMenuItem.SetSubmenu(testMenu)
	menu.AddItem(testMenuItem)
}

func setSystemBar(app appkit.Application) {
	bar := appkit.StatusBarClass.SystemStatusBar()
	item := bar.StatusItemWithLength(appkit.VariableStatusItemLength)
	button := item.Button()
	button.SetTitle("TestTray")

	menu := appkit.MenuClass.Alloc().InitWithTitle("main")
	menu.AddItem(coface.NewMenuItemWithAction("Hide", "h", app.Hide))
	menu.AddItem(coface.NewMenuItemWithAction("Quit", "q", app.Terminate))
	item.SetMenu(menu)
}

func main() {
	initAndRun()
}
