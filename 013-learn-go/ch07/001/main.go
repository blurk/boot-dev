package main

func bulkSend(numMessages int) float64 {
	totalCost := 0.00

	for i := range numMessages {
		totalCost += 1.0 + (0.01 * float64(i))
	}

	return totalCost
}
