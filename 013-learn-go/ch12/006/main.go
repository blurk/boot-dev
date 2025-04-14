package main

func concurrentFib(n int) []int {
	fibonacciChan := make(chan int)
	fibonacciSequence := make([]int, 0)

	go fibonacci(n, fibonacciChan)

	for num := range fibonacciChan {
		fibonacciSequence = append(fibonacciSequence, num)
	}

	return fibonacciSequence
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
