package cmds

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
	"github.com/tiaguinho/cli-bud/cgpt"
)

// chat is the struct for the chat command
type chat struct {
	Messages []openai.ChatCompletionMessage
}

// Start starts the chat with ChatGPT
func (c *chat) Start() {
	for {
		var response string

		fmt.Print("You: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Println("--------------------")
		fmt.Printf("\n")
		c.Messages, response = cgpt.Chat(c.Messages, text)
		fmt.Printf("ChatGPT: %s \n", response)
		fmt.Printf("--------------------\n\n")
	}
}

// StartChat starts a new chat with ChatGPT
func StartChat() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chat",
		Short: "Chat with ChatGPT",
		Run: func(cmd *cobra.Command, args []string) {
			chat := &chat{
				Messages: []openai.ChatCompletionMessage{},
			}

			chat.Start()
		},
	}

	return cmd
}
