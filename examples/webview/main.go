package main

import (
	"embed"
	"io/fs"
	"runtime"
	"time"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/webkit"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	app := appkit.ApplicationClass.SharedApplication()
	w := appkit.NewWindowWithSize(600, 400)
	w.SetTitle("Test FS WebView")

	_fs, _ := fs.Sub(assetsFS, "assets")

	configuration := webkit.NewWebViewConfiguration()
	configuration.Preferences().SetJavaScriptEnabled(true)
	gofsHandler := &webkit.FileSystemURLSchemeHandler{FS: _fs}
	configuration.SetURLSchemeHandler_ForURLScheme(webkit.WrapURLSchemeHandler(gofsHandler), "gofs")

	view := webkit.WebViewClass.Alloc().InitWithFrame_Configuration(foundation.Rect{}, configuration)
	webkit.AddScriptMessageHandlerWithReply(view, "greet", func(message objc.Object) (objc.IObject, error) {
		param := message.Description()
		return foundation.NewString("hello: " + param), nil
	})
	w.SetContentView(view)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	creator := appkit.NewApplicationDelegateCreator("MyApplicationDelegate")
	creator.SetApplicationDidFinishLaunching(func(o objc.Object, notification foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
		webkit.LoadURL(view, "gofs:/index.html")
	})
	creator.SetApplicationShouldTerminateAfterLastWindowClosed(func(o objc.Object, sender appkit.Application) bool {
		return true
	})
	app.SetDelegate(creator.Create())

	go func() {
		time.Sleep(time.Second * 1)
		runtime.GC()
	}()
	app.Run()
}
