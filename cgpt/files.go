package cgpt

import (
	"context"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/webnuke/cli-bud/auth"
)

// UploadFile uploads a file to the OpenAI API
func UploadFile(path string) (openai.File, error) {
	client := openai.NewClient(auth.GetToken("openai-token"))

	finfo, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}

	return client.CreateFile(context.Background(), openai.FileRequest{
		FileName: finfo.Name(),
		FilePath: path,
		Purpose:  string(openai.PurposeAssistants),
	})
}
