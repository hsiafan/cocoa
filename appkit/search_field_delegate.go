// auto-generated code, do not modify
package appkit

import (
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

func WrapSearchFieldDelegate(v SearchFieldDelegate) objc.Object {
	return objc.WrapAsProtocol("NSSearchFieldDelegate", v)
}

type SearchFieldDelegateBase struct {
	TextFieldDelegateBase
}

func (p *SearchFieldDelegateBase) ImplementsSearchFieldDidStartSearching() bool {
	return false
}

func (p *SearchFieldDelegateBase) SearchFieldDidStartSearching(sender SearchField) {
	panic("unimpemented protocol method")
}

func (p *SearchFieldDelegateBase) ImplementsSearchFieldDidEndSearching() bool {
	return false
}

func (p *SearchFieldDelegateBase) SearchFieldDidEndSearching(sender SearchField) {
	panic("unimpemented protocol method")
}
