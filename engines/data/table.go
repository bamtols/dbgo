package data

import "github.com/bamtols/dbgo/cli/cliTypes"

type (
	TableMetadata struct {
		TableNm string
		Version uint
		Columns []ColumnMetadata
	}

	ColumnMetadata struct {
		Nm      string
		DBType  string
		Options string
		Tags    cliTypes.ColumnTagList
	}
)
