package cliFields

type (
	ColumnNm string
)

func NewColumnNm(v string) *ColumnNm {
	res := ColumnNm(v)
	return &res
}
