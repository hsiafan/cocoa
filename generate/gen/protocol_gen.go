package gen

import "C"
import (
	"strings"

	"github.com/hsiafan/cocoa/generate/internal/set"
	"github.com/hsiafan/cocoa/generate/typing"
)

var _ CodeGen = (*Protocol)(nil)

// Protocol is code generator for objc protocol.
type Protocol struct {
	Type       *typing.ProtocolType
	Parents    []*Protocol
	Methods    []*Method
	Properties []*Property

	init bool
}

// Copy implements CodeGen
func (p *Protocol) Copy() CodeGen {
	parents := make([]*Protocol, len(p.Parents))
	for i := 0; i < len(p.Parents); i++ {
		parents[i] = p.Parents[i].Copy().(*Protocol)
	}
	return &Protocol{
		Type:       p.Type,
		Parents:    p.Parents,
		Methods:    p.Methods,
		Properties: p.Properties,
	}
}

func (p *Protocol) String() string {
	return p.Type.Name
}

func (p *Protocol) Init() {
	if p.init {
		return
	}

	for _, m := range p.Methods {
		if _, ok := m.ReturnType.(*typing.InstanceType); ok {
			m.ReturnType = typing.Object
		}
	}

	for _, parent := range p.Parents {
		if parent != nil && parent.Type.Name != "NSObject" {
			parent.Init()
		}
	}

	var methodSet = set.New[string]()

	var methods []*Method
	for _, m := range p.Methods {
		if methodSet.Contains(m.Selector()) {
			continue
		}
		methodSet.Add(m.Selector())
		methods = append(methods, m)
	}
	p.Methods = methods

	var propertySet = set.New[string]()

	var properties []*Property
	for _, pp := range p.Properties {
		if propertySet.Contains(pp.Name) {
			continue
		}
		propertySet.Add(pp.Name)
		properties = append(properties, pp)
	}
	p.Properties = properties

	p.init = true
}

func (p *Protocol) GoImports() set.Set[string] {
	imports := set.New("github.com/hsiafan/cocoa/objc")
	for _, parent := range p.Parents {
		imports.Add("github.com/hsiafan/cocoa/" + parent.Type.Module.Package)
	}
	for _, m := range p.Methods {
		imports.AddSet(m.GoImports())
	}
	for _, p := range p.Properties {
		imports.AddSet(p.getter().GoImports())
	}
	return imports
}

func (p *Protocol) WriteGoCode(w *CodeWriter) {
	w.WriteLine("")

	if p.Type.OnlyUse {
		// Protocol Wrapper
		p.writeProtocolWrapperStruct(w)
	} else {
		// Protocol Interface
		p.writeProtocolInterface(w)

		// Prototol base impl struct
		w.WriteLine("")
		p.writeProtocolBaseStruct(w)

		if strings.Contains(p.Type.Name, "Delegate") {
			p.writeProtocolCreator(w)
		}
	}
}

func (p *Protocol) writeProtocolInterface(w *CodeWriter) {
	w.WriteLine("type " + p.Type.GName + " interface {")
	w.Indent()
	for _, pp := range p.Parents {
		if pp.Type.Name != "NSObject" {
			w.WriteLine(typing.FullGoName(*pp.Type.Module, pp.Type.GName, *p.Type.Module))
		}
	}
	for _, m := range p.allMethods() {
		if !m.Required {
			w.WriteLine("Implements" + m.ProtocolGoFuncName() + "() bool")
			w.WriteLine("// optional")
		} else {
			w.WriteLine("// required")
		}
		if m.Deprecated {
			w.WriteLine("// deprecated")
		}
		w.WriteLine(m.ProtocolGoFuncName() + m.ProtocolGoFuncFieldType(p.Type.Module))
	}

	w.UnIndent()
	w.WriteLine("}")

	// a convert method
	w.WriteLine("")
	w.WriteLineF("func Wrap%v(v %v) objc.Object {", p.Type.GName, p.Type.GName)
	w.Indent()
	w.WriteLineF("return objc.WrapAsProtocol(\"%v\", v)", p.Type.Name)
	w.UnIndent()
	w.WriteLine("}")
}

// Protocol base struct for all optional methods. Make implement protocol interface easier.
func (p *Protocol) writeProtocolBaseStruct(w *CodeWriter) {
	protocolBaseName := p.Type.GName + "Base"
	w.WriteLine("type " + protocolBaseName + " struct {")
	for _, pp := range p.Parents {
		if pp.Type.Name != "NSObject" {
			w.WriteLine(typing.FullGoName(*pp.Type.Module, pp.Type.GName, *p.Type.Module) + "Base")
		}
	}
	w.WriteLine("}")

	receiver := "p"
	for _, m := range p.allMethods() {
		if !m.Required {
			w.WriteLine("")
			w.WriteLineF("func (%s *%s) Implements%s() bool {", receiver, protocolBaseName, m.ProtocolGoFuncName())
			w.WriteLine("\t return false")
			w.WriteLine("}")
			w.WriteLine("")
			if m.Deprecated {
				w.WriteLine("// deprecated")
			}
			w.WriteLineF("func (%s *%s) %s%s {", receiver, protocolBaseName, m.ProtocolGoFuncName(), m.ProtocolGoFuncFieldType(p.Type.Module))
			w.WriteLine("\tpanic(\"unimpemented protocol method\")")
			w.WriteLine("}")
		}
	}
}

// A Helper for create protocol instance, no need to wrap a go interface
func (p *Protocol) writeProtocolCreator(w *CodeWriter) {
	creatorName := p.Type.GName + "Creator"
	w.WriteLine("type " + creatorName + " struct {")
	w.WriteLine("\tclassName string")
	w.WriteLine("\tclass objc.Class")
	w.WriteLine("}")

	w.WriteLineF("func New%s(name string) *%s {", creatorName, creatorName)
	w.Indent()
	w.WriteLineF("class := objc.AllocateClassPair(objc.GetClass(\"NSObject\"), name, 0)")
	w.WriteLine("objc.RegisterClassPair(class)")
	w.WriteLineF("return &%s{className: name, class: class}", creatorName)
	w.UnIndent()
	w.WriteLine("}")

	receiver := "c"
	for _, m := range p.allMethods() {
		if !m.Required {
			w.WriteLine("")
			if m.Deprecated {
				w.WriteLine("// deprecated")
			}
			funcSign := "(o objc.Object, " + m.ProtocolGoFuncFieldType(p.Type.Module)[1:]
			w.WriteLineF("func (%s *%s) Set%s(handle func %s) {", receiver, creatorName, m.ProtocolGoFuncName(), funcSign)
			w.WriteLineF("\tobjc.AddMethod(c.class, objc.SEL(\"%s\"), handle)", m.Selector())
			w.WriteLine("}")
		}
	}

	w.WriteLine("")
	w.WriteLineF("func (%s *%s) Create() objc.Object {", receiver, creatorName)
	w.WriteLineF("\treturn %s.class.CreateInstance(0)", receiver)
	w.WriteLine("}")
}

// generate protocol instance wrapper, for protocol which passed from objc to go code.
func (p *Protocol) writeProtocolWrapperStruct(w *CodeWriter) {
	typeName := p.Type.GName
	w.WriteLine("type " + typeName + " struct {")
	w.Indent()
	w.WriteLine("objc.Object")
	w.UnIndent()
	w.WriteLine("}")

	for _, m := range p.allMethodsRecursively() {
		w.WriteLine("")
		m.WriteGoCallCode(p.Type.Module, typeName, w)
	}

}

// include property getter/setter
func (p *Protocol) allMethods() []*Method {
	var allMethods = make([]*Method, len(p.Methods))
	copy(allMethods, p.Methods)
	for _, pp := range p.Properties {
		if !pp.ReadOnly {
			allMethods = append(allMethods, (*Method)(pp.setter()))
		}

		allMethods = append(allMethods, (*Method)(pp.getter()))
	}
	return allMethods
}

// method include parents..
func (p *Protocol) allMethodsRecursively() []*Method {
	var allMethods = p.allMethods()
	for _, pp := range p.Parents {
		if pp.Type.Name != "NSObject" {
			allMethods = append(allMethods, pp.allMethodsRecursively()...)
		}
	}
	return allMethods
}
