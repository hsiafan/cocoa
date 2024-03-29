// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
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
	UIDelegate() objc.Object
	SetUIDelegate(value objc.IObject)
	NavigationDelegate() objc.Object
	SetNavigationDelegate(value objc.IObject)
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
	rv := objc.CallMethod[WebView](w_, objc.SEL("initWithFrame:configuration:"), frame, objc.ExtractPtr(configuration))
	return rv
}

func (w_ WebView) InitWithFrame(frameRect foundation.Rect) WebView {
	rv := objc.CallMethod[WebView](w_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (w_ WebView) Init() WebView {
	rv := objc.CallMethod[WebView](w_, objc.SEL("init"))
	return rv
}

func (wc _WebViewClass) Alloc() WebView {
	rv := objc.CallMethod[WebView](wc, objc.SEL("alloc"))
	return rv
}

func (wc _WebViewClass) New() WebView {
	rv := objc.CallMethod[WebView](wc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewWebView() WebView {
	return WebViewClass.New()
}

func (wc _WebViewClass) HandlesURLScheme(urlScheme string) bool {
	rv := objc.CallMethod[bool](wc, objc.SEL("handlesURLScheme:"), urlScheme)
	return rv
}

func (w_ WebView) LoadRequest(request foundation.IURLRequest) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("loadRequest:"), objc.ExtractPtr(request))
	return rv
}

func (w_ WebView) LoadData_MIMEType_CharacterEncodingName_BaseURL(data []byte, MIMEType string, characterEncodingName string, baseURL foundation.IURL) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("loadData:MIMEType:characterEncodingName:baseURL:"), data, MIMEType, characterEncodingName, objc.ExtractPtr(baseURL))
	return rv
}

func (w_ WebView) LoadHTMLString_BaseURL(string_ string, baseURL foundation.IURL) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("loadHTMLString:baseURL:"), string_, objc.ExtractPtr(baseURL))
	return rv
}

func (w_ WebView) LoadFileRequest_AllowingReadAccessToURL(request foundation.IURLRequest, readAccessURL foundation.IURL) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("loadFileRequest:allowingReadAccessToURL:"), objc.ExtractPtr(request), objc.ExtractPtr(readAccessURL))
	return rv
}

func (w_ WebView) LoadFileURL_AllowingReadAccessToURL(URL foundation.IURL, readAccessURL foundation.IURL) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("loadFileURL:allowingReadAccessToURL:"), objc.ExtractPtr(URL), objc.ExtractPtr(readAccessURL))
	return rv
}

func (w_ WebView) LoadSimulatedRequest_Response_ResponseData(request foundation.IURLRequest, response foundation.IURLResponse, data []byte) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("loadSimulatedRequest:response:responseData:"), objc.ExtractPtr(request), objc.ExtractPtr(response), data)
	return rv
}

func (w_ WebView) LoadSimulatedRequest_ResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("loadSimulatedRequest:responseHTMLString:"), objc.ExtractPtr(request), string_)
	return rv
}

func (w_ WebView) Reload() Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("reload"))
	return rv
}

func (w_ WebView) Reload1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("reload:"), objc.ExtractPtr(sender))
}

func (w_ WebView) ReloadFromOrigin() Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("reloadFromOrigin"))
	return rv
}

func (w_ WebView) ReloadFromOrigin1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("reloadFromOrigin:"), objc.ExtractPtr(sender))
}

func (w_ WebView) StopLoading() {
	objc.CallMethod[objc.Void](w_, objc.SEL("stopLoading"))
}

func (w_ WebView) StopLoading1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("stopLoading:"), objc.ExtractPtr(sender))
}

func (w_ WebView) StartDownloadUsingRequest_CompletionHandler(request foundation.IURLRequest, completionHandler func(param1 Download)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("startDownloadUsingRequest:completionHandler:"), objc.ExtractPtr(request), completionHandler)
}

func (w_ WebView) ResumeDownloadFromResumeData_CompletionHandler(resumeData []byte, completionHandler func(param1 Download)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("resumeDownloadFromResumeData:completionHandler:"), resumeData, completionHandler)
}

func (w_ WebView) SetMagnification_CenteredAtPoint(magnification float64, point coregraphics.Point) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMagnification:centeredAtPoint:"), magnification, point)
}

func (w_ WebView) PauseAllMediaPlaybackWithCompletionHandler(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("pauseAllMediaPlaybackWithCompletionHandler:"), completionHandler)
}

// deprecated
func (w_ WebView) PauseAllMediaPlayback(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("pauseAllMediaPlayback:"), completionHandler)
}

func (w_ WebView) RequestMediaPlaybackStateWithCompletionHandler(completionHandler func(param1 MediaPlaybackState)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("requestMediaPlaybackStateWithCompletionHandler:"), completionHandler)
}

// deprecated
func (w_ WebView) RequestMediaPlaybackState(completionHandler func(param1 MediaPlaybackState)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("requestMediaPlaybackState:"), completionHandler)
}

func (w_ WebView) SetAllMediaPlaybackSuspended_CompletionHandler(suspended bool, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAllMediaPlaybackSuspended:completionHandler:"), suspended, completionHandler)
}

// deprecated
func (w_ WebView) SuspendAllMediaPlayback(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("suspendAllMediaPlayback:"), completionHandler)
}

// deprecated
func (w_ WebView) ResumeAllMediaPlayback(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("resumeAllMediaPlayback:"), completionHandler)
}

func (w_ WebView) CloseAllMediaPresentationsWithCompletionHandler(completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("closeAllMediaPresentationsWithCompletionHandler:"), completionHandler)
}

func (w_ WebView) SetCameraCaptureState_CompletionHandler(state MediaCaptureState, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setCameraCaptureState:completionHandler:"), state, completionHandler)
}

func (w_ WebView) SetMicrophoneCaptureState_CompletionHandler(state MediaCaptureState, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMicrophoneCaptureState:completionHandler:"), state, completionHandler)
}

func (w_ WebView) FindString_WithConfiguration_CompletionHandler(string_ string, configuration IFindConfiguration, completionHandler func(result FindResult)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("findString:withConfiguration:completionHandler:"), string_, objc.ExtractPtr(configuration), completionHandler)
}

func (w_ WebView) GoBack1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("goBack:"), objc.ExtractPtr(sender))
}

func (w_ WebView) GoBack() Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("goBack"))
	return rv
}

func (w_ WebView) GoForward1(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("goForward:"), objc.ExtractPtr(sender))
}

func (w_ WebView) GoForward() Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("goForward"))
	return rv
}

func (w_ WebView) GoToBackForwardListItem(item IBackForwardListItem) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("goToBackForwardListItem:"), objc.ExtractPtr(item))
	return rv
}

func (w_ WebView) EvaluateJavaScript_CompletionHandler(javaScriptString string, completionHandler func(param1 objc.Object, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("evaluateJavaScript:completionHandler:"), javaScriptString, completionHandler)
}

func (w_ WebView) EvaluateJavaScript_InFrame_InContentWorld_CompletionHandler(javaScriptString string, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(param1 objc.Object, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("evaluateJavaScript:inFrame:inContentWorld:completionHandler:"), javaScriptString, objc.ExtractPtr(frame), objc.ExtractPtr(contentWorld), completionHandler)
}

func (w_ WebView) CallAsyncJavaScript_Arguments_InFrame_InContentWorld_CompletionHandler(functionBody string, arguments map[string]objc.IObject, frame IFrameInfo, contentWorld IContentWorld, completionHandler func(param1 objc.Object, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("callAsyncJavaScript:arguments:inFrame:inContentWorld:completionHandler:"), functionBody, arguments, objc.ExtractPtr(frame), objc.ExtractPtr(contentWorld), completionHandler)
}

func (w_ WebView) TakeSnapshotWithConfiguration_CompletionHandler(snapshotConfiguration ISnapshotConfiguration, completionHandler func(snapshotImage appkit.Image, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("takeSnapshotWithConfiguration:completionHandler:"), objc.ExtractPtr(snapshotConfiguration), completionHandler)
}

func (w_ WebView) CreatePDFWithConfiguration_CompletionHandler(pdfConfiguration IPDFConfiguration, completionHandler func(pdfDocumentData []byte, error foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("createPDFWithConfiguration:completionHandler:"), objc.ExtractPtr(pdfConfiguration), completionHandler)
}

func (w_ WebView) CreateWebArchiveDataWithCompletionHandler(completionHandler func(param1 []byte, param2 foundation.Error)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("createWebArchiveDataWithCompletionHandler:"), completionHandler)
}

func (w_ WebView) PrintOperationWithPrintInfo(printInfo appkit.IPrintInfo) appkit.PrintOperation {
	rv := objc.CallMethod[appkit.PrintOperation](w_, objc.SEL("printOperationWithPrintInfo:"), objc.ExtractPtr(printInfo))
	return rv
}

func (w_ WebView) SetMinimumViewportInset_MaximumViewportInset(minimumViewportInset foundation.EdgeInsets, maximumViewportInset foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMinimumViewportInset:maximumViewportInset:"), minimumViewportInset, maximumViewportInset)
}

// deprecated
func (w_ WebView) CloseAllMediaPresentations() {
	objc.CallMethod[objc.Void](w_, objc.SEL("closeAllMediaPresentations"))
}

// deprecated
func (w_ WebView) LoadSimulatedRequest_WithResponse_ResponseData(request foundation.IURLRequest, response foundation.IURLResponse, data []byte) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("loadSimulatedRequest:withResponse:responseData:"), objc.ExtractPtr(request), objc.ExtractPtr(response), data)
	return rv
}

// deprecated
func (w_ WebView) LoadSimulatedRequest_WithResponseHTMLString(request foundation.IURLRequest, string_ string) Navigation {
	rv := objc.CallMethod[Navigation](w_, objc.SEL("loadSimulatedRequest:withResponseHTMLString:"), objc.ExtractPtr(request), string_)
	return rv
}

func (w_ WebView) Configuration() WebViewConfiguration {
	rv := objc.CallMethod[WebViewConfiguration](w_, objc.SEL("configuration"))
	return rv
}

// weak property
func (w_ WebView) UIDelegate() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("UIDelegate"))
	return rv
}

// weak property
func (w_ WebView) SetUIDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setUIDelegate:"), objc.ExtractPtr(value))
}

// weak property
func (w_ WebView) NavigationDelegate() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("navigationDelegate"))
	return rv
}

// weak property
func (w_ WebView) SetNavigationDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setNavigationDelegate:"), objc.ExtractPtr(value))
}

func (w_ WebView) IsLoading() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isLoading"))
	return rv
}

func (w_ WebView) EstimatedProgress() float64 {
	rv := objc.CallMethod[float64](w_, objc.SEL("estimatedProgress"))
	return rv
}

func (w_ WebView) Title() string {
	rv := objc.CallMethod[string](w_, objc.SEL("title"))
	return rv
}

func (w_ WebView) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.SEL("URL"))
	return rv
}

func (w_ WebView) MediaType() string {
	rv := objc.CallMethod[string](w_, objc.SEL("mediaType"))
	return rv
}

func (w_ WebView) SetMediaType(value string) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMediaType:"), value)
}

func (w_ WebView) CustomUserAgent() string {
	rv := objc.CallMethod[string](w_, objc.SEL("customUserAgent"))
	return rv
}

func (w_ WebView) SetCustomUserAgent(value string) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setCustomUserAgent:"), value)
}

func (w_ WebView) HasOnlySecureContent() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("hasOnlySecureContent"))
	return rv
}

func (w_ WebView) ThemeColor() appkit.Color {
	rv := objc.CallMethod[appkit.Color](w_, objc.SEL("themeColor"))
	return rv
}

func (w_ WebView) UnderPageBackgroundColor() appkit.Color {
	rv := objc.CallMethod[appkit.Color](w_, objc.SEL("underPageBackgroundColor"))
	return rv
}

func (w_ WebView) SetUnderPageBackgroundColor(value appkit.IColor) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setUnderPageBackgroundColor:"), objc.ExtractPtr(value))
}

func (w_ WebView) PageZoom() float64 {
	rv := objc.CallMethod[float64](w_, objc.SEL("pageZoom"))
	return rv
}

func (w_ WebView) SetPageZoom(value float64) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setPageZoom:"), value)
}

func (w_ WebView) AllowsMagnification() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("allowsMagnification"))
	return rv
}

func (w_ WebView) SetAllowsMagnification(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAllowsMagnification:"), value)
}

func (w_ WebView) Magnification() float64 {
	rv := objc.CallMethod[float64](w_, objc.SEL("magnification"))
	return rv
}

func (w_ WebView) SetMagnification(value float64) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMagnification:"), value)
}

func (w_ WebView) CameraCaptureState() MediaCaptureState {
	rv := objc.CallMethod[MediaCaptureState](w_, objc.SEL("cameraCaptureState"))
	return rv
}

func (w_ WebView) MicrophoneCaptureState() MediaCaptureState {
	rv := objc.CallMethod[MediaCaptureState](w_, objc.SEL("microphoneCaptureState"))
	return rv
}

func (w_ WebView) AllowsBackForwardNavigationGestures() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("allowsBackForwardNavigationGestures"))
	return rv
}

func (w_ WebView) SetAllowsBackForwardNavigationGestures(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAllowsBackForwardNavigationGestures:"), value)
}

func (w_ WebView) BackForwardList() BackForwardList {
	rv := objc.CallMethod[BackForwardList](w_, objc.SEL("backForwardList"))
	return rv
}

func (w_ WebView) CanGoBack() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("canGoBack"))
	return rv
}

func (w_ WebView) CanGoForward() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("canGoForward"))
	return rv
}

func (w_ WebView) AllowsLinkPreview() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("allowsLinkPreview"))
	return rv
}

func (w_ WebView) SetAllowsLinkPreview(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAllowsLinkPreview:"), value)
}

func (w_ WebView) InteractionState() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("interactionState"))
	return rv
}

func (w_ WebView) SetInteractionState(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setInteractionState:"), objc.ExtractPtr(value))
}

func (w_ WebView) FullscreenState() FullscreenState {
	rv := objc.CallMethod[FullscreenState](w_, objc.SEL("fullscreenState"))
	return rv
}

func (w_ WebView) MaximumViewportInset() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](w_, objc.SEL("maximumViewportInset"))
	return rv
}

func (w_ WebView) MinimumViewportInset() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](w_, objc.SEL("minimumViewportInset"))
	return rv
}

// deprecated
func (w_ WebView) CertificateChain() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](w_, objc.SEL("certificateChain"))
	return rv
}
