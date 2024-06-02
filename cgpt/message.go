package cgpt

import (
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

var (
	// Model is the OpenAI model name
	Model string
	// Debug is the debug mode
	Debug bool
)

// SendMessage formats a message to send to the OpenAI API
func SendMessage(system string, msgs ...string) string {
	messages := make([]openai.ChatCompletionMessage, 0)
	if system != "" {
		messages = append(messages, openai.ChatCompletionMessage{Role: "system", Content: system})
	}

	for _, msg := range msgs {
		messages = append(messages, openai.ChatCompletionMessage{Role: "user", Content: msg})
	}

	if Debug {
		fmt.Println("MODEL: ", Model)
		fmt.Println("SYSTEM: ", system)
		fmt.Println("MESSAGES: ", len(msgs), len(messages), messages)
	}

	return send(messages)
}
