package main

import (
	"time"
)

func processMessages(messages []string) []string {
	numMessages := len(messages)
	ch := make(chan string, numMessages)

	for _, msg := range messages {
		go func(m string) {
			processed := process(m)
			ch <- processed
		}(msg)
	}

	result := make([]string, numMessages)
	for i := range numMessages {
		result[i] = <-ch
	}
	return result
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
