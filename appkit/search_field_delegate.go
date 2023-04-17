// auto-generated code, do not modify
package appkit

import (
	"github.com/hsiafan/cocoa/ffi"
	"github.com/hsiafan/cocoa/objc"
)

type SearchFieldDelegate interface {
	TextFieldDelegate
	ImplementsSearchFieldDidStartSearching() bool
	// optional
	SearchFieldDidStartSearching(sender SearchField)
	ImplementsSearchFieldDidEndSearching() bool
	// optional
	SearchFieldDidEndSearching(sender SearchField)
}

type SearchFieldDelegateImpl struct {
	TextFieldDelegateImpl
	_SearchFieldDidStartSearching func(sender SearchField)
	_SearchFieldDidEndSearching   func(sender SearchField)
}

func (di *SearchFieldDelegateImpl) ImplementsSearchFieldDidStartSearching() bool {
	return di._SearchFieldDidStartSearching != nil
}

func (di *SearchFieldDelegateImpl) SetSearchFieldDidStartSearching(f func(sender SearchField)) {
	di._SearchFieldDidStartSearching = f
}

func (di *SearchFieldDelegateImpl) SearchFieldDidStartSearching(sender SearchField) {
	di._SearchFieldDidStartSearching(sender)
}
func (di *SearchFieldDelegateImpl) ImplementsSearchFieldDidEndSearching() bool {
	return di._SearchFieldDidEndSearching != nil
}

func (di *SearchFieldDelegateImpl) SetSearchFieldDidEndSearching(f func(sender SearchField)) {
	di._SearchFieldDidEndSearching = f
}

func (di *SearchFieldDelegateImpl) SearchFieldDidEndSearching(sender SearchField) {
	di._SearchFieldDidEndSearching(sender)
}

type SearchFieldDelegateWrapper struct {
	TextFieldDelegateWrapper
}

func (s_ *SearchFieldDelegateWrapper) ImplementsSearchFieldDidStartSearching() bool {
	return s_.RespondsToSelector(objc.GetSelector("searchFieldDidStartSearching:"))
}

func (s_ SearchFieldDelegateWrapper) SearchFieldDidStartSearching(sender ISearchField) {
	ffi.CallMethod[ffi.Void](s_, "searchFieldDidStartSearching:", sender)
}

func (s_ *SearchFieldDelegateWrapper) ImplementsSearchFieldDidEndSearching() bool {
	return s_.RespondsToSelector(objc.GetSelector("searchFieldDidEndSearching:"))
}

func (s_ SearchFieldDelegateWrapper) SearchFieldDidEndSearching(sender ISearchField) {
	ffi.CallMethod[ffi.Void](s_, "searchFieldDidEndSearching:", sender)
}
