// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/quartzcore"
)

var ViewClass = _ViewClass{objc.GetClass("NSView")}

type _ViewClass struct {
	objc.Class
}

type IView interface {
	IResponder
	PrepareForReuse()
	IsDescendantOf(view IView) bool
	AncestorSharedWithView(view IView) View
	AddSubview(view IView)
	AddSubview_Positioned_RelativeTo(view IView, place WindowOrderingMode, otherView IView)
	RemoveFromSuperview()
	RemoveFromSuperviewWithoutNeedingDisplay()
	ReplaceSubview_With(oldView IView, newView IView)
	DidAddSubview(subview IView)
	ViewDidMoveToSuperview()
	ViewDidMoveToWindow()
	ViewWillMoveToSuperview(newSuperview IView)
	ViewWillMoveToWindow(newWindow IWindow)
	WillRemoveSubview(subview IView)
	ViewWithTag(tag int) View
	SetFrameOrigin(newOrigin foundation.Point)
	SetFrameSize(newSize foundation.Size)
	SetBoundsOrigin(newOrigin foundation.Point)
	SetBoundsSize(newSize foundation.Size)
	TranslateOriginToPoint(translation foundation.Point)
	ScaleUnitSquareToSize(newUnitSize foundation.Size)
	RotateByAngle(angle float64)
	BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect
	ConvertPointFromBacking(point foundation.Point) foundation.Point
	ConvertPointToBacking(point foundation.Point) foundation.Point
	ConvertPointFromLayer(point foundation.Point) foundation.Point
	ConvertPointToLayer(point foundation.Point) foundation.Point
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	ConvertRectFromLayer(rect foundation.Rect) foundation.Rect
	ConvertRectToLayer(rect foundation.Rect) foundation.Rect
	ConvertSizeFromBacking(size foundation.Size) foundation.Size
	ConvertSizeToBacking(size foundation.Size) foundation.Size
	ConvertSizeFromLayer(size foundation.Size) foundation.Size
	ConvertSizeToLayer(size foundation.Size) foundation.Size
	ConvertPoint_FromView(point foundation.Point, view IView) foundation.Point
	ConvertPoint_ToView(point foundation.Point, view IView) foundation.Point
	ConvertSize_FromView(size foundation.Size, view IView) foundation.Size
	ConvertSize_ToView(size foundation.Size, view IView) foundation.Size
	ConvertRect_FromView(rect foundation.Rect, view IView) foundation.Rect
	ConvertRect_ToView(rect foundation.Rect, view IView) foundation.Rect
	CenterScanRect(rect foundation.Rect) foundation.Rect
	ViewDidHide()
	ViewDidUnhide()
	ViewDidChangeEffectiveAppearance()
	ViewDidChangeBackingProperties()
	DrawFocusRingMask()
	NoteFocusRingMaskChanged()
	SetKeyboardFocusRingNeedsDisplayInRect(rect foundation.Rect)
	MakeBackingLayer() quartzcore.Layer
	MenuForEvent(event IEvent) Menu
	WillOpenMenu_WithEvent(menu IMenu, event IEvent)
	DidCloseMenu_WithEvent(menu IMenu, event IEvent)
	AddCursorRect_Cursor(rect foundation.Rect, object ICursor)
	RemoveCursorRect_Cursor(rect foundation.Rect, object ICursor)
	DiscardCursorRects()
	ResetCursorRects()
	AddToolTipRect_Owner_UserData(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer) ToolTipTag
	RemoveAllToolTips()
	RemoveToolTip(tag ToolTipTag)
	ShowDefinitionForAttributedString_AtPoint(attrString foundation.IAttributedString, textBaselineOrigin foundation.Point)
	ShowDefinitionForAttributedString_Range_Options_BaselineOriginProvider(attrString foundation.IAttributedString, targetRange foundation.Range, options map[DefinitionOptionKey]objc.IObject, originProvider func(adjustedRange foundation.Range) foundation.Point)
	RulerView_DidAddMarker(ruler IRulerView, marker IRulerMarker)
	RulerView_DidMoveMarker(ruler IRulerView, marker IRulerMarker)
	RulerView_DidRemoveMarker(ruler IRulerView, marker IRulerMarker)
	RulerView_HandleMouseDown(ruler IRulerView, event IEvent)
	RulerView_LocationForPoint(ruler IRulerView, point foundation.Point) float64
	RulerView_PointForLocation(ruler IRulerView, point float64) foundation.Point
	RulerView_ShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerView_ShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerView_ShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerView_WillAddMarker_AtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64
	RulerView_WillMoveMarker_ToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64
	RulerView_WillSetClientView(ruler IRulerView, newClient IView)
	AddConstraint(constraint ILayoutConstraint)
	AddConstraints(constraints []ILayoutConstraint)
	RemoveConstraint(constraint ILayoutConstraint)
	RemoveConstraints(constraints []ILayoutConstraint)
	InvalidateIntrinsicContentSize()
	ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	SetContentCompressionResistancePriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation)
	ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	SetContentHuggingPriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation)
	AddLayoutGuide(guide ILayoutGuide)
	RemoveLayoutGuide(guide ILayoutGuide)
	AlignmentRectForFrame(frame foundation.Rect) foundation.Rect
	FrameForAlignmentRect(alignmentRect foundation.Rect) foundation.Rect
	Layout()
	LayoutSubtreeIfNeeded()
	UpdateConstraints()
	UpdateConstraintsForSubtreeIfNeeded()
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
	ExerciseAmbiguityInLayout()
	ResizeSubviewsWithOldSize(oldSize foundation.Size)
	ResizeWithOldSuperviewSize(oldSize foundation.Size)
	UpdateLayer()
	DrawRect(dirtyRect foundation.Rect)
	NeedsToDrawRect(rect foundation.Rect) bool
	BitmapImageRepForCachingDisplayInRect(rect foundation.Rect) BitmapImageRep
	CacheDisplayInRect_ToBitmapImageRep(rect foundation.Rect, bitmapImageRep IBitmapImageRep)
	EnterFullScreenMode_WithOptions(screen IScreen, options map[ViewFullScreenModeOptionKey]objc.IObject) bool
	ExitFullScreenModeWithOptions(options map[ViewFullScreenModeOptionKey]objc.IObject)
	SetNeedsDisplayInRect(invalidRect foundation.Rect)
	Display()
	DisplayRect(rect foundation.Rect)
	DisplayRectIgnoringOpacity(rect foundation.Rect)
	DisplayRectIgnoringOpacity_InContext(rect foundation.Rect, context IGraphicsContext)
	DisplayIfNeeded()
	DisplayIfNeededInRect(rect foundation.Rect)
	DisplayIfNeededIgnoringOpacity()
	DisplayIfNeededInRectIgnoringOpacity(rect foundation.Rect)
	TranslateRectsNeedingDisplayInRect_By(clipRect foundation.Rect, delta foundation.Size)
	ViewWillDraw()
	GetRectsExposedDuringLiveResize_Count(exposedRects *foundation.Rect, count *int)
	ViewWillStartLiveResize()
	ViewDidEndLiveResize()
	Print(sender objc.IObject)
	BeginPageInRect_AtPlacement(rect foundation.Rect, location foundation.Point)
	DataWithEPSInsideRect(rect foundation.Rect) []byte
	DataWithPDFInsideRect(rect foundation.Rect) []byte
	WriteEPSInsideRect_ToPasteboard(rect foundation.Rect, pasteboard IPasteboard)
	WritePDFInsideRect_ToPasteboard(rect foundation.Rect, pasteboard IPasteboard)
	DrawPageBorderWithSize(borderSize foundation.Size)
	// deprecated
	DrawSheetBorderWithSize(borderSize foundation.Size)
	AdjustPageWidthNew_Left_Right_Limit(newRight *float64, oldLeft float64, oldRight float64, rightLimit float64)
	AdjustPageHeightNew_Top_Bottom_Limit(newBottom *float64, oldTop float64, oldBottom float64, bottomLimit float64)
	KnowsPageRange(range_ *foundation.Range) bool
	RectForPage(page int) foundation.Rect
	LocationOfPrintRect(rect foundation.Rect) foundation.Point
	BeginDocument()
	EndDocument()
	EndPage()
	AcceptsFirstMouse(event IEvent) bool
	HitTest(point foundation.Point) View
	Mouse_InRect(point foundation.Point, rect foundation.Rect) bool
	AddGestureRecognizer(gestureRecognizer IGestureRecognizer)
	RemoveGestureRecognizer(gestureRecognizer IGestureRecognizer)
	RectForSmartMagnificationAtPoint_InRect(location foundation.Point, visibleRect foundation.Rect) foundation.Rect
	AddTrackingArea(trackingArea ITrackingArea)
	RemoveTrackingArea(trackingArea ITrackingArea)
	UpdateTrackingAreas()
	AddTrackingRect_Owner_UserData_AssumeInside(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer, flag bool) TrackingRectTag
	RemoveTrackingRect(tag TrackingRectTag)
	PrepareContentInRect(rect foundation.Rect)
	ScrollPoint(point foundation.Point)
	ScrollRectToVisible(rect foundation.Rect) bool
	Autoscroll(event IEvent) bool
	AdjustScroll(newVisible foundation.Rect) foundation.Rect
	ScrollClipView_ToPoint(clipView IClipView, point foundation.Point)
	ReflectScrolledClipView(clipView IClipView)
	RegisterForDraggedTypes(newTypes []PasteboardType)
	UnregisterDraggedTypes()
	BeginDraggingSessionWithItems_Event_Source(items []IDraggingItem, event IEvent, source DraggingSource) DraggingSession
	BeginDraggingSessionWithItems0_Event_Source(items []IDraggingItem, event IEvent, source objc.IObject) DraggingSession
	ShouldDelayWindowOrderingForEvent(event IEvent) bool
	// deprecated
	LockFocus()
	// deprecated
	LockFocusIfCanDraw() bool
	// deprecated
	LockFocusIfCanDrawInContext(context IGraphicsContext) bool
	// deprecated
	UnlockFocus()
	// deprecated
	ScrollRect_By(rect foundation.Rect, delta foundation.Size)
	// deprecated
	ConvertPointToBase(point foundation.Point) foundation.Point
	// deprecated
	ConvertPointFromBase(point foundation.Point) foundation.Point
	// deprecated
	ConvertSizeToBase(size foundation.Size) foundation.Size
	// deprecated
	ConvertSizeFromBase(size foundation.Size) foundation.Size
	// deprecated
	ConvertRectToBase(rect foundation.Rect) foundation.Rect
	// deprecated
	ConvertRectFromBase(rect foundation.Rect) foundation.Rect
	// deprecated
	ShouldDrawColor() bool
	// deprecated
	AllocateGState()
	// deprecated
	GState() int
	// deprecated
	SetUpGState()
	// deprecated
	RenewGState()
	// deprecated
	DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(image IImage, viewLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool)
	// deprecated
	DragFile_FromRect_SlideBack_Event(filename string, rect foundation.Rect, flag bool, event IEvent) bool
	// deprecated
	DragPromisedFilesOfTypes_FromRect_Source_SlideBack_Event(typeArray []string, rect foundation.Rect, sourceObject objc.IObject, flag bool, event IEvent) bool
	Superview() View
	Subviews() []View
	SetSubviews(value []IView)
	Window() Window
	OpaqueAncestor() View
	EnclosingMenuItem() MenuItem
	EnclosingScrollView() ScrollView
	Tag() int
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
	FrameRotation() float64
	SetFrameRotation(value float64)
	PostsFrameChangedNotifications() bool
	SetPostsFrameChangedNotifications(value bool)
	Bounds() foundation.Rect
	SetBounds(value foundation.Rect)
	BoundsRotation() float64
	SetBoundsRotation(value float64)
	PostsBoundsChangedNotifications() bool
	SetPostsBoundsChangedNotifications(value bool)
	IsFlipped() bool
	IsRotatedFromBase() bool
	IsRotatedOrScaledFromBase() bool
	IsHidden() bool
	SetHidden(value bool)
	IsHiddenOrHasHiddenAncestor() bool
	AllowsVibrancy() bool
	FocusRingType() FocusRingType
	SetFocusRingType(value FocusRingType)
	FocusRingMaskBounds() foundation.Rect
	IsDrawingFindIndicator() bool
	WantsLayer() bool
	SetWantsLayer(value bool)
	WantsUpdateLayer() bool
	Layer() quartzcore.Layer
	SetLayer(value quartzcore.ILayer)
	LayerContentsPlacement() ViewLayerContentsPlacement
	SetLayerContentsPlacement(value ViewLayerContentsPlacement)
	LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy
	SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy)
	CanDrawSubviewsIntoLayer() bool
	SetCanDrawSubviewsIntoLayer(value bool)
	LayerUsesCoreImageFilters() bool
	SetLayerUsesCoreImageFilters(value bool)
	AlphaValue() float64
	SetAlphaValue(value float64)
	FrameCenterRotation() float64
	SetFrameCenterRotation(value float64)
	Shadow() Shadow
	SetShadow(value IShadow)
	ToolTip() string
	SetToolTip(value string)
	SafeAreaRect() foundation.Rect
	SafeAreaInsets() foundation.EdgeInsets
	AdditionalSafeAreaInsets() foundation.EdgeInsets
	SetAdditionalSafeAreaInsets(value foundation.EdgeInsets)
	SafeAreaLayoutGuide() LayoutGuide
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	TranslatesAutoresizingMaskIntoConstraints() bool
	SetTranslatesAutoresizingMaskIntoConstraints(value bool)
	BottomAnchor() LayoutYAxisAnchor
	CenterXAnchor() LayoutXAxisAnchor
	CenterYAnchor() LayoutYAxisAnchor
	FirstBaselineAnchor() LayoutYAxisAnchor
	HeightAnchor() LayoutDimension
	LastBaselineAnchor() LayoutYAxisAnchor
	LeadingAnchor() LayoutXAxisAnchor
	LeftAnchor() LayoutXAxisAnchor
	RightAnchor() LayoutXAxisAnchor
	TopAnchor() LayoutYAxisAnchor
	TrailingAnchor() LayoutXAxisAnchor
	WidthAnchor() LayoutDimension
	Constraints() []LayoutConstraint
	FittingSize() foundation.Size
	IntrinsicContentSize() foundation.Size
	LayoutGuides() []LayoutGuide
	LayoutMarginsGuide() LayoutGuide
	AlignmentRectInsets() foundation.EdgeInsets
	BaselineOffsetFromBottom() float64
	FirstBaselineOffsetFromTop() float64
	LastBaselineOffsetFromBottom() float64
	NeedsLayout() bool
	SetNeedsLayout(value bool)
	NeedsUpdateConstraints() bool
	SetNeedsUpdateConstraints(value bool)
	IsHorizontalContentSizeConstraintActive() bool
	SetHorizontalContentSizeConstraintActive(value bool)
	IsVerticalContentSizeConstraintActive() bool
	SetVerticalContentSizeConstraintActive(value bool)
	HasAmbiguousLayout() bool
	AutoresizesSubviews() bool
	SetAutoresizesSubviews(value bool)
	AutoresizingMask() AutoresizingMaskOptions
	SetAutoresizingMask(value AutoresizingMaskOptions)
	CanDrawConcurrently() bool
	SetCanDrawConcurrently(value bool)
	VisibleRect() foundation.Rect
	WantsDefaultClipping() bool
	IsInFullScreenMode() bool
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
	IsOpaque() bool
	InLiveResize() bool
	PreservesContentDuringLiveResize() bool
	RectPreservedDuringLiveResize() foundation.Rect
	PrintJobTitle() string
	PageHeader() foundation.AttributedString
	PageFooter() foundation.AttributedString
	HeightAdjustLimit() float64
	WidthAdjustLimit() float64
	MouseDownCanMoveWindow() bool
	InputContext() TextInputContext
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
	WantsRestingTouches() bool
	SetWantsRestingTouches(value bool)
	CandidateListTouchBarItem() CandidateListTouchBarItem
	GestureRecognizers() []GestureRecognizer
	SetGestureRecognizers(value []IGestureRecognizer)
	CanBecomeKeyView() bool
	NeedsPanelToBecomeKey() bool
	NextKeyView() View
	SetNextKeyView(value IView)
	NextValidKeyView() View
	PreviousKeyView() View
	PreviousValidKeyView() View
	TrackingAreas() []TrackingArea
	PreparedContentRect() foundation.Rect
	SetPreparedContentRect(value foundation.Rect)
	PressureConfiguration() PressureConfiguration
	SetPressureConfiguration(value IPressureConfiguration)
	RegisteredDraggedTypes() []PasteboardType
	// deprecated
	WantsExtendedDynamicRangeOpenGLSurface() bool
	// deprecated
	SetWantsExtendedDynamicRangeOpenGLSurface(value bool)
	// deprecated
	AcceptsTouchEvents() bool
	// deprecated
	SetAcceptsTouchEvents(value bool)
	// deprecated
	CanDraw() bool
	// deprecated
	WantsBestResolutionOpenGLSurface() bool
	// deprecated
	SetWantsBestResolutionOpenGLSurface(value bool)
}

type View struct {
	Responder
}

func MakeView(ptr unsafe.Pointer) View {
	return View{
		Responder: MakeResponder(ptr),
	}
}

func (v_ View) InitWithFrame(frameRect foundation.Rect) View {
	rv := ffi.CallMethod[View](v_, "initWithFrame:", frameRect)
	return rv
}

func (v_ View) Init() View {
	rv := ffi.CallMethod[View](v_, "init")
	return rv
}

func (vc _ViewClass) Alloc() View {
	rv := ffi.CallMethod[View](vc, "alloc")
	return rv
}

func (vc _ViewClass) New() View {
	rv := ffi.CallMethod[View](vc, "new")
	rv.Autorelease()
	return rv
}

func NewView() View {
	return ViewClass.New()
}

func (v_ View) PrepareForReuse() {
	ffi.CallMethod[ffi.Void](v_, "prepareForReuse")
}

func (v_ View) IsDescendantOf(view IView) bool {
	rv := ffi.CallMethod[bool](v_, "isDescendantOf:", view)
	return rv
}

func (v_ View) AncestorSharedWithView(view IView) View {
	rv := ffi.CallMethod[View](v_, "ancestorSharedWithView:", view)
	return rv
}

func (v_ View) AddSubview(view IView) {
	ffi.CallMethod[ffi.Void](v_, "addSubview:", view)
}

func (v_ View) AddSubview_Positioned_RelativeTo(view IView, place WindowOrderingMode, otherView IView) {
	ffi.CallMethod[ffi.Void](v_, "addSubview:positioned:relativeTo:", view, place, otherView)
}

func (v_ View) RemoveFromSuperview() {
	ffi.CallMethod[ffi.Void](v_, "removeFromSuperview")
}

func (v_ View) RemoveFromSuperviewWithoutNeedingDisplay() {
	ffi.CallMethod[ffi.Void](v_, "removeFromSuperviewWithoutNeedingDisplay")
}

func (v_ View) ReplaceSubview_With(oldView IView, newView IView) {
	ffi.CallMethod[ffi.Void](v_, "replaceSubview:with:", oldView, newView)
}

func (v_ View) DidAddSubview(subview IView) {
	ffi.CallMethod[ffi.Void](v_, "didAddSubview:", subview)
}

func (v_ View) ViewDidMoveToSuperview() {
	ffi.CallMethod[ffi.Void](v_, "viewDidMoveToSuperview")
}

func (v_ View) ViewDidMoveToWindow() {
	ffi.CallMethod[ffi.Void](v_, "viewDidMoveToWindow")
}

func (v_ View) ViewWillMoveToSuperview(newSuperview IView) {
	ffi.CallMethod[ffi.Void](v_, "viewWillMoveToSuperview:", newSuperview)
}

func (v_ View) ViewWillMoveToWindow(newWindow IWindow) {
	ffi.CallMethod[ffi.Void](v_, "viewWillMoveToWindow:", newWindow)
}

func (v_ View) WillRemoveSubview(subview IView) {
	ffi.CallMethod[ffi.Void](v_, "willRemoveSubview:", subview)
}

func (v_ View) ViewWithTag(tag int) View {
	rv := ffi.CallMethod[View](v_, "viewWithTag:", tag)
	return rv
}

func (v_ View) SetFrameOrigin(newOrigin foundation.Point) {
	ffi.CallMethod[ffi.Void](v_, "setFrameOrigin:", newOrigin)
}

func (v_ View) SetFrameSize(newSize foundation.Size) {
	ffi.CallMethod[ffi.Void](v_, "setFrameSize:", newSize)
}

func (v_ View) SetBoundsOrigin(newOrigin foundation.Point) {
	ffi.CallMethod[ffi.Void](v_, "setBoundsOrigin:", newOrigin)
}

func (v_ View) SetBoundsSize(newSize foundation.Size) {
	ffi.CallMethod[ffi.Void](v_, "setBoundsSize:", newSize)
}

func (v_ View) TranslateOriginToPoint(translation foundation.Point) {
	ffi.CallMethod[ffi.Void](v_, "translateOriginToPoint:", translation)
}

func (v_ View) ScaleUnitSquareToSize(newUnitSize foundation.Size) {
	ffi.CallMethod[ffi.Void](v_, "scaleUnitSquareToSize:", newUnitSize)
}

func (v_ View) RotateByAngle(angle float64) {
	ffi.CallMethod[ffi.Void](v_, "rotateByAngle:", angle)
}

func (v_ View) BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "backingAlignedRect:options:", rect, options)
	return rv
}

func (v_ View) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "convertPointFromBacking:", point)
	return rv
}

func (v_ View) ConvertPointToBacking(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "convertPointToBacking:", point)
	return rv
}

func (v_ View) ConvertPointFromLayer(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "convertPointFromLayer:", point)
	return rv
}

func (v_ View) ConvertPointToLayer(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "convertPointToLayer:", point)
	return rv
}

func (v_ View) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "convertRectFromBacking:", rect)
	return rv
}

func (v_ View) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "convertRectToBacking:", rect)
	return rv
}

func (v_ View) ConvertRectFromLayer(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "convertRectFromLayer:", rect)
	return rv
}

func (v_ View) ConvertRectToLayer(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "convertRectToLayer:", rect)
	return rv
}

func (v_ View) ConvertSizeFromBacking(size foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "convertSizeFromBacking:", size)
	return rv
}

func (v_ View) ConvertSizeToBacking(size foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "convertSizeToBacking:", size)
	return rv
}

func (v_ View) ConvertSizeFromLayer(size foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "convertSizeFromLayer:", size)
	return rv
}

func (v_ View) ConvertSizeToLayer(size foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "convertSizeToLayer:", size)
	return rv
}

func (v_ View) ConvertPoint_FromView(point foundation.Point, view IView) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "convertPoint:fromView:", point, view)
	return rv
}

func (v_ View) ConvertPoint_ToView(point foundation.Point, view IView) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "convertPoint:toView:", point, view)
	return rv
}

func (v_ View) ConvertSize_FromView(size foundation.Size, view IView) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "convertSize:fromView:", size, view)
	return rv
}

func (v_ View) ConvertSize_ToView(size foundation.Size, view IView) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "convertSize:toView:", size, view)
	return rv
}

func (v_ View) ConvertRect_FromView(rect foundation.Rect, view IView) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "convertRect:fromView:", rect, view)
	return rv
}

func (v_ View) ConvertRect_ToView(rect foundation.Rect, view IView) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "convertRect:toView:", rect, view)
	return rv
}

func (v_ View) CenterScanRect(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "centerScanRect:", rect)
	return rv
}

func (v_ View) ViewDidHide() {
	ffi.CallMethod[ffi.Void](v_, "viewDidHide")
}

func (v_ View) ViewDidUnhide() {
	ffi.CallMethod[ffi.Void](v_, "viewDidUnhide")
}

func (v_ View) ViewDidChangeEffectiveAppearance() {
	ffi.CallMethod[ffi.Void](v_, "viewDidChangeEffectiveAppearance")
}

func (v_ View) ViewDidChangeBackingProperties() {
	ffi.CallMethod[ffi.Void](v_, "viewDidChangeBackingProperties")
}

func (v_ View) DrawFocusRingMask() {
	ffi.CallMethod[ffi.Void](v_, "drawFocusRingMask")
}

func (v_ View) NoteFocusRingMaskChanged() {
	ffi.CallMethod[ffi.Void](v_, "noteFocusRingMaskChanged")
}

func (v_ View) SetKeyboardFocusRingNeedsDisplayInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "setKeyboardFocusRingNeedsDisplayInRect:", rect)
}

func (v_ View) MakeBackingLayer() quartzcore.Layer {
	rv := ffi.CallMethod[quartzcore.Layer](v_, "makeBackingLayer")
	return rv
}

func (v_ View) MenuForEvent(event IEvent) Menu {
	rv := ffi.CallMethod[Menu](v_, "menuForEvent:", event)
	return rv
}

func (v_ View) WillOpenMenu_WithEvent(menu IMenu, event IEvent) {
	ffi.CallMethod[ffi.Void](v_, "willOpenMenu:withEvent:", menu, event)
}

func (v_ View) DidCloseMenu_WithEvent(menu IMenu, event IEvent) {
	ffi.CallMethod[ffi.Void](v_, "didCloseMenu:withEvent:", menu, event)
}

func (v_ View) AddCursorRect_Cursor(rect foundation.Rect, object ICursor) {
	ffi.CallMethod[ffi.Void](v_, "addCursorRect:cursor:", rect, object)
}

func (v_ View) RemoveCursorRect_Cursor(rect foundation.Rect, object ICursor) {
	ffi.CallMethod[ffi.Void](v_, "removeCursorRect:cursor:", rect, object)
}

func (v_ View) DiscardCursorRects() {
	ffi.CallMethod[ffi.Void](v_, "discardCursorRects")
}

func (v_ View) ResetCursorRects() {
	ffi.CallMethod[ffi.Void](v_, "resetCursorRects")
}

func (v_ View) AddToolTipRect_Owner_UserData(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer) ToolTipTag {
	rv := ffi.CallMethod[ToolTipTag](v_, "addToolTipRect:owner:userData:", rect, owner, data)
	return rv
}

func (v_ View) RemoveAllToolTips() {
	ffi.CallMethod[ffi.Void](v_, "removeAllToolTips")
}

func (v_ View) RemoveToolTip(tag ToolTipTag) {
	ffi.CallMethod[ffi.Void](v_, "removeToolTip:", tag)
}

func (v_ View) ShowDefinitionForAttributedString_AtPoint(attrString foundation.IAttributedString, textBaselineOrigin foundation.Point) {
	ffi.CallMethod[ffi.Void](v_, "showDefinitionForAttributedString:atPoint:", attrString, textBaselineOrigin)
}

func (v_ View) ShowDefinitionForAttributedString_Range_Options_BaselineOriginProvider(attrString foundation.IAttributedString, targetRange foundation.Range, options map[DefinitionOptionKey]objc.IObject, originProvider func(adjustedRange foundation.Range) foundation.Point) {
	ffi.CallMethod[ffi.Void](v_, "showDefinitionForAttributedString:range:options:baselineOriginProvider:", attrString, targetRange, options, originProvider)
}

func (v_ View) RulerView_DidAddMarker(ruler IRulerView, marker IRulerMarker) {
	ffi.CallMethod[ffi.Void](v_, "rulerView:didAddMarker:", ruler, marker)
}

func (v_ View) RulerView_DidMoveMarker(ruler IRulerView, marker IRulerMarker) {
	ffi.CallMethod[ffi.Void](v_, "rulerView:didMoveMarker:", ruler, marker)
}

func (v_ View) RulerView_DidRemoveMarker(ruler IRulerView, marker IRulerMarker) {
	ffi.CallMethod[ffi.Void](v_, "rulerView:didRemoveMarker:", ruler, marker)
}

func (v_ View) RulerView_HandleMouseDown(ruler IRulerView, event IEvent) {
	ffi.CallMethod[ffi.Void](v_, "rulerView:handleMouseDown:", ruler, event)
}

func (v_ View) RulerView_LocationForPoint(ruler IRulerView, point foundation.Point) float64 {
	rv := ffi.CallMethod[float64](v_, "rulerView:locationForPoint:", ruler, point)
	return rv
}

func (v_ View) RulerView_PointForLocation(ruler IRulerView, point float64) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "rulerView:pointForLocation:", ruler, point)
	return rv
}

func (v_ View) RulerView_ShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := ffi.CallMethod[bool](v_, "rulerView:shouldAddMarker:", ruler, marker)
	return rv
}

func (v_ View) RulerView_ShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := ffi.CallMethod[bool](v_, "rulerView:shouldMoveMarker:", ruler, marker)
	return rv
}

func (v_ View) RulerView_ShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := ffi.CallMethod[bool](v_, "rulerView:shouldRemoveMarker:", ruler, marker)
	return rv
}

func (v_ View) RulerView_WillAddMarker_AtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := ffi.CallMethod[float64](v_, "rulerView:willAddMarker:atLocation:", ruler, marker, location)
	return rv
}

func (v_ View) RulerView_WillMoveMarker_ToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := ffi.CallMethod[float64](v_, "rulerView:willMoveMarker:toLocation:", ruler, marker, location)
	return rv
}

func (v_ View) RulerView_WillSetClientView(ruler IRulerView, newClient IView) {
	ffi.CallMethod[ffi.Void](v_, "rulerView:willSetClientView:", ruler, newClient)
}

func (v_ View) AddConstraint(constraint ILayoutConstraint) {
	ffi.CallMethod[ffi.Void](v_, "addConstraint:", constraint)
}

func (v_ View) AddConstraints(constraints []ILayoutConstraint) {
	ffi.CallMethod[ffi.Void](v_, "addConstraints:", constraints)
}

func (v_ View) RemoveConstraint(constraint ILayoutConstraint) {
	ffi.CallMethod[ffi.Void](v_, "removeConstraint:", constraint)
}

func (v_ View) RemoveConstraints(constraints []ILayoutConstraint) {
	ffi.CallMethod[ffi.Void](v_, "removeConstraints:", constraints)
}

func (v_ View) InvalidateIntrinsicContentSize() {
	ffi.CallMethod[ffi.Void](v_, "invalidateIntrinsicContentSize")
}

func (v_ View) ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := ffi.CallMethod[LayoutPriority](v_, "contentCompressionResistancePriorityForOrientation:", orientation)
	return rv
}

func (v_ View) SetContentCompressionResistancePriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	ffi.CallMethod[ffi.Void](v_, "setContentCompressionResistancePriority:forOrientation:", priority, orientation)
}

func (v_ View) ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := ffi.CallMethod[LayoutPriority](v_, "contentHuggingPriorityForOrientation:", orientation)
	return rv
}

func (v_ View) SetContentHuggingPriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	ffi.CallMethod[ffi.Void](v_, "setContentHuggingPriority:forOrientation:", priority, orientation)
}

func (v_ View) AddLayoutGuide(guide ILayoutGuide) {
	ffi.CallMethod[ffi.Void](v_, "addLayoutGuide:", guide)
}

func (v_ View) RemoveLayoutGuide(guide ILayoutGuide) {
	ffi.CallMethod[ffi.Void](v_, "removeLayoutGuide:", guide)
}

func (v_ View) AlignmentRectForFrame(frame foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "alignmentRectForFrame:", frame)
	return rv
}

func (v_ View) FrameForAlignmentRect(alignmentRect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "frameForAlignmentRect:", alignmentRect)
	return rv
}

func (v_ View) Layout() {
	ffi.CallMethod[ffi.Void](v_, "layout")
}

func (v_ View) LayoutSubtreeIfNeeded() {
	ffi.CallMethod[ffi.Void](v_, "layoutSubtreeIfNeeded")
}

func (v_ View) UpdateConstraints() {
	ffi.CallMethod[ffi.Void](v_, "updateConstraints")
}

func (v_ View) UpdateConstraintsForSubtreeIfNeeded() {
	ffi.CallMethod[ffi.Void](v_, "updateConstraintsForSubtreeIfNeeded")
}

func (v_ View) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := ffi.CallMethod[[]LayoutConstraint](v_, "constraintsAffectingLayoutForOrientation:", orientation)
	return rv
}

func (v_ View) ExerciseAmbiguityInLayout() {
	ffi.CallMethod[ffi.Void](v_, "exerciseAmbiguityInLayout")
}

func (v_ View) ResizeSubviewsWithOldSize(oldSize foundation.Size) {
	ffi.CallMethod[ffi.Void](v_, "resizeSubviewsWithOldSize:", oldSize)
}

func (v_ View) ResizeWithOldSuperviewSize(oldSize foundation.Size) {
	ffi.CallMethod[ffi.Void](v_, "resizeWithOldSuperviewSize:", oldSize)
}

func (v_ View) UpdateLayer() {
	ffi.CallMethod[ffi.Void](v_, "updateLayer")
}

func (v_ View) DrawRect(dirtyRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "drawRect:", dirtyRect)
}

func (v_ View) NeedsToDrawRect(rect foundation.Rect) bool {
	rv := ffi.CallMethod[bool](v_, "needsToDrawRect:", rect)
	return rv
}

func (v_ View) BitmapImageRepForCachingDisplayInRect(rect foundation.Rect) BitmapImageRep {
	rv := ffi.CallMethod[BitmapImageRep](v_, "bitmapImageRepForCachingDisplayInRect:", rect)
	return rv
}

func (v_ View) CacheDisplayInRect_ToBitmapImageRep(rect foundation.Rect, bitmapImageRep IBitmapImageRep) {
	ffi.CallMethod[ffi.Void](v_, "cacheDisplayInRect:toBitmapImageRep:", rect, bitmapImageRep)
}

func (v_ View) EnterFullScreenMode_WithOptions(screen IScreen, options map[ViewFullScreenModeOptionKey]objc.IObject) bool {
	rv := ffi.CallMethod[bool](v_, "enterFullScreenMode:withOptions:", screen, options)
	return rv
}

func (v_ View) ExitFullScreenModeWithOptions(options map[ViewFullScreenModeOptionKey]objc.IObject) {
	ffi.CallMethod[ffi.Void](v_, "exitFullScreenModeWithOptions:", options)
}

func (v_ View) SetNeedsDisplayInRect(invalidRect foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "setNeedsDisplayInRect:", invalidRect)
}

func (v_ View) Display() {
	ffi.CallMethod[ffi.Void](v_, "display")
}

func (v_ View) DisplayRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "displayRect:", rect)
}

func (v_ View) DisplayRectIgnoringOpacity(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "displayRectIgnoringOpacity:", rect)
}

func (v_ View) DisplayRectIgnoringOpacity_InContext(rect foundation.Rect, context IGraphicsContext) {
	ffi.CallMethod[ffi.Void](v_, "displayRectIgnoringOpacity:inContext:", rect, context)
}

func (v_ View) DisplayIfNeeded() {
	ffi.CallMethod[ffi.Void](v_, "displayIfNeeded")
}

func (v_ View) DisplayIfNeededInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "displayIfNeededInRect:", rect)
}

func (v_ View) DisplayIfNeededIgnoringOpacity() {
	ffi.CallMethod[ffi.Void](v_, "displayIfNeededIgnoringOpacity")
}

func (v_ View) DisplayIfNeededInRectIgnoringOpacity(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "displayIfNeededInRectIgnoringOpacity:", rect)
}

func (v_ View) TranslateRectsNeedingDisplayInRect_By(clipRect foundation.Rect, delta foundation.Size) {
	ffi.CallMethod[ffi.Void](v_, "translateRectsNeedingDisplayInRect:by:", clipRect, delta)
}

func (v_ View) ViewWillDraw() {
	ffi.CallMethod[ffi.Void](v_, "viewWillDraw")
}

func (v_ View) GetRectsExposedDuringLiveResize_Count(exposedRects *foundation.Rect, count *int) {
	ffi.CallMethod[ffi.Void](v_, "getRectsExposedDuringLiveResize:count:", exposedRects, count)
}

func (v_ View) ViewWillStartLiveResize() {
	ffi.CallMethod[ffi.Void](v_, "viewWillStartLiveResize")
}

func (v_ View) ViewDidEndLiveResize() {
	ffi.CallMethod[ffi.Void](v_, "viewDidEndLiveResize")
}

func (v_ View) Print(sender objc.IObject) {
	ffi.CallMethod[ffi.Void](v_, "print:", sender)
}

func (v_ View) BeginPageInRect_AtPlacement(rect foundation.Rect, location foundation.Point) {
	ffi.CallMethod[ffi.Void](v_, "beginPageInRect:atPlacement:", rect, location)
}

func (v_ View) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	rv := ffi.CallMethod[[]byte](v_, "dataWithEPSInsideRect:", rect)
	return rv
}

func (v_ View) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	rv := ffi.CallMethod[[]byte](v_, "dataWithPDFInsideRect:", rect)
	return rv
}

func (v_ View) WriteEPSInsideRect_ToPasteboard(rect foundation.Rect, pasteboard IPasteboard) {
	ffi.CallMethod[ffi.Void](v_, "writeEPSInsideRect:toPasteboard:", rect, pasteboard)
}

func (v_ View) WritePDFInsideRect_ToPasteboard(rect foundation.Rect, pasteboard IPasteboard) {
	ffi.CallMethod[ffi.Void](v_, "writePDFInsideRect:toPasteboard:", rect, pasteboard)
}

func (v_ View) DrawPageBorderWithSize(borderSize foundation.Size) {
	ffi.CallMethod[ffi.Void](v_, "drawPageBorderWithSize:", borderSize)
}

// deprecated
func (v_ View) DrawSheetBorderWithSize(borderSize foundation.Size) {
	ffi.CallMethod[ffi.Void](v_, "drawSheetBorderWithSize:", borderSize)
}

func (v_ View) AdjustPageWidthNew_Left_Right_Limit(newRight *float64, oldLeft float64, oldRight float64, rightLimit float64) {
	ffi.CallMethod[ffi.Void](v_, "adjustPageWidthNew:left:right:limit:", newRight, oldLeft, oldRight, rightLimit)
}

func (v_ View) AdjustPageHeightNew_Top_Bottom_Limit(newBottom *float64, oldTop float64, oldBottom float64, bottomLimit float64) {
	ffi.CallMethod[ffi.Void](v_, "adjustPageHeightNew:top:bottom:limit:", newBottom, oldTop, oldBottom, bottomLimit)
}

func (v_ View) KnowsPageRange(range_ *foundation.Range) bool {
	rv := ffi.CallMethod[bool](v_, "knowsPageRange:", range_)
	return rv
}

func (v_ View) RectForPage(page int) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "rectForPage:", page)
	return rv
}

func (v_ View) LocationOfPrintRect(rect foundation.Rect) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "locationOfPrintRect:", rect)
	return rv
}

func (v_ View) BeginDocument() {
	ffi.CallMethod[ffi.Void](v_, "beginDocument")
}

func (v_ View) EndDocument() {
	ffi.CallMethod[ffi.Void](v_, "endDocument")
}

func (v_ View) EndPage() {
	ffi.CallMethod[ffi.Void](v_, "endPage")
}

func (v_ View) AcceptsFirstMouse(event IEvent) bool {
	rv := ffi.CallMethod[bool](v_, "acceptsFirstMouse:", event)
	return rv
}

func (v_ View) HitTest(point foundation.Point) View {
	rv := ffi.CallMethod[View](v_, "hitTest:", point)
	return rv
}

func (v_ View) Mouse_InRect(point foundation.Point, rect foundation.Rect) bool {
	rv := ffi.CallMethod[bool](v_, "mouse:inRect:", point, rect)
	return rv
}

func (v_ View) AddGestureRecognizer(gestureRecognizer IGestureRecognizer) {
	ffi.CallMethod[ffi.Void](v_, "addGestureRecognizer:", gestureRecognizer)
}

func (v_ View) RemoveGestureRecognizer(gestureRecognizer IGestureRecognizer) {
	ffi.CallMethod[ffi.Void](v_, "removeGestureRecognizer:", gestureRecognizer)
}

func (v_ View) RectForSmartMagnificationAtPoint_InRect(location foundation.Point, visibleRect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "rectForSmartMagnificationAtPoint:inRect:", location, visibleRect)
	return rv
}

func (v_ View) AddTrackingArea(trackingArea ITrackingArea) {
	ffi.CallMethod[ffi.Void](v_, "addTrackingArea:", trackingArea)
}

func (v_ View) RemoveTrackingArea(trackingArea ITrackingArea) {
	ffi.CallMethod[ffi.Void](v_, "removeTrackingArea:", trackingArea)
}

func (v_ View) UpdateTrackingAreas() {
	ffi.CallMethod[ffi.Void](v_, "updateTrackingAreas")
}

func (v_ View) AddTrackingRect_Owner_UserData_AssumeInside(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer, flag bool) TrackingRectTag {
	rv := ffi.CallMethod[TrackingRectTag](v_, "addTrackingRect:owner:userData:assumeInside:", rect, owner, data, flag)
	return rv
}

func (v_ View) RemoveTrackingRect(tag TrackingRectTag) {
	ffi.CallMethod[ffi.Void](v_, "removeTrackingRect:", tag)
}

func (v_ View) PrepareContentInRect(rect foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "prepareContentInRect:", rect)
}

func (v_ View) ScrollPoint(point foundation.Point) {
	ffi.CallMethod[ffi.Void](v_, "scrollPoint:", point)
}

func (v_ View) ScrollRectToVisible(rect foundation.Rect) bool {
	rv := ffi.CallMethod[bool](v_, "scrollRectToVisible:", rect)
	return rv
}

func (v_ View) Autoscroll(event IEvent) bool {
	rv := ffi.CallMethod[bool](v_, "autoscroll:", event)
	return rv
}

func (v_ View) AdjustScroll(newVisible foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "adjustScroll:", newVisible)
	return rv
}

func (v_ View) ScrollClipView_ToPoint(clipView IClipView, point foundation.Point) {
	ffi.CallMethod[ffi.Void](v_, "scrollClipView:toPoint:", clipView, point)
}

func (v_ View) ReflectScrolledClipView(clipView IClipView) {
	ffi.CallMethod[ffi.Void](v_, "reflectScrolledClipView:", clipView)
}

func (v_ View) RegisterForDraggedTypes(newTypes []PasteboardType) {
	ffi.CallMethod[ffi.Void](v_, "registerForDraggedTypes:", newTypes)
}

func (v_ View) UnregisterDraggedTypes() {
	ffi.CallMethod[ffi.Void](v_, "unregisterDraggedTypes")
}

func (v_ View) BeginDraggingSessionWithItems_Event_Source(items []IDraggingItem, event IEvent, source DraggingSource) DraggingSession {
	po := ffi.CreateProtocol("NSDraggingSource", source)
	defer po.Release()
	rv := ffi.CallMethod[DraggingSession](v_, "beginDraggingSessionWithItems:event:source:", items, event, po)
	return rv
}

func (v_ View) BeginDraggingSessionWithItems0_Event_Source(items []IDraggingItem, event IEvent, source objc.IObject) DraggingSession {
	rv := ffi.CallMethod[DraggingSession](v_, "beginDraggingSessionWithItems:event:source:", items, event, source)
	return rv
}

func (v_ View) ShouldDelayWindowOrderingForEvent(event IEvent) bool {
	rv := ffi.CallMethod[bool](v_, "shouldDelayWindowOrderingForEvent:", event)
	return rv
}

// deprecated
func (v_ View) LockFocus() {
	ffi.CallMethod[ffi.Void](v_, "lockFocus")
}

// deprecated
func (v_ View) LockFocusIfCanDraw() bool {
	rv := ffi.CallMethod[bool](v_, "lockFocusIfCanDraw")
	return rv
}

// deprecated
func (v_ View) LockFocusIfCanDrawInContext(context IGraphicsContext) bool {
	rv := ffi.CallMethod[bool](v_, "lockFocusIfCanDrawInContext:", context)
	return rv
}

// deprecated
func (v_ View) UnlockFocus() {
	ffi.CallMethod[ffi.Void](v_, "unlockFocus")
}

// deprecated
func (v_ View) ScrollRect_By(rect foundation.Rect, delta foundation.Size) {
	ffi.CallMethod[ffi.Void](v_, "scrollRect:by:", rect, delta)
}

// deprecated
func (v_ View) ConvertPointToBase(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "convertPointToBase:", point)
	return rv
}

// deprecated
func (v_ View) ConvertPointFromBase(point foundation.Point) foundation.Point {
	rv := ffi.CallMethod[foundation.Point](v_, "convertPointFromBase:", point)
	return rv
}

// deprecated
func (v_ View) ConvertSizeToBase(size foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "convertSizeToBase:", size)
	return rv
}

// deprecated
func (v_ View) ConvertSizeFromBase(size foundation.Size) foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "convertSizeFromBase:", size)
	return rv
}

// deprecated
func (v_ View) ConvertRectToBase(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "convertRectToBase:", rect)
	return rv
}

// deprecated
func (v_ View) ConvertRectFromBase(rect foundation.Rect) foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "convertRectFromBase:", rect)
	return rv
}

// deprecated
func (v_ View) ShouldDrawColor() bool {
	rv := ffi.CallMethod[bool](v_, "shouldDrawColor")
	return rv
}

// deprecated
func (v_ View) AllocateGState() {
	ffi.CallMethod[ffi.Void](v_, "allocateGState")
}

// deprecated
func (v_ View) GState() int {
	rv := ffi.CallMethod[int](v_, "gState")
	return rv
}

// deprecated
func (v_ View) SetUpGState() {
	ffi.CallMethod[ffi.Void](v_, "setUpGState")
}

// deprecated
func (v_ View) RenewGState() {
	ffi.CallMethod[ffi.Void](v_, "renewGState")
}

// deprecated
func (v_ View) DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(image IImage, viewLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool) {
	ffi.CallMethod[ffi.Void](v_, "dragImage:at:offset:event:pasteboard:source:slideBack:", image, viewLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}

// deprecated
func (v_ View) DragFile_FromRect_SlideBack_Event(filename string, rect foundation.Rect, flag bool, event IEvent) bool {
	rv := ffi.CallMethod[bool](v_, "dragFile:fromRect:slideBack:event:", filename, rect, flag, event)
	return rv
}

// deprecated
func (v_ View) DragPromisedFilesOfTypes_FromRect_Source_SlideBack_Event(typeArray []string, rect foundation.Rect, sourceObject objc.IObject, flag bool, event IEvent) bool {
	rv := ffi.CallMethod[bool](v_, "dragPromisedFilesOfTypes:fromRect:source:slideBack:event:", typeArray, rect, sourceObject, flag, event)
	return rv
}

func (v_ View) Superview() View {
	rv := ffi.CallMethod[View](v_, "superview")
	return rv
}

func (v_ View) Subviews() []View {
	rv := ffi.CallMethod[[]View](v_, "subviews")
	return rv
}

func (v_ View) SetSubviews(value []IView) {
	ffi.CallMethod[ffi.Void](v_, "setSubviews:", value)
}

func (v_ View) Window() Window {
	rv := ffi.CallMethod[Window](v_, "window")
	return rv
}

func (v_ View) OpaqueAncestor() View {
	rv := ffi.CallMethod[View](v_, "opaqueAncestor")
	return rv
}

func (v_ View) EnclosingMenuItem() MenuItem {
	rv := ffi.CallMethod[MenuItem](v_, "enclosingMenuItem")
	return rv
}

func (v_ View) EnclosingScrollView() ScrollView {
	rv := ffi.CallMethod[ScrollView](v_, "enclosingScrollView")
	return rv
}

func (v_ View) Tag() int {
	rv := ffi.CallMethod[int](v_, "tag")
	return rv
}

func (v_ View) Frame() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "frame")
	return rv
}

func (v_ View) SetFrame(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "setFrame:", value)
}

func (v_ View) FrameRotation() float64 {
	rv := ffi.CallMethod[float64](v_, "frameRotation")
	return rv
}

func (v_ View) SetFrameRotation(value float64) {
	ffi.CallMethod[ffi.Void](v_, "setFrameRotation:", value)
}

func (v_ View) PostsFrameChangedNotifications() bool {
	rv := ffi.CallMethod[bool](v_, "postsFrameChangedNotifications")
	return rv
}

func (v_ View) SetPostsFrameChangedNotifications(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setPostsFrameChangedNotifications:", value)
}

func (v_ View) Bounds() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "bounds")
	return rv
}

func (v_ View) SetBounds(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "setBounds:", value)
}

func (v_ View) BoundsRotation() float64 {
	rv := ffi.CallMethod[float64](v_, "boundsRotation")
	return rv
}

func (v_ View) SetBoundsRotation(value float64) {
	ffi.CallMethod[ffi.Void](v_, "setBoundsRotation:", value)
}

func (v_ View) PostsBoundsChangedNotifications() bool {
	rv := ffi.CallMethod[bool](v_, "postsBoundsChangedNotifications")
	return rv
}

func (v_ View) SetPostsBoundsChangedNotifications(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setPostsBoundsChangedNotifications:", value)
}

func (v_ View) IsFlipped() bool {
	rv := ffi.CallMethod[bool](v_, "isFlipped")
	return rv
}

func (v_ View) IsRotatedFromBase() bool {
	rv := ffi.CallMethod[bool](v_, "isRotatedFromBase")
	return rv
}

func (v_ View) IsRotatedOrScaledFromBase() bool {
	rv := ffi.CallMethod[bool](v_, "isRotatedOrScaledFromBase")
	return rv
}

func (v_ View) IsHidden() bool {
	rv := ffi.CallMethod[bool](v_, "isHidden")
	return rv
}

func (v_ View) SetHidden(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setHidden:", value)
}

func (v_ View) IsHiddenOrHasHiddenAncestor() bool {
	rv := ffi.CallMethod[bool](v_, "isHiddenOrHasHiddenAncestor")
	return rv
}

func (v_ View) AllowsVibrancy() bool {
	rv := ffi.CallMethod[bool](v_, "allowsVibrancy")
	return rv
}

func (v_ View) FocusRingType() FocusRingType {
	rv := ffi.CallMethod[FocusRingType](v_, "focusRingType")
	return rv
}

func (v_ View) SetFocusRingType(value FocusRingType) {
	ffi.CallMethod[ffi.Void](v_, "setFocusRingType:", value)
}

func (v_ View) FocusRingMaskBounds() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "focusRingMaskBounds")
	return rv
}

func (vc _ViewClass) DefaultFocusRingType() FocusRingType {
	rv := ffi.CallMethod[FocusRingType](vc, "defaultFocusRingType")
	return rv
}

func (v_ View) IsDrawingFindIndicator() bool {
	rv := ffi.CallMethod[bool](v_, "isDrawingFindIndicator")
	return rv
}

func (v_ View) WantsLayer() bool {
	rv := ffi.CallMethod[bool](v_, "wantsLayer")
	return rv
}

func (v_ View) SetWantsLayer(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setWantsLayer:", value)
}

func (v_ View) WantsUpdateLayer() bool {
	rv := ffi.CallMethod[bool](v_, "wantsUpdateLayer")
	return rv
}

func (v_ View) Layer() quartzcore.Layer {
	rv := ffi.CallMethod[quartzcore.Layer](v_, "layer")
	return rv
}

func (v_ View) SetLayer(value quartzcore.ILayer) {
	ffi.CallMethod[ffi.Void](v_, "setLayer:", value)
}

func (v_ View) LayerContentsPlacement() ViewLayerContentsPlacement {
	rv := ffi.CallMethod[ViewLayerContentsPlacement](v_, "layerContentsPlacement")
	return rv
}

func (v_ View) SetLayerContentsPlacement(value ViewLayerContentsPlacement) {
	ffi.CallMethod[ffi.Void](v_, "setLayerContentsPlacement:", value)
}

func (v_ View) LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy {
	rv := ffi.CallMethod[ViewLayerContentsRedrawPolicy](v_, "layerContentsRedrawPolicy")
	return rv
}

func (v_ View) SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy) {
	ffi.CallMethod[ffi.Void](v_, "setLayerContentsRedrawPolicy:", value)
}

func (v_ View) CanDrawSubviewsIntoLayer() bool {
	rv := ffi.CallMethod[bool](v_, "canDrawSubviewsIntoLayer")
	return rv
}

func (v_ View) SetCanDrawSubviewsIntoLayer(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setCanDrawSubviewsIntoLayer:", value)
}

func (v_ View) LayerUsesCoreImageFilters() bool {
	rv := ffi.CallMethod[bool](v_, "layerUsesCoreImageFilters")
	return rv
}

func (v_ View) SetLayerUsesCoreImageFilters(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setLayerUsesCoreImageFilters:", value)
}

func (v_ View) AlphaValue() float64 {
	rv := ffi.CallMethod[float64](v_, "alphaValue")
	return rv
}

func (v_ View) SetAlphaValue(value float64) {
	ffi.CallMethod[ffi.Void](v_, "setAlphaValue:", value)
}

func (v_ View) FrameCenterRotation() float64 {
	rv := ffi.CallMethod[float64](v_, "frameCenterRotation")
	return rv
}

func (v_ View) SetFrameCenterRotation(value float64) {
	ffi.CallMethod[ffi.Void](v_, "setFrameCenterRotation:", value)
}

func (v_ View) Shadow() Shadow {
	rv := ffi.CallMethod[Shadow](v_, "shadow")
	return rv
}

func (v_ View) SetShadow(value IShadow) {
	ffi.CallMethod[ffi.Void](v_, "setShadow:", value)
}

func (vc _ViewClass) DefaultMenu() Menu {
	rv := ffi.CallMethod[Menu](vc, "defaultMenu")
	return rv
}

func (v_ View) ToolTip() string {
	rv := ffi.CallMethod[string](v_, "toolTip")
	return rv
}

func (v_ View) SetToolTip(value string) {
	ffi.CallMethod[ffi.Void](v_, "setToolTip:", value)
}

func (vc _ViewClass) FocusView() View {
	rv := ffi.CallMethod[View](vc, "focusView")
	return rv
}

func (v_ View) SafeAreaRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "safeAreaRect")
	return rv
}

func (v_ View) SafeAreaInsets() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](v_, "safeAreaInsets")
	return rv
}

func (v_ View) AdditionalSafeAreaInsets() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](v_, "additionalSafeAreaInsets")
	return rv
}

func (v_ View) SetAdditionalSafeAreaInsets(value foundation.EdgeInsets) {
	ffi.CallMethod[ffi.Void](v_, "setAdditionalSafeAreaInsets:", value)
}

func (v_ View) SafeAreaLayoutGuide() LayoutGuide {
	rv := ffi.CallMethod[LayoutGuide](v_, "safeAreaLayoutGuide")
	return rv
}

func (v_ View) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := ffi.CallMethod[UserInterfaceLayoutDirection](v_, "userInterfaceLayoutDirection")
	return rv
}

func (v_ View) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	ffi.CallMethod[ffi.Void](v_, "setUserInterfaceLayoutDirection:", value)
}

func (vc _ViewClass) RequiresConstraintBasedLayout() bool {
	rv := ffi.CallMethod[bool](vc, "requiresConstraintBasedLayout")
	return rv
}

func (v_ View) TranslatesAutoresizingMaskIntoConstraints() bool {
	rv := ffi.CallMethod[bool](v_, "translatesAutoresizingMaskIntoConstraints")
	return rv
}

func (v_ View) SetTranslatesAutoresizingMaskIntoConstraints(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setTranslatesAutoresizingMaskIntoConstraints:", value)
}

func (v_ View) BottomAnchor() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](v_, "bottomAnchor")
	return rv
}

func (v_ View) CenterXAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](v_, "centerXAnchor")
	return rv
}

func (v_ View) CenterYAnchor() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](v_, "centerYAnchor")
	return rv
}

func (v_ View) FirstBaselineAnchor() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](v_, "firstBaselineAnchor")
	return rv
}

func (v_ View) HeightAnchor() LayoutDimension {
	rv := ffi.CallMethod[LayoutDimension](v_, "heightAnchor")
	return rv
}

func (v_ View) LastBaselineAnchor() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](v_, "lastBaselineAnchor")
	return rv
}

func (v_ View) LeadingAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](v_, "leadingAnchor")
	return rv
}

func (v_ View) LeftAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](v_, "leftAnchor")
	return rv
}

func (v_ View) RightAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](v_, "rightAnchor")
	return rv
}

func (v_ View) TopAnchor() LayoutYAxisAnchor {
	rv := ffi.CallMethod[LayoutYAxisAnchor](v_, "topAnchor")
	return rv
}

func (v_ View) TrailingAnchor() LayoutXAxisAnchor {
	rv := ffi.CallMethod[LayoutXAxisAnchor](v_, "trailingAnchor")
	return rv
}

func (v_ View) WidthAnchor() LayoutDimension {
	rv := ffi.CallMethod[LayoutDimension](v_, "widthAnchor")
	return rv
}

func (v_ View) Constraints() []LayoutConstraint {
	rv := ffi.CallMethod[[]LayoutConstraint](v_, "constraints")
	return rv
}

func (v_ View) FittingSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "fittingSize")
	return rv
}

func (v_ View) IntrinsicContentSize() foundation.Size {
	rv := ffi.CallMethod[foundation.Size](v_, "intrinsicContentSize")
	return rv
}

func (v_ View) LayoutGuides() []LayoutGuide {
	rv := ffi.CallMethod[[]LayoutGuide](v_, "layoutGuides")
	return rv
}

func (v_ View) LayoutMarginsGuide() LayoutGuide {
	rv := ffi.CallMethod[LayoutGuide](v_, "layoutMarginsGuide")
	return rv
}

func (v_ View) AlignmentRectInsets() foundation.EdgeInsets {
	rv := ffi.CallMethod[foundation.EdgeInsets](v_, "alignmentRectInsets")
	return rv
}

func (v_ View) BaselineOffsetFromBottom() float64 {
	rv := ffi.CallMethod[float64](v_, "baselineOffsetFromBottom")
	return rv
}

func (v_ View) FirstBaselineOffsetFromTop() float64 {
	rv := ffi.CallMethod[float64](v_, "firstBaselineOffsetFromTop")
	return rv
}

func (v_ View) LastBaselineOffsetFromBottom() float64 {
	rv := ffi.CallMethod[float64](v_, "lastBaselineOffsetFromBottom")
	return rv
}

func (v_ View) NeedsLayout() bool {
	rv := ffi.CallMethod[bool](v_, "needsLayout")
	return rv
}

func (v_ View) SetNeedsLayout(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setNeedsLayout:", value)
}

func (v_ View) NeedsUpdateConstraints() bool {
	rv := ffi.CallMethod[bool](v_, "needsUpdateConstraints")
	return rv
}

func (v_ View) SetNeedsUpdateConstraints(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setNeedsUpdateConstraints:", value)
}

func (v_ View) IsHorizontalContentSizeConstraintActive() bool {
	rv := ffi.CallMethod[bool](v_, "isHorizontalContentSizeConstraintActive")
	return rv
}

func (v_ View) SetHorizontalContentSizeConstraintActive(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setHorizontalContentSizeConstraintActive:", value)
}

func (v_ View) IsVerticalContentSizeConstraintActive() bool {
	rv := ffi.CallMethod[bool](v_, "isVerticalContentSizeConstraintActive")
	return rv
}

func (v_ View) SetVerticalContentSizeConstraintActive(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setVerticalContentSizeConstraintActive:", value)
}

func (v_ View) HasAmbiguousLayout() bool {
	rv := ffi.CallMethod[bool](v_, "hasAmbiguousLayout")
	return rv
}

func (v_ View) AutoresizesSubviews() bool {
	rv := ffi.CallMethod[bool](v_, "autoresizesSubviews")
	return rv
}

func (v_ View) SetAutoresizesSubviews(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setAutoresizesSubviews:", value)
}

func (v_ View) AutoresizingMask() AutoresizingMaskOptions {
	rv := ffi.CallMethod[AutoresizingMaskOptions](v_, "autoresizingMask")
	return rv
}

func (v_ View) SetAutoresizingMask(value AutoresizingMaskOptions) {
	ffi.CallMethod[ffi.Void](v_, "setAutoresizingMask:", value)
}

func (v_ View) CanDrawConcurrently() bool {
	rv := ffi.CallMethod[bool](v_, "canDrawConcurrently")
	return rv
}

func (v_ View) SetCanDrawConcurrently(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setCanDrawConcurrently:", value)
}

func (v_ View) VisibleRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "visibleRect")
	return rv
}

func (v_ View) WantsDefaultClipping() bool {
	rv := ffi.CallMethod[bool](v_, "wantsDefaultClipping")
	return rv
}

func (v_ View) IsInFullScreenMode() bool {
	rv := ffi.CallMethod[bool](v_, "isInFullScreenMode")
	return rv
}

func (v_ View) NeedsDisplay() bool {
	rv := ffi.CallMethod[bool](v_, "needsDisplay")
	return rv
}

func (v_ View) SetNeedsDisplay(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setNeedsDisplay:", value)
}

func (v_ View) IsOpaque() bool {
	rv := ffi.CallMethod[bool](v_, "isOpaque")
	return rv
}

func (v_ View) InLiveResize() bool {
	rv := ffi.CallMethod[bool](v_, "inLiveResize")
	return rv
}

func (v_ View) PreservesContentDuringLiveResize() bool {
	rv := ffi.CallMethod[bool](v_, "preservesContentDuringLiveResize")
	return rv
}

func (v_ View) RectPreservedDuringLiveResize() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "rectPreservedDuringLiveResize")
	return rv
}

func (v_ View) PrintJobTitle() string {
	rv := ffi.CallMethod[string](v_, "printJobTitle")
	return rv
}

func (v_ View) PageHeader() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](v_, "pageHeader")
	return rv
}

func (v_ View) PageFooter() foundation.AttributedString {
	rv := ffi.CallMethod[foundation.AttributedString](v_, "pageFooter")
	return rv
}

func (v_ View) HeightAdjustLimit() float64 {
	rv := ffi.CallMethod[float64](v_, "heightAdjustLimit")
	return rv
}

func (v_ View) WidthAdjustLimit() float64 {
	rv := ffi.CallMethod[float64](v_, "widthAdjustLimit")
	return rv
}

func (v_ View) MouseDownCanMoveWindow() bool {
	rv := ffi.CallMethod[bool](v_, "mouseDownCanMoveWindow")
	return rv
}

func (v_ View) InputContext() TextInputContext {
	rv := ffi.CallMethod[TextInputContext](v_, "inputContext")
	return rv
}

func (v_ View) AllowedTouchTypes() TouchTypeMask {
	rv := ffi.CallMethod[TouchTypeMask](v_, "allowedTouchTypes")
	return rv
}

func (v_ View) SetAllowedTouchTypes(value TouchTypeMask) {
	ffi.CallMethod[ffi.Void](v_, "setAllowedTouchTypes:", value)
}

func (v_ View) WantsRestingTouches() bool {
	rv := ffi.CallMethod[bool](v_, "wantsRestingTouches")
	return rv
}

func (v_ View) SetWantsRestingTouches(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setWantsRestingTouches:", value)
}

func (v_ View) CandidateListTouchBarItem() CandidateListTouchBarItem {
	rv := ffi.CallMethod[CandidateListTouchBarItem](v_, "candidateListTouchBarItem")
	return rv
}

func (v_ View) GestureRecognizers() []GestureRecognizer {
	rv := ffi.CallMethod[[]GestureRecognizer](v_, "gestureRecognizers")
	return rv
}

func (v_ View) SetGestureRecognizers(value []IGestureRecognizer) {
	ffi.CallMethod[ffi.Void](v_, "setGestureRecognizers:", value)
}

func (v_ View) CanBecomeKeyView() bool {
	rv := ffi.CallMethod[bool](v_, "canBecomeKeyView")
	return rv
}

func (v_ View) NeedsPanelToBecomeKey() bool {
	rv := ffi.CallMethod[bool](v_, "needsPanelToBecomeKey")
	return rv
}

func (v_ View) NextKeyView() View {
	rv := ffi.CallMethod[View](v_, "nextKeyView")
	return rv
}

func (v_ View) SetNextKeyView(value IView) {
	ffi.CallMethod[ffi.Void](v_, "setNextKeyView:", value)
}

func (v_ View) NextValidKeyView() View {
	rv := ffi.CallMethod[View](v_, "nextValidKeyView")
	return rv
}

func (v_ View) PreviousKeyView() View {
	rv := ffi.CallMethod[View](v_, "previousKeyView")
	return rv
}

func (v_ View) PreviousValidKeyView() View {
	rv := ffi.CallMethod[View](v_, "previousValidKeyView")
	return rv
}

func (v_ View) TrackingAreas() []TrackingArea {
	rv := ffi.CallMethod[[]TrackingArea](v_, "trackingAreas")
	return rv
}

func (v_ View) PreparedContentRect() foundation.Rect {
	rv := ffi.CallMethod[foundation.Rect](v_, "preparedContentRect")
	return rv
}

func (v_ View) SetPreparedContentRect(value foundation.Rect) {
	ffi.CallMethod[ffi.Void](v_, "setPreparedContentRect:", value)
}

func (vc _ViewClass) IsCompatibleWithResponsiveScrolling() bool {
	rv := ffi.CallMethod[bool](vc, "isCompatibleWithResponsiveScrolling")
	return rv
}

func (v_ View) PressureConfiguration() PressureConfiguration {
	rv := ffi.CallMethod[PressureConfiguration](v_, "pressureConfiguration")
	return rv
}

func (v_ View) SetPressureConfiguration(value IPressureConfiguration) {
	ffi.CallMethod[ffi.Void](v_, "setPressureConfiguration:", value)
}

func (v_ View) RegisteredDraggedTypes() []PasteboardType {
	rv := ffi.CallMethod[[]PasteboardType](v_, "registeredDraggedTypes")
	return rv
}

// deprecated
func (v_ View) WantsExtendedDynamicRangeOpenGLSurface() bool {
	rv := ffi.CallMethod[bool](v_, "wantsExtendedDynamicRangeOpenGLSurface")
	return rv
}

// deprecated
func (v_ View) SetWantsExtendedDynamicRangeOpenGLSurface(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setWantsExtendedDynamicRangeOpenGLSurface:", value)
}

// deprecated
func (v_ View) AcceptsTouchEvents() bool {
	rv := ffi.CallMethod[bool](v_, "acceptsTouchEvents")
	return rv
}

// deprecated
func (v_ View) SetAcceptsTouchEvents(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setAcceptsTouchEvents:", value)
}

// deprecated
func (v_ View) CanDraw() bool {
	rv := ffi.CallMethod[bool](v_, "canDraw")
	return rv
}

// deprecated
func (v_ View) WantsBestResolutionOpenGLSurface() bool {
	rv := ffi.CallMethod[bool](v_, "wantsBestResolutionOpenGLSurface")
	return rv
}

// deprecated
func (v_ View) SetWantsBestResolutionOpenGLSurface(value bool) {
	ffi.CallMethod[ffi.Void](v_, "setWantsBestResolutionOpenGLSurface:", value)
}
