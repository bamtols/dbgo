package cliTmp

import "github.com/bamtols/dbgo/func/errFn"

const errPrefix = "cliTmp"

var (
	ErrAlreadyExistFile = errFn.New(errPrefix, "AlreadyExistFile")
)
