// auto-generated code, do not modify
package quartzcore

import (
	"unsafe"

	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var LayerClass = _LayerClass{objc.GetClass("CALayer")}

type _LayerClass struct {
	objc.Class
}

type ILayer interface {
	objc.IObject
	Display()
	DrawInContext(ctx coregraphics.ContextRef)
	ContentsAreFlipped() bool
	RenderInContext(ctx coregraphics.ContextRef)
	AffineTransform() coregraphics.AffineTransform
	SetAffineTransform(m coregraphics.AffineTransform)
	AddSublayer(layer ILayer)
	RemoveFromSuperlayer()
	InsertSublayer_AtIndex(layer ILayer, idx uint32)
	InsertSublayer_Below(layer ILayer, sibling ILayer)
	InsertSublayer_Above(layer ILayer, sibling ILayer)
	ReplaceSublayer_With(oldLayer ILayer, newLayer ILayer)
	SetNeedsDisplay()
	SetNeedsDisplayInRect(r coregraphics.Rect)
	DisplayIfNeeded()
	NeedsDisplay() bool
	AddAnimation_ForKey(anim IAnimation, key string)
	AnimationForKey(key string) Animation
	RemoveAllAnimations()
	RemoveAnimationForKey(key string)
	AnimationKeys() []string
	SetNeedsLayout()
	LayoutSublayers()
	LayoutIfNeeded()
	NeedsLayout() bool
	ResizeWithOldSuperlayerSize(size coregraphics.Size)
	ResizeSublayersWithOldSize(size coregraphics.Size)
	PreferredFrameSize() coregraphics.Size
	AddConstraint(c IConstraint)
	ActionForKey(event string) ActionWrapper
	ConvertPoint_FromLayer(p coregraphics.Point, l ILayer) coregraphics.Point
	ConvertPoint_ToLayer(p coregraphics.Point, l ILayer) coregraphics.Point
	ConvertRect_FromLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect
	ConvertRect_ToLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect
	HitTest(p coregraphics.Point) Layer
	ContainsPoint(p coregraphics.Point) bool
	ScrollPoint(p coregraphics.Point)
	ScrollRectToVisible(r coregraphics.Rect)
	ShouldArchiveValueForKey(key string) bool
	Delegate() LayerDelegateWrapper
	SetDelegate(value LayerDelegate)
	SetDelegate0(value objc.IObject)
	Contents() objc.Object
	SetContents(value objc.IObject)
	ContentsRect() coregraphics.Rect
	SetContentsRect(value coregraphics.Rect)
	ContentsCenter() coregraphics.Rect
	SetContentsCenter(value coregraphics.Rect)
	ContentsGravity() LayerContentsGravity
	SetContentsGravity(value LayerContentsGravity)
	Opacity() float32
	SetOpacity(value float32)
	IsHidden() bool
	SetHidden(value bool)
	MasksToBounds() bool
	SetMasksToBounds(value bool)
	Mask() Layer
	SetMask(value ILayer)
	IsDoubleSided() bool
	SetDoubleSided(value bool)
	CornerRadius() float64
	SetCornerRadius(value float64)
	MaskedCorners() CornerMask
	SetMaskedCorners(value CornerMask)
	BorderWidth() float64
	SetBorderWidth(value float64)
	BorderColor() coregraphics.ColorRef
	SetBorderColor(value coregraphics.ColorRef)
	BackgroundColor() coregraphics.ColorRef
	SetBackgroundColor(value coregraphics.ColorRef)
	ShadowOpacity() float32
	SetShadowOpacity(value float32)
	ShadowRadius() float64
	SetShadowRadius(value float64)
	ShadowOffset() coregraphics.Size
	SetShadowOffset(value coregraphics.Size)
	ShadowColor() coregraphics.ColorRef
	SetShadowColor(value coregraphics.ColorRef)
	ShadowPath() coregraphics.PathRef
	SetShadowPath(value coregraphics.PathRef)
	AllowsEdgeAntialiasing() bool
	SetAllowsEdgeAntialiasing(value bool)
	AllowsGroupOpacity() bool
	SetAllowsGroupOpacity(value bool)
	Filters() []objc.Object
	SetFilters(value []objc.IObject)
	CompositingFilter() objc.Object
	SetCompositingFilter(value objc.IObject)
	BackgroundFilters() []objc.Object
	SetBackgroundFilters(value []objc.IObject)
	MinificationFilter() LayerContentsFilter
	SetMinificationFilter(value LayerContentsFilter)
	MinificationFilterBias() float32
	SetMinificationFilterBias(value float32)
	MagnificationFilter() LayerContentsFilter
	SetMagnificationFilter(value LayerContentsFilter)
	IsOpaque() bool
	SetOpaque(value bool)
	EdgeAntialiasingMask() EdgeAntialiasingMask
	SetEdgeAntialiasingMask(value EdgeAntialiasingMask)
	IsGeometryFlipped() bool
	SetGeometryFlipped(value bool)
	DrawsAsynchronously() bool
	SetDrawsAsynchronously(value bool)
	ShouldRasterize() bool
	SetShouldRasterize(value bool)
	RasterizationScale() float64
	SetRasterizationScale(value float64)
	ContentsFormat() LayerContentsFormat
	SetContentsFormat(value LayerContentsFormat)
	Frame() coregraphics.Rect
	SetFrame(value coregraphics.Rect)
	Bounds() coregraphics.Rect
	SetBounds(value coregraphics.Rect)
	Position() coregraphics.Point
	SetPosition(value coregraphics.Point)
	ZPosition() float64
	SetZPosition(value float64)
	AnchorPointZ() float64
	SetAnchorPointZ(value float64)
	AnchorPoint() coregraphics.Point
	SetAnchorPoint(value coregraphics.Point)
	ContentsScale() float64
	SetContentsScale(value float64)
	Transform() Transform3D
	SetTransform(value Transform3D)
	SublayerTransform() Transform3D
	SetSublayerTransform(value Transform3D)
	Sublayers() []Layer
	SetSublayers(value []ILayer)
	Superlayer() Layer
	NeedsDisplayOnBoundsChange() bool
	SetNeedsDisplayOnBoundsChange(value bool)
	LayoutManager() LayoutManagerWrapper
	SetLayoutManager(value LayoutManager)
	SetLayoutManager0(value objc.IObject)
	AutoresizingMask() AutoresizingMask
	SetAutoresizingMask(value AutoresizingMask)
	Constraints() []Constraint
	SetConstraints(value []IConstraint)
	Actions() foundation.Dictionary
	SetActions(value foundation.IDictionary)
	VisibleRect() coregraphics.Rect
	Name() string
	SetName(value string)
	CornerCurve() LayerCornerCurve
	SetCornerCurve(value LayerCornerCurve)
}

type Layer struct {
	objc.Object
}

func MakeLayer(ptr unsafe.Pointer) Layer {
	return Layer{
		Object: objc.MakeObject(ptr),
	}
}

func (lc _LayerClass) Layer() Layer {
	rv := objc.CallMethod[Layer](lc, "layer")
	return rv
}

func (l_ Layer) Init() Layer {
	rv := objc.CallMethod[Layer](l_, "init")
	return rv
}

func (l_ Layer) InitWithLayer(layer objc.IObject) Layer {
	rv := objc.CallMethod[Layer](l_, "initWithLayer:", layer)
	return rv
}

func (l_ Layer) PresentationLayer() Layer {
	rv := objc.CallMethod[Layer](l_, "presentationLayer")
	return rv
}

func (l_ Layer) ModelLayer() Layer {
	rv := objc.CallMethod[Layer](l_, "modelLayer")
	return rv
}

func (lc _LayerClass) Alloc() Layer {
	rv := objc.CallMethod[Layer](lc, "alloc")
	return rv
}

func (lc _LayerClass) New() Layer {
	rv := objc.CallMethod[Layer](lc, "new")
	rv.Autorelease()
	return rv
}

func NewLayer() Layer {
	return LayerClass.New()
}

func (lc _LayerClass) LayerWithRemoteClientId(client_id uint32) Layer {
	rv := objc.CallMethod[Layer](lc, "layerWithRemoteClientId:", client_id)
	return rv
}

func (l_ Layer) Display() {
	objc.CallMethod[objc.Void](l_, "display")
}

func (l_ Layer) DrawInContext(ctx coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](l_, "drawInContext:", ctx)
}

func (l_ Layer) ContentsAreFlipped() bool {
	rv := objc.CallMethod[bool](l_, "contentsAreFlipped")
	return rv
}

func (l_ Layer) RenderInContext(ctx coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](l_, "renderInContext:", ctx)
}

func (l_ Layer) AffineTransform() coregraphics.AffineTransform {
	rv := objc.CallMethod[coregraphics.AffineTransform](l_, "affineTransform")
	return rv
}

func (l_ Layer) SetAffineTransform(m coregraphics.AffineTransform) {
	objc.CallMethod[objc.Void](l_, "setAffineTransform:", m)
}

func (l_ Layer) AddSublayer(layer ILayer) {
	objc.CallMethod[objc.Void](l_, "addSublayer:", layer)
}

func (l_ Layer) RemoveFromSuperlayer() {
	objc.CallMethod[objc.Void](l_, "removeFromSuperlayer")
}

func (l_ Layer) InsertSublayer_AtIndex(layer ILayer, idx uint32) {
	objc.CallMethod[objc.Void](l_, "insertSublayer:atIndex:", layer, idx)
}

func (l_ Layer) InsertSublayer_Below(layer ILayer, sibling ILayer) {
	objc.CallMethod[objc.Void](l_, "insertSublayer:below:", layer, sibling)
}

func (l_ Layer) InsertSublayer_Above(layer ILayer, sibling ILayer) {
	objc.CallMethod[objc.Void](l_, "insertSublayer:above:", layer, sibling)
}

func (l_ Layer) ReplaceSublayer_With(oldLayer ILayer, newLayer ILayer) {
	objc.CallMethod[objc.Void](l_, "replaceSublayer:with:", oldLayer, newLayer)
}

func (l_ Layer) SetNeedsDisplay() {
	objc.CallMethod[objc.Void](l_, "setNeedsDisplay")
}

func (l_ Layer) SetNeedsDisplayInRect(r coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, "setNeedsDisplayInRect:", r)
}

func (l_ Layer) DisplayIfNeeded() {
	objc.CallMethod[objc.Void](l_, "displayIfNeeded")
}

func (l_ Layer) NeedsDisplay() bool {
	rv := objc.CallMethod[bool](l_, "needsDisplay")
	return rv
}

func (lc _LayerClass) NeedsDisplayForKey(key string) bool {
	rv := objc.CallMethod[bool](lc, "needsDisplayForKey:", key)
	return rv
}

func (l_ Layer) AddAnimation_ForKey(anim IAnimation, key string) {
	objc.CallMethod[objc.Void](l_, "addAnimation:forKey:", anim, key)
}

func (l_ Layer) AnimationForKey(key string) Animation {
	rv := objc.CallMethod[Animation](l_, "animationForKey:", key)
	return rv
}

func (l_ Layer) RemoveAllAnimations() {
	objc.CallMethod[objc.Void](l_, "removeAllAnimations")
}

func (l_ Layer) RemoveAnimationForKey(key string) {
	objc.CallMethod[objc.Void](l_, "removeAnimationForKey:", key)
}

func (l_ Layer) AnimationKeys() []string {
	rv := objc.CallMethod[[]string](l_, "animationKeys")
	return rv
}

func (l_ Layer) SetNeedsLayout() {
	objc.CallMethod[objc.Void](l_, "setNeedsLayout")
}

func (l_ Layer) LayoutSublayers() {
	objc.CallMethod[objc.Void](l_, "layoutSublayers")
}

func (l_ Layer) LayoutIfNeeded() {
	objc.CallMethod[objc.Void](l_, "layoutIfNeeded")
}

func (l_ Layer) NeedsLayout() bool {
	rv := objc.CallMethod[bool](l_, "needsLayout")
	return rv
}

func (l_ Layer) ResizeWithOldSuperlayerSize(size coregraphics.Size) {
	objc.CallMethod[objc.Void](l_, "resizeWithOldSuperlayerSize:", size)
}

func (l_ Layer) ResizeSublayersWithOldSize(size coregraphics.Size) {
	objc.CallMethod[objc.Void](l_, "resizeSublayersWithOldSize:", size)
}

func (l_ Layer) PreferredFrameSize() coregraphics.Size {
	rv := objc.CallMethod[coregraphics.Size](l_, "preferredFrameSize")
	return rv
}

func (l_ Layer) AddConstraint(c IConstraint) {
	objc.CallMethod[objc.Void](l_, "addConstraint:", c)
}

func (l_ Layer) ActionForKey(event string) ActionWrapper {
	rv := objc.CallMethod[ActionWrapper](l_, "actionForKey:", event)
	return rv
}

func (lc _LayerClass) DefaultActionForKey(event string) ActionWrapper {
	rv := objc.CallMethod[ActionWrapper](lc, "defaultActionForKey:", event)
	return rv
}

func (l_ Layer) ConvertPoint_FromLayer(p coregraphics.Point, l ILayer) coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](l_, "convertPoint:fromLayer:", p, l)
	return rv
}

func (l_ Layer) ConvertPoint_ToLayer(p coregraphics.Point, l ILayer) coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](l_, "convertPoint:toLayer:", p, l)
	return rv
}

func (l_ Layer) ConvertRect_FromLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, "convertRect:fromLayer:", r, l)
	return rv
}

func (l_ Layer) ConvertRect_ToLayer(r coregraphics.Rect, l ILayer) coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, "convertRect:toLayer:", r, l)
	return rv
}

func (l_ Layer) HitTest(p coregraphics.Point) Layer {
	rv := objc.CallMethod[Layer](l_, "hitTest:", p)
	return rv
}

func (l_ Layer) ContainsPoint(p coregraphics.Point) bool {
	rv := objc.CallMethod[bool](l_, "containsPoint:", p)
	return rv
}

func (l_ Layer) ScrollPoint(p coregraphics.Point) {
	objc.CallMethod[objc.Void](l_, "scrollPoint:", p)
}

func (l_ Layer) ScrollRectToVisible(r coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, "scrollRectToVisible:", r)
}

func (l_ Layer) ShouldArchiveValueForKey(key string) bool {
	rv := objc.CallMethod[bool](l_, "shouldArchiveValueForKey:", key)
	return rv
}

func (lc _LayerClass) DefaultValueForKey(key string) objc.Object {
	rv := objc.CallMethod[objc.Object](lc, "defaultValueForKey:", key)
	return rv
}

func (lc _LayerClass) CornerCurveExpansionFactor(curve LayerCornerCurve) float64 {
	rv := objc.CallMethod[float64](lc, "cornerCurveExpansionFactor:", curve)
	return rv
}

func (l_ Layer) Delegate() LayerDelegateWrapper {
	rv := objc.CallMethod[LayerDelegateWrapper](l_, "delegate")
	return rv
}

func (l_ Layer) SetDelegate(value LayerDelegate) {
	po := objc.CreateProtocol("CALayerDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(l_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](l_, "setDelegate:", po)
}

func (l_ Layer) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, "setDelegate:", value)
}

func (l_ Layer) Contents() objc.Object {
	rv := objc.CallMethod[objc.Object](l_, "contents")
	return rv
}

func (l_ Layer) SetContents(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, "setContents:", value)
}

func (l_ Layer) ContentsRect() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, "contentsRect")
	return rv
}

func (l_ Layer) SetContentsRect(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, "setContentsRect:", value)
}

func (l_ Layer) ContentsCenter() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, "contentsCenter")
	return rv
}

func (l_ Layer) SetContentsCenter(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, "setContentsCenter:", value)
}

func (l_ Layer) ContentsGravity() LayerContentsGravity {
	rv := objc.CallMethod[LayerContentsGravity](l_, "contentsGravity")
	return rv
}

func (l_ Layer) SetContentsGravity(value LayerContentsGravity) {
	objc.CallMethod[objc.Void](l_, "setContentsGravity:", value)
}

func (l_ Layer) Opacity() float32 {
	rv := objc.CallMethod[float32](l_, "opacity")
	return rv
}

func (l_ Layer) SetOpacity(value float32) {
	objc.CallMethod[objc.Void](l_, "setOpacity:", value)
}

func (l_ Layer) IsHidden() bool {
	rv := objc.CallMethod[bool](l_, "isHidden")
	return rv
}

func (l_ Layer) SetHidden(value bool) {
	objc.CallMethod[objc.Void](l_, "setHidden:", value)
}

func (l_ Layer) MasksToBounds() bool {
	rv := objc.CallMethod[bool](l_, "masksToBounds")
	return rv
}

func (l_ Layer) SetMasksToBounds(value bool) {
	objc.CallMethod[objc.Void](l_, "setMasksToBounds:", value)
}

func (l_ Layer) Mask() Layer {
	rv := objc.CallMethod[Layer](l_, "mask")
	return rv
}

func (l_ Layer) SetMask(value ILayer) {
	objc.CallMethod[objc.Void](l_, "setMask:", value)
}

func (l_ Layer) IsDoubleSided() bool {
	rv := objc.CallMethod[bool](l_, "isDoubleSided")
	return rv
}

func (l_ Layer) SetDoubleSided(value bool) {
	objc.CallMethod[objc.Void](l_, "setDoubleSided:", value)
}

func (l_ Layer) CornerRadius() float64 {
	rv := objc.CallMethod[float64](l_, "cornerRadius")
	return rv
}

func (l_ Layer) SetCornerRadius(value float64) {
	objc.CallMethod[objc.Void](l_, "setCornerRadius:", value)
}

func (l_ Layer) MaskedCorners() CornerMask {
	rv := objc.CallMethod[CornerMask](l_, "maskedCorners")
	return rv
}

func (l_ Layer) SetMaskedCorners(value CornerMask) {
	objc.CallMethod[objc.Void](l_, "setMaskedCorners:", value)
}

func (l_ Layer) BorderWidth() float64 {
	rv := objc.CallMethod[float64](l_, "borderWidth")
	return rv
}

func (l_ Layer) SetBorderWidth(value float64) {
	objc.CallMethod[objc.Void](l_, "setBorderWidth:", value)
}

func (l_ Layer) BorderColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](l_, "borderColor")
	return rv
}

func (l_ Layer) SetBorderColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](l_, "setBorderColor:", value)
}

func (l_ Layer) BackgroundColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](l_, "backgroundColor")
	return rv
}

func (l_ Layer) SetBackgroundColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](l_, "setBackgroundColor:", value)
}

func (l_ Layer) ShadowOpacity() float32 {
	rv := objc.CallMethod[float32](l_, "shadowOpacity")
	return rv
}

func (l_ Layer) SetShadowOpacity(value float32) {
	objc.CallMethod[objc.Void](l_, "setShadowOpacity:", value)
}

func (l_ Layer) ShadowRadius() float64 {
	rv := objc.CallMethod[float64](l_, "shadowRadius")
	return rv
}

func (l_ Layer) SetShadowRadius(value float64) {
	objc.CallMethod[objc.Void](l_, "setShadowRadius:", value)
}

func (l_ Layer) ShadowOffset() coregraphics.Size {
	rv := objc.CallMethod[coregraphics.Size](l_, "shadowOffset")
	return rv
}

func (l_ Layer) SetShadowOffset(value coregraphics.Size) {
	objc.CallMethod[objc.Void](l_, "setShadowOffset:", value)
}

func (l_ Layer) ShadowColor() coregraphics.ColorRef {
	rv := objc.CallMethod[coregraphics.ColorRef](l_, "shadowColor")
	return rv
}

func (l_ Layer) SetShadowColor(value coregraphics.ColorRef) {
	objc.CallMethod[objc.Void](l_, "setShadowColor:", value)
}

func (l_ Layer) ShadowPath() coregraphics.PathRef {
	rv := objc.CallMethod[coregraphics.PathRef](l_, "shadowPath")
	return rv
}

func (l_ Layer) SetShadowPath(value coregraphics.PathRef) {
	objc.CallMethod[objc.Void](l_, "setShadowPath:", value)
}

func (l_ Layer) AllowsEdgeAntialiasing() bool {
	rv := objc.CallMethod[bool](l_, "allowsEdgeAntialiasing")
	return rv
}

func (l_ Layer) SetAllowsEdgeAntialiasing(value bool) {
	objc.CallMethod[objc.Void](l_, "setAllowsEdgeAntialiasing:", value)
}

func (l_ Layer) AllowsGroupOpacity() bool {
	rv := objc.CallMethod[bool](l_, "allowsGroupOpacity")
	return rv
}

func (l_ Layer) SetAllowsGroupOpacity(value bool) {
	objc.CallMethod[objc.Void](l_, "setAllowsGroupOpacity:", value)
}

func (l_ Layer) Filters() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](l_, "filters")
	return rv
}

func (l_ Layer) SetFilters(value []objc.IObject) {
	objc.CallMethod[objc.Void](l_, "setFilters:", value)
}

func (l_ Layer) CompositingFilter() objc.Object {
	rv := objc.CallMethod[objc.Object](l_, "compositingFilter")
	return rv
}

func (l_ Layer) SetCompositingFilter(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, "setCompositingFilter:", value)
}

func (l_ Layer) BackgroundFilters() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](l_, "backgroundFilters")
	return rv
}

func (l_ Layer) SetBackgroundFilters(value []objc.IObject) {
	objc.CallMethod[objc.Void](l_, "setBackgroundFilters:", value)
}

func (l_ Layer) MinificationFilter() LayerContentsFilter {
	rv := objc.CallMethod[LayerContentsFilter](l_, "minificationFilter")
	return rv
}

func (l_ Layer) SetMinificationFilter(value LayerContentsFilter) {
	objc.CallMethod[objc.Void](l_, "setMinificationFilter:", value)
}

func (l_ Layer) MinificationFilterBias() float32 {
	rv := objc.CallMethod[float32](l_, "minificationFilterBias")
	return rv
}

func (l_ Layer) SetMinificationFilterBias(value float32) {
	objc.CallMethod[objc.Void](l_, "setMinificationFilterBias:", value)
}

func (l_ Layer) MagnificationFilter() LayerContentsFilter {
	rv := objc.CallMethod[LayerContentsFilter](l_, "magnificationFilter")
	return rv
}

func (l_ Layer) SetMagnificationFilter(value LayerContentsFilter) {
	objc.CallMethod[objc.Void](l_, "setMagnificationFilter:", value)
}

func (l_ Layer) IsOpaque() bool {
	rv := objc.CallMethod[bool](l_, "isOpaque")
	return rv
}

func (l_ Layer) SetOpaque(value bool) {
	objc.CallMethod[objc.Void](l_, "setOpaque:", value)
}

func (l_ Layer) EdgeAntialiasingMask() EdgeAntialiasingMask {
	rv := objc.CallMethod[EdgeAntialiasingMask](l_, "edgeAntialiasingMask")
	return rv
}

func (l_ Layer) SetEdgeAntialiasingMask(value EdgeAntialiasingMask) {
	objc.CallMethod[objc.Void](l_, "setEdgeAntialiasingMask:", value)
}

func (l_ Layer) IsGeometryFlipped() bool {
	rv := objc.CallMethod[bool](l_, "isGeometryFlipped")
	return rv
}

func (l_ Layer) SetGeometryFlipped(value bool) {
	objc.CallMethod[objc.Void](l_, "setGeometryFlipped:", value)
}

func (l_ Layer) DrawsAsynchronously() bool {
	rv := objc.CallMethod[bool](l_, "drawsAsynchronously")
	return rv
}

func (l_ Layer) SetDrawsAsynchronously(value bool) {
	objc.CallMethod[objc.Void](l_, "setDrawsAsynchronously:", value)
}

func (l_ Layer) ShouldRasterize() bool {
	rv := objc.CallMethod[bool](l_, "shouldRasterize")
	return rv
}

func (l_ Layer) SetShouldRasterize(value bool) {
	objc.CallMethod[objc.Void](l_, "setShouldRasterize:", value)
}

func (l_ Layer) RasterizationScale() float64 {
	rv := objc.CallMethod[float64](l_, "rasterizationScale")
	return rv
}

func (l_ Layer) SetRasterizationScale(value float64) {
	objc.CallMethod[objc.Void](l_, "setRasterizationScale:", value)
}

func (l_ Layer) ContentsFormat() LayerContentsFormat {
	rv := objc.CallMethod[LayerContentsFormat](l_, "contentsFormat")
	return rv
}

func (l_ Layer) SetContentsFormat(value LayerContentsFormat) {
	objc.CallMethod[objc.Void](l_, "setContentsFormat:", value)
}

func (l_ Layer) Frame() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, "frame")
	return rv
}

func (l_ Layer) SetFrame(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, "setFrame:", value)
}

func (l_ Layer) Bounds() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, "bounds")
	return rv
}

func (l_ Layer) SetBounds(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](l_, "setBounds:", value)
}

func (l_ Layer) Position() coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](l_, "position")
	return rv
}

func (l_ Layer) SetPosition(value coregraphics.Point) {
	objc.CallMethod[objc.Void](l_, "setPosition:", value)
}

func (l_ Layer) ZPosition() float64 {
	rv := objc.CallMethod[float64](l_, "zPosition")
	return rv
}

func (l_ Layer) SetZPosition(value float64) {
	objc.CallMethod[objc.Void](l_, "setZPosition:", value)
}

func (l_ Layer) AnchorPointZ() float64 {
	rv := objc.CallMethod[float64](l_, "anchorPointZ")
	return rv
}

func (l_ Layer) SetAnchorPointZ(value float64) {
	objc.CallMethod[objc.Void](l_, "setAnchorPointZ:", value)
}

func (l_ Layer) AnchorPoint() coregraphics.Point {
	rv := objc.CallMethod[coregraphics.Point](l_, "anchorPoint")
	return rv
}

func (l_ Layer) SetAnchorPoint(value coregraphics.Point) {
	objc.CallMethod[objc.Void](l_, "setAnchorPoint:", value)
}

func (l_ Layer) ContentsScale() float64 {
	rv := objc.CallMethod[float64](l_, "contentsScale")
	return rv
}

func (l_ Layer) SetContentsScale(value float64) {
	objc.CallMethod[objc.Void](l_, "setContentsScale:", value)
}

func (l_ Layer) Transform() Transform3D {
	rv := objc.CallMethod[Transform3D](l_, "transform")
	return rv
}

func (l_ Layer) SetTransform(value Transform3D) {
	objc.CallMethod[objc.Void](l_, "setTransform:", value)
}

func (l_ Layer) SublayerTransform() Transform3D {
	rv := objc.CallMethod[Transform3D](l_, "sublayerTransform")
	return rv
}

func (l_ Layer) SetSublayerTransform(value Transform3D) {
	objc.CallMethod[objc.Void](l_, "setSublayerTransform:", value)
}

func (l_ Layer) Sublayers() []Layer {
	rv := objc.CallMethod[[]Layer](l_, "sublayers")
	return rv
}

func (l_ Layer) SetSublayers(value []ILayer) {
	objc.CallMethod[objc.Void](l_, "setSublayers:", value)
}

func (l_ Layer) Superlayer() Layer {
	rv := objc.CallMethod[Layer](l_, "superlayer")
	return rv
}

func (l_ Layer) NeedsDisplayOnBoundsChange() bool {
	rv := objc.CallMethod[bool](l_, "needsDisplayOnBoundsChange")
	return rv
}

func (l_ Layer) SetNeedsDisplayOnBoundsChange(value bool) {
	objc.CallMethod[objc.Void](l_, "setNeedsDisplayOnBoundsChange:", value)
}

func (l_ Layer) LayoutManager() LayoutManagerWrapper {
	rv := objc.CallMethod[LayoutManagerWrapper](l_, "layoutManager")
	return rv
}

func (l_ Layer) SetLayoutManager(value LayoutManager) {
	po := objc.CreateProtocol("CALayoutManager", value)
	defer po.Release()
	objc.CallMethod[objc.Void](l_, "setLayoutManager:", po)
}

func (l_ Layer) SetLayoutManager0(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, "setLayoutManager:", value)
}

func (l_ Layer) AutoresizingMask() AutoresizingMask {
	rv := objc.CallMethod[AutoresizingMask](l_, "autoresizingMask")
	return rv
}

func (l_ Layer) SetAutoresizingMask(value AutoresizingMask) {
	objc.CallMethod[objc.Void](l_, "setAutoresizingMask:", value)
}

func (l_ Layer) Constraints() []Constraint {
	rv := objc.CallMethod[[]Constraint](l_, "constraints")
	return rv
}

func (l_ Layer) SetConstraints(value []IConstraint) {
	objc.CallMethod[objc.Void](l_, "setConstraints:", value)
}

func (l_ Layer) Actions() foundation.Dictionary {
	rv := objc.CallMethod[foundation.Dictionary](l_, "actions")
	return rv
}

func (l_ Layer) SetActions(value foundation.IDictionary) {
	objc.CallMethod[objc.Void](l_, "setActions:", value)
}

func (l_ Layer) VisibleRect() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](l_, "visibleRect")
	return rv
}

func (l_ Layer) Name() string {
	rv := objc.CallMethod[string](l_, "name")
	return rv
}

func (l_ Layer) SetName(value string) {
	objc.CallMethod[objc.Void](l_, "setName:", value)
}

func (l_ Layer) CornerCurve() LayerCornerCurve {
	rv := objc.CallMethod[LayerCornerCurve](l_, "cornerCurve")
	return rv
}

func (l_ Layer) SetCornerCurve(value LayerCornerCurve) {
	objc.CallMethod[objc.Void](l_, "setCornerCurve:", value)
}
