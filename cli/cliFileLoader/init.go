package cliFileLoader

import (
	"github.com/bamtols/dbgo/cli/cliFn"
	"github.com/bamtols/dbgo/cli/cliTypes"
	"path/filepath"
)

type (
	InitConfigurationOptions struct {
		ConfigPath string
		Version    string
	}
)

func NewInitConfiguration(i *InitConfigurationOptions) error {
	temp := cliTypes.NewConfigureTemplate(i.Version)

	if err := cliFn.SerializeByYaml(i.ConfigPath, temp); err != nil {
		return err
	}

	path := temp.AbsoluteFilePath()
	if err := cliFn.CreateFilePathByArray(
		path.MigrationPath,
		path.GoModelPath,
		path.DefinitionPath,
	); err != nil {
		return err
	}

	model := cliTypes.NewTableTemplate(&cliTypes.ICreateTable{
		Version: i.Version,
		ModelNm: "user",
	})

	modelFp := filepath.Join(path.DefinitionPath, "user.yml")
	if err := cliFn.SerializeByYaml(modelFp, model); err != nil {
		return err
	}

	return nil
}

//func CreateConfigure(i *InitConfigurationOptions) (err error) {
//	if _, err := os.Stat(i.FilePath); err == nil {
//		return ErrAlreadyExistConfigureFile
//	}
//
//	file, err := os.Create(i.FilePath)
//	if err != nil {
//		return err
//	}
//
//	defer func() {
//		_ = file.Close()
//	}()
//
//	template := cliTypes.NewConfigureTemplate(i.Version)
//	if err := yaml.NewEncoder(file).Encode(template); err != nil {
//		return ErrFailGenerateConfigureFile
//	}
//
//	return
//}
//
//func CreateDefaultTable(fp string) error {
//	template, err := LoadConfig(fp)
//	if err != nil {
//		return err
//	}
//
//	databaseNm := cliTypes.DatabaseNm("sample")
//	modelPath, err := template.GetDatabaseModelPath(databaseNm)
//	if err != nil {
//		return err
//	}
//
//	// 파일 생성
//	if err := CreateFilePath(modelPath); err != nil {
//		return err
//	}
//
//	file, err := CreateFile(filepath.Join(modelPath, "sample.yml"))
//	if err != nil {
//		return err
//	}
//
//	defer file.Close()
//
//	if err := yaml.NewEncoder(file).Encode(cliTypes.NewDefTable(&cliTypes.ICreateTable{
//		Version:  template.Version,
//		Database: databaseNm,
//	})); err != nil {
//		return err
//	}
//
//	return nil
//}
