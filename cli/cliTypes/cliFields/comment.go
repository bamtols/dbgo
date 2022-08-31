package cliFields

type (
	Comment string
)

func NewComment(v string) *Comment {
	res := Comment(v)
	return &res
}
