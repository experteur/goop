package views

import (
	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui/components"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type BoardView struct {
	*tview.Flex
	app            *tview.Application
	header         *components.ProjectHeader
	projectPreview *components.ProjectPreview
	columns        []*components.KanbanColumn
	footer         *components.Footer
	focusedColumn  int
	project        *domain.Project
	onBack         func()
}

func NewBoardView(app *tview.Application) *BoardView { // needs app for SetFocus
	header := components.NewProjectHeader()
	footer := components.NewFooter([]*components.Shortcut{
		{Key: "h/l", Label: "switch column"},
		{Key: "j/k", Label: "navigate tasks"},
		// {Key: "o", Label: "overview"},
		{Key: "q", Label: "back"},
	})
	columns := []*components.KanbanColumn{
		components.NewKanbanColumn("BACKLOG", domain.StatusBacklog),
		components.NewKanbanColumn("IN PROGRESS", domain.StatusInProgress),
		components.NewKanbanColumn("BLOCKED", domain.StatusBlocked),
		components.NewKanbanColumn("DONE", domain.StatusDone),
	}

	// a.boardView.SetInputCapture(a.handleGlobalKeys)
	// Layout columns horizontally
	columnsRow := tview.NewFlex().SetDirection(tview.FlexColumn)
	for _, col := range columns {
		columnsRow.AddItem(col, 0, 1, true)
	}

	// Main layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 1, false).
		AddItem(columnsRow, 0, 8, true).
		AddItem(footer, 0, 1, false)

	bv := &BoardView{
		Flex:    flex,
		app:     app,
		header:  header,
		columns: columns,
		footer:  footer,
	}

    bv.SetInputCapture(bv.handleKeyEvent)
	// list.OnBack(func() {
	// 	if bv.onBack != nil {
	// 		bv.onBack()
	// 	}
	// })
	return bv
}

func (bv *BoardView) handleKeyEvent(event *tcell.EventKey) *tcell.EventKey {
    currentList := bv.columns[bv.focusedColumn]
	switch event.Key() {
	case tcell.KeyLeft:
		bv.focusPreviousColumn()
		return nil
	case tcell.KeyRight:
		bv.focusNextColumn()
		return nil
	case tcell.KeyRune:
		switch event.Rune() {
		case 'h':
			bv.focusPreviousColumn()
			return nil
		case 'l':
			bv.focusNextColumn()
			return nil
		case 'j':
			current := currentList.GetCurrentItem()
			if current < currentList.GetItemCount()-1 {
				currentList.SetCurrentItem(current + 1)
			} else {
				currentList.SetCurrentItem(0)
			}
			return nil
		case 'k':
			current := currentList.GetCurrentItem()
			if current > 0 {
				currentList.SetCurrentItem(current - 1)
			} else {
				currentList.SetCurrentItem(currentList.GetItemCount() - 1)
			}
			return nil
		case 'q':
			if bv.onBack != nil {
				bv.onBack()
			}
			return nil
			// case 'o':
			//     if bv.onSwitchToOverview != nil {
			//         bv.onSwitchToOverview()
			//     }
			//     return nil
		}
	}
	return event
}

func (bv *BoardView) focusNextColumn() {
	if bv.focusedColumn < len(bv.columns)-1 {
		bv.focusedColumn++
		bv.app.SetFocus(bv.columns[bv.focusedColumn])
	}
}

func (bv *BoardView) focusPreviousColumn() {
	if bv.focusedColumn > 0 {
		bv.focusedColumn--
		bv.app.SetFocus(bv.columns[bv.focusedColumn])
	}
}

func (bv *BoardView) SetProject(project *domain.Project) {
	bv.project = project
	bv.header.SetProject(project)

	// Distribute tasks to columns
	for _, col := range bv.columns {
		tasks := project.TasksByStatus(col.Status())
		col.SetTasks(tasks)
	}

	// Focus first column
	bv.focusedColumn = 0
	bv.app.SetFocus(bv.columns[0])
}

// Callback registration
func (bv *BoardView) OnBack(handler func()) {
	bv.onBack = handler
}
