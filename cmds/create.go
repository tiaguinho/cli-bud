package cmds

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/webnuke/cli-bud/prompts"
)

// Create is the command for the create action
func Create() *cobra.Command {
	var (
		path    string
		content string
		output  string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a file with the response from ChatGPT",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}

			action := args[0]
			prompt, err := prompts.GetPrompt("create", action)
			if err != nil {
				log.Fatal(err)
			}

			switch {
			// prompts with only path required
			case prompt.ValidateFlags([]string{"path"}):
				checkRequiredFlags("Path is required", path)

				withPathFlag(prompt, path, output)
				// prompts with content and output required
			case prompt.ValidateFlags([]string{"content", "output"}):
				checkRequiredFlags("Content and output are required", content, output)

				withContentFlag(prompt, content, output)
			}
		},
	}

	cmd.Flags().StringVarP(&path, "path", "p", "", "Path to file(s)")
	cmd.Flags().StringVarP(&content, "content", "c", "", "Message with the prompt")
	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file")

	return cmd
}
