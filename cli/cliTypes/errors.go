package cliTypes

import "github.com/bamtols/dbgo/func/errFn"

const errPrefix = "cliTypes"

var (
	ErrNotFound = errFn.New(errPrefix, "NotFound")
	ErrInvalid  = errFn.New(errPrefix, "Invalid")
)
