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

func list(args []string) {
	loadProjects()
	match := strings.Join(args, " ")
	project := projects.find(match)
	fmt.Printf("Showing tasks for %v/%v:\n", project.Workspace, project.Name)
	for i, t := range project.tasks() {
		fmt.Printf("%v. %v\n", i+1, t.Name)
	}
}

func add(args []string) {
	if len(args) < 2 {
		fmt.Println("add needs 2 arguments")
		os.Exit(2)
	}
	loadProjects()
	project := projects.find(args[0])
	project.addTask(strings.Join(args[1:], " "))
}

var commands = map[string]func([]string){
	"list": list,
	"add":  add,
}

func parseOptions() {
	if len(os.Args) < 2 {
		invalidArgsHandler()
	}

	command, ok := commands[os.Args[1]]
	if !ok {
		invalidArgsHandler()
	}
	command(os.Args[2:])
}

func invalidArgsHandler() {
	fmt.Println("Invalid args")
	fmt.Println("Usage: asana list <project-name>")
	os.Exit(1)
}
