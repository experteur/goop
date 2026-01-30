package views

import (
	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui/components"
	"github.com/rivo/tview"
)

type BoardView struct {
	*tview.Flex
    projectHeader  *components.ProjectHeader
	// projectList    *components.ProjectList
	projectPreview *components.ProjectPreview
    projectFooter  *components.Footer
	projects       []*domain.Project
	onBack         func()
}

func NewBoardView(app *tview.Application) *BoardView { // needs app for SetFocus
    header := components.NewProjectHeader()
    // list := components.NewProjectList()
    preview := components.NewProjectPreview()
    footer := components.NewFooter([]*components.Shortcut{
        {Key: "enter", Label: "open"},
        {Key: "q", Label: "quit"},
    })
    flex := tview.NewFlex().
        AddItem(header, 0, 2, true).
        AddItem(preview, 0, 8, false).
        AddItem(footer, 0, 1, false)
    flex.SetDirection(tview.FlexRow)

	bv := &BoardView{
		Flex:           flex,
        projectHeader: header,
		// projectList:    list,
		projectPreview: preview,
        projectFooter: footer,
		projects:       []*domain.Project{},
	}

	// list.OnBack(func() {
	// 	if bv.onBack != nil {
	// 		bv.onBack()
	// 	}
	// })
	return bv
}
func (bv *BoardView) SetProject(project *domain.Project) {
    bv.projectHeader.Update(project)
    // bv.projectFooter.Update(project)
    // if bv.projectHeader.Update != nil {
    //     bv.projectHeader.Update(project)
    // }

}
func (bv *BoardView) OnBack(handler func()) {
	bv.onBack = handler
}

// func (bv *BoardView) OnSwitchToOverview(handler func())
