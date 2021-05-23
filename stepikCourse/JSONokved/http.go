package main

import (
	"fmt"
	"io"
	"net/http"
)

func getJSON(str string) []byte {
	resp, err := http.Get(str)
	if err != nil {
		fmt.Println("Error is occured in http repsones")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read body response")
	}
	return body
}
