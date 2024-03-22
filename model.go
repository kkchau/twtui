package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
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
	MaxContext // Limit the context to here until more contexts are implemented
	Runs
	Tasks
)

type model struct {
	// The table.Model interface is used to represent the data
	// in a tabular format. This is used to display the data to
	// the user.
	table table.Model

	inputs []textinput.Model

	// Represents the data in the table field filtered down to some subset of the
	// original data. This is used to allow the user to navigate through the
	// data in a more focused manner. The original data is stored in the table
	// field when the user is done with the filtered data.
	filteredTable table.Model

	// The context represents the current state of the application.
	context Context

	// The model table is updated based on the selected row. Store the stack
	// of selected rows to allow the user to navigate back to the previous
	// context.
	rowQueryStack []table.Row
}

func (m *model) nextContext() {
	m.context += 1
	m.rowQueryStack = append(m.rowQueryStack, m.table.SelectedRow())
}
func (m *model) prevContext() table.Row {
	m.context -= 1

	// Pop the last row from the stack and return it to get back to the
	// previous context.
	row := m.rowQueryStack[len(m.rowQueryStack)-1]
	m.rowQueryStack = m.rowQueryStack[:len(m.rowQueryStack)-1]
	return row
}
func (m *model) updateTable(selectedRow table.Row) {
	switch m.context {
	case Workspaces:
		m.table = createWorkspacesTable(getWorkspaces())
		m.inputs = createWorkspacesFilter()
	case Workflows:
		m.table = createWorkflowsTable(getWorkflows(selectedRow[2]))
		m.inputs = createWorkflowsFilter()
	}
}
