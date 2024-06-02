package main

import (
	"log"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"

	"github.com/webnuke/cli-bud/auth"
	"github.com/webnuke/cli-bud/cgpt"
	"github.com/webnuke/cli-bud/cmds"
)

var rootCmd = &cobra.Command{}

func main() {
	rootCmd.PersistentFlags().StringVarP(&cgpt.Model, "model", "m", openai.GPT3Dot5Turbo, "OpenAI model name")
	rootCmd.PersistentFlags().StringVarP(&auth.Token, "token", "t", "", "OpenAI API token")
	rootCmd.PersistentFlags().BoolVarP(&cgpt.Debug, "debug", "d", false, "Enable debug mode")

	rootCmd.AddCommand(auth.SetToken())
	rootCmd.AddCommand(cmds.Create())
	rootCmd.AddCommand(cmds.Do())
	rootCmd.AddCommand(cmds.Message())
	rootCmd.AddCommand(cmds.StartChat())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
