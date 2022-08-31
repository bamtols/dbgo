package cliFileLoader

import (
	"github.com/bamtols/dbgo/func/errFn"
)

const errPrefix = "cliFileLoader"

var (
	ErrAlreadyExistConfigureFile = errFn.New(errPrefix, "AlreadyExistConfigureFile")
	ErrAlreadyDefaultTableFile   = errFn.New(errPrefix, "AlreadyDefaultTableFile")
	ErrFailGenerateConfigureFile = errFn.New(errPrefix, "FailGenerateConfigureFile ")
	ErrFailParsingConfigureFile  = errFn.New(errPrefix, "FailParsingConfigureFile ")
	ErrNotFoundConfigureFile     = errFn.New(errPrefix, "NotFoundConfigureFile")
	ErrFailParsing               = errFn.New(errPrefix, "FailParsing")
	ErrNoSupportVersion          = errFn.New(errPrefix, "NoSupportVersion")
)
