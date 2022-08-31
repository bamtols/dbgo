package cliFields

import (
	"github.com/bamtols/dbgo/func/errFn"
	"github.com/bamtols/dbgo/func/parserFn"
)

type (
	Column struct {
		GoType  GoType   `yaml:"GoType"`
		DbType  *DbType  `yaml:"DbType"`
		Tags    TagList  `yaml:"Tags"`
		Comment *Comment `yaml:"Comment"`
	}

	ColumnMap map[ColumnNm]Column
)

func (x *ColumnNm) String() string {
	return parserFn.GoPublicCamelCase(string(*x))
}

func (x *Column) Validate(goTypeMap GoTypeMap) (err error) {
	x.Tags = x.Tags.Validate()
	if !goTypeMap.HasType(x.GoType) {
		return errFn.BindF(ErrNotFound, "GoType : goType=%s", x.GoType)
	}
	return
}
