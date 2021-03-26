package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	const data = "./stepik-go/work_with_files/task_sep_1/task.data"
	file, err := os.Open(data)
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()

	inf := bufio.NewReader(file)
	i := 1
	//fmt.Println(inf)
	for {
		line, err := inf.ReadString(';')
		if err == io.EOF {
			break
		}
		if line == "0;" {
			fmt.Println(i)
		}
		i++
		//fmt.Println(line)
	}

}
