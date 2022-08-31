package cliFields

type (
	NameStrategy string
)

const (
	NameStrategyCamelCase NameStrategy = "camelCase"
	NameStrategySnakeCase NameStrategy = "snakeCase"
)

func (x NameStrategy) Ptr() *NameStrategy {
	return &x
}
