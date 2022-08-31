package cliGoPrinter

import (
	"fmt"
	"github.com/bamtols/dbgo/func/parserFn"
)

type (
	FieldPrinter struct {
		FieldNm   string
		FieldType string
	}
)

func NewFieldPrinter(fieldNm, fieldType string) PrinterFn {
	return &FieldPrinter{
		FieldNm:   parserFn.GoPublicCamelCase(fieldNm),
		FieldType: fieldType,
	}
}

func (x *FieldPrinter) PrintLn() string {
	return fmt.Sprintf("%s\n", x.Print())
}

func (x *FieldPrinter) Print() string {
	return fmt.Sprintf("%s %s", x.FieldNm, x.FieldType)
}
