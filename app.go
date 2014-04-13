package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	loadConfig()
	//-r should refresh the metadata
	parseOptions()
}

func list() {
	loadProjects()
	match := strings.Join(os.Args[2:], " ")
	fmt.Println(match)
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

var commands = map[string]func(){
	"list": list,
}

func parseOptions() {
	if len(os.Args) < 2 {
		invalidArgsHandler()
	}

	command, ok := commands[os.Args[1]]
	if !ok {
		invalidArgsHandler()
	}
	command()
}

func invalidArgsHandler() {
	fmt.Println("Invalid args")
	fmt.Println("Usage: asana list <project-name>")
	os.Exit(1)
}
