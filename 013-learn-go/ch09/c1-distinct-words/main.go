package main

import "strings"

func countDistinctWords(messages []string) int {
	wordMap := make(map[string]bool)

	for _, message := range messages {
		words := strings.Fields(message)

		for _, word := range words {
			lowerWord := strings.ToLower(word)
			wordMap[lowerWord] = true
		}
	}

	return len(wordMap)
}
