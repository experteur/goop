package components

import (
	"fmt"

	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui"
	"github.com/rivo/tview"
)

type ProjectHeader struct {
	*tview.TextView
}

func NewProjectHeader() *ProjectHeader {
	textView:= tview.NewTextView()
	textView.SetBorder(true)
    textView.SetTitleColor(ui.Theme.TitleColor)
	textView.SetBorderColor(ui.Theme.BorderColor)
	ph := &ProjectHeader{
		textView,
	}
	return ph
}

func (ph *ProjectHeader) SetProject(project *domain.Project) {
    paddedTitle := fmt.Sprintf(" %s ", project.Name)
	ph.SetTitle(paddedTitle)
	ph.SetText(ph.buildHeader(project))
}

func (ph *ProjectHeader) buildHeader(p *domain.Project) string {
	// `ProjectName   [ACTIVE]   14/22 (64%)   ⚠   @alex`
	string := fmt.Sprintf("%s\t[%s]\t%d/%d (%d)\t%s\t@%s",
        p.Name,
        p.Status,
        len(p.TasksByStatus(domain.TaskStatus(domain.StatusActive))),
        len(p.Tasks),
        p.CalculateProgress(),
        "x",
        p.Owner,
    )
	return string
}
