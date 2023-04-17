package typing

import (
	"github.com/hsiafan/cocoa/generate/internal/set"
)

var _ Type = (*StringType)(nil)

// StringType string
type StringType struct {
}

func (s *StringType) GoImports() set.Set[string] {
	return set.New("github.com/hsiafan/cocoa/foundation")
}

func (s *StringType) GoName(currentModule *Module, receiveFromObjc bool) string {
	return "string"
}

func (s *StringType) ObjcName() string {
	return "NSString*"
}

func (s *StringType) DeclareModule() *Module {
	return Foundation
}
