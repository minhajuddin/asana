package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	loadConfig()
	//-r should refresh the metadata
	loadProjects()
	//TODO handle invalid args
	if len(os.Args) < 2 {
		fmt.Println("Invalid number of args")
		os.Exit(1)
	}
	match := strings.Join(os.Args[1:], " ")
	project := projects.find(match)
	if project == nil {
		fmt.Println("Project not found")
		os.Exit(2)
	}
	fmt.Printf("Showing tasks for %v/%v:\n", project.Workspace, project.Name)
	for i, t := range project.tasks() {
		fmt.Printf("%v. %v\n", i+1, t.Name)
	}
}
