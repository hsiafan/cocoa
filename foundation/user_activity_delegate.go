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

type UserActivityDelegateCreator struct {
	className string
	class     objc.Class
}

func NewUserActivityDelegateCreator(name string) *UserActivityDelegateCreator {
	class := objc.AllocateClassPair(objc.GetClass("NSObject"), name, 0)
	objc.RegisterClassPair(class)
	return &UserActivityDelegateCreator{className: name, class: class}
}

func (c *UserActivityDelegateCreator) SetUserActivity_DidReceiveInputStream_OutputStream(handle func(o objc.Object, userActivity UserActivity, inputStream InputStream, outputStream OutputStream)) {
	objc.AddMethod(c.class, objc.SEL("userActivity:didReceiveInputStream:outputStream:"), handle)
}

func (c *UserActivityDelegateCreator) SetUserActivityWasContinued(handle func(o objc.Object, userActivity UserActivity)) {
	objc.AddMethod(c.class, objc.SEL("userActivityWasContinued:"), handle)
}

func (c *UserActivityDelegateCreator) SetUserActivityWillSave(handle func(o objc.Object, userActivity UserActivity)) {
	objc.AddMethod(c.class, objc.SEL("userActivityWillSave:"), handle)
}

func (c *UserActivityDelegateCreator) Create() objc.Object {
	return c.class.CreateInstance(0)
}
