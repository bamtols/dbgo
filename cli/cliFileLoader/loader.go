package cliFileLoader

import (
	"github.com/bamtols/dbgo/cli/cliTypes"
	"github.com/bamtols/dbgo/func/errFn"
	"github.com/bamtols/dbgo/func/parserFn"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
)

func LoadConfig(fp string, version string) (*cliTypes.Configure, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, ErrNotFoundConfigureFile
	}

	config := cliTypes.NewConfigure(version)
	if err := yaml.NewDecoder(file).Decode(config); err != nil {
		return nil, ErrFailParsingConfigureFile
	}

	if config.Version.String() != version {
		return nil, ErrNoSupportVersion
	}

	path := config.AbsoluteFilePath()

	fileList, err := os.ReadDir(path.DefinitionPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range fileList {
		if entry.IsDir() {
			continue
		}

		modelFp := filepath.Join(path.DefinitionPath, entry.Name())

		l := strings.Split(entry.Name(), ".")
		if len(l) != 2 {
			return nil, errFn.BindF(ErrFailParsing, "ModelNm : fileName=%s", entry.Name())
		}

		modelNm := parserFn.GoPublicCamelCase(l[0])
		fileType := strings.ToLower(l[1])

		switch fileType {
		case "yml", "yaml":
			modelConfig, err := loadModel(
				modelFp,
				version,
				modelNm,
			)

			if err != nil {
				return nil, err
			}

			config.ModelConfigure().ModelMap[cliTypes.ModelNm(modelNm)] = *modelConfig
		default:
			continue
		}
	}

	return config, nil
}

func loadModel(fp, version, modelNm string) (res *cliTypes.TmpModel, err error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	res = cliTypes.NewModel(version, modelNm)
	if err := yaml.NewDecoder(file).Decode(res); err != nil {
		return nil, err
	}

	return
}
