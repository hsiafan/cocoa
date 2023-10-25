// auto-generated code, do not modify
package foundation

import (
	"github.com/hsiafan/cocoa/objc"
)

type UserActivityDelegate interface {
	ImplementsUserActivity_DidReceiveInputStream_OutputStream() bool
	// optional
	UserActivity_DidReceiveInputStream_OutputStream(userActivity UserActivity, inputStream InputStream, outputStream OutputStream)
	ImplementsUserActivityWasContinued() bool
	// optional
	UserActivityWasContinued(userActivity UserActivity)
	ImplementsUserActivityWillSave() bool
	// optional
	UserActivityWillSave(userActivity UserActivity)
}

func WrapUserActivityDelegate(v UserActivityDelegate) objc.Object {
	return objc.WrapAsProtocol("NSUserActivityDelegate", v)
}

type UserActivityDelegateBase struct {
}

func (p *UserActivityDelegateBase) ImplementsUserActivity_DidReceiveInputStream_OutputStream() bool {
	return false
}

func (p *UserActivityDelegateBase) UserActivity_DidReceiveInputStream_OutputStream(userActivity UserActivity, inputStream InputStream, outputStream OutputStream) {
	panic("unimpemented protocol method")
}

func (p *UserActivityDelegateBase) ImplementsUserActivityWasContinued() bool {
	return false
}

func (p *UserActivityDelegateBase) UserActivityWasContinued(userActivity UserActivity) {
	panic("unimpemented protocol method")
}

func (p *UserActivityDelegateBase) ImplementsUserActivityWillSave() bool {
	return false
}

func (p *UserActivityDelegateBase) UserActivityWillSave(userActivity UserActivity) {
	panic("unimpemented protocol method")
}
