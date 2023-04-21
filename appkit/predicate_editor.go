// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
)

var PredicateEditorClass = _PredicateEditorClass{objc.GetClass("NSPredicateEditor")}

type _PredicateEditorClass struct {
	objc.Class
}

type IPredicateEditor interface {
	IRuleEditor
	RowTemplates() []PredicateEditorRowTemplate
	SetRowTemplates(value []IPredicateEditorRowTemplate)
}

type PredicateEditor struct {
	RuleEditor
}

func MakePredicateEditor(ptr unsafe.Pointer) PredicateEditor {
	return PredicateEditor{
		RuleEditor: MakeRuleEditor(ptr),
	}
}

func (p_ PredicateEditor) InitWithFrame(frameRect foundation.Rect) PredicateEditor {
	rv := objc.CallMethod[PredicateEditor](p_, "initWithFrame:", frameRect)
	return rv
}

func (p_ PredicateEditor) Init() PredicateEditor {
	rv := objc.CallMethod[PredicateEditor](p_, "init")
	return rv
}

func (pc _PredicateEditorClass) Alloc() PredicateEditor {
	rv := objc.CallMethod[PredicateEditor](pc, "alloc")
	return rv
}

func (pc _PredicateEditorClass) New() PredicateEditor {
	rv := objc.CallMethod[PredicateEditor](pc, "new")
	rv.Autorelease()
	return rv
}

func NewPredicateEditor() PredicateEditor {
	return PredicateEditorClass.New()
}

func (p_ PredicateEditor) RowTemplates() []PredicateEditorRowTemplate {
	rv := objc.CallMethod[[]PredicateEditorRowTemplate](p_, "rowTemplates")
	return rv
}

func (p_ PredicateEditor) SetRowTemplates(value []IPredicateEditorRowTemplate) {
	objc.CallMethod[objc.Void](p_, "setRowTemplates:", value)
}
