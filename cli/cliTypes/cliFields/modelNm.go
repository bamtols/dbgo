package cliFields

import "github.com/bamtols/dbgo/func/parserFn"

type (
	ModelNm     string
	ModelNmList []ModelNm
)

func (x *ModelNm) String() string {
	return parserFn.GoPublicCamelCase(string(*x))
}

func (x *ModelNmList) HasModel(m ModelNm) bool {
	for _, nm := range *x {
		if m.String() == nm.String() {
			return true
		}
	}
	return false
}
