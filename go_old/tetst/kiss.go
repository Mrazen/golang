package main

import "fmt"

func main() {

	// объявляем переменные
	var N, a, first, second, k uint
	fmt.Scan(&N, &a)
	k = 1
	for N > 0 {
		first = N % 10
		N = N / 10
		if first == a {
			continue
		}
		second = first*k + second
		k *= 10
		// second += first
	}
	fmt.Println(second)
}
