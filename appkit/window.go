// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	SetFrameFromString(string_ WindowPersistableFrameDescriptor)
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
	SetDelegate0(value objc.IObject)
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
	Drawers() []Drawer
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
	rv := objc.CallMethod[Window](wc, "windowWithContentViewController:", contentViewController)
	return rv
}

func (w_ Window) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	rv := objc.CallMethod[Window](w_, "initWithContentRect:styleMask:backing:defer:", contentRect, style, backingStoreType, flag)
	return rv
}

func (w_ Window) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Window {
	rv := objc.CallMethod[Window](w_, "initWithContentRect:styleMask:backing:defer:screen:", contentRect, style, backingStoreType, flag, screen)
	return rv
}

func (w_ Window) Init() Window {
	rv := objc.CallMethod[Window](w_, "init")
	return rv
}

func (wc _WindowClass) Alloc() Window {
	rv := objc.CallMethod[Window](wc, "alloc")
	return rv
}

func (wc _WindowClass) New() Window {
	rv := objc.CallMethod[Window](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWindow() Window {
	return WindowClass.New()
}

func (w_ Window) ToggleFullScreen(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "toggleFullScreen:", sender)
}

func (w_ Window) SetDynamicDepthLimit(flag bool) {
	objc.CallMethod[objc.Void](w_, "setDynamicDepthLimit:", flag)
}

func (w_ Window) InvalidateShadow() {
	objc.CallMethod[objc.Void](w_, "invalidateShadow")
}

func (w_ Window) AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool {
	rv := objc.CallMethod[bool](w_, "autorecalculatesContentBorderThicknessForEdge:", edge)
	return rv
}

func (w_ Window) SetAutorecalculatesContentBorderThickness_ForEdge(flag bool, edge foundation.RectEdge) {
	objc.CallMethod[objc.Void](w_, "setAutorecalculatesContentBorderThickness:forEdge:", flag, edge)
}

func (w_ Window) ContentBorderThicknessForEdge(edge foundation.RectEdge) float64 {
	rv := objc.CallMethod[float64](w_, "contentBorderThicknessForEdge:", edge)
	return rv
}

func (w_ Window) SetContentBorderThickness_ForEdge(thickness float64, edge foundation.RectEdge) {
	objc.CallMethod[objc.Void](w_, "setContentBorderThickness:forEdge:", thickness, edge)
}

func (wc _WindowClass) WindowNumbersWithOptions(options WindowNumberListOptions) []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](wc, "windowNumbersWithOptions:", options)
	return rv
}

func (wc _WindowClass) ContentRectForFrameRect_StyleMask(fRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](wc, "contentRectForFrameRect:styleMask:", fRect, style)
	return rv
}

func (wc _WindowClass) FrameRectForContentRect_StyleMask(cRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](wc, "frameRectForContentRect:styleMask:", cRect, style)
	return rv
}

func (wc _WindowClass) MinFrameWidthWithTitle_StyleMask(title string, style WindowStyleMask) float64 {
	rv := objc.CallMethod[float64](wc, "minFrameWidthWithTitle:styleMask:", title, style)
	return rv
}

func (w_ Window) ContentRectForFrameRect(frameRect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "contentRectForFrameRect:", frameRect)
	return rv
}

func (w_ Window) FrameRectForContentRect(contentRect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "frameRectForContentRect:", contentRect)
	return rv
}

func (w_ Window) BeginSheet_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.CallMethod[objc.Void](w_, "beginSheet:completionHandler:", sheetWindow, handler)
}

func (w_ Window) BeginCriticalSheet_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.CallMethod[objc.Void](w_, "beginCriticalSheet:completionHandler:", sheetWindow, handler)
}

func (w_ Window) EndSheet(sheetWindow IWindow) {
	objc.CallMethod[objc.Void](w_, "endSheet:", sheetWindow)
}

func (w_ Window) EndSheet_ReturnCode(sheetWindow IWindow, returnCode ModalResponse) {
	objc.CallMethod[objc.Void](w_, "endSheet:returnCode:", sheetWindow, returnCode)
}

func (w_ Window) SetFrameOrigin(point foundation.Point) {
	objc.CallMethod[objc.Void](w_, "setFrameOrigin:", point)
}

func (w_ Window) SetFrameTopLeftPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](w_, "setFrameTopLeftPoint:", point)
}

func (w_ Window) ConstrainFrameRect_ToScreen(frameRect foundation.Rect, screen IScreen) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "constrainFrameRect:toScreen:", frameRect, screen)
	return rv
}

func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, "cascadeTopLeftFromPoint:", topLeftPoint)
	return rv
}

func (w_ Window) SetFrame_Display(frameRect foundation.Rect, flag bool) {
	objc.CallMethod[objc.Void](w_, "setFrame:display:", frameRect, flag)
}

func (w_ Window) SetFrame_Display_Animate(frameRect foundation.Rect, displayFlag bool, animateFlag bool) {
	objc.CallMethod[objc.Void](w_, "setFrame:display:animate:", frameRect, displayFlag, animateFlag)
}

func (w_ Window) AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](w_, "animationResizeTime:", newFrame)
	return rv
}

func (w_ Window) PerformZoom(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "performZoom:", sender)
}

func (w_ Window) Zoom(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "zoom:", sender)
}

func (w_ Window) SetContentSize(size foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setContentSize:", size)
}

func (w_ Window) OrderOut(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "orderOut:", sender)
}

func (w_ Window) OrderBack(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "orderBack:", sender)
}

func (w_ Window) OrderFront(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "orderFront:", sender)
}

func (w_ Window) OrderFrontRegardless() {
	objc.CallMethod[objc.Void](w_, "orderFrontRegardless")
}

func (w_ Window) OrderWindow_RelativeTo(place WindowOrderingMode, otherWin int) {
	objc.CallMethod[objc.Void](w_, "orderWindow:relativeTo:", place, otherWin)
}

func (wc _WindowClass) RemoveFrameUsingName(name WindowFrameAutosaveName) {
	objc.CallMethod[objc.Void](wc, "removeFrameUsingName:", name)
}

func (w_ Window) SetFrameUsingName(name WindowFrameAutosaveName) bool {
	rv := objc.CallMethod[bool](w_, "setFrameUsingName:", name)
	return rv
}

func (w_ Window) SetFrameUsingName_Force(name WindowFrameAutosaveName, force bool) bool {
	rv := objc.CallMethod[bool](w_, "setFrameUsingName:force:", name, force)
	return rv
}

func (w_ Window) SaveFrameUsingName(name WindowFrameAutosaveName) {
	objc.CallMethod[objc.Void](w_, "saveFrameUsingName:", name)
}

func (w_ Window) SetFrameAutosaveName(name WindowFrameAutosaveName) bool {
	rv := objc.CallMethod[bool](w_, "setFrameAutosaveName:", name)
	return rv
}

func (w_ Window) SetFrameFromString(string_ WindowPersistableFrameDescriptor) {
	objc.CallMethod[objc.Void](w_, "setFrameFromString:", string_)
}

func (w_ Window) MakeKeyWindow() {
	objc.CallMethod[objc.Void](w_, "makeKeyWindow")
}

func (w_ Window) MakeKeyAndOrderFront(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "makeKeyAndOrderFront:", sender)
}

func (w_ Window) BecomeKeyWindow() {
	objc.CallMethod[objc.Void](w_, "becomeKeyWindow")
}

func (w_ Window) ResignKeyWindow() {
	objc.CallMethod[objc.Void](w_, "resignKeyWindow")
}

func (w_ Window) MakeMainWindow() {
	objc.CallMethod[objc.Void](w_, "makeMainWindow")
}

func (w_ Window) BecomeMainWindow() {
	objc.CallMethod[objc.Void](w_, "becomeMainWindow")
}

func (w_ Window) ResignMainWindow() {
	objc.CallMethod[objc.Void](w_, "resignMainWindow")
}

func (w_ Window) ToggleToolbarShown(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "toggleToolbarShown:", sender)
}

func (w_ Window) RunToolbarCustomizationPalette(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "runToolbarCustomizationPalette:", sender)
}

func (w_ Window) AddChildWindow_Ordered(childWin IWindow, place WindowOrderingMode) {
	objc.CallMethod[objc.Void](w_, "addChildWindow:ordered:", childWin, place)
}

func (w_ Window) RemoveChildWindow(childWin IWindow) {
	objc.CallMethod[objc.Void](w_, "removeChildWindow:", childWin)
}

func (w_ Window) EnableKeyEquivalentForDefaultButtonCell() {
	objc.CallMethod[objc.Void](w_, "enableKeyEquivalentForDefaultButtonCell")
}

func (w_ Window) DisableKeyEquivalentForDefaultButtonCell() {
	objc.CallMethod[objc.Void](w_, "disableKeyEquivalentForDefaultButtonCell")
}

func (w_ Window) FieldEditor_ForObject(createFlag bool, object objc.IObject) Text {
	rv := objc.CallMethod[Text](w_, "fieldEditor:forObject:", createFlag, object)
	return rv
}

func (w_ Window) EndEditingFor(object objc.IObject) {
	objc.CallMethod[objc.Void](w_, "endEditingFor:", object)
}

func (w_ Window) EnableCursorRects() {
	objc.CallMethod[objc.Void](w_, "enableCursorRects")
}

func (w_ Window) DisableCursorRects() {
	objc.CallMethod[objc.Void](w_, "disableCursorRects")
}

func (w_ Window) DiscardCursorRects() {
	objc.CallMethod[objc.Void](w_, "discardCursorRects")
}

func (w_ Window) InvalidateCursorRectsForView(view IView) {
	objc.CallMethod[objc.Void](w_, "invalidateCursorRectsForView:", view)
}

func (w_ Window) ResetCursorRects() {
	objc.CallMethod[objc.Void](w_, "resetCursorRects")
}

func (wc _WindowClass) StandardWindowButton_ForStyleMask(b WindowButton, styleMask WindowStyleMask) Button {
	rv := objc.CallMethod[Button](wc, "standardWindowButton:forStyleMask:", b, styleMask)
	return rv
}

func (w_ Window) StandardWindowButton(b WindowButton) Button {
	rv := objc.CallMethod[Button](w_, "standardWindowButton:", b)
	return rv
}

func (w_ Window) AddTitlebarAccessoryViewController(childViewController ITitlebarAccessoryViewController) {
	objc.CallMethod[objc.Void](w_, "addTitlebarAccessoryViewController:", childViewController)
}

func (w_ Window) InsertTitlebarAccessoryViewController_AtIndex(childViewController ITitlebarAccessoryViewController, index int) {
	objc.CallMethod[objc.Void](w_, "insertTitlebarAccessoryViewController:atIndex:", childViewController, index)
}

func (w_ Window) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	objc.CallMethod[objc.Void](w_, "removeTitlebarAccessoryViewControllerAtIndex:", index)
}

func (w_ Window) AddTabbedWindow_Ordered(window IWindow, ordered WindowOrderingMode) {
	objc.CallMethod[objc.Void](w_, "addTabbedWindow:ordered:", window, ordered)
}

func (w_ Window) MergeAllWindows(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "mergeAllWindows:", sender)
}

func (w_ Window) SelectNextTab(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "selectNextTab:", sender)
}

func (w_ Window) SelectPreviousTab(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "selectPreviousTab:", sender)
}

func (w_ Window) MoveTabToNewWindow(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "moveTabToNewWindow:", sender)
}

func (w_ Window) ToggleTabBar(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "toggleTabBar:", sender)
}

func (w_ Window) ToggleTabOverview(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "toggleTabOverview:", sender)
}

func (w_ Window) NextEventMatchingMask(mask EventMask) Event {
	rv := objc.CallMethod[Event](w_, "nextEventMatchingMask:", mask)
	return rv
}

func (w_ Window) NextEventMatchingMask_UntilDate_InMode_Dequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event {
	rv := objc.CallMethod[Event](w_, "nextEventMatchingMask:untilDate:inMode:dequeue:", mask, expiration, mode, deqFlag)
	return rv
}

func (w_ Window) DiscardEventsMatchingMask_BeforeEvent(mask EventMask, lastEvent IEvent) {
	objc.CallMethod[objc.Void](w_, "discardEventsMatchingMask:beforeEvent:", mask, lastEvent)
}

func (w_ Window) PostEvent_AtStart(event IEvent, flag bool) {
	objc.CallMethod[objc.Void](w_, "postEvent:atStart:", event, flag)
}

func (w_ Window) SendEvent(event IEvent) {
	objc.CallMethod[objc.Void](w_, "sendEvent:", event)
}

func (w_ Window) MakeFirstResponder(responder IResponder) bool {
	rv := objc.CallMethod[bool](w_, "makeFirstResponder:", responder)
	return rv
}

func (w_ Window) SelectKeyViewPrecedingView(view IView) {
	objc.CallMethod[objc.Void](w_, "selectKeyViewPrecedingView:", view)
}

func (w_ Window) SelectKeyViewFollowingView(view IView) {
	objc.CallMethod[objc.Void](w_, "selectKeyViewFollowingView:", view)
}

func (w_ Window) SelectPreviousKeyView(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "selectPreviousKeyView:", sender)
}

func (w_ Window) SelectNextKeyView(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "selectNextKeyView:", sender)
}

func (w_ Window) RecalculateKeyViewLoop() {
	objc.CallMethod[objc.Void](w_, "recalculateKeyViewLoop")
}

func (wc _WindowClass) WindowNumberAtPoint_BelowWindowWithWindowNumber(point foundation.Point, windowNumber int) int {
	rv := objc.CallMethod[int](wc, "windowNumberAtPoint:belowWindowWithWindowNumber:", point, windowNumber)
	return rv
}

func (w_ Window) TrackEventsMatchingMask_Timeout_Mode_Handler(mask EventMask, timeout foundation.TimeInterval, mode foundation.RunLoopMode, trackingHandler func(event Event, stop *bool)) {
	objc.CallMethod[objc.Void](w_, "trackEventsMatchingMask:timeout:mode:handler:", mask, timeout, mode, trackingHandler)
}

func (w_ Window) PerformWindowDragWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](w_, "performWindowDragWithEvent:", event)
}

func (w_ Window) DisableSnapshotRestoration() {
	objc.CallMethod[objc.Void](w_, "disableSnapshotRestoration")
}

func (w_ Window) EnableSnapshotRestoration() {
	objc.CallMethod[objc.Void](w_, "enableSnapshotRestoration")
}

func (w_ Window) Display() {
	objc.CallMethod[objc.Void](w_, "display")
}

func (w_ Window) DisplayIfNeeded() {
	objc.CallMethod[objc.Void](w_, "displayIfNeeded")
}

func (w_ Window) DisableScreenUpdatesUntilFlush() {
	objc.CallMethod[objc.Void](w_, "disableScreenUpdatesUntilFlush")
}

func (w_ Window) Update() {
	objc.CallMethod[objc.Void](w_, "update")
}

func (w_ Window) DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(image IImage, baseLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool) {
	objc.CallMethod[objc.Void](w_, "dragImage:at:offset:event:pasteboard:source:slideBack:", image, baseLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}

func (w_ Window) RegisterForDraggedTypes(newTypes []PasteboardType) {
	objc.CallMethod[objc.Void](w_, "registerForDraggedTypes:", newTypes)
}

func (w_ Window) UnregisterDraggedTypes() {
	objc.CallMethod[objc.Void](w_, "unregisterDraggedTypes")
}

func (w_ Window) BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "backingAlignedRect:options:", rect, options)
	return rv
}

func (w_ Window) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "convertRectFromBacking:", rect)
	return rv
}

func (w_ Window) ConvertRectFromScreen(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "convertRectFromScreen:", rect)
	return rv
}

func (w_ Window) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, "convertPointFromBacking:", point)
	return rv
}

func (w_ Window) ConvertPointFromScreen(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, "convertPointFromScreen:", point)
	return rv
}

func (w_ Window) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "convertRectToBacking:", rect)
	return rv
}

func (w_ Window) ConvertRectToScreen(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "convertRectToScreen:", rect)
	return rv
}

func (w_ Window) ConvertPointToBacking(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, "convertPointToBacking:", point)
	return rv
}

func (w_ Window) ConvertPointToScreen(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, "convertPointToScreen:", point)
	return rv
}

func (w_ Window) SetTitleWithRepresentedFilename(filename string) {
	objc.CallMethod[objc.Void](w_, "setTitleWithRepresentedFilename:", filename)
}

func (w_ Window) Center() {
	objc.CallMethod[objc.Void](w_, "center")
}

func (w_ Window) PerformClose(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "performClose:", sender)
}

func (w_ Window) Close() {
	objc.CallMethod[objc.Void](w_, "close")
}

func (w_ Window) PerformMiniaturize(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "performMiniaturize:", sender)
}

func (w_ Window) Miniaturize(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "miniaturize:", sender)
}

func (w_ Window) Deminiaturize(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "deminiaturize:", sender)
}

func (w_ Window) Print(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, "print:", sender)
}

func (w_ Window) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	rv := objc.CallMethod[[]byte](w_, "dataWithEPSInsideRect:", rect)
	return rv
}

func (w_ Window) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	rv := objc.CallMethod[[]byte](w_, "dataWithPDFInsideRect:", rect)
	return rv
}

func (w_ Window) UpdateConstraintsIfNeeded() {
	objc.CallMethod[objc.Void](w_, "updateConstraintsIfNeeded")
}

func (w_ Window) LayoutIfNeeded() {
	objc.CallMethod[objc.Void](w_, "layoutIfNeeded")
}

func (w_ Window) VisualizeConstraints(constraints []ILayoutConstraint) {
	objc.CallMethod[objc.Void](w_, "visualizeConstraints:", constraints)
}

func (w_ Window) AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](w_, "anchorAttributeForOrientation:", orientation)
	return rv
}

func (w_ Window) SetAnchorAttribute_ForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](w_, "setAnchorAttribute:forOrientation:", attr, orientation)
}

func (w_ Window) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.CallMethod[bool](w_, "canRepresentDisplayGamut:", displayGamut)
	return rv
}

func (w_ Window) SetIsMiniaturized(flag bool) {
	objc.CallMethod[objc.Void](w_, "setIsMiniaturized:", flag)
}

func (w_ Window) SetIsVisible(flag bool) {
	objc.CallMethod[objc.Void](w_, "setIsVisible:", flag)
}

func (w_ Window) SetIsZoomed(flag bool) {
	objc.CallMethod[objc.Void](w_, "setIsZoomed:", flag)
}

func (w_ Window) HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, "handleCloseScriptCommand:", command)
	return rv
}

func (w_ Window) HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, "handlePrintScriptCommand:", command)
	return rv
}

func (w_ Window) HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, "handleSaveScriptCommand:", command)
	return rv
}

// deprecated
func (w_ Window) GState() int {
	rv := objc.CallMethod[int](w_, "gState")
	return rv
}

// deprecated
func (w_ Window) CanStoreColor() bool {
	rv := objc.CallMethod[bool](w_, "canStoreColor")
	return rv
}

// deprecated
func (w_ Window) EnableFlushWindow() {
	objc.CallMethod[objc.Void](w_, "enableFlushWindow")
}

// deprecated
func (w_ Window) DisableFlushWindow() {
	objc.CallMethod[objc.Void](w_, "disableFlushWindow")
}

// deprecated
func (w_ Window) FlushWindow() {
	objc.CallMethod[objc.Void](w_, "flushWindow")
}

// deprecated
func (w_ Window) FlushWindowIfNeeded() {
	objc.CallMethod[objc.Void](w_, "flushWindowIfNeeded")
}

// deprecated
func (wc _WindowClass) MenuChanged(menu IMenu) {
	objc.CallMethod[objc.Void](wc, "menuChanged:", menu)
}

// deprecated
func (w_ Window) CacheImageInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](w_, "cacheImageInRect:", rect)
}

// deprecated
func (w_ Window) RestoreCachedImage() {
	objc.CallMethod[objc.Void](w_, "restoreCachedImage")
}

// deprecated
func (w_ Window) DiscardCachedImage() {
	objc.CallMethod[objc.Void](w_, "discardCachedImage")
}

// deprecated
func (w_ Window) UseOptimizedDrawing(flag bool) {
	objc.CallMethod[objc.Void](w_, "useOptimizedDrawing:", flag)
}

// deprecated
func (w_ Window) ConvertBaseToScreen(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, "convertBaseToScreen:", point)
	return rv
}

// deprecated
func (w_ Window) ConvertScreenToBase(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, "convertScreenToBase:", point)
	return rv
}

// deprecated
func (w_ Window) UserSpaceScaleFactor() float64 {
	rv := objc.CallMethod[float64](w_, "userSpaceScaleFactor")
	return rv
}

func (w_ Window) InitWithWindowRef(windowRef unsafe.Pointer) Window {
	rv := objc.CallMethod[Window](w_, "initWithWindowRef:", windowRef)
	return rv
}

func (w_ Window) Delegate() WindowDelegateWrapper {
	rv := objc.CallMethod[WindowDelegateWrapper](w_, "delegate")
	return rv
}

func (w_ Window) SetDelegate(value WindowDelegate) {
	po := objc.CreateProtocol("NSWindowDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(w_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](w_, "setDelegate:", po)
}

func (w_ Window) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, "setDelegate:", value)
}

func (w_ Window) ContentViewController() ViewController {
	rv := objc.CallMethod[ViewController](w_, "contentViewController")
	return rv
}

func (w_ Window) SetContentViewController(value IViewController) {
	objc.CallMethod[objc.Void](w_, "setContentViewController:", value)
}

func (w_ Window) ContentView() View {
	rv := objc.CallMethod[View](w_, "contentView")
	return rv
}

func (w_ Window) SetContentView(value IView) {
	objc.CallMethod[objc.Void](w_, "setContentView:", value)
}

func (w_ Window) StyleMask() WindowStyleMask {
	rv := objc.CallMethod[WindowStyleMask](w_, "styleMask")
	return rv
}

func (w_ Window) SetStyleMask(value WindowStyleMask) {
	objc.CallMethod[objc.Void](w_, "setStyleMask:", value)
}

func (w_ Window) WorksWhenModal() bool {
	rv := objc.CallMethod[bool](w_, "worksWhenModal")
	return rv
}

func (w_ Window) AlphaValue() float64 {
	rv := objc.CallMethod[float64](w_, "alphaValue")
	return rv
}

func (w_ Window) SetAlphaValue(value float64) {
	objc.CallMethod[objc.Void](w_, "setAlphaValue:", value)
}

func (w_ Window) BackgroundColor() Color {
	rv := objc.CallMethod[Color](w_, "backgroundColor")
	return rv
}

func (w_ Window) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](w_, "setBackgroundColor:", value)
}

func (w_ Window) ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](w_, "colorSpace")
	return rv
}

func (w_ Window) SetColorSpace(value IColorSpace) {
	objc.CallMethod[objc.Void](w_, "setColorSpace:", value)
}

func (w_ Window) CanHide() bool {
	rv := objc.CallMethod[bool](w_, "canHide")
	return rv
}

func (w_ Window) SetCanHide(value bool) {
	objc.CallMethod[objc.Void](w_, "setCanHide:", value)
}

func (w_ Window) IsOnActiveSpace() bool {
	rv := objc.CallMethod[bool](w_, "isOnActiveSpace")
	return rv
}

func (w_ Window) HidesOnDeactivate() bool {
	rv := objc.CallMethod[bool](w_, "hidesOnDeactivate")
	return rv
}

func (w_ Window) SetHidesOnDeactivate(value bool) {
	objc.CallMethod[objc.Void](w_, "setHidesOnDeactivate:", value)
}

func (w_ Window) CollectionBehavior() WindowCollectionBehavior {
	rv := objc.CallMethod[WindowCollectionBehavior](w_, "collectionBehavior")
	return rv
}

func (w_ Window) SetCollectionBehavior(value WindowCollectionBehavior) {
	objc.CallMethod[objc.Void](w_, "setCollectionBehavior:", value)
}

func (w_ Window) IsOpaque() bool {
	rv := objc.CallMethod[bool](w_, "isOpaque")
	return rv
}

func (w_ Window) SetOpaque(value bool) {
	objc.CallMethod[objc.Void](w_, "setOpaque:", value)
}

func (w_ Window) HasShadow() bool {
	rv := objc.CallMethod[bool](w_, "hasShadow")
	return rv
}

func (w_ Window) SetHasShadow(value bool) {
	objc.CallMethod[objc.Void](w_, "setHasShadow:", value)
}

func (w_ Window) PreventsApplicationTerminationWhenModal() bool {
	rv := objc.CallMethod[bool](w_, "preventsApplicationTerminationWhenModal")
	return rv
}

func (w_ Window) SetPreventsApplicationTerminationWhenModal(value bool) {
	objc.CallMethod[objc.Void](w_, "setPreventsApplicationTerminationWhenModal:", value)
}

func (w_ Window) DepthLimit() WindowDepth {
	rv := objc.CallMethod[WindowDepth](w_, "depthLimit")
	return rv
}

func (w_ Window) SetDepthLimit(value WindowDepth) {
	objc.CallMethod[objc.Void](w_, "setDepthLimit:", value)
}

func (w_ Window) HasDynamicDepthLimit() bool {
	rv := objc.CallMethod[bool](w_, "hasDynamicDepthLimit")
	return rv
}

func (wc _WindowClass) DefaultDepthLimit() WindowDepth {
	rv := objc.CallMethod[WindowDepth](wc, "defaultDepthLimit")
	return rv
}

func (w_ Window) WindowNumber() int {
	rv := objc.CallMethod[int](w_, "windowNumber")
	return rv
}

func (w_ Window) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.CallMethod[map[DeviceDescriptionKey]objc.Object](w_, "deviceDescription")
	return rv
}

func (w_ Window) CanBecomeVisibleWithoutLogin() bool {
	rv := objc.CallMethod[bool](w_, "canBecomeVisibleWithoutLogin")
	return rv
}

func (w_ Window) SetCanBecomeVisibleWithoutLogin(value bool) {
	objc.CallMethod[objc.Void](w_, "setCanBecomeVisibleWithoutLogin:", value)
}

func (w_ Window) SharingType() WindowSharingType {
	rv := objc.CallMethod[WindowSharingType](w_, "sharingType")
	return rv
}

func (w_ Window) SetSharingType(value WindowSharingType) {
	objc.CallMethod[objc.Void](w_, "setSharingType:", value)
}

func (w_ Window) BackingType() BackingStoreType {
	rv := objc.CallMethod[BackingStoreType](w_, "backingType")
	return rv
}

func (w_ Window) SetBackingType(value BackingStoreType) {
	objc.CallMethod[objc.Void](w_, "setBackingType:", value)
}

func (w_ Window) WindowController() WindowController {
	rv := objc.CallMethod[WindowController](w_, "windowController")
	return rv
}

func (w_ Window) SetWindowController(value IWindowController) {
	objc.CallMethod[objc.Void](w_, "setWindowController:", value)
}

func (w_ Window) AttachedSheet() Window {
	rv := objc.CallMethod[Window](w_, "attachedSheet")
	return rv
}

func (w_ Window) IsSheet() bool {
	rv := objc.CallMethod[bool](w_, "isSheet")
	return rv
}

func (w_ Window) SheetParent() Window {
	rv := objc.CallMethod[Window](w_, "sheetParent")
	return rv
}

func (w_ Window) Sheets() []Window {
	rv := objc.CallMethod[[]Window](w_, "sheets")
	return rv
}

func (w_ Window) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "frame")
	return rv
}

func (w_ Window) AspectRatio() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "aspectRatio")
	return rv
}

func (w_ Window) SetAspectRatio(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setAspectRatio:", value)
}

func (w_ Window) MinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "minSize")
	return rv
}

func (w_ Window) SetMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setMinSize:", value)
}

func (w_ Window) MaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "maxSize")
	return rv
}

func (w_ Window) SetMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setMaxSize:", value)
}

func (w_ Window) IsZoomed() bool {
	rv := objc.CallMethod[bool](w_, "isZoomed")
	return rv
}

func (w_ Window) ResizeFlags() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](w_, "resizeFlags")
	return rv
}

func (w_ Window) ResizeIncrements() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "resizeIncrements")
	return rv
}

func (w_ Window) SetResizeIncrements(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setResizeIncrements:", value)
}

func (w_ Window) PreservesContentDuringLiveResize() bool {
	rv := objc.CallMethod[bool](w_, "preservesContentDuringLiveResize")
	return rv
}

func (w_ Window) SetPreservesContentDuringLiveResize(value bool) {
	objc.CallMethod[objc.Void](w_, "setPreservesContentDuringLiveResize:", value)
}

func (w_ Window) InLiveResize() bool {
	rv := objc.CallMethod[bool](w_, "inLiveResize")
	return rv
}

func (w_ Window) ContentAspectRatio() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "contentAspectRatio")
	return rv
}

func (w_ Window) SetContentAspectRatio(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setContentAspectRatio:", value)
}

func (w_ Window) ContentMinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "contentMinSize")
	return rv
}

func (w_ Window) SetContentMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setContentMinSize:", value)
}

func (w_ Window) ContentMaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "contentMaxSize")
	return rv
}

func (w_ Window) SetContentMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setContentMaxSize:", value)
}

func (w_ Window) ContentResizeIncrements() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "contentResizeIncrements")
	return rv
}

func (w_ Window) SetContentResizeIncrements(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setContentResizeIncrements:", value)
}

func (w_ Window) ContentLayoutGuide() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, "contentLayoutGuide")
	return rv
}

func (w_ Window) ContentLayoutRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, "contentLayoutRect")
	return rv
}

func (w_ Window) MaxFullScreenContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "maxFullScreenContentSize")
	return rv
}

func (w_ Window) SetMaxFullScreenContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setMaxFullScreenContentSize:", value)
}

func (w_ Window) MinFullScreenContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, "minFullScreenContentSize")
	return rv
}

func (w_ Window) SetMinFullScreenContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, "setMinFullScreenContentSize:", value)
}

func (w_ Window) Level() WindowLevel {
	rv := objc.CallMethod[WindowLevel](w_, "level")
	return rv
}

func (w_ Window) SetLevel(value WindowLevel) {
	objc.CallMethod[objc.Void](w_, "setLevel:", value)
}

func (w_ Window) IsVisible() bool {
	rv := objc.CallMethod[bool](w_, "isVisible")
	return rv
}

func (w_ Window) OcclusionState() WindowOcclusionState {
	rv := objc.CallMethod[WindowOcclusionState](w_, "occlusionState")
	return rv
}

func (w_ Window) FrameAutosaveName() WindowFrameAutosaveName {
	rv := objc.CallMethod[WindowFrameAutosaveName](w_, "frameAutosaveName")
	return rv
}

func (w_ Window) StringWithSavedFrame() WindowPersistableFrameDescriptor {
	rv := objc.CallMethod[WindowPersistableFrameDescriptor](w_, "stringWithSavedFrame")
	return rv
}

func (w_ Window) IsKeyWindow() bool {
	rv := objc.CallMethod[bool](w_, "isKeyWindow")
	return rv
}

func (w_ Window) CanBecomeKeyWindow() bool {
	rv := objc.CallMethod[bool](w_, "canBecomeKeyWindow")
	return rv
}

func (w_ Window) IsMainWindow() bool {
	rv := objc.CallMethod[bool](w_, "isMainWindow")
	return rv
}

func (w_ Window) CanBecomeMainWindow() bool {
	rv := objc.CallMethod[bool](w_, "canBecomeMainWindow")
	return rv
}

func (w_ Window) Toolbar() Toolbar {
	rv := objc.CallMethod[Toolbar](w_, "toolbar")
	return rv
}

func (w_ Window) SetToolbar(value IToolbar) {
	objc.CallMethod[objc.Void](w_, "setToolbar:", value)
}

func (w_ Window) ChildWindows() []Window {
	rv := objc.CallMethod[[]Window](w_, "childWindows")
	return rv
}

func (w_ Window) ParentWindow() Window {
	rv := objc.CallMethod[Window](w_, "parentWindow")
	return rv
}

func (w_ Window) SetParentWindow(value IWindow) {
	objc.CallMethod[objc.Void](w_, "setParentWindow:", value)
}

func (w_ Window) DefaultButtonCell() ButtonCell {
	rv := objc.CallMethod[ButtonCell](w_, "defaultButtonCell")
	return rv
}

func (w_ Window) SetDefaultButtonCell(value IButtonCell) {
	objc.CallMethod[objc.Void](w_, "setDefaultButtonCell:", value)
}

func (w_ Window) IsExcludedFromWindowsMenu() bool {
	rv := objc.CallMethod[bool](w_, "isExcludedFromWindowsMenu")
	return rv
}

func (w_ Window) SetExcludedFromWindowsMenu(value bool) {
	objc.CallMethod[objc.Void](w_, "setExcludedFromWindowsMenu:", value)
}

func (w_ Window) AreCursorRectsEnabled() bool {
	rv := objc.CallMethod[bool](w_, "areCursorRectsEnabled")
	return rv
}

// deprecated
func (w_ Window) ShowsToolbarButton() bool {
	rv := objc.CallMethod[bool](w_, "showsToolbarButton")
	return rv
}

// deprecated
func (w_ Window) SetShowsToolbarButton(value bool) {
	objc.CallMethod[objc.Void](w_, "setShowsToolbarButton:", value)
}

func (w_ Window) TitlebarAppearsTransparent() bool {
	rv := objc.CallMethod[bool](w_, "titlebarAppearsTransparent")
	return rv
}

func (w_ Window) SetTitlebarAppearsTransparent(value bool) {
	objc.CallMethod[objc.Void](w_, "setTitlebarAppearsTransparent:", value)
}

func (w_ Window) ToolbarStyle() WindowToolbarStyle {
	rv := objc.CallMethod[WindowToolbarStyle](w_, "toolbarStyle")
	return rv
}

func (w_ Window) SetToolbarStyle(value WindowToolbarStyle) {
	objc.CallMethod[objc.Void](w_, "setToolbarStyle:", value)
}

func (w_ Window) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	rv := objc.CallMethod[TitlebarSeparatorStyle](w_, "titlebarSeparatorStyle")
	return rv
}

func (w_ Window) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	objc.CallMethod[objc.Void](w_, "setTitlebarSeparatorStyle:", value)
}

func (w_ Window) WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](w_, "windowTitlebarLayoutDirection")
	return rv
}

func (w_ Window) TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController {
	rv := objc.CallMethod[[]TitlebarAccessoryViewController](w_, "titlebarAccessoryViewControllers")
	return rv
}

func (w_ Window) SetTitlebarAccessoryViewControllers(value []ITitlebarAccessoryViewController) {
	objc.CallMethod[objc.Void](w_, "setTitlebarAccessoryViewControllers:", value)
}

func (wc _WindowClass) AllowsAutomaticWindowTabbing() bool {
	rv := objc.CallMethod[bool](wc, "allowsAutomaticWindowTabbing")
	return rv
}

func (wc _WindowClass) SetAllowsAutomaticWindowTabbing(value bool) {
	objc.CallMethod[objc.Void](wc, "setAllowsAutomaticWindowTabbing:", value)
}

func (wc _WindowClass) UserTabbingPreference() WindowUserTabbingPreference {
	rv := objc.CallMethod[WindowUserTabbingPreference](wc, "userTabbingPreference")
	return rv
}

func (w_ Window) Tab() WindowTab {
	rv := objc.CallMethod[WindowTab](w_, "tab")
	return rv
}

func (w_ Window) TabbingIdentifier() WindowTabbingIdentifier {
	rv := objc.CallMethod[WindowTabbingIdentifier](w_, "tabbingIdentifier")
	return rv
}

func (w_ Window) SetTabbingIdentifier(value WindowTabbingIdentifier) {
	objc.CallMethod[objc.Void](w_, "setTabbingIdentifier:", value)
}

func (w_ Window) TabbingMode() WindowTabbingMode {
	rv := objc.CallMethod[WindowTabbingMode](w_, "tabbingMode")
	return rv
}

func (w_ Window) SetTabbingMode(value WindowTabbingMode) {
	objc.CallMethod[objc.Void](w_, "setTabbingMode:", value)
}

func (w_ Window) TabbedWindows() []Window {
	rv := objc.CallMethod[[]Window](w_, "tabbedWindows")
	return rv
}

func (w_ Window) TabGroup() WindowTabGroup {
	rv := objc.CallMethod[WindowTabGroup](w_, "tabGroup")
	return rv
}

func (w_ Window) AllowsToolTipsWhenApplicationIsInactive() bool {
	rv := objc.CallMethod[bool](w_, "allowsToolTipsWhenApplicationIsInactive")
	return rv
}

func (w_ Window) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	objc.CallMethod[objc.Void](w_, "setAllowsToolTipsWhenApplicationIsInactive:", value)
}

func (w_ Window) CurrentEvent() Event {
	rv := objc.CallMethod[Event](w_, "currentEvent")
	return rv
}

func (w_ Window) InitialFirstResponder() View {
	rv := objc.CallMethod[View](w_, "initialFirstResponder")
	return rv
}

func (w_ Window) SetInitialFirstResponder(value IView) {
	objc.CallMethod[objc.Void](w_, "setInitialFirstResponder:", value)
}

func (w_ Window) FirstResponder() Responder {
	rv := objc.CallMethod[Responder](w_, "firstResponder")
	return rv
}

func (w_ Window) KeyViewSelectionDirection() SelectionDirection {
	rv := objc.CallMethod[SelectionDirection](w_, "keyViewSelectionDirection")
	return rv
}

func (w_ Window) AutorecalculatesKeyViewLoop() bool {
	rv := objc.CallMethod[bool](w_, "autorecalculatesKeyViewLoop")
	return rv
}

func (w_ Window) SetAutorecalculatesKeyViewLoop(value bool) {
	objc.CallMethod[objc.Void](w_, "setAutorecalculatesKeyViewLoop:", value)
}

func (w_ Window) AcceptsMouseMovedEvents() bool {
	rv := objc.CallMethod[bool](w_, "acceptsMouseMovedEvents")
	return rv
}

func (w_ Window) SetAcceptsMouseMovedEvents(value bool) {
	objc.CallMethod[objc.Void](w_, "setAcceptsMouseMovedEvents:", value)
}

func (w_ Window) IgnoresMouseEvents() bool {
	rv := objc.CallMethod[bool](w_, "ignoresMouseEvents")
	return rv
}

func (w_ Window) SetIgnoresMouseEvents(value bool) {
	objc.CallMethod[objc.Void](w_, "setIgnoresMouseEvents:", value)
}

func (w_ Window) MouseLocationOutsideOfEventStream() foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, "mouseLocationOutsideOfEventStream")
	return rv
}

func (w_ Window) IsRestorable() bool {
	rv := objc.CallMethod[bool](w_, "isRestorable")
	return rv
}

func (w_ Window) SetRestorable(value bool) {
	objc.CallMethod[objc.Void](w_, "setRestorable:", value)
}

func (w_ Window) ViewsNeedDisplay() bool {
	rv := objc.CallMethod[bool](w_, "viewsNeedDisplay")
	return rv
}

func (w_ Window) SetViewsNeedDisplay(value bool) {
	objc.CallMethod[objc.Void](w_, "setViewsNeedDisplay:", value)
}

func (w_ Window) AllowsConcurrentViewDrawing() bool {
	rv := objc.CallMethod[bool](w_, "allowsConcurrentViewDrawing")
	return rv
}

func (w_ Window) SetAllowsConcurrentViewDrawing(value bool) {
	objc.CallMethod[objc.Void](w_, "setAllowsConcurrentViewDrawing:", value)
}

func (w_ Window) AnimationBehavior() WindowAnimationBehavior {
	rv := objc.CallMethod[WindowAnimationBehavior](w_, "animationBehavior")
	return rv
}

func (w_ Window) SetAnimationBehavior(value WindowAnimationBehavior) {
	objc.CallMethod[objc.Void](w_, "setAnimationBehavior:", value)
}

func (w_ Window) IsDocumentEdited() bool {
	rv := objc.CallMethod[bool](w_, "isDocumentEdited")
	return rv
}

func (w_ Window) SetDocumentEdited(value bool) {
	objc.CallMethod[objc.Void](w_, "setDocumentEdited:", value)
}

func (w_ Window) BackingScaleFactor() float64 {
	rv := objc.CallMethod[float64](w_, "backingScaleFactor")
	return rv
}

func (w_ Window) Title() string {
	rv := objc.CallMethod[string](w_, "title")
	return rv
}

func (w_ Window) SetTitle(value string) {
	objc.CallMethod[objc.Void](w_, "setTitle:", value)
}

func (w_ Window) Subtitle() string {
	rv := objc.CallMethod[string](w_, "subtitle")
	return rv
}

func (w_ Window) SetSubtitle(value string) {
	objc.CallMethod[objc.Void](w_, "setSubtitle:", value)
}

func (w_ Window) TitleVisibility() WindowTitleVisibility {
	rv := objc.CallMethod[WindowTitleVisibility](w_, "titleVisibility")
	return rv
}

func (w_ Window) SetTitleVisibility(value WindowTitleVisibility) {
	objc.CallMethod[objc.Void](w_, "setTitleVisibility:", value)
}

func (w_ Window) RepresentedFilename() string {
	rv := objc.CallMethod[string](w_, "representedFilename")
	return rv
}

func (w_ Window) SetRepresentedFilename(value string) {
	objc.CallMethod[objc.Void](w_, "setRepresentedFilename:", value)
}

func (w_ Window) RepresentedURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, "representedURL")
	return rv
}

func (w_ Window) SetRepresentedURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](w_, "setRepresentedURL:", value)
}

func (w_ Window) Screen() Screen {
	rv := objc.CallMethod[Screen](w_, "screen")
	return rv
}

func (w_ Window) DeepestScreen() Screen {
	rv := objc.CallMethod[Screen](w_, "deepestScreen")
	return rv
}

func (w_ Window) DisplaysWhenScreenProfileChanges() bool {
	rv := objc.CallMethod[bool](w_, "displaysWhenScreenProfileChanges")
	return rv
}

func (w_ Window) SetDisplaysWhenScreenProfileChanges(value bool) {
	objc.CallMethod[objc.Void](w_, "setDisplaysWhenScreenProfileChanges:", value)
}

func (w_ Window) IsMovableByWindowBackground() bool {
	rv := objc.CallMethod[bool](w_, "isMovableByWindowBackground")
	return rv
}

func (w_ Window) SetMovableByWindowBackground(value bool) {
	objc.CallMethod[objc.Void](w_, "setMovableByWindowBackground:", value)
}

func (w_ Window) IsMovable() bool {
	rv := objc.CallMethod[bool](w_, "isMovable")
	return rv
}

func (w_ Window) SetMovable(value bool) {
	objc.CallMethod[objc.Void](w_, "setMovable:", value)
}

func (w_ Window) IsReleasedWhenClosed() bool {
	rv := objc.CallMethod[bool](w_, "isReleasedWhenClosed")
	return rv
}

func (w_ Window) SetReleasedWhenClosed(value bool) {
	objc.CallMethod[objc.Void](w_, "setReleasedWhenClosed:", value)
}

func (w_ Window) IsMiniaturized() bool {
	rv := objc.CallMethod[bool](w_, "isMiniaturized")
	return rv
}

func (w_ Window) MiniwindowImage() Image {
	rv := objc.CallMethod[Image](w_, "miniwindowImage")
	return rv
}

func (w_ Window) SetMiniwindowImage(value IImage) {
	objc.CallMethod[objc.Void](w_, "setMiniwindowImage:", value)
}

func (w_ Window) MiniwindowTitle() string {
	rv := objc.CallMethod[string](w_, "miniwindowTitle")
	return rv
}

func (w_ Window) SetMiniwindowTitle(value string) {
	objc.CallMethod[objc.Void](w_, "setMiniwindowTitle:", value)
}

func (w_ Window) DockTile() DockTile {
	rv := objc.CallMethod[DockTile](w_, "dockTile")
	return rv
}

func (w_ Window) HasCloseBox() bool {
	rv := objc.CallMethod[bool](w_, "hasCloseBox")
	return rv
}

func (w_ Window) HasTitleBar() bool {
	rv := objc.CallMethod[bool](w_, "hasTitleBar")
	return rv
}

func (w_ Window) IsModalPanel() bool {
	rv := objc.CallMethod[bool](w_, "isModalPanel")
	return rv
}

func (w_ Window) IsFloatingPanel() bool {
	rv := objc.CallMethod[bool](w_, "isFloatingPanel")
	return rv
}

func (w_ Window) IsZoomable() bool {
	rv := objc.CallMethod[bool](w_, "isZoomable")
	return rv
}

func (w_ Window) IsResizable() bool {
	rv := objc.CallMethod[bool](w_, "isResizable")
	return rv
}

func (w_ Window) IsMiniaturizable() bool {
	rv := objc.CallMethod[bool](w_, "isMiniaturizable")
	return rv
}

func (w_ Window) OrderedIndex() int {
	rv := objc.CallMethod[int](w_, "orderedIndex")
	return rv
}

func (w_ Window) SetOrderedIndex(value int) {
	objc.CallMethod[objc.Void](w_, "setOrderedIndex:", value)
}

// deprecated
func (w_ Window) BackingLocation() WindowBackingLocation {
	rv := objc.CallMethod[WindowBackingLocation](w_, "backingLocation")
	return rv
}

// deprecated
func (w_ Window) PreferredBackingLocation() WindowBackingLocation {
	rv := objc.CallMethod[WindowBackingLocation](w_, "preferredBackingLocation")
	return rv
}

// deprecated
func (w_ Window) SetPreferredBackingLocation(value WindowBackingLocation) {
	objc.CallMethod[objc.Void](w_, "setPreferredBackingLocation:", value)
}

// deprecated
func (w_ Window) IsOneShot() bool {
	rv := objc.CallMethod[bool](w_, "isOneShot")
	return rv
}

// deprecated
func (w_ Window) SetOneShot(value bool) {
	objc.CallMethod[objc.Void](w_, "setOneShot:", value)
}

// deprecated
func (w_ Window) Drawers() []Drawer {
	rv := objc.CallMethod[[]Drawer](w_, "drawers")
	return rv
}

// deprecated
func (w_ Window) ShowsResizeIndicator() bool {
	rv := objc.CallMethod[bool](w_, "showsResizeIndicator")
	return rv
}

// deprecated
func (w_ Window) SetShowsResizeIndicator(value bool) {
	objc.CallMethod[objc.Void](w_, "setShowsResizeIndicator:", value)
}

// deprecated
func (w_ Window) IsFlushWindowDisabled() bool {
	rv := objc.CallMethod[bool](w_, "isFlushWindowDisabled")
	return rv
}

// deprecated
func (w_ Window) IsAutodisplay() bool {
	rv := objc.CallMethod[bool](w_, "isAutodisplay")
	return rv
}

// deprecated
func (w_ Window) SetAutodisplay(value bool) {
	objc.CallMethod[objc.Void](w_, "setAutodisplay:", value)
}

// deprecated
func (w_ Window) GraphicsContext() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](w_, "graphicsContext")
	return rv
}

func (w_ Window) WindowRef() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](w_, "windowRef")
	return rv
}
