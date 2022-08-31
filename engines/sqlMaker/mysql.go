package sqlMaker

import (
	"fmt"
)

type (
	MySQL struct {
	}
)

func (x *MySQL) CreateTable(v any, columnNm ...fmt.Stringer) error {
	//TODO implement me
	panic("implement me")
}
