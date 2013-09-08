package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const BASE_URL = "https://app.asana.com/api/1.0"

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

func urlFor(args ...string) string {
	return strings.Join([]string{BASE_URL, strings.Join(args, "/")}, "/")
}
