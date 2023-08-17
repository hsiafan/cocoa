package main

import (
	"fmt"
	"runtime"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

// NSTableView

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

// https://gist.github.com/wozuo/53a475e67dd11c60cfb1e4f62ea91d32
func initAndRun() {
	app := appkit.ApplicationClass.SharedApplication()
	w := appkit.NewWindowWithSize(600, 400)
	w.SetTitle("Test")

	// text field
	sv := appkit.NewScrollView()
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	sv.SetBackgroundColor(appkit.ColorClass.ClearColor())
	sv.SetDrawsBackground(false)
	w.ContentView().AddSubview(sv)
	w.ContentView().LeadingAnchor().ConstraintEqualToAnchor_Constant(sv.LeadingAnchor(), -10).SetActive(true)
	w.ContentView().TopAnchor().ConstraintEqualToAnchor_Constant(sv.TopAnchor(), -10).SetActive(true)
	w.ContentView().TrailingAnchor().ConstraintEqualToAnchor_Constant(sv.TrailingAnchor(), 10).SetActive(true)
	w.ContentView().BottomAnchor().ConstraintEqualToAnchor_Constant(sv.BottomAnchor(), 10).SetActive(true)

	tv := appkit.NewTableView()
	tv.SetFrame(sv.Bounds())
	tv.SetBackgroundColor(appkit.ColorClass.ClearColor())
	tv.SetDataSource(appkit.WrapTableViewDataSource(&myTableDataSource{}))

	col := appkit.NewTableColumn()
	col.SetTitle("The Column1")
	col.SetMinWidth(20)
	tv.AddTableColumn(col)

	sv.SetDocumentView(tv)
	sv.SetHasHorizontalScroller(false)
	sv.SetHasVerticalScroller(true)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	app.SetDelegate(appkit.WrapApplicationDelegate(&myApplicationDelegate{app: app}))

	app.Run()
}

type myApplicationDelegate struct {
	appkit.ApplicationDelegateBase
	app appkit.Application
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

type myTableDataSource struct {
	appkit.TableViewDataSourceBase
}

func (d *myTableDataSource) ImplementsNumberOfRowsInTableView() bool {
	return true
}

func (d *myTableDataSource) NumberOfRowsInTableView(tableView appkit.TableView) int {
	return 100
}

func (d *myTableDataSource) ImplementsTableView_ObjectValueForTableColumn_Row() bool {
	return true
}

func (d *myTableDataSource) TableView_ObjectValueForTableColumn_Row(tableView appkit.TableView, tableColumn appkit.TableColumn, row int) objc.IObject {
	return foundation.NewString(fmt.Sprintf("%v", row))
}

func main() {
	initAndRun()
}
