package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const API_KEY = "xxx"
const PROJECT_ID = "xx"
const TASKS_URL = "https://app.asana.com/api/1.0/projects/xxx/tasks"

type Task struct {
	Id   int64  `id`
	Name string `name`
}

type Payload struct {
	Data []Task `data`
}

func main() {
	request, _ := http.NewRequest("GET", TASKS_URL, nil)
	request.SetBasicAuth(API_KEY, "")
	r, _ := http.DefaultClient.Do(request)
	dec := json.NewDecoder(r.Body)
	p := Payload{}
	dec.Decode(&p)
	for i, task := range p.Data {
		fmt.Printf("%02d) %s\n", i+1, task.Name)
	}
}
