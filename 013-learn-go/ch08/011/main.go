package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	result := []float64{}

	for _, cost := range costs {
		if cost.day == day {
			result = append(result, cost.value)
		}
	}

	return result
}
