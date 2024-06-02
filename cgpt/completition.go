package cgpt

import (
	"context"
	"log"
	"time"

	"github.com/briandowns/spinner"
	openai "github.com/sashabaranov/go-openai"
	"github.com/webnuke/cli-bud/auth"
)

// send a message to the OpenAI API
func send(messages []openai.ChatCompletionMessage) string {
	client := openai.NewClient(auth.GetToken("openai-token"))

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    Model,
			Messages: messages,
		},
	)
	if err != nil {
		log.Fatalf("ChatCompletion error: %v\n", err)
	}

	return resp.Choices[0].Message.Content
}
