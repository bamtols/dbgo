package errFn

import (
	"errors"
	"fmt"
	"strings"
)

func Bind(err error, value string) error {
	return errors.New(fmt.Sprintf("%s : %s", err.Error(), value))
}

func BindF(err error, format string, args ...any) error {
	msg := fmt.Sprintf(format, args...)
	return errors.New(fmt.Sprintf("%s:%s", err.Error(), msg))
}

func BindAll(err ...error) error {
	msg := ""

	for _, e := range err {
		if e == nil {
			continue
		}
		if msg == "" {
			msg = e.Error()
		} else {
			msg = fmt.Sprintf("%s,%s", msg, e.Error())
		}
	}

	v := strings.ReplaceAll(msg, " ", "")

	return errors.New(v)
}

func New(prefix, msg string) error {
	return errors.New(fmt.Sprintf("%s:%s", prefix, msg))
}
