package main

import "strings"

func removeProfanity(message *string) {
	bardWords := map[string]string{
		"fubb":  "****",
		"shiz":  "****",
		"witch": "*****",
	}

	for bardWord, replacement := range bardWords {
		*message = strings.ReplaceAll(*message, bardWord, replacement)
	}
}
