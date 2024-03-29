// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var WebsiteDataStoreClass = _WebsiteDataStoreClass{objc.GetClass("WKWebsiteDataStore")}

type _WebsiteDataStoreClass struct {
	objc.Class
}

type IWebsiteDataStore interface {
	objc.IObject
	FetchDataRecordsOfTypes_CompletionHandler(dataTypes foundation.ISet, completionHandler func(param1 []WebsiteDataRecord))
	RemoveDataOfTypes_ForDataRecords_CompletionHandler(dataTypes foundation.ISet, dataRecords []IWebsiteDataRecord, completionHandler func())
	RemoveDataOfTypes_ModifiedSince_CompletionHandler(dataTypes foundation.ISet, date foundation.IDate, completionHandler func())
	IsPersistent() bool
	HttpCookieStore() HTTPCookieStore
}

type WebsiteDataStore struct {
	objc.Object
}

func MakeWebsiteDataStore(ptr unsafe.Pointer) WebsiteDataStore {
	return WebsiteDataStore{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WebsiteDataStoreClass) Alloc() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](wc, objc.SEL("alloc"))
	return rv
}

func (wc _WebsiteDataStoreClass) New() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](wc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewWebsiteDataStore() WebsiteDataStore {
	return WebsiteDataStoreClass.New()
}

func (w_ WebsiteDataStore) Init() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](w_, objc.SEL("init"))
	return rv
}

func (wc _WebsiteDataStoreClass) DefaultDataStore() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](wc, objc.SEL("defaultDataStore"))
	return rv
}

func (wc _WebsiteDataStoreClass) NonPersistentDataStore() WebsiteDataStore {
	rv := objc.CallMethod[WebsiteDataStore](wc, objc.SEL("nonPersistentDataStore"))
	return rv
}

func (w_ WebsiteDataStore) FetchDataRecordsOfTypes_CompletionHandler(dataTypes foundation.ISet, completionHandler func(param1 []WebsiteDataRecord)) {
	objc.CallMethod[objc.Void](w_, objc.SEL("fetchDataRecordsOfTypes:completionHandler:"), objc.ExtractPtr(dataTypes), completionHandler)
}

func (wc _WebsiteDataStoreClass) AllWebsiteDataTypes() foundation.Set {
	rv := objc.CallMethod[foundation.Set](wc, objc.SEL("allWebsiteDataTypes"))
	return rv
}

func (w_ WebsiteDataStore) RemoveDataOfTypes_ForDataRecords_CompletionHandler(dataTypes foundation.ISet, dataRecords []IWebsiteDataRecord, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("removeDataOfTypes:forDataRecords:completionHandler:"), objc.ExtractPtr(dataTypes), dataRecords, completionHandler)
}

func (w_ WebsiteDataStore) RemoveDataOfTypes_ModifiedSince_CompletionHandler(dataTypes foundation.ISet, date foundation.IDate, completionHandler func()) {
	objc.CallMethod[objc.Void](w_, objc.SEL("removeDataOfTypes:modifiedSince:completionHandler:"), objc.ExtractPtr(dataTypes), objc.ExtractPtr(date), completionHandler)
}

func (w_ WebsiteDataStore) IsPersistent() bool {
	rv := objc.CallMethod[bool](w_, objc.SEL("isPersistent"))
	return rv
}

func (w_ WebsiteDataStore) HttpCookieStore() HTTPCookieStore {
	rv := objc.CallMethod[HTTPCookieStore](w_, objc.SEL("httpCookieStore"))
	return rv
}
