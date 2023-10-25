package main

import (
	"runtime"

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
	w := appkit.NewWindowWithSize(600, 400)
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

	creator := appkit.NewApplicationDelegateCreator("MyApplicationDelegate")
	creator.SetApplicationWillFinishLaunching(func(o objc.Object, notification foundation.Notification) {
		setMainMenu(app)
	})
	creator.SetApplicationDidFinishLaunching(func(o objc.Object, notification foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})
	creator.SetApplicationShouldTerminateAfterLastWindowClosed(func(o objc.Object, sender appkit.Application) bool {
		return true
	})
	app.SetDelegate(creator.Create())

	// should set system bar after window show
	setSystemBar(app)

	app.Run()
}

func setMainMenu(app appkit.Application) {
	menu := appkit.MenuClass.Alloc().InitWithTitle("main")
	app.SetMainMenu(menu)

	mainMenu := appkit.NewMenu()
	mainMenu.SetTitle("App")
	mainMenu.AddItem(appkit.NewMenuItemWithAction("Hide Test", "h", app.Hide))
	hideOthersItem := appkit.NewMenuItemWithAction("Hide Others", "h", app.HideOtherApplications)
	hideOthersItem.SetKeyEquivalentModifierMask(appkit.EventModifierFlagOption | appkit.EventModifierFlagCommand)
	mainMenu.AddItem(hideOthersItem)
	mainMenu.AddItem(appkit.NewMenuItemWithAction("Show All", "", app.UnhideAllApplications))
	mainMenu.AddItem(appkit.MenuItemClass.SeparatorItem())
	mainMenu.AddItem(appkit.NewMenuItemWithAction("Quit", "q", app.Terminate))
	menu.AddItem(appkit.NewSubMenuItem(mainMenu))

	testMenu := appkit.MenuClass.Alloc().InitWithTitle("Edit")
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Select All", "a", objc.SEL("selectAll:")))
	testMenu.AddItem(appkit.MenuItemClass.SeparatorItem())
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Copy", "c", objc.SEL("copy:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Paste", "v", objc.SEL("paste:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Cut", "x", objc.SEL("cut:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Undo", "z", objc.SEL("undo:")))
	testMenu.AddItem(appkit.NewMenuItemWithSelector("Redo", "Z", objc.SEL("redo:")))
	menu.AddItem(appkit.NewSubMenuItem(testMenu))
}

func setSystemBar(app appkit.Application) {
	bar := appkit.StatusBarClass.SystemStatusBar()
	item := bar.StatusItemWithLength(appkit.VariableStatusItemLength)
	button := item.Button()
	button.SetTitle("TestTray")

	menu := appkit.MenuClass.Alloc().InitWithTitle("main")
	menu.AddItem(appkit.NewMenuItemWithAction("Hide", "h", app.Hide))
	menu.AddItem(appkit.NewMenuItemWithAction("Quit", "q", app.Terminate))
	item.SetMenu(menu)
}

func main() {
	initAndRun()
}
