package components

import (
	"fmt"
	"sort"

	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ProjectList struct {
	*tview.List
	projects       []*domain.Project
	onSelected     func(*domain.Project)
	statusSections map[domain.ProjectStatus]int
}

func NewProjectList() *ProjectList {
	list := tview.NewList()
	list.SetBorder(true)
	list.SetTitle(" Projects ")
	list.SetTitleColor(ui.Theme.TitleColor)
	list.SetBorderColor(ui.Theme.BorderColor)

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'j':
			current := list.GetCurrentItem()
			if current < list.GetItemCount()-1 {
				list.SetCurrentItem(current + 1)
			} else {
				list.SetCurrentItem(0)
			}
			return nil
		case 'k':
			current := list.GetCurrentItem()
			if current > 0 {
				list.SetCurrentItem(current - 1)
			} else {
				list.SetCurrentItem(list.GetItemCount() - 1)
			}
			return nil
		}
		return event
	})

	return &ProjectList{
		List:           list,
		statusSections: make(map[domain.ProjectStatus]int),
	}
}

func (pl *ProjectList) OnSelected(callback func(*domain.Project)) {
	pl.onSelected = callback
}

func (pl *ProjectList) SetProjects(projects []*domain.Project) error {
	pl.projects = projects
	pl.statusSections = make(map[domain.ProjectStatus]int)
	pl.Clear()

	// Group projects by status
	grouped := make(map[domain.ProjectStatus][]*domain.Project)
	for _, p := range projects {
		grouped[p.Status] = append(grouped[p.Status], p)
	}

	// Sort projects within each group by name
	for _, group := range grouped {
		sort.Slice(group, func(i, j int) bool {
			return group[i].Name < group[j].Name
		})
	}

	// Display in order: active, inactive, archived
	statuses := []domain.ProjectStatus{
		domain.StatusActive,
		domain.StatusInactive,
		domain.StatusArchived,
	}

	currentIndex := 0
	for _, status := range statuses {
		projectsInStatus := grouped[status]
		if len(projectsInStatus) == 0 {
			continue
		}

		// Add section header
		pl.statusSections[status] = currentIndex
		pl.addSectionHeader(status, len(projectsInStatus))
		currentIndex++

		// Add projects in this section
		for _, project := range projectsInStatus {
			pl.addProject(project)
			currentIndex++
		}
	}

	// Set up selection handler
	pl.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		if pl.onSelected != nil {
			selectedProject := pl.getProjectAtIndex(index)
			if selectedProject != nil {
				pl.onSelected(selectedProject)
			}
		}
	})

	// Set up change handler (for navigation updates)
	pl.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		if pl.onSelected != nil {
			selectedProject := pl.getProjectAtIndex(index)
			if selectedProject != nil {
				pl.onSelected(selectedProject)
			}
		}
	})

	return nil
}

func (pl *ProjectList) addSectionHeader(status domain.ProjectStatus, count int) {
	header := fmt.Sprintf("[::b]%s (%d)[-::-]", statusName(status), count)
	pl.AddItem(header, "", 0, nil)
	pl.List.SetItemText(pl.GetItemCount()-1, header, "")
}

func (pl *ProjectList) addProject(project *domain.Project) {
	// statusColor := pl.getStatusColor(project.Status)

	// Format: "• ProjectName"
	// mainText := fmt.Sprintf("[%s]• [-]%s", colorToTag(statusColor), project.Name)
	mainText := fmt.Sprintf("- %s", project.Name)

	// Secondary text shows owner if present
	secondaryText := ""
	if project.Owner != "" {
		secondaryText = fmt.Sprintf("  @%s", project.Owner)
	}

	pl.AddItem(mainText, secondaryText, 0, nil)
}

// Convert status to string
func statusName(status domain.ProjectStatus) string {
	switch status {
	case domain.StatusActive:
		return "ACTIVE"
	case domain.StatusInactive:
		return "INACTIVE"
	case domain.StatusArchived:
		return "ARCHIVED"
	default:
		return "UNKNOWN"
	}
}

func (pl *ProjectList) getProjectAtIndex(index int) *domain.Project {
	grouped := make(map[domain.ProjectStatus][]*domain.Project)
	for _, p := range pl.projects {
		grouped[p.Status] = append(grouped[p.Status], p)
	}

	// Sort projects within each group
	for _, group := range grouped {
		sort.Slice(group, func(i, j int) bool {
			return group[i].Name < group[j].Name
		})
	}

	statuses := []domain.ProjectStatus{
		domain.StatusActive,
		domain.StatusInactive,
		domain.StatusArchived,
	}

	currentListIndex := 0
	for _, status := range statuses {
		projectsInStatus := grouped[status]
		if len(projectsInStatus) == 0 {
			continue
		}

		// Section header
		if currentListIndex == index {
			return nil // This is a header, not a project
		}
		currentListIndex++

		// Projects in section
		for _, project := range projectsInStatus {
			if currentListIndex == index {
				return project
			}
			currentListIndex++
		}
	}

	return nil
}
