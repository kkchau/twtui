package main

type userInfoResponse struct {
	User struct {
		Id int `json:"id"`
	} `json:"user"`
}

type workspaceResponse struct {
	OrgsAndWorkspaces []struct {
		OrgId             int    `json:"orgId"`
		OrgName           string `json:"orgName"`
		WorkspaceId       int    `json:"workspaceId"`
		WorkspaceName     string `json:"workspaceName"`
		WorkspaceFullName string `json:"workspaceFullName"`
	} `json:"orgsAndWorkspaces"`
}

type workflowsResponse struct {
	Workflows []struct {
		OrgName     string `json:"orgName"`
		WorkspaceId int    `json:"workspaceId"`
		Workflow    struct {
			Id       string `json:"id"`
			RunName  string `json:"runName"`
			Status   string `json:"status"`
			Submit   string `json:"submit"`
			Start    string `json:"start"`
			Complete string `json:"complete"`
			Stats    struct {
				Cached    int `json:"cachedCount"`
				Succeeded int `json:"succeedCount"`
				Failed    int `json:"failedCount"`
			} `json:"stats"`
		} `json:"workflow"`
	} `json:"workflows"`
}

type workflowResponse struct {
	Workflow struct {
		Id       string `json:"id"`
		RunName  string `json:"runName"`
		Status   string `json:"status"`
		Duration string `json:"duration"`
	} `json:"workflow"`
}

type tasksResponse struct {
	Tasks []struct {
		Task struct {
			Id       string `json:"id"`
			Name     string `json:"name"`
			Status   string `json:"status"`
			Attempt  int    `json:"attempt"`
			Duration string `json:"duration"`
			Tag      string `json:"tag"`
		} `json:"task"`
	}
}
