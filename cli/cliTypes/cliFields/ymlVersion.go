package cliFields

type (
	YmlVersion     string
	YmlVersionList []YmlVersion
)

const (
	YmlVersionNow YmlVersion = "1.0"
)

var (
	YmlVersionAvailableList = YmlVersionList{
		YmlVersionNow,
	}
)

func (x *YmlVersion) IsUsable() bool {
	for _, version := range YmlVersionAvailableList {
		if *x == version {
			return true
		}
	}
	return false
}
