package cgpt

import (
	openai "github.com/sashabaranov/go-openai"
)

// Chat format the chat messages and send to the OpenAI API
func Chat(history []openai.ChatCompletionMessage, message string) ([]openai.ChatCompletionMessage, string) {
	messages := append(history, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: message})

	response := send(messages)

	history = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: response})

	return history, response
}
