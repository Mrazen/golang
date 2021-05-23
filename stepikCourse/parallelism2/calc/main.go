package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		select {
		case arg := <-firstChan:
			out <- arg * arg
		case arg := <-secondChan:
			out <- arg * 3
		case <-stopChan:
			return
		default:
			return
		}
	}()
	return out
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})

	r := calculator(ch1, ch2, stop)
	ch1 <- 3
	//ch2 <- 3
	//close(stop)
	fmt.Println(<-r)

}
