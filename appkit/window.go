// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[Window](wc, objc.SEL("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func (w_ Window) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	rv := objc.CallMethod[Window](w_, objc.SEL("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func (w_ Window) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Window {
	rv := objc.CallMethod[Window](w_, objc.SEL("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func (w_ Window) Init() Window {
	rv := objc.CallMethod[Window](w_, objc.SEL("init"))
	return rv
}

func (wc _WindowClass) Alloc() Window {
	rv := objc.CallMethod[Window](wc, objc.SEL("alloc"))
	return rv
}

func (wc _WindowClass) New() Window {
	rv := objc.CallMethod[Window](wc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewWindow() Window {
	return WindowClass.New()
}

func (w_ Window) ToggleFullScreen(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("toggleFullScreen:"), objc.ExtractPtr(sender))
}

func (w_ Window) SetDynamicDepthLimit(flag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setDynamicDepthLimit:"), flag)
}

func (w_ Window) InvalidateShadow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("invalidateShadow"))
}

func (w_ Window) AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("autorecalculatesContentBorderThicknessForEdge:"), edge)
	return rv
}

func (w_ Window) SetAutorecalculatesContentBorderThickness_ForEdge(flag bool, edge foundation.RectEdge) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAutorecalculatesContentBorderThickness:forEdge:"), flag, edge)
}

func (w_ Window) ContentBorderThicknessForEdge(edge foundation.RectEdge) float64 {
	rv := objc.CallMethod[float64](w_, objc.SEL("contentBorderThicknessForEdge:"), edge)
	return rv
}

func (w_ Window) SetContentBorderThickness_ForEdge(thickness float64, edge foundation.RectEdge) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setContentBorderThickness:forEdge:"), thickness, edge)
}

func (wc _WindowClass) WindowNumbersWithOptions(options WindowNumberListOptions) []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](wc, objc.SEL("windowNumbersWithOptions:"), options)
	return rv
}

func (wc _WindowClass) ContentRectForFrameRect_StyleMask(fRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](wc, objc.SEL("contentRectForFrameRect:styleMask:"), fRect, style)
	return rv
}

func (wc _WindowClass) FrameRectForContentRect_StyleMask(cRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](wc, objc.SEL("frameRectForContentRect:styleMask:"), cRect, style)
	return rv
}

func (wc _WindowClass) MinFrameWidthWithTitle_StyleMask(title string, style WindowStyleMask) float64 {
	rv := objc.CallMethod[float64](wc, objc.SEL("minFrameWidthWithTitle:styleMask:"), title, style)
	return rv
}

func (w_ Window) ContentRectForFrameRect(frameRect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("contentRectForFrameRect:"), frameRect)
	return rv
}

func (w_ Window) FrameRectForContentRect(contentRect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("frameRectForContentRect:"), contentRect)
	return rv
}

func (w_ Window) BeginSheet_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("beginSheet:completionHandler:"), objc.ExtractPtr(sheetWindow), handler)
}

func (w_ Window) BeginCriticalSheet_CompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("beginCriticalSheet:completionHandler:"), objc.ExtractPtr(sheetWindow), handler)
}

func (w_ Window) EndSheet(sheetWindow IWindow) {
	objc.CallMethod[objc.Void](w_, objc.SEL("endSheet:"), objc.ExtractPtr(sheetWindow))
}

func (w_ Window) EndSheet_ReturnCode(sheetWindow IWindow, returnCode ModalResponse) {
	objc.CallMethod[objc.Void](w_, objc.SEL("endSheet:returnCode:"), objc.ExtractPtr(sheetWindow), returnCode)
}

func (w_ Window) SetFrameOrigin(point foundation.Point) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setFrameOrigin:"), point)
}

func (w_ Window) SetFrameTopLeftPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setFrameTopLeftPoint:"), point)
}

func (w_ Window) ConstrainFrameRect_ToScreen(frameRect foundation.Rect, screen IScreen) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("constrainFrameRect:toScreen:"), frameRect, objc.ExtractPtr(screen))
	return rv
}

func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.SEL("cascadeTopLeftFromPoint:"), topLeftPoint)
	return rv
}

func (w_ Window) SetFrame_Display(frameRect foundation.Rect, flag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setFrame:display:"), frameRect, flag)
}

func (w_ Window) SetFrame_Display_Animate(frameRect foundation.Rect, displayFlag bool, animateFlag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setFrame:display:animate:"), frameRect, displayFlag, animateFlag)
}

func (w_ Window) AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](w_, objc.SEL("animationResizeTime:"), newFrame)
	return rv
}

func (w_ Window) PerformZoom(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("performZoom:"), objc.ExtractPtr(sender))
}

func (w_ Window) Zoom(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("zoom:"), objc.ExtractPtr(sender))
}

func (w_ Window) SetContentSize(size foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setContentSize:"), size)
}

func (w_ Window) OrderOut(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("orderOut:"), objc.ExtractPtr(sender))
}

func (w_ Window) OrderBack(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("orderBack:"), objc.ExtractPtr(sender))
}

func (w_ Window) OrderFront(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("orderFront:"), objc.ExtractPtr(sender))
}

func (w_ Window) OrderFrontRegardless() {
	objc.CallMethod[objc.Void](w_, objc.SEL("orderFrontRegardless"))
}

func (w_ Window) OrderWindow_RelativeTo(place WindowOrderingMode, otherWin int) {
	objc.CallMethod[objc.Void](w_, objc.SEL("orderWindow:relativeTo:"), place, otherWin)
}

func (wc _WindowClass) RemoveFrameUsingName(name WindowFrameAutosaveName) {
	objc.CallMethod[objc.Void](wc, objc.SEL("removeFrameUsingName:"), name)
}

func (w_ Window) SetFrameUsingName(name WindowFrameAutosaveName) bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("setFrameUsingName:"), name)
	return rv
}

func (w_ Window) SetFrameUsingName_Force(name WindowFrameAutosaveName, force bool) bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("setFrameUsingName:force:"), name, force)
	return rv
}

func (w_ Window) SaveFrameUsingName(name WindowFrameAutosaveName) {
	objc.CallMethod[objc.Void](w_, objc.SEL("saveFrameUsingName:"), name)
}

func (w_ Window) SetFrameAutosaveName(name WindowFrameAutosaveName) bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("setFrameAutosaveName:"), name)
	return rv
}

func (w_ Window) SetFrameFromString(string_ WindowPersistableFrameDescriptor) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setFrameFromString:"), string_)
}

func (w_ Window) MakeKeyWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("makeKeyWindow"))
}

func (w_ Window) MakeKeyAndOrderFront(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("makeKeyAndOrderFront:"), objc.ExtractPtr(sender))
}

func (w_ Window) BecomeKeyWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("becomeKeyWindow"))
}

func (w_ Window) ResignKeyWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("resignKeyWindow"))
}

func (w_ Window) MakeMainWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("makeMainWindow"))
}

func (w_ Window) BecomeMainWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("becomeMainWindow"))
}

func (w_ Window) ResignMainWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("resignMainWindow"))
}

func (w_ Window) ToggleToolbarShown(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("toggleToolbarShown:"), objc.ExtractPtr(sender))
}

func (w_ Window) RunToolbarCustomizationPalette(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("runToolbarCustomizationPalette:"), objc.ExtractPtr(sender))
}

func (w_ Window) AddChildWindow_Ordered(childWin IWindow, place WindowOrderingMode) {
	objc.CallMethod[objc.Void](w_, objc.SEL("addChildWindow:ordered:"), objc.ExtractPtr(childWin), place)
}

func (w_ Window) RemoveChildWindow(childWin IWindow) {
	objc.CallMethod[objc.Void](w_, objc.SEL("removeChildWindow:"), objc.ExtractPtr(childWin))
}

func (w_ Window) EnableKeyEquivalentForDefaultButtonCell() {
	objc.CallMethod[objc.Void](w_, objc.SEL("enableKeyEquivalentForDefaultButtonCell"))
}

func (w_ Window) DisableKeyEquivalentForDefaultButtonCell() {
	objc.CallMethod[objc.Void](w_, objc.SEL("disableKeyEquivalentForDefaultButtonCell"))
}

func (w_ Window) FieldEditor_ForObject(createFlag bool, object objc.IObject) Text {
	rv := objc.CallMethod[Text](w_, objc.SEL("fieldEditor:forObject:"), createFlag, objc.ExtractPtr(object))
	return rv
}

func (w_ Window) EndEditingFor(object objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("endEditingFor:"), objc.ExtractPtr(object))
}

func (w_ Window) EnableCursorRects() {
	objc.CallMethod[objc.Void](w_, objc.SEL("enableCursorRects"))
}

func (w_ Window) DisableCursorRects() {
	objc.CallMethod[objc.Void](w_, objc.SEL("disableCursorRects"))
}

func (w_ Window) DiscardCursorRects() {
	objc.CallMethod[objc.Void](w_, objc.SEL("discardCursorRects"))
}

func (w_ Window) InvalidateCursorRectsForView(view IView) {
	objc.CallMethod[objc.Void](w_, objc.SEL("invalidateCursorRectsForView:"), objc.ExtractPtr(view))
}

func (w_ Window) ResetCursorRects() {
	objc.CallMethod[objc.Void](w_, objc.SEL("resetCursorRects"))
}

func (wc _WindowClass) StandardWindowButton_ForStyleMask(b WindowButton, styleMask WindowStyleMask) Button {
	rv := objc.CallMethod[Button](wc, objc.SEL("standardWindowButton:forStyleMask:"), b, styleMask)
	return rv
}

func (w_ Window) StandardWindowButton(b WindowButton) Button {
	rv := objc.CallMethod[Button](w_, objc.SEL("standardWindowButton:"), b)
	return rv
}

func (w_ Window) AddTitlebarAccessoryViewController(childViewController ITitlebarAccessoryViewController) {
	objc.CallMethod[objc.Void](w_, objc.SEL("addTitlebarAccessoryViewController:"), objc.ExtractPtr(childViewController))
}

func (w_ Window) InsertTitlebarAccessoryViewController_AtIndex(childViewController ITitlebarAccessoryViewController, index int) {
	objc.CallMethod[objc.Void](w_, objc.SEL("insertTitlebarAccessoryViewController:atIndex:"), objc.ExtractPtr(childViewController), index)
}

func (w_ Window) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	objc.CallMethod[objc.Void](w_, objc.SEL("removeTitlebarAccessoryViewControllerAtIndex:"), index)
}

func (w_ Window) AddTabbedWindow_Ordered(window IWindow, ordered WindowOrderingMode) {
	objc.CallMethod[objc.Void](w_, objc.SEL("addTabbedWindow:ordered:"), objc.ExtractPtr(window), ordered)
}

func (w_ Window) MergeAllWindows(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("mergeAllWindows:"), objc.ExtractPtr(sender))
}

func (w_ Window) SelectNextTab(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("selectNextTab:"), objc.ExtractPtr(sender))
}

func (w_ Window) SelectPreviousTab(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("selectPreviousTab:"), objc.ExtractPtr(sender))
}

func (w_ Window) MoveTabToNewWindow(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("moveTabToNewWindow:"), objc.ExtractPtr(sender))
}

func (w_ Window) ToggleTabBar(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("toggleTabBar:"), objc.ExtractPtr(sender))
}

func (w_ Window) ToggleTabOverview(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("toggleTabOverview:"), objc.ExtractPtr(sender))
}

func (w_ Window) NextEventMatchingMask(mask EventMask) Event {
	rv := objc.CallMethod[Event](w_, objc.SEL("nextEventMatchingMask:"), mask)
	return rv
}

func (w_ Window) NextEventMatchingMask_UntilDate_InMode_Dequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event {
	rv := objc.CallMethod[Event](w_, objc.SEL("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, objc.ExtractPtr(expiration), mode, deqFlag)
	return rv
}

func (w_ Window) DiscardEventsMatchingMask_BeforeEvent(mask EventMask, lastEvent IEvent) {
	objc.CallMethod[objc.Void](w_, objc.SEL("discardEventsMatchingMask:beforeEvent:"), mask, objc.ExtractPtr(lastEvent))
}

func (w_ Window) PostEvent_AtStart(event IEvent, flag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("postEvent:atStart:"), objc.ExtractPtr(event), flag)
}

func (w_ Window) SendEvent(event IEvent) {
	objc.CallMethod[objc.Void](w_, objc.SEL("sendEvent:"), objc.ExtractPtr(event))
}

func (w_ Window) MakeFirstResponder(responder IResponder) bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("makeFirstResponder:"), objc.ExtractPtr(responder))
	return rv
}

func (w_ Window) SelectKeyViewPrecedingView(view IView) {
	objc.CallMethod[objc.Void](w_, objc.SEL("selectKeyViewPrecedingView:"), objc.ExtractPtr(view))
}

func (w_ Window) SelectKeyViewFollowingView(view IView) {
	objc.CallMethod[objc.Void](w_, objc.SEL("selectKeyViewFollowingView:"), objc.ExtractPtr(view))
}

func (w_ Window) SelectPreviousKeyView(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("selectPreviousKeyView:"), objc.ExtractPtr(sender))
}

func (w_ Window) SelectNextKeyView(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("selectNextKeyView:"), objc.ExtractPtr(sender))
}

func (w_ Window) RecalculateKeyViewLoop() {
	objc.CallMethod[objc.Void](w_, objc.SEL("recalculateKeyViewLoop"))
}

func (wc _WindowClass) WindowNumberAtPoint_BelowWindowWithWindowNumber(point foundation.Point, windowNumber int) int {
	rv := objc.CallMethod[int](wc, objc.SEL("windowNumberAtPoint:belowWindowWithWindowNumber:"), point, windowNumber)
	return rv
}

func (w_ Window) TrackEventsMatchingMask_Timeout_Mode_Handler(mask EventMask, timeout foundation.TimeInterval, mode foundation.RunLoopMode, trackingHandler func(event Event, stop *bool)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("trackEventsMatchingMask:timeout:mode:handler:"), mask, timeout, mode, trackingHandler)
}

func (w_ Window) PerformWindowDragWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](w_, objc.SEL("performWindowDragWithEvent:"), objc.ExtractPtr(event))
}

func (w_ Window) DisableSnapshotRestoration() {
	objc.CallMethod[objc.Void](w_, objc.SEL("disableSnapshotRestoration"))
}

func (w_ Window) EnableSnapshotRestoration() {
	objc.CallMethod[objc.Void](w_, objc.SEL("enableSnapshotRestoration"))
}

func (w_ Window) Display() {
	objc.CallMethod[objc.Void](w_, objc.SEL("display"))
}

func (w_ Window) DisplayIfNeeded() {
	objc.CallMethod[objc.Void](w_, objc.SEL("displayIfNeeded"))
}

func (w_ Window) DisableScreenUpdatesUntilFlush() {
	objc.CallMethod[objc.Void](w_, objc.SEL("disableScreenUpdatesUntilFlush"))
}

func (w_ Window) Update() {
	objc.CallMethod[objc.Void](w_, objc.SEL("update"))
}

func (w_ Window) DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(image IImage, baseLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("dragImage:at:offset:event:pasteboard:source:slideBack:"), objc.ExtractPtr(image), baseLocation, initialOffset, objc.ExtractPtr(event), objc.ExtractPtr(pboard), objc.ExtractPtr(sourceObj), slideFlag)
}

func (w_ Window) RegisterForDraggedTypes(newTypes []PasteboardType) {
	objc.CallMethod[objc.Void](w_, objc.SEL("registerForDraggedTypes:"), newTypes)
}

func (w_ Window) UnregisterDraggedTypes() {
	objc.CallMethod[objc.Void](w_, objc.SEL("unregisterDraggedTypes"))
}

func (w_ Window) BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("backingAlignedRect:options:"), rect, options)
	return rv
}

func (w_ Window) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("convertRectFromBacking:"), rect)
	return rv
}

func (w_ Window) ConvertRectFromScreen(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("convertRectFromScreen:"), rect)
	return rv
}

func (w_ Window) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.SEL("convertPointFromBacking:"), point)
	return rv
}

func (w_ Window) ConvertPointFromScreen(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.SEL("convertPointFromScreen:"), point)
	return rv
}

func (w_ Window) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("convertRectToBacking:"), rect)
	return rv
}

func (w_ Window) ConvertRectToScreen(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("convertRectToScreen:"), rect)
	return rv
}

func (w_ Window) ConvertPointToBacking(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.SEL("convertPointToBacking:"), point)
	return rv
}

func (w_ Window) ConvertPointToScreen(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.SEL("convertPointToScreen:"), point)
	return rv
}

func (w_ Window) SetTitleWithRepresentedFilename(filename string) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setTitleWithRepresentedFilename:"), filename)
}

func (w_ Window) Center() {
	objc.CallMethod[objc.Void](w_, objc.SEL("center"))
}

func (w_ Window) PerformClose(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("performClose:"), objc.ExtractPtr(sender))
}

func (w_ Window) Close() {
	objc.CallMethod[objc.Void](w_, objc.SEL("close"))
}

func (w_ Window) PerformMiniaturize(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("performMiniaturize:"), objc.ExtractPtr(sender))
}

func (w_ Window) Miniaturize(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("miniaturize:"), objc.ExtractPtr(sender))
}

func (w_ Window) Deminiaturize(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("deminiaturize:"), objc.ExtractPtr(sender))
}

func (w_ Window) Print(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("print:"), objc.ExtractPtr(sender))
}

func (w_ Window) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	rv := objc.CallMethod[[]byte](w_, objc.SEL("dataWithEPSInsideRect:"), rect)
	return rv
}

func (w_ Window) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	rv := objc.CallMethod[[]byte](w_, objc.SEL("dataWithPDFInsideRect:"), rect)
	return rv
}

func (w_ Window) UpdateConstraintsIfNeeded() {
	objc.CallMethod[objc.Void](w_, objc.SEL("updateConstraintsIfNeeded"))
}

func (w_ Window) LayoutIfNeeded() {
	objc.CallMethod[objc.Void](w_, objc.SEL("layoutIfNeeded"))
}

func (w_ Window) VisualizeConstraints(constraints []ILayoutConstraint) {
	objc.CallMethod[objc.Void](w_, objc.SEL("visualizeConstraints:"), constraints)
}

func (w_ Window) AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](w_, objc.SEL("anchorAttributeForOrientation:"), orientation)
	return rv
}

func (w_ Window) SetAnchorAttribute_ForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAnchorAttribute:forOrientation:"), attr, orientation)
}

func (w_ Window) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("canRepresentDisplayGamut:"), displayGamut)
	return rv
}

func (w_ Window) SetIsMiniaturized(flag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setIsMiniaturized:"), flag)
}

func (w_ Window) SetIsVisible(flag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setIsVisible:"), flag)
}

func (w_ Window) SetIsZoomed(flag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setIsZoomed:"), flag)
}

func (w_ Window) HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("handleCloseScriptCommand:"), objc.ExtractPtr(command))
	return rv
}

func (w_ Window) HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("handlePrintScriptCommand:"), objc.ExtractPtr(command))
	return rv
}

func (w_ Window) HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("handleSaveScriptCommand:"), objc.ExtractPtr(command))
	return rv
}

// deprecated
func (w_ Window) GState() int {
	rv := objc.CallMethod[int](w_, objc.SEL("gState"))
	return rv
}

// deprecated
func (w_ Window) CanStoreColor() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("canStoreColor"))
	return rv
}

// deprecated
func (w_ Window) EnableFlushWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("enableFlushWindow"))
}

// deprecated
func (w_ Window) DisableFlushWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("disableFlushWindow"))
}

// deprecated
func (w_ Window) FlushWindow() {
	objc.CallMethod[objc.Void](w_, objc.SEL("flushWindow"))
}

// deprecated
func (w_ Window) FlushWindowIfNeeded() {
	objc.CallMethod[objc.Void](w_, objc.SEL("flushWindowIfNeeded"))
}

// deprecated
func (wc _WindowClass) MenuChanged(menu IMenu) {
	objc.CallMethod[objc.Void](wc, objc.SEL("menuChanged:"), objc.ExtractPtr(menu))
}

// deprecated
func (w_ Window) CacheImageInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](w_, objc.SEL("cacheImageInRect:"), rect)
}

// deprecated
func (w_ Window) RestoreCachedImage() {
	objc.CallMethod[objc.Void](w_, objc.SEL("restoreCachedImage"))
}

// deprecated
func (w_ Window) DiscardCachedImage() {
	objc.CallMethod[objc.Void](w_, objc.SEL("discardCachedImage"))
}

// deprecated
func (w_ Window) UseOptimizedDrawing(flag bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("useOptimizedDrawing:"), flag)
}

// deprecated
func (w_ Window) ConvertBaseToScreen(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.SEL("convertBaseToScreen:"), point)
	return rv
}

// deprecated
func (w_ Window) ConvertScreenToBase(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.SEL("convertScreenToBase:"), point)
	return rv
}

// deprecated
func (w_ Window) UserSpaceScaleFactor() float64 {
	rv := objc.CallMethod[float64](w_, objc.SEL("userSpaceScaleFactor"))
	return rv
}

func (w_ Window) InitWithWindowRef(windowRef unsafe.Pointer) Window {
	rv := objc.CallMethod[Window](w_, objc.SEL("initWithWindowRef:"), windowRef)
	return rv
}

// weak property
func (w_ Window) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("delegate"))
	return rv
}

// weak property
func (w_ Window) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (w_ Window) ContentViewController() ViewController {
	rv := objc.CallMethod[ViewController](w_, objc.SEL("contentViewController"))
	return rv
}

func (w_ Window) SetContentViewController(value IViewController) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setContentViewController:"), objc.ExtractPtr(value))
}

func (w_ Window) ContentView() View {
	rv := objc.CallMethod[View](w_, objc.SEL("contentView"))
	return rv
}

func (w_ Window) SetContentView(value IView) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setContentView:"), objc.ExtractPtr(value))
}

func (w_ Window) StyleMask() WindowStyleMask {
	rv := objc.CallMethod[WindowStyleMask](w_, objc.SEL("styleMask"))
	return rv
}

func (w_ Window) SetStyleMask(value WindowStyleMask) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setStyleMask:"), value)
}

func (w_ Window) WorksWhenModal() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("worksWhenModal"))
	return rv
}

func (w_ Window) AlphaValue() float64 {
	rv := objc.CallMethod[float64](w_, objc.SEL("alphaValue"))
	return rv
}

func (w_ Window) SetAlphaValue(value float64) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAlphaValue:"), value)
}

func (w_ Window) BackgroundColor() Color {
	rv := objc.CallMethod[Color](w_, objc.SEL("backgroundColor"))
	return rv
}

func (w_ Window) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (w_ Window) ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](w_, objc.SEL("colorSpace"))
	return rv
}

func (w_ Window) SetColorSpace(value IColorSpace) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setColorSpace:"), objc.ExtractPtr(value))
}

func (w_ Window) CanHide() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("canHide"))
	return rv
}

func (w_ Window) SetCanHide(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setCanHide:"), value)
}

func (w_ Window) IsOnActiveSpace() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isOnActiveSpace"))
	return rv
}

func (w_ Window) HidesOnDeactivate() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("hidesOnDeactivate"))
	return rv
}

func (w_ Window) SetHidesOnDeactivate(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setHidesOnDeactivate:"), value)
}

func (w_ Window) CollectionBehavior() WindowCollectionBehavior {
	rv := objc.CallMethod[WindowCollectionBehavior](w_, objc.SEL("collectionBehavior"))
	return rv
}

func (w_ Window) SetCollectionBehavior(value WindowCollectionBehavior) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setCollectionBehavior:"), value)
}

func (w_ Window) IsOpaque() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isOpaque"))
	return rv
}

func (w_ Window) SetOpaque(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setOpaque:"), value)
}

func (w_ Window) HasShadow() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("hasShadow"))
	return rv
}

func (w_ Window) SetHasShadow(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setHasShadow:"), value)
}

func (w_ Window) PreventsApplicationTerminationWhenModal() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("preventsApplicationTerminationWhenModal"))
	return rv
}

func (w_ Window) SetPreventsApplicationTerminationWhenModal(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setPreventsApplicationTerminationWhenModal:"), value)
}

func (w_ Window) DepthLimit() WindowDepth {
	rv := objc.CallMethod[WindowDepth](w_, objc.SEL("depthLimit"))
	return rv
}

func (w_ Window) SetDepthLimit(value WindowDepth) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setDepthLimit:"), value)
}

func (w_ Window) HasDynamicDepthLimit() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("hasDynamicDepthLimit"))
	return rv
}

func (wc _WindowClass) DefaultDepthLimit() WindowDepth {
	rv := objc.CallMethod[WindowDepth](wc, objc.SEL("defaultDepthLimit"))
	return rv
}

func (w_ Window) WindowNumber() int {
	rv := objc.CallMethod[int](w_, objc.SEL("windowNumber"))
	return rv
}

func (w_ Window) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.CallMethod[map[DeviceDescriptionKey]objc.Object](w_, objc.SEL("deviceDescription"))
	return rv
}

func (w_ Window) CanBecomeVisibleWithoutLogin() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("canBecomeVisibleWithoutLogin"))
	return rv
}

func (w_ Window) SetCanBecomeVisibleWithoutLogin(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setCanBecomeVisibleWithoutLogin:"), value)
}

func (w_ Window) SharingType() WindowSharingType {
	rv := objc.CallMethod[WindowSharingType](w_, objc.SEL("sharingType"))
	return rv
}

func (w_ Window) SetSharingType(value WindowSharingType) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setSharingType:"), value)
}

func (w_ Window) BackingType() BackingStoreType {
	rv := objc.CallMethod[BackingStoreType](w_, objc.SEL("backingType"))
	return rv
}

func (w_ Window) SetBackingType(value BackingStoreType) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setBackingType:"), value)
}

// weak property
func (w_ Window) WindowController() WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.SEL("windowController"))
	return rv
}

// weak property
func (w_ Window) SetWindowController(value IWindowController) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setWindowController:"), objc.ExtractPtr(value))
}

func (w_ Window) AttachedSheet() Window {
	rv := objc.CallMethod[Window](w_, objc.SEL("attachedSheet"))
	return rv
}

func (w_ Window) IsSheet() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isSheet"))
	return rv
}

func (w_ Window) SheetParent() Window {
	rv := objc.CallMethod[Window](w_, objc.SEL("sheetParent"))
	return rv
}

func (w_ Window) Sheets() []Window {
	rv := objc.CallMethod[[]Window](w_, objc.SEL("sheets"))
	return rv
}

func (w_ Window) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("frame"))
	return rv
}

func (w_ Window) AspectRatio() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("aspectRatio"))
	return rv
}

func (w_ Window) SetAspectRatio(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAspectRatio:"), value)
}

func (w_ Window) MinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("minSize"))
	return rv
}

func (w_ Window) SetMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMinSize:"), value)
}

func (w_ Window) MaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("maxSize"))
	return rv
}

func (w_ Window) SetMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMaxSize:"), value)
}

func (w_ Window) IsZoomed() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isZoomed"))
	return rv
}

func (w_ Window) ResizeFlags() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](w_, objc.SEL("resizeFlags"))
	return rv
}

func (w_ Window) ResizeIncrements() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("resizeIncrements"))
	return rv
}

func (w_ Window) SetResizeIncrements(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setResizeIncrements:"), value)
}

func (w_ Window) PreservesContentDuringLiveResize() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("preservesContentDuringLiveResize"))
	return rv
}

func (w_ Window) SetPreservesContentDuringLiveResize(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setPreservesContentDuringLiveResize:"), value)
}

func (w_ Window) InLiveResize() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("inLiveResize"))
	return rv
}

func (w_ Window) ContentAspectRatio() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("contentAspectRatio"))
	return rv
}

func (w_ Window) SetContentAspectRatio(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setContentAspectRatio:"), value)
}

func (w_ Window) ContentMinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("contentMinSize"))
	return rv
}

func (w_ Window) SetContentMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setContentMinSize:"), value)
}

func (w_ Window) ContentMaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("contentMaxSize"))
	return rv
}

func (w_ Window) SetContentMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setContentMaxSize:"), value)
}

func (w_ Window) ContentResizeIncrements() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("contentResizeIncrements"))
	return rv
}

func (w_ Window) SetContentResizeIncrements(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setContentResizeIncrements:"), value)
}

func (w_ Window) ContentLayoutGuide() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.SEL("contentLayoutGuide"))
	return rv
}

func (w_ Window) ContentLayoutRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.SEL("contentLayoutRect"))
	return rv
}

func (w_ Window) MaxFullScreenContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("maxFullScreenContentSize"))
	return rv
}

func (w_ Window) SetMaxFullScreenContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMaxFullScreenContentSize:"), value)
}

func (w_ Window) MinFullScreenContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.SEL("minFullScreenContentSize"))
	return rv
}

func (w_ Window) SetMinFullScreenContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMinFullScreenContentSize:"), value)
}

func (w_ Window) Level() WindowLevel {
	rv := objc.CallMethod[WindowLevel](w_, objc.SEL("level"))
	return rv
}

func (w_ Window) SetLevel(value WindowLevel) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setLevel:"), value)
}

func (w_ Window) IsVisible() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isVisible"))
	return rv
}

func (w_ Window) OcclusionState() WindowOcclusionState {
	rv := objc.CallMethod[WindowOcclusionState](w_, objc.SEL("occlusionState"))
	return rv
}

func (w_ Window) FrameAutosaveName() WindowFrameAutosaveName {
	rv := objc.CallMethod[WindowFrameAutosaveName](w_, objc.SEL("frameAutosaveName"))
	return rv
}

func (w_ Window) StringWithSavedFrame() WindowPersistableFrameDescriptor {
	rv := objc.CallMethod[WindowPersistableFrameDescriptor](w_, objc.SEL("stringWithSavedFrame"))
	return rv
}

func (w_ Window) IsKeyWindow() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isKeyWindow"))
	return rv
}

func (w_ Window) CanBecomeKeyWindow() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("canBecomeKeyWindow"))
	return rv
}

func (w_ Window) IsMainWindow() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isMainWindow"))
	return rv
}

func (w_ Window) CanBecomeMainWindow() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("canBecomeMainWindow"))
	return rv
}

func (w_ Window) Toolbar() Toolbar {
	rv := objc.CallMethod[Toolbar](w_, objc.SEL("toolbar"))
	return rv
}

func (w_ Window) SetToolbar(value IToolbar) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setToolbar:"), objc.ExtractPtr(value))
}

func (w_ Window) ChildWindows() []Window {
	rv := objc.CallMethod[[]Window](w_, objc.SEL("childWindows"))
	return rv
}

// weak property
func (w_ Window) ParentWindow() Window {
	rv := objc.CallMethod[Window](w_, objc.SEL("parentWindow"))
	return rv
}

// weak property
func (w_ Window) SetParentWindow(value IWindow) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setParentWindow:"), objc.ExtractPtr(value))
}

func (w_ Window) DefaultButtonCell() ButtonCell {
	rv := objc.CallMethod[ButtonCell](w_, objc.SEL("defaultButtonCell"))
	return rv
}

func (w_ Window) SetDefaultButtonCell(value IButtonCell) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setDefaultButtonCell:"), objc.ExtractPtr(value))
}

func (w_ Window) IsExcludedFromWindowsMenu() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isExcludedFromWindowsMenu"))
	return rv
}

func (w_ Window) SetExcludedFromWindowsMenu(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setExcludedFromWindowsMenu:"), value)
}

func (w_ Window) AreCursorRectsEnabled() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("areCursorRectsEnabled"))
	return rv
}

// deprecated
func (w_ Window) ShowsToolbarButton() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("showsToolbarButton"))
	return rv
}

// deprecated
func (w_ Window) SetShowsToolbarButton(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setShowsToolbarButton:"), value)
}

func (w_ Window) TitlebarAppearsTransparent() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("titlebarAppearsTransparent"))
	return rv
}

func (w_ Window) SetTitlebarAppearsTransparent(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setTitlebarAppearsTransparent:"), value)
}

func (w_ Window) ToolbarStyle() WindowToolbarStyle {
	rv := objc.CallMethod[WindowToolbarStyle](w_, objc.SEL("toolbarStyle"))
	return rv
}

func (w_ Window) SetToolbarStyle(value WindowToolbarStyle) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setToolbarStyle:"), value)
}

func (w_ Window) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	rv := objc.CallMethod[TitlebarSeparatorStyle](w_, objc.SEL("titlebarSeparatorStyle"))
	return rv
}

func (w_ Window) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setTitlebarSeparatorStyle:"), value)
}

func (w_ Window) WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](w_, objc.SEL("windowTitlebarLayoutDirection"))
	return rv
}

func (w_ Window) TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController {
	rv := objc.CallMethod[[]TitlebarAccessoryViewController](w_, objc.SEL("titlebarAccessoryViewControllers"))
	return rv
}

func (w_ Window) SetTitlebarAccessoryViewControllers(value []ITitlebarAccessoryViewController) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setTitlebarAccessoryViewControllers:"), value)
}

func (wc _WindowClass) AllowsAutomaticWindowTabbing() bool {
	rv := objc.CallMethod[bool](wc, objc.SEL("allowsAutomaticWindowTabbing"))
	return rv
}

func (wc _WindowClass) SetAllowsAutomaticWindowTabbing(value bool) {
	objc.CallMethod[objc.Void](wc, objc.SEL("setAllowsAutomaticWindowTabbing:"), value)
}

func (wc _WindowClass) UserTabbingPreference() WindowUserTabbingPreference {
	rv := objc.CallMethod[WindowUserTabbingPreference](wc, objc.SEL("userTabbingPreference"))
	return rv
}

func (w_ Window) Tab() WindowTab {
	rv := objc.CallMethod[WindowTab](w_, objc.SEL("tab"))
	return rv
}

func (w_ Window) TabbingIdentifier() WindowTabbingIdentifier {
	rv := objc.CallMethod[WindowTabbingIdentifier](w_, objc.SEL("tabbingIdentifier"))
	return rv
}

func (w_ Window) SetTabbingIdentifier(value WindowTabbingIdentifier) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setTabbingIdentifier:"), value)
}

func (w_ Window) TabbingMode() WindowTabbingMode {
	rv := objc.CallMethod[WindowTabbingMode](w_, objc.SEL("tabbingMode"))
	return rv
}

func (w_ Window) SetTabbingMode(value WindowTabbingMode) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setTabbingMode:"), value)
}

func (w_ Window) TabbedWindows() []Window {
	rv := objc.CallMethod[[]Window](w_, objc.SEL("tabbedWindows"))
	return rv
}

// weak property
func (w_ Window) TabGroup() WindowTabGroup {
	rv := objc.CallMethod[WindowTabGroup](w_, objc.SEL("tabGroup"))
	return rv
}

func (w_ Window) AllowsToolTipsWhenApplicationIsInactive() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("allowsToolTipsWhenApplicationIsInactive"))
	return rv
}

func (w_ Window) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAllowsToolTipsWhenApplicationIsInactive:"), value)
}

func (w_ Window) CurrentEvent() Event {
	rv := objc.CallMethod[Event](w_, objc.SEL("currentEvent"))
	return rv
}

// weak property
func (w_ Window) InitialFirstResponder() View {
	rv := objc.CallMethod[View](w_, objc.SEL("initialFirstResponder"))
	return rv
}

// weak property
func (w_ Window) SetInitialFirstResponder(value IView) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setInitialFirstResponder:"), objc.ExtractPtr(value))
}

// weak property
func (w_ Window) FirstResponder() Responder {
	rv := objc.CallMethod[Responder](w_, objc.SEL("firstResponder"))
	return rv
}

func (w_ Window) KeyViewSelectionDirection() SelectionDirection {
	rv := objc.CallMethod[SelectionDirection](w_, objc.SEL("keyViewSelectionDirection"))
	return rv
}

func (w_ Window) AutorecalculatesKeyViewLoop() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("autorecalculatesKeyViewLoop"))
	return rv
}

func (w_ Window) SetAutorecalculatesKeyViewLoop(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAutorecalculatesKeyViewLoop:"), value)
}

func (w_ Window) AcceptsMouseMovedEvents() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("acceptsMouseMovedEvents"))
	return rv
}

func (w_ Window) SetAcceptsMouseMovedEvents(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAcceptsMouseMovedEvents:"), value)
}

func (w_ Window) IgnoresMouseEvents() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("ignoresMouseEvents"))
	return rv
}

func (w_ Window) SetIgnoresMouseEvents(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setIgnoresMouseEvents:"), value)
}

func (w_ Window) MouseLocationOutsideOfEventStream() foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.SEL("mouseLocationOutsideOfEventStream"))
	return rv
}

func (w_ Window) IsRestorable() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isRestorable"))
	return rv
}

func (w_ Window) SetRestorable(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setRestorable:"), value)
}

func (w_ Window) ViewsNeedDisplay() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("viewsNeedDisplay"))
	return rv
}

func (w_ Window) SetViewsNeedDisplay(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setViewsNeedDisplay:"), value)
}

func (w_ Window) AllowsConcurrentViewDrawing() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("allowsConcurrentViewDrawing"))
	return rv
}

func (w_ Window) SetAllowsConcurrentViewDrawing(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAllowsConcurrentViewDrawing:"), value)
}

func (w_ Window) AnimationBehavior() WindowAnimationBehavior {
	rv := objc.CallMethod[WindowAnimationBehavior](w_, objc.SEL("animationBehavior"))
	return rv
}

func (w_ Window) SetAnimationBehavior(value WindowAnimationBehavior) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAnimationBehavior:"), value)
}

func (w_ Window) IsDocumentEdited() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isDocumentEdited"))
	return rv
}

func (w_ Window) SetDocumentEdited(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setDocumentEdited:"), value)
}

func (w_ Window) BackingScaleFactor() float64 {
	rv := objc.CallMethod[float64](w_, objc.SEL("backingScaleFactor"))
	return rv
}

func (w_ Window) Title() string {
	rv := objc.CallMethod[string](w_, objc.SEL("title"))
	return rv
}

func (w_ Window) SetTitle(value string) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setTitle:"), value)
}

func (w_ Window) Subtitle() string {
	rv := objc.CallMethod[string](w_, objc.SEL("subtitle"))
	return rv
}

func (w_ Window) SetSubtitle(value string) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setSubtitle:"), value)
}

func (w_ Window) TitleVisibility() WindowTitleVisibility {
	rv := objc.CallMethod[WindowTitleVisibility](w_, objc.SEL("titleVisibility"))
	return rv
}

func (w_ Window) SetTitleVisibility(value WindowTitleVisibility) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setTitleVisibility:"), value)
}

func (w_ Window) RepresentedFilename() string {
	rv := objc.CallMethod[string](w_, objc.SEL("representedFilename"))
	return rv
}

func (w_ Window) SetRepresentedFilename(value string) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setRepresentedFilename:"), value)
}

func (w_ Window) RepresentedURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.SEL("representedURL"))
	return rv
}

func (w_ Window) SetRepresentedURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setRepresentedURL:"), objc.ExtractPtr(value))
}

func (w_ Window) Screen() Screen {
	rv := objc.CallMethod[Screen](w_, objc.SEL("screen"))
	return rv
}

func (w_ Window) DeepestScreen() Screen {
	rv := objc.CallMethod[Screen](w_, objc.SEL("deepestScreen"))
	return rv
}

func (w_ Window) DisplaysWhenScreenProfileChanges() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("displaysWhenScreenProfileChanges"))
	return rv
}

func (w_ Window) SetDisplaysWhenScreenProfileChanges(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setDisplaysWhenScreenProfileChanges:"), value)
}

func (w_ Window) IsMovableByWindowBackground() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isMovableByWindowBackground"))
	return rv
}

func (w_ Window) SetMovableByWindowBackground(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMovableByWindowBackground:"), value)
}

func (w_ Window) IsMovable() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isMovable"))
	return rv
}

func (w_ Window) SetMovable(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMovable:"), value)
}

func (w_ Window) IsReleasedWhenClosed() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isReleasedWhenClosed"))
	return rv
}

func (w_ Window) SetReleasedWhenClosed(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setReleasedWhenClosed:"), value)
}

func (w_ Window) IsMiniaturized() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isMiniaturized"))
	return rv
}

func (w_ Window) MiniwindowImage() Image {
	rv := objc.CallMethod[Image](w_, objc.SEL("miniwindowImage"))
	return rv
}

func (w_ Window) SetMiniwindowImage(value IImage) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMiniwindowImage:"), objc.ExtractPtr(value))
}

func (w_ Window) MiniwindowTitle() string {
	rv := objc.CallMethod[string](w_, objc.SEL("miniwindowTitle"))
	return rv
}

func (w_ Window) SetMiniwindowTitle(value string) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setMiniwindowTitle:"), value)
}

func (w_ Window) DockTile() DockTile {
	rv := objc.CallMethod[DockTile](w_, objc.SEL("dockTile"))
	return rv
}

func (w_ Window) HasCloseBox() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("hasCloseBox"))
	return rv
}

func (w_ Window) HasTitleBar() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("hasTitleBar"))
	return rv
}

func (w_ Window) IsModalPanel() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isModalPanel"))
	return rv
}

func (w_ Window) IsFloatingPanel() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isFloatingPanel"))
	return rv
}

func (w_ Window) IsZoomable() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isZoomable"))
	return rv
}

func (w_ Window) IsResizable() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isResizable"))
	return rv
}

func (w_ Window) IsMiniaturizable() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isMiniaturizable"))
	return rv
}

func (w_ Window) OrderedIndex() int {
	rv := objc.CallMethod[int](w_, objc.SEL("orderedIndex"))
	return rv
}

func (w_ Window) SetOrderedIndex(value int) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setOrderedIndex:"), value)
}

// deprecated
func (w_ Window) BackingLocation() WindowBackingLocation {
	rv := objc.CallMethod[WindowBackingLocation](w_, objc.SEL("backingLocation"))
	return rv
}

// deprecated
func (w_ Window) PreferredBackingLocation() WindowBackingLocation {
	rv := objc.CallMethod[WindowBackingLocation](w_, objc.SEL("preferredBackingLocation"))
	return rv
}

// deprecated
func (w_ Window) SetPreferredBackingLocation(value WindowBackingLocation) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setPreferredBackingLocation:"), value)
}

// deprecated
func (w_ Window) IsOneShot() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isOneShot"))
	return rv
}

// deprecated
func (w_ Window) SetOneShot(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setOneShot:"), value)
}

// deprecated
func (w_ Window) Drawers() []Drawer {
	rv := objc.CallMethod[[]Drawer](w_, objc.SEL("drawers"))
	return rv
}

// deprecated
func (w_ Window) ShowsResizeIndicator() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("showsResizeIndicator"))
	return rv
}

// deprecated
func (w_ Window) SetShowsResizeIndicator(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setShowsResizeIndicator:"), value)
}

// deprecated
func (w_ Window) IsFlushWindowDisabled() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isFlushWindowDisabled"))
	return rv
}

// deprecated
func (w_ Window) IsAutodisplay() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isAutodisplay"))
	return rv
}

// deprecated
func (w_ Window) SetAutodisplay(value bool) {
	objc.CallMethod[objc.Void](w_, objc.SEL("setAutodisplay:"), value)
}

// deprecated
func (w_ Window) GraphicsContext() GraphicsContext {
	rv := objc.CallMethod[GraphicsContext](w_, objc.SEL("graphicsContext"))
	return rv
}

func (w_ Window) WindowRef() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](w_, objc.SEL("windowRef"))
	return rv
}
