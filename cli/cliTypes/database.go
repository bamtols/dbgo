package cliTypes

type (
	Database struct {
		Connection Connection  `yaml:"Connection" json:"connection"`
		Models     ModelNmList `yaml:"Generated" json:"models"`
	}

	Connection struct {
		Host     string `yaml:"Host" json:"host"`
		Port     string `yaml:"Port" json:"port"`
		Username string `yaml:"Username" json:"username"`
		Password string `yaml:"Password" json:"password"`
		Driver   string `yaml:"Driver" json:"driver"`
	}

	DatabaseNm  string
	DatabaseMap map[DatabaseNm]Database
)

var (
	driverList = []string{
		"mysql",
	}
)
