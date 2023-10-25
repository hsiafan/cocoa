// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var ContentRuleListClass = _ContentRuleListClass{objc.GetClass("WKContentRuleList")}

type _ContentRuleListClass struct {
	objc.Class
}

type IContentRuleList interface {
	objc.IObject
	Identifier() string
}

type ContentRuleList struct {
	objc.Object
}

func MakeContentRuleList(ptr unsafe.Pointer) ContentRuleList {
	return ContentRuleList{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ContentRuleListClass) Alloc() ContentRuleList {
	rv := objc.CallMethod[ContentRuleList](cc, objc.SEL("alloc"))
	return rv
}

func (cc _ContentRuleListClass) New() ContentRuleList {
	rv := objc.CallMethod[ContentRuleList](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewContentRuleList() ContentRuleList {
	return ContentRuleListClass.New()
}

func (c_ ContentRuleList) Init() ContentRuleList {
	rv := objc.CallMethod[ContentRuleList](c_, objc.SEL("init"))
	return rv
}

func (c_ ContentRuleList) Identifier() string {
	rv := objc.CallMethod[string](c_, objc.SEL("identifier"))
	return rv
}
