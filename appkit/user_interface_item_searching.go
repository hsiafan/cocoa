// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/objc"
)

type UserInterfaceItemSearching interface {
	// required
	LocalizedTitlesForItem(item objc.Object) []string
	ImplementsShowAllHelpTopicsForSearchString() bool
	// optional
	ShowAllHelpTopicsForSearchString(searchString string)
	// required
	SearchForItemsWithSearchString_ResultLimit_MatchedItemHandler(searchString string, resultLimit int, handleMatchedItems func(items []objc.IObject))
	ImplementsPerformActionForItem() bool
	// optional
	PerformActionForItem(item objc.Object)
}

func WrapUserInterfaceItemSearching(v UserInterfaceItemSearching) objc.Object {
	return objc.WrapAsProtocol("NSUserInterfaceItemSearching", v)
}

type UserInterfaceItemSearchingBase struct {
}

func (p *UserInterfaceItemSearchingBase) ImplementsShowAllHelpTopicsForSearchString() bool {
	return false
}

func (p *UserInterfaceItemSearchingBase) ShowAllHelpTopicsForSearchString(searchString string) {
	panic("unimpemented protocol method")
}

func (p *UserInterfaceItemSearchingBase) ImplementsPerformActionForItem() bool {
	return false
}

func (p *UserInterfaceItemSearchingBase) PerformActionForItem(item objc.Object) {
	panic("unimpemented protocol method")
}

type UserInterfaceItemSearchingWrapper struct {
	objc.Object
}

func (u_ UserInterfaceItemSearchingWrapper) LocalizedTitlesForItem(item objc.IObject) []string {
	rv := objc.CallMethod[[]string](u_, objc.GetSelector("localizedTitlesForItem:"), objc.ExtractPtr(item))
	return rv
}

func (u_ UserInterfaceItemSearchingWrapper) ShowAllHelpTopicsForSearchString(searchString string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("showAllHelpTopicsForSearchString:"), searchString)
}

func (u_ UserInterfaceItemSearchingWrapper) SearchForItemsWithSearchString_ResultLimit_MatchedItemHandler(searchString string, resultLimit int, handleMatchedItems func(items []objc.Object)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("searchForItemsWithSearchString:resultLimit:matchedItemHandler:"), searchString, resultLimit, handleMatchedItems)
}

func (u_ UserInterfaceItemSearchingWrapper) PerformActionForItem(item objc.IObject) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("performActionForItem:"), objc.ExtractPtr(item))
}
