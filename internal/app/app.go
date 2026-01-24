package app

import (
	"fmt"

	"github.com/experteur/goop/internal/markdown"
	"github.com/rivo/tview"
)

type App struct {
	tviewApp   *tview.Application
	projectDir string
}

func New(projectsDir string) *App {
	return &App{
		tviewApp:   tview.NewApplication(),
		projectDir: projectsDir,
	}
}

func (a *App) Run() error {
	projects, err := markdown.LoadProjects(a.projectDir)
	if err != nil {
		return fmt.Errorf("projects failed to load: %v", err)
	}
	list := tview.NewList()
	for i, project := range projects {
		list.AddItem(project.Name, project.Description, rune(i), nil)
	}
	a.tviewApp.SetRoot(list, true)
	if err := a.tviewApp.Run(); err != nil {
		return fmt.Errorf("app running failure: %v", err)
	}
	return nil
}
