package cliTmp

import "github.com/bamtols/dbgo/cli/cliTypes/cliFields"

type (
	Database struct {
		Host     cliFields.EnvString   `yaml:"Host"`
		Port     cliFields.EnvString   `yaml:"Port"`
		Username cliFields.EnvString   `yaml:"Username"`
		Password cliFields.EnvString   `yaml:"Password"`
		Driver   cliFields.DbDriver    `yaml:"Driver"`
		Models   cliFields.ModelNmList `yaml:"Models"`
	}

	DatabaseMap map[DatabaseNm]Database
	DatabaseNm  string
)
