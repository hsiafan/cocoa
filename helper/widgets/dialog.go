package widgets

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/helper/action"
	"github.com/hsiafan/cocoa/helper/layout"
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
	return NewDialogWithEdgeInsets(width, height, foundation.EdgeInsets{
		Top:    10,
		Left:   10,
		Bottom: 10,
		Right:  10,
	})
}

// NewDialogWithEdgeInsets create new Dialog width edgeInsets
func NewDialogWithEdgeInsets(width, height float64, edgeInsets foundation.EdgeInsets) *Dialog {
	panel := appkit.PanelClass.New()
	panel.SetFrame_Display(foundation.Rect{Size: foundation.Size{Width: width, Height: height}}, true)

	contentView := panel.ContentView()
	view := appkit.ViewClass.New()
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	ok := appkit.NewButtonWithTitle("OK")
	ok.SetTranslatesAutoresizingMaskIntoConstraints(false)
	cancel := appkit.NewButtonWithTitle("Cancel")
	cancel.SetTranslatesAutoresizingMaskIntoConstraints(false)
	contentView.AddSubview(view)
	contentView.AddSubview(ok)
	contentView.AddSubview(cancel)

	layout.PinAnchorTo(ok.BottomAnchor(), contentView.BottomAnchor(), -edgeInsets.Bottom)
	layout.PinAnchorTo(ok.RightAnchor(), contentView.RightAnchor(), -edgeInsets.Right)
	layout.PinAnchorTo(cancel.BottomAnchor(), contentView.BottomAnchor(), -edgeInsets.Bottom)
	layout.PinAnchorTo(cancel.RightAnchor(), ok.LeftAnchor(), -10)
	layout.AlignWidth(ok, cancel)
	layout.PinAnchorTo(view.BottomAnchor(), ok.TopAnchor(), -10)
	layout.PinAnchorTo(view.LeftAnchor(), contentView.LeftAnchor(), edgeInsets.Left)
	layout.PinAnchorTo(view.TopAnchor(), contentView.TopAnchor(), edgeInsets.Top)
	layout.PinAnchorTo(view.RightAnchor(), contentView.RightAnchor(), -edgeInsets.Right)

	return &Dialog{
		Panel:   panel,
		content: view,
		ok:      ok,
		cancel:  cancel,
	}
}

// SetView set inner content view for Dialog
func (d *Dialog) SetView(view appkit.IView) {
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	d.content.AddSubview(view)
	view.LeftAnchor().ConstraintEqualToAnchor(d.content.LeftAnchor()).SetActive(true)
	view.TopAnchor().ConstraintEqualToAnchor(d.content.TopAnchor()).SetActive(true)
	view.RightAnchor().ConstraintEqualToAnchor(d.content.RightAnchor()).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor(d.content.BottomAnchor()).SetActive(true)
}

// Show display dialog in non-modal mode
func (d *Dialog) Show(handle func()) {
	action.Set(d.ok, func(sender objc.IObject) {
		handle()
		d.Close()
	})

	action.Set(d.cancel, func(sender objc.IObject) {
		d.Close()
	})

	d.MakeKeyAndOrderFront(d.Panel)
}

// RunModal display dialog in modal mode
func (d *Dialog) RunModal() appkit.ModalResponse {
	app := appkit.ApplicationClass.SharedApplication()

	action.Set(d.ok, func(sender objc.IObject) {
		app.StopModalWithCode(appkit.ModalResponseOK)
		d.Close()
	})

	action.Set(d.cancel, func(sender objc.IObject) {
		app.StopModalWithCode(appkit.ModalResponseCancel)
		d.Close()
	})

	dwdo := appkit.WrapWindowDelegate(&DialogWindowDelegate{app: app})
	d.SetDelegate(dwdo)
	objc.SetRetainDelegateAssociated(d, dwdo)

	return app.RunModalForWindow(d)
}

type DialogWindowDelegate struct {
	appkit.WindowDelegateBase
	app appkit.Application
}

func (p *DialogWindowDelegate) ImplementsWindowShouldClose() bool {
	return true
}

func (p *DialogWindowDelegate) WindowShouldClose(sender appkit.Window) bool {
	p.app.StopModalWithCode(appkit.ModalResponseCancel)
	return true
}
