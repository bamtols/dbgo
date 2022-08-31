package cliTmp

import "github.com/bamtols/dbgo/cli/cliTypes/cliFields"

type (
	ModelMigration struct {
		Job    cliFields.MigJob  `yaml:"Job"`
		Column cliFields.Column  `yaml:"Column"`
		GoType cliFields.GoType  `yaml:"GoType"`
		DbType *cliFields.DbType `yaml:"DbType"`
		Tags   cliFields.TagList `yaml:"Tags"`
	}

	ModelMigrationList []ModelMigration
)
