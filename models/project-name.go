package models

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	validStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")) // Green
	invalidStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")) // Red
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")) // Orange (Warning)
	nameStyle    = lipgloss.NewStyle().Bold(true)                            // Bold name
	mutedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#808080"))
)

type ProjectName struct {
	name        string
	err         bool
	errMsg      string
	firstRender bool
}

func (m ProjectName) Init() tea.Cmd {
	return nil
}

func (m ProjectName) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// ✅ Return tea.Batch() to execute both quitting and starting the next model
			return m, tea.Batch(tea.Quit, func() tea.Msg {
				return StartProjectTypeSelectionMsg{}
			})
		case "backspace":
			if len(m.name) > 0 {
				m.name = m.name[:len(m.name)-1]
			} else {
				m.name = "project-name"
			}
		case "ctrl+c", "esc":
			os.Exit(0)
		default:
			if len(msg.String()) == 1 {
				m.name += msg.String()
			}
		}
		// Revalidate after every name change
		m.err, m.errMsg = validateProjectName(m.name)
	}
	return m, nil
}

var icon = mutedStyle.Render("?")
var count = 0

func (m ProjectName) View() string {
	if count <= 1 {
		m.firstRender = true
	} else {
		m.firstRender = false
	}
	if m.err {
		icon = invalidStyle.Render("✖")
	} else if !m.err && !m.firstRender {
		icon = validStyle.Render("✔")
	}
	displayText := nameStyle.Render(m.name)
	if m.firstRender || m.name == "" || m.name == "project-name" {
		displayText = mutedStyle.Render("project-name")
	}
	s := fmt.Sprintf("%s Enter your project name: %s\n", icon, displayText)

	if m.err {
		icon = errorStyle.Render("!")
		s += fmt.Sprintf("\n%s %s\n", icon, m.errMsg)
	}

	s += "\nPress Enter to confirm, Backspace to delete, ESC to quit."

	count += 1
	return s
}

// Function to get project name interactively
func GetProjectName() string {
	p := tea.NewProgram(ProjectName{
		firstRender: true,
	})
	m, err := p.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return m.(ProjectName).name
}

// Validate project name in real-time
func validateProjectName(name string) (bool, string) {
	name = strings.TrimSpace(name)
	if len(name) < 3 {
		return true, "Project name must be at least 3 characters long"
	}
	if len(name) > 50 {
		return true, "Project name must not exceed 50 characters"
	}
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9-_]+$`, name); !matched {
		return true, "Only letters, numbers, hyphens, and underscores are allowed"
	}
	return false, ""
}
