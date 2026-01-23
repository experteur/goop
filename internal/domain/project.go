package domain

import "time"

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

	// LastUpdated is the modification time of the project.md file
	LastUpdated time.Time

	// Tags are project categorization labels (Optional)
	Tags []string
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
