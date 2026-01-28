package views

import (
	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui/components"
	"github.com/rivo/tview"
)

type BoardView struct {
	*tview.Flex
	projectList    *components.ProjectList
	projectPreview *components.ProjectPreview
	projects       []*domain.Project
	onBack         func()
}

func NewBoardView(app *tview.Application) *BoardView { // needs app for SetFocus
    list := components.NewProjectList()
    preview := components.NewProjectPreview()
    flex := tview.NewFlex().
        AddItem(list, 0, 6, true).
        AddItem(preview, 0, 4, false)

	bv := &BoardView{
		Flex:           flex,
		projectList:    &components.ProjectList{},
		projectPreview: &components.ProjectPreview{},
		projects:       []*domain.Project{},
	}
	return bv
}
func (bv *BoardView) SetProject(project *domain.Project) {

}
func (bv *BoardView) OnBack(handler func()) {
	bv.onBack = handler
}

// func (bv *BoardView) OnSwitchToOverview(handler func())
