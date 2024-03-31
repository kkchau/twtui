package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

func makeGetRequest(endpoint string) (*http.Request, error) {
	bearer := "Bearer " + token
	req, err := http.NewRequest("GET", url+endpoint, nil)
	req.Header.Add("Authorization", bearer)
	return req, err
}

func getUserInfo() userInfoResponse {
	req, _ := makeGetRequest("/user-info")

	client := &http.Client{Timeout: 10 * time.Second}
	res, _ := client.Do(req)

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	userInfo := userInfoResponse{}
	json.Unmarshal(body, &userInfo)
	return userInfo
}

func getWorkspacesFromUserInfo(userId int) workspaceResponse {
	req, _ := makeGetRequest("/user/" + strconv.Itoa(userId) + "/workspaces")

	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	workspaces := workspaceResponse{}
	json.Unmarshal(body, &workspaces)
	return workspaces
}

func getWorkspaces() workspaceResponse {
	return getWorkspacesFromUserInfo(getUserInfo().User.Id)
}

func getWorkflows(workspaceId int) workflowsResponse {
	req, _ := makeGetRequest("/workflow")
	query := req.URL.Query()
	query.Add("workspaceId", strconv.Itoa(workspaceId))
	req.URL.RawQuery = query.Encode()

	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	workflows := workflowsResponse{}
	json.Unmarshal(body, &workflows)
	return workflows
}

func getWorkflowTasks(workspaceId int, workflowId string, pageOffset int) tasksResponse {
	req, _ := makeGetRequest("/workflow/" + workflowId + "/tasks")
	query := req.URL.Query()
	query.Add("workspaceId", strconv.Itoa(workspaceId))
	query.Add("offset", strconv.Itoa(pageOffset))
	req.URL.RawQuery = query.Encode()
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	tasks := tasksResponse{}
	json.Unmarshal(body, &tasks)

	// Store the workspaceId and workflowId in the tasks struct
	// for subsequent requests
	tasks.WorkspaceId = workspaceId
	tasks.WorkflowId = workflowId
	return tasks
}
