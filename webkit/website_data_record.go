// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var WebsiteDataRecordClass = _WebsiteDataRecordClass{objc.GetClass("WKWebsiteDataRecord")}

type _WebsiteDataRecordClass struct {
	objc.Class
}

type IWebsiteDataRecord interface {
	objc.IObject
	DisplayName() string
	DataTypes() foundation.Set
}

type WebsiteDataRecord struct {
	objc.Object
}

func MakeWebsiteDataRecord(ptr unsafe.Pointer) WebsiteDataRecord {
	return WebsiteDataRecord{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WebsiteDataRecordClass) Alloc() WebsiteDataRecord {
	rv := ffi.CallMethod[WebsiteDataRecord](wc, "alloc")
	return rv
}

func (wc _WebsiteDataRecordClass) New() WebsiteDataRecord {
	rv := ffi.CallMethod[WebsiteDataRecord](wc, "new")
	rv.Autorelease()
	return rv
}

func NewWebsiteDataRecord() WebsiteDataRecord {
	return WebsiteDataRecordClass.New()
}

func (w_ WebsiteDataRecord) Init() WebsiteDataRecord {
	rv := ffi.CallMethod[WebsiteDataRecord](w_, "init")
	return rv
}

func (w_ WebsiteDataRecord) DisplayName() string {
	rv := ffi.CallMethod[string](w_, "displayName")
	return rv
}

func (w_ WebsiteDataRecord) DataTypes() foundation.Set {
	rv := ffi.CallMethod[foundation.Set](w_, "dataTypes")
	return rv
}
