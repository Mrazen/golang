package main

import "fmt"

func main() {

	var a, b, i, j, arepo, tenet, count int
	fmt.Scan(&a, &b) // 564 8954
	for a > 0 {
		i = a % 10 // 4 Отсекаем последнюю цифру
		a = a / 10 // 56 Удаляем последнюю цифру
		tenet *= 10
		tenet += i
		// fmt.Println(tenet)
	}

	for b > 0 {
		j = b % 10 //4
		b = b / 10 // 895
		arepo *= 10
		arepo += j
		// fmt.Println(arepo)

	}
	for tenet > 0 {
		i = tenet % 10     // 5
		tenet = tenet / 10 // 46
		count = arepo
		for arepo > 0 {
			j = arepo % 10     // 8 // 9 // 5  // 4
			arepo = arepo / 10 // 459 // 45 // 4 // 0
			if i == j {
				fmt.Print(i, " ")
			}

		}
		arepo = count
	}
}
