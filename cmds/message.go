package cmds

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tiaguinho/cb/cgpt"
	"github.com/tiaguinho/cb/ioutil"
)

// Message sends a message to ChatGPT
func Message() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:   "message",
		Short: "Send a message",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}

			content := strings.Join(args, " ")

			response := cgpt.SendMessage("", content)
			if response != "" {
				if output != "" {
					ioutil.WriteAIFile(ioutil.AIFile{Content: response}, output)
				} else {
					println(response)
				}
			}
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "Output file")

	return cmd
}
