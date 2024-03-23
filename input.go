package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
)

func createTextInput(fields []string) []textinput.Model {
	inputs := make([]textinput.Model, len(fields))

	var t textinput.Model
	for i, field := range fields {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32
		t.Placeholder = field
		t.Blur()
		t.PromptStyle = focusedStyle
		t.TextStyle = focusedStyle
		inputs[i] = t
	}
	return inputs
}

func blurInputs(inputs []textinput.Model, reset bool) {
	for i := range inputs {
		inputs[i].Blur()
		if reset {
			inputs[i].Reset()
		}
	}
}

func (m *model) cycleInputs() {
	focusNext := false
	for i := range m.inputs {
		if focusNext {
			m.inputs[i].Focus()
			return
		}
		if m.inputs[i].Focused() {
			m.inputs[i].Blur()
			focusNext = true
		}
	}
	if focusNext && len(m.inputs) > 0 {
		m.inputs[0].Focus()
	}
}

func createWorkspacesFilter() []textinput.Model {
	return createTextInput([]string{"orgId", "orgName", "workspaceId", "workspaceName", "workspaceFullName"})
}
func createWorkflowsFilter() []textinput.Model {
	return createTextInput([]string{"runName", "status"})
}
func createTasksFilter() []textinput.Model {
	return createTextInput([]string{"name", "status", "tag"})
}
