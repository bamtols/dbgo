package cliTmp

import (
	"github.com/bamtols/dbgo/cli/cliTypes/cliFields"
	"github.com/bamtols/dbgo/func/errFn"
	"github.com/bamtols/dbgo/func/fsFn"
	"github.com/bamtols/dbgo/func/isFn"
	"github.com/bamtols/dbgo/func/serializeFn"
	"os"
)

type (
	Configure struct {
		Version      cliFields.YmlVersion       `yaml:"Version"`
		Paths        *ConfigurePath             `yaml:"Paths,omitempty"`
		NameStrategy *cliFields.NameStrategy    `yaml:"NameStrategy,omitempty"`
		GoTypes      cliFields.GoTypePackageMap `yaml:"GoTypes,omitempty"`
		Databases    DatabaseMap                `yaml:"Databases,omitempty"`
	}
)

func NewConfigureTemplate() *Configure {
	return &Configure{
		Version: cliFields.YmlVersionNow,
		Paths: &ConfigurePath{
			Definition: cliFields.NewFilePath("dbgos/definitions"),
			Migration:  cliFields.NewFilePath("dbgos/migrations"),
			Generated:  cliFields.NewFilePath("dbgos/dbgos"),
		},
		NameStrategy: cliFields.NameStrategySnakeCase.Ptr(),
		GoTypes: cliFields.GoTypePackageMap{
			"Time": "time.Time",
		},
		Databases: DatabaseMap{
			"Service": {
				Connection: DatabaseConnection{
					Host:     cliFields.NewEnvString("!Env DB_HOST"),
					Port:     cliFields.NewEnvString("!Env DB_PORT"),
					Username: cliFields.NewEnvString("!Env DB_USERNAME"),
					Password: cliFields.NewEnvString("!Env DB_PASSWORD"),
					Driver:   cliFields.DbDriverMySQL,
				},
				Models: cliFields.ModelNmList{
					"User",
				},
			},
		},
	}
}

func (x *Configure) Save(fp string, force ...bool) error {
	if fsFn.IsExistFile(fp) {
		if !isFn.IsTrue(force) {
			return errFn.BindF(ErrAlreadyExistFile, "filePath=%s", fp)
		}

		if err := fsFn.RemoveFile(fp); err != nil {
			return err
		}
	}

	file, err := os.Create(fp)
	if err != nil {
		return err
	}

	defer file.Close()

	return serializeFn.EncodeYaml(file, x)
}
