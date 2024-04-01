package main

import (
	"fmt"

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

	// The pageOffset is used to keep track of the current page of data where applicable
	pageOffset int

	// The loadMore is used to indicate if the user can load more data
	// than what is currently visible in the table
	loadMore string

	// The model table is updated based on the selected row. Store the stack
	// of selected rows to allow the user to navigate back to the previous
	// context.
	rowQueryStack []table.Row

	windowWidth  int
	windowHeight int
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
func (m *model) updateTable(highlightedRow table.Row, backstep bool) {
	var response PaginatedResponse
	m.loadMore = ""
	switch m.context {
	case Workspaces:
		response = getWorkspaces()
		m.table = createWorkspacesTable(response.(workspaceResponse), m.windowWidth)
	case Workflows:
		response = getWorkflows(highlightedRow.Data["workspaceId"].(int))
		m.table = createWorkflowsTable(response.(workflowsResponse), m.windowWidth)
	case Tasks:
		if backstep {
			if m.pageOffset == 0 {
				return
			}

			m.pageOffset -= len(m.table.GetVisibleRows()) * 2
			if m.pageOffset < 0 {
				m.pageOffset = 0
			}
		}
		response = getWorkflowTasks(
			highlightedRow.Data["workspaceId"].(int),
			highlightedRow.Data["workflowId"].(string),
			m.pageOffset,
		)
		m.table = createTasksTable(response.(tasksResponse), m.windowWidth)
		m.loadMore = fmt.Sprintf(
			"Loaded %d - %d.",
			m.pageOffset+1,
			m.pageOffset+response.GetPageSize(),
		)
	}
	if m.loadMore != "" {
		m.pageOffset += response.GetPageSize()
	} else {
		m.pageOffset = 0
	}
}
