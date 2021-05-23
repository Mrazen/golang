package main

import (
	"encoding/json"
	"fmt"
)

var (
	//buf = new(bytes.Buffer)

	dst okved
	//dst1 okved
)

func main() {
	strDecode := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

	if err := json.Unmarshal(getJSON(strDecode), &dst); err != nil {
		fmt.Print(err)
		return
	}
	sum := 0
	for _, valeu := range dst {
		sum += valeu.GlobalID
	}
	fmt.Print(sum)

}
