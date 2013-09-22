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

func getProjects(cache bool) []Project {
	p := ProjectPayload{}
	get(&p, "projects")
	cacheProjects(p.Data)
	return p.Data
}

func listProjects() {
	for i, project := range getProjects(true) {
		fmt.Printf("%02d) %s (%v)\n", i+1, project.Name, project.Id)
	}
}

func cacheProjects(projects []Project) {
}
