package cliFn

import "github.com/bamtols/dbgo/func/errFn"

const errPrefix = "cliFn"

var (
	ErrAlreadyExist = errFn.New(errPrefix, "AlreadyExist")
)
