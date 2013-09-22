package main

import (
	"os"
)

func main() {
	readConfig()
	listProjects()
}

func initialize() {
	cachePath := relativeFromHome(".asana")
	//os.Mkdir(cachePath, os.)
}
