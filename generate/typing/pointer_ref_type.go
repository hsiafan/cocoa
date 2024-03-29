package typing

import (
	"github.com/hsiafan/cocoa/generate/internal/set"
)

var _ Type = (*PointerRefType)(nil)

// PointerRefType c pointer type def
type PointerRefType struct {
	Name   string  // objc type name
	GName  string  // Go name, usually is objc type name without prefix 'NS'
	Module *Module // object-c module
}

func (c *PointerRefType) GoImports() set.Set[string] {
	return set.New("github.com/hsiafan/cocoa/" + c.Module.Package)
}

func (c *PointerRefType) GoName(currentModule *Module, receiveFromObjc bool) string {
	return FullGoName(*c.Module, c.GName, *currentModule)
}

func (c *PointerRefType) ObjcName() string {
	return c.Name
}

func (c *PointerRefType) DeclareModule() *Module {
	return c.Module
}
