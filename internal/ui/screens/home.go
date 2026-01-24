package screens

import (
	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui/components"
	"github.com/rivo/tview"
)

type HomeScreen struct {
	*tview.Flex
	projectList    *components.ProjectList
	projectPreview *components.ProjectPreview
	projects       []*domain.Project
}

func NewHomeScreen() *HomeScreen {
	projectList := components.NewProjectList()
	projectPreview := components.NewProjectPreview()

	flex := tview.NewFlex().
		AddItem(projectList, 0, 4, true).    // 40% width, focusable
		AddItem(projectPreview, 0, 6, false) // 60% width, not focusable

    hs := &HomeScreen{
		Flex:           flex,
		projectList:    projectList,
		projectPreview: projectPreview,
		projects:       []*domain.Project{},
	}

	projectList.OnSelected(func(project *domain.Project) {
		projectPreview.SetProject(project)
	})

    return hs
}

func (hs *HomeScreen) SetProjects(projects []*domain.Project) error {
    hs.projects = projects
    hs.projectList.SetProjects(projects)
	return nil
}
