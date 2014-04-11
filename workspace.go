package main

import (
	"fmt"
)

type Workspace struct {
	Id       int64  `id`
	Name     string `name`
	Projects []Project
}

type WorkspacePayload struct {
	Data []Workspace `data`
}

func getWorkspaces() []Workspace {
	w := WorkspacePayload{}
	get(&w, "workspaces")
	for _, w := range w.Data {
		go func() {
			w.Projects = getProjects(w)
		}()
	}
	return w.Data
}

func listWorkspaces() {
	for i, w := range getWorkspaces() {
		fmt.Printf("%02d) %s (%v)\n", i+1, w.Name, w.Id)
	}
}
