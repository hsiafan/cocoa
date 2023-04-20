// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var WebViewClass = _WebViewClass{objc.GetClass("WKWebView")}

type _WebViewClass struct {
	objc.Class
}

type IWebView interface {
	appkit.IView
	LoadRequest(request foundation.IURLRequest) Navigation
	LoadData_MIMEType_CharacterEncodingName_BaseURL(data []byte, MIMEType string, characterEncodingName string, baseURL foundation.IURL) Navigation
	LoadHTMLString_BaseURL(string_ string, baseURL foundation.IURL) Navigation
	LoadFileRequest_AllowingReadAccessToURL(request foundation.IURLRequest, readAccessURL foundation.IURL) Navigation
	LoadFileURL_AllowingReadAccessToURL(URL foundation.IURL, readAccessURL foundation.IURL) Navigation
	LoadSimulatedRequest_Response_ResponseData(request foundation.IURLRequest, response foundation.IURLResponse, data []byte) Navigation
	LoadSimulatedRequest_ResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation
	Reload() Navigation
	Reload1(sender objc.IObject)
	ReloadFromOrigin() Navigation
	ReloadFromOrigin1(sender objc.IObject)
	StopLoading()
	StopLoading1(sender objc.IObject)
	StartDownloadUsingRequest_CompletionHandler(request foundation.IURLRequest, completionHandler func(param1 Download))
	ResumeDownloadFromResumeData_CompletionHandler(resumeData []byte, completionHandler func(param1 Download))
	SetMagnification_CenteredAtPoint(magnification float64, point coregraphics.Point)
	PauseAllMediaPlaybackWithCompletionHandler(completionHandler func())
	// deprecated
	PauseAllMediaPlayback(completionHandler func())
	RequestMediaPlaybackStateWithCompletionHandler(completionHandler func(param1 MediaPlaybackState))
	// deprecated
	RequestMediaPlaybackState(completionHandler func(param1 MediaPlaybackState))
	SetAllMediaPlaybackSuspended_CompletionHandler(suspended bool, completionHandler func())
	// deprecated
	SuspendAllMediaPlayback(completionHandler func())
	// deprecated
	ResumeAllMediaPlayback(completionHandler func())
	CloseAllMediaPresentationsWithCompletionHandler(completionHandler func())
	SetCameraCaptureState_CompletionHandler(state MediaCaptureState, completionHandler func())
	SetMicrophoneCaptureState_CompletionHandler(state MediaCaptureState, completionHandler func())
	FindString_WithConfiguration_CompletionHandler(string_ string, configuration IFindConfiguration, completionHandler func(result FindResult))
	GoBack1(sender objc.IObject)
	GoBack() Navigation
	GoForward1(sender objc.IObject)
	GoForward() Navigation
	GoToBackForwardListItem(item IBackForwardListItem) Navigation
	EvaluateJavaScript_CompletionHandler(javaScriptString string, completionHandler func(param1 objc.Object, error foundation.Error))
	EvaluateJavaScript_InFrame_InContentWorld_CompletionHandler(javaScriptString string, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(param1 objc.Object, error foundation.Error))
	CallAsyncJavaScript_Arguments_InFrame_InContentWorld_CompletionHandler(functionBody string, arguments map[string]objc.IObject, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(param1 objc.Object, error foundation.Error))
	TakeSnapshotWithConfiguration_CompletionHandler(snapshotConfiguration ISnapshotConfiguration, completionHandler func(snapshotImage appkit.Image, error foundation.Error))
	CreatePDFWithConfiguration_CompletionHandler(pdfConfiguration IPDFConfiguration, completionHandler func(pdfDocumentData []byte, error foundation.Error))
	CreateWebArchiveDataWithCompletionHandler(completionHandler func(param1 []byte, param2 foundation.Error))
	PrintOperationWithPrintInfo(printInfo appkit.IPrintInfo) appkit.PrintOperation
	SetMinimumViewportInset_MaximumViewportInset(minimumViewportInset foundation.EdgeInsets, maximumViewportInset foundation.EdgeInsets)
	// deprecated
	CloseAllMediaPresentations()
	// deprecated
	LoadSimulatedRequest_WithResponse_ResponseData(request foundation.IURLRequest, response foundation.IURLResponse, data []byte) Navigation
	// deprecated
	LoadSimulatedRequest_WithResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation
	Configuration() WebViewConfiguration
	UIDelegate() UIDelegateWrapper
	SetUIDelegate(value UIDelegate)
	SetUIDelegate0(value objc.IObject)
	NavigationDelegate() NavigationDelegateWrapper
	SetNavigationDelegate(value NavigationDelegate)
	SetNavigationDelegate0(value objc.IObject)
	IsLoading() bool
	EstimatedProgress() float64
	Title() string
	URL() foundation.URL
	MediaType() string
	SetMediaType(value string)
	CustomUserAgent() string
	SetCustomUserAgent(value string)
	HasOnlySecureContent() bool
	ThemeColor() appkit.Color
	UnderPageBackgroundColor() appkit.Color
	SetUnderPageBackgroundColor(value appkit.IColor)
	PageZoom() float64
	SetPageZoom(value float64)
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	Magnification() float64
	SetMagnification(value float64)
	CameraCaptureState() MediaCaptureState
	MicrophoneCaptureState() MediaCaptureState
	AllowsBackForwardNavigationGestures() bool
	SetAllowsBackForwardNavigationGestures(value bool)
	BackForwardList() BackForwardList
	CanGoBack() bool
	CanGoForward() bool
	AllowsLinkPreview() bool
	SetAllowsLinkPreview(value bool)
	InteractionState() objc.Object
	SetInteractionState(value objc.IObject)
	FullscreenState() FullscreenState
	MaximumViewportInset() foundation.EdgeInsets
	MinimumViewportInset() foundation.EdgeInsets
	// deprecated
	CertificateChain() []objc.Object
}

type WebView struct {
	appkit.View
}

func MakeWebView(ptr unsafe.Pointer) WebView {
	return WebView{
		View: appkit.MakeView(ptr),
	}
}

func (w_ WebView) InitWithFrame_Configuration(frame coregraphics.Rect, configuration IWebViewConfiguration) WebView {
	rv := ffi.CallMethod[WebView](w_, "initWithFrame:configuration:", frame, configuration)
	return rv
}

func (w_ WebView) InitWithFrame(frameRect foundation.Rect) WebView {
	rv := ffi.CallMethod[WebView](w_, "initWithFrame:", frameRect)
	return rv
}

func (w_ WebView) Init() WebView {
	rv := ffi.CallMethod[WebView](w_, "init")
	return rv
}

func (wc _WebViewClass) Alloc() WebView {
	rv := ffi.CallMethod[WebView](wc, "alloc")
	return rv
}

func (wc _WebViewClass) New() WebView {
	rv := ffi.CallMethod[WebView](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWebView() WebView {
	return WebViewClass.New()
}

func (wc _WebViewClass) HandlesURLScheme(urlScheme string) bool {
	rv := ffi.CallMethod[bool](wc, "handlesURLScheme:", urlScheme)
	return rv
}

func (w_ WebView) LoadRequest(request foundation.IURLRequest) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "loadRequest:", request)
	return rv
}

func (w_ WebView) LoadData_MIMEType_CharacterEncodingName_BaseURL(data []byte, MIMEType string, characterEncodingName string, baseURL foundation.IURL) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "loadData:MIMEType:characterEncodingName:baseURL:", data, MIMEType, characterEncodingName, baseURL)
	return rv
}

func (w_ WebView) LoadHTMLString_BaseURL(string_ string, baseURL foundation.IURL) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "loadHTMLString:baseURL:", string_, baseURL)
	return rv
}

func (w_ WebView) LoadFileRequest_AllowingReadAccessToURL(request foundation.IURLRequest, readAccessURL foundation.IURL) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "loadFileRequest:allowingReadAccessToURL:", request, readAccessURL)
	return rv
}

func (w_ WebView) LoadFileURL_AllowingReadAccessToURL(URL foundation.IURL, readAccessURL foundation.IURL) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "loadFileURL:allowingReadAccessToURL:", URL, readAccessURL)
	return rv
}

func (w_ WebView) LoadSimulatedRequest_Response_ResponseData(request foundation.IURLRequest, response foundation.IURLResponse, data []byte) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "loadSimulatedRequest:response:responseData:", request, response, data)
	return rv
}

func (w_ WebView) LoadSimulatedRequest_ResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "loadSimulatedRequest:responseHTMLString:", request, string_)
	return rv
}

func (w_ WebView) Reload() Navigation {
	rv := ffi.CallMethod[Navigation](w_, "reload")
	return rv
}

func (w_ WebView) Reload1(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "reload:", sender)
}

func (w_ WebView) ReloadFromOrigin() Navigation {
	rv := ffi.CallMethod[Navigation](w_, "reloadFromOrigin")
	return rv
}

func (w_ WebView) ReloadFromOrigin1(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "reloadFromOrigin:", sender)
}

func (w_ WebView) StopLoading() {
	ffi.CallMethod[ffi.Void](w_, "stopLoading")
}

func (w_ WebView) StopLoading1(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "stopLoading:", sender)
}

func (w_ WebView) StartDownloadUsingRequest_CompletionHandler(request foundation.IURLRequest, completionHandler func(param1 Download)) {
	ffi.CallMethod[ffi.Void](w_, "startDownloadUsingRequest:completionHandler:", request, completionHandler)
}

func (w_ WebView) ResumeDownloadFromResumeData_CompletionHandler(resumeData []byte, completionHandler func(param1 Download)) {
	ffi.CallMethod[ffi.Void](w_, "resumeDownloadFromResumeData:completionHandler:", resumeData, completionHandler)
}

func (w_ WebView) SetMagnification_CenteredAtPoint(magnification float64, point coregraphics.Point) {
	ffi.CallMethod[ffi.Void](w_, "setMagnification:centeredAtPoint:", magnification, point)
}

func (w_ WebView) PauseAllMediaPlaybackWithCompletionHandler(completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "pauseAllMediaPlaybackWithCompletionHandler:", completionHandler)
}

// deprecated
func (w_ WebView) PauseAllMediaPlayback(completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "pauseAllMediaPlayback:", completionHandler)
}

func (w_ WebView) RequestMediaPlaybackStateWithCompletionHandler(completionHandler func(param1 MediaPlaybackState)) {
	ffi.CallMethod[ffi.Void](w_, "requestMediaPlaybackStateWithCompletionHandler:", completionHandler)
}

// deprecated
func (w_ WebView) RequestMediaPlaybackState(completionHandler func(param1 MediaPlaybackState)) {
	ffi.CallMethod[ffi.Void](w_, "requestMediaPlaybackState:", completionHandler)
}

func (w_ WebView) SetAllMediaPlaybackSuspended_CompletionHandler(suspended bool, completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "setAllMediaPlaybackSuspended:completionHandler:", suspended, completionHandler)
}

// deprecated
func (w_ WebView) SuspendAllMediaPlayback(completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "suspendAllMediaPlayback:", completionHandler)
}

// deprecated
func (w_ WebView) ResumeAllMediaPlayback(completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "resumeAllMediaPlayback:", completionHandler)
}

func (w_ WebView) CloseAllMediaPresentationsWithCompletionHandler(completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "closeAllMediaPresentationsWithCompletionHandler:", completionHandler)
}

func (w_ WebView) SetCameraCaptureState_CompletionHandler(state MediaCaptureState, completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "setCameraCaptureState:completionHandler:", state, completionHandler)
}

func (w_ WebView) SetMicrophoneCaptureState_CompletionHandler(state MediaCaptureState, completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "setMicrophoneCaptureState:completionHandler:", state, completionHandler)
}

func (w_ WebView) FindString_WithConfiguration_CompletionHandler(string_ string, configuration IFindConfiguration, completionHandler func(result FindResult)) {
	ffi.CallMethod[ffi.Void](w_, "findString:withConfiguration:completionHandler:", string_, configuration, completionHandler)
}

func (w_ WebView) GoBack1(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "goBack:", sender)
}

func (w_ WebView) GoBack() Navigation {
	rv := ffi.CallMethod[Navigation](w_, "goBack")
	return rv
}

func (w_ WebView) GoForward1(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "goForward:", sender)
}

func (w_ WebView) GoForward() Navigation {
	rv := ffi.CallMethod[Navigation](w_, "goForward")
	return rv
}

func (w_ WebView) GoToBackForwardListItem(item IBackForwardListItem) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "goToBackForwardListItem:", item)
	return rv
}

func (w_ WebView) EvaluateJavaScript_CompletionHandler(javaScriptString string, completionHandler func(param1 objc.Object, error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "evaluateJavaScript:completionHandler:", javaScriptString, completionHandler)
}

func (w_ WebView) EvaluateJavaScript_InFrame_InContentWorld_CompletionHandler(javaScriptString string, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(param1 objc.Object, error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "evaluateJavaScript:inFrame:inContentWorld:completionHandler:", javaScriptString, frame, contentWorld, completionHandler)
}

func (w_ WebView) CallAsyncJavaScript_Arguments_InFrame_InContentWorld_CompletionHandler(functionBody string, arguments map[string]objc.IObject, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(param1 objc.Object, error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "callAsyncJavaScript:arguments:inFrame:inContentWorld:completionHandler:", functionBody, arguments, frame, contentWorld, completionHandler)
}

func (w_ WebView) TakeSnapshotWithConfiguration_CompletionHandler(snapshotConfiguration ISnapshotConfiguration, completionHandler func(snapshotImage appkit.Image, error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "takeSnapshotWithConfiguration:completionHandler:", snapshotConfiguration, completionHandler)
}

func (w_ WebView) CreatePDFWithConfiguration_CompletionHandler(pdfConfiguration IPDFConfiguration, completionHandler func(pdfDocumentData []byte, error foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "createPDFWithConfiguration:completionHandler:", pdfConfiguration, completionHandler)
}

func (w_ WebView) CreateWebArchiveDataWithCompletionHandler(completionHandler func(param1 []byte, param2 foundation.Error)) {
	ffi.CallMethod[ffi.Void](w_, "createWebArchiveDataWithCompletionHandler:", completionHandler)
}

func (w_ WebView) PrintOperationWithPrintInfo(printInfo appkit.IPrintInfo) appkit.PrintOperation {
	rv := ffi.CallMethod[appkit.PrintOperation](w_, "printOperationWithPrintInfo:", printInfo)
	return rv
}

func (w_ WebView) SetMinimumViewportInset_MaximumViewportInset(minimumViewportInset foundation.EdgeInsets, maximumViewportInset foundation.EdgeInsets) {
	ffi.CallMethod[ffi.Void](w_, "setMinimumViewportInset:maximumViewportInset:", minimumViewportInset, maximumViewportInset)
}

// deprecated
func (w_ WebView) CloseAllMediaPresentations() {
	ffi.CallMethod[ffi.Void](w_, "closeAllMediaPresentations")
}

// deprecated
func (w_ WebView) LoadSimulatedRequest_WithResponse_ResponseData(request foundation.IURLRequest, response foundation.IURLResponse, data []byte) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "loadSimulatedRequest:withResponse:responseData:", request, response, data)
	return rv
}

// deprecated
func (w_ WebView) LoadSimulatedRequest_WithResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation {
	rv := ffi.CallMethod[Navigation](w_, "loadSimulatedRequest:withResponseHTMLString:", request, string_)
	return rv
}

func (w_ WebView) Configuration() WebViewConfiguration {
	rv := ffi.CallMethod[WebViewConfiguration](w_, "configuration")
	return rv
}

func (w_ WebView) UIDelegate() UIDelegateWrapper {
	rv := ffi.CallMethod[UIDelegateWrapper](w_, "UIDelegate")
	return rv
}

func (w_ WebView) SetUIDelegate(value UIDelegate) {
	po := ffi.CreateProtocol("WKUIDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(w_, internal.AssociationKey("setUIDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](w_, "setUIDelegate:", po)
}

func (w_ WebView) SetUIDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "setUIDelegate:", value)
}

func (w_ WebView) NavigationDelegate() NavigationDelegateWrapper {
	rv := ffi.CallMethod[NavigationDelegateWrapper](w_, "navigationDelegate")
	return rv
}

func (w_ WebView) SetNavigationDelegate(value NavigationDelegate) {
	po := ffi.CreateProtocol("WKNavigationDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(w_, internal.AssociationKey("setNavigationDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](w_, "setNavigationDelegate:", po)
}

func (w_ WebView) SetNavigationDelegate0(value objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "setNavigationDelegate:", value)
}

func (w_ WebView) IsLoading() bool {
	rv := ffi.CallMethod[bool](w_, "isLoading")
	return rv
}

func (w_ WebView) EstimatedProgress() float64 {
	rv := ffi.CallMethod[float64](w_, "estimatedProgress")
	return rv
}

func (w_ WebView) Title() string {
	rv := ffi.CallMethod[string](w_, "title")
	return rv
}

func (w_ WebView) URL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](w_, "URL")
	return rv
}

func (w_ WebView) MediaType() string {
	rv := ffi.CallMethod[string](w_, "mediaType")
	return rv
}

func (w_ WebView) SetMediaType(value string) {
	ffi.CallMethod[ffi.Void](w_, "setMediaType:", value)
}

func (w_ WebView) CustomUserAgent() string {
	rv := ffi.CallMethod[string](w_, "customUserAgent")
	return rv
}

func (w_ WebView) SetCustomUserAgent(value string) {
	ffi.CallMethod[ffi.Void](w_, "setCustomUserAgent:", value)
}

func (w_ WebView) HasOnlySecureContent() bool {
	rv := ffi.CallMethod[bool](w_, "hasOnlySecureContent")
	return rv
}

func (w_ WebView) ThemeColor() appkit.Color {
	rv := ffi.CallMethod[appkit.Color](w_, "themeColor")
	return rv
}

func (w_ WebView) UnderPageBackgroundColor() appkit.Color {
	rv := ffi.CallMethod[appkit.Color](w_, "underPageBackgroundColor")
	return rv
}

func (w_ WebView) SetUnderPageBackgroundColor(value appkit.IColor) {
	ffi.CallMethod[ffi.Void](w_, "setUnderPageBackgroundColor:", value)
}

func (w_ WebView) PageZoom() float64 {
	rv := ffi.CallMethod[float64](w_, "pageZoom")
	return rv
}

func (w_ WebView) SetPageZoom(value float64) {
	ffi.CallMethod[ffi.Void](w_, "setPageZoom:", value)
}

func (w_ WebView) AllowsMagnification() bool {
	rv := ffi.CallMethod[bool](w_, "allowsMagnification")
	return rv
}

func (w_ WebView) SetAllowsMagnification(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAllowsMagnification:", value)
}

func (w_ WebView) Magnification() float64 {
	rv := ffi.CallMethod[float64](w_, "magnification")
	return rv
}

func (w_ WebView) SetMagnification(value float64) {
	ffi.CallMethod[ffi.Void](w_, "setMagnification:", value)
}

func (w_ WebView) CameraCaptureState() MediaCaptureState {
	rv := ffi.CallMethod[MediaCaptureState](w_, "cameraCaptureState")
	return rv
}

func (w_ WebView) MicrophoneCaptureState() MediaCaptureState {
	rv := ffi.CallMethod[MediaCaptureState](w_, "microphoneCaptureState")
	return rv
}

func (w_ WebView) AllowsBackForwardNavigationGestures() bool {
	rv := ffi.CallMethod[bool](w_, "allowsBackForwardNavigationGestures")
	return rv
}

func (w_ WebView) SetAllowsBackForwardNavigationGestures(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAllowsBackForwardNavigationGestures:", value)
}

func (w_ WebView) BackForwardList() BackForwardList {
	rv := ffi.CallMethod[BackForwardList](w_, "backForwardList")
	return rv
}

func (w_ WebView) CanGoBack() bool {
	rv := ffi.CallMethod[bool](w_, "canGoBack")
	return rv
}

func (w_ WebView) CanGoForward() bool {
	rv := ffi.CallMethod[bool](w_, "canGoForward")
	return rv
}

func (w_ WebView) AllowsLinkPreview() bool {
	rv := ffi.CallMethod[bool](w_, "allowsLinkPreview")
	return rv
}

func (w_ WebView) SetAllowsLinkPreview(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAllowsLinkPreview:", value)
}

func (w_ WebView) InteractionState() objc.Object {
	rv := ffi.CallMethod[objc.Object](w_, "interactionState")
	return rv
}

func (w_ WebView) SetInteractionState(value objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "setInteractionState:", value)
}

func (w_ WebView) FullscreenState() FullscreenState {
	rv := ffi.CallMethod[FullscreenState](w_, "fullscreenState")
	return rv
}

func (w_ WebView) MaximumViewportInset() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](w_, "maximumViewportInset")
	return rv
}

func (w_ WebView) MinimumViewportInset() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](w_, "minimumViewportInset")
	return rv
}

// deprecated
func (w_ WebView) CertificateChain() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](w_, "certificateChain")
	return rv
}
