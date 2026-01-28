package views

import (
	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui/components"
	"github.com/rivo/tview"
)

type HomeView struct {
	*tview.Flex
	table             *components.ProjectTable
	onProjectSelected func(*domain.Project)
	projects          []*domain.Project
}

func NewHomeView() *HomeView {
	table := components.NewProjectTable()

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true)

	hv := &HomeView{
		Flex:     flex,
		table:    table,
		projects: []*domain.Project{},
	}

	table.OnSelected(func(project *domain.Project) {
		if hv.onProjectSelected != nil {
			hv.onProjectSelected(project)
		}
	})

	return hv
}

func (hv *HomeView) SetProjects(projects []*domain.Project) error {
	hv.projects = projects
	hv.table.SetProjects(projects)
	return nil
}
func (hv *HomeView) OnProjectSelected(handler func(*domain.Project)) {
	hv.onProjectSelected = handler
}

// func (hv *HomeView) OnQuit(handler func())
