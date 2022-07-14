package coface

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

// Dialog is appkit.Panel with optional OK|CANCEL buttons
type Dialog struct {
	appkit.Panel
	content appkit.View
	ok      appkit.Button
	cancel  appkit.Button
}

// NewDialog create new Dialog
func NewDialog(width, height float64) *Dialog {
	panel := appkit.PanelClass.New()
	panel.SetFrame_Display(foundation.Rect{Size: foundation.Size{Width: width, Height: height}}, true)

	contentView := panel.ContentView()
	view := appkit.ViewClass.New()
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	ok := NewButton("OK")
	cancel := NewButton("Cancel")
	contentView.AddSubview(view)
	contentView.AddSubview(ok)
	contentView.AddSubview(cancel)

	ok.BottomAnchor().ConstraintEqualToAnchor_Constant(contentView.BottomAnchor(), -10).SetActive(true)
	ok.RightAnchor().ConstraintEqualToAnchor_Constant(contentView.RightAnchor(), -15).SetActive(true)
	cancel.BottomAnchor().ConstraintEqualToAnchor(ok.BottomAnchor()).SetActive(true)
	cancel.RightAnchor().ConstraintEqualToAnchor_Constant(ok.LeftAnchor(), -10).SetActive(true)
	ok.WidthAnchor().ConstraintEqualToAnchor(cancel.WidthAnchor()).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor_Constant(ok.TopAnchor(), -10).SetActive(true)
	view.LeftAnchor().ConstraintEqualToAnchor_Constant(contentView.LeftAnchor(), 10).SetActive(true)
	view.TopAnchor().ConstraintEqualToAnchor_Constant(contentView.TopAnchor(), 10).SetActive(true)
	view.RightAnchor().ConstraintEqualToAnchor_Constant(contentView.RightAnchor(), -10).SetActive(true)

	return &Dialog{
		Panel:   panel,
		content: view,
		ok:      ok,
		cancel:  cancel,
	}
}

// SetView set inner content view for Dialog
func (d *Dialog) SetView(view appkit.IView) {
	d.content.AddSubview(view)
	view.LeftAnchor().ConstraintEqualToAnchor(d.content.LeftAnchor()).SetActive(true)
	view.TopAnchor().ConstraintEqualToAnchor(d.content.TopAnchor()).SetActive(true)
	view.RightAnchor().ConstraintEqualToAnchor(d.content.RightAnchor()).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor(d.content.BottomAnchor()).SetActive(true)
}

// Show display dialog in non-modal mode
func (d *Dialog) Show(handle func()) {
	objc.SetAction(d.ok, func(sender objc.IObject) {
		handle()
		d.Close()
	})

	objc.SetAction(d.cancel, func(sender objc.IObject) {
		d.Close()
	})

	d.MakeKeyAndOrderFront(d.Panel)
}

// RunModal display dialog in modal mode
func (d *Dialog) RunModal() appkit.ModalResponse {
	app := appkit.ApplicationClass.SharedApplication()

	objc.SetAction(d.ok, func(sender objc.IObject) {
		app.StopModalWithCode(appkit.ModalResponseOK)
		d.Close()
	})

	objc.SetAction(d.cancel, func(sender objc.IObject) {
		app.StopModalWithCode(appkit.ModalResponseCancel)
		d.Close()
	})

	dialogDelegate := &appkit.WindowDelegateImpl{}
	dialogDelegate.SetWindowShouldClose(func(sender appkit.Window) bool {
		app.StopModalWithCode(appkit.ModalResponseCancel)
		return true
	})

	d.SetDelegate(dialogDelegate)

	return app.RunModalForWindow(d)
}
