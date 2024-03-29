package gen

import (
	"fmt"
	"strings"

	"github.com/hsiafan/cocoa/generate/internal/set"
	"github.com/hsiafan/cocoa/generate/internal/stringx"

	"github.com/hsiafan/cocoa/generate/typing"
)

// Method is code generator for objective-c method
type Method struct {
	Name         string // the first part of objc method name
	GoName       string
	Params       []*Param
	ReturnType   typing.Type
	ClassMethod  bool // true if is class method
	WeakProperty bool // if is a weak property setter
	Deprecated   bool // if has been deprecated
	Required     bool // If this method is required. only for protocol method.
	InitMethod   bool // method that return instancetype

	goFuncName string
	identifier string
}

func (m *Method) needRelease() bool {
	return strings.HasPrefix(m.Name, "new") || !strings.HasPrefix(m.Name, "init") && strings.HasPrefix(m.Name, "Initial") ||
		strings.HasPrefix(m.Name, "copy") || strings.HasPrefix(m.Name, "mutableCopy")
}

// Selector return full Objc function name
func (m *Method) Selector() string {
	if m.identifier == "" {
		var sb strings.Builder
		sb.WriteString(m.Name)
		for idx, p := range m.Params {
			if idx > 0 {
				sb.WriteString(p.FieldName)
			}
			sb.WriteString(":")
		}
		m.identifier = sb.String()
	}
	return m.identifier
}

func (m *Method) String() string {
	return m.Selector()
}

// NormalizeInstanceTypeMethod return new init method.
func (m *Method) NormalizeInstanceTypeMethod(returnType *typing.ClassType) *Method {
	nm := &Method{
		Name:        m.Name,
		GoName:      m.GoName,
		Params:      m.Params,
		ReturnType:  returnType,
		ClassMethod: m.ClassMethod,
		InitMethod:  m.InitMethod,
		goFuncName:  m.goFuncName,
	}
	return nm
}

// WriteGoCallCode generate go method code to call c wrapper code
func (m *Method) WriteGoCallCode(currentModule *typing.Module, typeName string, cw *CodeWriter) {
	funcDeclare := m.GoFuncDeclare(currentModule, typeName)

	if m.WeakProperty {
		cw.WriteLine("// weak property")
	}
	if m.Deprecated {
		cw.WriteLine("// deprecated")
	}

	var receiver string
	if m.ClassMethod {
		receiver = strings.ToLower(typeName[0:1]) + "c"
		cw.WriteLine("func (" + receiver + " _" + typeName + "Class) " + funcDeclare + " {")
	} else {
		receiver = strings.ToLower(typeName[0:1] + "_")
		cw.WriteLine("func (" + receiver + " " + typeName + ") " + funcDeclare + " {")
	}

	cw.Indent()

	var returnTypeStr string
	rt := typing.UnwrapAlias(m.ReturnType)
	switch rt.(type) {
	case *typing.VoidType:
		returnTypeStr = "objc.Void"
	default:
		returnTypeStr = m.ReturnType.GoName(currentModule, true)
	}
	callCode := fmt.Sprintf("objc.CallMethod[%s](%s, objc.SEL(\"%s\")", returnTypeStr, receiver, m.Selector())
	var sb strings.Builder
	for _, p := range m.Params {
		sb.WriteString(", ")
		switch tt := p.Type.(type) {
		case *typing.ClassType:
			sb.WriteString("objc.ExtractPtr(" + p.GoName() + ")")
		case *typing.ProtocolType:
			sb.WriteString("objc.ExtractPtr(" + p.GoName() + ")")
		case *typing.PointerType:
			switch tt.Type.(type) {
			case *typing.ClassType: //object pointer convert to unsafe.Pointer, avoiding ffi treat it as PointerHolder
				sb.WriteString(fmt.Sprintf("unsafe.Pointer(%s)", p.GoName()))
			default:
				sb.WriteString(p.GoName())
			}
		default:
			sb.WriteString(p.GoName())
		}
	}
	callCode += sb.String() + ")"

	switch rt.(type) {
	case *typing.VoidType:
		cw.WriteLine(callCode)
	default:
		var resultName = "rv"
		cw.WriteLine(resultName + " := " + callCode)
		if m.needRelease() {
			cw.WriteLine(resultName + ".Autorelease()")
		}
		cw.WriteLine("return " + resultName)
	}
	cw.UnIndent()
	cw.WriteLine("}")
}

// WriteGoInterfaceCode generate go interface method signature code
func (m *Method) WriteGoInterfaceCode(currentModule *typing.Module, classType *typing.ClassType, w *CodeWriter) {
	if m.ClassMethod {
		return
	}
	if m.Deprecated {
		w.WriteLine("// deprecated")
	}
	funcDeclare := m.GoFuncDeclare(currentModule, classType.GName)
	w.WriteLine(funcDeclare)
}

// GoFuncDeclare generate go function declaration
func (m *Method) GoFuncDeclare(currentModule *typing.Module, goTypeName string) string {
	var paramStrs []string
	for _, p := range m.Params {
		paramStrs = append(paramStrs, p.GoDeclare(currentModule, false))
	}

	var returnType = m.ReturnType.GoName(currentModule, true)
	return m.GoFuncName(goTypeName) + "(" + strings.Join(paramStrs, ", ") + ")" + " " + returnType
}

// GoFuncName return go func name
func (m *Method) GoFuncName(goTypeName string) string {
	if m.goFuncName == "" {
		var sb strings.Builder
		name := m.GoName
		sb.WriteString(stringx.Capitalize(name))

		for idx, p := range m.Params {
			if idx > 0 {
				sb.WriteString("_")
			}
			sb.WriteString(stringx.Capitalize(p.FieldName))
		}
		m.goFuncName = sb.String()
	}

	return m.goFuncName
}

// ProtocolGoFuncFieldType generate go function declaration for protocol struct impl field
func (m *Method) ProtocolGoFuncFieldType(currentModule *typing.Module) string {
	var paramStrs []string
	for _, p := range m.Params {
		paramStrs = append(paramStrs, p.GoDeclare(currentModule, true))
	}

	return "(" + strings.Join(paramStrs, ", ") + ")" + " " + m.ReturnType.GoName(currentModule, false)
}

// ProtocolGoFuncName return go protocol func name
func (m *Method) ProtocolGoFuncName() string {
	if m.goFuncName == "" {
		var sb strings.Builder
		sb.WriteString(stringx.Capitalize(m.Name))
		for idx, p := range m.Params {
			if idx == 0 {
				continue
			}
			sb.WriteByte('_')
			sb.WriteString(stringx.Capitalize(p.FieldName))
		}
		m.goFuncName = sb.String()
	}

	return m.goFuncName
}

// GoImports return all imports for go file
func (m *Method) GoImports() set.Set[string] {
	var imports = set.New("github.com/hsiafan/cocoa/objc")
	for _, param := range m.Params {
		imports.AddSet(param.Type.GoImports())
		if _, ok := param.Type.(*typing.ClassType); ok {
			imports.Add("github.com/hsiafan/cocoa/" + typing.ObjCRuntime.Package)
		}
	}
	if m.WeakProperty {
		imports.Add("github.com/hsiafan/cocoa/internal")
	}
	imports.AddSet(m.ReturnType.GoImports())
	return imports
}
