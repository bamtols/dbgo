package cliFields

type (
	FilePath string
)

func NewFilePath(v string) *FilePath {
	res := FilePath(v)
	return &res
}
