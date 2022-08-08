// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var WindowClass = _WindowClass{objc.GetClass("NSWindow")}

type _WindowClass struct {
	objc.Class
}

type IWindow interface {
	IResponder
	ToggleFullScreen(sender objc.IObject)
	SetDynamicDepthLimit(flag bool)
	InvalidateShadow()
	AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool
	SetAutorecalculatesContentBorderThickness_ForEdge(flag bool, edge foundation.RectEdge)
	ContentBorderThicknessForEdge(edge foundation.RectEdge) float64
	SetContentBorderThickness_ForEdge(thickness float64, edge foundation.RectEdge)
	ContentRectForFrameRect(frameRect foundation.Rect) foundation.Rect
	FrameRectForContentRect(contentRect foundation.Rect) foundation.Rect
	BeginSheet_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse))
	BeginCriticalSheet_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse))
	EndSheet(sheetWindow IWindow)
	EndSheet_ReturnCode(sheetWindow IWindow, returnCode ModalResponse)
	SetFrameOrigin(point foundation.Point)
	SetFrameTopLeftPoint(point foundation.Point)
	ConstrainFrameRect_ToScreen(frameRect foundation.Rect, screen IScreen) foundation.Rect
	CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point
	SetFrame_Display(frameRect foundation.Rect, flag bool)
	SetFrame_Display_Animate(frameRect foundation.Rect, displayFlag bool, animateFlag bool)
	AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval
	PerformZoom(sender objc.IObject)
	Zoom(sender objc.IObject)
	SetContentSize(size foundation.Size)
	OrderOut(sender objc.IObject)
	OrderBack(sender objc.IObject)
	OrderFront(sender objc.IObject)
	OrderFrontRegardless()
	OrderWindow_RelativeTo(place WindowOrderingMode, otherWin int)
	SetFrameUsingName(name WindowFrameAutosaveName) bool
	SetFrameUsingName_Force(name WindowFrameAutosaveName, force bool) bool
	SaveFrameUsingName(name WindowFrameAutosaveName)
	SetFrameAutosaveName(name WindowFrameAutosaveName) bool
	SetFrameFromString(_string WindowPersistableFrameDescriptor)
	MakeKeyWindow()
	MakeKeyAndOrderFront(sender objc.IObject)
	BecomeKeyWindow()
	ResignKeyWindow()
	MakeMainWindow()
	BecomeMainWindow()
	ResignMainWindow()
	ToggleToolbarShown(sender objc.IObject)
	RunToolbarCustomizationPalette(sender objc.IObject)
	AddChildWindow_Ordered(childWin IWindow, place WindowOrderingMode)
	RemoveChildWindow(childWin IWindow)
	EnableKeyEquivalentForDefaultButtonCell()
	DisableKeyEquivalentForDefaultButtonCell()
	FieldEditor_ForObject(createFlag bool, object objc.IObject) Text
	EndEditingFor(object objc.IObject)
	EnableCursorRects()
	DisableCursorRects()
	DiscardCursorRects()
	InvalidateCursorRectsForView(view IView)
	ResetCursorRects()
	StandardWindowButton(b WindowButton) Button
	AddTitlebarAccessoryViewController(childViewController ITitlebarAccessoryViewController)
	InsertTitlebarAccessoryViewController_AtIndex(childViewController ITitlebarAccessoryViewController, index int)
	RemoveTitlebarAccessoryViewControllerAtIndex(index int)
	AddTabbedWindow_Ordered(window IWindow, ordered WindowOrderingMode)
	MergeAllWindows(sender objc.IObject)
	SelectNextTab(sender objc.IObject)
	SelectPreviousTab(sender objc.IObject)
	MoveTabToNewWindow(sender objc.IObject)
	ToggleTabBar(sender objc.IObject)
	ToggleTabOverview(sender objc.IObject)
	NextEventMatchingMask(mask EventMask) Event
	NextEventMatchingMask_UntilDate_InMode_Dequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event
	DiscardEventsMatchingMask_BeforeEvent(mask EventMask, lastEvent IEvent)
	PostEvent_AtStart(event IEvent, flag bool)
	SendEvent(event IEvent)
	MakeFirstResponder(responder IResponder) bool
	SelectKeyViewPrecedingView(view IView)
	SelectKeyViewFollowingView(view IView)
	SelectPreviousKeyView(sender objc.IObject)
	SelectNextKeyView(sender objc.IObject)
	RecalculateKeyViewLoop()
	TrackEventsMatchingMask_Timeout_Mode_Handler(mask EventMask, timeout foundation.TimeInterval, mode foundation.RunLoopMode, trackingHandler func(event Event, stop *bool))
	PerformWindowDragWithEvent(event IEvent)
	DisableSnapshotRestoration()
	EnableSnapshotRestoration()
	Display()
	DisplayIfNeeded()
	DisableScreenUpdatesUntilFlush()
	Update()
	DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(image IImage, baseLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool)
	RegisterForDraggedTypes(newTypes []PasteboardType)
	UnregisterDraggedTypes()
	BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectFromScreen(rect foundation.Rect) foundation.Rect
	ConvertPointFromBacking(point foundation.Point) foundation.Point
	ConvertPointFromScreen(point foundation.Point) foundation.Point
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToScreen(rect foundation.Rect) foundation.Rect
	ConvertPointToBacking(point foundation.Point) foundation.Point
	ConvertPointToScreen(point foundation.Point) foundation.Point
	SetTitleWithRepresentedFilename(filename string)
	Center()
	PerformClose(sender objc.IObject)
	Close()
	PerformMiniaturize(sender objc.IObject)
	Miniaturize(sender objc.IObject)
	Deminiaturize(sender objc.IObject)
	Print(sender objc.IObject)
	DataWithEPSInsideRect(rect foundation.Rect) []byte
	DataWithPDFInsideRect(rect foundation.Rect) []byte
	UpdateConstraintsIfNeeded()
	LayoutIfNeeded()
	VisualizeConstraints(constraints []ILayoutConstraint)
	AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute
	SetAnchorAttribute_ForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation)
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	SetIsMiniaturized(flag bool)
	SetIsVisible(flag bool)
	SetIsZoomed(flag bool)
	HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object
	HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object
	HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object
	// deprecated
	GState() int
	// deprecated
	CanStoreColor() bool
	// deprecated
	EnableFlushWindow()
	// deprecated
	DisableFlushWindow()
	// deprecated
	FlushWindow()
	// deprecated
	FlushWindowIfNeeded()
	// deprecated
	CacheImageInRect(rect foundation.Rect)
	// deprecated
	RestoreCachedImage()
	// deprecated
	DiscardCachedImage()
	// deprecated
	UseOptimizedDrawing(flag bool)
	// deprecated
	ConvertBaseToScreen(point foundation.Point) foundation.Point
	// deprecated
	ConvertScreenToBase(point foundation.Point) foundation.Point
	// deprecated
	UserSpaceScaleFactor() float64
	InitWithWindowRef(windowRef unsafe.Pointer) Window
	Delegate() WindowDelegateWrapper
	SetDelegate(value WindowDelegate)
	ContentViewController() ViewController
	SetContentViewController(value IViewController)
	ContentView() View
	SetContentView(value IView)
	StyleMask() WindowStyleMask
	SetStyleMask(value WindowStyleMask)
	WorksWhenModal() bool
	AlphaValue() float64
	SetAlphaValue(value float64)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	ColorSpace() ColorSpace
	SetColorSpace(value IColorSpace)
	CanHide() bool
	SetCanHide(value bool)
	IsOnActiveSpace() bool
	HidesOnDeactivate() bool
	SetHidesOnDeactivate(value bool)
	CollectionBehavior() WindowCollectionBehavior
	SetCollectionBehavior(value WindowCollectionBehavior)
	IsOpaque() bool
	SetOpaque(value bool)
	HasShadow() bool
	SetHasShadow(value bool)
	PreventsApplicationTerminationWhenModal() bool
	SetPreventsApplicationTerminationWhenModal(value bool)
	DepthLimit() WindowDepth
	SetDepthLimit(value WindowDepth)
	HasDynamicDepthLimit() bool
	WindowNumber() int
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
	CanBecomeVisibleWithoutLogin() bool
	SetCanBecomeVisibleWithoutLogin(value bool)
	SharingType() WindowSharingType
	SetSharingType(value WindowSharingType)
	BackingType() BackingStoreType
	SetBackingType(value BackingStoreType)
	WindowController() WindowController
	SetWindowController(value IWindowController)
	AttachedSheet() Window
	IsSheet() bool
	SheetParent() Window
	Sheets() []Window
	Frame() foundation.Rect
	AspectRatio() foundation.Size
	SetAspectRatio(value foundation.Size)
	MinSize() foundation.Size
	SetMinSize(value foundation.Size)
	MaxSize() foundation.Size
	SetMaxSize(value foundation.Size)
	IsZoomed() bool
	ResizeFlags() EventModifierFlags
	ResizeIncrements() foundation.Size
	SetResizeIncrements(value foundation.Size)
	PreservesContentDuringLiveResize() bool
	SetPreservesContentDuringLiveResize(value bool)
	InLiveResize() bool
	ContentAspectRatio() foundation.Size
	SetContentAspectRatio(value foundation.Size)
	ContentMinSize() foundation.Size
	SetContentMinSize(value foundation.Size)
	ContentMaxSize() foundation.Size
	SetContentMaxSize(value foundation.Size)
	ContentResizeIncrements() foundation.Size
	SetContentResizeIncrements(value foundation.Size)
	ContentLayoutGuide() objc.Object
	ContentLayoutRect() foundation.Rect
	MaxFullScreenContentSize() foundation.Size
	SetMaxFullScreenContentSize(value foundation.Size)
	MinFullScreenContentSize() foundation.Size
	SetMinFullScreenContentSize(value foundation.Size)
	Level() WindowLevel
	SetLevel(value WindowLevel)
	IsVisible() bool
	OcclusionState() WindowOcclusionState
	FrameAutosaveName() WindowFrameAutosaveName
	StringWithSavedFrame() WindowPersistableFrameDescriptor
	IsKeyWindow() bool
	CanBecomeKeyWindow() bool
	IsMainWindow() bool
	CanBecomeMainWindow() bool
	Toolbar() Toolbar
	SetToolbar(value IToolbar)
	ChildWindows() []Window
	ParentWindow() Window
	SetParentWindow(value IWindow)
	DefaultButtonCell() ButtonCell
	SetDefaultButtonCell(value IButtonCell)
	IsExcludedFromWindowsMenu() bool
	SetExcludedFromWindowsMenu(value bool)
	AreCursorRectsEnabled() bool
	// deprecated
	ShowsToolbarButton() bool
	// deprecated
	SetShowsToolbarButton(value bool)
	TitlebarAppearsTransparent() bool
	SetTitlebarAppearsTransparent(value bool)
	ToolbarStyle() WindowToolbarStyle
	SetToolbarStyle(value WindowToolbarStyle)
	TitlebarSeparatorStyle() TitlebarSeparatorStyle
	SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle)
	WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection
	TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController
	SetTitlebarAccessoryViewControllers(value []ITitlebarAccessoryViewController)
	Tab() WindowTab
	TabbingIdentifier() WindowTabbingIdentifier
	SetTabbingIdentifier(value WindowTabbingIdentifier)
	TabbingMode() WindowTabbingMode
	SetTabbingMode(value WindowTabbingMode)
	TabbedWindows() []Window
	TabGroup() WindowTabGroup
	AllowsToolTipsWhenApplicationIsInactive() bool
	SetAllowsToolTipsWhenApplicationIsInactive(value bool)
	CurrentEvent() Event
	InitialFirstResponder() View
	SetInitialFirstResponder(value IView)
	FirstResponder() Responder
	KeyViewSelectionDirection() SelectionDirection
	AutorecalculatesKeyViewLoop() bool
	SetAutorecalculatesKeyViewLoop(value bool)
	AcceptsMouseMovedEvents() bool
	SetAcceptsMouseMovedEvents(value bool)
	IgnoresMouseEvents() bool
	SetIgnoresMouseEvents(value bool)
	MouseLocationOutsideOfEventStream() foundation.Point
	IsRestorable() bool
	SetRestorable(value bool)
	ViewsNeedDisplay() bool
	SetViewsNeedDisplay(value bool)
	AllowsConcurrentViewDrawing() bool
	SetAllowsConcurrentViewDrawing(value bool)
	AnimationBehavior() WindowAnimationBehavior
	SetAnimationBehavior(value WindowAnimationBehavior)
	IsDocumentEdited() bool
	SetDocumentEdited(value bool)
	BackingScaleFactor() float64
	Title() string
	SetTitle(value string)
	Subtitle() string
	SetSubtitle(value string)
	TitleVisibility() WindowTitleVisibility
	SetTitleVisibility(value WindowTitleVisibility)
	RepresentedFilename() string
	SetRepresentedFilename(value string)
	RepresentedURL() foundation.URL
	SetRepresentedURL(value foundation.IURL)
	Screen() Screen
	DeepestScreen() Screen
	DisplaysWhenScreenProfileChanges() bool
	SetDisplaysWhenScreenProfileChanges(value bool)
	IsMovableByWindowBackground() bool
	SetMovableByWindowBackground(value bool)
	IsMovable() bool
	SetMovable(value bool)
	IsReleasedWhenClosed() bool
	SetReleasedWhenClosed(value bool)
	IsMiniaturized() bool
	MiniwindowImage() Image
	SetMiniwindowImage(value IImage)
	MiniwindowTitle() string
	SetMiniwindowTitle(value string)
	DockTile() DockTile
	HasCloseBox() bool
	HasTitleBar() bool
	IsModalPanel() bool
	IsFloatingPanel() bool
	IsZoomable() bool
	IsResizable() bool
	IsMiniaturizable() bool
	OrderedIndex() int
	SetOrderedIndex(value int)
	// deprecated
	BackingLocation() WindowBackingLocation
	// deprecated
	PreferredBackingLocation() WindowBackingLocation
	// deprecated
	SetPreferredBackingLocation(value WindowBackingLocation)
	// deprecated
	IsOneShot() bool
	// deprecated
	SetOneShot(value bool)
	// deprecated
	ShowsResizeIndicator() bool
	// deprecated
	SetShowsResizeIndicator(value bool)
	// deprecated
	IsFlushWindowDisabled() bool
	// deprecated
	IsAutodisplay() bool
	// deprecated
	SetAutodisplay(value bool)
	// deprecated
	GraphicsContext() GraphicsContext
	WindowRef() unsafe.Pointer
}

type Window struct {
	Responder
}

func MakeWindow(ptr unsafe.Pointer) Window {
	return Window{
		Responder: MakeResponder(ptr),
	}
}

func (wc _WindowClass) WindowWithContentViewController(contentViewController IViewController) Window {
	rv := ffi.CallMethod[Window](wc, "windowWithContentViewController:", contentViewController)
	return rv
}

func (w_ Window) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	rv := ffi.CallMethod[Window](w_, "initWithContentRect:styleMask:backing:defer:", contentRect, style, backingStoreType, flag)
	return rv
}

func (w_ Window) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Window {
	rv := ffi.CallMethod[Window](w_, "initWithContentRect:styleMask:backing:defer:screen:", contentRect, style, backingStoreType, flag, screen)
	return rv
}

func (w_ Window) Init() Window {
	rv := ffi.CallMethod[Window](w_, "init")
	return rv
}

func (wc _WindowClass) Alloc() Window {
	rv := ffi.CallMethod[Window](wc, "alloc")
	return rv
}

func (wc _WindowClass) New() Window {
	rv := ffi.CallMethod[Window](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWindow() Window {
	return WindowClass.New()
}

func (w_ Window) ToggleFullScreen(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "toggleFullScreen:", sender)
}

func (w_ Window) SetDynamicDepthLimit(flag bool) {
	ffi.CallMethod[ffi.Void](w_, "setDynamicDepthLimit:", flag)
}

func (w_ Window) InvalidateShadow() {
	ffi.CallMethod[ffi.Void](w_, "invalidateShadow")
}

func (w_ Window) AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool {
	rv := ffi.CallMethod[bool](w_, "autorecalculatesContentBorderThicknessForEdge:", edge)
	return rv
}

func (w_ Window) SetAutorecalculatesContentBorderThickness_ForEdge(flag bool, edge foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](w_, "setAutorecalculatesContentBorderThickness:forEdge:", flag, edge)
}

func (w_ Window) ContentBorderThicknessForEdge(edge foundation.RectEdge) float64 {
	rv := ffi.CallMethod[float64](w_, "contentBorderThicknessForEdge:", edge)
	return rv
}

func (w_ Window) SetContentBorderThickness_ForEdge(thickness float64, edge foundation.RectEdge) {
	ffi.CallMethod[ffi.Void](w_, "setContentBorderThickness:forEdge:", thickness, edge)
}

func (wc _WindowClass) WindowNumbersWithOptions(options WindowNumberListOptions) []foundation.Number {
	rv := ffi.CallMethod[[]foundation.Number](wc, "windowNumbersWithOptions:", options)
	return rv
}

func (wc _WindowClass) ContentRectForFrameRect_StyleMask(fRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](wc, "contentRectForFrameRect:styleMask:", fRect, style)
	return rv
}

func (wc _WindowClass) FrameRectForContentRect_StyleMask(cRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](wc, "frameRectForContentRect:styleMask:", cRect, style)
	return rv
}

func (wc _WindowClass) MinFrameWidthWithTitle_StyleMask(title string, style WindowStyleMask) float64 {
	rv := ffi.CallMethod[float64](wc, "minFrameWidthWithTitle:styleMask:", title, style)
	return rv
}

func (w_ Window) ContentRectForFrameRect(frameRect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "contentRectForFrameRect:", frameRect)
	return rv
}

func (w_ Window) FrameRectForContentRect(contentRect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "frameRectForContentRect:", contentRect)
	return rv
}

func (w_ Window) BeginSheet_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	ffi.CallMethod[ffi.Void](w_, "beginSheet:completionHandler:", sheetWindow, handler)
}

func (w_ Window) BeginCriticalSheet_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	ffi.CallMethod[ffi.Void](w_, "beginCriticalSheet:completionHandler:", sheetWindow, handler)
}

func (w_ Window) EndSheet(sheetWindow IWindow) {
	ffi.CallMethod[ffi.Void](w_, "endSheet:", sheetWindow)
}

func (w_ Window) EndSheet_ReturnCode(sheetWindow IWindow, returnCode ModalResponse) {
	ffi.CallMethod[ffi.Void](w_, "endSheet:returnCode:", sheetWindow, returnCode)
}

func (w_ Window) SetFrameOrigin(point foundation.Point) {
	ffi.CallMethod[ffi.Void](w_, "setFrameOrigin:", point)
}

func (w_ Window) SetFrameTopLeftPoint(point foundation.Point) {
	ffi.CallMethod[ffi.Void](w_, "setFrameTopLeftPoint:", point)
}

func (w_ Window) ConstrainFrameRect_ToScreen(frameRect foundation.Rect, screen IScreen) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "constrainFrameRect:toScreen:", frameRect, screen)
	return rv
}

func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](w_, "cascadeTopLeftFromPoint:", topLeftPoint)
	return rv
}

func (w_ Window) SetFrame_Display(frameRect foundation.Rect, flag bool) {
	ffi.CallMethod[ffi.Void](w_, "setFrame:display:", frameRect, flag)
}

func (w_ Window) SetFrame_Display_Animate(frameRect foundation.Rect, displayFlag bool, animateFlag bool) {
	ffi.CallMethod[ffi.Void](w_, "setFrame:display:animate:", frameRect, displayFlag, animateFlag)
}

func (w_ Window) AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval {
	rv := ffi.CallMethod[foundation.TimeInterval](w_, "animationResizeTime:", newFrame)
	return rv
}

func (w_ Window) PerformZoom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "performZoom:", sender)
}

func (w_ Window) Zoom(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "zoom:", sender)
}

func (w_ Window) SetContentSize(size foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setContentSize:", size)
}

func (w_ Window) OrderOut(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "orderOut:", sender)
}

func (w_ Window) OrderBack(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "orderBack:", sender)
}

func (w_ Window) OrderFront(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "orderFront:", sender)
}

func (w_ Window) OrderFrontRegardless() {
	ffi.CallMethod[ffi.Void](w_, "orderFrontRegardless")
}

func (w_ Window) OrderWindow_RelativeTo(place WindowOrderingMode, otherWin int) {
	ffi.CallMethod[ffi.Void](w_, "orderWindow:relativeTo:", place, otherWin)
}

func (wc _WindowClass) RemoveFrameUsingName(name WindowFrameAutosaveName) {
	ffi.CallMethod[ffi.Void](wc, "removeFrameUsingName:", name)
}

func (w_ Window) SetFrameUsingName(name WindowFrameAutosaveName) bool {
	rv := ffi.CallMethod[bool](w_, "setFrameUsingName:", name)
	return rv
}

func (w_ Window) SetFrameUsingName_Force(name WindowFrameAutosaveName, force bool) bool {
	rv := ffi.CallMethod[bool](w_, "setFrameUsingName:force:", name, force)
	return rv
}

func (w_ Window) SaveFrameUsingName(name WindowFrameAutosaveName) {
	ffi.CallMethod[ffi.Void](w_, "saveFrameUsingName:", name)
}

func (w_ Window) SetFrameAutosaveName(name WindowFrameAutosaveName) bool {
	rv := ffi.CallMethod[bool](w_, "setFrameAutosaveName:", name)
	return rv
}

func (w_ Window) SetFrameFromString(_string WindowPersistableFrameDescriptor) {
	ffi.CallMethod[ffi.Void](w_, "setFrameFromString:", _string)
}

func (w_ Window) MakeKeyWindow() {
	ffi.CallMethod[ffi.Void](w_, "makeKeyWindow")
}

func (w_ Window) MakeKeyAndOrderFront(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "makeKeyAndOrderFront:", sender)
}

func (w_ Window) BecomeKeyWindow() {
	ffi.CallMethod[ffi.Void](w_, "becomeKeyWindow")
}

func (w_ Window) ResignKeyWindow() {
	ffi.CallMethod[ffi.Void](w_, "resignKeyWindow")
}

func (w_ Window) MakeMainWindow() {
	ffi.CallMethod[ffi.Void](w_, "makeMainWindow")
}

func (w_ Window) BecomeMainWindow() {
	ffi.CallMethod[ffi.Void](w_, "becomeMainWindow")
}

func (w_ Window) ResignMainWindow() {
	ffi.CallMethod[ffi.Void](w_, "resignMainWindow")
}

func (w_ Window) ToggleToolbarShown(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "toggleToolbarShown:", sender)
}

func (w_ Window) RunToolbarCustomizationPalette(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "runToolbarCustomizationPalette:", sender)
}

func (w_ Window) AddChildWindow_Ordered(childWin IWindow, place WindowOrderingMode) {
	ffi.CallMethod[ffi.Void](w_, "addChildWindow:ordered:", childWin, place)
}

func (w_ Window) RemoveChildWindow(childWin IWindow) {
	ffi.CallMethod[ffi.Void](w_, "removeChildWindow:", childWin)
}

func (w_ Window) EnableKeyEquivalentForDefaultButtonCell() {
	ffi.CallMethod[ffi.Void](w_, "enableKeyEquivalentForDefaultButtonCell")
}

func (w_ Window) DisableKeyEquivalentForDefaultButtonCell() {
	ffi.CallMethod[ffi.Void](w_, "disableKeyEquivalentForDefaultButtonCell")
}

func (w_ Window) FieldEditor_ForObject(createFlag bool, object objc.IObject) Text {
	rv := ffi.CallMethod[Text](w_, "fieldEditor:forObject:", createFlag, object)
	return rv
}

func (w_ Window) EndEditingFor(object objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "endEditingFor:", object)
}

func (w_ Window) EnableCursorRects() {
	ffi.CallMethod[ffi.Void](w_, "enableCursorRects")
}

func (w_ Window) DisableCursorRects() {
	ffi.CallMethod[ffi.Void](w_, "disableCursorRects")
}

func (w_ Window) DiscardCursorRects() {
	ffi.CallMethod[ffi.Void](w_, "discardCursorRects")
}

func (w_ Window) InvalidateCursorRectsForView(view IView) {
	ffi.CallMethod[ffi.Void](w_, "invalidateCursorRectsForView:", view)
}

func (w_ Window) ResetCursorRects() {
	ffi.CallMethod[ffi.Void](w_, "resetCursorRects")
}

func (wc _WindowClass) StandardWindowButton_ForStyleMask(b WindowButton, styleMask WindowStyleMask) Button {
	rv := ffi.CallMethod[Button](wc, "standardWindowButton:forStyleMask:", b, styleMask)
	return rv
}

func (w_ Window) StandardWindowButton(b WindowButton) Button {
	rv := ffi.CallMethod[Button](w_, "standardWindowButton:", b)
	return rv
}

func (w_ Window) AddTitlebarAccessoryViewController(childViewController ITitlebarAccessoryViewController) {
	ffi.CallMethod[ffi.Void](w_, "addTitlebarAccessoryViewController:", childViewController)
}

func (w_ Window) InsertTitlebarAccessoryViewController_AtIndex(childViewController ITitlebarAccessoryViewController, index int) {
	ffi.CallMethod[ffi.Void](w_, "insertTitlebarAccessoryViewController:atIndex:", childViewController, index)
}

func (w_ Window) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	ffi.CallMethod[ffi.Void](w_, "removeTitlebarAccessoryViewControllerAtIndex:", index)
}

func (w_ Window) AddTabbedWindow_Ordered(window IWindow, ordered WindowOrderingMode) {
	ffi.CallMethod[ffi.Void](w_, "addTabbedWindow:ordered:", window, ordered)
}

func (w_ Window) MergeAllWindows(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "mergeAllWindows:", sender)
}

func (w_ Window) SelectNextTab(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "selectNextTab:", sender)
}

func (w_ Window) SelectPreviousTab(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "selectPreviousTab:", sender)
}

func (w_ Window) MoveTabToNewWindow(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "moveTabToNewWindow:", sender)
}

func (w_ Window) ToggleTabBar(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "toggleTabBar:", sender)
}

func (w_ Window) ToggleTabOverview(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "toggleTabOverview:", sender)
}

func (w_ Window) NextEventMatchingMask(mask EventMask) Event {
	rv := ffi.CallMethod[Event](w_, "nextEventMatchingMask:", mask)
	return rv
}

func (w_ Window) NextEventMatchingMask_UntilDate_InMode_Dequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event {
	rv := ffi.CallMethod[Event](w_, "nextEventMatchingMask:untilDate:inMode:dequeue:", mask, expiration, mode, deqFlag)
	return rv
}

func (w_ Window) DiscardEventsMatchingMask_BeforeEvent(mask EventMask, lastEvent IEvent) {
	ffi.CallMethod[ffi.Void](w_, "discardEventsMatchingMask:beforeEvent:", mask, lastEvent)
}

func (w_ Window) PostEvent_AtStart(event IEvent, flag bool) {
	ffi.CallMethod[ffi.Void](w_, "postEvent:atStart:", event, flag)
}

func (w_ Window) SendEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](w_, "sendEvent:", event)
}

func (w_ Window) MakeFirstResponder(responder IResponder) bool {
	rv := ffi.CallMethod[bool](w_, "makeFirstResponder:", responder)
	return rv
}

func (w_ Window) SelectKeyViewPrecedingView(view IView) {
	ffi.CallMethod[ffi.Void](w_, "selectKeyViewPrecedingView:", view)
}

func (w_ Window) SelectKeyViewFollowingView(view IView) {
	ffi.CallMethod[ffi.Void](w_, "selectKeyViewFollowingView:", view)
}

func (w_ Window) SelectPreviousKeyView(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "selectPreviousKeyView:", sender)
}

func (w_ Window) SelectNextKeyView(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "selectNextKeyView:", sender)
}

func (w_ Window) RecalculateKeyViewLoop() {
	ffi.CallMethod[ffi.Void](w_, "recalculateKeyViewLoop")
}

func (wc _WindowClass) WindowNumberAtPoint_BelowWindowWithWindowNumber(point foundation.Point, windowNumber int) int {
	rv := ffi.CallMethod[int](wc, "windowNumberAtPoint:belowWindowWithWindowNumber:", point, windowNumber)
	return rv
}

func (w_ Window) TrackEventsMatchingMask_Timeout_Mode_Handler(mask EventMask, timeout foundation.TimeInterval, mode foundation.RunLoopMode, trackingHandler func(event Event, stop *bool)) {
	ffi.CallMethod[ffi.Void](w_, "trackEventsMatchingMask:timeout:mode:handler:", mask, timeout, mode, trackingHandler)
}

func (w_ Window) PerformWindowDragWithEvent(event IEvent) {
	ffi.CallMethod[ffi.Void](w_, "performWindowDragWithEvent:", event)
}

func (w_ Window) DisableSnapshotRestoration() {
	ffi.CallMethod[ffi.Void](w_, "disableSnapshotRestoration")
}

func (w_ Window) EnableSnapshotRestoration() {
	ffi.CallMethod[ffi.Void](w_, "enableSnapshotRestoration")
}

func (w_ Window) Display() {
	ffi.CallMethod[ffi.Void](w_, "display")
}

func (w_ Window) DisplayIfNeeded() {
	ffi.CallMethod[ffi.Void](w_, "displayIfNeeded")
}

func (w_ Window) DisableScreenUpdatesUntilFlush() {
	ffi.CallMethod[ffi.Void](w_, "disableScreenUpdatesUntilFlush")
}

func (w_ Window) Update() {
	ffi.CallMethod[ffi.Void](w_, "update")
}

func (w_ Window) DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(image IImage, baseLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool) {
	ffi.CallMethod[ffi.Void](w_, "dragImage:at:offset:event:pasteboard:source:slideBack:", image, baseLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}

func (w_ Window) RegisterForDraggedTypes(newTypes []PasteboardType) {
	ffi.CallMethod[ffi.Void](w_, "registerForDraggedTypes:", newTypes)
}

func (w_ Window) UnregisterDraggedTypes() {
	ffi.CallMethod[ffi.Void](w_, "unregisterDraggedTypes")
}

func (w_ Window) BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "backingAlignedRect:options:", rect, options)
	return rv
}

func (w_ Window) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "convertRectFromBacking:", rect)
	return rv
}

func (w_ Window) ConvertRectFromScreen(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "convertRectFromScreen:", rect)
	return rv
}

func (w_ Window) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](w_, "convertPointFromBacking:", point)
	return rv
}

func (w_ Window) ConvertPointFromScreen(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](w_, "convertPointFromScreen:", point)
	return rv
}

func (w_ Window) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "convertRectToBacking:", rect)
	return rv
}

func (w_ Window) ConvertRectToScreen(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "convertRectToScreen:", rect)
	return rv
}

func (w_ Window) ConvertPointToBacking(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](w_, "convertPointToBacking:", point)
	return rv
}

func (w_ Window) ConvertPointToScreen(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](w_, "convertPointToScreen:", point)
	return rv
}

func (w_ Window) SetTitleWithRepresentedFilename(filename string) {
	ffi.CallMethod[ffi.Void](w_, "setTitleWithRepresentedFilename:", filename)
}

func (w_ Window) Center() {
	ffi.CallMethod[ffi.Void](w_, "center")
}

func (w_ Window) PerformClose(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "performClose:", sender)
}

func (w_ Window) Close() {
	ffi.CallMethod[ffi.Void](w_, "close")
}

func (w_ Window) PerformMiniaturize(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "performMiniaturize:", sender)
}

func (w_ Window) Miniaturize(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "miniaturize:", sender)
}

func (w_ Window) Deminiaturize(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "deminiaturize:", sender)
}

func (w_ Window) Print(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](w_, "print:", sender)
}

func (w_ Window) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	rv := ffi.CallMethod[[]byte](w_, "dataWithEPSInsideRect:", rect)
	return rv
}

func (w_ Window) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	rv := ffi.CallMethod[[]byte](w_, "dataWithPDFInsideRect:", rect)
	return rv
}

func (w_ Window) UpdateConstraintsIfNeeded() {
	ffi.CallMethod[ffi.Void](w_, "updateConstraintsIfNeeded")
}

func (w_ Window) LayoutIfNeeded() {
	ffi.CallMethod[ffi.Void](w_, "layoutIfNeeded")
}

func (w_ Window) VisualizeConstraints(constraints []ILayoutConstraint) {
	ffi.CallMethod[ffi.Void](w_, "visualizeConstraints:", constraints)
}

func (w_ Window) AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute {
	rv := ffi.CallMethod[LayoutAttribute](w_, "anchorAttributeForOrientation:", orientation)
	return rv
}

func (w_ Window) SetAnchorAttribute_ForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation) {
	ffi.CallMethod[ffi.Void](w_, "setAnchorAttribute:forOrientation:", attr, orientation)
}

func (w_ Window) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := ffi.CallMethod[bool](w_, "canRepresentDisplayGamut:", displayGamut)
	return rv
}

func (w_ Window) SetIsMiniaturized(flag bool) {
	ffi.CallMethod[ffi.Void](w_, "setIsMiniaturized:", flag)
}

func (w_ Window) SetIsVisible(flag bool) {
	ffi.CallMethod[ffi.Void](w_, "setIsVisible:", flag)
}

func (w_ Window) SetIsZoomed(flag bool) {
	ffi.CallMethod[ffi.Void](w_, "setIsZoomed:", flag)
}

func (w_ Window) HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object {
	rv := ffi.CallMethod[objc.Object](w_, "handleCloseScriptCommand:", command)
	return rv
}

func (w_ Window) HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := ffi.CallMethod[objc.Object](w_, "handlePrintScriptCommand:", command)
	return rv
}

func (w_ Window) HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := ffi.CallMethod[objc.Object](w_, "handleSaveScriptCommand:", command)
	return rv
}

// deprecated
func (w_ Window) GState() int {
	rv := ffi.CallMethod[int](w_, "gState")
	return rv
}

// deprecated
func (w_ Window) CanStoreColor() bool {
	rv := ffi.CallMethod[bool](w_, "canStoreColor")
	return rv
}

// deprecated
func (w_ Window) EnableFlushWindow() {
	ffi.CallMethod[ffi.Void](w_, "enableFlushWindow")
}

// deprecated
func (w_ Window) DisableFlushWindow() {
	ffi.CallMethod[ffi.Void](w_, "disableFlushWindow")
}

// deprecated
func (w_ Window) FlushWindow() {
	ffi.CallMethod[ffi.Void](w_, "flushWindow")
}

// deprecated
func (w_ Window) FlushWindowIfNeeded() {
	ffi.CallMethod[ffi.Void](w_, "flushWindowIfNeeded")
}

// deprecated
func (wc _WindowClass) MenuChanged(menu IMenu) {
	ffi.CallMethod[ffi.Void](wc, "menuChanged:", menu)
}

// deprecated
func (w_ Window) CacheImageInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](w_, "cacheImageInRect:", rect)
}

// deprecated
func (w_ Window) RestoreCachedImage() {
	ffi.CallMethod[ffi.Void](w_, "restoreCachedImage")
}

// deprecated
func (w_ Window) DiscardCachedImage() {
	ffi.CallMethod[ffi.Void](w_, "discardCachedImage")
}

// deprecated
func (w_ Window) UseOptimizedDrawing(flag bool) {
	ffi.CallMethod[ffi.Void](w_, "useOptimizedDrawing:", flag)
}

// deprecated
func (w_ Window) ConvertBaseToScreen(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](w_, "convertBaseToScreen:", point)
	return rv
}

// deprecated
func (w_ Window) ConvertScreenToBase(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](w_, "convertScreenToBase:", point)
	return rv
}

// deprecated
func (w_ Window) UserSpaceScaleFactor() float64 {
	rv := ffi.CallMethod[float64](w_, "userSpaceScaleFactor")
	return rv
}

func (w_ Window) InitWithWindowRef(windowRef unsafe.Pointer) Window {
	rv := ffi.CallMethod[Window](w_, "initWithWindowRef:", windowRef)
	return rv
}

func (w_ Window) Delegate() WindowDelegateWrapper {
	rv := ffi.CallMethod[WindowDelegateWrapper](w_, "delegate")
	return rv
}

func (w_ Window) SetDelegate(value WindowDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(w_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](w_, "setDelegate:", po)
}

func (w_ Window) ContentViewController() ViewController {
	rv := ffi.CallMethod[ViewController](w_, "contentViewController")
	return rv
}

func (w_ Window) SetContentViewController(value IViewController) {
	ffi.CallMethod[ffi.Void](w_, "setContentViewController:", value)
}

func (w_ Window) ContentView() View {
	rv := ffi.CallMethod[View](w_, "contentView")
	return rv
}

func (w_ Window) SetContentView(value IView) {
	ffi.CallMethod[ffi.Void](w_, "setContentView:", value)
}

func (w_ Window) StyleMask() WindowStyleMask {
	rv := ffi.CallMethod[WindowStyleMask](w_, "styleMask")
	return rv
}

func (w_ Window) SetStyleMask(value WindowStyleMask) {
	ffi.CallMethod[ffi.Void](w_, "setStyleMask:", value)
}

func (w_ Window) WorksWhenModal() bool {
	rv := ffi.CallMethod[bool](w_, "worksWhenModal")
	return rv
}

func (w_ Window) AlphaValue() float64 {
	rv := ffi.CallMethod[float64](w_, "alphaValue")
	return rv
}

func (w_ Window) SetAlphaValue(value float64) {
	ffi.CallMethod[ffi.Void](w_, "setAlphaValue:", value)
}

func (w_ Window) BackgroundColor() Color {
	rv := ffi.CallMethod[Color](w_, "backgroundColor")
	return rv
}

func (w_ Window) SetBackgroundColor(value IColor) {
	ffi.CallMethod[ffi.Void](w_, "setBackgroundColor:", value)
}

func (w_ Window) ColorSpace() ColorSpace {
	rv := ffi.CallMethod[ColorSpace](w_, "colorSpace")
	return rv
}

func (w_ Window) SetColorSpace(value IColorSpace) {
	ffi.CallMethod[ffi.Void](w_, "setColorSpace:", value)
}

func (w_ Window) CanHide() bool {
	rv := ffi.CallMethod[bool](w_, "canHide")
	return rv
}

func (w_ Window) SetCanHide(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setCanHide:", value)
}

func (w_ Window) IsOnActiveSpace() bool {
	rv := ffi.CallMethod[bool](w_, "isOnActiveSpace")
	return rv
}

func (w_ Window) HidesOnDeactivate() bool {
	rv := ffi.CallMethod[bool](w_, "hidesOnDeactivate")
	return rv
}

func (w_ Window) SetHidesOnDeactivate(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setHidesOnDeactivate:", value)
}

func (w_ Window) CollectionBehavior() WindowCollectionBehavior {
	rv := ffi.CallMethod[WindowCollectionBehavior](w_, "collectionBehavior")
	return rv
}

func (w_ Window) SetCollectionBehavior(value WindowCollectionBehavior) {
	ffi.CallMethod[ffi.Void](w_, "setCollectionBehavior:", value)
}

func (w_ Window) IsOpaque() bool {
	rv := ffi.CallMethod[bool](w_, "isOpaque")
	return rv
}

func (w_ Window) SetOpaque(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setOpaque:", value)
}

func (w_ Window) HasShadow() bool {
	rv := ffi.CallMethod[bool](w_, "hasShadow")
	return rv
}

func (w_ Window) SetHasShadow(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setHasShadow:", value)
}

func (w_ Window) PreventsApplicationTerminationWhenModal() bool {
	rv := ffi.CallMethod[bool](w_, "preventsApplicationTerminationWhenModal")
	return rv
}

func (w_ Window) SetPreventsApplicationTerminationWhenModal(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setPreventsApplicationTerminationWhenModal:", value)
}

func (w_ Window) DepthLimit() WindowDepth {
	rv := ffi.CallMethod[WindowDepth](w_, "depthLimit")
	return rv
}

func (w_ Window) SetDepthLimit(value WindowDepth) {
	ffi.CallMethod[ffi.Void](w_, "setDepthLimit:", value)
}

func (w_ Window) HasDynamicDepthLimit() bool {
	rv := ffi.CallMethod[bool](w_, "hasDynamicDepthLimit")
	return rv
}

func (wc _WindowClass) DefaultDepthLimit() WindowDepth {
	rv := ffi.CallMethod[WindowDepth](wc, "defaultDepthLimit")
	return rv
}

func (w_ Window) WindowNumber() int {
	rv := ffi.CallMethod[int](w_, "windowNumber")
	return rv
}

func (w_ Window) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := ffi.CallMethod[map[DeviceDescriptionKey]objc.Object](w_, "deviceDescription")
	return rv
}

func (w_ Window) CanBecomeVisibleWithoutLogin() bool {
	rv := ffi.CallMethod[bool](w_, "canBecomeVisibleWithoutLogin")
	return rv
}

func (w_ Window) SetCanBecomeVisibleWithoutLogin(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setCanBecomeVisibleWithoutLogin:", value)
}

func (w_ Window) SharingType() WindowSharingType {
	rv := ffi.CallMethod[WindowSharingType](w_, "sharingType")
	return rv
}

func (w_ Window) SetSharingType(value WindowSharingType) {
	ffi.CallMethod[ffi.Void](w_, "setSharingType:", value)
}

func (w_ Window) BackingType() BackingStoreType {
	rv := ffi.CallMethod[BackingStoreType](w_, "backingType")
	return rv
}

func (w_ Window) SetBackingType(value BackingStoreType) {
	ffi.CallMethod[ffi.Void](w_, "setBackingType:", value)
}

func (w_ Window) WindowController() WindowController {
	rv := ffi.CallMethod[WindowController](w_, "windowController")
	return rv
}

func (w_ Window) SetWindowController(value IWindowController) {
	ffi.CallMethod[ffi.Void](w_, "setWindowController:", value)
}

func (w_ Window) AttachedSheet() Window {
	rv := ffi.CallMethod[Window](w_, "attachedSheet")
	return rv
}

func (w_ Window) IsSheet() bool {
	rv := ffi.CallMethod[bool](w_, "isSheet")
	return rv
}

func (w_ Window) SheetParent() Window {
	rv := ffi.CallMethod[Window](w_, "sheetParent")
	return rv
}

func (w_ Window) Sheets() []Window {
	rv := ffi.CallMethod[[]Window](w_, "sheets")
	return rv
}

func (w_ Window) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "frame")
	return rv
}

func (w_ Window) AspectRatio() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "aspectRatio")
	return rv
}

func (w_ Window) SetAspectRatio(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setAspectRatio:", value)
}

func (w_ Window) MinSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "minSize")
	return rv
}

func (w_ Window) SetMinSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setMinSize:", value)
}

func (w_ Window) MaxSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "maxSize")
	return rv
}

func (w_ Window) SetMaxSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setMaxSize:", value)
}

func (w_ Window) IsZoomed() bool {
	rv := ffi.CallMethod[bool](w_, "isZoomed")
	return rv
}

func (w_ Window) ResizeFlags() EventModifierFlags {
	rv := ffi.CallMethod[EventModifierFlags](w_, "resizeFlags")
	return rv
}

func (w_ Window) ResizeIncrements() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "resizeIncrements")
	return rv
}

func (w_ Window) SetResizeIncrements(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setResizeIncrements:", value)
}

func (w_ Window) PreservesContentDuringLiveResize() bool {
	rv := ffi.CallMethod[bool](w_, "preservesContentDuringLiveResize")
	return rv
}

func (w_ Window) SetPreservesContentDuringLiveResize(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setPreservesContentDuringLiveResize:", value)
}

func (w_ Window) InLiveResize() bool {
	rv := ffi.CallMethod[bool](w_, "inLiveResize")
	return rv
}

func (w_ Window) ContentAspectRatio() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "contentAspectRatio")
	return rv
}

func (w_ Window) SetContentAspectRatio(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setContentAspectRatio:", value)
}

func (w_ Window) ContentMinSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "contentMinSize")
	return rv
}

func (w_ Window) SetContentMinSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setContentMinSize:", value)
}

func (w_ Window) ContentMaxSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "contentMaxSize")
	return rv
}

func (w_ Window) SetContentMaxSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setContentMaxSize:", value)
}

func (w_ Window) ContentResizeIncrements() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "contentResizeIncrements")
	return rv
}

func (w_ Window) SetContentResizeIncrements(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setContentResizeIncrements:", value)
}

func (w_ Window) ContentLayoutGuide() objc.Object {
	rv := ffi.CallMethod[objc.Object](w_, "contentLayoutGuide")
	return rv
}

func (w_ Window) ContentLayoutRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "contentLayoutRect")
	return rv
}

func (w_ Window) MaxFullScreenContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "maxFullScreenContentSize")
	return rv
}

func (w_ Window) SetMaxFullScreenContentSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setMaxFullScreenContentSize:", value)
}

func (w_ Window) MinFullScreenContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "minFullScreenContentSize")
	return rv
}

func (w_ Window) SetMinFullScreenContentSize(value foundation.Size) {
	ffi.CallMethod[ffi.Void](w_, "setMinFullScreenContentSize:", value)
}

func (w_ Window) Level() WindowLevel {
	rv := ffi.CallMethod[WindowLevel](w_, "level")
	return rv
}

func (w_ Window) SetLevel(value WindowLevel) {
	ffi.CallMethod[ffi.Void](w_, "setLevel:", value)
}

func (w_ Window) IsVisible() bool {
	rv := ffi.CallMethod[bool](w_, "isVisible")
	return rv
}

func (w_ Window) OcclusionState() WindowOcclusionState {
	rv := ffi.CallMethod[WindowOcclusionState](w_, "occlusionState")
	return rv
}

func (w_ Window) FrameAutosaveName() WindowFrameAutosaveName {
	rv := ffi.CallMethod[WindowFrameAutosaveName](w_, "frameAutosaveName")
	return rv
}

func (w_ Window) StringWithSavedFrame() WindowPersistableFrameDescriptor {
	rv := ffi.CallMethod[WindowPersistableFrameDescriptor](w_, "stringWithSavedFrame")
	return rv
}

func (w_ Window) IsKeyWindow() bool {
	rv := ffi.CallMethod[bool](w_, "isKeyWindow")
	return rv
}

func (w_ Window) CanBecomeKeyWindow() bool {
	rv := ffi.CallMethod[bool](w_, "canBecomeKeyWindow")
	return rv
}

func (w_ Window) IsMainWindow() bool {
	rv := ffi.CallMethod[bool](w_, "isMainWindow")
	return rv
}

func (w_ Window) CanBecomeMainWindow() bool {
	rv := ffi.CallMethod[bool](w_, "canBecomeMainWindow")
	return rv
}

func (w_ Window) Toolbar() Toolbar {
	rv := ffi.CallMethod[Toolbar](w_, "toolbar")
	return rv
}

func (w_ Window) SetToolbar(value IToolbar) {
	ffi.CallMethod[ffi.Void](w_, "setToolbar:", value)
}

func (w_ Window) ChildWindows() []Window {
	rv := ffi.CallMethod[[]Window](w_, "childWindows")
	return rv
}

func (w_ Window) ParentWindow() Window {
	rv := ffi.CallMethod[Window](w_, "parentWindow")
	return rv
}

func (w_ Window) SetParentWindow(value IWindow) {
	ffi.CallMethod[ffi.Void](w_, "setParentWindow:", value)
}

func (w_ Window) DefaultButtonCell() ButtonCell {
	rv := ffi.CallMethod[ButtonCell](w_, "defaultButtonCell")
	return rv
}

func (w_ Window) SetDefaultButtonCell(value IButtonCell) {
	ffi.CallMethod[ffi.Void](w_, "setDefaultButtonCell:", value)
}

func (w_ Window) IsExcludedFromWindowsMenu() bool {
	rv := ffi.CallMethod[bool](w_, "isExcludedFromWindowsMenu")
	return rv
}

func (w_ Window) SetExcludedFromWindowsMenu(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setExcludedFromWindowsMenu:", value)
}

func (w_ Window) AreCursorRectsEnabled() bool {
	rv := ffi.CallMethod[bool](w_, "areCursorRectsEnabled")
	return rv
}

// deprecated
func (w_ Window) ShowsToolbarButton() bool {
	rv := ffi.CallMethod[bool](w_, "showsToolbarButton")
	return rv
}

// deprecated
func (w_ Window) SetShowsToolbarButton(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setShowsToolbarButton:", value)
}

func (w_ Window) TitlebarAppearsTransparent() bool {
	rv := ffi.CallMethod[bool](w_, "titlebarAppearsTransparent")
	return rv
}

func (w_ Window) SetTitlebarAppearsTransparent(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setTitlebarAppearsTransparent:", value)
}

func (w_ Window) ToolbarStyle() WindowToolbarStyle {
	rv := ffi.CallMethod[WindowToolbarStyle](w_, "toolbarStyle")
	return rv
}

func (w_ Window) SetToolbarStyle(value WindowToolbarStyle) {
	ffi.CallMethod[ffi.Void](w_, "setToolbarStyle:", value)
}

func (w_ Window) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	rv := ffi.CallMethod[TitlebarSeparatorStyle](w_, "titlebarSeparatorStyle")
	return rv
}

func (w_ Window) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	ffi.CallMethod[ffi.Void](w_, "setTitlebarSeparatorStyle:", value)
}

func (w_ Window) WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection {
	rv := ffi.CallMethod[UserInterfaceLayoutDirection](w_, "windowTitlebarLayoutDirection")
	return rv
}

func (w_ Window) TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController {
	rv := ffi.CallMethod[[]TitlebarAccessoryViewController](w_, "titlebarAccessoryViewControllers")
	return rv
}

func (w_ Window) SetTitlebarAccessoryViewControllers(value []ITitlebarAccessoryViewController) {
	ffi.CallMethod[ffi.Void](w_, "setTitlebarAccessoryViewControllers:", value)
}

func (wc _WindowClass) AllowsAutomaticWindowTabbing() bool {
	rv := ffi.CallMethod[bool](wc, "allowsAutomaticWindowTabbing")
	return rv
}

func (wc _WindowClass) SetAllowsAutomaticWindowTabbing(value bool) {
	ffi.CallMethod[ffi.Void](wc, "setAllowsAutomaticWindowTabbing:", value)
}

func (wc _WindowClass) UserTabbingPreference() WindowUserTabbingPreference {
	rv := ffi.CallMethod[WindowUserTabbingPreference](wc, "userTabbingPreference")
	return rv
}

func (w_ Window) Tab() WindowTab {
	rv := ffi.CallMethod[WindowTab](w_, "tab")
	return rv
}

func (w_ Window) TabbingIdentifier() WindowTabbingIdentifier {
	rv := ffi.CallMethod[WindowTabbingIdentifier](w_, "tabbingIdentifier")
	return rv
}

func (w_ Window) SetTabbingIdentifier(value WindowTabbingIdentifier) {
	ffi.CallMethod[ffi.Void](w_, "setTabbingIdentifier:", value)
}

func (w_ Window) TabbingMode() WindowTabbingMode {
	rv := ffi.CallMethod[WindowTabbingMode](w_, "tabbingMode")
	return rv
}

func (w_ Window) SetTabbingMode(value WindowTabbingMode) {
	ffi.CallMethod[ffi.Void](w_, "setTabbingMode:", value)
}

func (w_ Window) TabbedWindows() []Window {
	rv := ffi.CallMethod[[]Window](w_, "tabbedWindows")
	return rv
}

func (w_ Window) TabGroup() WindowTabGroup {
	rv := ffi.CallMethod[WindowTabGroup](w_, "tabGroup")
	return rv
}

func (w_ Window) AllowsToolTipsWhenApplicationIsInactive() bool {
	rv := ffi.CallMethod[bool](w_, "allowsToolTipsWhenApplicationIsInactive")
	return rv
}

func (w_ Window) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAllowsToolTipsWhenApplicationIsInactive:", value)
}

func (w_ Window) CurrentEvent() Event {
	rv := ffi.CallMethod[Event](w_, "currentEvent")
	return rv
}

func (w_ Window) InitialFirstResponder() View {
	rv := ffi.CallMethod[View](w_, "initialFirstResponder")
	return rv
}

func (w_ Window) SetInitialFirstResponder(value IView) {
	ffi.CallMethod[ffi.Void](w_, "setInitialFirstResponder:", value)
}

func (w_ Window) FirstResponder() Responder {
	rv := ffi.CallMethod[Responder](w_, "firstResponder")
	return rv
}

func (w_ Window) KeyViewSelectionDirection() SelectionDirection {
	rv := ffi.CallMethod[SelectionDirection](w_, "keyViewSelectionDirection")
	return rv
}

func (w_ Window) AutorecalculatesKeyViewLoop() bool {
	rv := ffi.CallMethod[bool](w_, "autorecalculatesKeyViewLoop")
	return rv
}

func (w_ Window) SetAutorecalculatesKeyViewLoop(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAutorecalculatesKeyViewLoop:", value)
}

func (w_ Window) AcceptsMouseMovedEvents() bool {
	rv := ffi.CallMethod[bool](w_, "acceptsMouseMovedEvents")
	return rv
}

func (w_ Window) SetAcceptsMouseMovedEvents(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAcceptsMouseMovedEvents:", value)
}

func (w_ Window) IgnoresMouseEvents() bool {
	rv := ffi.CallMethod[bool](w_, "ignoresMouseEvents")
	return rv
}

func (w_ Window) SetIgnoresMouseEvents(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setIgnoresMouseEvents:", value)
}

func (w_ Window) MouseLocationOutsideOfEventStream() foundation.Point {
	rv := ffi.CallMethod[foundation.Point](w_, "mouseLocationOutsideOfEventStream")
	return rv
}

func (w_ Window) IsRestorable() bool {
	rv := ffi.CallMethod[bool](w_, "isRestorable")
	return rv
}

func (w_ Window) SetRestorable(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setRestorable:", value)
}

func (w_ Window) ViewsNeedDisplay() bool {
	rv := ffi.CallMethod[bool](w_, "viewsNeedDisplay")
	return rv
}

func (w_ Window) SetViewsNeedDisplay(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setViewsNeedDisplay:", value)
}

func (w_ Window) AllowsConcurrentViewDrawing() bool {
	rv := ffi.CallMethod[bool](w_, "allowsConcurrentViewDrawing")
	return rv
}

func (w_ Window) SetAllowsConcurrentViewDrawing(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAllowsConcurrentViewDrawing:", value)
}

func (w_ Window) AnimationBehavior() WindowAnimationBehavior {
	rv := ffi.CallMethod[WindowAnimationBehavior](w_, "animationBehavior")
	return rv
}

func (w_ Window) SetAnimationBehavior(value WindowAnimationBehavior) {
	ffi.CallMethod[ffi.Void](w_, "setAnimationBehavior:", value)
}

func (w_ Window) IsDocumentEdited() bool {
	rv := ffi.CallMethod[bool](w_, "isDocumentEdited")
	return rv
}

func (w_ Window) SetDocumentEdited(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setDocumentEdited:", value)
}

func (w_ Window) BackingScaleFactor() float64 {
	rv := ffi.CallMethod[float64](w_, "backingScaleFactor")
	return rv
}

func (w_ Window) Title() string {
	rv := ffi.CallMethod[string](w_, "title")
	return rv
}

func (w_ Window) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](w_, "setTitle:", value)
}

func (w_ Window) Subtitle() string {
	rv := ffi.CallMethod[string](w_, "subtitle")
	return rv
}

func (w_ Window) SetSubtitle(value string) {
	ffi.CallMethod[ffi.Void](w_, "setSubtitle:", value)
}

func (w_ Window) TitleVisibility() WindowTitleVisibility {
	rv := ffi.CallMethod[WindowTitleVisibility](w_, "titleVisibility")
	return rv
}

func (w_ Window) SetTitleVisibility(value WindowTitleVisibility) {
	ffi.CallMethod[ffi.Void](w_, "setTitleVisibility:", value)
}

func (w_ Window) RepresentedFilename() string {
	rv := ffi.CallMethod[string](w_, "representedFilename")
	return rv
}

func (w_ Window) SetRepresentedFilename(value string) {
	ffi.CallMethod[ffi.Void](w_, "setRepresentedFilename:", value)
}

func (w_ Window) RepresentedURL() foundation.URL {
	rv := ffi.CallMethod[foundation.URL](w_, "representedURL")
	return rv
}

func (w_ Window) SetRepresentedURL(value foundation.IURL) {
	ffi.CallMethod[ffi.Void](w_, "setRepresentedURL:", value)
}

func (w_ Window) Screen() Screen {
	rv := ffi.CallMethod[Screen](w_, "screen")
	return rv
}

func (w_ Window) DeepestScreen() Screen {
	rv := ffi.CallMethod[Screen](w_, "deepestScreen")
	return rv
}

func (w_ Window) DisplaysWhenScreenProfileChanges() bool {
	rv := ffi.CallMethod[bool](w_, "displaysWhenScreenProfileChanges")
	return rv
}

func (w_ Window) SetDisplaysWhenScreenProfileChanges(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setDisplaysWhenScreenProfileChanges:", value)
}

func (w_ Window) IsMovableByWindowBackground() bool {
	rv := ffi.CallMethod[bool](w_, "isMovableByWindowBackground")
	return rv
}

func (w_ Window) SetMovableByWindowBackground(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setMovableByWindowBackground:", value)
}

func (w_ Window) IsMovable() bool {
	rv := ffi.CallMethod[bool](w_, "isMovable")
	return rv
}

func (w_ Window) SetMovable(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setMovable:", value)
}

func (w_ Window) IsReleasedWhenClosed() bool {
	rv := ffi.CallMethod[bool](w_, "isReleasedWhenClosed")
	return rv
}

func (w_ Window) SetReleasedWhenClosed(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setReleasedWhenClosed:", value)
}

func (w_ Window) IsMiniaturized() bool {
	rv := ffi.CallMethod[bool](w_, "isMiniaturized")
	return rv
}

func (w_ Window) MiniwindowImage() Image {
	rv := ffi.CallMethod[Image](w_, "miniwindowImage")
	return rv
}

func (w_ Window) SetMiniwindowImage(value IImage) {
	ffi.CallMethod[ffi.Void](w_, "setMiniwindowImage:", value)
}

func (w_ Window) MiniwindowTitle() string {
	rv := ffi.CallMethod[string](w_, "miniwindowTitle")
	return rv
}

func (w_ Window) SetMiniwindowTitle(value string) {
	ffi.CallMethod[ffi.Void](w_, "setMiniwindowTitle:", value)
}

func (w_ Window) DockTile() DockTile {
	rv := ffi.CallMethod[DockTile](w_, "dockTile")
	return rv
}

func (w_ Window) HasCloseBox() bool {
	rv := ffi.CallMethod[bool](w_, "hasCloseBox")
	return rv
}

func (w_ Window) HasTitleBar() bool {
	rv := ffi.CallMethod[bool](w_, "hasTitleBar")
	return rv
}

func (w_ Window) IsModalPanel() bool {
	rv := ffi.CallMethod[bool](w_, "isModalPanel")
	return rv
}

func (w_ Window) IsFloatingPanel() bool {
	rv := ffi.CallMethod[bool](w_, "isFloatingPanel")
	return rv
}

func (w_ Window) IsZoomable() bool {
	rv := ffi.CallMethod[bool](w_, "isZoomable")
	return rv
}

func (w_ Window) IsResizable() bool {
	rv := ffi.CallMethod[bool](w_, "isResizable")
	return rv
}

func (w_ Window) IsMiniaturizable() bool {
	rv := ffi.CallMethod[bool](w_, "isMiniaturizable")
	return rv
}

func (w_ Window) OrderedIndex() int {
	rv := ffi.CallMethod[int](w_, "orderedIndex")
	return rv
}

func (w_ Window) SetOrderedIndex(value int) {
	ffi.CallMethod[ffi.Void](w_, "setOrderedIndex:", value)
}

// deprecated
func (w_ Window) BackingLocation() WindowBackingLocation {
	rv := ffi.CallMethod[WindowBackingLocation](w_, "backingLocation")
	return rv
}

// deprecated
func (w_ Window) PreferredBackingLocation() WindowBackingLocation {
	rv := ffi.CallMethod[WindowBackingLocation](w_, "preferredBackingLocation")
	return rv
}

// deprecated
func (w_ Window) SetPreferredBackingLocation(value WindowBackingLocation) {
	ffi.CallMethod[ffi.Void](w_, "setPreferredBackingLocation:", value)
}

// deprecated
func (w_ Window) IsOneShot() bool {
	rv := ffi.CallMethod[bool](w_, "isOneShot")
	return rv
}

// deprecated
func (w_ Window) SetOneShot(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setOneShot:", value)
}

// deprecated
func (w_ Window) ShowsResizeIndicator() bool {
	rv := ffi.CallMethod[bool](w_, "showsResizeIndicator")
	return rv
}

// deprecated
func (w_ Window) SetShowsResizeIndicator(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setShowsResizeIndicator:", value)
}

// deprecated
func (w_ Window) IsFlushWindowDisabled() bool {
	rv := ffi.CallMethod[bool](w_, "isFlushWindowDisabled")
	return rv
}

// deprecated
func (w_ Window) IsAutodisplay() bool {
	rv := ffi.CallMethod[bool](w_, "isAutodisplay")
	return rv
}

// deprecated
func (w_ Window) SetAutodisplay(value bool) {
	ffi.CallMethod[ffi.Void](w_, "setAutodisplay:", value)
}

// deprecated
func (w_ Window) GraphicsContext() GraphicsContext {
	rv := ffi.CallMethod[GraphicsContext](w_, "graphicsContext")
	return rv
}

func (w_ Window) WindowRef() unsafe.Pointer {
	rv := ffi.CallMethod[unsafe.Pointer](w_, "windowRef")
	return rv
}

type WindowDelegate interface {
	ImplementsWindow_WillPositionSheet_UsingRect() bool
	// optional
	Window_WillPositionSheet_UsingRect(window Window, sheet Window, rect foundation.Rect) foundation.Rect
	ImplementsWindowWillBeginSheet() bool
	// optional
	WindowWillBeginSheet(notification foundation.Notification)
	ImplementsWindowDidEndSheet() bool
	// optional
	WindowDidEndSheet(notification foundation.Notification)
	ImplementsWindowWillResize_ToSize() bool
	// optional
	WindowWillResize_ToSize(sender Window, frameSize foundation.Size) foundation.Size
	ImplementsWindowDidResize() bool
	// optional
	WindowDidResize(notification foundation.Notification)
	ImplementsWindowWillStartLiveResize() bool
	// optional
	WindowWillStartLiveResize(notification foundation.Notification)
	ImplementsWindowDidEndLiveResize() bool
	// optional
	WindowDidEndLiveResize(notification foundation.Notification)
	ImplementsWindowWillMiniaturize() bool
	// optional
	WindowWillMiniaturize(notification foundation.Notification)
	ImplementsWindowDidMiniaturize() bool
	// optional
	WindowDidMiniaturize(notification foundation.Notification)
	ImplementsWindowDidDeminiaturize() bool
	// optional
	WindowDidDeminiaturize(notification foundation.Notification)
	ImplementsWindowWillUseStandardFrame_DefaultFrame() bool
	// optional
	WindowWillUseStandardFrame_DefaultFrame(window Window, newFrame foundation.Rect) foundation.Rect
	ImplementsWindowShouldZoom_ToFrame() bool
	// optional
	WindowShouldZoom_ToFrame(window Window, newFrame foundation.Rect) bool
	ImplementsWindow_WillUseFullScreenContentSize() bool
	// optional
	Window_WillUseFullScreenContentSize(window Window, proposedSize foundation.Size) foundation.Size
	ImplementsWindow_WillUseFullScreenPresentationOptions() bool
	// optional
	Window_WillUseFullScreenPresentationOptions(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	ImplementsWindowWillEnterFullScreen() bool
	// optional
	WindowWillEnterFullScreen(notification foundation.Notification)
	ImplementsWindowDidEnterFullScreen() bool
	// optional
	WindowDidEnterFullScreen(notification foundation.Notification)
	ImplementsWindowWillExitFullScreen() bool
	// optional
	WindowWillExitFullScreen(notification foundation.Notification)
	ImplementsWindowDidExitFullScreen() bool
	// optional
	WindowDidExitFullScreen(notification foundation.Notification)
	ImplementsCustomWindowsToEnterFullScreenForWindow() bool
	// optional
	CustomWindowsToEnterFullScreenForWindow(window Window) []IWindow
	ImplementsCustomWindowsToEnterFullScreenForWindow_OnScreen() bool
	// optional
	CustomWindowsToEnterFullScreenForWindow_OnScreen(window Window, screen Screen) []IWindow
	ImplementsWindow_StartCustomAnimationToEnterFullScreenWithDuration() bool
	// optional
	Window_StartCustomAnimationToEnterFullScreenWithDuration(window Window, duration foundation.TimeInterval)
	ImplementsWindow_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration() bool
	// optional
	Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(window Window, screen Screen, duration foundation.TimeInterval)
	ImplementsWindowDidFailToEnterFullScreen() bool
	// optional
	WindowDidFailToEnterFullScreen(window Window)
	ImplementsCustomWindowsToExitFullScreenForWindow() bool
	// optional
	CustomWindowsToExitFullScreenForWindow(window Window) []IWindow
	ImplementsWindow_StartCustomAnimationToExitFullScreenWithDuration() bool
	// optional
	Window_StartCustomAnimationToExitFullScreenWithDuration(window Window, duration foundation.TimeInterval)
	ImplementsWindowDidFailToExitFullScreen() bool
	// optional
	WindowDidFailToExitFullScreen(window Window)
	ImplementsWindowWillMove() bool
	// optional
	WindowWillMove(notification foundation.Notification)
	ImplementsWindowDidMove() bool
	// optional
	WindowDidMove(notification foundation.Notification)
	ImplementsWindowDidChangeScreen() bool
	// optional
	WindowDidChangeScreen(notification foundation.Notification)
	ImplementsWindowDidChangeScreenProfile() bool
	// optional
	WindowDidChangeScreenProfile(notification foundation.Notification)
	ImplementsWindowDidChangeBackingProperties() bool
	// optional
	WindowDidChangeBackingProperties(notification foundation.Notification)
	ImplementsWindowShouldClose() bool
	// optional
	WindowShouldClose(sender Window) bool
	ImplementsWindowWillClose() bool
	// optional
	WindowWillClose(notification foundation.Notification)
	ImplementsWindowDidBecomeKey() bool
	// optional
	WindowDidBecomeKey(notification foundation.Notification)
	ImplementsWindowDidResignKey() bool
	// optional
	WindowDidResignKey(notification foundation.Notification)
	ImplementsWindowDidBecomeMain() bool
	// optional
	WindowDidBecomeMain(notification foundation.Notification)
	ImplementsWindowDidResignMain() bool
	// optional
	WindowDidResignMain(notification foundation.Notification)
	ImplementsWindowWillReturnFieldEditor_ToObject() bool
	// optional
	WindowWillReturnFieldEditor_ToObject(sender Window, client objc.Object) objc.IObject
	ImplementsWindowDidUpdate() bool
	// optional
	WindowDidUpdate(notification foundation.Notification)
	ImplementsWindowDidExpose() bool
	// optional
	WindowDidExpose(notification foundation.Notification)
	ImplementsWindowDidChangeOcclusionState() bool
	// optional
	WindowDidChangeOcclusionState(notification foundation.Notification)
	ImplementsWindow_ShouldDragDocumentWithEvent_From_WithPasteboard() bool
	// optional
	Window_ShouldDragDocumentWithEvent_From_WithPasteboard(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool
	ImplementsWindowWillReturnUndoManager() bool
	// optional
	WindowWillReturnUndoManager(window Window) foundation.IUndoManager
	ImplementsWindow_ShouldPopUpDocumentPathMenu() bool
	// optional
	Window_ShouldPopUpDocumentPathMenu(window Window, menu Menu) bool
	ImplementsWindow_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize() bool
	// optional
	Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size
	ImplementsWindowWillEnterVersionBrowser() bool
	// optional
	WindowWillEnterVersionBrowser(notification foundation.Notification)
	ImplementsWindowDidEnterVersionBrowser() bool
	// optional
	WindowDidEnterVersionBrowser(notification foundation.Notification)
	ImplementsWindowWillExitVersionBrowser() bool
	// optional
	WindowWillExitVersionBrowser(notification foundation.Notification)
	ImplementsWindowDidExitVersionBrowser() bool
	// optional
	WindowDidExitVersionBrowser(notification foundation.Notification)
}

type WindowDelegateImpl struct {
	_Window_WillPositionSheet_UsingRect                                    func(window Window, sheet Window, rect foundation.Rect) foundation.Rect
	_WindowWillBeginSheet                                                  func(notification foundation.Notification)
	_WindowDidEndSheet                                                     func(notification foundation.Notification)
	_WindowWillResize_ToSize                                               func(sender Window, frameSize foundation.Size) foundation.Size
	_WindowDidResize                                                       func(notification foundation.Notification)
	_WindowWillStartLiveResize                                             func(notification foundation.Notification)
	_WindowDidEndLiveResize                                                func(notification foundation.Notification)
	_WindowWillMiniaturize                                                 func(notification foundation.Notification)
	_WindowDidMiniaturize                                                  func(notification foundation.Notification)
	_WindowDidDeminiaturize                                                func(notification foundation.Notification)
	_WindowWillUseStandardFrame_DefaultFrame                               func(window Window, newFrame foundation.Rect) foundation.Rect
	_WindowShouldZoom_ToFrame                                              func(window Window, newFrame foundation.Rect) bool
	_Window_WillUseFullScreenContentSize                                   func(window Window, proposedSize foundation.Size) foundation.Size
	_Window_WillUseFullScreenPresentationOptions                           func(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	_WindowWillEnterFullScreen                                             func(notification foundation.Notification)
	_WindowDidEnterFullScreen                                              func(notification foundation.Notification)
	_WindowWillExitFullScreen                                              func(notification foundation.Notification)
	_WindowDidExitFullScreen                                               func(notification foundation.Notification)
	_CustomWindowsToEnterFullScreenForWindow                               func(window Window) []IWindow
	_CustomWindowsToEnterFullScreenForWindow_OnScreen                      func(window Window, screen Screen) []IWindow
	_Window_StartCustomAnimationToEnterFullScreenWithDuration              func(window Window, duration foundation.TimeInterval)
	_Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration     func(window Window, screen Screen, duration foundation.TimeInterval)
	_WindowDidFailToEnterFullScreen                                        func(window Window)
	_CustomWindowsToExitFullScreenForWindow                                func(window Window) []IWindow
	_Window_StartCustomAnimationToExitFullScreenWithDuration               func(window Window, duration foundation.TimeInterval)
	_WindowDidFailToExitFullScreen                                         func(window Window)
	_WindowWillMove                                                        func(notification foundation.Notification)
	_WindowDidMove                                                         func(notification foundation.Notification)
	_WindowDidChangeScreen                                                 func(notification foundation.Notification)
	_WindowDidChangeScreenProfile                                          func(notification foundation.Notification)
	_WindowDidChangeBackingProperties                                      func(notification foundation.Notification)
	_WindowShouldClose                                                     func(sender Window) bool
	_WindowWillClose                                                       func(notification foundation.Notification)
	_WindowDidBecomeKey                                                    func(notification foundation.Notification)
	_WindowDidResignKey                                                    func(notification foundation.Notification)
	_WindowDidBecomeMain                                                   func(notification foundation.Notification)
	_WindowDidResignMain                                                   func(notification foundation.Notification)
	_WindowWillReturnFieldEditor_ToObject                                  func(sender Window, client objc.Object) objc.IObject
	_WindowDidUpdate                                                       func(notification foundation.Notification)
	_WindowDidExpose                                                       func(notification foundation.Notification)
	_WindowDidChangeOcclusionState                                         func(notification foundation.Notification)
	_Window_ShouldDragDocumentWithEvent_From_WithPasteboard                func(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool
	_WindowWillReturnUndoManager                                           func(window Window) foundation.IUndoManager
	_Window_ShouldPopUpDocumentPathMenu                                    func(window Window, menu Menu) bool
	_Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize func(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size
	_WindowWillEnterVersionBrowser                                         func(notification foundation.Notification)
	_WindowDidEnterVersionBrowser                                          func(notification foundation.Notification)
	_WindowWillExitVersionBrowser                                          func(notification foundation.Notification)
	_WindowDidExitVersionBrowser                                           func(notification foundation.Notification)
}

func (di *WindowDelegateImpl) ImplementsWindow_WillPositionSheet_UsingRect() bool {
	return di._Window_WillPositionSheet_UsingRect != nil
}

func (di *WindowDelegateImpl) SetWindow_WillPositionSheet_UsingRect(f func(window Window, sheet Window, rect foundation.Rect) foundation.Rect) {
	di._Window_WillPositionSheet_UsingRect = f
}

func (di *WindowDelegateImpl) Window_WillPositionSheet_UsingRect(window Window, sheet Window, rect foundation.Rect) foundation.Rect {
	return di._Window_WillPositionSheet_UsingRect(window, sheet, rect)
}
func (di *WindowDelegateImpl) ImplementsWindowWillBeginSheet() bool {
	return di._WindowWillBeginSheet != nil
}

func (di *WindowDelegateImpl) SetWindowWillBeginSheet(f func(notification foundation.Notification)) {
	di._WindowWillBeginSheet = f
}

func (di *WindowDelegateImpl) WindowWillBeginSheet(notification foundation.Notification) {
	di._WindowWillBeginSheet(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidEndSheet() bool {
	return di._WindowDidEndSheet != nil
}

func (di *WindowDelegateImpl) SetWindowDidEndSheet(f func(notification foundation.Notification)) {
	di._WindowDidEndSheet = f
}

func (di *WindowDelegateImpl) WindowDidEndSheet(notification foundation.Notification) {
	di._WindowDidEndSheet(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillResize_ToSize() bool {
	return di._WindowWillResize_ToSize != nil
}

func (di *WindowDelegateImpl) SetWindowWillResize_ToSize(f func(sender Window, frameSize foundation.Size) foundation.Size) {
	di._WindowWillResize_ToSize = f
}

func (di *WindowDelegateImpl) WindowWillResize_ToSize(sender Window, frameSize foundation.Size) foundation.Size {
	return di._WindowWillResize_ToSize(sender, frameSize)
}
func (di *WindowDelegateImpl) ImplementsWindowDidResize() bool {
	return di._WindowDidResize != nil
}

func (di *WindowDelegateImpl) SetWindowDidResize(f func(notification foundation.Notification)) {
	di._WindowDidResize = f
}

func (di *WindowDelegateImpl) WindowDidResize(notification foundation.Notification) {
	di._WindowDidResize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillStartLiveResize() bool {
	return di._WindowWillStartLiveResize != nil
}

func (di *WindowDelegateImpl) SetWindowWillStartLiveResize(f func(notification foundation.Notification)) {
	di._WindowWillStartLiveResize = f
}

func (di *WindowDelegateImpl) WindowWillStartLiveResize(notification foundation.Notification) {
	di._WindowWillStartLiveResize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidEndLiveResize() bool {
	return di._WindowDidEndLiveResize != nil
}

func (di *WindowDelegateImpl) SetWindowDidEndLiveResize(f func(notification foundation.Notification)) {
	di._WindowDidEndLiveResize = f
}

func (di *WindowDelegateImpl) WindowDidEndLiveResize(notification foundation.Notification) {
	di._WindowDidEndLiveResize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillMiniaturize() bool {
	return di._WindowWillMiniaturize != nil
}

func (di *WindowDelegateImpl) SetWindowWillMiniaturize(f func(notification foundation.Notification)) {
	di._WindowWillMiniaturize = f
}

func (di *WindowDelegateImpl) WindowWillMiniaturize(notification foundation.Notification) {
	di._WindowWillMiniaturize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidMiniaturize() bool {
	return di._WindowDidMiniaturize != nil
}

func (di *WindowDelegateImpl) SetWindowDidMiniaturize(f func(notification foundation.Notification)) {
	di._WindowDidMiniaturize = f
}

func (di *WindowDelegateImpl) WindowDidMiniaturize(notification foundation.Notification) {
	di._WindowDidMiniaturize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidDeminiaturize() bool {
	return di._WindowDidDeminiaturize != nil
}

func (di *WindowDelegateImpl) SetWindowDidDeminiaturize(f func(notification foundation.Notification)) {
	di._WindowDidDeminiaturize = f
}

func (di *WindowDelegateImpl) WindowDidDeminiaturize(notification foundation.Notification) {
	di._WindowDidDeminiaturize(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillUseStandardFrame_DefaultFrame() bool {
	return di._WindowWillUseStandardFrame_DefaultFrame != nil
}

func (di *WindowDelegateImpl) SetWindowWillUseStandardFrame_DefaultFrame(f func(window Window, newFrame foundation.Rect) foundation.Rect) {
	di._WindowWillUseStandardFrame_DefaultFrame = f
}

func (di *WindowDelegateImpl) WindowWillUseStandardFrame_DefaultFrame(window Window, newFrame foundation.Rect) foundation.Rect {
	return di._WindowWillUseStandardFrame_DefaultFrame(window, newFrame)
}
func (di *WindowDelegateImpl) ImplementsWindowShouldZoom_ToFrame() bool {
	return di._WindowShouldZoom_ToFrame != nil
}

func (di *WindowDelegateImpl) SetWindowShouldZoom_ToFrame(f func(window Window, newFrame foundation.Rect) bool) {
	di._WindowShouldZoom_ToFrame = f
}

func (di *WindowDelegateImpl) WindowShouldZoom_ToFrame(window Window, newFrame foundation.Rect) bool {
	return di._WindowShouldZoom_ToFrame(window, newFrame)
}
func (di *WindowDelegateImpl) ImplementsWindow_WillUseFullScreenContentSize() bool {
	return di._Window_WillUseFullScreenContentSize != nil
}

func (di *WindowDelegateImpl) SetWindow_WillUseFullScreenContentSize(f func(window Window, proposedSize foundation.Size) foundation.Size) {
	di._Window_WillUseFullScreenContentSize = f
}

func (di *WindowDelegateImpl) Window_WillUseFullScreenContentSize(window Window, proposedSize foundation.Size) foundation.Size {
	return di._Window_WillUseFullScreenContentSize(window, proposedSize)
}
func (di *WindowDelegateImpl) ImplementsWindow_WillUseFullScreenPresentationOptions() bool {
	return di._Window_WillUseFullScreenPresentationOptions != nil
}

func (di *WindowDelegateImpl) SetWindow_WillUseFullScreenPresentationOptions(f func(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions) {
	di._Window_WillUseFullScreenPresentationOptions = f
}

func (di *WindowDelegateImpl) Window_WillUseFullScreenPresentationOptions(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions {
	return di._Window_WillUseFullScreenPresentationOptions(window, proposedOptions)
}
func (di *WindowDelegateImpl) ImplementsWindowWillEnterFullScreen() bool {
	return di._WindowWillEnterFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowWillEnterFullScreen(f func(notification foundation.Notification)) {
	di._WindowWillEnterFullScreen = f
}

func (di *WindowDelegateImpl) WindowWillEnterFullScreen(notification foundation.Notification) {
	di._WindowWillEnterFullScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidEnterFullScreen() bool {
	return di._WindowDidEnterFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidEnterFullScreen(f func(notification foundation.Notification)) {
	di._WindowDidEnterFullScreen = f
}

func (di *WindowDelegateImpl) WindowDidEnterFullScreen(notification foundation.Notification) {
	di._WindowDidEnterFullScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillExitFullScreen() bool {
	return di._WindowWillExitFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowWillExitFullScreen(f func(notification foundation.Notification)) {
	di._WindowWillExitFullScreen = f
}

func (di *WindowDelegateImpl) WindowWillExitFullScreen(notification foundation.Notification) {
	di._WindowWillExitFullScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidExitFullScreen() bool {
	return di._WindowDidExitFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidExitFullScreen(f func(notification foundation.Notification)) {
	di._WindowDidExitFullScreen = f
}

func (di *WindowDelegateImpl) WindowDidExitFullScreen(notification foundation.Notification) {
	di._WindowDidExitFullScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsCustomWindowsToEnterFullScreenForWindow() bool {
	return di._CustomWindowsToEnterFullScreenForWindow != nil
}

func (di *WindowDelegateImpl) SetCustomWindowsToEnterFullScreenForWindow(f func(window Window) []IWindow) {
	di._CustomWindowsToEnterFullScreenForWindow = f
}

func (di *WindowDelegateImpl) CustomWindowsToEnterFullScreenForWindow(window Window) []IWindow {
	return di._CustomWindowsToEnterFullScreenForWindow(window)
}
func (di *WindowDelegateImpl) ImplementsCustomWindowsToEnterFullScreenForWindow_OnScreen() bool {
	return di._CustomWindowsToEnterFullScreenForWindow_OnScreen != nil
}

func (di *WindowDelegateImpl) SetCustomWindowsToEnterFullScreenForWindow_OnScreen(f func(window Window, screen Screen) []IWindow) {
	di._CustomWindowsToEnterFullScreenForWindow_OnScreen = f
}

func (di *WindowDelegateImpl) CustomWindowsToEnterFullScreenForWindow_OnScreen(window Window, screen Screen) []IWindow {
	return di._CustomWindowsToEnterFullScreenForWindow_OnScreen(window, screen)
}
func (di *WindowDelegateImpl) ImplementsWindow_StartCustomAnimationToEnterFullScreenWithDuration() bool {
	return di._Window_StartCustomAnimationToEnterFullScreenWithDuration != nil
}

func (di *WindowDelegateImpl) SetWindow_StartCustomAnimationToEnterFullScreenWithDuration(f func(window Window, duration foundation.TimeInterval)) {
	di._Window_StartCustomAnimationToEnterFullScreenWithDuration = f
}

func (di *WindowDelegateImpl) Window_StartCustomAnimationToEnterFullScreenWithDuration(window Window, duration foundation.TimeInterval) {
	di._Window_StartCustomAnimationToEnterFullScreenWithDuration(window, duration)
}
func (di *WindowDelegateImpl) ImplementsWindow_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration() bool {
	return di._Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration != nil
}

func (di *WindowDelegateImpl) SetWindow_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(f func(window Window, screen Screen, duration foundation.TimeInterval)) {
	di._Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration = f
}

func (di *WindowDelegateImpl) Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(window Window, screen Screen, duration foundation.TimeInterval) {
	di._Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(window, screen, duration)
}
func (di *WindowDelegateImpl) ImplementsWindowDidFailToEnterFullScreen() bool {
	return di._WindowDidFailToEnterFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidFailToEnterFullScreen(f func(window Window)) {
	di._WindowDidFailToEnterFullScreen = f
}

func (di *WindowDelegateImpl) WindowDidFailToEnterFullScreen(window Window) {
	di._WindowDidFailToEnterFullScreen(window)
}
func (di *WindowDelegateImpl) ImplementsCustomWindowsToExitFullScreenForWindow() bool {
	return di._CustomWindowsToExitFullScreenForWindow != nil
}

func (di *WindowDelegateImpl) SetCustomWindowsToExitFullScreenForWindow(f func(window Window) []IWindow) {
	di._CustomWindowsToExitFullScreenForWindow = f
}

func (di *WindowDelegateImpl) CustomWindowsToExitFullScreenForWindow(window Window) []IWindow {
	return di._CustomWindowsToExitFullScreenForWindow(window)
}
func (di *WindowDelegateImpl) ImplementsWindow_StartCustomAnimationToExitFullScreenWithDuration() bool {
	return di._Window_StartCustomAnimationToExitFullScreenWithDuration != nil
}

func (di *WindowDelegateImpl) SetWindow_StartCustomAnimationToExitFullScreenWithDuration(f func(window Window, duration foundation.TimeInterval)) {
	di._Window_StartCustomAnimationToExitFullScreenWithDuration = f
}

func (di *WindowDelegateImpl) Window_StartCustomAnimationToExitFullScreenWithDuration(window Window, duration foundation.TimeInterval) {
	di._Window_StartCustomAnimationToExitFullScreenWithDuration(window, duration)
}
func (di *WindowDelegateImpl) ImplementsWindowDidFailToExitFullScreen() bool {
	return di._WindowDidFailToExitFullScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidFailToExitFullScreen(f func(window Window)) {
	di._WindowDidFailToExitFullScreen = f
}

func (di *WindowDelegateImpl) WindowDidFailToExitFullScreen(window Window) {
	di._WindowDidFailToExitFullScreen(window)
}
func (di *WindowDelegateImpl) ImplementsWindowWillMove() bool {
	return di._WindowWillMove != nil
}

func (di *WindowDelegateImpl) SetWindowWillMove(f func(notification foundation.Notification)) {
	di._WindowWillMove = f
}

func (di *WindowDelegateImpl) WindowWillMove(notification foundation.Notification) {
	di._WindowWillMove(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidMove() bool {
	return di._WindowDidMove != nil
}

func (di *WindowDelegateImpl) SetWindowDidMove(f func(notification foundation.Notification)) {
	di._WindowDidMove = f
}

func (di *WindowDelegateImpl) WindowDidMove(notification foundation.Notification) {
	di._WindowDidMove(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidChangeScreen() bool {
	return di._WindowDidChangeScreen != nil
}

func (di *WindowDelegateImpl) SetWindowDidChangeScreen(f func(notification foundation.Notification)) {
	di._WindowDidChangeScreen = f
}

func (di *WindowDelegateImpl) WindowDidChangeScreen(notification foundation.Notification) {
	di._WindowDidChangeScreen(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidChangeScreenProfile() bool {
	return di._WindowDidChangeScreenProfile != nil
}

func (di *WindowDelegateImpl) SetWindowDidChangeScreenProfile(f func(notification foundation.Notification)) {
	di._WindowDidChangeScreenProfile = f
}

func (di *WindowDelegateImpl) WindowDidChangeScreenProfile(notification foundation.Notification) {
	di._WindowDidChangeScreenProfile(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidChangeBackingProperties() bool {
	return di._WindowDidChangeBackingProperties != nil
}

func (di *WindowDelegateImpl) SetWindowDidChangeBackingProperties(f func(notification foundation.Notification)) {
	di._WindowDidChangeBackingProperties = f
}

func (di *WindowDelegateImpl) WindowDidChangeBackingProperties(notification foundation.Notification) {
	di._WindowDidChangeBackingProperties(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowShouldClose() bool {
	return di._WindowShouldClose != nil
}

func (di *WindowDelegateImpl) SetWindowShouldClose(f func(sender Window) bool) {
	di._WindowShouldClose = f
}

func (di *WindowDelegateImpl) WindowShouldClose(sender Window) bool {
	return di._WindowShouldClose(sender)
}
func (di *WindowDelegateImpl) ImplementsWindowWillClose() bool {
	return di._WindowWillClose != nil
}

func (di *WindowDelegateImpl) SetWindowWillClose(f func(notification foundation.Notification)) {
	di._WindowWillClose = f
}

func (di *WindowDelegateImpl) WindowWillClose(notification foundation.Notification) {
	di._WindowWillClose(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidBecomeKey() bool {
	return di._WindowDidBecomeKey != nil
}

func (di *WindowDelegateImpl) SetWindowDidBecomeKey(f func(notification foundation.Notification)) {
	di._WindowDidBecomeKey = f
}

func (di *WindowDelegateImpl) WindowDidBecomeKey(notification foundation.Notification) {
	di._WindowDidBecomeKey(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidResignKey() bool {
	return di._WindowDidResignKey != nil
}

func (di *WindowDelegateImpl) SetWindowDidResignKey(f func(notification foundation.Notification)) {
	di._WindowDidResignKey = f
}

func (di *WindowDelegateImpl) WindowDidResignKey(notification foundation.Notification) {
	di._WindowDidResignKey(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidBecomeMain() bool {
	return di._WindowDidBecomeMain != nil
}

func (di *WindowDelegateImpl) SetWindowDidBecomeMain(f func(notification foundation.Notification)) {
	di._WindowDidBecomeMain = f
}

func (di *WindowDelegateImpl) WindowDidBecomeMain(notification foundation.Notification) {
	di._WindowDidBecomeMain(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidResignMain() bool {
	return di._WindowDidResignMain != nil
}

func (di *WindowDelegateImpl) SetWindowDidResignMain(f func(notification foundation.Notification)) {
	di._WindowDidResignMain = f
}

func (di *WindowDelegateImpl) WindowDidResignMain(notification foundation.Notification) {
	di._WindowDidResignMain(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillReturnFieldEditor_ToObject() bool {
	return di._WindowWillReturnFieldEditor_ToObject != nil
}

func (di *WindowDelegateImpl) SetWindowWillReturnFieldEditor_ToObject(f func(sender Window, client objc.Object) objc.IObject) {
	di._WindowWillReturnFieldEditor_ToObject = f
}

func (di *WindowDelegateImpl) WindowWillReturnFieldEditor_ToObject(sender Window, client objc.Object) objc.IObject {
	return di._WindowWillReturnFieldEditor_ToObject(sender, client)
}
func (di *WindowDelegateImpl) ImplementsWindowDidUpdate() bool {
	return di._WindowDidUpdate != nil
}

func (di *WindowDelegateImpl) SetWindowDidUpdate(f func(notification foundation.Notification)) {
	di._WindowDidUpdate = f
}

func (di *WindowDelegateImpl) WindowDidUpdate(notification foundation.Notification) {
	di._WindowDidUpdate(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidExpose() bool {
	return di._WindowDidExpose != nil
}

func (di *WindowDelegateImpl) SetWindowDidExpose(f func(notification foundation.Notification)) {
	di._WindowDidExpose = f
}

func (di *WindowDelegateImpl) WindowDidExpose(notification foundation.Notification) {
	di._WindowDidExpose(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidChangeOcclusionState() bool {
	return di._WindowDidChangeOcclusionState != nil
}

func (di *WindowDelegateImpl) SetWindowDidChangeOcclusionState(f func(notification foundation.Notification)) {
	di._WindowDidChangeOcclusionState = f
}

func (di *WindowDelegateImpl) WindowDidChangeOcclusionState(notification foundation.Notification) {
	di._WindowDidChangeOcclusionState(notification)
}
func (di *WindowDelegateImpl) ImplementsWindow_ShouldDragDocumentWithEvent_From_WithPasteboard() bool {
	return di._Window_ShouldDragDocumentWithEvent_From_WithPasteboard != nil
}

func (di *WindowDelegateImpl) SetWindow_ShouldDragDocumentWithEvent_From_WithPasteboard(f func(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool) {
	di._Window_ShouldDragDocumentWithEvent_From_WithPasteboard = f
}

func (di *WindowDelegateImpl) Window_ShouldDragDocumentWithEvent_From_WithPasteboard(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool {
	return di._Window_ShouldDragDocumentWithEvent_From_WithPasteboard(window, event, dragImageLocation, pasteboard)
}
func (di *WindowDelegateImpl) ImplementsWindowWillReturnUndoManager() bool {
	return di._WindowWillReturnUndoManager != nil
}

func (di *WindowDelegateImpl) SetWindowWillReturnUndoManager(f func(window Window) foundation.IUndoManager) {
	di._WindowWillReturnUndoManager = f
}

func (di *WindowDelegateImpl) WindowWillReturnUndoManager(window Window) foundation.IUndoManager {
	return di._WindowWillReturnUndoManager(window)
}
func (di *WindowDelegateImpl) ImplementsWindow_ShouldPopUpDocumentPathMenu() bool {
	return di._Window_ShouldPopUpDocumentPathMenu != nil
}

func (di *WindowDelegateImpl) SetWindow_ShouldPopUpDocumentPathMenu(f func(window Window, menu Menu) bool) {
	di._Window_ShouldPopUpDocumentPathMenu = f
}

func (di *WindowDelegateImpl) Window_ShouldPopUpDocumentPathMenu(window Window, menu Menu) bool {
	return di._Window_ShouldPopUpDocumentPathMenu(window, menu)
}
func (di *WindowDelegateImpl) ImplementsWindow_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize() bool {
	return di._Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize != nil
}

func (di *WindowDelegateImpl) SetWindow_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(f func(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size) {
	di._Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize = f
}

func (di *WindowDelegateImpl) Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size {
	return di._Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(window, maxPreferredFrameSize, maxAllowedFrameSize)
}
func (di *WindowDelegateImpl) ImplementsWindowWillEnterVersionBrowser() bool {
	return di._WindowWillEnterVersionBrowser != nil
}

func (di *WindowDelegateImpl) SetWindowWillEnterVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowWillEnterVersionBrowser = f
}

func (di *WindowDelegateImpl) WindowWillEnterVersionBrowser(notification foundation.Notification) {
	di._WindowWillEnterVersionBrowser(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidEnterVersionBrowser() bool {
	return di._WindowDidEnterVersionBrowser != nil
}

func (di *WindowDelegateImpl) SetWindowDidEnterVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowDidEnterVersionBrowser = f
}

func (di *WindowDelegateImpl) WindowDidEnterVersionBrowser(notification foundation.Notification) {
	di._WindowDidEnterVersionBrowser(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowWillExitVersionBrowser() bool {
	return di._WindowWillExitVersionBrowser != nil
}

func (di *WindowDelegateImpl) SetWindowWillExitVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowWillExitVersionBrowser = f
}

func (di *WindowDelegateImpl) WindowWillExitVersionBrowser(notification foundation.Notification) {
	di._WindowWillExitVersionBrowser(notification)
}
func (di *WindowDelegateImpl) ImplementsWindowDidExitVersionBrowser() bool {
	return di._WindowDidExitVersionBrowser != nil
}

func (di *WindowDelegateImpl) SetWindowDidExitVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowDidExitVersionBrowser = f
}

func (di *WindowDelegateImpl) WindowDidExitVersionBrowser(notification foundation.Notification) {
	di._WindowDidExitVersionBrowser(notification)
}

type WindowDelegateWrapper struct {
	objc.Object
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_WillPositionSheet_UsingRect() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willPositionSheet:usingRect:"))
}

func (w_ WindowDelegateWrapper) Window_WillPositionSheet_UsingRect(window IWindow, sheet IWindow, rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "window:willPositionSheet:usingRect:", window, sheet, rect)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillBeginSheet() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillBeginSheet:"))
}

func (w_ WindowDelegateWrapper) WindowWillBeginSheet(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillBeginSheet:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidEndSheet() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEndSheet:"))
}

func (w_ WindowDelegateWrapper) WindowDidEndSheet(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidEndSheet:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillResize_ToSize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillResize:toSize:"))
}

func (w_ WindowDelegateWrapper) WindowWillResize_ToSize(sender IWindow, frameSize foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "windowWillResize:toSize:", sender, frameSize)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidResize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidResize:"))
}

func (w_ WindowDelegateWrapper) WindowDidResize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidResize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillStartLiveResize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillStartLiveResize:"))
}

func (w_ WindowDelegateWrapper) WindowWillStartLiveResize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillStartLiveResize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidEndLiveResize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEndLiveResize:"))
}

func (w_ WindowDelegateWrapper) WindowDidEndLiveResize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidEndLiveResize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillMiniaturize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillMiniaturize:"))
}

func (w_ WindowDelegateWrapper) WindowWillMiniaturize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillMiniaturize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidMiniaturize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidMiniaturize:"))
}

func (w_ WindowDelegateWrapper) WindowDidMiniaturize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidMiniaturize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidDeminiaturize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidDeminiaturize:"))
}

func (w_ WindowDelegateWrapper) WindowDidDeminiaturize(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidDeminiaturize:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillUseStandardFrame_DefaultFrame() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillUseStandardFrame:defaultFrame:"))
}

func (w_ WindowDelegateWrapper) WindowWillUseStandardFrame_DefaultFrame(window IWindow, newFrame foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](w_, "windowWillUseStandardFrame:defaultFrame:", window, newFrame)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowShouldZoom_ToFrame() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowShouldZoom:toFrame:"))
}

func (w_ WindowDelegateWrapper) WindowShouldZoom_ToFrame(window IWindow, newFrame foundation.Rect) bool {
	rv := ffi.CallMethod[bool](w_, "windowShouldZoom:toFrame:", window, newFrame)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_WillUseFullScreenContentSize() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willUseFullScreenContentSize:"))
}

func (w_ WindowDelegateWrapper) Window_WillUseFullScreenContentSize(window IWindow, proposedSize foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "window:willUseFullScreenContentSize:", window, proposedSize)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_WillUseFullScreenPresentationOptions() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willUseFullScreenPresentationOptions:"))
}

func (w_ WindowDelegateWrapper) Window_WillUseFullScreenPresentationOptions(window IWindow, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions {
	rv := ffi.CallMethod[ApplicationPresentationOptions](w_, "window:willUseFullScreenPresentationOptions:", window, proposedOptions)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillEnterFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowWillEnterFullScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillEnterFullScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEnterFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidEnterFullScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidEnterFullScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillExitFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillExitFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowWillExitFullScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillExitFullScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidExitFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidExitFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidExitFullScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidExitFullScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsCustomWindowsToEnterFullScreenForWindow() bool {
	return w_.RespondsToSelector(objc.GetSelector("customWindowsToEnterFullScreenForWindow:"))
}

func (w_ WindowDelegateWrapper) CustomWindowsToEnterFullScreenForWindow(window IWindow) []Window {
	rv := ffi.CallMethod[[]Window](w_, "customWindowsToEnterFullScreenForWindow:", window)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsCustomWindowsToEnterFullScreenForWindow_OnScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("customWindowsToEnterFullScreenForWindow:onScreen:"))
}

func (w_ WindowDelegateWrapper) CustomWindowsToEnterFullScreenForWindow_OnScreen(window IWindow, screen IScreen) []Window {
	rv := ffi.CallMethod[[]Window](w_, "customWindowsToEnterFullScreenForWindow:onScreen:", window, screen)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_StartCustomAnimationToEnterFullScreenWithDuration() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:startCustomAnimationToEnterFullScreenWithDuration:"))
}

func (w_ WindowDelegateWrapper) Window_StartCustomAnimationToEnterFullScreenWithDuration(window IWindow, duration foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](w_, "window:startCustomAnimationToEnterFullScreenWithDuration:", window, duration)
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:"))
}

func (w_ WindowDelegateWrapper) Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(window IWindow, screen IScreen, duration foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](w_, "window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:", window, screen, duration)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidFailToEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidFailToEnterFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidFailToEnterFullScreen(window IWindow) {
	ffi.CallMethod[ffi.Void](w_, "windowDidFailToEnterFullScreen:", window)
}

func (w_ *WindowDelegateWrapper) ImplementsCustomWindowsToExitFullScreenForWindow() bool {
	return w_.RespondsToSelector(objc.GetSelector("customWindowsToExitFullScreenForWindow:"))
}

func (w_ WindowDelegateWrapper) CustomWindowsToExitFullScreenForWindow(window IWindow) []Window {
	rv := ffi.CallMethod[[]Window](w_, "customWindowsToExitFullScreenForWindow:", window)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_StartCustomAnimationToExitFullScreenWithDuration() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:startCustomAnimationToExitFullScreenWithDuration:"))
}

func (w_ WindowDelegateWrapper) Window_StartCustomAnimationToExitFullScreenWithDuration(window IWindow, duration foundation.TimeInterval) {
	ffi.CallMethod[ffi.Void](w_, "window:startCustomAnimationToExitFullScreenWithDuration:", window, duration)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidFailToExitFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidFailToExitFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidFailToExitFullScreen(window IWindow) {
	ffi.CallMethod[ffi.Void](w_, "windowDidFailToExitFullScreen:", window)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillMove() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillMove:"))
}

func (w_ WindowDelegateWrapper) WindowWillMove(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillMove:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidMove() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidMove:"))
}

func (w_ WindowDelegateWrapper) WindowDidMove(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidMove:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidChangeScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeScreen(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidChangeScreen:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidChangeScreenProfile() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeScreenProfile:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeScreenProfile(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidChangeScreenProfile:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidChangeBackingProperties() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeBackingProperties:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeBackingProperties(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidChangeBackingProperties:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowShouldClose() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowShouldClose:"))
}

func (w_ WindowDelegateWrapper) WindowShouldClose(sender IWindow) bool {
	rv := ffi.CallMethod[bool](w_, "windowShouldClose:", sender)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillClose() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillClose:"))
}

func (w_ WindowDelegateWrapper) WindowWillClose(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillClose:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidBecomeKey() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidBecomeKey:"))
}

func (w_ WindowDelegateWrapper) WindowDidBecomeKey(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidBecomeKey:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidResignKey() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidResignKey:"))
}

func (w_ WindowDelegateWrapper) WindowDidResignKey(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidResignKey:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidBecomeMain() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidBecomeMain:"))
}

func (w_ WindowDelegateWrapper) WindowDidBecomeMain(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidBecomeMain:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidResignMain() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidResignMain:"))
}

func (w_ WindowDelegateWrapper) WindowDidResignMain(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidResignMain:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillReturnFieldEditor_ToObject() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillReturnFieldEditor:toObject:"))
}

func (w_ WindowDelegateWrapper) WindowWillReturnFieldEditor_ToObject(sender IWindow, client objc.IObject) objc.Object {
	rv := ffi.CallMethod[objc.Object](w_, "windowWillReturnFieldEditor:toObject:", sender, client)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidUpdate() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidUpdate:"))
}

func (w_ WindowDelegateWrapper) WindowDidUpdate(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidUpdate:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidExpose() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidExpose:"))
}

func (w_ WindowDelegateWrapper) WindowDidExpose(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidExpose:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidChangeOcclusionState() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeOcclusionState:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeOcclusionState(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidChangeOcclusionState:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_ShouldDragDocumentWithEvent_From_WithPasteboard() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:shouldDragDocumentWithEvent:from:withPasteboard:"))
}

func (w_ WindowDelegateWrapper) Window_ShouldDragDocumentWithEvent_From_WithPasteboard(window IWindow, event IEvent, dragImageLocation foundation.Point, pasteboard IPasteboard) bool {
	rv := ffi.CallMethod[bool](w_, "window:shouldDragDocumentWithEvent:from:withPasteboard:", window, event, dragImageLocation, pasteboard)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillReturnUndoManager() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillReturnUndoManager:"))
}

func (w_ WindowDelegateWrapper) WindowWillReturnUndoManager(window IWindow) foundation.UndoManager {
	rv := ffi.CallMethod[foundation.UndoManager](w_, "windowWillReturnUndoManager:", window)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_ShouldPopUpDocumentPathMenu() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:shouldPopUpDocumentPathMenu:"))
}

func (w_ WindowDelegateWrapper) Window_ShouldPopUpDocumentPathMenu(window IWindow, menu IMenu) bool {
	rv := ffi.CallMethod[bool](w_, "window:shouldPopUpDocumentPathMenu:", window, menu)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindow_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:"))
}

func (w_ WindowDelegateWrapper) Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(window IWindow, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](w_, "window:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:", window, maxPreferredFrameSize, maxAllowedFrameSize)
	return rv
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillEnterVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillEnterVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowWillEnterVersionBrowser(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillEnterVersionBrowser:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidEnterVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEnterVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowDidEnterVersionBrowser(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidEnterVersionBrowser:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowWillExitVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillExitVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowWillExitVersionBrowser(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowWillExitVersionBrowser:", notification)
}

func (w_ *WindowDelegateWrapper) ImplementsWindowDidExitVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidExitVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowDidExitVersionBrowser(notification foundation.INotification) {
	ffi.CallMethod[ffi.Void](w_, "windowDidExitVersionBrowser:", notification)
}
