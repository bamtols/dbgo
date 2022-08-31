package migrator

import "github.com/bamtols/dbgo/func/errFn"

const errPrefix = "migrator"

var (
	ErrFailTx = errFn.New(errPrefix, "FailTx")
)
