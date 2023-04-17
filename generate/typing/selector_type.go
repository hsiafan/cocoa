package typing

import (
	"github.com/hsiafan/cocoa/generate/internal/set"
)

var _ Type = (*SelectorType)(nil)

// SelectorType objc selector type
type SelectorType struct {
}

func (s *SelectorType) GoImports() set.Set[string] {
	return set.New("github.com/hsiafan/cocoa/" + ObjCRuntime.Package)
}

func (s *SelectorType) GoName(currentModule *Module, receiveFromObjc bool) string {
	if *currentModule == *ObjCRuntime {
		return "Selector"
	} else {
		return "objc.Selector"
	}
}

func (s *SelectorType) ObjcName() string {
	return "SEL"
}

func (s *SelectorType) DeclareModule() *Module {
	return ObjCRuntime
}
