package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ApiKey string `apiKey`
}

var config Config

func readConfig() {
	file, err := os.Open(relativeFromHome(".asana-config.json"))
	if err != nil {
		log.Panicln("Unable to open config file", err)
	}
	dec := json.NewDecoder(file)
	config = Config{}
	dec.Decode(&config)
}
