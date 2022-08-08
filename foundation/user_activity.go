// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/internal"
	"github.com/hsiafan/cocoa/objc"
)

var UserActivityClass = _UserActivityClass{objc.GetClass("NSUserActivity")}

type _UserActivityClass struct {
	objc.Class
}

type IUserActivity interface {
	objc.IObject
	BecomeCurrent()
	ResignCurrent()
	Invalidate()
	GetContinuationStreamsWithCompletionHandler(completionHandler func(inputStream InputStream, outputStream OutputStream, error Error))
	ActivityType() string
	Title() string
	SetTitle(value string)
	RequiredUserInfoKeys() Set
	SetRequiredUserInfoKeys(value ISet)
	TargetContentIdentifier() string
	SetTargetContentIdentifier(value string)
	NeedsSave() bool
	SetNeedsSave(value bool)
	Keywords() Set
	SetKeywords(value ISet)
	PersistentIdentifier() UserActivityPersistentIdentifier
	SetPersistentIdentifier(value UserActivityPersistentIdentifier)
	IsEligibleForHandoff() bool
	SetEligibleForHandoff(value bool)
	IsEligibleForSearch() bool
	SetEligibleForSearch(value bool)
	IsEligibleForPublicIndexing() bool
	SetEligibleForPublicIndexing(value bool)
	ExpirationDate() Date
	SetExpirationDate(value IDate)
	Delegate() UserActivityDelegateWrapper
	SetDelegate(value UserActivityDelegate)
	SupportsContinuationStreams() bool
	SetSupportsContinuationStreams(value bool)
	WebpageURL() URL
	SetWebpageURL(value IURL)
	ReferrerURL() URL
	SetReferrerURL(value IURL)
}

type UserActivity struct {
	objc.Object
}

func MakeUserActivity(ptr unsafe.Pointer) UserActivity {
	return UserActivity{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ UserActivity) InitWithActivityType(activityType string) UserActivity {
	rv := ffi.CallMethod[UserActivity](u_, "initWithActivityType:", activityType)
	return rv
}

func (u_ UserActivity) Init() UserActivity {
	rv := ffi.CallMethod[UserActivity](u_, "init")
	return rv
}

func (uc _UserActivityClass) Alloc() UserActivity {
	rv := ffi.CallMethod[UserActivity](uc, "alloc")
	return rv
}

func (uc _UserActivityClass) New() UserActivity {
	rv := ffi.CallMethod[UserActivity](uc, "new")
	rv.Autorelease()
	return rv
}

func NewUserActivity() UserActivity {
	return UserActivityClass.New()
}

func (u_ UserActivity) BecomeCurrent() {
	ffi.CallMethod[ffi.Void](u_, "becomeCurrent")
}

func (u_ UserActivity) ResignCurrent() {
	ffi.CallMethod[ffi.Void](u_, "resignCurrent")
}

func (u_ UserActivity) Invalidate() {
	ffi.CallMethod[ffi.Void](u_, "invalidate")
}

func (uc _UserActivityClass) DeleteAllSavedUserActivitiesWithCompletionHandler(handler func()) {
	ffi.CallMethod[ffi.Void](uc, "deleteAllSavedUserActivitiesWithCompletionHandler:", handler)
}

func (uc _UserActivityClass) DeleteSavedUserActivitiesWithPersistentIdentifiers_CompletionHandler(persistentIdentifiers []UserActivityPersistentIdentifier, handler func()) {
	ffi.CallMethod[ffi.Void](uc, "deleteSavedUserActivitiesWithPersistentIdentifiers:completionHandler:", persistentIdentifiers, handler)
}

func (u_ UserActivity) GetContinuationStreamsWithCompletionHandler(completionHandler func(inputStream InputStream, outputStream OutputStream, error Error)) {
	ffi.CallMethod[ffi.Void](u_, "getContinuationStreamsWithCompletionHandler:", completionHandler)
}

func (u_ UserActivity) ActivityType() string {
	rv := ffi.CallMethod[string](u_, "activityType")
	return rv
}

func (u_ UserActivity) Title() string {
	rv := ffi.CallMethod[string](u_, "title")
	return rv
}

func (u_ UserActivity) SetTitle(value string) {
	ffi.CallMethod[ffi.Void](u_, "setTitle:", value)
}

func (u_ UserActivity) RequiredUserInfoKeys() Set {
	rv := ffi.CallMethod[Set](u_, "requiredUserInfoKeys")
	return rv
}

func (u_ UserActivity) SetRequiredUserInfoKeys(value ISet) {
	ffi.CallMethod[ffi.Void](u_, "setRequiredUserInfoKeys:", value)
}

func (u_ UserActivity) TargetContentIdentifier() string {
	rv := ffi.CallMethod[string](u_, "targetContentIdentifier")
	return rv
}

func (u_ UserActivity) SetTargetContentIdentifier(value string) {
	ffi.CallMethod[ffi.Void](u_, "setTargetContentIdentifier:", value)
}

func (u_ UserActivity) NeedsSave() bool {
	rv := ffi.CallMethod[bool](u_, "needsSave")
	return rv
}

func (u_ UserActivity) SetNeedsSave(value bool) {
	ffi.CallMethod[ffi.Void](u_, "setNeedsSave:", value)
}

func (u_ UserActivity) Keywords() Set {
	rv := ffi.CallMethod[Set](u_, "keywords")
	return rv
}

func (u_ UserActivity) SetKeywords(value ISet) {
	ffi.CallMethod[ffi.Void](u_, "setKeywords:", value)
}

func (u_ UserActivity) PersistentIdentifier() UserActivityPersistentIdentifier {
	rv := ffi.CallMethod[UserActivityPersistentIdentifier](u_, "persistentIdentifier")
	return rv
}

func (u_ UserActivity) SetPersistentIdentifier(value UserActivityPersistentIdentifier) {
	ffi.CallMethod[ffi.Void](u_, "setPersistentIdentifier:", value)
}

func (u_ UserActivity) IsEligibleForHandoff() bool {
	rv := ffi.CallMethod[bool](u_, "isEligibleForHandoff")
	return rv
}

func (u_ UserActivity) SetEligibleForHandoff(value bool) {
	ffi.CallMethod[ffi.Void](u_, "setEligibleForHandoff:", value)
}

func (u_ UserActivity) IsEligibleForSearch() bool {
	rv := ffi.CallMethod[bool](u_, "isEligibleForSearch")
	return rv
}

func (u_ UserActivity) SetEligibleForSearch(value bool) {
	ffi.CallMethod[ffi.Void](u_, "setEligibleForSearch:", value)
}

func (u_ UserActivity) IsEligibleForPublicIndexing() bool {
	rv := ffi.CallMethod[bool](u_, "isEligibleForPublicIndexing")
	return rv
}

func (u_ UserActivity) SetEligibleForPublicIndexing(value bool) {
	ffi.CallMethod[ffi.Void](u_, "setEligibleForPublicIndexing:", value)
}

func (u_ UserActivity) ExpirationDate() Date {
	rv := ffi.CallMethod[Date](u_, "expirationDate")
	return rv
}

func (u_ UserActivity) SetExpirationDate(value IDate) {
	ffi.CallMethod[ffi.Void](u_, "setExpirationDate:", value)
}

func (u_ UserActivity) Delegate() UserActivityDelegateWrapper {
	rv := ffi.CallMethod[UserActivityDelegateWrapper](u_, "delegate")
	return rv
}

func (u_ UserActivity) SetDelegate(value UserActivityDelegate) {
	po := ffi.CreateProtocol(value)
	defer po.Release()
	objc.SetAssociatedObject(u_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	ffi.CallMethod[ffi.Void](u_, "setDelegate:", po)
}

func (u_ UserActivity) SupportsContinuationStreams() bool {
	rv := ffi.CallMethod[bool](u_, "supportsContinuationStreams")
	return rv
}

func (u_ UserActivity) SetSupportsContinuationStreams(value bool) {
	ffi.CallMethod[ffi.Void](u_, "setSupportsContinuationStreams:", value)
}

func (u_ UserActivity) WebpageURL() URL {
	rv := ffi.CallMethod[URL](u_, "webpageURL")
	return rv
}

func (u_ UserActivity) SetWebpageURL(value IURL) {
	ffi.CallMethod[ffi.Void](u_, "setWebpageURL:", value)
}

func (u_ UserActivity) ReferrerURL() URL {
	rv := ffi.CallMethod[URL](u_, "referrerURL")
	return rv
}

func (u_ UserActivity) SetReferrerURL(value IURL) {
	ffi.CallMethod[ffi.Void](u_, "setReferrerURL:", value)
}

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
	ffi.CallMethod[ffi.Void](u_, "userActivity:didReceiveInputStream:outputStream:", userActivity, inputStream, outputStream)
}

func (u_ *UserActivityDelegateWrapper) ImplementsUserActivityWasContinued() bool {
	return u_.RespondsToSelector(objc.GetSelector("userActivityWasContinued:"))
}

func (u_ UserActivityDelegateWrapper) UserActivityWasContinued(userActivity IUserActivity) {
	ffi.CallMethod[ffi.Void](u_, "userActivityWasContinued:", userActivity)
}

func (u_ *UserActivityDelegateWrapper) ImplementsUserActivityWillSave() bool {
	return u_.RespondsToSelector(objc.GetSelector("userActivityWillSave:"))
}

func (u_ UserActivityDelegateWrapper) UserActivityWillSave(userActivity IUserActivity) {
	ffi.CallMethod[ffi.Void](u_, "userActivityWillSave:", userActivity)
}
