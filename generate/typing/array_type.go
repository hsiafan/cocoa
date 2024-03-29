package typing

import (
	"github.com/hsiafan/cocoa/generate/internal/set"
)

var _ Type = (*ArrayType)(nil)

// ArrayType the element should be StringType or ClassType
type ArrayType struct {
	Type Type
}

func (a *ArrayType) GoImports() set.Set[string] {
	imports := set.New("github.com/hsiafan/cocoa/foundation")
	imports.AddSet(a.Type.GoImports())
	return imports
}

func (a *ArrayType) GoName(currentModule *Module, receiveFromObjc bool) string {
	return "[]" + a.Type.GoName(currentModule, receiveFromObjc)
}

func (a *ArrayType) ObjcName() string {
	return "NSArray*"
}

func (a *ArrayType) DeclareModule() *Module {
	return a.Type.DeclareModule()
}
