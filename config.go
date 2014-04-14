package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

var config struct {
	Key string `json:"key"`
}

var ASANA_DIR = ""

func loadConfig() {
	ASANA_DIR = path.Join(os.Getenv("HOME"), ".asana")
	bytes, err := ioutil.ReadFile(path.Join(ASANA_DIR, "config.json"))

	handleError(err)
	json.Unmarshal(bytes, &config)
}
