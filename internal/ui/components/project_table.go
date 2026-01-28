package components

import (
	"fmt"
	"sort"
	"time"

	"github.com/experteur/goop/internal/domain"
	"github.com/experteur/goop/internal/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ProjectTable struct {
	*tview.Table
	projects   []*domain.Project
	onSelected func(*domain.Project)
	onNavigate func(*domain.Project)
}

func NewProjectTable() *ProjectTable {
	table := tview.NewTable()
	table.SetBorder(true)
	table.SetTitle(" Projects ")
	table.SetTitleColor(ui.Theme.TitleColor)
	table.SetBorderColor(ui.Theme.BorderColor)
	table.SetSelectable(true, false) // rows selectable, columns not
	table.SetFixed(1, 0)             // fix header row

	pt := &ProjectTable{
		Table:      table,
		projects:   []*domain.Project{},
		onSelected: nil,
		onNavigate: nil,
	}
	// Set up keyboard handling
	table.SetInputCapture(pt.handleKeyEvent)

	// Set up selection handler
	table.SetSelectedFunc(func(row, col int) {
		if row > 0 && pt.onSelected != nil { // Skip header row
			project := pt.getProjectAtRow(row)
			if project != nil {
				pt.onSelected(project)
			}
		}
	})

	// Fire onNavigate callback when cursor moves (for preview updates, etc.)
	table.SetSelectionChangedFunc(func(row, col int) {
		if row > 0 && pt.onNavigate != nil {
			project := pt.getProjectAtRow(row)
			if project != nil {
				pt.onNavigate(project)
			}
		}
	})

	return pt
}

// SetProjects populates the table with projects
func (pt *ProjectTable) SetProjects(projects []*domain.Project) {
	pt.projects = projects
	pt.Clear()

	// Add header row
	pt.addHeaderRow()

	// Add separator row (optional, for visual appeal)
	// pt.addSeparatorRow()

	pt.sortByStatus()

	// Add project rows
	for _, project := range pt.projects {
		pt.addProjectRow(project)
	}

	// Select first project by default (skip header)
	// if len(projects) > 0 {
	// 	pt.Select(1, 0)
	// }
}

// OnSelected registers a callback for when a project is selected (Enter key)
func (pt *ProjectTable) OnSelected(handler func(*domain.Project)) {
	pt.onSelected = handler
}

// OnNavigate registers a callback for when the cursor moves to a different project
func (pt *ProjectTable) OnNavigate(handler func(*domain.Project)) {
	pt.onNavigate = handler
}

// GetSelectedProject returns the currently selected project
func (pt *ProjectTable) GetSelectedProject() *domain.Project {
	row, _ := pt.GetSelection()
	return pt.getProjectAtRow(row)
}

// Private helper methods
func (pt *ProjectTable) addHeaderRow() {
	headers := []string{"STATUS", "HEALTH", "PROG", "NAME", "OWNER", "UPDATED"}

	for col, header := range headers {
		cell := tview.NewTableCell(header).
			SetTextColor(ui.Theme.TextSecondary).
			SetAlign(tview.AlignLeft).
			SetSelectable(false).
			SetAttributes(tcell.AttrBold)
		pt.SetCell(0, col, cell)
	}
}
func (pt *ProjectTable) addProjectRow(project *domain.Project) {
	row := pt.GetRowCount()

	// Column 0: STATUS
	statusCell := tview.NewTableCell(pt.formatStatus(project.Status)).
		SetTextColor(pt.getStatusColor(project.Status)).
		SetAlign(tview.AlignLeft)
	pt.SetCell(row, 0, statusCell)

	// Column 1: HEALTH
	healthCell := tview.NewTableCell(pt.formatHealth(project)).
		SetTextColor(pt.getHealthColor(project)).
		SetAlign(tview.AlignCenter)
	pt.SetCell(row, 1, healthCell)

	// Column 2: PROG (Progress)
	progCell := tview.NewTableCell(project.ProgressDisplay()).
		SetTextColor(ui.Theme.TextPrimary).
		SetAlign(tview.AlignRight)
	pt.SetCell(row, 2, progCell)

	// Column 3: NAME
	nameCell := tview.NewTableCell(project.Name).
		SetTextColor(ui.Theme.TextPrimary).
		SetAlign(tview.AlignLeft).
		SetExpansion(1) // Allow this column to expand
	pt.SetCell(row, 3, nameCell)

	// Column 4: OWNER
	owner := project.Owner
	if owner == "" {
		owner = "–"
	}
	ownerCell := tview.NewTableCell(owner).
		SetTextColor(ui.Theme.TextSecondary).
		SetAlign(tview.AlignLeft)
	pt.SetCell(row, 4, ownerCell)

	// Column 5: UPDATED
	updatedCell := tview.NewTableCell(formatRelativeTime(project.LastUpdated)).
		SetTextColor(ui.Theme.TextDim).
		SetAlign(tview.AlignRight)
	pt.SetCell(row, 5, updatedCell)
}
func (pt *ProjectTable) formatStatus(status domain.ProjectStatus) string {
	switch status {
	case domain.StatusActive:
		return "Active"
	case domain.StatusInactive:
		return "On Hold"
	case domain.StatusArchived:
		return "Archived"
	default:
		return "Unknown"
	}
}
func (pt *ProjectTable) getStatusColor(status domain.ProjectStatus) tcell.Color {
	switch status {
	case domain.StatusActive:
		return ui.Theme.StatusActive
	case domain.StatusInactive:
		return ui.Theme.StatusInactive
	case domain.StatusArchived:
		return ui.Theme.StatusArchived
	default:
		return ui.Theme.TextDim
	}
}
func (pt *ProjectTable) formatHealth(project *domain.Project) string {
	// TODO: Add Health field to Project model
	// For now, calculate based on tasks or return default

	// Option 1: Check if there are blocked tasks
	if len(project.TasksByStatus(domain.StatusBlocked)) > 0 {
		return "⚠" // Warning
	}

	// Option 2: Check if project has tasks and all are done
	if len(project.Tasks) > 0 && project.CalculateProgress() == 100 {
		return "✓" // Healthy
	}

	// Option 3: Check if progress is low
	if project.CalculateProgress() < 30 && len(project.Tasks) > 0 {
		return "✕" // Critical
	}

	// Default
	if len(project.Tasks) == 0 {
		return "–" // None
	}

	return "✓" // Healthy
}
func (pt *ProjectTable) getHealthColor(project *domain.Project) tcell.Color {
	health := pt.formatHealth(project)
	switch health {
	case "✓":
		return tcell.ColorGreen
	case "⚠":
		return tcell.ColorYellow
	case "✕":
		return tcell.ColorRed
	default:
		return tcell.ColorGray
	}
}
func (pt *ProjectTable) getProjectAtRow(row int) *domain.Project {
	if row <= 0 || row > len(pt.projects) {
		return nil
	}
	return pt.projects[row-1] // -1 because row 0 is header
}
func (pt *ProjectTable) handleKeyEvent(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'j':
		// Move down
		row, col := pt.GetSelection()
		if row < pt.GetRowCount()-1 {
			pt.Select(row+1, col)
		}
		return nil
	case 'k':
		// Move up
		row, col := pt.GetSelection()
		if row > 1 { // Don't go to header
			pt.Select(row-1, col)
		}
		return nil
	}
	return event
}

// Utility function for relative time formatting
func formatRelativeTime(t time.Time) string {
	if t.IsZero() {
		return "–"
	}

	now := time.Now()
	diff := now.Sub(t)

	switch {
	case diff < time.Minute:
		return "just now"
	case diff < time.Hour:
		minutes := int(diff.Minutes())
		return fmt.Sprintf("%dm ago", minutes)
	case diff < 24*time.Hour:
		hours := int(diff.Hours())
		return fmt.Sprintf("%dh ago", hours)
	case diff < 30*24*time.Hour:
		days := int(diff.Hours() / 24)
		return fmt.Sprintf("%dd ago", days)
	case diff < 365*24*time.Hour:
		months := int(diff.Hours() / 24 / 30)
		return fmt.Sprintf("%dmo ago", months)
	default:
		years := int(diff.Hours() / 24 / 365)
		return fmt.Sprintf("%dy ago", years)
	}
}

// )
// Sort the projects by status
func (pt *ProjectTable) sortByStatus() {
	var statusOrder = map[domain.ProjectStatus]int{
		domain.StatusActive:   0,
		domain.StatusInactive: 1,
		domain.StatusArchived: 2,
	}
	sort.Slice(pt.projects, func(i, j int) bool {
		return statusOrder[pt.projects[i].Status] <
			statusOrder[pt.projects[j].Status]
	})
}
