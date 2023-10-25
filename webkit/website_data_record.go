// auto-generated code, do not modify
package webkit

import (
	"unsafe"

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
	rv := objc.CallMethod[WebsiteDataRecord](wc, objc.SEL("alloc"))
	return rv
}

func (wc _WebsiteDataRecordClass) New() WebsiteDataRecord {
	rv := objc.CallMethod[WebsiteDataRecord](wc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewWebsiteDataRecord() WebsiteDataRecord {
	return WebsiteDataRecordClass.New()
}

func (w_ WebsiteDataRecord) Init() WebsiteDataRecord {
	rv := objc.CallMethod[WebsiteDataRecord](w_, objc.SEL("init"))
	return rv
}

func (w_ WebsiteDataRecord) DisplayName() string {
	rv := objc.CallMethod[string](w_, objc.SEL("displayName"))
	return rv
}

func (w_ WebsiteDataRecord) DataTypes() foundation.Set {
	rv := objc.CallMethod[foundation.Set](w_, objc.SEL("dataTypes"))
	return rv
}
