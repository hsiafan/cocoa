// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
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
	rv := ffi.CallMethod[WebsiteDataStore](wc, "alloc")
	return rv
}

func (wc _WebsiteDataStoreClass) New() WebsiteDataStore {
	rv := ffi.CallMethod[WebsiteDataStore](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWebsiteDataStore() WebsiteDataStore {
	return WebsiteDataStoreClass.New()
}

func (w_ WebsiteDataStore) Init() WebsiteDataStore {
	rv := ffi.CallMethod[WebsiteDataStore](w_, "init")
	return rv
}

func (wc _WebsiteDataStoreClass) DefaultDataStore() WebsiteDataStore {
	rv := ffi.CallMethod[WebsiteDataStore](wc, "defaultDataStore")
	return rv
}

func (wc _WebsiteDataStoreClass) NonPersistentDataStore() WebsiteDataStore {
	rv := ffi.CallMethod[WebsiteDataStore](wc, "nonPersistentDataStore")
	return rv
}

func (w_ WebsiteDataStore) FetchDataRecordsOfTypes_CompletionHandler(dataTypes foundation.ISet, completionHandler func(param1 []WebsiteDataRecord)) {
	ffi.CallMethod[ffi.Void](w_, "fetchDataRecordsOfTypes:completionHandler:", dataTypes, completionHandler)
}

func (wc _WebsiteDataStoreClass) AllWebsiteDataTypes() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](wc, "allWebsiteDataTypes")
	return rv
}

func (w_ WebsiteDataStore) RemoveDataOfTypes_ForDataRecords_CompletionHandler(dataTypes foundation.ISet, dataRecords []IWebsiteDataRecord, completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "removeDataOfTypes:forDataRecords:completionHandler:", dataTypes, dataRecords, completionHandler)
}

func (w_ WebsiteDataStore) RemoveDataOfTypes_ModifiedSince_CompletionHandler(dataTypes foundation.ISet, date foundation.IDate, completionHandler func()) {
	ffi.CallMethod[ffi.Void](w_, "removeDataOfTypes:modifiedSince:completionHandler:", dataTypes, date, completionHandler)
}

func (w_ WebsiteDataStore) IsPersistent() bool {
	rv := ffi.CallMethod[bool](w_, "isPersistent")
	return rv
}

func (w_ WebsiteDataStore) HttpCookieStore() HTTPCookieStore {
	rv := ffi.CallMethod[HTTPCookieStore](w_, "httpCookieStore")
	return rv
}
