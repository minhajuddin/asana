package main

import (
	"encoding/json"
	"log"
	"net/http"
	//"os/user"
	//"path"
	"strings"
)

const BASE_URL = "https://app.asana.com/api/1.0"

func get(payload interface{}, urlTokens ...string) error {
	url := urlFor(urlTokens...)
	request, _ := http.NewRequest("GET", url, nil)
	request.SetBasicAuth(config.Key, "")
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("HTTP error", err)
	}
	if r.StatusCode != 200 {
		log.Println("ERROR", r)
	}
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&payload)
	if err != nil {
		log.Println("error in decoding", err)
	}
	return err
}

func urlFor(args ...string) string {
	return strings.Join([]string{BASE_URL, strings.Join(args, "/")}, "/")
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
