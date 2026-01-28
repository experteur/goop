package domain

import (
	"fmt"
	"time"
)

type ProjectStatus string

const (
	StatusActive   ProjectStatus = "active"
	StatusInactive ProjectStatus = "inactive"
	StatusArchived ProjectStatus = "archived"
)

// Representation of a single project.md
type Project struct {
	// Name is the title of the project extracted from the first H1 heading
	Name string

	// Description pulled from the first paragraph after Title
	Description string

	// Status is the lifecycle state (active|inactive|archived)
	Status ProjectStatus

	// Owner of the project (Optional)
	Owner string

	// Path is the absolute filestystem path to the project.md file
	Path string

	// TargetDate is the planned completion date (Optional)
	TargetDate time.Time

	// LastUpdated is the modification time of the project.md file
	LastUpdated time.Time

	// Tags are project categorization labels (Optional)
	Tags []string

	// Tasks are the lists of tasks for a project(Optional)
	Tasks []*Task
}

func (p *Project) IsActive() bool {
	return p.Status == StatusActive
}

func (p *Project) IsInactive() bool {
	return p.Status == StatusInactive
}

func (p *Project) IsArchived() bool {
	return p.Status == StatusArchived
}

func (p *Project) TasksByStatus(status TaskStatus) []*Task {
	var filtered []*Task
	for _, task := range p.Tasks {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

// CalculateProgress returns progress percentage based on completed tasks
func (p *Project) CalculateProgress() int {
	if len(p.Tasks) == 0 {
		return 0
	}
	done := len(p.TasksByStatus(StatusDone))
	return (done * 100) / len(p.Tasks)
}

// ProgressDisplay returns formatted string like "14/22 (64%)"
func (p *Project) ProgressDisplay() string {
	done := len(p.TasksByStatus(StatusDone))
	total := len(p.Tasks)
	progress := p.CalculateProgress()
	if total == 0 {
		return "-"
	}
	return fmt.Sprintf("%d/%d (%d%%)", done, total, progress)
}
