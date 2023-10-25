package typing

import (
	"github.com/hsiafan/cocoa/generate/internal/set"
)

var _ Type = (*ProtocolType)(nil)

// ProtocolType objective-c protocol type
type ProtocolType struct {
	Name    string  // the objc type name
	GName   string  // Go name, usually is objc type name without prefix 'NS'
	OnlyUse bool    // if this protocol is just for use, no need to implement it
	Module  *Module // object-c module name
}

func (p *ProtocolType) GoName(currentModule *Module, receiveFromObjc bool) string {
	if receiveFromObjc {

		if p.OnlyUse {
			return FullGoName(*p.Module, p.GName, *currentModule)
		}
	}
	return Object.GoName(currentModule, receiveFromObjc)
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
