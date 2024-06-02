package prompts

import (
	"embed"
	"fmt"
	"log"
	"slices"

	"gopkg.in/yaml.v2"
)

//go:embed templates/*.yaml
var promptsFolder embed.FS

// Prompt is a struct that represents a prompt
type Prompt struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Required []string `yaml:"required"`
	Output   bool     `yaml:"output"`
	Suffix   string   `yaml:"suffix"`
	System   string   `yaml:"system"`
}

// ValidateFlags validates the list of flags
func (p *Prompt) ValidateFlags(key []string) bool {
	if len(p.Required) != len(key) {
		return false
	}

	for _, r := range p.Required {
		if !slices.Contains(key, r) {
			return false
		}
	}

	return true
}

// GetPrompt gets a prompt from the file
func GetPrompt(file, action string) (*Prompt, error) {
	// Read the file
	data, err := promptsFolder.ReadFile(fmt.Sprintf("templates/%s.yaml", file))
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the data
	lprompts := make([]*Prompt, 0)
	err = yaml.Unmarshal(data, &lprompts)
	if err != nil {
		log.Fatal(err)
	}

	// Get the prompt
	var prompt *Prompt
	for _, p := range lprompts {
		if p.Type == action {
			prompt = p
			break
		}
	}

	if prompt == nil {
		return nil, fmt.Errorf("Prompt not found")
	}

	return prompt, nil
}
