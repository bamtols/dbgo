package cliGoPrinter

import (
	"fmt"
	"github.com/bamtols/dbgo/func/parserFn"
)

type (
	MethodPrinter struct {
		StructNm    string
		ReceiverNm  string
		MethodNm    string
		ParamValue  PrinterFn
		ReturnValue PrinterFn
		Body        PrinterFn
	}
)

func NewMethodPrinter(structNm, methodNm string) PrinterFn {
	return &MethodPrinter{
		StructNm:    parserFn.GoPublicCamelCase(structNm),
		ReceiverNm:  "x",
		MethodNm:    parserFn.GoPublicCamelCase(methodNm),
		ParamValue:  nil,
		ReturnValue: nil,
		Body:        nil,
	}
}

func (x *MethodPrinter) PrintLn() string {
	return fmt.Sprintf("%s\n", x.Print())
}

func (x *MethodPrinter) Print() string {
	body := fmt.Sprintf(
		"func (%s *%s) %s ",
		x.ReceiverNm,
		x.StructNm,
		x.MethodNm,
	)

	if x.ParamValue != nil {
		body += x.ParamValue.Print()
	} else {
		body += "()"
	}

	if x.ReturnValue != nil {
		body += x.ReturnValue.PrintLn()
	}

	if x.Body != nil {
		body += x.Body.PrintLn()
	} else {
		body += "{\n}"
	}

	return body
}
