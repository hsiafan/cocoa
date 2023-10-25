// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var CandidateListTouchBarItemClass = _CandidateListTouchBarItemClass{objc.GetClass("NSCandidateListTouchBarItem")}

type _CandidateListTouchBarItemClass struct {
	objc.Class
}

type ICandidateListTouchBarItem interface {
	ITouchBarItem
	UpdateWithInsertionPointVisibility(isVisible bool)
	Client() View
	SetClient(value IView)
	Delegate() objc.Object
	SetDelegate(value objc.IObject)
	AllowsTextInputContextCandidates() bool
	SetAllowsTextInputContextCandidates(value bool)
	AllowsCollapsing() bool
	SetAllowsCollapsing(value bool)
	IsCollapsed() bool
	SetCollapsed(value bool)
	IsCandidateListVisible() bool
	SetCustomizationLabel(value string)
}

type CandidateListTouchBarItem struct {
	TouchBarItem
}

func MakeCandidateListTouchBarItem(ptr unsafe.Pointer) CandidateListTouchBarItem {
	return CandidateListTouchBarItem{
		TouchBarItem: MakeTouchBarItem(ptr),
	}
}

func (c_ CandidateListTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) CandidateListTouchBarItem {
	rv := objc.CallMethod[CandidateListTouchBarItem](c_, objc.SEL("initWithIdentifier:"), identifier)
	return rv
}

func (cc _CandidateListTouchBarItemClass) Alloc() CandidateListTouchBarItem {
	rv := objc.CallMethod[CandidateListTouchBarItem](cc, objc.SEL("alloc"))
	return rv
}

func (cc _CandidateListTouchBarItemClass) New() CandidateListTouchBarItem {
	rv := objc.CallMethod[CandidateListTouchBarItem](cc, objc.SEL("new"))
	rv.Autorelease()
	return rv
}

func NewCandidateListTouchBarItem() CandidateListTouchBarItem {
	return CandidateListTouchBarItemClass.New()
}

func (c_ CandidateListTouchBarItem) Init() CandidateListTouchBarItem {
	rv := objc.CallMethod[CandidateListTouchBarItem](c_, objc.SEL("init"))
	return rv
}

func (c_ CandidateListTouchBarItem) UpdateWithInsertionPointVisibility(isVisible bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("updateWithInsertionPointVisibility:"), isVisible)
}

// weak property
func (c_ CandidateListTouchBarItem) Client() View {
	rv := objc.CallMethod[View](c_, objc.SEL("client"))
	return rv
}

// weak property
func (c_ CandidateListTouchBarItem) SetClient(value IView) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setClient:"), objc.ExtractPtr(value))
}

// weak property
func (c_ CandidateListTouchBarItem) Delegate() objc.Object {
	rv := objc.CallMethod[objc.Object](c_, objc.SEL("delegate"))
	return rv
}

// weak property
func (c_ CandidateListTouchBarItem) SetDelegate(value objc.IObject) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setDelegate:"), objc.ExtractPtr(value))
}

func (c_ CandidateListTouchBarItem) AllowsTextInputContextCandidates() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("allowsTextInputContextCandidates"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetAllowsTextInputContextCandidates(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAllowsTextInputContextCandidates:"), value)
}

func (c_ CandidateListTouchBarItem) AllowsCollapsing() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("allowsCollapsing"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetAllowsCollapsing(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setAllowsCollapsing:"), value)
}

func (c_ CandidateListTouchBarItem) IsCollapsed() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isCollapsed"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetCollapsed(value bool) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setCollapsed:"), value)
}

func (c_ CandidateListTouchBarItem) IsCandidateListVisible() bool {
	rv := objc.CallMethod[bool](c_, objc.SEL("isCandidateListVisible"))
	return rv
}

func (c_ CandidateListTouchBarItem) SetCustomizationLabel(value string) {
	objc.CallMethod[objc.Void](c_, objc.SEL("setCustomizationLabel:"), value)
}
