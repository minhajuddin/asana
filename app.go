package main

import (
	"encoding/json"
	_ "fmt"
	"log"
	"net/http"
	"strings"
)

const API_KEY = "xxx"
const PROJECT_ID = "xx"

//const TASKS_URL =

type Task struct {
	Id   int64  `id`
	Name string `name`
}

type Payload struct {
	Data []Task `data`
}

func get(payload interface{}, urlTokens ...string) error {
	url := urlFor(urlTokens...)
	request, _ := http.NewRequest("GET", url, nil)
	request.SetBasicAuth(config.ApiKey, "")
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("HTTP error", err)
	}
	dec := json.NewDecoder(r.Body)
	return dec.Decode(&payload)
}

//func getTasks() []Task {
//request, _ := http.NewRequest("GET", urlFor("projects"), nil)
//request.SetBasicAuth(config.ApiKey, "")
//r, _ := http.DefaultClient.Do(request)
//dec := json.NewDecoder(r.Body)
//p := Payload{}
//dec.Decode(&p)
//return p.Data
//}

//func listTasks() {
//tasks := getTasks()
//for i, task := range tasks {
//fmt.Printf("%02d) %s\n", i+1, task.Name)
//}
//}

//func listTasks() {
//tasks := getTasks()
//for i, task := range tasks {
//fmt.Printf("%02d) %s\n", i+1, task.Name)
//}

const BASE_URL = "https://app.asana.com/api/1.0"

func urlFor(args ...string) string {
	return strings.Join([]string{BASE_URL, strings.Join(args, "/")}, "/")
}

func main() {
	readConfig()
	listProjects()
}
