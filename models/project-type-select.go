package models

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles for UI
var (
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#B4A7D6")).Bold(true).Underline(true)
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	pointer       = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF")).Render("❯")
)

// Menu items
var projectTypes = []string{"Frontend", "Backend"}

type StartProjectTypeSelectionMsg struct{}

type menuModel struct {
	selected bool
	cursor   int
}

// Init method
func (m menuModel) Init() tea.Cmd {
	return nil
}

// Update handles key presses
func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StartProjectTypeSelectionMsg:
		return m, nil // ✅ Start model immediately without a newlin
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(projectTypes)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = true
			return m, tea.Quit
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View method renders the UI
func (m menuModel) View() string {
	if m.selected {
		return "Selected project type: " + projectTypes[m.cursor]
	}
	s := "Select the project type:\n"
	for i, option := range projectTypes {
		if i == m.cursor {
			s += fmt.Sprintf("%s %s\n", pointer, selectedStyle.Render(option))
		} else {
			s += fmt.Sprintf("  %s\n", normalStyle.Render(option))
		}
	}
	s += "Use ↑/↓ to navigate, Enter to select."
	return s
}

// GetProjectType runs the menu and returns the selected project type
func GetProjectType() (string, error) {
	p := tea.NewProgram(menuModel{})
	m, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("error running menu: %w", err)
	}
	fmt.Print("\r")
	return projectTypes[m.(menuModel).cursor], nil
}

