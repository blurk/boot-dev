package main

func getNameCounts(names []string) map[rune]map[string]int {
	nameCountsMap := make(map[rune]map[string]int)

	for _, name := range names {
		runes := []rune(name)
		firstChar := runes[0]

		if _, exists := nameCountsMap[firstChar]; !exists {
			nameCountsMap[firstChar] = make(map[string]int)
		}

		nameCountsMap[firstChar][name]++

	}
	return nameCountsMap
}
