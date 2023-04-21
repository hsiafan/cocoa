// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/hsiafan/cocoa/objc"
)

var TextSelectionNavigationClass = _TextSelectionNavigationClass{objc.GetClass("NSTextSelectionNavigation")}

type _TextSelectionNavigationClass struct {
	objc.Class
}

type ITextSelectionNavigation interface {
	objc.IObject
	TextSelectionForSelectionGranularity_EnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) TextSelection
	DestinationSelectionForTextSelection_Direction_Destination_Extending_Confined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) TextSelection
	FlushLayoutCache()
	DeletionRangesForTextSelection_Direction_Destination_AllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange
	AllowsNonContiguousRanges() bool
	SetAllowsNonContiguousRanges(value bool)
	RotatesCoordinateSystemForLayoutOrientation() bool
	SetRotatesCoordinateSystemForLayoutOrientation(value bool)
}

type TextSelectionNavigation struct {
	objc.Object
}

func MakeTextSelectionNavigation(ptr unsafe.Pointer) TextSelectionNavigation {
	return TextSelectionNavigation{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TextSelectionNavigationClass) Alloc() TextSelectionNavigation {
	rv := objc.CallMethod[TextSelectionNavigation](tc, "alloc")
	return rv
}

func (tc _TextSelectionNavigationClass) New() TextSelectionNavigation {
	rv := objc.CallMethod[TextSelectionNavigation](tc, "new")
	rv.Autorelease()
	return rv
}

func NewTextSelectionNavigation() TextSelectionNavigation {
	return TextSelectionNavigationClass.New()
}

func (t_ TextSelectionNavigation) Init() TextSelectionNavigation {
	rv := objc.CallMethod[TextSelectionNavigation](t_, "init")
	return rv
}

func (t_ TextSelectionNavigation) TextSelectionForSelectionGranularity_EnclosingTextSelection(selectionGranularity TextSelectionGranularity, textSelection ITextSelection) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, "textSelectionForSelectionGranularity:enclosingTextSelection:", selectionGranularity, textSelection)
	return rv
}

func (t_ TextSelectionNavigation) DestinationSelectionForTextSelection_Direction_Destination_Extending_Confined(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, extending bool, confined bool) TextSelection {
	rv := objc.CallMethod[TextSelection](t_, "destinationSelectionForTextSelection:direction:destination:extending:confined:", textSelection, direction, destination, extending, confined)
	return rv
}

func (t_ TextSelectionNavigation) FlushLayoutCache() {
	objc.CallMethod[objc.Void](t_, "flushLayoutCache")
}

func (t_ TextSelectionNavigation) DeletionRangesForTextSelection_Direction_Destination_AllowsDecomposition(textSelection ITextSelection, direction TextSelectionNavigationDirection, destination TextSelectionNavigationDestination, allowsDecomposition bool) []TextRange {
	rv := objc.CallMethod[[]TextRange](t_, "deletionRangesForTextSelection:direction:destination:allowsDecomposition:", textSelection, direction, destination, allowsDecomposition)
	return rv
}

func (t_ TextSelectionNavigation) AllowsNonContiguousRanges() bool {
	rv := objc.CallMethod[bool](t_, "allowsNonContiguousRanges")
	return rv
}

func (t_ TextSelectionNavigation) SetAllowsNonContiguousRanges(value bool) {
	objc.CallMethod[objc.Void](t_, "setAllowsNonContiguousRanges:", value)
}

func (t_ TextSelectionNavigation) RotatesCoordinateSystemForLayoutOrientation() bool {
	rv := objc.CallMethod[bool](t_, "rotatesCoordinateSystemForLayoutOrientation")
	return rv
}

func (t_ TextSelectionNavigation) SetRotatesCoordinateSystemForLayoutOrientation(value bool) {
	objc.CallMethod[objc.Void](t_, "setRotatesCoordinateSystemForLayoutOrientation:", value)
}
