package main

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"path"
)

type Config struct {
	ApiKey string `apiKey`
}

var config Config

func readConfig() {
	u, err := user.Current()
	if err != nil {
		log.Println("Unable to get user info", err)
	}
	file, err := os.Open(path.Join(u.HomeDir, ".asana-config.json"))
	if err != nil {
		log.Println("Unable to open config file", err)
	}
	dec := json.NewDecoder(file)
	config = Config{}
	dec.Decode(&config)
}
