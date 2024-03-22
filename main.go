package main

import (
	"fmt"
	"os"
	"strings"

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
		case "/":
			if m.table.Focused() && len(m.inputs) > 0 {
				m.table.Blur()
				m.inputs[0].Focus()
			}
			return m, nil
		case "tab":
			m.cycleInputs()
		case "esc":
			if !m.table.Focused() && len(m.inputs) > 0 {
				m.table.Focus()
				blurInputs(m.inputs, false)
				return m, nil
			}
			if m.context > MinContext+1 {
				prevRow := m.prevContext()
				m.updateTable(prevRow)
			}
			return m, nil
		case "enter":
			if !m.table.Focused() {
				return m, nil
			}
			blurInputs(m.inputs, true)
			if m.context < MaxContext-1 {
				m.nextContext()
				m.updateTable(m.table.SelectedRow())
			}
			return m, nil
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	// If the focus is on the text input, update the text input.
	for i := range m.inputs {
		if m.inputs[i].Focused() {
			m.inputs[i], cmd = m.inputs[i].Update(msg)
			return m, tea.Batch(cmd)
		}
	}

	// Otherwise, update the table.
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m model) View() string {
	inputStrings := make([]string, len(m.inputs))
	for i := range m.inputs {
		builder := strings.Builder{}
		builder.WriteString(m.inputs[i].View())
		builder.WriteString("\t")
		inputStrings[i] = builder.String()
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		baseStyle.Render(m.table.View()),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			inputStrings...,
		),
	) + "\n"
}

func (m model) Init() tea.Cmd { return nil }

func main() {
	userWorkspaces := getWorkspaces()
	workspacesTable := createWorkspacesTable(userWorkspaces)
	workspacesFilter := createWorkspacesFilter()
	if _, err := tea.NewProgram(model{workspacesTable, workspacesFilter, table.Model{}, Workspaces, []table.Row{}}).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
