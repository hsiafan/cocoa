package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/hsiafan/cocoa/dispatch"
	"github.com/hsiafan/cocoa/helper/action"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.ApplicationClass.SharedApplication()
	w := appkit.NewWindowWithSize(600, 400)

	w.SetTitle("Test widgets")

	filePathField := appkit.TextFieldClass.Alloc().InitWithFrame(rectOf(10, 330, 200, 20))
	cell := filePathField.Cell()
	cell.SetWraps(false)
	cell.SetScrollable(true)
	filePathField.SetEditable(false)
	w.ContentView().AddSubview(filePathField)

	saveButton := appkit.NewButtonWithTitle("Save...")
	saveButton.SetFrame(rectOf(250, 330, 80, 20))
	action.Set(saveButton, func(sender objc.IObject) {
		savePanel := appkit.SavePanelClass.New()
		// if savePanel.RunModal() == appkit.ModalResponseOK {
		// 	filePathField.SetStringValue(savePanel.URL().Path())
		// }
		savePanel.BeginWithCompletionHandler(func(response appkit.ModalResponse) {
			if response == appkit.ModalResponseOK {
				filePathField.SetStringValue(savePanel.URL().Path())
			}
		})
	})
	w.ContentView().AddSubview(saveButton)

	presentationTF := appkit.TextFieldClass.Alloc().InitWithFrame(rectOf(10, 290, 100, 25))
	w.ContentView().AddSubview(presentationTF)

	stepper := appkit.StepperClass.Alloc().InitWithFrame(rectOf(130, 290, 16, 25))
	stepper.SetDoubleValue(100)
	w.ContentView().AddSubview(stepper)

	colorWell := appkit.ColorWellClass.Alloc().InitWithFrame(rectOf(160, 290, 30, 25))
	w.ContentView().AddSubview(colorWell)

	comboBox := appkit.ComboBoxClass.Alloc().InitWithFrame(rectOf(210, 290, 100, 25))
	comboBox.AddItemsWithObjectValues([]objc.IObject{
		foundation.NewString("Test1"),
		foundation.NewString("Test2"),
	})
	comboBox.SelectItemAtIndex(0)
	w.ContentView().AddSubview(comboBox)

	slider := appkit.SliderClass.Alloc().InitWithFrame(rectOf(330, 290, 100, 25))
	action.Set(slider, func(sender objc.IObject) {
		presentationTF.SetDoubleValue(slider.DoubleValue())
	})
	slider.SetMaxValue(10)
	w.ContentView().AddSubview(slider)

	datePicker := appkit.DatePickerClass.Alloc().InitWithFrame(rectOf(450, 290, 140, 25))
	w.ContentView().AddSubview(datePicker)

	// buttons
	cb := appkit.NewCheckBox("check box")
	cb.SetFrame(rectOf(10, 250, 80, 25))
	w.ContentView().AddSubview(cb)

	rb := appkit.NewRadioButton("radio button")
	rb.SetFrame(rectOf(150, 250, 120, 25))
	w.ContentView().AddSubview(rb)

	sw := appkit.SwitchClass.Alloc().InitWithFrame(rectOf(260, 250, 120, 25))
	w.ContentView().AddSubview(sw)

	li := appkit.LevelIndicatorClass.Alloc().InitWithFrame(rectOf(370, 250, 120, 25))
	li.SetCriticalValue(4)
	li.SetDoubleValue(3)
	w.ContentView().AddSubview(li)

	btn := appkit.NewButtonWithTitle("change color")
	btn.SetFrame(rectOf(10, 160, 120, 25))
	w.ContentView().AddSubview(btn)

	quitBtn := appkit.NewButtonWithTitle("Quit")
	quitBtn.SetFrame(rectOf(10, 130, 80, 25))
	action.Set(quitBtn, func(sender objc.IObject) {
		app.Terminate(nil)
	})
	w.ContentView().AddSubview(quitBtn)

	// text field
	tf := appkit.NewTextField()
	w.ContentView().AddSubview(tf)
	tf.SetFrame(rectOf(10, 100, 150, 25))

	// label
	label := appkit.NewLabel("")
	label.SetFrame(rectOf(170, 100, 150, 25))
	w.ContentView().AddSubview(label)
	tf.SetDelegate(appkit.WrapTextFieldDelegate(&myTextFieldDelegate{label: label, tf: tf}))
	action.Set(btn, func(sender objc.IObject) {
		label.SetTextColor(appkit.ColorClass.RedColor())
	})

	// password
	stf := appkit.NewSecureTextField()
	stf.SetFrame(rectOf(340, 100, 150, 25))
	stf.SetDelegate(appkit.WrapTextFieldDelegate(&myTextFieldDelegate{label: label, tf: stf}))
	w.ContentView().AddSubview(stf)

	// progress indicator
	indicator := appkit.ProgressIndicatorClass.Alloc().InitWithFrame(rectOf(10, 70, 350, 25))
	indicator.SetMinValue(0)
	indicator.SetMaxValue(1)
	indicator.SetIndeterminate(false)
	indicator.SetDisplayedWhenStopped(false)
	w.ContentView().AddSubview(indicator)
	go func() {
		dispatch.GetMainQueue().DispatchAsync(func() {
			indicator.StartAnimation(indicator)
		})
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			dispatch.GetMainQueue().DispatchAsync(func() {
				indicator.SetDoubleValue(0.1 * float64(i))
			})
		}
		dispatch.GetMainQueue().DispatchAsync(func() {
			indicator.StopAnimation(indicator)
		})
	}()

	// text view & scroll view
	sv := appkit.TextViewClass.ScrollableTextView()
	sv.SetFrame(rectOf(10, 200, 200, 30))
	appkit.MakeTextView(sv.DocumentView().Ptr()).SetAllowsUndo(true)
	w.ContentView().AddSubview(sv)

	sv2 := appkit.TextViewClass.ScrollableTextView()
	sv2.SetFrame(rectOf(250, 200, 200, 30))
	appkit.MakeTextView(sv2.DocumentView().Ptr()).SetAllowsUndo(true)
	w.ContentView().AddSubview(sv2)

	wdCreator := appkit.NewWindowDelegateCreator("MyWindowDelegate")
	wdCreator.SetWindowDidMove(func(o objc.Object, notification foundation.Notification) {
		frame := w.Frame()
		origin := frame.Origin
		fmt.Println("window move to ", origin.X, origin.Y)
	})
	w.SetDelegate(wdCreator.Create())
	w.Center()
	w.MakeKeyAndOrderFront(nil)

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

type myTextFieldDelegate struct {
	appkit.TextFieldDelegateBase
	tf    appkit.ITextField
	label appkit.TextField
}

func (p *myTextFieldDelegate) ImplementsControlTextDidChange() bool {
	return true
}

func (p *myTextFieldDelegate) ControlTextDidChange(obj foundation.Notification) {
	dispatch.GetMainQueue().DispatchAsync(func() {
		p.label.SetStringValue(p.tf.StringValue())
	})
}

func main() {
	initAndRun()
}

func rectOf(x, y, width, height float64) foundation.Rect {
	return foundation.Rect{Origin: foundation.Point{X: x, Y: y}, Size: foundation.Size{Width: width, Height: height}}
}
