package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/helper/action"
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
	w := appkit.NewWindowWithSize(600, 400)
	w.SetTitle("Test widgets")

	sv := appkit.NewVerticalStackView()
	w.SetContentView(sv)

	webView := webkit.WebViewClass.New()
	webView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	webView.LoadHTMLString_BaseURL(html, foundation.URLClass.URLWithString(url))
	sv.AddView_InGravity(webView, appkit.StackViewGravityTop)

	snapshotButton := appkit.NewButtonWithTitle("capture")

	snapshotWin := appkit.NewWindowWithSize(0, 0)
	snapshotWin.SetTitle("Test widgets")

	snapshotWebView := webkit.WebViewClass.New()
	snapshotWebView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	snapshotWin.SetContentView(snapshotWebView)

	webView.SetNavigationDelegate(webkit.WrapNavigationDelegate(&myNavigationDelegate{
		onFinishNavigation: func(webView webkit.WebView, navigation webkit.Navigation) {
			script := `var rect = {"width":document.body.scrollWidth, "height":document.body.scrollHeight}; rect`
			webView.EvaluateJavaScript_CompletionHandler(script, func(value objc.Object, err foundation.Error) {
				rect := foundation.DictToMap[string, foundation.Number](foundation.MakeDictionary(value.Ptr()))
				width := rect["width"].DoubleValue()
				height := rect["height"].DoubleValue()
				snapshotWin.SetFrame_Display(foundation.Rect{Size: foundation.Size{Width: width, Height: height}}, true)
				snapshotWebView.LoadHTMLString_BaseURL(html, foundation.URLClass.URLWithString(url))
			})
		},
	}))

	snapshotWebView.SetNavigationDelegate(webkit.WrapNavigationDelegate(&myNavigationDelegate{
		onFinishNavigation: func(webView webkit.WebView, navigation webkit.Navigation) {
			snapshotButton.SetEnabled(true)
		},
	}))

	action.Set(snapshotButton, func(sender objc.IObject) {
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
	w.SetDelegate(appkit.WrapWindowDelegate(&myWindowDelegate{snapshotWin: snapshotWin}))

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	app.SetDelegate(appkit.WrapApplicationDelegate(&myApplicationDelegate{app: app}))

	app.Run()
}

type myNavigationDelegate struct {
	webkit.NavigationDelegateBase
	onFinishNavigation func(webView webkit.WebView, navigation webkit.Navigation)
}

func (p *myNavigationDelegate) ImplementsWebView_DidFinishNavigation() bool {
	return true
}

func (p *myNavigationDelegate) WebView_DidFinishNavigation(webView webkit.WebView, navigation webkit.Navigation) {
	p.onFinishNavigation(webView, navigation)
}

type myWindowDelegate struct {
	appkit.WindowDelegateBase
	snapshotWin appkit.Window
}

func (p *myWindowDelegate) ImplementsWindowWillClose() bool {
	return true
}

func (p *myWindowDelegate) WindowWillClose(notification foundation.Notification) {
	p.snapshotWin.Close()
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

func main() {
	initAndRun()
}
