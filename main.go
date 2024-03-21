package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var url = os.Getenv("TOWER_API_ENDPOINT")
var token = os.Getenv("TOWER_ACCESS_TOKEN")

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
                        if m.context > MinContext + 1 {
                            prevRow := m.prevContext()
                            m.updateTable(prevRow)
                        }
			return m, nil
		case "enter":
                        if m.context < MaxContext - 1 {
                            m.nextContext()
                            m.updateTable(m.table.SelectedRow())
                        }
			return m, nil
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func (m model) Init() tea.Cmd { return nil }

func main() {
	userWorkspaces := getWorkspaces()
	workspacesTable := createWorkspacesTable(userWorkspaces)
	if _, err := tea.NewProgram(model{workspacesTable, Workspaces, []table.Row{}}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
