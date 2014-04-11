package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type Project struct {
	Name        string
	ID          int64
	Workspace   string
	WorkspaceID int64
}

type Projects []Project

type NameIDResponse struct {
	Data []struct {
		ID   int64  `id`
		Name string `name`
	} `data`
}

func loadProjects() {
	//TODO: fetch
	projects := fetchProjects()
	file, err := os.OpenFile("/tmp/asana.js", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
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
