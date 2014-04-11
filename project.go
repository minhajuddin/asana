package main

import (
	"fmt"
)

type Project struct {
	Id        int64  `id`
	Name      string `name`
	Workspace string `workspace`
}

type ProjectPayload struct {
	Data []Project `data`
}

func getProjects(workspace Workspace) []Project {
	p := ProjectPayload{}
	get(&p, fmt.Sprintf("workspaces/%v/projects", workspace.Id))
	return p.Data
}
