package cliGoPrinter

import "fmt"

type (
	ImportPrinter []string
)

func NewImportPrinter() *ImportPrinter {
	res := make(ImportPrinter, 0)
	return &res
}

func (x *ImportPrinter) AddPackage(packageNm string) {
	*x = append(*x, packageNm)
}

func (x *ImportPrinter) PrintLn() string {
	return fmt.Sprintf("%s\n", x.Print())
}

func (x *ImportPrinter) Print() string {
	body := fmt.Sprint("import (\n")
	for _, packageNm := range *x {
		body += fmt.Sprintf("\t\"%s\"\n", packageNm)
	}
	body += ")"
	return body
}
