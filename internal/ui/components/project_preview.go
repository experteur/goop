package components

import (
	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui"
	"github.com/rivo/tview"
)

type ProjectPreview struct {
	*tview.TextView
	currentProject *domain.Project
}

func NewProjectPreview() *ProjectPreview {
    textView := tview.NewTextView()
	textView.SetBorder(true)
	textView.SetTitle(" Preview ")
	textView.SetDynamicColors(true)
    textView.SetTitleColor(ui.Theme.TitleColor)
	textView.SetBorderColor(ui.Theme.BorderColor)
	textView.SetWordWrap(true)
	textView.SetText("\n  Select a project to view details")
    return &ProjectPreview{
        TextView: textView,
    }
}

func (pp *ProjectPreview) SetProject(project *domain.Project) {
    pp.currentProject = project
    pp.render()
}

func (pp *ProjectPreview) render() {
    pp.SetText(pp.currentProject.Name)
}
