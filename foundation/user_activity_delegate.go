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

type UserActivityDelegateImpl struct {
	_UserActivity_DidReceiveInputStream_OutputStream func(userActivity UserActivity, inputStream InputStream, outputStream OutputStream)
	_UserActivityWasContinued                        func(userActivity UserActivity)
	_UserActivityWillSave                            func(userActivity UserActivity)
}

func (di *UserActivityDelegateImpl) ImplementsUserActivity_DidReceiveInputStream_OutputStream() bool {
	return di._UserActivity_DidReceiveInputStream_OutputStream != nil
}

func (di *UserActivityDelegateImpl) SetUserActivity_DidReceiveInputStream_OutputStream(f func(userActivity UserActivity, inputStream InputStream, outputStream OutputStream)) {
	di._UserActivity_DidReceiveInputStream_OutputStream = f
}

func (di *UserActivityDelegateImpl) UserActivity_DidReceiveInputStream_OutputStream(userActivity UserActivity, inputStream InputStream, outputStream OutputStream) {
	di._UserActivity_DidReceiveInputStream_OutputStream(userActivity, inputStream, outputStream)
}
func (di *UserActivityDelegateImpl) ImplementsUserActivityWasContinued() bool {
	return di._UserActivityWasContinued != nil
}

func (di *UserActivityDelegateImpl) SetUserActivityWasContinued(f func(userActivity UserActivity)) {
	di._UserActivityWasContinued = f
}

func (di *UserActivityDelegateImpl) UserActivityWasContinued(userActivity UserActivity) {
	di._UserActivityWasContinued(userActivity)
}
func (di *UserActivityDelegateImpl) ImplementsUserActivityWillSave() bool {
	return di._UserActivityWillSave != nil
}

func (di *UserActivityDelegateImpl) SetUserActivityWillSave(f func(userActivity UserActivity)) {
	di._UserActivityWillSave = f
}

func (di *UserActivityDelegateImpl) UserActivityWillSave(userActivity UserActivity) {
	di._UserActivityWillSave(userActivity)
}

type UserActivityDelegateWrapper struct {
	objc.Object
}

func (u_ *UserActivityDelegateWrapper) ImplementsUserActivity_DidReceiveInputStream_OutputStream() bool {
	return u_.RespondsToSelector(objc.GetSelector("userActivity:didReceiveInputStream:outputStream:"))
}

func (u_ UserActivityDelegateWrapper) UserActivity_DidReceiveInputStream_OutputStream(userActivity IUserActivity, inputStream IInputStream, outputStream IOutputStream) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("userActivity:didReceiveInputStream:outputStream:"), objc.ExtractPtr(userActivity), objc.ExtractPtr(inputStream), objc.ExtractPtr(outputStream))
}

func (u_ *UserActivityDelegateWrapper) ImplementsUserActivityWasContinued() bool {
	return u_.RespondsToSelector(objc.GetSelector("userActivityWasContinued:"))
}

func (u_ UserActivityDelegateWrapper) UserActivityWasContinued(userActivity IUserActivity) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("userActivityWasContinued:"), objc.ExtractPtr(userActivity))
}

func (u_ *UserActivityDelegateWrapper) ImplementsUserActivityWillSave() bool {
	return u_.RespondsToSelector(objc.GetSelector("userActivityWillSave:"))
}

func (u_ UserActivityDelegateWrapper) UserActivityWillSave(userActivity IUserActivity) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("userActivityWillSave:"), objc.ExtractPtr(userActivity))
}
