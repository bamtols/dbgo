package cliFields

import (
	"fmt"
	"github.com/bamtols/dbgo/func/errFn"
	"strings"
)

type (
	GoType           string
	GoTypeList       []GoType
	GoTypeMap        map[GoType]GoTypeDefinition
	GoTypeDefinition struct {
		Import *string
		Type   string
	}
)

const (
	GoTypeString  GoType = "string"
	GoTypeBool    GoType = "bool"
	GoTypeInt     GoType = "int"
	GoTypeInt8    GoType = "int8"
	GoTypeInt16   GoType = "int16"
	GoTypeInt32   GoType = "int32"
	GoTypeInt64   GoType = "int64"
	GoTypeUInt    GoType = "uint"
	GoTypeUInt8   GoType = "uint8"
	GoTypeUInt16  GoType = "uint16"
	GoTypeUInt32  GoType = "uint32"
	GoTypeUInt64  GoType = "uint64"
	GoTypeFloat32 GoType = "float32"
	GoTypeFloat64 GoType = "float64"
)

var (
	GoTypeDefaultAll = GoTypeList{
		GoTypeString,
		GoTypeBool,
		GoTypeInt,
		GoTypeInt8,
		GoTypeInt16,
		GoTypeInt32,
		GoTypeInt64,
		GoTypeUInt,
		GoTypeUInt8,
		GoTypeUInt16,
		GoTypeUInt32,
		GoTypeUInt64,
		GoTypeFloat32,
		GoTypeFloat64,
	}
)

func NewGoTypeMap() GoTypeMap {
	res := make(GoTypeMap)
	res.addDefaultType()
	return res
}

func (x *GoTypeMap) addDefaultType() {
	for _, goType := range GoTypeDefaultAll {
		(*x)[goType] = GoTypeDefinition{
			Type: goType.String(),
		}
	}
}

func (x *GoTypeMap) HasType(goType GoType) bool {
	_, isOk := (*x)[goType]
	if !isOk {
		return false
	}
	return true
}

func (x *GoTypeMap) AddType(goType GoType, goPackage string) {
	if goPackage == "" {
		panic(errFn.BindF(ErrNotFound, "GoTypePackage : goType=%s", goType))
	}

	s := strings.Split(goPackage, "/")
	if len(s) == 0 {
		panic(errFn.BindF(ErrInvalid, "GoTypePackage : package=%s", goPackage))
	}

	p := strings.Split(goPackage, ".")
	if len(p) == 0 {
		panic(errFn.BindF(ErrInvalid, "GoTypePackage : package=%s", goPackage))
	}

	importGoCode := strings.TrimSuffix(goPackage, fmt.Sprintf(".%s", p[len(p)-1]))

	(*x)[goType] = GoTypeDefinition{
		Import: &importGoCode,
		Type:   s[len(s)-1],
	}
}

func (x *GoTypeMap) Definition(t GoType) GoTypeDefinition {
	def, isOk := (*x)[t]
	if !isOk {
		panic(errFn.BindF(ErrNotFound, "goType : goTypeKey=%s", t))
	}
	return def
}

func (x *GoTypeMap) PrintGoType(t GoType) string {
	def := x.Definition(t)
	return def.Type
}

func (x *GoTypeMap) PrintPtrGoType(t GoType) string {
	def := x.Definition(t)
	return fmt.Sprintf("*%s", def.Type)
}

func (x *GoTypeMap) PrintSliceGoType(t GoType) string {
	def := x.Definition(t)
	return fmt.Sprintf("[]%s", def.Type)
}

func (x *GoType) String() string {
	return string(*x)
}
