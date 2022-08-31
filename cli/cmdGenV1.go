package main

import (
	"fmt"
	"github.com/bamtols/dbgo/cli/cliCodeGenerator"
	"github.com/bamtols/dbgo/cli/cliFileLoader"
	"github.com/bamtols/dbgo/cli/cliFn"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func newCmdGenV1(pwd string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "generate",
		Short: "generate go codes",
	}

	configPath := filepath.Join(pwd, "config.yml")
	rootCmd.Flags().StringVarP(&configPath, "config", "c", configPath, fmt.Sprintf("configure file path (default : %s)", configPath))

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		config, err := cliFileLoader.LoadConfig(configPath, version)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}

		genV1, err := cliCodeGenerator.NewV1(config)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}

		// 파일로 저장
		genFp := filepath.Join(config.AbsoluteFilePath().GoModelPath, "generated.go")
		file, err := cliFn.CreateFile(genFp)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}
		defer file.Close()

		_, err = file.WriteString(genV1.Print())
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}

		_, _ = fmt.Fprintf(os.Stderr, "generated!! %s\n", genFp)
	}

	return rootCmd
}
