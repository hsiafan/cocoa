package main

import (
	"log"
	"strings"

	"github.com/hsiafan/cocoa/generate/data"
	"github.com/hsiafan/cocoa/generate/gen"
	"github.com/hsiafan/cocoa/generate/typing"
)

var genCache = map[string]gen.CodeGen{}

func getClassGen(ci *data.Class) *gen.Class {
	key := ci.Module + "." + ci.Name
	cg, ok := genCache[key]
	if ok {
		return cg.Copy().(*gen.Class)
	}

	type_ := &typing.ClassType{
		Name:   ci.Name,
		GName:  toGoSymbolName(ci.Name),
		Module: typing.FindModule(ci.Module),
	}
	classGen := &gen.Class{
		Type: type_,
	}
	if ci.Parent != "" {
		parentStr := strings.TrimRightFunc(ci.Parent, func(r rune) bool {
			return r == '*' || r == ' '
		})
		module, name, found := strings.Cut(parentStr, ".")
		if !found {
			panic("invalid class name:" + parentStr)
		}
		parentInfo := loadOneTypeInfo(module, name).(*data.Class)
		classGen.Parent = getClassGen(parentInfo)
	}

	classGen.Properties = convertToGenProperties(ci.Properties)
	classGen.Methods = convertToGenMethods(ci.Methods)

	genCache[key] = classGen
	return classGen

}

func getProtocolGen(pi *data.Protocol) *gen.Protocol {
	key := pi.Module + "." + pi.Name
	cg, ok := genCache[key]
	if ok {
		return cg.Copy().(*gen.Protocol)
	}

	type_ := &typing.ProtocolType{
		Name:    pi.Name,
		OnlyUse: pi.OnlyUse,
		GName:   toGoSymbolName(pi.Name),
		Module:  typing.FindModule(pi.Module),
	}
	protocolGen := &gen.Protocol{
		Type: type_,
	}
	for _, parent := range pi.Parents {
		module, name, found := strings.Cut(parent, ".")
		if !found {
			panic("invalid protocol name:" + parent)
		}
		if name == "NSObject" {
			continue
		}
		parentInfo := loadOneTypeInfo(module, name).(*data.Protocol)
		protocolGen.Parents = append(protocolGen.Parents, getProtocolGen(parentInfo))
	}

	protocolGen.Properties = convertToGenProperties(pi.Properties)
	protocolGen.Methods = convertToGenMethods(pi.Methods)

	genCache[key] = protocolGen
	return protocolGen

}

func convertToGenProperties(ps []*data.Property) []*gen.Property {
	var gps []*gen.Property
	for _, p := range ps {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println("convert property:", p.Name, "error:", err)
				}
			}()
			goName := p.GoName
			if goName == "" {
				goName = p.Name
			}
			gps = append(gps, &gen.Property{
				Name:          p.Name,
				GoName:        goName,
				ReadOnly:      p.ReadOnly,
				GetterName:    p.GetterName,
				Type:          ParseType(p.Type),
				ClassProperty: p.ClassProperty,
				Weak:          p.Weak,
				Deprecated:    p.Deprecated,
				Required:      p.Required,
			})
		}()
	}
	return gps
}

func convertToGenMethods(ms []*data.Method) []*gen.Method {
	var gms []*gen.Method
	for _, m := range ms {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println("convert method:", m.Name, "error:", err)
				}
			}()
			var params []*gen.Param
			for _, p := range m.Params {
				gp := &gen.Param{
					Name:      p.VariableName,
					Type:      ParseType(p.Type),
					FieldName: p.Name,
				}
				params = append(params, gp)
			}
			goName := m.GoName
			if goName == "" {
				goName = m.Name
			}
			gms = append(gms, &gen.Method{
				Name:        m.Name,
				GoName:      goName,
				Params:      params,
				ReturnType:  ParseType(m.ReturnType),
				ClassMethod: m.ClassMethod,
				Deprecated:  m.Deprecated,
				Required:    m.Required,
			})
		}()

	}
	return gms
}

func getEnumInfo(ai *data.Alias) *gen.EnumInfo {
	at := loadTypeCached(ai.Module, ai.Name).(*typing.AliasType)
	var enumValues []*gen.EnumValue
	for _, ev := range ai.Values {
		enumValues = append(enumValues, &gen.EnumValue{
			Name:       ev.Name,
			GoName:     toGoSymbolName(ev.Name),
			Value:      ev.Value,
			Arm64Value: ev.Arm64Value,
			Module:     typing.FindModule(ev.Module),
		})
	}
	ei := &gen.EnumInfo{
		AliasType: *at,
		Values:    enumValues,
	}
	ei.GName = toGoSymbolName(ai.Name)
	return ei
}
