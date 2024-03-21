package main

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func createTable(columns []table.Column, rows []table.Row) table.Model {
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
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	t.SetStyles(s)
	return t
}

func createWorkspacesTable(workspaces workspaceResponse) table.Model {
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
		{Title: "Org ID", Width: 15},
		{Title: "Org Name", Width: 10},
		{Title: "Workspace ID", Width: 15},
		{Title: "Workspace Name", Width: 15},
		{Title: "Workspace Full Name", Width: 20},
	}, rows)
}

func createWorkflowsTable(workflows workflowsResponse) table.Model {
	rows := []table.Row{}
	for _, workflow := range workflows.Workflows {
		rows = append(rows,
			table.Row{
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
		{Title: "runName", Width: 15},
		{Title: "status", Width: 15},
                {Title: "submit", Width: 15},
                {Title: "start", Width: 15},
                {Title: "complete", Width: 15},
		{Title: "cached", Width: 15},
		{Title: "succeeded", Width: 15},
		{Title: "failed", Width: 15},
	}, rows)
}
