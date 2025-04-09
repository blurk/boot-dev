package main

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	a := []int{1, 2, 3, 1, 1}

	println(sum(a...))
}
