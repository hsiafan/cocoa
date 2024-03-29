// auto-generated code, do not modify
package appkit

import (
	"unsafe"

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
	BeginDraggingSessionWithItems_Event_Source(items []IDraggingItem, event IEvent, source objc.IObject) DraggingSession
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
	rv := objc.CallMethod[View](v_, objc.SEL("initWithFrame:"), frameRect)
	return rv
}

func (v_ View) Init() View {
	rv := objc.CallMethod[View](v_, objc.SEL("init"))
	return rv
}

func (vc _ViewClass) Alloc() View {
	rv := objc.CallMethod[View](vc, objc.SEL("alloc"))
	return rv
}

func (vc _ViewClass) New() View {
	rv := objc.CallMethod[View](vc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewView() View {
	return ViewClass.New()
}

func (v_ View) PrepareForReuse() {
	objc.CallMethod[objc.Void](v_, objc.SEL("prepareForReuse"))
}

func (v_ View) IsDescendantOf(view IView) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isDescendantOf:"), objc.ExtractPtr(view))
	return rv
}

func (v_ View) AncestorSharedWithView(view IView) View {
	rv := objc.CallMethod[View](v_, objc.SEL("ancestorSharedWithView:"), objc.ExtractPtr(view))
	return rv
}

func (v_ View) AddSubview(view IView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("addSubview:"), objc.ExtractPtr(view))
}

func (v_ View) AddSubview_Positioned_RelativeTo(view IView, place WindowOrderingMode, otherView IView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("addSubview:positioned:relativeTo:"), objc.ExtractPtr(view), place, objc.ExtractPtr(otherView))
}

func (v_ View) RemoveFromSuperview() {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeFromSuperview"))
}

func (v_ View) RemoveFromSuperviewWithoutNeedingDisplay() {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeFromSuperviewWithoutNeedingDisplay"))
}

func (v_ View) ReplaceSubview_With(oldView IView, newView IView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("replaceSubview:with:"), objc.ExtractPtr(oldView), objc.ExtractPtr(newView))
}

func (v_ View) DidAddSubview(subview IView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("didAddSubview:"), objc.ExtractPtr(subview))
}

func (v_ View) ViewDidMoveToSuperview() {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewDidMoveToSuperview"))
}

func (v_ View) ViewDidMoveToWindow() {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewDidMoveToWindow"))
}

func (v_ View) ViewWillMoveToSuperview(newSuperview IView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewWillMoveToSuperview:"), objc.ExtractPtr(newSuperview))
}

func (v_ View) ViewWillMoveToWindow(newWindow IWindow) {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewWillMoveToWindow:"), objc.ExtractPtr(newWindow))
}

func (v_ View) WillRemoveSubview(subview IView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("willRemoveSubview:"), objc.ExtractPtr(subview))
}

func (v_ View) ViewWithTag(tag int) View {
	rv := objc.CallMethod[View](v_, objc.SEL("viewWithTag:"), tag)
	return rv
}

func (v_ View) SetFrameOrigin(newOrigin foundation.Point) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setFrameOrigin:"), newOrigin)
}

func (v_ View) SetFrameSize(newSize foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setFrameSize:"), newSize)
}

func (v_ View) SetBoundsOrigin(newOrigin foundation.Point) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setBoundsOrigin:"), newOrigin)
}

func (v_ View) SetBoundsSize(newSize foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setBoundsSize:"), newSize)
}

func (v_ View) TranslateOriginToPoint(translation foundation.Point) {
	objc.CallMethod[objc.Void](v_, objc.SEL("translateOriginToPoint:"), translation)
}

func (v_ View) ScaleUnitSquareToSize(newUnitSize foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.SEL("scaleUnitSquareToSize:"), newUnitSize)
}

func (v_ View) RotateByAngle(angle float64) {
	objc.CallMethod[objc.Void](v_, objc.SEL("rotateByAngle:"), angle)
}

func (v_ View) BackingAlignedRect_Options(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("backingAlignedRect:options:"), rect, options)
	return rv
}

func (v_ View) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("convertPointFromBacking:"), point)
	return rv
}

func (v_ View) ConvertPointToBacking(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("convertPointToBacking:"), point)
	return rv
}

func (v_ View) ConvertPointFromLayer(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("convertPointFromLayer:"), point)
	return rv
}

func (v_ View) ConvertPointToLayer(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("convertPointToLayer:"), point)
	return rv
}

func (v_ View) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("convertRectFromBacking:"), rect)
	return rv
}

func (v_ View) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("convertRectToBacking:"), rect)
	return rv
}

func (v_ View) ConvertRectFromLayer(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("convertRectFromLayer:"), rect)
	return rv
}

func (v_ View) ConvertRectToLayer(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("convertRectToLayer:"), rect)
	return rv
}

func (v_ View) ConvertSizeFromBacking(size foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("convertSizeFromBacking:"), size)
	return rv
}

func (v_ View) ConvertSizeToBacking(size foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("convertSizeToBacking:"), size)
	return rv
}

func (v_ View) ConvertSizeFromLayer(size foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("convertSizeFromLayer:"), size)
	return rv
}

func (v_ View) ConvertSizeToLayer(size foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("convertSizeToLayer:"), size)
	return rv
}

func (v_ View) ConvertPoint_FromView(point foundation.Point, view IView) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("convertPoint:fromView:"), point, objc.ExtractPtr(view))
	return rv
}

func (v_ View) ConvertPoint_ToView(point foundation.Point, view IView) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("convertPoint:toView:"), point, objc.ExtractPtr(view))
	return rv
}

func (v_ View) ConvertSize_FromView(size foundation.Size, view IView) foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("convertSize:fromView:"), size, objc.ExtractPtr(view))
	return rv
}

func (v_ View) ConvertSize_ToView(size foundation.Size, view IView) foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("convertSize:toView:"), size, objc.ExtractPtr(view))
	return rv
}

func (v_ View) ConvertRect_FromView(rect foundation.Rect, view IView) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("convertRect:fromView:"), rect, objc.ExtractPtr(view))
	return rv
}

func (v_ View) ConvertRect_ToView(rect foundation.Rect, view IView) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("convertRect:toView:"), rect, objc.ExtractPtr(view))
	return rv
}

func (v_ View) CenterScanRect(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("centerScanRect:"), rect)
	return rv
}

func (v_ View) ViewDidHide() {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewDidHide"))
}

func (v_ View) ViewDidUnhide() {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewDidUnhide"))
}

func (v_ View) ViewDidChangeEffectiveAppearance() {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewDidChangeEffectiveAppearance"))
}

func (v_ View) ViewDidChangeBackingProperties() {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewDidChangeBackingProperties"))
}

func (v_ View) DrawFocusRingMask() {
	objc.CallMethod[objc.Void](v_, objc.SEL("drawFocusRingMask"))
}

func (v_ View) NoteFocusRingMaskChanged() {
	objc.CallMethod[objc.Void](v_, objc.SEL("noteFocusRingMaskChanged"))
}

func (v_ View) SetKeyboardFocusRingNeedsDisplayInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setKeyboardFocusRingNeedsDisplayInRect:"), rect)
}

func (v_ View) MakeBackingLayer() quartzcore.Layer {
	rv := objc.CallMethod[quartzcore.Layer](v_, objc.SEL("makeBackingLayer"))
	return rv
}

func (v_ View) MenuForEvent(event IEvent) Menu {
	rv := objc.CallMethod[Menu](v_, objc.SEL("menuForEvent:"), objc.ExtractPtr(event))
	return rv
}

func (v_ View) WillOpenMenu_WithEvent(menu IMenu, event IEvent) {
	objc.CallMethod[objc.Void](v_, objc.SEL("willOpenMenu:withEvent:"), objc.ExtractPtr(menu), objc.ExtractPtr(event))
}

func (v_ View) DidCloseMenu_WithEvent(menu IMenu, event IEvent) {
	objc.CallMethod[objc.Void](v_, objc.SEL("didCloseMenu:withEvent:"), objc.ExtractPtr(menu), objc.ExtractPtr(event))
}

func (v_ View) AddCursorRect_Cursor(rect foundation.Rect, object ICursor) {
	objc.CallMethod[objc.Void](v_, objc.SEL("addCursorRect:cursor:"), rect, objc.ExtractPtr(object))
}

func (v_ View) RemoveCursorRect_Cursor(rect foundation.Rect, object ICursor) {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeCursorRect:cursor:"), rect, objc.ExtractPtr(object))
}

func (v_ View) DiscardCursorRects() {
	objc.CallMethod[objc.Void](v_, objc.SEL("discardCursorRects"))
}

func (v_ View) ResetCursorRects() {
	objc.CallMethod[objc.Void](v_, objc.SEL("resetCursorRects"))
}

func (v_ View) AddToolTipRect_Owner_UserData(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer) ToolTipTag {
	rv := objc.CallMethod[ToolTipTag](v_, objc.SEL("addToolTipRect:owner:userData:"), rect, objc.ExtractPtr(owner), data)
	return rv
}

func (v_ View) RemoveAllToolTips() {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeAllToolTips"))
}

func (v_ View) RemoveToolTip(tag ToolTipTag) {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeToolTip:"), tag)
}

func (v_ View) ShowDefinitionForAttributedString_AtPoint(attrString foundation.IAttributedString, textBaselineOrigin foundation.Point) {
	objc.CallMethod[objc.Void](v_, objc.SEL("showDefinitionForAttributedString:atPoint:"), objc.ExtractPtr(attrString), textBaselineOrigin)
}

func (v_ View) ShowDefinitionForAttributedString_Range_Options_BaselineOriginProvider(attrString foundation.IAttributedString, targetRange foundation.Range, options map[DefinitionOptionKey]objc.IObject, originProvider func(adjustedRange foundation.Range) foundation.Point) {
	objc.CallMethod[objc.Void](v_, objc.SEL("showDefinitionForAttributedString:range:options:baselineOriginProvider:"), objc.ExtractPtr(attrString), targetRange, options, originProvider)
}

func (v_ View) RulerView_DidAddMarker(ruler IRulerView, marker IRulerMarker) {
	objc.CallMethod[objc.Void](v_, objc.SEL("rulerView:didAddMarker:"), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
}

func (v_ View) RulerView_DidMoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.CallMethod[objc.Void](v_, objc.SEL("rulerView:didMoveMarker:"), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
}

func (v_ View) RulerView_DidRemoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.CallMethod[objc.Void](v_, objc.SEL("rulerView:didRemoveMarker:"), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
}

func (v_ View) RulerView_HandleMouseDown(ruler IRulerView, event IEvent) {
	objc.CallMethod[objc.Void](v_, objc.SEL("rulerView:handleMouseDown:"), objc.ExtractPtr(ruler), objc.ExtractPtr(event))
}

func (v_ View) RulerView_LocationForPoint(ruler IRulerView, point foundation.Point) float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("rulerView:locationForPoint:"), objc.ExtractPtr(ruler), point)
	return rv
}

func (v_ View) RulerView_PointForLocation(ruler IRulerView, point float64) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("rulerView:pointForLocation:"), objc.ExtractPtr(ruler), point)
	return rv
}

func (v_ View) RulerView_ShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("rulerView:shouldAddMarker:"), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
	return rv
}

func (v_ View) RulerView_ShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("rulerView:shouldMoveMarker:"), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
	return rv
}

func (v_ View) RulerView_ShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("rulerView:shouldRemoveMarker:"), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
	return rv
}

func (v_ View) RulerView_WillAddMarker_AtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("rulerView:willAddMarker:atLocation:"), objc.ExtractPtr(ruler), objc.ExtractPtr(marker), location)
	return rv
}

func (v_ View) RulerView_WillMoveMarker_ToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("rulerView:willMoveMarker:toLocation:"), objc.ExtractPtr(ruler), objc.ExtractPtr(marker), location)
	return rv
}

func (v_ View) RulerView_WillSetClientView(ruler IRulerView, newClient IView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("rulerView:willSetClientView:"), objc.ExtractPtr(ruler), objc.ExtractPtr(newClient))
}

func (v_ View) AddConstraint(constraint ILayoutConstraint) {
	objc.CallMethod[objc.Void](v_, objc.SEL("addConstraint:"), objc.ExtractPtr(constraint))
}

func (v_ View) AddConstraints(constraints []ILayoutConstraint) {
	objc.CallMethod[objc.Void](v_, objc.SEL("addConstraints:"), constraints)
}

func (v_ View) RemoveConstraint(constraint ILayoutConstraint) {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeConstraint:"), objc.ExtractPtr(constraint))
}

func (v_ View) RemoveConstraints(constraints []ILayoutConstraint) {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeConstraints:"), constraints)
}

func (v_ View) InvalidateIntrinsicContentSize() {
	objc.CallMethod[objc.Void](v_, objc.SEL("invalidateIntrinsicContentSize"))
}

func (v_ View) ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](v_, objc.SEL("contentCompressionResistancePriorityForOrientation:"), orientation)
	return rv
}

func (v_ View) SetContentCompressionResistancePriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setContentCompressionResistancePriority:forOrientation:"), priority, orientation)
}

func (v_ View) ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.CallMethod[LayoutPriority](v_, objc.SEL("contentHuggingPriorityForOrientation:"), orientation)
	return rv
}

func (v_ View) SetContentHuggingPriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setContentHuggingPriority:forOrientation:"), priority, orientation)
}

func (v_ View) AddLayoutGuide(guide ILayoutGuide) {
	objc.CallMethod[objc.Void](v_, objc.SEL("addLayoutGuide:"), objc.ExtractPtr(guide))
}

func (v_ View) RemoveLayoutGuide(guide ILayoutGuide) {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeLayoutGuide:"), objc.ExtractPtr(guide))
}

func (v_ View) AlignmentRectForFrame(frame foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("alignmentRectForFrame:"), frame)
	return rv
}

func (v_ View) FrameForAlignmentRect(alignmentRect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("frameForAlignmentRect:"), alignmentRect)
	return rv
}

func (v_ View) Layout() {
	objc.CallMethod[objc.Void](v_, objc.SEL("layout"))
}

func (v_ View) LayoutSubtreeIfNeeded() {
	objc.CallMethod[objc.Void](v_, objc.SEL("layoutSubtreeIfNeeded"))
}

func (v_ View) UpdateConstraints() {
	objc.CallMethod[objc.Void](v_, objc.SEL("updateConstraints"))
}

func (v_ View) UpdateConstraintsForSubtreeIfNeeded() {
	objc.CallMethod[objc.Void](v_, objc.SEL("updateConstraintsForSubtreeIfNeeded"))
}

func (v_ View) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](v_, objc.SEL("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}

func (v_ View) ExerciseAmbiguityInLayout() {
	objc.CallMethod[objc.Void](v_, objc.SEL("exerciseAmbiguityInLayout"))
}

func (v_ View) ResizeSubviewsWithOldSize(oldSize foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.SEL("resizeSubviewsWithOldSize:"), oldSize)
}

func (v_ View) ResizeWithOldSuperviewSize(oldSize foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.SEL("resizeWithOldSuperviewSize:"), oldSize)
}

func (v_ View) UpdateLayer() {
	objc.CallMethod[objc.Void](v_, objc.SEL("updateLayer"))
}

func (v_ View) DrawRect(dirtyRect foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("drawRect:"), dirtyRect)
}

func (v_ View) NeedsToDrawRect(rect foundation.Rect) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("needsToDrawRect:"), rect)
	return rv
}

func (v_ View) BitmapImageRepForCachingDisplayInRect(rect foundation.Rect) BitmapImageRep {
	rv := objc.CallMethod[BitmapImageRep](v_, objc.SEL("bitmapImageRepForCachingDisplayInRect:"), rect)
	return rv
}

func (v_ View) CacheDisplayInRect_ToBitmapImageRep(rect foundation.Rect, bitmapImageRep IBitmapImageRep) {
	objc.CallMethod[objc.Void](v_, objc.SEL("cacheDisplayInRect:toBitmapImageRep:"), rect, objc.ExtractPtr(bitmapImageRep))
}

func (v_ View) EnterFullScreenMode_WithOptions(screen IScreen, options map[ViewFullScreenModeOptionKey]objc.IObject) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("enterFullScreenMode:withOptions:"), objc.ExtractPtr(screen), options)
	return rv
}

func (v_ View) ExitFullScreenModeWithOptions(options map[ViewFullScreenModeOptionKey]objc.IObject) {
	objc.CallMethod[objc.Void](v_, objc.SEL("exitFullScreenModeWithOptions:"), options)
}

func (v_ View) SetNeedsDisplayInRect(invalidRect foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setNeedsDisplayInRect:"), invalidRect)
}

func (v_ View) Display() {
	objc.CallMethod[objc.Void](v_, objc.SEL("display"))
}

func (v_ View) DisplayRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("displayRect:"), rect)
}

func (v_ View) DisplayRectIgnoringOpacity(rect foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("displayRectIgnoringOpacity:"), rect)
}

func (v_ View) DisplayRectIgnoringOpacity_InContext(rect foundation.Rect, context IGraphicsContext) {
	objc.CallMethod[objc.Void](v_, objc.SEL("displayRectIgnoringOpacity:inContext:"), rect, objc.ExtractPtr(context))
}

func (v_ View) DisplayIfNeeded() {
	objc.CallMethod[objc.Void](v_, objc.SEL("displayIfNeeded"))
}

func (v_ View) DisplayIfNeededInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("displayIfNeededInRect:"), rect)
}

func (v_ View) DisplayIfNeededIgnoringOpacity() {
	objc.CallMethod[objc.Void](v_, objc.SEL("displayIfNeededIgnoringOpacity"))
}

func (v_ View) DisplayIfNeededInRectIgnoringOpacity(rect foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("displayIfNeededInRectIgnoringOpacity:"), rect)
}

func (v_ View) TranslateRectsNeedingDisplayInRect_By(clipRect foundation.Rect, delta foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.SEL("translateRectsNeedingDisplayInRect:by:"), clipRect, delta)
}

func (v_ View) ViewWillDraw() {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewWillDraw"))
}

func (v_ View) GetRectsExposedDuringLiveResize_Count(exposedRects *foundation.Rect, count *int) {
	objc.CallMethod[objc.Void](v_, objc.SEL("getRectsExposedDuringLiveResize:count:"), exposedRects, count)
}

func (v_ View) ViewWillStartLiveResize() {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewWillStartLiveResize"))
}

func (v_ View) ViewDidEndLiveResize() {
	objc.CallMethod[objc.Void](v_, objc.SEL("viewDidEndLiveResize"))
}

func (v_ View) Print(sender objc.IObject) {
	objc.CallMethod[objc.Void](v_, objc.SEL("print:"), objc.ExtractPtr(sender))
}

func (v_ View) BeginPageInRect_AtPlacement(rect foundation.Rect, location foundation.Point) {
	objc.CallMethod[objc.Void](v_, objc.SEL("beginPageInRect:atPlacement:"), rect, location)
}

func (v_ View) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	rv := objc.CallMethod[[]byte](v_, objc.SEL("dataWithEPSInsideRect:"), rect)
	return rv
}

func (v_ View) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	rv := objc.CallMethod[[]byte](v_, objc.SEL("dataWithPDFInsideRect:"), rect)
	return rv
}

func (v_ View) WriteEPSInsideRect_ToPasteboard(rect foundation.Rect, pasteboard IPasteboard) {
	objc.CallMethod[objc.Void](v_, objc.SEL("writeEPSInsideRect:toPasteboard:"), rect, objc.ExtractPtr(pasteboard))
}

func (v_ View) WritePDFInsideRect_ToPasteboard(rect foundation.Rect, pasteboard IPasteboard) {
	objc.CallMethod[objc.Void](v_, objc.SEL("writePDFInsideRect:toPasteboard:"), rect, objc.ExtractPtr(pasteboard))
}

func (v_ View) DrawPageBorderWithSize(borderSize foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.SEL("drawPageBorderWithSize:"), borderSize)
}

// deprecated
func (v_ View) DrawSheetBorderWithSize(borderSize foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.SEL("drawSheetBorderWithSize:"), borderSize)
}

func (v_ View) AdjustPageWidthNew_Left_Right_Limit(newRight *float64, oldLeft float64, oldRight float64, rightLimit float64) {
	objc.CallMethod[objc.Void](v_, objc.SEL("adjustPageWidthNew:left:right:limit:"), newRight, oldLeft, oldRight, rightLimit)
}

func (v_ View) AdjustPageHeightNew_Top_Bottom_Limit(newBottom *float64, oldTop float64, oldBottom float64, bottomLimit float64) {
	objc.CallMethod[objc.Void](v_, objc.SEL("adjustPageHeightNew:top:bottom:limit:"), newBottom, oldTop, oldBottom, bottomLimit)
}

func (v_ View) KnowsPageRange(range_ *foundation.Range) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("knowsPageRange:"), range_)
	return rv
}

func (v_ View) RectForPage(page int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("rectForPage:"), page)
	return rv
}

func (v_ View) LocationOfPrintRect(rect foundation.Rect) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("locationOfPrintRect:"), rect)
	return rv
}

func (v_ View) BeginDocument() {
	objc.CallMethod[objc.Void](v_, objc.SEL("beginDocument"))
}

func (v_ View) EndDocument() {
	objc.CallMethod[objc.Void](v_, objc.SEL("endDocument"))
}

func (v_ View) EndPage() {
	objc.CallMethod[objc.Void](v_, objc.SEL("endPage"))
}

func (v_ View) AcceptsFirstMouse(event IEvent) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("acceptsFirstMouse:"), objc.ExtractPtr(event))
	return rv
}

func (v_ View) HitTest(point foundation.Point) View {
	rv := objc.CallMethod[View](v_, objc.SEL("hitTest:"), point)
	return rv
}

func (v_ View) Mouse_InRect(point foundation.Point, rect foundation.Rect) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("mouse:inRect:"), point, rect)
	return rv
}

func (v_ View) AddGestureRecognizer(gestureRecognizer IGestureRecognizer) {
	objc.CallMethod[objc.Void](v_, objc.SEL("addGestureRecognizer:"), objc.ExtractPtr(gestureRecognizer))
}

func (v_ View) RemoveGestureRecognizer(gestureRecognizer IGestureRecognizer) {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeGestureRecognizer:"), objc.ExtractPtr(gestureRecognizer))
}

func (v_ View) RectForSmartMagnificationAtPoint_InRect(location foundation.Point, visibleRect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("rectForSmartMagnificationAtPoint:inRect:"), location, visibleRect)
	return rv
}

func (v_ View) AddTrackingArea(trackingArea ITrackingArea) {
	objc.CallMethod[objc.Void](v_, objc.SEL("addTrackingArea:"), objc.ExtractPtr(trackingArea))
}

func (v_ View) RemoveTrackingArea(trackingArea ITrackingArea) {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeTrackingArea:"), objc.ExtractPtr(trackingArea))
}

func (v_ View) UpdateTrackingAreas() {
	objc.CallMethod[objc.Void](v_, objc.SEL("updateTrackingAreas"))
}

func (v_ View) AddTrackingRect_Owner_UserData_AssumeInside(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer, flag bool) TrackingRectTag {
	rv := objc.CallMethod[TrackingRectTag](v_, objc.SEL("addTrackingRect:owner:userData:assumeInside:"), rect, objc.ExtractPtr(owner), data, flag)
	return rv
}

func (v_ View) RemoveTrackingRect(tag TrackingRectTag) {
	objc.CallMethod[objc.Void](v_, objc.SEL("removeTrackingRect:"), tag)
}

func (v_ View) PrepareContentInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("prepareContentInRect:"), rect)
}

func (v_ View) ScrollPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](v_, objc.SEL("scrollPoint:"), point)
}

func (v_ View) ScrollRectToVisible(rect foundation.Rect) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("scrollRectToVisible:"), rect)
	return rv
}

func (v_ View) Autoscroll(event IEvent) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("autoscroll:"), objc.ExtractPtr(event))
	return rv
}

func (v_ View) AdjustScroll(newVisible foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("adjustScroll:"), newVisible)
	return rv
}

func (v_ View) ScrollClipView_ToPoint(clipView IClipView, point foundation.Point) {
	objc.CallMethod[objc.Void](v_, objc.SEL("scrollClipView:toPoint:"), objc.ExtractPtr(clipView), point)
}

func (v_ View) ReflectScrolledClipView(clipView IClipView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("reflectScrolledClipView:"), objc.ExtractPtr(clipView))
}

func (v_ View) RegisterForDraggedTypes(newTypes []PasteboardType) {
	objc.CallMethod[objc.Void](v_, objc.SEL("registerForDraggedTypes:"), newTypes)
}

func (v_ View) UnregisterDraggedTypes() {
	objc.CallMethod[objc.Void](v_, objc.SEL("unregisterDraggedTypes"))
}

func (v_ View) BeginDraggingSessionWithItems_Event_Source(items []IDraggingItem, event IEvent, source objc.IObject) DraggingSession {
	rv := objc.CallMethod[DraggingSession](v_, objc.SEL("beginDraggingSessionWithItems:event:source:"), items, objc.ExtractPtr(event), objc.ExtractPtr(source))
	return rv
}

func (v_ View) ShouldDelayWindowOrderingForEvent(event IEvent) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("shouldDelayWindowOrderingForEvent:"), objc.ExtractPtr(event))
	return rv
}

// deprecated
func (v_ View) LockFocus() {
	objc.CallMethod[objc.Void](v_, objc.SEL("lockFocus"))
}

// deprecated
func (v_ View) LockFocusIfCanDraw() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("lockFocusIfCanDraw"))
	return rv
}

// deprecated
func (v_ View) LockFocusIfCanDrawInContext(context IGraphicsContext) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("lockFocusIfCanDrawInContext:"), objc.ExtractPtr(context))
	return rv
}

// deprecated
func (v_ View) UnlockFocus() {
	objc.CallMethod[objc.Void](v_, objc.SEL("unlockFocus"))
}

// deprecated
func (v_ View) ScrollRect_By(rect foundation.Rect, delta foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.SEL("scrollRect:by:"), rect, delta)
}

// deprecated
func (v_ View) ConvertPointToBase(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("convertPointToBase:"), point)
	return rv
}

// deprecated
func (v_ View) ConvertPointFromBase(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.SEL("convertPointFromBase:"), point)
	return rv
}

// deprecated
func (v_ View) ConvertSizeToBase(size foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("convertSizeToBase:"), size)
	return rv
}

// deprecated
func (v_ View) ConvertSizeFromBase(size foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("convertSizeFromBase:"), size)
	return rv
}

// deprecated
func (v_ View) ConvertRectToBase(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("convertRectToBase:"), rect)
	return rv
}

// deprecated
func (v_ View) ConvertRectFromBase(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("convertRectFromBase:"), rect)
	return rv
}

// deprecated
func (v_ View) ShouldDrawColor() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("shouldDrawColor"))
	return rv
}

// deprecated
func (v_ View) AllocateGState() {
	objc.CallMethod[objc.Void](v_, objc.SEL("allocateGState"))
}

// deprecated
func (v_ View) GState() int {
	rv := objc.CallMethod[int](v_, objc.SEL("gState"))
	return rv
}

// deprecated
func (v_ View) SetUpGState() {
	objc.CallMethod[objc.Void](v_, objc.SEL("setUpGState"))
}

// deprecated
func (v_ View) RenewGState() {
	objc.CallMethod[objc.Void](v_, objc.SEL("renewGState"))
}

// deprecated
func (v_ View) DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(image IImage, viewLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("dragImage:at:offset:event:pasteboard:source:slideBack:"), objc.ExtractPtr(image), viewLocation, initialOffset, objc.ExtractPtr(event), objc.ExtractPtr(pboard), objc.ExtractPtr(sourceObj), slideFlag)
}

// deprecated
func (v_ View) DragFile_FromRect_SlideBack_Event(filename string, rect foundation.Rect, flag bool, event IEvent) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("dragFile:fromRect:slideBack:event:"), filename, rect, flag, objc.ExtractPtr(event))
	return rv
}

// deprecated
func (v_ View) DragPromisedFilesOfTypes_FromRect_Source_SlideBack_Event(typeArray []string, rect foundation.Rect, sourceObject objc.IObject, flag bool, event IEvent) bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("dragPromisedFilesOfTypes:fromRect:source:slideBack:event:"), typeArray, rect, objc.ExtractPtr(sourceObject), flag, objc.ExtractPtr(event))
	return rv
}

// weak property
func (v_ View) Superview() View {
	rv := objc.CallMethod[View](v_, objc.SEL("superview"))
	return rv
}

func (v_ View) Subviews() []View {
	rv := objc.CallMethod[[]View](v_, objc.SEL("subviews"))
	return rv
}

func (v_ View) SetSubviews(value []IView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setSubviews:"), value)
}

// weak property
func (v_ View) Window() Window {
	rv := objc.CallMethod[Window](v_, objc.SEL("window"))
	return rv
}

// weak property
func (v_ View) OpaqueAncestor() View {
	rv := objc.CallMethod[View](v_, objc.SEL("opaqueAncestor"))
	return rv
}

func (v_ View) EnclosingMenuItem() MenuItem {
	rv := objc.CallMethod[MenuItem](v_, objc.SEL("enclosingMenuItem"))
	return rv
}

func (v_ View) EnclosingScrollView() ScrollView {
	rv := objc.CallMethod[ScrollView](v_, objc.SEL("enclosingScrollView"))
	return rv
}

func (v_ View) Tag() int {
	rv := objc.CallMethod[int](v_, objc.SEL("tag"))
	return rv
}

func (v_ View) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("frame"))
	return rv
}

func (v_ View) SetFrame(value foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setFrame:"), value)
}

func (v_ View) FrameRotation() float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("frameRotation"))
	return rv
}

func (v_ View) SetFrameRotation(value float64) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setFrameRotation:"), value)
}

func (v_ View) PostsFrameChangedNotifications() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("postsFrameChangedNotifications"))
	return rv
}

func (v_ View) SetPostsFrameChangedNotifications(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setPostsFrameChangedNotifications:"), value)
}

func (v_ View) Bounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("bounds"))
	return rv
}

func (v_ View) SetBounds(value foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setBounds:"), value)
}

func (v_ View) BoundsRotation() float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("boundsRotation"))
	return rv
}

func (v_ View) SetBoundsRotation(value float64) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setBoundsRotation:"), value)
}

func (v_ View) PostsBoundsChangedNotifications() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("postsBoundsChangedNotifications"))
	return rv
}

func (v_ View) SetPostsBoundsChangedNotifications(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setPostsBoundsChangedNotifications:"), value)
}

func (v_ View) IsFlipped() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isFlipped"))
	return rv
}

func (v_ View) IsRotatedFromBase() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isRotatedFromBase"))
	return rv
}

func (v_ View) IsRotatedOrScaledFromBase() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isRotatedOrScaledFromBase"))
	return rv
}

func (v_ View) IsHidden() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isHidden"))
	return rv
}

func (v_ View) SetHidden(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setHidden:"), value)
}

func (v_ View) IsHiddenOrHasHiddenAncestor() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isHiddenOrHasHiddenAncestor"))
	return rv
}

func (v_ View) AllowsVibrancy() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("allowsVibrancy"))
	return rv
}

func (v_ View) FocusRingType() FocusRingType {
	rv := objc.CallMethod[FocusRingType](v_, objc.SEL("focusRingType"))
	return rv
}

func (v_ View) SetFocusRingType(value FocusRingType) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setFocusRingType:"), value)
}

func (v_ View) FocusRingMaskBounds() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("focusRingMaskBounds"))
	return rv
}

func (vc _ViewClass) DefaultFocusRingType() FocusRingType {
	rv := objc.CallMethod[FocusRingType](vc, objc.SEL("defaultFocusRingType"))
	return rv
}

func (v_ View) IsDrawingFindIndicator() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isDrawingFindIndicator"))
	return rv
}

func (v_ View) WantsLayer() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("wantsLayer"))
	return rv
}

func (v_ View) SetWantsLayer(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setWantsLayer:"), value)
}

func (v_ View) WantsUpdateLayer() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("wantsUpdateLayer"))
	return rv
}

func (v_ View) Layer() quartzcore.Layer {
	rv := objc.CallMethod[quartzcore.Layer](v_, objc.SEL("layer"))
	return rv
}

func (v_ View) SetLayer(value quartzcore.ILayer) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setLayer:"), objc.ExtractPtr(value))
}

func (v_ View) LayerContentsPlacement() ViewLayerContentsPlacement {
	rv := objc.CallMethod[ViewLayerContentsPlacement](v_, objc.SEL("layerContentsPlacement"))
	return rv
}

func (v_ View) SetLayerContentsPlacement(value ViewLayerContentsPlacement) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setLayerContentsPlacement:"), value)
}

func (v_ View) LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy {
	rv := objc.CallMethod[ViewLayerContentsRedrawPolicy](v_, objc.SEL("layerContentsRedrawPolicy"))
	return rv
}

func (v_ View) SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setLayerContentsRedrawPolicy:"), value)
}

func (v_ View) CanDrawSubviewsIntoLayer() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("canDrawSubviewsIntoLayer"))
	return rv
}

func (v_ View) SetCanDrawSubviewsIntoLayer(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setCanDrawSubviewsIntoLayer:"), value)
}

func (v_ View) LayerUsesCoreImageFilters() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("layerUsesCoreImageFilters"))
	return rv
}

func (v_ View) SetLayerUsesCoreImageFilters(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setLayerUsesCoreImageFilters:"), value)
}

func (v_ View) AlphaValue() float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("alphaValue"))
	return rv
}

func (v_ View) SetAlphaValue(value float64) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setAlphaValue:"), value)
}

func (v_ View) FrameCenterRotation() float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("frameCenterRotation"))
	return rv
}

func (v_ View) SetFrameCenterRotation(value float64) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setFrameCenterRotation:"), value)
}

func (v_ View) Shadow() Shadow {
	rv := objc.CallMethod[Shadow](v_, objc.SEL("shadow"))
	return rv
}

func (v_ View) SetShadow(value IShadow) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setShadow:"), objc.ExtractPtr(value))
}

func (vc _ViewClass) DefaultMenu() Menu {
	rv := objc.CallMethod[Menu](vc, objc.SEL("defaultMenu"))
	return rv
}

func (v_ View) ToolTip() string {
	rv := objc.CallMethod[string](v_, objc.SEL("toolTip"))
	return rv
}

func (v_ View) SetToolTip(value string) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setToolTip:"), value)
}

func (vc _ViewClass) FocusView() View {
	rv := objc.CallMethod[View](vc, objc.SEL("focusView"))
	return rv
}

func (v_ View) SafeAreaRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("safeAreaRect"))
	return rv
}

func (v_ View) SafeAreaInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](v_, objc.SEL("safeAreaInsets"))
	return rv
}

func (v_ View) AdditionalSafeAreaInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](v_, objc.SEL("additionalSafeAreaInsets"))
	return rv
}

func (v_ View) SetAdditionalSafeAreaInsets(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setAdditionalSafeAreaInsets:"), value)
}

func (v_ View) SafeAreaLayoutGuide() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](v_, objc.SEL("safeAreaLayoutGuide"))
	return rv
}

func (v_ View) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](v_, objc.SEL("userInterfaceLayoutDirection"))
	return rv
}

func (v_ View) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setUserInterfaceLayoutDirection:"), value)
}

func (vc _ViewClass) RequiresConstraintBasedLayout() bool {
	rv := objc.CallMethod[bool](vc, objc.SEL("requiresConstraintBasedLayout"))
	return rv
}

func (v_ View) TranslatesAutoresizingMaskIntoConstraints() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("translatesAutoresizingMaskIntoConstraints"))
	return rv
}

func (v_ View) SetTranslatesAutoresizingMaskIntoConstraints(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setTranslatesAutoresizingMaskIntoConstraints:"), value)
}

func (v_ View) BottomAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](v_, objc.SEL("bottomAnchor"))
	return rv
}

func (v_ View) CenterXAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](v_, objc.SEL("centerXAnchor"))
	return rv
}

func (v_ View) CenterYAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](v_, objc.SEL("centerYAnchor"))
	return rv
}

func (v_ View) FirstBaselineAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](v_, objc.SEL("firstBaselineAnchor"))
	return rv
}

func (v_ View) HeightAnchor() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](v_, objc.SEL("heightAnchor"))
	return rv
}

func (v_ View) LastBaselineAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](v_, objc.SEL("lastBaselineAnchor"))
	return rv
}

func (v_ View) LeadingAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](v_, objc.SEL("leadingAnchor"))
	return rv
}

func (v_ View) LeftAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](v_, objc.SEL("leftAnchor"))
	return rv
}

func (v_ View) RightAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](v_, objc.SEL("rightAnchor"))
	return rv
}

func (v_ View) TopAnchor() LayoutYAxisAnchor {
	rv := objc.CallMethod[LayoutYAxisAnchor](v_, objc.SEL("topAnchor"))
	return rv
}

func (v_ View) TrailingAnchor() LayoutXAxisAnchor {
	rv := objc.CallMethod[LayoutXAxisAnchor](v_, objc.SEL("trailingAnchor"))
	return rv
}

func (v_ View) WidthAnchor() LayoutDimension {
	rv := objc.CallMethod[LayoutDimension](v_, objc.SEL("widthAnchor"))
	return rv
}

func (v_ View) Constraints() []LayoutConstraint {
	rv := objc.CallMethod[[]LayoutConstraint](v_, objc.SEL("constraints"))
	return rv
}

func (v_ View) FittingSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("fittingSize"))
	return rv
}

func (v_ View) IntrinsicContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.SEL("intrinsicContentSize"))
	return rv
}

func (v_ View) LayoutGuides() []LayoutGuide {
	rv := objc.CallMethod[[]LayoutGuide](v_, objc.SEL("layoutGuides"))
	return rv
}

func (v_ View) LayoutMarginsGuide() LayoutGuide {
	rv := objc.CallMethod[LayoutGuide](v_, objc.SEL("layoutMarginsGuide"))
	return rv
}

func (v_ View) AlignmentRectInsets() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](v_, objc.SEL("alignmentRectInsets"))
	return rv
}

func (v_ View) BaselineOffsetFromBottom() float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("baselineOffsetFromBottom"))
	return rv
}

func (v_ View) FirstBaselineOffsetFromTop() float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("firstBaselineOffsetFromTop"))
	return rv
}

func (v_ View) LastBaselineOffsetFromBottom() float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("lastBaselineOffsetFromBottom"))
	return rv
}

func (v_ View) NeedsLayout() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("needsLayout"))
	return rv
}

func (v_ View) SetNeedsLayout(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setNeedsLayout:"), value)
}

func (v_ View) NeedsUpdateConstraints() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("needsUpdateConstraints"))
	return rv
}

func (v_ View) SetNeedsUpdateConstraints(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setNeedsUpdateConstraints:"), value)
}

func (v_ View) IsHorizontalContentSizeConstraintActive() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isHorizontalContentSizeConstraintActive"))
	return rv
}

func (v_ View) SetHorizontalContentSizeConstraintActive(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setHorizontalContentSizeConstraintActive:"), value)
}

func (v_ View) IsVerticalContentSizeConstraintActive() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isVerticalContentSizeConstraintActive"))
	return rv
}

func (v_ View) SetVerticalContentSizeConstraintActive(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setVerticalContentSizeConstraintActive:"), value)
}

func (v_ View) HasAmbiguousLayout() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("hasAmbiguousLayout"))
	return rv
}

func (v_ View) AutoresizesSubviews() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("autoresizesSubviews"))
	return rv
}

func (v_ View) SetAutoresizesSubviews(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setAutoresizesSubviews:"), value)
}

func (v_ View) AutoresizingMask() AutoresizingMaskOptions {
	rv := objc.CallMethod[AutoresizingMaskOptions](v_, objc.SEL("autoresizingMask"))
	return rv
}

func (v_ View) SetAutoresizingMask(value AutoresizingMaskOptions) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setAutoresizingMask:"), value)
}

func (v_ View) CanDrawConcurrently() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("canDrawConcurrently"))
	return rv
}

func (v_ View) SetCanDrawConcurrently(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setCanDrawConcurrently:"), value)
}

func (v_ View) VisibleRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("visibleRect"))
	return rv
}

func (v_ View) WantsDefaultClipping() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("wantsDefaultClipping"))
	return rv
}

func (v_ View) IsInFullScreenMode() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isInFullScreenMode"))
	return rv
}

func (v_ View) NeedsDisplay() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("needsDisplay"))
	return rv
}

func (v_ View) SetNeedsDisplay(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setNeedsDisplay:"), value)
}

func (v_ View) IsOpaque() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("isOpaque"))
	return rv
}

func (v_ View) InLiveResize() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("inLiveResize"))
	return rv
}

func (v_ View) PreservesContentDuringLiveResize() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("preservesContentDuringLiveResize"))
	return rv
}

func (v_ View) RectPreservedDuringLiveResize() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("rectPreservedDuringLiveResize"))
	return rv
}

func (v_ View) PrintJobTitle() string {
	rv := objc.CallMethod[string](v_, objc.SEL("printJobTitle"))
	return rv
}

func (v_ View) PageHeader() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](v_, objc.SEL("pageHeader"))
	return rv
}

func (v_ View) PageFooter() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](v_, objc.SEL("pageFooter"))
	return rv
}

func (v_ View) HeightAdjustLimit() float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("heightAdjustLimit"))
	return rv
}

func (v_ View) WidthAdjustLimit() float64 {
	rv := objc.CallMethod[float64](v_, objc.SEL("widthAdjustLimit"))
	return rv
}

func (v_ View) MouseDownCanMoveWindow() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("mouseDownCanMoveWindow"))
	return rv
}

func (v_ View) InputContext() TextInputContext {
	rv := objc.CallMethod[TextInputContext](v_, objc.SEL("inputContext"))
	return rv
}

func (v_ View) AllowedTouchTypes() TouchTypeMask {
	rv := objc.CallMethod[TouchTypeMask](v_, objc.SEL("allowedTouchTypes"))
	return rv
}

func (v_ View) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setAllowedTouchTypes:"), value)
}

func (v_ View) WantsRestingTouches() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("wantsRestingTouches"))
	return rv
}

func (v_ View) SetWantsRestingTouches(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setWantsRestingTouches:"), value)
}

func (v_ View) CandidateListTouchBarItem() CandidateListTouchBarItem {
	rv := objc.CallMethod[CandidateListTouchBarItem](v_, objc.SEL("candidateListTouchBarItem"))
	return rv
}

func (v_ View) GestureRecognizers() []GestureRecognizer {
	rv := objc.CallMethod[[]GestureRecognizer](v_, objc.SEL("gestureRecognizers"))
	return rv
}

func (v_ View) SetGestureRecognizers(value []IGestureRecognizer) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setGestureRecognizers:"), value)
}

func (v_ View) CanBecomeKeyView() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("canBecomeKeyView"))
	return rv
}

func (v_ View) NeedsPanelToBecomeKey() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("needsPanelToBecomeKey"))
	return rv
}

// weak property
func (v_ View) NextKeyView() View {
	rv := objc.CallMethod[View](v_, objc.SEL("nextKeyView"))
	return rv
}

// weak property
func (v_ View) SetNextKeyView(value IView) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setNextKeyView:"), objc.ExtractPtr(value))
}

// weak property
func (v_ View) NextValidKeyView() View {
	rv := objc.CallMethod[View](v_, objc.SEL("nextValidKeyView"))
	return rv
}

// weak property
func (v_ View) PreviousKeyView() View {
	rv := objc.CallMethod[View](v_, objc.SEL("previousKeyView"))
	return rv
}

// weak property
func (v_ View) PreviousValidKeyView() View {
	rv := objc.CallMethod[View](v_, objc.SEL("previousValidKeyView"))
	return rv
}

func (v_ View) TrackingAreas() []TrackingArea {
	rv := objc.CallMethod[[]TrackingArea](v_, objc.SEL("trackingAreas"))
	return rv
}

func (v_ View) PreparedContentRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](v_, objc.SEL("preparedContentRect"))
	return rv
}

func (v_ View) SetPreparedContentRect(value foundation.Rect) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setPreparedContentRect:"), value)
}

func (vc _ViewClass) IsCompatibleWithResponsiveScrolling() bool {
	rv := objc.CallMethod[bool](vc, objc.SEL("isCompatibleWithResponsiveScrolling"))
	return rv
}

func (v_ View) PressureConfiguration() PressureConfiguration {
	rv := objc.CallMethod[PressureConfiguration](v_, objc.SEL("pressureConfiguration"))
	return rv
}

func (v_ View) SetPressureConfiguration(value IPressureConfiguration) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setPressureConfiguration:"), objc.ExtractPtr(value))
}

func (v_ View) RegisteredDraggedTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](v_, objc.SEL("registeredDraggedTypes"))
	return rv
}

// deprecated
func (v_ View) WantsExtendedDynamicRangeOpenGLSurface() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("wantsExtendedDynamicRangeOpenGLSurface"))
	return rv
}

// deprecated
func (v_ View) SetWantsExtendedDynamicRangeOpenGLSurface(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setWantsExtendedDynamicRangeOpenGLSurface:"), value)
}

// deprecated
func (v_ View) AcceptsTouchEvents() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("acceptsTouchEvents"))
	return rv
}

// deprecated
func (v_ View) SetAcceptsTouchEvents(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setAcceptsTouchEvents:"), value)
}

// deprecated
func (v_ View) CanDraw() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("canDraw"))
	return rv
}

// deprecated
func (v_ View) WantsBestResolutionOpenGLSurface() bool {
	rv := objc.CallMethod[bool](v_, objc.SEL("wantsBestResolutionOpenGLSurface"))
	return rv
}

// deprecated
func (v_ View) SetWantsBestResolutionOpenGLSurface(value bool) {
	objc.CallMethod[objc.Void](v_, objc.SEL("setWantsBestResolutionOpenGLSurface:"), value)
}
