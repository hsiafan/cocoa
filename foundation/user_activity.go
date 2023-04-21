// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	SetDelegate0(value objc.IObject)
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
	rv := objc.CallMethod[UserActivity](u_, "initWithActivityType:", activityType)
	return rv
}

func (u_ UserActivity) Init() UserActivity {
	rv := objc.CallMethod[UserActivity](u_, "init")
	return rv
}

func (uc _UserActivityClass) Alloc() UserActivity {
	rv := objc.CallMethod[UserActivity](uc, "alloc")
	return rv
}

func (uc _UserActivityClass) New() UserActivity {
	rv := objc.CallMethod[UserActivity](uc, "new")
	rv.Autorelease()
	return rv
}

func NewUserActivity() UserActivity {
	return UserActivityClass.New()
}

func (u_ UserActivity) BecomeCurrent() {
	objc.CallMethod[objc.Void](u_, "becomeCurrent")
}

func (u_ UserActivity) ResignCurrent() {
	objc.CallMethod[objc.Void](u_, "resignCurrent")
}

func (u_ UserActivity) Invalidate() {
	objc.CallMethod[objc.Void](u_, "invalidate")
}

func (uc _UserActivityClass) DeleteAllSavedUserActivitiesWithCompletionHandler(handler func()) {
	objc.CallMethod[objc.Void](uc, "deleteAllSavedUserActivitiesWithCompletionHandler:", handler)
}

func (uc _UserActivityClass) DeleteSavedUserActivitiesWithPersistentIdentifiers_CompletionHandler(persistentIdentifiers []UserActivityPersistentIdentifier, handler func()) {
	objc.CallMethod[objc.Void](uc, "deleteSavedUserActivitiesWithPersistentIdentifiers:completionHandler:", persistentIdentifiers, handler)
}

func (u_ UserActivity) GetContinuationStreamsWithCompletionHandler(completionHandler func(inputStream InputStream, outputStream OutputStream, error Error)) {
	objc.CallMethod[objc.Void](u_, "getContinuationStreamsWithCompletionHandler:", completionHandler)
}

func (u_ UserActivity) ActivityType() string {
	rv := objc.CallMethod[string](u_, "activityType")
	return rv
}

func (u_ UserActivity) Title() string {
	rv := objc.CallMethod[string](u_, "title")
	return rv
}

func (u_ UserActivity) SetTitle(value string) {
	objc.CallMethod[objc.Void](u_, "setTitle:", value)
}

func (u_ UserActivity) RequiredUserInfoKeys() Set {
	rv := objc.CallMethod[Set](u_, "requiredUserInfoKeys")
	return rv
}

func (u_ UserActivity) SetRequiredUserInfoKeys(value ISet) {
	objc.CallMethod[objc.Void](u_, "setRequiredUserInfoKeys:", value)
}

func (u_ UserActivity) TargetContentIdentifier() string {
	rv := objc.CallMethod[string](u_, "targetContentIdentifier")
	return rv
}

func (u_ UserActivity) SetTargetContentIdentifier(value string) {
	objc.CallMethod[objc.Void](u_, "setTargetContentIdentifier:", value)
}

func (u_ UserActivity) NeedsSave() bool {
	rv := objc.CallMethod[bool](u_, "needsSave")
	return rv
}

func (u_ UserActivity) SetNeedsSave(value bool) {
	objc.CallMethod[objc.Void](u_, "setNeedsSave:", value)
}

func (u_ UserActivity) Keywords() Set {
	rv := objc.CallMethod[Set](u_, "keywords")
	return rv
}

func (u_ UserActivity) SetKeywords(value ISet) {
	objc.CallMethod[objc.Void](u_, "setKeywords:", value)
}

func (u_ UserActivity) PersistentIdentifier() UserActivityPersistentIdentifier {
	rv := objc.CallMethod[UserActivityPersistentIdentifier](u_, "persistentIdentifier")
	return rv
}

func (u_ UserActivity) SetPersistentIdentifier(value UserActivityPersistentIdentifier) {
	objc.CallMethod[objc.Void](u_, "setPersistentIdentifier:", value)
}

func (u_ UserActivity) IsEligibleForHandoff() bool {
	rv := objc.CallMethod[bool](u_, "isEligibleForHandoff")
	return rv
}

func (u_ UserActivity) SetEligibleForHandoff(value bool) {
	objc.CallMethod[objc.Void](u_, "setEligibleForHandoff:", value)
}

func (u_ UserActivity) IsEligibleForSearch() bool {
	rv := objc.CallMethod[bool](u_, "isEligibleForSearch")
	return rv
}

func (u_ UserActivity) SetEligibleForSearch(value bool) {
	objc.CallMethod[objc.Void](u_, "setEligibleForSearch:", value)
}

func (u_ UserActivity) IsEligibleForPublicIndexing() bool {
	rv := objc.CallMethod[bool](u_, "isEligibleForPublicIndexing")
	return rv
}

func (u_ UserActivity) SetEligibleForPublicIndexing(value bool) {
	objc.CallMethod[objc.Void](u_, "setEligibleForPublicIndexing:", value)
}

func (u_ UserActivity) ExpirationDate() Date {
	rv := objc.CallMethod[Date](u_, "expirationDate")
	return rv
}

func (u_ UserActivity) SetExpirationDate(value IDate) {
	objc.CallMethod[objc.Void](u_, "setExpirationDate:", value)
}

func (u_ UserActivity) Delegate() UserActivityDelegateWrapper {
	rv := objc.CallMethod[UserActivityDelegateWrapper](u_, "delegate")
	return rv
}

func (u_ UserActivity) SetDelegate(value UserActivityDelegate) {
	po := objc.CreateProtocol("NSUserActivityDelegate", value)
	defer po.Release()
	objc.SetAssociatedObject(u_, internal.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](u_, "setDelegate:", po)
}

func (u_ UserActivity) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](u_, "setDelegate:", value)
}

func (u_ UserActivity) SupportsContinuationStreams() bool {
	rv := objc.CallMethod[bool](u_, "supportsContinuationStreams")
	return rv
}

func (u_ UserActivity) SetSupportsContinuationStreams(value bool) {
	objc.CallMethod[objc.Void](u_, "setSupportsContinuationStreams:", value)
}

func (u_ UserActivity) WebpageURL() URL {
	rv := objc.CallMethod[URL](u_, "webpageURL")
	return rv
}

func (u_ UserActivity) SetWebpageURL(value IURL) {
	objc.CallMethod[objc.Void](u_, "setWebpageURL:", value)
}

func (u_ UserActivity) ReferrerURL() URL {
	rv := objc.CallMethod[URL](u_, "referrerURL")
	return rv
}

func (u_ UserActivity) SetReferrerURL(value IURL) {
	objc.CallMethod[objc.Void](u_, "setReferrerURL:", value)
}
