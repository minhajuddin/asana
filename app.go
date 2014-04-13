package main

import (
	"fmt"
	"os"
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
	match := os.Args[1]
	//TODO handle match not found
	project := projects.find(match)
	fmt.Println("Showing tasks for ", project.Name)
	for _, t := range project.tasks() {
		fmt.Println(t.Name)
	}
}
