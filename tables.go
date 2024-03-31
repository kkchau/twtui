package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

func createTable(columns []table.Column, rowDatas []table.RowData, width int) (t table.Model) {
	rows := make([]table.Row, len(rowDatas))
	for i, rowData := range rowDatas {
		rows[i] = table.NewRow(rowData)
	}
	t = table.New(columns).
		WithRows(rows).
		Filtered(true).
		Focused(true).
		WithPageSize(10).
		WithTargetWidth(width).
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

func createWorkspacesTable(workspaces workspaceResponse, width int) table.Model {
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
		width,
	)
}

func createWorkflowsTable(workflows workflowsResponse, width int) table.Model {
	rowDatas := make([]table.RowData, len(workflows.Workflows))
	for i, workflow := range workflows.Workflows {
		rowDatas[i] = table.RowData{
			"workflowId":  workflow.Workflow.Id,
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
			table.NewFlexColumn("runName", "Run Name", 1).WithFiltered(true),
			table.NewFlexColumn("status", "Status", 1).WithFiltered(true),
			table.NewFlexColumn("submit", "Submit", 1),
			table.NewFlexColumn("start", "Start", 1),
			table.NewFlexColumn("complete", "Complete", 1),
			table.NewFlexColumn("cached", "Cached", 1),
			table.NewFlexColumn("succeeded", "Succeeded", 1),
			table.NewFlexColumn("failed", "Failed", 1),
		},
		rowDatas,
		width,
	)
}

func createTasksTable(tasks tasksResponse, width int) table.Model {
	rowDatas := make([]table.RowData, len(tasks.Tasks))
	for i, task := range tasks.Tasks {
		rowDatas[i] = table.RowData{
			"workspaceId": tasks.WorkspaceId,
			"workflowId":  tasks.WorkflowId,
			"id":          task.Task.Id,
			"name":        task.Task.Name,
			"status":      task.Task.Status,
			"attempt":     task.Task.Attempt,
			"duration":    task.Task.Duration,
			"tag":         task.Task.Tag,
			"total":       tasks.Total,
		}
	}
	return createTable(
		[]table.Column{
			table.NewFlexColumn("name", "Name", 10).WithFiltered(true),
			table.NewFlexColumn("status", "Status", 2).WithFiltered(true),
			table.NewFlexColumn("attempt", "Attempt", 1),
			table.NewFlexColumn("duration", "Duration", 2),
			table.NewFlexColumn("tag", "Tag", 5),
		},
		rowDatas,
		width,
	)
}
