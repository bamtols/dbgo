package cliFields

type (
	DbType string
)

func NewDbType(v string) *DbType {
	res := DbType(v)
	return &res
}
