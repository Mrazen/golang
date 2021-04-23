package main

import "fmt"

func task2(c chan string, S string) {
	for i := 0; i < 5; i++ {
		c <- S + " "
	}

}

func removeDuplicates(inputStream chan string, outputStream chan string) {
	cache := ""
	defer close(outputStream)
	for data := range inputStream {
		if cache != data {
			outputStream <- data
		}
		cache = data

	}

}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	go func() {
		defer close(inputStream)

		for _, r := range "1122334456" {
			inputStream <- string(r)
		}
	}()
	for x := range outputStream {
		fmt.Print(x)
	}
}
