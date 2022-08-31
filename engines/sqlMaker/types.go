package sqlMaker

import (
	"fmt"
)

type (
	SqlMaker interface {
		CreateDatabase() (string, error)
		CreateTable(v any, columnNm ...fmt.Stringer) (string, error)

		Select() (string, error)
		Insert() (string, error)
		Update() (string, error)
		Delete() (string, error)
	}
)
