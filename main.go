package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

var url = os.Getenv("TOWER_API_ENDPOINT")
var token = os.Getenv("TOWER_ACCESS_TOKEN")

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.context > MinContext+1 && !m.table.GetIsFilterInputFocused() {
				prevRow := m.prevContext()
				m.updateTable(prevRow)
                return m, nil
			}
		case "enter":
			if m.context < MaxContext-1 && !m.table.GetIsFilterInputFocused() {
				m.nextContext()
				m.updateTable(m.table.HighlightedRow())
                return m, nil
			}
		case "ctrl+c":
			return m, tea.Quit
		}
    case tea.WindowSizeMsg:
        m.windowWidth = msg.Width
        m.windowHeight = msg.Height
        m.table = m.table.WithTargetWidth(int(math.Floor(float64(msg.Width) * 0.9)))
	}
	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	body := strings.Builder{}
	body.WriteString(m.table.View())
	return body.String()
}

func (m model) Init() tea.Cmd { return nil }

func main() {
	userWorkspaces := getWorkspaces()
	workspacesTable := createWorkspacesTable(userWorkspaces, 0)
	if _, err := tea.NewProgram(
		model{
			workspacesTable,
			Workspaces,
			[]table.Row{},
            0,
            0,
		},
	).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
