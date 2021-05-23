package main

import "fmt"

func main() {

	var n, count int

	fmt.Scan(&n)
	a := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	// fmt.Println(a)
	min := a[0]

	for _, v := range a {
		// fmt.Println(v)
		if v < min {
			min = v
			count = 1
		} else if v == min {
			count = count + 1
		}
	}
	fmt.Println(count)
}
