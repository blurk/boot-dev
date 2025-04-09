package main

func getMessageCosts(messages []string) []float64 {
	messageCosts := make([]float64, len(messages))

	for i, message := range messages {
		messageCosts[i] = float64(len(message)) * 0.01
	}

	return messageCosts
}
