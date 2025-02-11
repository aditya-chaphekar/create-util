package models

import tea "github.com/charmbracelet/bubbletea"

type RestartModelDirCheck struct {
	dir    string
	err    bool
	errMsg string
}

func (m RestartModelDirCheck) Init() tea.Cmd {
	return nil
}

func (m RestartModelDirCheck) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m RestartModelDirCheck) View() string {
	return "RestartModelDirCheck"
}
