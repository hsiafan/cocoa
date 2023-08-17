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

	app.SetDelegate(appkit.WrapApplicationDelegate(&myApplicationDelegate{app: app, webView: view}))

	go func() {
		time.Sleep(time.Second * 1)
		runtime.GC()
	}()
	app.Run()
}

type myApplicationDelegate struct {
	appkit.ApplicationDelegateBase
	app     appkit.Application
	webView webkit.WebView
}

func (p *myApplicationDelegate) ImplementsApplicationDidFinishLaunching() bool {
	return true
}

func (p *myApplicationDelegate) ApplicationDidFinishLaunching(notification foundation.Notification) {
	p.app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	p.app.ActivateIgnoringOtherApps(true)
	webkit.LoadURL(p.webView, "gofs:/index.html")
}

func (p *myApplicationDelegate) ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool {
	return true
}

func (p *myApplicationDelegate) ApplicationShouldTerminateAfterLastWindowClosed(sender appkit.Application) bool {
	return true
}
