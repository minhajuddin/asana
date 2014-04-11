package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("doing it")
	loadConfig()
	//-r should refresh the metadata
	loadProjects()
	//TODO handle invalid args
	match := os.Args[1]
	//TODO handle match not found
	project := projects.find(match)
	for _, t := range project.tasks() {
		fmt.Println(t.Name)
	}
}
