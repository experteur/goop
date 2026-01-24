package app

import (
	"fmt"

	"github.com/rivo/tview"
)

type App struct {
	tviewApp *tview.Application
}

func New() *App {
	return &App{
		tviewApp: tview.NewApplication(),
	}
}

func (a *App) Run() error {
	list := tview.NewList()
	list.AddItem("Project 1", "", '1', nil)
	list.AddItem("Project 2", "", '2', nil)
    a.tviewApp.SetRoot(list, true)
    if err := a.tviewApp.Run(); err != nil {
        return fmt.Errorf("app running failure: %v", err)
    }
    return nil
}
