package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Univer struct {
	ID       int
	Number   string
	Year     int
	Students []map[string]interface{}
}

func main() {
	var s Univer
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("keep going")

	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)

	}
	for _, el := range s.Students {
		if v, ok := el["Rating"].([]interface{}); ok {
			fmt.Println(ok)
			fmt.Println(v)
		}
		fmt.Printf("%T", el["Rating"])
	}
	//fmt.Println(s.Students[0]["Rating"])

}
