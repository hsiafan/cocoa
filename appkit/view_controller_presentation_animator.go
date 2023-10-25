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

func WrapViewControllerPresentationAnimator(v ViewControllerPresentationAnimator) objc.Object {
	return objc.WrapAsProtocol("NSViewControllerPresentationAnimator", v)
}

type ViewControllerPresentationAnimatorBase struct {
}
