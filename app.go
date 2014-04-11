package main

import (
	"fmt"
)

func main() {
	loadConfig()
	fmt.Println(config.Key)
	//-r should refresh the metadata
	//i.e. the workspaces and project list
	//and it should run in the background after showing the list
	//of tasks if lastmod of the cache file is > 1.day
	//getMeta()
	//listProjects()
	listWorkspaces()
}

//func initialize() {
//cachePath := relativeFromHome(".asana")
//os.Mkdir(cachePath, os.)
//}
