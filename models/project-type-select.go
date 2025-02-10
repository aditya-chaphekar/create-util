package models

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles for UI
var (
    selectedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#B4A7D6")).Bold(true).Underline(true) // Green & Bold
    normalStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))            // White
    pointer         = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF")).Render("❯") // Yellow Pointer
)

// Menu items
var projectTypes = []string{"Frontend", "Backend"}

type menuModel struct {
    cursor int
}

// Init method
func (m menuModel) Init() tea.Cmd {
    return nil
}

// Update handles key presses
func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "up", "k": // Move up
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j": // Move down
            if m.cursor < len(projectTypes)-1 {
                m.cursor++
            }
        case "enter": // Select option
            return m, tea.Quit
        case "ctrl+c", "esc": // Exit program
            os.Exit(0)
        }
    }
    return m, nil
}

// View method renders the UI
func (m menuModel) View() string {
    s := "Select the project type:\n\n"
    for i, option := range projectTypes {
        if i == m.cursor {
            s += fmt.Sprintf("%s \t%s\n", pointer, selectedStyle.Render(option))
        } else {
            s += fmt.Sprintf("   \t%s\n", normalStyle.Render(option))
        }
    }
    s += "\nUse ↑/↓ to navigate, Enter to select."
    return s
}

// Function to get the selected project type
func GetProjectType() string {
    p := tea.NewProgram(menuModel{})
    m, err := p.Run()
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    return projectTypes[m.(menuModel).cursor]
}
