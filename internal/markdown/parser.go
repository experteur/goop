package markdown

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
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

func extractTitleAndDescription(body []byte) (string, string, error){
    md := goldmark.New()
	reader := text.NewReader(body)
	doc := md.Parser().Parse(reader)

	var title string
	var description string

    ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
        if !entering {
            return ast.WalkContinue, nil
        }
        switch node := n.(type) {
        case *ast.Heading:
            if node.Level == 1 && title == "" {
                title = extractText(node, body)
            }
        case *ast.Paragraph:
            if description == "" && title != "" {
                description = extractText(node, body)
            }
        }
        return ast.WalkContinue, nil
    })

	if title == "" {
		return "", "", fmt.Errorf("no H1 heading found in markdown")
	}

	return title, description, nil
}

func extractText(node ast.Node, source []byte) string {
	var buf bytes.Buffer
	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		if text, ok := child.(*ast.Text); ok {
			buf.Write(text.Segment.Value(source))
		}
	}
	return strings.TrimSpace(buf.String())
}

