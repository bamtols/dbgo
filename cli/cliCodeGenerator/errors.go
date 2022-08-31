package cliCodeGenerator

import "github.com/bamtols/dbgo/func/errFn"

const errPrefix = "cliGens"

var (
	ErrNotFound = errFn.New(errPrefix, "NotFound")
	ErrInvalid  = errFn.New(errPrefix, "Invalid")
)
