package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var typeFilestring = `
export type EnvConfigType = {
  baseUrl: string;
  clientId: string;
};
`

var configFilestring = `
import { type EnvConfigType } from "./env.type.ts";

const EnvConfig: EnvConfigType = {
  baseUrl: "https://api.<stage>.example.com",
  clientId: "asfjhajshnfj871263478",
};

export default EnvConfig;
`

func CreateFile(filename *string, filecontent *string) {
	file, err := os.Create(*filename)

	if err != nil {
		panic(fmt.Sprintf("error creating %v", *filename))
	}

	defer file.Close()

	_, err = file.WriteString(*filecontent)

	if err != nil {
		panic(fmt.Sprintf("error writing to %v", *filename))
	}

	fmt.Println("created ", *filename)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize the env.type file and env.config files for all input stages",
	Long:  "Creates env.type.ts file and env.<stage>.config.ts file(s) for all input stages",
	Run: func(cmd *cobra.Command, args []string) {
		stages, err := cmd.Flags().GetStringArray("stage")

		if err != nil {
			panic("no stages detected")
		}

		typeFilename := "env.type.ts"
		configFilename := "env.config.ts"

		CreateFile(&typeFilename, &typeFilestring)
		CreateFile(&configFilename, &configFilestring)

		for _, stage := range stages {
			filename := "env." + stage + ".config.ts"
			CreateFile(&filename, &configFilestring)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCommand)

	initCommand.Flags().StringArray("stage", []string{}, "--stage dev")

}
