package main

import "fmt"

func main() {

	var i, count int
	i = 1
	max := 1

	for i != 0 {
		fmt.Scan(&i)

		// un := iйфрнр
		//max := 1

		if i > max {
			max = i
			count = 1
		} else if i == max {
			count++
		}
		fmt.Println(count)
	}
	fmt.Println(count)
}
