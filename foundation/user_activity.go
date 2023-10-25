// auto-generated code, do not modify
package foundation

import (
	"unsafe"

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
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
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
	rv := objc.CallMethod[UserActivity](u_, objc.SEL("initWithActivityType:"), activityType)
	return rv
}

func (u_ UserActivity) Init() UserActivity {
	rv := objc.CallMethod[UserActivity](u_, objc.SEL("init"))
	return rv
}

func (uc _UserActivityClass) Alloc() UserActivity {
	rv := objc.CallMethod[UserActivity](uc, objc.SEL("alloc"))
	return rv
}

func (uc _UserActivityClass) New() UserActivity {
	rv := objc.CallMethod[UserActivity](uc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewUserActivity() UserActivity {
	return UserActivityClass.New()
}

func (u_ UserActivity) BecomeCurrent() {
	objc.CallMethod[objc.Void](u_, objc.SEL("becomeCurrent"))
}

func (u_ UserActivity) ResignCurrent() {
	objc.CallMethod[objc.Void](u_, objc.SEL("resignCurrent"))
}

func (u_ UserActivity) Invalidate() {
	objc.CallMethod[objc.Void](u_, objc.SEL("invalidate"))
}

func (uc _UserActivityClass) DeleteAllSavedUserActivitiesWithCompletionHandler(handler func()) {
	objc.CallMethod[objc.Void](uc, objc.SEL("deleteAllSavedUserActivitiesWithCompletionHandler:"), handler)
}

func (uc _UserActivityClass) DeleteSavedUserActivitiesWithPersistentIdentifiers_CompletionHandler(persistentIdentifiers []UserActivityPersistentIdentifier, handler func()) {
	objc.CallMethod[objc.Void](uc, objc.SEL("deleteSavedUserActivitiesWithPersistentIdentifiers:completionHandler:"), persistentIdentifiers, handler)
}

func (u_ UserActivity) GetContinuationStreamsWithCompletionHandler(completionHandler func(inputStream InputStream, outputStream OutputStream, error Error)) {
	objc.CallMethod[objc.Void](u_, objc.SEL("getContinuationStreamsWithCompletionHandler:"), completionHandler)
}

func (u_ UserActivity) ActivityType() string {
	rv := objc.CallMethod[string](u_, objc.SEL("activityType"))
	return rv
}

func (u_ UserActivity) Title() string {
	rv := objc.CallMethod[string](u_, objc.SEL("title"))
	return rv
}

func (u_ UserActivity) SetTitle(value string) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setTitle:"), value)
}

func (u_ UserActivity) RequiredUserInfoKeys() Set {
	rv := objc.CallMethod[Set](u_, objc.SEL("requiredUserInfoKeys"))
	return rv
}

func (u_ UserActivity) SetRequiredUserInfoKeys(value ISet) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setRequiredUserInfoKeys:"), objc.ExtractPtr(value))
}

func (u_ UserActivity) TargetContentIdentifier() string {
	rv := objc.CallMethod[string](u_, objc.SEL("targetContentIdentifier"))
	return rv
}

func (u_ UserActivity) SetTargetContentIdentifier(value string) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setTargetContentIdentifier:"), value)
}

// weak property
func (u_ UserActivity) NeedsSave() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("needsSave"))
	return rv
}

// weak property
func (u_ UserActivity) SetNeedsSave(value bool) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setNeedsSave:"), value)
}

func (u_ UserActivity) Keywords() Set {
	rv := objc.CallMethod[Set](u_, objc.SEL("keywords"))
	return rv
}

func (u_ UserActivity) SetKeywords(value ISet) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setKeywords:"), objc.ExtractPtr(value))
}

func (u_ UserActivity) PersistentIdentifier() UserActivityPersistentIdentifier {
	rv := objc.CallMethod[UserActivityPersistentIdentifier](u_, objc.SEL("persistentIdentifier"))
	return rv
}

func (u_ UserActivity) SetPersistentIdentifier(value UserActivityPersistentIdentifier) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setPersistentIdentifier:"), value)
}

func (u_ UserActivity) IsEligibleForHandoff() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("isEligibleForHandoff"))
	return rv
}

func (u_ UserActivity) SetEligibleForHandoff(value bool) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setEligibleForHandoff:"), value)
}

func (u_ UserActivity) IsEligibleForSearch() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("isEligibleForSearch"))
	return rv
}

func (u_ UserActivity) SetEligibleForSearch(value bool) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setEligibleForSearch:"), value)
}

func (u_ UserActivity) IsEligibleForPublicIndexing() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("isEligibleForPublicIndexing"))
	return rv
}

func (u_ UserActivity) SetEligibleForPublicIndexing(value bool) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setEligibleForPublicIndexing:"), value)
}

func (u_ UserActivity) ExpirationDate() Date {
	rv := objc.CallMethod[Date](u_, objc.SEL("expirationDate"))
	return rv
}

func (u_ UserActivity) SetExpirationDate(value IDate) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setExpirationDate:"), objc.ExtractPtr(value))
}

// weak property
func (u_ UserActivity) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](u_, objc.SEL("delegate"))
	return rv
}

// weak property
func (u_ UserActivity) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (u_ UserActivity) SupportsContinuationStreams() bool {
	rv := objc.CallMethod[bool](u_, objc.SEL("supportsContinuationStreams"))
	return rv
}

func (u_ UserActivity) SetSupportsContinuationStreams(value bool) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setSupportsContinuationStreams:"), value)
}

func (u_ UserActivity) WebpageURL() URL {
	rv := objc.CallMethod[URL](u_, objc.SEL("webpageURL"))
	return rv
}

func (u_ UserActivity) SetWebpageURL(value IURL) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setWebpageURL:"), objc.ExtractPtr(value))
}

func (u_ UserActivity) ReferrerURL() URL {
	rv := objc.CallMethod[URL](u_, objc.SEL("referrerURL"))
	return rv
}

func (u_ UserActivity) SetReferrerURL(value IURL) {
	objc.CallMethod[objc.Void](u_, objc.SEL("setReferrerURL:"), objc.ExtractPtr(value))
}
