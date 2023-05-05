// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type ViewControllerPresentationAnimator interface {
	// required
	AnimatePresentationOfViewController_FromViewController(viewController ViewController, fromViewController ViewController)
	// required
	AnimateDismissalOfViewController_FromViewController(viewController ViewController, fromViewController ViewController)
}

type ViewControllerPresentationAnimatorWrapper struct {
	objc.Object
}

func (v_ ViewControllerPresentationAnimatorWrapper) AnimatePresentationOfViewController_FromViewController(viewController IViewController, fromViewController IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("animatePresentationOfViewController:fromViewController:"), objc.ExtractPtr(viewController), objc.ExtractPtr(fromViewController))
}

func (v_ ViewControllerPresentationAnimatorWrapper) AnimateDismissalOfViewController_FromViewController(viewController IViewController, fromViewController IViewController) {
	objc.CallMethod[objc.Void](v_, objc.GetSelector("animateDismissalOfViewController:fromViewController:"), objc.ExtractPtr(viewController), objc.ExtractPtr(fromViewController))
}
