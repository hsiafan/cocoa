// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var TextContentManagerClass = _TextContentManagerClass{objc.GetClass("NSTextContentManager")}

type _TextContentManagerClass struct {
	objc.Class
}

type ITextContentManager interface {
	objc.IObject
	PerformEditingTransactionUsingBlock(transaction func())
	RecordEditActionInRange_NewTextRange(originalTextRange ITextRange, newTextRange ITextRange)
	AddTextLayoutManager(textLayoutManager ITextLayoutManager)
	RemoveTextLayoutManager(textLayoutManager ITextLayoutManager)
	SynchronizeTextLayoutManagers(completionHandler func(error foundation.Error))
	TextElementsForRange(range_ ITextRange) []TextElement
	AutomaticallySynchronizesToBackingStore() bool
	SetAutomaticallySynchronizesToBackingStore(value bool)
	HasEditingTransaction() bool
	PrimaryTextLayoutManager() TextLayoutManager
	SetPrimaryTextLayoutManager(value ITextLayoutManager)
	TextLayoutManagers() []TextLayoutManager
	AutomaticallySynchronizesTextLayoutManagers() bool
	SetAutomaticallySynchronizesTextLayoutManagers(value bool)
}

type TextContentManager struct {
	objc.Object
}

func MakeTextContentManager(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextContentManager) Init() TextContentManager {
	rv := ffi.CallMethod[TextContentManager](t_, "init")
	return rv
}

func (tc _TextContentManagerClass) Alloc() TextContentManager {
	rv := ffi.CallMethod[TextContentManager](tc, "alloc")
	return rv
}

func (tc _TextContentManagerClass) New() TextContentManager {
	rv := ffi.CallMethod[TextContentManager](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextContentManager() TextContentManager {
	return TextContentManagerClass.New()
}

func (t_ TextContentManager) PerformEditingTransactionUsingBlock(transaction func()) {
	ffi.CallMethod[ffi.Void](t_, "performEditingTransactionUsingBlock:", transaction)
}

func (t_ TextContentManager) RecordEditActionInRange_NewTextRange(originalTextRange ITextRange, newTextRange ITextRange) {
	ffi.CallMethod[ffi.Void](t_, "recordEditActionInRange:newTextRange:", originalTextRange, newTextRange)
}

func (t_ TextContentManager) AddTextLayoutManager(textLayoutManager ITextLayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "addTextLayoutManager:", textLayoutManager)
}

func (t_ TextContentManager) RemoveTextLayoutManager(textLayoutManager ITextLayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "removeTextLayoutManager:", textLayoutManager)
}

func (t_ TextContentManager) SynchronizeTextLayoutManagers(completionHandler func(error foundation.Error)) {
	ffi.CallMethod[ffi.Void](t_, "synchronizeTextLayoutManagers:", completionHandler)
}

func (t_ TextContentManager) TextElementsForRange(range_ ITextRange) []TextElement {
	rv := ffi.CallMethod[[]TextElement](t_, "textElementsForRange:", range_)
	return rv
}

func (t_ TextContentManager) AutomaticallySynchronizesToBackingStore() bool {
	rv := ffi.CallMethod[bool](t_, "automaticallySynchronizesToBackingStore")
	return rv
}

func (t_ TextContentManager) SetAutomaticallySynchronizesToBackingStore(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticallySynchronizesToBackingStore:", value)
}

func (t_ TextContentManager) HasEditingTransaction() bool {
	rv := ffi.CallMethod[bool](t_, "hasEditingTransaction")
	return rv
}

func (t_ TextContentManager) PrimaryTextLayoutManager() TextLayoutManager {
	rv := ffi.CallMethod[TextLayoutManager](t_, "primaryTextLayoutManager")
	return rv
}

func (t_ TextContentManager) SetPrimaryTextLayoutManager(value ITextLayoutManager) {
	ffi.CallMethod[ffi.Void](t_, "setPrimaryTextLayoutManager:", value)
}

func (t_ TextContentManager) TextLayoutManagers() []TextLayoutManager {
	rv := ffi.CallMethod[[]TextLayoutManager](t_, "textLayoutManagers")
	return rv
}

func (t_ TextContentManager) AutomaticallySynchronizesTextLayoutManagers() bool {
	rv := ffi.CallMethod[bool](t_, "automaticallySynchronizesTextLayoutManagers")
	return rv
}

func (t_ TextContentManager) SetAutomaticallySynchronizesTextLayoutManagers(value bool) {
	ffi.CallMethod[ffi.Void](t_, "setAutomaticallySynchronizesTextLayoutManagers:", value)
}
