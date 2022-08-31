package cliFields

import "github.com/bamtols/dbgo/func/errFn"

const errPrefix = "cliField"

var (
	ErrNotFound = errFn.New(errPrefix, "NotFound")
	ErrInvalid  = errFn.New(errPrefix, "Invalid")
)
