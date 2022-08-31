package cliTypes

import (
	"fmt"
	"github.com/bamtols/dbgo/func/parserFn"
)

type (
	Column struct {
		GoType  GoType        `yaml:"GoType" json:"goType"`
		DbType  *string       `yaml:"DbType,omitempty" json:"dbType,omitempty"`
		Tags    ColumnTagList `yaml:"Tags,omitempty" json:"tags,omitempty"`
		Comment *string       `yaml:"Comment,omitempty" json:"comment,omitempty"`
	}

	Hooks struct {
		TimeAt TimeAtList `yaml:"TimeAt,omitempty" json:"timeAt,omitempty"`
	}

	Bind struct {
		Model      ModelNm   `yaml:"Model" json:"model"`
		BindType   BindType  `yaml:"BindType" json:"bindType"`
		ForeignKey *ColumnNm `yaml:"ForeignKey" json:"foreignKey"`
		ReferKey   *ColumnNm `yaml:"ReferKey" json:"referKey"`
	}

	Migration struct {
		Job     MigrationJob  `yaml:"Job" json:"job"`
		GoType  GoType        `yaml:"GoType" json:"goType"`
		DbType  *string       `yaml:"DbType,omitempty" json:"dbType,omitempty"`
		Tags    ColumnTagList `yaml:"Tags,omitempty" json:"tags,omitempty"`
		Comment *string       `yaml:"Comment,omitempty" json:"comment,omitempty"`
	}

	DbColumnType  string
	ColumnTag     string
	ColumnMap     map[ColumnNm]Column
	ColumnTagList []ColumnTag
	ColumnNmList  []ColumnNm
	ColumnNmMap   map[string][]ColumnNm
	ColumnNm      string
	ModelNm       string
	ModelNmList   []ModelNm
	ModelMap      map[ModelNm]TmpModel
	TimeAt        string
	TimeAtList    []TimeAt
	BindType      string
	BindTypeList  []BindType
	BindMap       map[ColumnNm]Bind
	PackageNm     string
	MigrationList []Migration

	ICreateTable struct {
		Version string
		ModelNm ModelNm
	}
)

const (
	ColumnTagPrimaryKey ColumnTag = "primaryKey"
	ColumnTagOmit       ColumnTag = "omit"
	ColumnTagUnique     ColumnTag = "unique"
	ColumnTagNullable   ColumnTag = "nullable"
)

var (
	ColumnTagAll = ColumnTagList{
		ColumnTagPrimaryKey,
		ColumnTagOmit,
		ColumnTagUnique,
		ColumnTagNullable,
	}
)

const (
	TimeAtCreatedAt TimeAt = "createdAt"
	TimeAtUpdatedAt TimeAt = "updatedAt"
	TimeAtDeletedAt TimeAt = "deletedAt"
)

var (
	TimeAtAll = TimeAtList{
		TimeAtCreatedAt,
		TimeAtUpdatedAt,
		TimeAtDeletedAt,
	}
)

const (
	BindTypeHasOne    BindType = "hasOne"
	BindTypeHasMany   BindType = "hasMany"
	BindTypeBelongsTo BindType = "belongsTo"
)

var (
	BindTypeAll = BindTypeList{
		BindTypeHasOne,
		BindTypeHasMany,
		BindTypeBelongsTo,
	}
)

var (
	defTableNm = ModelNm("sample")
)

func NewTableTemplate(i *ICreateTable) *TmpModel {
	return &TmpModel{
		Version: Version(i.Version),
		Indexes: ColumnNmMap{
			"CreatedAt": []ColumnNm{"createdAt"},
			"SignIn":    []ColumnNm{"username", "password"},
		},
		Hooks: &Hooks{
			TimeAt: TimeAtList{
				TimeAtCreatedAt,
				TimeAtUpdatedAt,
				TimeAtDeletedAt,
			},
		},
		Columns: map[ColumnNm]Column{
			"Id": {
				GoType: GoTypeUInt,
				Tags: ColumnTagList{
					ColumnTagPrimaryKey,
				},
				Comment: parserFn.PtrString("등록번호"),
			},
			"Username": {
				GoType: "string",
				DbType: parserFn.PtrString("varchar(100)"),
				Tags: ColumnTagList{
					ColumnTagUnique,
				},
				Comment: parserFn.PtrString("아이디"),
			},
			"Password": {
				GoType:  "string",
				DbType:  parserFn.PtrString("varchar(500)"),
				Comment: parserFn.PtrString("비밀번호"),
			},
			"Point": {
				GoType:  "Decimal",
				DbType:  parserFn.PtrString("decimal(30,8)"),
				Comment: parserFn.PtrString("포인트"),
			},
			"JoinAt": {
				GoType: "Time",
			},
		},
	}
}

func NewModel(version string, modelNm string) *TmpModel {
	return &TmpModel{
		Version: Version(version),
		Columns: make(ColumnMap),
		Indexes: make(ColumnNmMap),
		Uniques: make(ColumnNmMap),
		Binds:   make(BindMap),
		Hooks:   new(Hooks),
		ModelNm: ModelNm(modelNm),
	}
}

func (x *ModelNm) String() string {
	return string(*x)
}

func (x *ColumnNm) String() string {
	return string(*x)
}

func (x *ColumnNm) GoPublicModelNm() string {
	return parserFn.GoPublicCamelCase(string(*x))
}

func (x ColumnNm) PrintPtrGoType() string {
	return fmt.Sprintf("*%s", x.PrintGoType())
}

func (x *ColumnNm) PrintSliceGoType() string {
	return fmt.Sprintf("[]%s", x.PrintGoType())
}

func (x *ColumnNm) PrintGoType() string {
	return fmt.Sprintf("%s", x.GoPublicModelNm())
}

func (x *TimeAt) String() string {
	return string(*x)
}

func (x *ModelNm) GoPublicModelNm() string {
	return parserFn.GoPublicCamelCase(string(*x))
}
