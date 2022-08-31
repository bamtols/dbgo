package main

import (
	"fmt"
	"github.com/bamtols/dbgo/cli/cliTypes/cliTmp"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func newCmdInit(pwd string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "init",
		Short: "generate initial configuration file",
	}

	configPath := filepath.Join(pwd, "config.yml")
	rootCmd.Flags().StringVarP(&configPath, "config", "c", configPath, fmt.Sprintf("configure file path (default : %s)", configPath))

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		config := cliTmp.NewConfigureTemplate()
		err := config.Save(configPath, true)
		onErrorExit(err)
		_, _ = fmt.Fprintf(os.Stderr, "generated configure file path : %s\n", configPath)
	}

	return rootCmd
}
