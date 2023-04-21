package typing

import (
	"github.com/hsiafan/cocoa/generate/internal/set"
)

var _ Type = (*ProtocolType)(nil)

// ProtocolType objective-c protocol type
type ProtocolType struct {
	Name   string  // the objc type name
	GName  string  // Go name, usually is objc type name without prefix 'NS'
	Module *Module // object-c module name
}

func (p *ProtocolType) GoName(currentModule *Module, receiveFromObjc bool) string {
	name := FullGoName(*p.Module, p.GName, *currentModule)
	if receiveFromObjc {
		if p.GName != "Object" {
			name += "Wrapper"
		}
	}
	return name
}

func (p *ProtocolType) ObjcName() string {
	return "id<" + p.Name + ">"
}

func (p *ProtocolType) GoImports() set.Set[string] {
	return set.New("github.com/hsiafan/cocoa/objc", "github.com/hsiafan/cocoa/"+p.Module.Package)
}

func (p *ProtocolType) DeclareModule() *Module {
	return p.Module
}
