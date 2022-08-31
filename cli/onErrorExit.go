package main

import (
	"fmt"
	"os"
)

func onErrorExit(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
