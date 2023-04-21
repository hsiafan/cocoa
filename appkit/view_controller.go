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
	rv := objc.CallMethod[ViewController](v_, "initWithNibName:bundle:", nibNameOrNil, nibBundleOrNil)
	return rv
}

func (v_ ViewController) Init() ViewController {
	rv := objc.CallMethod[ViewController](v_, "init")
	return rv
}

func (vc _ViewControllerClass) Alloc() ViewController {
	rv := objc.CallMethod[ViewController](vc, "alloc")
	return rv
}

func (vc _ViewControllerClass) New() ViewController {
	rv := objc.CallMethod[ViewController](vc, "new")
	rv.Autorelease()
	return rv
}

func NewViewController() ViewController {
	return ViewControllerClass.New()
}

func (v_ ViewController) LoadView() {
	objc.CallMethod[objc.Void](v_, "loadView")
}

func (v_ ViewController) CommitEditingWithDelegate_DidCommitSelector_ContextInfo(delegate objc.IObject, didCommitSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.CallMethod[objc.Void](v_, "commitEditingWithDelegate:didCommitSelector:contextInfo:", delegate, didCommitSelector, contextInfo)
}

func (v_ ViewController) CommitEditing() bool {
	rv := objc.CallMethod[bool](v_, "commitEditing")
	return rv
}

func (v_ ViewController) DiscardEditing() {
	objc.CallMethod[objc.Void](v_, "discardEditing")
}

func (v_ ViewController) DismissController(sender objc.IObject) {
	objc.CallMethod[objc.Void](v_, "dismissController:", sender)
}

func (v_ ViewController) ViewDidLoad() {
	objc.CallMethod[objc.Void](v_, "viewDidLoad")
}

func (v_ ViewController) ViewWillAppear() {
	objc.CallMethod[objc.Void](v_, "viewWillAppear")
}

func (v_ ViewController) ViewDidAppear() {
	objc.CallMethod[objc.Void](v_, "viewDidAppear")
}

func (v_ ViewController) ViewWillDisappear() {
	objc.CallMethod[objc.Void](v_, "viewWillDisappear")
}

func (v_ ViewController) ViewDidDisappear() {
	objc.CallMethod[objc.Void](v_, "viewDidDisappear")
}

func (v_ ViewController) UpdateViewConstraints() {
	objc.CallMethod[objc.Void](v_, "updateViewConstraints")
}

func (v_ ViewController) ViewWillLayout() {
	objc.CallMethod[objc.Void](v_, "viewWillLayout")
}

func (v_ ViewController) ViewDidLayout() {
	objc.CallMethod[objc.Void](v_, "viewDidLayout")
}

func (v_ ViewController) AddChildViewController(childViewController IViewController) {
	objc.CallMethod[objc.Void](v_, "addChildViewController:", childViewController)
}

func (v_ ViewController) TransitionFromViewController_ToViewController_Options_CompletionHandler(fromViewController IViewController, toViewController IViewController, options ViewControllerTransitionOptions, completion func()) {
	objc.CallMethod[objc.Void](v_, "transitionFromViewController:toViewController:options:completionHandler:", fromViewController, toViewController, options, completion)
}

func (v_ ViewController) InsertChildViewController_AtIndex(childViewController IViewController, index int) {
	objc.CallMethod[objc.Void](v_, "insertChildViewController:atIndex:", childViewController, index)
}

func (v_ ViewController) RemoveChildViewControllerAtIndex(index int) {
	objc.CallMethod[objc.Void](v_, "removeChildViewControllerAtIndex:", index)
}

func (v_ ViewController) RemoveFromParentViewController() {
	objc.CallMethod[objc.Void](v_, "removeFromParentViewController")
}

func (v_ ViewController) PreferredContentSizeDidChangeForViewController(viewController IViewController) {
	objc.CallMethod[objc.Void](v_, "preferredContentSizeDidChangeForViewController:", viewController)
}

func (v_ ViewController) PresentViewController_Animator(viewController IViewController, animator ViewControllerPresentationAnimator) {
	po := objc.CreateProtocol("NSViewControllerPresentationAnimator", animator)
	defer po.Release()
	objc.CallMethod[objc.Void](v_, "presentViewController:animator:", viewController, po)
}

func (v_ ViewController) PresentViewController0_Animator(viewController IViewController, animator objc.IObject) {
	objc.CallMethod[objc.Void](v_, "presentViewController:animator:", viewController, animator)
}

func (v_ ViewController) DismissViewController(viewController IViewController) {
	objc.CallMethod[objc.Void](v_, "dismissViewController:", viewController)
}

func (v_ ViewController) PresentViewController_AsPopoverRelativeToRect_OfView_PreferredEdge_Behavior(viewController IViewController, positioningRect foundation.Rect, positioningView IView, preferredEdge foundation.RectEdge, behavior PopoverBehavior) {
	objc.CallMethod[objc.Void](v_, "presentViewController:asPopoverRelativeToRect:ofView:preferredEdge:behavior:", viewController, positioningRect, positioningView, preferredEdge, behavior)
}

func (v_ ViewController) PresentViewControllerAsModalWindow(viewController IViewController) {
	objc.CallMethod[objc.Void](v_, "presentViewControllerAsModalWindow:", viewController)
}

func (v_ ViewController) PresentViewControllerAsSheet(viewController IViewController) {
	objc.CallMethod[objc.Void](v_, "presentViewControllerAsSheet:", viewController)
}

func (v_ ViewController) ViewWillTransitionToSize(newSize foundation.Size) {
	objc.CallMethod[objc.Void](v_, "viewWillTransitionToSize:", newSize)
}

func (v_ ViewController) RepresentedObject() objc.Object {
	rv := objc.CallMethod[objc.Object](v_, "representedObject")
	return rv
}

func (v_ ViewController) SetRepresentedObject(value objc.IObject) {
	objc.CallMethod[objc.Void](v_, "setRepresentedObject:", value)
}

func (v_ ViewController) NibBundle() foundation.Bundle {
	rv := objc.CallMethod[foundation.Bundle](v_, "nibBundle")
	return rv
}

func (v_ ViewController) NibName() NibName {
	rv := objc.CallMethod[NibName](v_, "nibName")
	return rv
}

func (v_ ViewController) View() View {
	rv := objc.CallMethod[View](v_, "view")
	return rv
}

func (v_ ViewController) SetView(value IView) {
	objc.CallMethod[objc.Void](v_, "setView:", value)
}

func (v_ ViewController) Title() string {
	rv := objc.CallMethod[string](v_, "title")
	return rv
}

func (v_ ViewController) SetTitle(value string) {
	objc.CallMethod[objc.Void](v_, "setTitle:", value)
}

func (v_ ViewController) Storyboard() Storyboard {
	rv := objc.CallMethod[Storyboard](v_, "storyboard")
	return rv
}

func (v_ ViewController) IsViewLoaded() bool {
	rv := objc.CallMethod[bool](v_, "isViewLoaded")
	return rv
}

func (v_ ViewController) PreferredContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, "preferredContentSize")
	return rv
}

func (v_ ViewController) SetPreferredContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](v_, "setPreferredContentSize:", value)
}

func (v_ ViewController) ChildViewControllers() []ViewController {
	rv := objc.CallMethod[[]ViewController](v_, "childViewControllers")
	return rv
}

func (v_ ViewController) SetChildViewControllers(value []IViewController) {
	objc.CallMethod[objc.Void](v_, "setChildViewControllers:", value)
}

func (v_ ViewController) ParentViewController() ViewController {
	rv := objc.CallMethod[ViewController](v_, "parentViewController")
	return rv
}

func (v_ ViewController) PresentedViewControllers() []ViewController {
	rv := objc.CallMethod[[]ViewController](v_, "presentedViewControllers")
	return rv
}

func (v_ ViewController) PresentingViewController() ViewController {
	rv := objc.CallMethod[ViewController](v_, "presentingViewController")
	return rv
}

func (v_ ViewController) ExtensionContext() foundation.ExtensionContext {
	rv := objc.CallMethod[foundation.ExtensionContext](v_, "extensionContext")
	return rv
}

func (v_ ViewController) PreferredScreenOrigin() foundation.Point {
	rv := objc.CallMethod[foundation.Point](v_, "preferredScreenOrigin")
	return rv
}

func (v_ ViewController) SetPreferredScreenOrigin(value foundation.Point) {
	objc.CallMethod[objc.Void](v_, "setPreferredScreenOrigin:", value)
}

func (v_ ViewController) PreferredMaximumSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, "preferredMaximumSize")
	return rv
}

func (v_ ViewController) PreferredMinimumSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](v_, "preferredMinimumSize")
	return rv
}

func (v_ ViewController) SourceItemView() View {
	rv := objc.CallMethod[View](v_, "sourceItemView")
	return rv
}

func (v_ ViewController) SetSourceItemView(value IView) {
	objc.CallMethod[objc.Void](v_, "setSourceItemView:", value)
}
