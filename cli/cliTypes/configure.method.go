package cliTypes

import (
	"github.com/bamtols/dbgo/cli/cliFn"
	"github.com/bamtols/dbgo/func/errFn"
	"strings"
)

type (
	FilePath struct {
		GoModelPath    string
		DefinitionPath string
		MigrationPath  string
	}
)

func (x *Configure) SetDefault() {
	x.GoModelPath = cliFn.NilDefault(x.GoModelPath, defGoModelPath)
	x.DefinitionPath = cliFn.NilDefault(x.DefinitionPath, defDefinitionPath)
	x.MigrationPath = cliFn.NilDefault(x.MigrationPath, defMigrationPath)
	x.NameStrategy = cliFn.NilDefault(x.NameStrategy, defNameStrategy)
	x.GoType = make(GoTypePackageMap)
	x.Database = make(DatabaseMap)
	x.modelConfigure = &ModelConfigure{
		GoTypeMap: make(GoTypeMap),
		ModelMap:  make(ModelMap),
	}
}

func (x *Configure) GetDatabase(nm DatabaseNm) (*Database, error) {
	res, isOk := x.Database[nm]
	if !isOk {
		return nil, errFn.BindF(ErrNotFound, "Database : key=%s", nm)
	}
	return &res, nil
}

func (x *Configure) DefaultDB() (*Database, error) {
	key, err := x.Database.DefaultDBKey()
	if err != nil {
		return nil, err
	}

	db := x.Database[key]
	return &db, nil
}

func (x *Configure) AbsoluteFilePath() *FilePath {
	return &FilePath{
		GoModelPath:    cliFn.GetAbsolutePathWithDefault(x.GoModelPath, defGoModelPath),
		DefinitionPath: cliFn.GetAbsolutePathWithDefault(x.DefinitionPath, defDefinitionPath),
		MigrationPath:  cliFn.GetAbsolutePathWithDefault(x.MigrationPath, defMigrationPath),
	}
}

func (x *Configure) ModelConfigure() *ModelConfigure {
	if x.modelConfigure == nil {
		panic(errFn.BindF(ErrNotFound, "modelConfigure is nil"))
	}

	if len(x.modelConfigure.GoTypeMap) == 0 {
		x.modelConfigure.GoTypeMap = NewGoTypeCollection(x)
	}

	return x.modelConfigure
}

func (x *Configure) GoPackageNm() string {
	path := x.AbsoluteFilePath()
	split := strings.Split(path.GoModelPath, "/")
	if len(split) < 1 {
		panic(errFn.BindF(ErrNotFound, "go package name : %s", path.GoModelPath))
	}
	return split[len(split)-1]
}
