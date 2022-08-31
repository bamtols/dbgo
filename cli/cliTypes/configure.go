package cliTypes

type (
	GoPackage        string
	GoTypePackageMap map[GoType]GoPackage
	NameStrategy     string
	NameStrategyList []NameStrategy
	Version          string
)

var (
	nameStrategySnakeCase = NameStrategy("snakeCase")
	nameStrategyCamelCase = NameStrategy("camelCase")

	nameStrategyList = NameStrategyList{
		nameStrategySnakeCase,
		nameStrategyCamelCase,
	}
)

// config default values
var (
	defNameStrategy   = nameStrategySnakeCase
	defGoModelPath    = "dbgos/models"
	defMigrationPath  = "dbgos/migrations"
	defDefinitionPath = "dbgos/defs"
)

func NewConfigureTemplate(version string) *Configure {
	res := NewConfigure(version)
	res.Database = DatabaseMap{
		"sample": {
			Connection{
				Host:     "110.110.111.123",
				Port:     "3306",
				Username: "root",
				Password: "password",
				Driver:   "mysql",
			},
			ModelNmList{
				"user",
			},
		},
	}
	res.GoType[GoType("Decimal")] = "github.com/shopspring/decimal.Decimal"
	return res
}

func NewConfigure(version string) *Configure {
	res := &Configure{
		Version: Version(version),
	}
	res.SetDefault()
	return res
}

func (x DatabaseNm) String() string {
	return string(x)
}

func (x *GoTypePackageMap) IsEmpty() bool {
	return len(*x) == 0
}

func (x GoPackage) String() string {
	return string(x)
}

func (x Version) String() string {
	return string(x)
}
