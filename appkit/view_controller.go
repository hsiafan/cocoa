// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var ViewControllerClass = _ViewControllerClass{objc.GetClass("NSViewController")}

type _ViewControllerClass struct {
	objc.Class
}

type IViewController interface {
	IResponder
	LoadView()
	CommitEditingWithDelegate_DidCommitSelector_ContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer)
	CommitEditing() bool
	DiscardEditing()
	DismissController(sender objc.IObject)
	ViewDidLoad()
	ViewWillAppear()
	ViewDidAppear()
	ViewWillDisappear()
	ViewDidDisappear()
	UpdateViewConstraints()
	ViewWillLayout()
	ViewDidLayout()
	AddChildViewController(childViewController IViewController)
	TransitionFromViewController_ToViewController_Options_CompletionHandler(fromViewController IViewController, toViewController IViewController, options ViewControllerTransitionOptions, completion func())
	InsertChildViewController_AtIndex(childViewController IViewController, index int)
	RemoveChildViewControllerAtIndex(index int)
	RemoveFromParentViewController()
	PreferredContentSizeDidChangeForViewController(viewController IViewController)
	PresentViewController_Animator(viewController IViewController, animator ViewControllerPresentationAnimator)
	PresentViewController0_Animator(viewController IViewController, animator objc.IObject)
	DismissViewController(viewController IViewController)
	PresentViewController_AsPopoverRelativeToRect_OfView_PreferredEdge_Behavior(viewController IViewController, positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge, behavior PopoverBehavior)
	PresentViewControllerAsModalWindow(viewController IViewController)
	PresentViewControllerAsSheet(viewController IViewController)
	ViewWillTransitionToSize(newSize foundation.Size)
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.IObject)
	NibBundle() foundation.Bundle
	NibName() NibName
	View() View
	SetView(value IView)
	Title() string
	SetTitle(value string)
	Storyboard() Storyboard
	IsViewLoaded() bool
	PreferredContentSize() foundation.Size
	SetPreferredContentSize(value foundation.Size)
	ChildViewControllers() []ViewController
	SetChildViewControllers(value []IViewController)
	ParentViewController() ViewController
	PresentedViewControllers() []ViewController
	PresentingViewController() ViewController
	ExtensionContext() foundation.ExtensionContext
	PreferredScreenOrigin() foundation.Point
	SetPreferredScreenOrigin(value foundation.Point)
	PreferredMaximumSize() foundation.Size
	PreferredMinimumSize() foundation.Size
	SourceItemView() View
	SetSourceItemView(value IView)
}

type ViewController struct {
	Responder
}

func MakeViewController(ptr unsafe.Pointer) ViewController {
	return ViewController{
		Responder: MakeResponder(ptr),
	}
}

func (v_ ViewController) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) ViewController {
	rv := objc.CallMethod[ViewController](v_, objc.GetSelector("initWithNibName:bundle:"), nibNameOrNil, nibBundleOrNil)
	return rv
}

func (v_ ViewController) Init() ViewController {
	rv := objc.CallMethod[ViewController](v_, objc.GetSelector("init"))
	return rv
}

func (vc _ViewControllerClass) Alloc() ViewController {
	rv := objc.CallMethod[ViewController](vc, objc.GetSelector("alloc"))
	return rv
}

func (vc _ViewControllerClass) New() ViewController {
	rv := objc.CallMethod[ViewController](vc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewViewController() ViewController {
	return ViewControllerClass.New()
}

func (v_ ViewController) LoadView() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("loadView"))
}

func (v_ ViewController) CommitEditingWithDelegate_DidCommitSelector_ContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
}

func (v_ ViewController) CommitEditing() bool {
	rv := objc.CallMethod[bool](v_, objc.GetSelector("commitEditing"))
	return rv
}

func (v_ ViewController) DiscardEditing() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("discardEditing"))
}

func (v_ ViewController) DismissController(sender objc.IObject) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("dismissController:"), sender)
}

func (v_ ViewController) ViewDidLoad() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("viewDidLoad"))
}

func (v_ ViewController) ViewWillAppear() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("viewWillAppear"))
}

func (v_ ViewController) ViewDidAppear() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("viewDidAppear"))
}

func (v_ ViewController) ViewWillDisappear() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("viewWillDisappear"))
}

func (v_ ViewController) ViewDidDisappear() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("viewDidDisappear"))
}

func (v_ ViewController) UpdateViewConstraints() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("updateViewConstraints"))
}

func (v_ ViewController) ViewWillLayout() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("viewWillLayout"))
}

func (v_ ViewController) ViewDidLayout() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("viewDidLayout"))
}

func (v_ ViewController) AddChildViewController(childViewController IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("addChildViewController:"), childViewController)
}

func (v_ ViewController) TransitionFromViewController_ToViewController_Options_CompletionHandler(fromViewController IViewController, toViewController IViewController, options ViewControllerTransitionOptions, completion func()) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("transitionFromViewController:toViewController:options:completionHandler:"), fromViewController, toViewController, options, completion)
}

func (v_ ViewController) InsertChildViewController_AtIndex(childViewController IViewController, index int) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("insertChildViewController:atIndex:"), childViewController, index)
}

func (v_ ViewController) RemoveChildViewControllerAtIndex(index int) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("removeChildViewControllerAtIndex:"), index)
}

func (v_ ViewController) RemoveFromParentViewController() {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("removeFromParentViewController"))
}

func (v_ ViewController) PreferredContentSizeDidChangeForViewController(viewController IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("preferredContentSizeDidChangeForViewController:"), viewController)
}

func (v_ ViewController) PresentViewController_Animator(viewController IViewController, animator ViewControllerPresentationAnimator) {
	po := objc.CreateProtocol("NSViewControllerPresentationAnimator", animator)
	defer po.Release()
	objc.CallMethod[objc.Void](v_, objc.GetSelector("presentViewController:animator:"), viewController, po)
}

func (v_ ViewController) PresentViewController0_Animator(viewController IViewController, animator objc.IObject) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("presentViewController:animator:"), viewController, animator)
}

func (v_ ViewController) DismissViewController(viewController IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("dismissViewController:"), viewController)
}

func (v_ ViewController) PresentViewController_AsPopoverRelativeToRect_OfView_PreferredEdge_Behavior(viewController IViewController, positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge, behavior PopoverBehavior) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:"), viewController, positioningRect, positioningView, preferredEdge, behavior)
}

func (v_ ViewController) PresentViewControllerAsModalWindow(viewController IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("presentViewControllerAsModalWindow:"), viewController)
}

func (v_ ViewController) PresentViewControllerAsSheet(viewController IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("presentViewControllerAsSheet:"), viewController)
}

func (v_ ViewController) ViewWillTransitionToSize(newSize foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("viewWillTransitionToSize:"), newSize)
}

func (v_ ViewController) RepresentedObject() objc.Object {
	rv := objc.CallMethod[objc.Object](v_, objc.GetSelector("representedObject"))
	return rv
}

func (v_ ViewController) SetRepresentedObject(value objc.IObject) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setRepresentedObject:"), value)
}

func (v_ ViewController) NibBundle() foundation.Bundle {
	rv := objc.CallMethod[foundation.Bundle](v_, objc.GetSelector("nibBundle"))
	return rv
}

func (v_ ViewController) NibName() NibName {
	rv := objc.CallMethod[NibName](v_, objc.GetSelector("nibName"))
	return rv
}

func (v_ ViewController) View() View {
	rv := objc.CallMethod[View](v_, objc.GetSelector("view"))
	return rv
}

func (v_ ViewController) SetView(value IView) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setView:"), value)
}

func (v_ ViewController) Title() string {
	rv := objc.CallMethod[string](v_, objc.GetSelector("title"))
	return rv
}

func (v_ ViewController) SetTitle(value string) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setTitle:"), value)
}

func (v_ ViewController) Storyboard() Storyboard {
	rv := objc.CallMethod[Storyboard](v_, objc.GetSelector("storyboard"))
	return rv
}

func (v_ ViewController) IsViewLoaded() bool {
	rv := objc.CallMethod[bool](v_, objc.GetSelector("isViewLoaded"))
	return rv
}

func (v_ ViewController) PreferredContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.GetSelector("preferredContentSize"))
	return rv
}

func (v_ ViewController) SetPreferredContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setPreferredContentSize:"), value)
}

func (v_ ViewController) ChildViewControllers() []ViewController {
	rv := objc.CallMethod[[]ViewController](v_, objc.GetSelector("childViewControllers"))
	return rv
}

func (v_ ViewController) SetChildViewControllers(value []IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setChildViewControllers:"), value)
}

func (v_ ViewController) ParentViewController() ViewController {
	rv := objc.CallMethod[ViewController](v_, objc.GetSelector("parentViewController"))
	return rv
}

func (v_ ViewController) PresentedViewControllers() []ViewController {
	rv := objc.CallMethod[[]ViewController](v_, objc.GetSelector("presentedViewControllers"))
	return rv
}

func (v_ ViewController) PresentingViewController() ViewController {
	rv := objc.CallMethod[ViewController](v_, objc.GetSelector("presentingViewController"))
	return rv
}

func (v_ ViewController) ExtensionContext() foundation.ExtensionContext {
	rv := objc.CallMethod[foundation.ExtensionContext](v_, objc.GetSelector("extensionContext"))
	return rv
}

func (v_ ViewController) PreferredScreenOrigin() foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, objc.GetSelector("preferredScreenOrigin"))
	return rv
}

func (v_ ViewController) SetPreferredScreenOrigin(value foundation.Point) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setPreferredScreenOrigin:"), value)
}

func (v_ ViewController) PreferredMaximumSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.GetSelector("preferredMaximumSize"))
	return rv
}

func (v_ ViewController) PreferredMinimumSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, objc.GetSelector("preferredMinimumSize"))
	return rv
}

func (v_ ViewController) SourceItemView() View {
	rv := objc.CallMethod[View](v_, objc.GetSelector("sourceItemView"))
	return rv
}

func (v_ ViewController) SetSourceItemView(value IView) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("setSourceItemView:"), value)
}
