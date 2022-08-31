package cliGoPrinter

import "fmt"

type (
	Param struct {
		ParamNm   string
		ParamType string
	}

	ParamPrinter []Param
)

func NewParamPrinter() ParamPrinter {
	res := make(ParamPrinter, 0)
	return res
}

func (x *ParamPrinter) PrintLn() string {
	return fmt.Sprintf("%s\n", x.Print())
}

func (x *ParamPrinter) Print() string {
	body := "(\n"
	for _, param := range *x {
		body += fmt.Sprintf("%s %s,\n", param.ParamNm, param.ParamType)
	}
	body += ")"
	return body
}
