package markdown

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/experteur/goop/internal/domain"
	"gopkg.in/yaml.v3"
)

func LoadProjects(baseDir string) ([]*domain.Project, error) {
	// Verify base directory exists
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("base directory does not exist: %s", baseDir)
	}

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read base directory: %w", err)
	}

	var projects []*domain.Project
	var errors []error

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		projectDir := filepath.Join(baseDir, entry.Name())
		projectFilePath := filepath.Join(projectDir, "project.md")
		projectTaskPath := filepath.Join(projectDir, "tasks.yaml")

		if _, err := os.Stat(projectFilePath); os.IsNotExist(err) {
			continue
		}
        tasks, err := loadTasks(projectTaskPath)

		project, err := loadProject(projectFilePath)

        project.Tasks = tasks

		if err != nil {
			errors = append(errors, fmt.Errorf("failed to load %s: %w", entry.Name(), err))
			continue
		}

		projects = append(projects, project)

	}
	if len(errors) > 0 {
		if len(projects) == 0 {
			return nil, fmt.Errorf("failed to load any projects: %v", errors)
		}
		// TODO: Logging of errors
	}

	return projects, nil
}

func loadProject(path string) (*domain.Project, error) {
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

	title, description, err := extractTitleAndDescription(body)

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

func loadTasks(path string) ([]*domain.Task, error) {
    var tasks []*domain.Task
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return tasks, nil
    }

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
    err = yaml.Unmarshal(data, &tasks)

    fmt.Printf("Topics: %v\n", tasks)
	for i, topic := range tasks {
		fmt.Printf("Topic %d: %s\n", i+1, topic)
	}
    return tasks, nil
}
