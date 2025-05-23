package main

import (
	"fmt"
	"math"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		costPerMessage float64
		numMessages    int
		expected       float64
	}
	runCases := []testCase{
		{2.55, 89, 226.95},
		{2.25, 204, 459},
	}
	submitCases := append(runCases, []testCase{
		{1, 1428, 1285.2},
		{5, 1000, 5000},
		{5, 1001, 4504.5},
		{3, 0, 0},
		{3, 7421, 17810.4},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := calculateFinalBill(test.costPerMessage, test.numMessages)
		output = math.Round(output*100) / 100
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%.2f, %d)
Expecting:  %.2f
Actual:     %.2f
Fail
`, test.costPerMessage, test.numMessages, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%.2f, %d)
Expecting:  %.2f
Actual:     %.2f
Pass
`, test.costPerMessage, test.numMessages, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
