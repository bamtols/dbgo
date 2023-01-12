package cmdDef

type (
	FilePath string

	ConfigYAML struct {
		Version  Version  `yaml:"version"`
		Path     FilePath `yaml:"path"`
		Database Database `yaml:"database"`
		GoType   GoType   `yaml:"goType"`
	}
)
