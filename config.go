package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var config struct {
	Key string `json:"key"`
}

func loadConfig() {
	homeDir := os.Getenv("HOME")
	bytes, err := ioutil.ReadFile(path.Join(homeDir, ".asana.json"))

	if err != nil {
		log.Fatalln(err, "config file not found")
	}
	json.Unmarshal(bytes, &config)
}
