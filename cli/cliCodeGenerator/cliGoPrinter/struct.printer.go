package cliGoPrinter

import (
	"fmt"
	"github.com/bamtols/dbgo/func/parserFn"
)

type (
	StructPrinter struct {
		StructNm  string
		FieldList []PrinterFn
		Method    []PrinterFn
	}
)

func NewStructPrinter(structNm string) *StructPrinter {
	return &StructPrinter{
		StructNm:  parserFn.GoPublicCamelCase(structNm),
		FieldList: make([]PrinterFn, 0),
		Method:    make([]PrinterFn, 0),
	}
}

func (x *StructPrinter) AddField(fieldNm, fieldType string) {
	x.FieldList = append(x.FieldList, NewFieldPrinter(fieldNm, fieldType))
}

func (x *StructPrinter) AddMethod(m PrinterFn) {
	x.Method = append(x.Method, m)
}

func (x *StructPrinter) PrintLn() string {
	return fmt.Sprintf("%s\n", x.Print())
}

func (x *StructPrinter) Print() string {
	body := fmt.Sprintf("type %s struct {\n", x.StructNm)

	for _, fieldPrinter := range x.FieldList {
		body += fmt.Sprintf("\t%s", fieldPrinter.PrintLn())
	}

	body += "}"
	return body
}
