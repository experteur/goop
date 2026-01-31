package app

import (
	"fmt"

	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/markdown"
	"github.com/experteur/goop/internal/ui/views"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app            *tview.Application
	pages          *tview.Pages
	homeView       *views.HomeView
	boardView      *views.BoardView
	projects       []*domain.Project
	currentProject *domain.Project
}

func New(projectsDir string) (*App, error) {
	a := &App{
		app:   tview.NewApplication(),
		pages: tview.NewPages(),
	}
    var err error
    a.projects, err = markdown.LoadProjects(projectsDir)
	if err != nil {
		return nil, fmt.Errorf("projects failed to load: %v", err)
	}
    a.homeView = views.NewHomeView()
    a.boardView = views.NewBoardView(a.app)

    a.homeView.OnProjectSelected(func(p *domain.Project) {
        a.currentProject = p
        a.boardView.SetProject(p)
        a.pages.SwitchToPage("board")
        a.app.SetFocus(a.boardView)
    })

    a.boardView.OnBack(func() {
        a.pages.SwitchToPage("home")
        a.app.SetFocus(a.homeView)
    })

    a.pages.AddPage("home", a.homeView, true, true)
    a.pages.AddPage("board", a.boardView, true, false)

    a.homeView.SetProjects(a.projects)

    return a, nil
}

func (a *App) Run() error {
	a.homeView.SetInputCapture(a.handleGlobalKeys)
	a.boardView.SetInputCapture(a.handleGlobalKeys)
	a.app.SetRoot(a.pages, true)
	if err := a.app.Run(); err != nil {
		return fmt.Errorf("app running failure: %v", err)
	}
	return nil
}


func (a *App) handleGlobalKeys(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'q', 'Q':
		// Quit the application
		a.app.Stop()
		return nil
	}

	// Let the event propagate to focused components
	return event
}
