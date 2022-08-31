package cliTmp

import (
	"github.com/bamtols/dbgo/cli/cliTypes/cliFields"
)

type (
	Model struct {
		Version    cliFields.YmlVersion  `yaml:"Version"`
		Columns    cliFields.ColumnMap   `yaml:"Columns"`
		Binds      cliFields.BindMap     `yaml:"Binds"`
		Indexes    cliFields.IndexNmMap  `yaml:"Indexes"`
		Uniques    cliFields.UniqueNmMap `yaml:"Uniques"`
		Migrations ModelMigrationList    `yaml:"Migrations"`
	}
)

func NewModelTemplate() *Model {
	return &Model{
		Version: cliFields.YmlVersionNow,
		Columns: cliFields.ColumnMap{
			"Id": {
				GoType: "uint",
				Tags: cliFields.TagList{
					cliFields.TagPrimaryKey,
				},
				Comment: cliFields.NewComment("등록번호"),
			},
			"UserId": {
				GoType:  "uint",
				Comment: cliFields.NewComment("사용자 등록번호"),
			},
			"Username": {
				GoType: "string",
				DbType: cliFields.NewDbType("varchar(100)"),
				Tags: cliFields.TagList{
					cliFields.TagUnique,
				},
			},
		},
		Binds: cliFields.BindMap{
			"UserSession": {
				Model:      "UserSession",
				BindType:   cliFields.BindTypeBelongsTo,
				ForeignKey: cliFields.NewColumnNm("UserId"),
				ReferKey:   cliFields.NewColumnNm("Id"),
			},
		},
		Indexes:    nil,
		Uniques:    nil,
		Migrations: nil,
	}
}
