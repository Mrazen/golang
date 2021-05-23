package main

import "fmt"

func fibonacci() int {
	//print your code here
	var a, n, prev, next int
	fmt.Scan(&n)
	prev, next = 1, 1
	for a != n {
		prev, next = next, prev+next
		a++
	}
	fmt.Println(next)
	return next

}

func main() {

	fibonacci()

}
