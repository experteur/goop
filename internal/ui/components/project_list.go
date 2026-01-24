package components

import (
	"github.com/experteur/goop/internal/domain"
	"github.com/rivo/tview"
)

type ProjectList struct {
    *tview.List
    projects []*domain.Project
}
