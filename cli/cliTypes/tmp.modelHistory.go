package cliTypes

type (
	TmpModelHistory struct {
		Version        string           `yaml:"Version" json:"version"`
		HistoryVersion string           `yaml:"HistoryVersion" json:"historyVersion"`
		HistoryType    ModelHistoryType `yaml:"HistoryType" json:"historyType"`
		ModelNm        ModelNm          `yaml:"ModelNm" json:"modelNm"`
		Columns        ColumnMap        `yaml:"Columns" json:"columns"`
		Indexes        ColumnNmMap      `yaml:"Indexes,omitempty" json:"indexes,omitempty"`
		Uniques        ColumnNmMap      `yaml:"Uniques,omitempty" json:"uniques,omitempty"`
		Binds          BindMap          `yaml:"Binds,omitempty" json:"binds,omitempty"`
		Hooks          *Hooks           `yaml:"Hooks,omitempty" json:"hook,omitempty"`
		Migrations     MigrationList    `yaml:"Migrations,omitempty" json:"migrations,omitempty"`
	}

	ModelHistoryType string
)

const (
	ModelHistoryTypeInit      ModelHistoryType = "Init"
	ModelHistoryTypeMigration ModelHistoryType = "Migration"
)
