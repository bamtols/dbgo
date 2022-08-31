package cliTmp

import "github.com/bamtols/dbgo/cli/cliTypes/cliFields"

type (
	ConfigurePath struct {
		Definition *cliFields.FilePath `yaml:"Definition,omitempty"`
		Migration  *cliFields.FilePath `yaml:"Migration,omitempty"`
		Generated  *cliFields.FilePath `yaml:"Generated,omitempty"`
	}
)
