// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var UserDefaultsClass = _UserDefaultsClass{objc.GetClass("NSUserDefaults")}

type _UserDefaultsClass struct {
	objc.Class
}

type IUserDefaults interface {
	objc.IObject
	ObjectForKey(defaultName string) objc.Object
	URLForKey(defaultName string) URL
	ArrayForKey(defaultName string) []objc.Object
	DictionaryForKey(defaultName string) map[string]objc.Object
	StringForKey(defaultName string) string
	StringArrayForKey(defaultName string) []string
	DataForKey(defaultName string) []byte
	BoolForKey(defaultName string) bool
	IntegerForKey(defaultName string) int
	FloatForKey(defaultName string) float32
	DoubleForKey(defaultName string) float64
	DictionaryRepresentation() map[string]objc.Object
	SetObject_ForKey(value objc.IObject, defaultName string)
	SetFloat_ForKey(value float32, defaultName string)
	SetDouble_ForKey(value float64, defaultName string)
	SetInteger_ForKey(value int, defaultName string)
	SetBool_ForKey(value bool, defaultName string)
	SetURL_ForKey(url IURL, defaultName string)
	RemoveObjectForKey(defaultName string)
	AddSuiteNamed(suiteName string)
	RemoveSuiteNamed(suiteName string)
	RegisterDefaults(registrationDictionary map[string]objc.IObject)
	PersistentDomainForName(domainName string) map[string]objc.Object
	SetPersistentDomain_ForName(domain map[string]objc.IObject, domainName string)
	RemovePersistentDomainForName(domainName string)
	// deprecated
	PersistentDomainNames() []objc.Object
	VolatileDomainForName(domainName string) map[string]objc.Object
	SetVolatileDomain_ForName(domain map[string]objc.IObject, domainName string)
	RemoveVolatileDomainForName(domainName string)
	ObjectIsForcedForKey(key string) bool
	ObjectIsForcedForKey_InDomain(key string, domain string) bool
	// deprecated
	InitWithUser(username string) objc.Object
	Synchronize() bool
	VolatileDomainNames() []string
}

type UserDefaults struct {
	objc.Object
}

func MakeUserDefaults(ptr unsafe.Pointer) UserDefaults {
	return UserDefaults{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ UserDefaults) Init() UserDefaults {
	rv := objc.CallMethod[UserDefaults](u_, objc.GetSelector("init"))
	return rv
}

func (u_ UserDefaults) InitWithSuiteName(suitename string) UserDefaults {
	rv := objc.CallMethod[UserDefaults](u_, objc.GetSelector("initWithSuiteName:"), suitename)
	return rv
}

func (uc _UserDefaultsClass) Alloc() UserDefaults {
	rv := objc.CallMethod[UserDefaults](uc, objc.GetSelector("alloc"))
	return rv
}

func (uc _UserDefaultsClass) New() UserDefaults {
	rv := objc.CallMethod[UserDefaults](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewUserDefaults() UserDefaults {
	return UserDefaultsClass.New()
}

func (u_ UserDefaults) ObjectForKey(defaultName string) objc.Object {
	rv := objc.CallMethod[objc.Object](u_, objc.GetSelector("objectForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) URLForKey(defaultName string) URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("URLForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) ArrayForKey(defaultName string) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](u_, objc.GetSelector("arrayForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) DictionaryForKey(defaultName string) map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](u_, objc.GetSelector("dictionaryForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) StringForKey(defaultName string) string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("stringForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) StringArrayForKey(defaultName string) []string {
	rv := objc.CallMethod[[]string](u_, objc.GetSelector("stringArrayForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) DataForKey(defaultName string) []byte {
	rv := objc.CallMethod[[]byte](u_, objc.GetSelector("dataForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) BoolForKey(defaultName string) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("boolForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) IntegerForKey(defaultName string) int {
	rv := objc.CallMethod[int](u_, objc.GetSelector("integerForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) FloatForKey(defaultName string) float32 {
	rv := objc.CallMethod[float32](u_, objc.GetSelector("floatForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) DoubleForKey(defaultName string) float64 {
	rv := objc.CallMethod[float64](u_, objc.GetSelector("doubleForKey:"), defaultName)
	return rv
}

func (u_ UserDefaults) DictionaryRepresentation() map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](u_, objc.GetSelector("dictionaryRepresentation"))
	return rv
}

func (u_ UserDefaults) SetObject_ForKey(value objc.IObject, defaultName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setObject:forKey:"), value, defaultName)
}

func (u_ UserDefaults) SetFloat_ForKey(value float32, defaultName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setFloat:forKey:"), value, defaultName)
}

func (u_ UserDefaults) SetDouble_ForKey(value float64, defaultName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setDouble:forKey:"), value, defaultName)
}

func (u_ UserDefaults) SetInteger_ForKey(value int, defaultName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setInteger:forKey:"), value, defaultName)
}

func (u_ UserDefaults) SetBool_ForKey(value bool, defaultName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setBool:forKey:"), value, defaultName)
}

func (u_ UserDefaults) SetURL_ForKey(url IURL, defaultName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setURL:forKey:"), url, defaultName)
}

func (u_ UserDefaults) RemoveObjectForKey(defaultName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeObjectForKey:"), defaultName)
}

func (u_ UserDefaults) AddSuiteNamed(suiteName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("addSuiteNamed:"), suiteName)
}

func (u_ UserDefaults) RemoveSuiteNamed(suiteName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeSuiteNamed:"), suiteName)
}

func (u_ UserDefaults) RegisterDefaults(registrationDictionary map[string]objc.IObject) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("registerDefaults:"), registrationDictionary)
}

func (u_ UserDefaults) PersistentDomainForName(domainName string) map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](u_, objc.GetSelector("persistentDomainForName:"), domainName)
	return rv
}

func (u_ UserDefaults) SetPersistentDomain_ForName(domain map[string]objc.IObject, domainName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setPersistentDomain:forName:"), domain, domainName)
}

func (u_ UserDefaults) RemovePersistentDomainForName(domainName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removePersistentDomainForName:"), domainName)
}

// deprecated
func (u_ UserDefaults) PersistentDomainNames() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](u_, objc.GetSelector("persistentDomainNames"))
	return rv
}

func (u_ UserDefaults) VolatileDomainForName(domainName string) map[string]objc.Object {
	rv := objc.CallMethod[map[string]objc.Object](u_, objc.GetSelector("volatileDomainForName:"), domainName)
	return rv
}

func (u_ UserDefaults) SetVolatileDomain_ForName(domain map[string]objc.IObject, domainName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setVolatileDomain:forName:"), domain, domainName)
}

func (u_ UserDefaults) RemoveVolatileDomainForName(domainName string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("removeVolatileDomainForName:"), domainName)
}

func (u_ UserDefaults) ObjectIsForcedForKey(key string) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("objectIsForcedForKey:"), key)
	return rv
}

func (u_ UserDefaults) ObjectIsForcedForKey_InDomain(key string, domain string) bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("objectIsForcedForKey:inDomain:"), key, domain)
	return rv
}

// deprecated
func (u_ UserDefaults) InitWithUser(username string) objc.Object {
	rv := objc.CallMethod[objc.Object](u_, objc.GetSelector("initWithUser:"), username)
	return rv
}

func (u_ UserDefaults) Synchronize() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("synchronize"))
	return rv
}

func (uc _UserDefaultsClass) ResetStandardUserDefaults() {
	objc.CallMethod[objc.Void](uc, objc.GetSelector("resetStandardUserDefaults"))
}

func (uc _UserDefaultsClass) StandardUserDefaults() UserDefaults {
	rv := objc.CallMethod[UserDefaults](uc, objc.GetSelector("standardUserDefaults"))
	return rv
}

func (u_ UserDefaults) VolatileDomainNames() []string {
	rv := objc.CallMethod[[]string](u_, objc.GetSelector("volatileDomainNames"))
	return rv
}
