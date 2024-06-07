package cmds

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tiaguinho/cb/prompts"
)

// Do is the command for the do action
func Do() *cobra.Command {
	var path string

	cmd := &cobra.Command{
		Use:   "do",
		Short: "Do something with a file",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}

			action := args[0]
			prompt, err := prompts.GetPrompt("do", action)
			if err != nil {
				log.Fatal(err)
			}

			switch {
			// prompts with only path required
			case prompt.ValidateFlags([]string{"path"}):
				checkRequiredFlags("Path are required", path)

				withPathFlag(prompt, path, "")
			}
		},
	}

	cmd.Flags().StringVarP(&path, "path", "p", "", "Path to file(s)")

	return cmd
}
