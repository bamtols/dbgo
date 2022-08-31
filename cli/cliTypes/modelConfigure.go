package cliTypes

type (
	ModelConfigure struct {
		GoTypeMap GoTypeMap `yaml:"-" json:"-"`
		ModelMap  ModelMap  `yaml:"-" json:"-"`
	}
)

func (x *ModelConfigure) HasModel(nm ModelNm) bool {
	_, isOk := x.ModelMap[nm]
	return isOk
}
