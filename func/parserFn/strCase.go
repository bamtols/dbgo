package parserFn

type StrCase struct {
	caseNm string
}

func NewStrCase(caseNm string) *StrCase {
	return &StrCase{
		caseNm: caseNm,
	}
}

func (x *StrCase) GoPublicNm(v string) string {
	return toCamelInitCase(v, true)
}

func (x *StrCase) Parse(v string) string {
	switch x.caseNm {
	case "camelCase":
		return CamelCase(v)
	case "snakeCase":
		return SnakeCase(v)
	default:
		panic("notFoundCase")
	}
}
