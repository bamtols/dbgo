package cliGoPrinter

import "github.com/bamtols/dbgo/func/errFn"

const errPrefix = "cliGenV1"

var (
	ErrInvalid          = errFn.New(errPrefix, "Invalid")
	ErrAlreadyGenerated = errFn.New(errPrefix, "AlreadyGenerated")
)
