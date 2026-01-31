package components

import (
	"fmt"

	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui"
	"github.com/rivo/tview"
)

type KanbanColumn struct {
	*tview.List
	status domain.TaskStatus
	tasks  []*domain.Task
}

func NewKanbanColumn(title string, status domain.TaskStatus) *KanbanColumn {
	list := tview.NewList()
	list.SetBorder(true)
	list.SetTitle(fmt.Sprintf(" %s ", title))
	list.SetTitleColor(ui.Theme.TitleColor)
	list.SetBorderColor(ui.Theme.BorderColor)
	return &KanbanColumn{
		List:   list,
		status: status,
		tasks:  nil,
	}
}

func (k *KanbanColumn) Status() domain.TaskStatus {
	return k.status
}

func (k *KanbanColumn) SetTasks(tasks []*domain.Task) {
	k.tasks = tasks
	for _, task := range tasks {
		k.AddItem(task.Title, "- " + task.Description, 0, nil)
	}
}
