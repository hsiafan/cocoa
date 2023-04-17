package gen

import (
	"github.com/hsiafan/cocoa/generate/typing"
)

// EnumInfo enum type and values
type EnumInfo struct {
	typing.AliasType
	Values []*EnumValue
}

// IsString return if is a string type Enum
func (e *EnumInfo) IsString() bool {
	_, ok := e.Type.(*typing.StringType)
	return ok
}

// EnumValue the enum name and value
type EnumValue struct {
	Name   string         // the objc enum name
	GoName string         // the go name of enum
	Value  string         // the value
	Module *typing.Module // the module enum value defined in
}
