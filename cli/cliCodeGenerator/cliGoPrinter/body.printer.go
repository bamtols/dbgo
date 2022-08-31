package cliGoPrinter

import "fmt"

type (
	BodyPrinter string
)

func NewBodyPrinter() *BodyPrinter {
	res := BodyPrinter("")
	return &res
}

func (x *BodyPrinter) Tab() *BodyPrinter {
	*x = BodyPrinter(fmt.Sprintf("%s\t", *x))
	return x
}

func (x *BodyPrinter) Ln() *BodyPrinter {
	*x = BodyPrinter(fmt.Sprintf("%s\n", *x))
	return x
}

func (x *BodyPrinter) PrintLn() string {
	return fmt.Sprintf("%s\n", x.Print())
}

func (x *BodyPrinter) Print() string {
	return fmt.Sprintf("{\n%s\n}", *x)
}
