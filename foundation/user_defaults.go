// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[UserDefaults](u_, "init")
	return rv
}

func (u_ UserDefaults) InitWithSuiteName(suitename string) UserDefaults {
	rv := ffi.CallMethod[UserDefaults](u_, "initWithSuiteName:", suitename)
	return rv
}

func (uc _UserDefaultsClass) Alloc() UserDefaults {
	rv := ffi.CallMethod[UserDefaults](uc, "alloc")
	return rv
}

func (uc _UserDefaultsClass) New() UserDefaults {
	rv := ffi.CallMethod[UserDefaults](uc, "new")
	rv.Autorelease()
	return rv
}

func NewUserDefaults() UserDefaults {
	return UserDefaultsClass.New()
}

func (u_ UserDefaults) ObjectForKey(defaultName string) objc.Object {
	rv := ffi.CallMethod[objc.Object](u_, "objectForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) URLForKey(defaultName string) URL {
	rv := ffi.CallMethod[URL](u_, "URLForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) ArrayForKey(defaultName string) []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](u_, "arrayForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) DictionaryForKey(defaultName string) map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](u_, "dictionaryForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) StringForKey(defaultName string) string {
	rv := ffi.CallMethod[string](u_, "stringForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) StringArrayForKey(defaultName string) []string {
	rv := ffi.CallMethod[[]string](u_, "stringArrayForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) DataForKey(defaultName string) []byte {
	rv := ffi.CallMethod[[]byte](u_, "dataForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) BoolForKey(defaultName string) bool {
	rv := ffi.CallMethod[bool](u_, "boolForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) IntegerForKey(defaultName string) int {
	rv := ffi.CallMethod[int](u_, "integerForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) FloatForKey(defaultName string) float32 {
	rv := ffi.CallMethod[float32](u_, "floatForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) DoubleForKey(defaultName string) float64 {
	rv := ffi.CallMethod[float64](u_, "doubleForKey:", defaultName)
	return rv
}

func (u_ UserDefaults) DictionaryRepresentation() map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](u_, "dictionaryRepresentation")
	return rv
}

func (u_ UserDefaults) SetObject_ForKey(value objc.IObject, defaultName string) {
	ffi.CallMethod[ffi.Void](u_, "setObject:forKey:", value, defaultName)
}

func (u_ UserDefaults) SetFloat_ForKey(value float32, defaultName string) {
	ffi.CallMethod[ffi.Void](u_, "setFloat:forKey:", value, defaultName)
}

func (u_ UserDefaults) SetDouble_ForKey(value float64, defaultName string) {
	ffi.CallMethod[ffi.Void](u_, "setDouble:forKey:", value, defaultName)
}

func (u_ UserDefaults) SetInteger_ForKey(value int, defaultName string) {
	ffi.CallMethod[ffi.Void](u_, "setInteger:forKey:", value, defaultName)
}

func (u_ UserDefaults) SetBool_ForKey(value bool, defaultName string) {
	ffi.CallMethod[ffi.Void](u_, "setBool:forKey:", value, defaultName)
}

func (u_ UserDefaults) SetURL_ForKey(url IURL, defaultName string) {
	ffi.CallMethod[ffi.Void](u_, "setURL:forKey:", url, defaultName)
}

func (u_ UserDefaults) RemoveObjectForKey(defaultName string) {
	ffi.CallMethod[ffi.Void](u_, "removeObjectForKey:", defaultName)
}

func (u_ UserDefaults) AddSuiteNamed(suiteName string) {
	ffi.CallMethod[ffi.Void](u_, "addSuiteNamed:", suiteName)
}

func (u_ UserDefaults) RemoveSuiteNamed(suiteName string) {
	ffi.CallMethod[ffi.Void](u_, "removeSuiteNamed:", suiteName)
}

func (u_ UserDefaults) RegisterDefaults(registrationDictionary map[string]objc.IObject) {
	ffi.CallMethod[ffi.Void](u_, "registerDefaults:", registrationDictionary)
}

func (u_ UserDefaults) PersistentDomainForName(domainName string) map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](u_, "persistentDomainForName:", domainName)
	return rv
}

func (u_ UserDefaults) SetPersistentDomain_ForName(domain map[string]objc.IObject, domainName string) {
	ffi.CallMethod[ffi.Void](u_, "setPersistentDomain:forName:", domain, domainName)
}

func (u_ UserDefaults) RemovePersistentDomainForName(domainName string) {
	ffi.CallMethod[ffi.Void](u_, "removePersistentDomainForName:", domainName)
}

// deprecated
func (u_ UserDefaults) PersistentDomainNames() []objc.Object {
	rv := ffi.CallMethod[[]objc.Object](u_, "persistentDomainNames")
	return rv
}

func (u_ UserDefaults) VolatileDomainForName(domainName string) map[string]objc.Object {
	rv := ffi.CallMethod[map[string]objc.Object](u_, "volatileDomainForName:", domainName)
	return rv
}

func (u_ UserDefaults) SetVolatileDomain_ForName(domain map[string]objc.IObject, domainName string) {
	ffi.CallMethod[ffi.Void](u_, "setVolatileDomain:forName:", domain, domainName)
}

func (u_ UserDefaults) RemoveVolatileDomainForName(domainName string) {
	ffi.CallMethod[ffi.Void](u_, "removeVolatileDomainForName:", domainName)
}

func (u_ UserDefaults) ObjectIsForcedForKey(key string) bool {
	rv := ffi.CallMethod[bool](u_, "objectIsForcedForKey:", key)
	return rv
}

func (u_ UserDefaults) ObjectIsForcedForKey_InDomain(key string, domain string) bool {
	rv := ffi.CallMethod[bool](u_, "objectIsForcedForKey:inDomain:", key, domain)
	return rv
}

// deprecated
func (u_ UserDefaults) InitWithUser(username string) objc.Object {
	rv := ffi.CallMethod[objc.Object](u_, "initWithUser:", username)
	return rv
}

func (u_ UserDefaults) Synchronize() bool {
	rv := ffi.CallMethod[bool](u_, "synchronize")
	return rv
}

func (uc _UserDefaultsClass) ResetStandardUserDefaults() {
	ffi.CallMethod[ffi.Void](uc, "resetStandardUserDefaults")
}

func (uc _UserDefaultsClass) StandardUserDefaults() UserDefaults {
	rv := ffi.CallMethod[UserDefaults](uc, "standardUserDefaults")
	return rv
}

func (u_ UserDefaults) VolatileDomainNames() []string {
	rv := ffi.CallMethod[[]string](u_, "volatileDomainNames")
	return rv
}
