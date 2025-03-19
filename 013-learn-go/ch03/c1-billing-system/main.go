package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages)

	discount := baseBill * calculateDiscountRate(numMessages)
	finalBill := baseBill - discount

	return finalBill
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent <= 1000 {
		return 0
	}

	if messagesSent > 5000 {
		return 0.2
	}

	return 0.1
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
