package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

//const BASE_URL = "http://localhost:4000/api/1.0"

const BASE_URL = "https://app.asana.com/api/1.0"

type Form struct {
	Data interface{} `json:"data"`
}

func post(form interface{}, urlTokens ...string) error {
	body, err := json.Marshal(&Form{Data: form})
	url := urlFor(urlTokens...)
	request, _ := http.NewRequest("POST", url, strings.NewReader(string(body)))
	request.Header.Add("Content-Type", "application/json")
	request.SetBasicAuth(config.Key, "")
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("HTTP error", err)
	}
	if r.StatusCode != 201 {
		fmt.Println("ERROR", r)
		bytes, _ := httputil.DumpResponse(r, true)
		fmt.Println(string(bytes))
	}
	return err
}

//order of args isn't natural
func get(payload interface{}, urlTokens ...string) error {
	url := urlFor(urlTokens...)
	request, _ := http.NewRequest("GET", url, nil)
	request.SetBasicAuth(config.Key, "")
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("HTTP error", err)
		return err
	}
	if r.StatusCode != 200 {
		fmt.Println("ERROR", r)
		return errors.New("Invalid return code")
	}
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&payload)
	if err != nil {
		fmt.Println("error in decoding", err)
	}
	return err
}

func urlFor(args ...string) string {
	return strings.Join([]string{BASE_URL, strings.Join(args, "/")}, "/")
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
