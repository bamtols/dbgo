package cliCodeGenerator

import "fmt"

type (
	StrStream string
)

func (x *StrStream) Write(p []byte) (n int, err error) {
	*x = StrStream(fmt.Sprintf("%s", p))
	return len(string(p)), nil
}
