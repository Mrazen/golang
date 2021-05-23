package main

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	out := make(chan int)
	var sum int
	defer close(out)
	go func() {
	out:
		for {
			select {
			case arg := <-arguments:
				sum += arg
			case <-done:
				out <- sum
				break out
			}
		}
	}()
	return out
}

func main() {

}
