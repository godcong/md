package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/godcong/md/internal/configs"
)

var (
	rootCmd = &cobra.Command{
		Use:   "md",
		Short: "Markdown files manager",
		Long:  `A markdown file manager that helps organize and index your markdown files.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(configs.InitConfig)
	rootCmd.PersistentFlags().StringVar(&configs.ConfigFile, "config", "", "config file (default is .markdown.toml)")
}
