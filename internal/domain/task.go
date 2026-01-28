package domain

import (
	"time"
)

type TaskStatus string

const (
	StatusBacklog    TaskStatus = "backlog"
	StatusInProgress TaskStatus = "in_progress"
	StatusBlocked    TaskStatus = "blocked"
	StatusDone       TaskStatus = "done"
)

// - id: task-1
//   title: "Migrate logging pipeline"
//   status: in_progress
//   description: "Move from old to new logging system"
//   notes: "Blocked waiting on vendor access"
//   created: 2026-01-15T10:00:00Z
//   updated: 2026-01-22T14:30:00Z

type Task struct {
	ID          string     `yaml:"id"`
	Title       string     `yaml:"title"`
	Status      TaskStatus `yaml:"status"`
	Description string     `yaml:"description"`
	Notes       string     `yaml:"notes"`
	CreatedAt   time.Time  `yaml:"createdat"`
	UpdatedAt   time.Time  `yaml:"updatedat"`
}
