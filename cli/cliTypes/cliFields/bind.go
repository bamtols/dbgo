package cliFields

type (
	Bind struct {
		Model      ModelNm   `yaml:"Model"`
		BindType   BindType  `yaml:"BindType"`
		ForeignKey *ColumnNm `yaml:"ForeignKey"`
		ReferKey   *ColumnNm `yaml:"ReferKey"`
	}

	BindMap      map[ColumnNm]Bind
	BindType     string
	BindTypeList []BindType
)

const (
	BindTypeHasOne    BindType = "hasOne"
	BindTypeHasMany   BindType = "hasMany"
	BindTypeBelongsTo BindType = "belongsTo"
)

var (
	BindTypeAll = BindTypeList{
		BindTypeHasOne,
		BindTypeHasMany,
		BindTypeBelongsTo,
	}
)
