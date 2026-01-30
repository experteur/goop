package components

import (
	"fmt"

	"github.com/experteur/goop/internal/domain"
	"github.com/rivo/tview"
)

type ProjectHeader struct {
	*tview.TextView
}

func NewProjectHeader() *ProjectHeader {
	text:= tview.NewTextView()
	text.SetBorder(true)
	ph := &ProjectHeader{
		text,
	}
	return ph
}

func (ph *ProjectHeader) Update(project *domain.Project) {
	ph.SetTitle(project.Name)
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
