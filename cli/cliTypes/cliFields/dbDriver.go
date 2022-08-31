package cliFields

type (
	DbDriver string
)

const (
	DbDriverMySQL DbDriver = "mysql"
)

var (
	DbDriverAll = []DbDriver{
		DbDriverMySQL,
	}
)
