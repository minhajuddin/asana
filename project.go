package main

import (
	"fmt"
)

type Project struct {
	Id   int64  `id`
	Name string `name`
}

type ProjectPayload struct {
	Data []Project `data`
}

func getProjects() []Project {
	p := ProjectPayload{}
	get(&p, "projects")
	return p.Data
}

func listProjects() {
	for i, project := range getProjects() {
		fmt.Printf("%02d) %s\n", i+1, project.Name)
	}
}
