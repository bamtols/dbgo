package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	version = "1.0"
)

func main() {
	cmdRoot := &cobra.Command{
		Use:     "dbgo",
		Example: "dbgo [command] [options]",
		Version: "0.1",
		Short:   "A Generator cli for DbGO",
	}

	pwd, err := os.Getwd()
	cobra.CheckErr(err)

	cmdRoot.AddCommand(newCmdInit(pwd))
	cmdRoot.AddCommand(newCmdGenV1(pwd))
	cmdRoot.AddCommand(versionCmd)

	if err := cmdRoot.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}

var (
	versionCmd = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprintf(os.Stderr, "version : %s\n", version)
		},
	}
)
