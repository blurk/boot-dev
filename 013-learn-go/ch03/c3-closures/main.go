package main

func adder() func(int) int {
	sum := 0

	return func(increment int) int {
		sum += increment
		return sum
	}
}
