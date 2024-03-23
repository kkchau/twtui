package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func createTable(columns []table.Column, rows []table.Row) (t table.Model, h []string) {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t = table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	t.SetStyles(s)

	h = make([]string, len(columns))
	for i, col := range columns {
		h[i] = col.Title
	}
	return t, h
}

func createWorkspacesTable(workspaces workspaceResponse) (table.Model, []string) {
	rows := []table.Row{}
	for _, workspace := range workspaces.OrgsAndWorkspaces {
		rows = append(rows,
			table.Row{
				strconv.Itoa(workspace.OrgId),
				workspace.OrgName,
				strconv.Itoa(workspace.WorkspaceId),
				workspace.WorkspaceName,
				workspace.WorkspaceFullName,
			},
		)
	}
	return createTable([]table.Column{
		{Title: "orgId", Width: 15},
		{Title: "orgName", Width: 10},
		{Title: "workspaceId", Width: 15},
		{Title: "workspaceName", Width: 15},
		{Title: "workspaceFullName", Width: 25},
	}, rows)
}

func createWorkflowsTable(workflows workflowsResponse) (table.Model, []string) {
	rows := []table.Row{}
	for _, workflow := range workflows.Workflows {
		rows = append(
			rows,
			table.Row{
				strconv.Itoa(workflow.WorkspaceId),
				workflow.Workflow.Id,
				workflow.Workflow.RunName,
				workflow.Workflow.Status,
				workflow.Workflow.Submit,
				workflow.Workflow.Start,
				workflow.Workflow.Complete,
				strconv.Itoa(workflow.Workflow.Stats.Cached),
				strconv.Itoa(workflow.Workflow.Stats.Succeeded),
				strconv.Itoa(workflow.Workflow.Stats.Failed),
			},
		)
	}
	return createTable([]table.Column{
		{Title: "workspaceId", Width: 0},
		{Title: "id", Width: 0},
		{Title: "runName", Width: 15},
		{Title: "status", Width: 10},
		{Title: "submit", Width: 20},
		{Title: "start", Width: 20},
		{Title: "complete", Width: 20},
		{Title: "cached", Width: 10},
		{Title: "succeeded", Width: 10},
		{Title: "failed", Width: 10},
	}, rows)
}
func createTasksTable(tasks tasksResponse) (table.Model, []string) {
	rows := []table.Row{}
	for _, task := range tasks.Tasks {
		rows = append(
			rows,
			table.Row{
				task.Task.Id,
				task.Task.Name,
				task.Task.Status,
				strconv.Itoa(task.Task.Attempt),
				task.Task.Duration,
				task.Task.Tag,
			},
		)
	}
	return createTable([]table.Column{
		{Title: "id", Width: 0},
		{Title: "name", Width: 15},
		{Title: "status", Width: 10},
		{Title: "attempt", Width: 10},
		{Title: "duration", Width: 10},
		{Title: "tag", Width: 10},
	}, rows)
}

func queryRows(inputRows []table.Row, tableColumnIndex int, query string) (rows []table.Row) {
	for _, row := range inputRows {
		if strings.Contains(row[tableColumnIndex], query) {
			rows = append(rows, row)
		}
	}
	return rows
}
