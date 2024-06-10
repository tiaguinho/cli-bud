package cmds

import (
	"log"

	"github.com/tiaguinho/cli-bud/cgpt"
	"github.com/tiaguinho/cli-bud/ioutil"
	"github.com/tiaguinho/cli-bud/prompts"
)

func withPathFlag(prompt *prompts.Prompt, path, output string) {
	files := ioutil.ReadPath(path)
	if len(files) == 0 {
		log.Fatal("No files found")
	}

	for _, file := range files {
		log.Println("Processing:", file.Filename)
		response := cgpt.SendMessage(prompt.System, file.Content)

		// If the prompt requires output, write the file
		if prompt.Output {
			aifile := ioutil.AIFile{
				Original: &file,
				Suffix:   prompt.Suffix,
				Content:  response,
			}

			ioutil.WriteAIFile(aifile, output)
		} else {
			// Otherwise, print the response
			log.Println(response)
		}
	}
}

func withContentFlag(prompt *prompts.Prompt, content, output string) {
	response := cgpt.SendMessage(prompt.System, content)

	// If the prompt requires output, write the file
	if prompt.Output {
		aifile := ioutil.AIFile{
			Content: response,
			Suffix:  prompt.Suffix,
		}

		ioutil.WriteAIFile(aifile, output)
	} else {
		// Otherwise, print the response
		log.Println(response)
	}
}
