package main

import (
	"github.com/evertras/bubble-table/table"
)

// Stores the current context of the application
// Each successive context represents a deeper level of detail
// for the user to interact with. These can also be interpreted as
// bi-directional states for the model.
type Context int

const (
	MinContext Context = iota
	Workspaces
	Workflows
	Tasks
	MaxContext
)

type model struct {
	// The table.Model interface is used to represent the data
	// in a tabular format. This is used to display the data to
	// the user.
	table table.Model

	// The context represents the current state of the application.
	context Context

	// The model table is updated based on the selected row. Store the stack
	// of selected rows to allow the user to navigate back to the previous
	// context.
	rowQueryStack []table.Row
}

func (m *model) nextContext() {
	m.context += 1
	m.rowQueryStack = append(m.rowQueryStack, m.table.HighlightedRow())
}
func (m *model) prevContext() table.Row {
	m.context -= 1

	// Pop the last row from the stack and return it to get back to the
	// previous context.
	row := m.rowQueryStack[len(m.rowQueryStack)-1]
	m.rowQueryStack = m.rowQueryStack[:len(m.rowQueryStack)-1]
	return row
}
func (m *model) updateTable(highlightedRow table.Row) {
	switch m.context {
	case Workspaces:
		m.table = createWorkspacesTable(getWorkspaces())
	case Workflows:
		m.table = createWorkflowsTable(getWorkflows(highlightedRow.Data["workspaceId"].(int)))
	case Tasks:
		m.table = createTasksTable(
			getWorkflowTasks(highlightedRow.Data["workspaceId"].(int), highlightedRow.Data["workflowId"].(string)),
		)
	}
}
