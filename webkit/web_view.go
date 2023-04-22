// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/coregraphics"
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
	rv := objc.CallMethod[WebView](w_, objc.GetSelector("initWithFrame:configuration:"), frame, configuration)
	return rv
}

func (w_ WebView) InitWithFrame(frameRect foundation.Rect) WebView {
	rv := objc.CallMethod[WebView](w_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func (w_ WebView) Init() WebView {
	rv := objc.CallMethod[WebView](w_, objc.GetSelector("init"))
	return rv
}

func (wc _WebViewClass) Alloc() WebView {
	rv := objc.CallMethod[WebView](wc, objc.GetSelector("alloc"))
	return rv
}

func (wc _WebViewClass) New() WebView {
	rv := objc.CallMethod[WebView](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWebView() WebView {
	return WebViewClass.New()
}

func (wc _WebViewClass) HandlesURLScheme(urlScheme string) bool {
	rv := objc.CallMethod[bool](wc, objc.GetSelector("handlesURLScheme:"), urlScheme)
	return rv
}

func (w_ WebView) LoadRequest(request foundation.IURLRequest) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("loadRequest:"), request)
	return rv
}

func (w_ WebView) LoadData_MIMEType_CharacterEncodingName_BaseURL(data []byte, MIMEType string, characterEncodingName string, baseURL foundation.IURL) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("loadData:MIMEType:characterEncodingName:baseURL:"), data, MIMEType, characterEncodingName, baseURL)
	return rv
}

func (w_ WebView) LoadHTMLString_BaseURL(string_ string, baseURL foundation.IURL) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("loadHTMLString:baseURL:"), string_, baseURL)
	return rv
}

func (w_ WebView) LoadFileRequest_AllowingReadAccessToURL(request foundation.IURLRequest, readAccessURL foundation.IURL) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("loadFileRequest:allowingReadAccessToURL:"), request, readAccessURL)
	return rv
}

func (w_ WebView) LoadFileURL_AllowingReadAccessToURL(URL foundation.IURL, readAccessURL foundation.IURL) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("loadFileURL:allowingReadAccessToURL:"), URL, readAccessURL)
	return rv
}

func (w_ WebView) LoadSimulatedRequest_Response_ResponseData(request foundation.IURLRequest, response foundation.IURLResponse, data []byte) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("loadSimulatedRequest:response:responseData:"), request, response, data)
	return rv
}

func (w_ WebView) LoadSimulatedRequest_ResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("loadSimulatedRequest:responseHTMLString:"), request, string_)
	return rv
}

func (w_ WebView) Reload() Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("reload"))
	return rv
}

func (w_ WebView) Reload1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("reload:"), sender)
}

func (w_ WebView) ReloadFromOrigin() Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("reloadFromOrigin"))
	return rv
}

func (w_ WebView) ReloadFromOrigin1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("reloadFromOrigin:"), sender)
}

func (w_ WebView) StopLoading() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("stopLoading"))
}

func (w_ WebView) StopLoading1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("stopLoading:"), sender)
}

func (w_ WebView) StartDownloadUsingRequest_CompletionHandler(request foundation.IURLRequest, completionHandler func(param1 Download)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("startDownloadUsingRequest:completionHandler:"), request, completionHandler)
}

func (w_ WebView) ResumeDownloadFromResumeData_CompletionHandler(resumeData []byte, completionHandler func(param1 Download)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("resumeDownloadFromResumeData:completionHandler:"), resumeData, completionHandler)
}

func (w_ WebView) SetMagnification_CenteredAtPoint(magnification float64, point coregraphics.Point) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMagnification:centeredAtPoint:"), magnification, point)
}

func (w_ WebView) PauseAllMediaPlaybackWithCompletionHandler(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("pauseAllMediaPlaybackWithCompletionHandler:"), completionHandler)
}

// deprecated
func (w_ WebView) PauseAllMediaPlayback(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("pauseAllMediaPlayback:"), completionHandler)
}

func (w_ WebView) RequestMediaPlaybackStateWithCompletionHandler(completionHandler func(param1 MediaPlaybackState)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("requestMediaPlaybackStateWithCompletionHandler:"), completionHandler)
}

// deprecated
func (w_ WebView) RequestMediaPlaybackState(completionHandler func(param1 MediaPlaybackState)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("requestMediaPlaybackState:"), completionHandler)
}

func (w_ WebView) SetAllMediaPlaybackSuspended_CompletionHandler(suspended bool, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAllMediaPlaybackSuspended:completionHandler:"), suspended, completionHandler)
}

// deprecated
func (w_ WebView) SuspendAllMediaPlayback(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("suspendAllMediaPlayback:"), completionHandler)
}

// deprecated
func (w_ WebView) ResumeAllMediaPlayback(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("resumeAllMediaPlayback:"), completionHandler)
}

func (w_ WebView) CloseAllMediaPresentationsWithCompletionHandler(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("closeAllMediaPresentationsWithCompletionHandler:"), completionHandler)
}

func (w_ WebView) SetCameraCaptureState_CompletionHandler(state MediaCaptureState, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setCameraCaptureState:completionHandler:"), state, completionHandler)
}

func (w_ WebView) SetMicrophoneCaptureState_CompletionHandler(state MediaCaptureState, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMicrophoneCaptureState:completionHandler:"), state, completionHandler)
}

func (w_ WebView) FindString_WithConfiguration_CompletionHandler(string_ string, configuration IFindConfiguration, completionHandler func(result FindResult)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("findString:withConfiguration:completionHandler:"), string_, configuration, completionHandler)
}

func (w_ WebView) GoBack1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("goBack:"), sender)
}

func (w_ WebView) GoBack() Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("goBack"))
	return rv
}

func (w_ WebView) GoForward1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("goForward:"), sender)
}

func (w_ WebView) GoForward() Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("goForward"))
	return rv
}

func (w_ WebView) GoToBackForwardListItem(item IBackForwardListItem) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("goToBackForwardListItem:"), item)
	return rv
}

func (w_ WebView) EvaluateJavaScript_CompletionHandler(javaScriptString string, completionHandler func(param1 objc.Object, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("evaluateJavaScript:completionHandler:"), javaScriptString, completionHandler)
}

func (w_ WebView) EvaluateJavaScript_InFrame_InContentWorld_CompletionHandler(javaScriptString string, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(param1 objc.Object, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("evaluateJavaScript:inFrame:inContentWorld:completionHandler:"), javaScriptString, frame, contentWorld, completionHandler)
}

func (w_ WebView) CallAsyncJavaScript_Arguments_InFrame_InContentWorld_CompletionHandler(functionBody string, arguments map[string]objc.IObject, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(param1 objc.Object, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("callAsyncJavaScript:arguments:inFrame:inContentWorld:completionHandler:"), functionBody, arguments, frame, contentWorld, completionHandler)
}

func (w_ WebView) TakeSnapshotWithConfiguration_CompletionHandler(snapshotConfiguration ISnapshotConfiguration, completionHandler func(snapshotImage appkit.Image, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("takeSnapshotWithConfiguration:completionHandler:"), snapshotConfiguration, completionHandler)
}

func (w_ WebView) CreatePDFWithConfiguration_CompletionHandler(pdfConfiguration IPDFConfiguration, completionHandler func(pdfDocumentData []byte, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("createPDFWithConfiguration:completionHandler:"), pdfConfiguration, completionHandler)
}

func (w_ WebView) CreateWebArchiveDataWithCompletionHandler(completionHandler func(param1 []byte, param2 foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("createWebArchiveDataWithCompletionHandler:"), completionHandler)
}

func (w_ WebView) PrintOperationWithPrintInfo(printInfo appkit.IPrintInfo) appkit.PrintOperation {
	rv := objc.CallMethod[appkit.PrintOperation](w_, objc.GetSelector("printOperationWithPrintInfo:"), printInfo)
	return rv
}

func (w_ WebView) SetMinimumViewportInset_MaximumViewportInset(minimumViewportInset foundation.EdgeInsets, maximumViewportInset foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMinimumViewportInset:maximumViewportInset:"), minimumViewportInset, maximumViewportInset)
}

// deprecated
func (w_ WebView) CloseAllMediaPresentations() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("closeAllMediaPresentations"))
}

// deprecated
func (w_ WebView) LoadSimulatedRequest_WithResponse_ResponseData(request foundation.IURLRequest, response foundation.IURLResponse, data []byte) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("loadSimulatedRequest:withResponse:responseData:"), request, response, data)
	return rv
}

// deprecated
func (w_ WebView) LoadSimulatedRequest_WithResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.GetSelector("loadSimulatedRequest:withResponseHTMLString:"), request, string_)
	return rv
}

func (w_ WebView) Configuration() WebViewConfiguration {
	rv := objc.CallMethod[WebViewConfiguration](w_, objc.GetSelector("configuration"))
	return rv
}

func (w_ WebView) UIDelegate() UIDelegateWrapper {
	rv := objc.CallMethod[UIDelegateWrapper](w_, objc.GetSelector("UIDelegate"))
	return rv
}

func (w_ WebView) SetUIDelegate(value UIDelegate) {
	po := objc.CreateProtocol("WKUIDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(w_, internal.AssociationKey("setUIDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setUIDelegate:"), po)
}

func (w_ WebView) SetUIDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setUIDelegate:"), value)
}

func (w_ WebView) NavigationDelegate() NavigationDelegateWrapper {
	rv := objc.CallMethod[NavigationDelegateWrapper](w_, objc.GetSelector("navigationDelegate"))
	return rv
}

func (w_ WebView) SetNavigationDelegate(value NavigationDelegate) {
	po := objc.CreateProtocol("WKNavigationDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(w_, internal.AssociationKey("setNavigationDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setNavigationDelegate:"), po)
}

func (w_ WebView) SetNavigationDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setNavigationDelegate:"), value)
}

func (w_ WebView) IsLoading() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isLoading"))
	return rv
}

func (w_ WebView) EstimatedProgress() float64 {
	rv := objc.CallMethod[float64](w_, objc.GetSelector("estimatedProgress"))
	return rv
}

func (w_ WebView) Title() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("title"))
	return rv
}

func (w_ WebView) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("URL"))
	return rv
}

func (w_ WebView) MediaType() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("mediaType"))
	return rv
}

func (w_ WebView) SetMediaType(value string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMediaType:"), value)
}

func (w_ WebView) CustomUserAgent() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("customUserAgent"))
	return rv
}

func (w_ WebView) SetCustomUserAgent(value string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setCustomUserAgent:"), value)
}

func (w_ WebView) HasOnlySecureContent() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("hasOnlySecureContent"))
	return rv
}

func (w_ WebView) ThemeColor() appkit.Color {
	rv := objc.CallMethod[appkit.Color](w_, objc.GetSelector("themeColor"))
	return rv
}

func (w_ WebView) UnderPageBackgroundColor() appkit.Color {
	rv := objc.CallMethod[appkit.Color](w_, objc.GetSelector("underPageBackgroundColor"))
	return rv
}

func (w_ WebView) SetUnderPageBackgroundColor(value appkit.IColor) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setUnderPageBackgroundColor:"), value)
}

func (w_ WebView) PageZoom() float64 {
	rv := objc.CallMethod[float64](w_, objc.GetSelector("pageZoom"))
	return rv
}

func (w_ WebView) SetPageZoom(value float64) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setPageZoom:"), value)
}

func (w_ WebView) AllowsMagnification() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("allowsMagnification"))
	return rv
}

func (w_ WebView) SetAllowsMagnification(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAllowsMagnification:"), value)
}

func (w_ WebView) Magnification() float64 {
	rv := objc.CallMethod[float64](w_, objc.GetSelector("magnification"))
	return rv
}

func (w_ WebView) SetMagnification(value float64) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMagnification:"), value)
}

func (w_ WebView) CameraCaptureState() MediaCaptureState {
	rv := objc.CallMethod[MediaCaptureState](w_, objc.GetSelector("cameraCaptureState"))
	return rv
}

func (w_ WebView) MicrophoneCaptureState() MediaCaptureState {
	rv := objc.CallMethod[MediaCaptureState](w_, objc.GetSelector("microphoneCaptureState"))
	return rv
}

func (w_ WebView) AllowsBackForwardNavigationGestures() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("allowsBackForwardNavigationGestures"))
	return rv
}

func (w_ WebView) SetAllowsBackForwardNavigationGestures(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAllowsBackForwardNavigationGestures:"), value)
}

func (w_ WebView) BackForwardList() BackForwardList {
	rv := objc.CallMethod[BackForwardList](w_, objc.GetSelector("backForwardList"))
	return rv
}

func (w_ WebView) CanGoBack() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("canGoBack"))
	return rv
}

func (w_ WebView) CanGoForward() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("canGoForward"))
	return rv
}

func (w_ WebView) AllowsLinkPreview() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("allowsLinkPreview"))
	return rv
}

func (w_ WebView) SetAllowsLinkPreview(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAllowsLinkPreview:"), value)
}

func (w_ WebView) InteractionState() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.GetSelector("interactionState"))
	return rv
}

func (w_ WebView) SetInteractionState(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setInteractionState:"), value)
}

func (w_ WebView) FullscreenState() FullscreenState {
	rv := objc.CallMethod[FullscreenState](w_, objc.GetSelector("fullscreenState"))
	return rv
}

func (w_ WebView) MaximumViewportInset() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](w_, objc.GetSelector("maximumViewportInset"))
	return rv
}

func (w_ WebView) MinimumViewportInset() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](w_, objc.GetSelector("minimumViewportInset"))
	return rv
}

// deprecated
func (w_ WebView) CertificateChain() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](w_, objc.GetSelector("certificateChain"))
	return rv
}
