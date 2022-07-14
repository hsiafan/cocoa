package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/coface"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/webkit"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	var html = "<html><h1>Hello World!</h1></html>"
	var url = "https://www.test.com"

	app := appkit.ApplicationClass.SharedApplication()
	w := coface.NewWindow(600, 400)
	w.SetTitle("Test widgets")

	sv := coface.NewVerticalStackView()
	w.SetContentView(sv)

	webView := webkit.WebViewClass.New()
	webView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	webView.LoadHTMLString_BaseURL(html, foundation.URLClass.URLWithString(url))
	sv.AddView_InGravity(webView, appkit.StackViewGravityTop)

	snapshotButton := coface.NewButton("capture")

	snapshotWin := coface.NewWindow(0, 0)
	snapshotWin.SetTitle("Test widgets")

	snapshotWebView := webkit.WebViewClass.New()
	snapshotWebView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	snapshotWin.SetContentView(snapshotWebView)

	navigationDelegate := &webkit.NavigationDelegateImpl{}
	navigationDelegate.SetWebView_DidFinishNavigation(func(webView webkit.WebView, navigation webkit.Navigation) {
		script := `var rect = {"width":document.body.scrollWidth, "height":document.body.scrollHeight}; rect`
		webView.EvaluateJavaScript_CompletionHandler(script, func(value objc.Object, err foundation.Error) {
			rect := foundation.DictToMap[string, foundation.Number](foundation.MakeDictionary(value.Ptr()))
			width := rect["width"].DoubleValue()
			height := rect["height"].DoubleValue()
			snapshotWin.SetFrame_Display(foundation.Rect{Size: foundation.Size{Width: width, Height: height}}, true)
			snapshotWebView.LoadHTMLString_BaseURL(html, foundation.URLClass.URLWithString(url))
		})
	})
	webView.SetNavigationDelegate(navigationDelegate)

	ssnd := &webkit.NavigationDelegateImpl{}
	ssnd.SetWebView_DidFinishNavigation(func(webView webkit.WebView, navigation webkit.Navigation) {
		snapshotButton.SetEnabled(true)
	})
	snapshotWebView.SetNavigationDelegate(ssnd)

	objc.SetAction(snapshotButton, func(sender objc.IObject) {
		snapshotWebView.TakeSnapshotWithConfiguration_CompletionHandler(nil, func(image appkit.Image, err foundation.Error) {
			imageRef := image.CGImageForProposedRect_Context_Hints(nil, nil, nil)
			imageRepo := appkit.BitmapImageRepClass.Alloc().InitWithCGImage(imageRef)
			imageRepo.SetSize(image.Size())
			pngData := imageRepo.RepresentationUsingType_Properties(appkit.BitmapImageFileTypePNG, nil)

			if err := os.WriteFile("webview_screenshot.png", pngData, os.ModePerm); err != nil {
				fmt.Println("write image to file error:", err)
			} else {
				fmt.Println("image captured to webview_screenshot.png")
			}
		})
	})
	snapshotButton.SetEnabled(false)

	sv.AddView_InGravity(snapshotButton, appkit.StackViewGravityTop)

	wd := &appkit.WindowDelegateImpl{}
	wd.SetWindowWillClose(func(notification foundation.Notification) {
		snapshotWin.Close()
	})
	w.SetDelegate(wd)

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
