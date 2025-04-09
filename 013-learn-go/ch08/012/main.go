package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for index, m := range msg {
		for _, badWord := range badWords {
			if m == badWord {
				return index
			}
		}
	}

	return -1
}
