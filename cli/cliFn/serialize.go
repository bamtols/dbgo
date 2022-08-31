package cliFn

import (
	"gopkg.in/yaml.v2"
	"os"
)

func SerializeByYaml(fp string, v any) error {
	file, err := os.Create(fp)
	if err != nil {
		return err
	}

	if err := yaml.NewEncoder(file).Encode(v); err != nil {
		return err
	}

	return nil
}
