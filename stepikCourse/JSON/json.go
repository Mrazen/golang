package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Univer struct {
	ID       int
	Number   string
	Year     int
	Students []map[string]interface{}
}
type Average struct {
	Average float64
}

func main() {
	var s Univer
	var a Average

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("error1")
		return
	}
	// file, _ := os.Open("test.json")
	// data, _ := io.ReadAll(file)
	// fmt.Println("keep going")

	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)

	}
	studNumber := 0
	studSign := 0

	for _, el := range s.Students {
		studSign += len(el["Rating"].([]interface{}))
		studNumber++
	}
	a = Average{float64(studSign) / float64(studNumber)}
	outJSON, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", outJSON)
}
