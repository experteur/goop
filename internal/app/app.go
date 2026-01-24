package app

import (
	"fmt"

	"github.com/experteur/goop/internal/markdown"
	"github.com/experteur/goop/internal/ui/screens"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	tviewApp   *tview.Application
	homeScreen *screens.HomeScreen
	projectDir string
}

func New(projectsDir string) *App {
	return &App{
		tviewApp:   tview.NewApplication(),
		homeScreen: screens.NewHomeScreen(),
		projectDir: projectsDir,
	}
}

func (a *App) Run() error {
	projects, err := markdown.LoadProjects(a.projectDir)
	if err != nil {
		return fmt.Errorf("projects failed to load: %v", err)
	}

	a.homeScreen.SetProjects(projects)

	a.homeScreen.SetInputCapture(a.handleGlobalKeys)

	a.tviewApp.SetRoot(a.homeScreen, true)
	if err := a.tviewApp.Run(); err != nil {
		return fmt.Errorf("app running failure: %v", err)
	}
	return nil
}

func (a *App) handleGlobalKeys(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'q', 'Q':
		// Quit the application
		a.tviewApp.Stop()
		return nil
	}

	// Let the event propagate to focused components
	return event
}
