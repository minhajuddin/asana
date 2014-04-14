package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type Project struct {
	Name        string
	ID          int64
	Workspace   string
	WorkspaceID int64
}

type Task struct {
	ID   int64
	Name string
}

func (p *Project) addTask(name string) {
	form := struct {
		Name      string  `json:"name"`
		Workspace int64   `json:"workspace"`
		Projects  []int64 `json:"projects"`
	}{name, p.WorkspaceID, []int64{p.ID}}
	post(form, "tasks")
}

func (p *Project) tasks() []Task {
	tasks := make([]Task, 0, 100)
	tasksJson := NameIDResponse{}
	get(&tasksJson, "projects", strconv.FormatInt(p.ID, 10), "tasks")
	for _, tj := range tasksJson.Data {
		tasks = append(tasks, Task{ID: tj.ID, Name: tj.Name})
	}
	return tasks
}

type Projects []Project

func (projects Projects) find(match string) *Project {
	for _, p := range projects {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(match)) {
			return &p
		}
	}
	fmt.Printf("Project '%v' not found\n", match)
	os.Exit(2)
	return nil
}

type NameIDResponse struct {
	Data []struct {
		ID   int64  `id`
		Name string `name`
	} `data`
}

var projects Projects

func loadProjects() {
	//TODO: fetch
	cachePath := path.Join(ASANA_DIR, "projects.json")
	_, err := os.Stat(cachePath)
	//if cached load from file
	if err == nil {
		projects = make(Projects, 0, 100)
		file, err := os.Open(cachePath)
		handleError(err)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&projects)
		handleError(err)
		return
	}
	projects = fetchProjects()
	file, err := os.OpenFile(cachePath, os.O_RDWR|os.O_CREATE, 0600)
	handleError(err)
	encoder := json.NewEncoder(file)
	encoder.Encode(projects)
}

func fetchProjects() Projects {
	projects := make(Projects, 0, 100)
	workspaces := NameIDResponse{}
	get(&workspaces, "workspaces")
	for _, workspace := range workspaces.Data {
		pr := NameIDResponse{}
		get(&pr, "workspaces", strconv.FormatInt(workspace.ID, 10), "projects")
		for _, project := range pr.Data {
			p := Project{
				ID:          project.ID,
				Name:        project.Name,
				Workspace:   workspace.Name,
				WorkspaceID: workspace.ID,
			}
			projects = append(projects, p)
		}
	}
	return projects
}
