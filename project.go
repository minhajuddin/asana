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
	get("projects")
	return p.Data
}

func listProjects() {
	for _, project := range getProjects() {
		fmt.Println(project)
	}
}
