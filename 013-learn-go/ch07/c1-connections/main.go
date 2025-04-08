package main

func countConnections(groupSize int) int {
	if groupSize <= 1 {
		return 0
	}

	count := 0

	for i := 2; i <= groupSize; i++ {
		count += i - 1
	}

	return count
}
