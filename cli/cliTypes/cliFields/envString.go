package cliFields

import (
	"os"
	"strings"
)

type (
	EnvString string
)

func NewEnvString(v string) EnvString {
	return EnvString(v)
}

func (x *EnvString) String() string {
	str := string(*x)
	if strings.HasPrefix(str, "!Env") {
		str = os.Getenv(str)
	}
	return str
}
