package markdown

import (
	"os"
	"strings"

	"github.com/experteur/goop/internal/domain"
)

func LoadProject(path string) (*domain.Project, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	frontmatter, body, err := extractFrontmatter(data)
	status := normalizeStatus(frontmatter.Status)

	content := string(body)

	splitString := strings.Split(content, "\n\n")

	title := splitString[0]
	description := splitString[1]

	project := &domain.Project{
		Name:        title,
		Description: description,
		Status:      status,
		Owner:       frontmatter.Owner,
		Path:        path,
		LastUpdated: fileInfo.ModTime(),
		Tags:        frontmatter.Tags,
	}

	return project, nil
}

func normalizeStatus(status string) domain.ProjectStatus {
	status = strings.ToLower(strings.TrimSpace(status))
	switch status {
	case "active":
		return domain.StatusActive
	case "inactive":
		return domain.StatusInactive
	case "archived":
		return domain.StatusArchived
	default:
		return domain.StatusInactive
	}
}
