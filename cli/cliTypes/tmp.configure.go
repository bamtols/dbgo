package cliTypes

type (
	Configure struct {
		Version        Version          `yaml:"Version" json:"version"`
		GoModelPath    *string          `yaml:"GoModelPath" json:"goModelPath"`
		DefinitionPath *string          `yaml:"DefinitionPath" json:"definitionPath"`
		MigrationPath  *string          `yaml:"MigrationPath,omitempty" json:"migrationPath,omitempty"`
		NameStrategy   *NameStrategy    `yaml:"NameStrategy,omitempty" json:"nameStrategy,omitempty"`
		GoType         GoTypePackageMap `yaml:"GoTypes,omitempty" json:"goTypes,omitempty"`
		Database       DatabaseMap      `yaml:"Databases,omitempty" json:"databases,omitempty"`

		// Post set data
		modelConfigure *ModelConfigure
	}
)
