package main

func maxMessages(thresh int) int {
	count := 0
	total := 0

	for fee := 0; ; fee++ {
		total += 100 + fee

		if total > thresh {
			return count
		} else {
			count++
		}
	}
}
