package auth

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Token is the OpenAI API token
var Token string

// SetToken sets the OpenAI API token
func SetToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-token",
		Short: "Set the OpenAI API token",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}

			cacheDir, err := os.UserCacheDir()
			if err != nil {
				log.Fatal(err)
			}

			err = os.WriteFile(cacheDir+"/openai-token", []byte(args[0]), 0o644)
		},
	}

	return cmd
}

// GetToken gets the OpenAI API token
func GetToken(name string) string {
	// If the token is set as a flag, use that
	if Token != "" {
		return Token
	}

	// Otherwise, read from the cache
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}

	token, err := os.ReadFile(cacheDir + "/" + name)
	if err != nil {
		log.Fatal(err)
	}

	return string(token)
}
