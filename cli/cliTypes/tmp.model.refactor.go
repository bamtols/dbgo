package cliTypes

type (
	TmpModel struct {
		Version    Version       `yaml:"Version" json:"version"`
		Columns    ColumnMap     `yaml:"Columns" json:"columns"`
		Indexes    ColumnNmMap   `yaml:"Indexes,omitempty" json:"indexes,omitempty"`
		Uniques    ColumnNmMap   `yaml:"Uniques,omitempty" json:"uniques,omitempty"`
		Binds      BindMap       `yaml:"Binds,omitempty" json:"binds,omitempty"`
		Hooks      *Hooks        `yaml:"Hooks,omitempty" json:"hook,omitempty"`
		Migrations MigrationList `yaml:"Migrations,omitempty" json:"migrations,omitempty"`

		// ref
		ModelNm ModelNm `yaml:"-" json:"-"`
	}
)
