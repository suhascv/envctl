package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var useCommand = &cobra.Command{
	Use:   "use",
	Short: "use a stage config",
	Long: `use env config from specific state, useful if chained with your app commands 
    eg: envctl use dev && yarn dev`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			panic("Invalid number of arguments")
		}

		stageFilename := fmt.Sprintf("env.%v.config.ts", args[0])

		stageInfo, err := os.Stat(stageFilename)

		if err != nil || stageInfo.IsDir() {
			panic("Invalid stage file")
		}

		sourceFile, err := os.Open(stageFilename)
		if err != nil {
			panic("error reading " + stageFilename)
		}

		destnFile, err := os.Create("env.config.ts")

		if err != nil {
			panic("error creating env.config.ts")
		}

		_, err = io.Copy(destnFile, sourceFile)

		if err != nil {
			panic("error copying " + stageFilename + " to env.config.ts")
		}

		fmt.Println("using " + stageFilename)

	},
}

func init() {
	rootCmd.AddCommand(useCommand)
}
