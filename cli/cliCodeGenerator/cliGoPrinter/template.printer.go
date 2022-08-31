package cliGoPrinter

import (
	"fmt"
	"github.com/bamtols/dbgo/func/errFn"
	"github.com/bamtols/dbgo/func/parserFn"
	"go/format"
)

type (
	Template struct {
		packageNm string
		imports   *ImportPrinter
		body      []PrinterFn
	}
)

func NewTemplatePrinter(version, packageNm string) (*Template, error) {
	if version != "1.0" {
		return nil, errFn.BindF(ErrInvalid, "TemplateVersion: %s", version)
	}

	return &Template{
		packageNm: parserFn.GoPrivateCamelCase(packageNm),
		imports:   NewImportPrinter(),
		body:      make([]PrinterFn, 0),
	}, nil
}

func (x *Template) AddImport(packageNm string) {
	x.imports.AddPackage(packageNm)
}

func (x *Template) AddBody(printer PrinterFn) {
	x.body = append(x.body, printer)
}

func (x *Template) PrintLn() string {
	res, err := format.Source([]byte(fmt.Sprintf("%s\n", x.Print())))
	if err != nil {
		panic(errFn.BindF(ErrInvalid, "go fmt error : err=%s", err.Error()))
	}
	return string(res)
}

func (x *Template) Print() string {
	body := fmt.Sprintf("package %s\n\n", x.packageNm)
	body += x.imports.PrintLn()

	for _, fn := range x.body {
		body += fn.PrintLn()
	}
	return body
}
