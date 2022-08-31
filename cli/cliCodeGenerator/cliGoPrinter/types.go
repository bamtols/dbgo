package cliGoPrinter

type (
	PrinterFn interface {
		PrintLn() string
		Print() string
	}
)
