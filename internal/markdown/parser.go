package markdown

import (
	"bytes"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Frontmatter struct {
	Status  string   `yaml:"status"`
	Owner   string   `yaml:"owner"`
	Created string   `yaml:"created"`
	Tags    []string `yaml:"tags"`
}

func extractFrontmatter(content []byte) (*Frontmatter, []byte, error) {
	lines := bytes.Split(content, []byte("\n"))

	// Check if file starts with frontmatter delimiter
	if len(lines) == 0 || string(bytes.TrimSpace(lines[0])) != "---" {
		// No frontmatter, return empty struct and full content
		return &Frontmatter{}, content, nil
	}

	endIdx := -1
	for i := 1; i < len(lines); i++ {
		if string(bytes.TrimSpace(lines[i])) == "---" {
			endIdx = i
			break
		}
	}

	if endIdx == -1 {
		return nil, nil, fmt.Errorf("frontmatter closing delimiter not found")
	}

	// Front matter
	frontMatterBytes := bytes.Join(lines[1:endIdx], []byte("\n"))
	var fm Frontmatter
	if err := yaml.Unmarshal(frontMatterBytes, &fm); err != nil {
		return nil, nil, fmt.Errorf("failed to parse frontmatter YAML: %w", err)
	}

	body := bytes.Join(lines[endIdx+1:], []byte("\n"))

	return &fm, body, nil
}
