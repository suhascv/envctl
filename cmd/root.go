package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "envctl",
	Short: "envctl is a cli tool to create and manage your stage/env specific typescript config files with type safety",
	Long: `A Fast and Flexible tool to 
    create and manage env/stage configuration files in your typescript project
    designed to integrate with your app commands locally and in CI`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
