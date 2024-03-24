package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

func createTable(columns []table.Column, rowDatas []table.RowData) (t table.Model) {
	rows := make([]table.Row, len(rowDatas))
	for i, rowData := range rowDatas {
		rows[i] = table.NewRow(rowData)
	}
	t = table.New(columns).
		WithRows(rows).
		Filtered(true).
		Focused(true).
		WithPageSize(10).
		WithTargetWidth(100).
		WithBaseStyle(
			lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("#a38")).
				Foreground(lipgloss.Color("#a7a")).
				Align(lipgloss.Left),
		).
		HeaderStyle(
			lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true),
		)
	return
}

func createWorkspacesTable(workspaces workspaceResponse) table.Model {
	rowDatas := make([]table.RowData, len(workspaces.OrgsAndWorkspaces))
	for i, workspace := range workspaces.OrgsAndWorkspaces {
		rowDatas[i] = table.RowData{
			"orgId":             workspace.OrgId,
			"orgName":           workspace.OrgName,
			"workspaceId":       workspace.WorkspaceId,
			"workspaceName":     workspace.WorkspaceName,
			"workspaceFullName": workspace.WorkspaceFullName,
		}
	}
	return createTable(
		[]table.Column{
			table.NewFlexColumn("orgName", "Org Name", 1).WithFiltered(true),
			table.NewFlexColumn("workspaceName", "Workspace Name", 1),
			table.NewFlexColumn("workspaceFullName", "Workspace Full Name", 1).WithFiltered(true),
		},
		rowDatas,
	)
}

func createWorkflowsTable(workflows workflowsResponse) table.Model {
	rowDatas := make([]table.RowData, len(workflows.Workflows))
	for i, workflow := range workflows.Workflows {
		rowDatas[i] = table.RowData{
			"workspaceId": workflow.WorkspaceId,
			"id":          workflow.Workflow.Id,
			"runName":     workflow.Workflow.RunName,
			"status":      workflow.Workflow.Status,
			"submit":      workflow.Workflow.Submit,
			"start":       workflow.Workflow.Start,
			"complete":    workflow.Workflow.Complete,
			"cached":      workflow.Workflow.Stats.Cached,
			"succeeded":   workflow.Workflow.Stats.Succeeded,
			"failed":      workflow.Workflow.Stats.Failed,
		}
	}
	return createTable(
		[]table.Column{
			table.NewColumn("runName", "Run Name", 15).WithFiltered(true),
			table.NewColumn("status", "Status", 10).WithFiltered(true),
			table.NewColumn("submit", "Submit", 20),
			table.NewColumn("start", "Start", 20),
			table.NewColumn("complete", "Complete", 20),
			table.NewColumn("cached", "Cached", 10),
			table.NewColumn("succeeded", "Succeeded", 10),
			table.NewColumn("failed", "Failed", 10),
		},
		rowDatas,
	)
}
